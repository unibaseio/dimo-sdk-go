package contract

import (
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/lib/bls"
	"github.com/MOSSV2/dimo-sdk-go/lib/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func Set(sk *ecdsa.PrivateKey, _typ string, ca common.Address) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	bi, err := NewBank(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	tx, err := bi.Set(au, _typ, ca)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}
	logger.Debug(_typ, " set success to: ", ca.String())
	return nil
}

func GetEpoch(sk *ecdsa.PrivateKey) (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	ei, err := NewEpoch(ctx)
	if err != nil {
		return 0, err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return 0, err
	}

	tx, err := ei.Check(au)
	if err != nil {
		return 0, err
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return 0, err
	}

	return ei.Current(&bind.CallOpts{From: au.From})
}

func RegisterNode(sk *ecdsa.PrivateKey, _typ uint8, val *big.Int) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()
	ni, err := NewNode(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	_, err = ni.Check(&bind.CallOpts{From: au.From}, au.From, _typ)
	if err == nil {
		logger.Debugf("%s already pledge enough money in type %d", au.From, _typ)
		return nil
	}

	ti, err := NewToken(ctx)
	if err != nil {
		return err
	}

	if val == nil {
		pval, err := ni.GetMinPledge(&bind.CallOpts{From: au.From}, _typ)
		if err != nil {
			return err
		}
		pinfo, err := ni.GetPledge(&bind.CallOpts{From: au.From}, au.From, _typ)
		if err != nil {
			return err
		}

		pinfo.Value.Sub(pinfo.Value, pinfo.Lock)

		if pval.Cmp(pinfo.Value) > 0 {
			pval.Sub(pval, pinfo.Value)
			val = pval
		} else {
			logger.Debug("no need more pledge")
			return nil
		}
	}

	tx, err := ti.IncreaseAllowance(au, BankAddr, val)
	if err != nil {
		return err
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}
	tx, err = ni.Pledge(au, _typ, val)
	if err != nil {
		return err
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}
	_, err = ni.Check(&bind.CallOpts{From: au.From}, au.From, _typ)
	return err
}

func AddPiece(sk *ecdsa.PrivateKey, pc types.PieceCore) error {
	logger.Debug("add piece: ", pc)
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()
	fi, err := NewPiece(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	ce, err := GetEpoch(sk)
	if err != nil {
		return err
	}

	if pc.Expire == 0 {
		pc.Start = ce
		pc.Expire = ce + uint64(DefaultStoreEpoch) + 1
	}
	if pc.Price == nil {
		pc.Price = big.NewInt(int64(DefaultReplicaPrice))
	}

	_size := 1 + (pc.Size-1)/(31*int64(pc.Policy.K))
	_size = 1 + (_size-1)/(32*1024)

	val := big.NewInt(int64(pc.Expire-pc.Start) * _size)
	val.Mul(val, pc.Price)
	val.Add(val, big.NewInt(int64(DefaultStreamPrice)))
	val.Mul(val, big.NewInt(int64(pc.Policy.N)))

	ti, err := NewToken(ctx)
	if err != nil {
		return err
	}

	tx, err := ti.IncreaseAllowance(au, BankAddr, val)
	if err != nil {
		return err
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	pb, err := G1StringInSolidity(pc.Name)
	if err != nil {
		return err
	}

	logger.Debug("add piece: ", pc)
	tx, err = fi.AddPiece(au, pb, pc.Price, uint64(pc.Size), pc.Expire, pc.Policy.N, pc.Policy.K, pc.Streamer)
	if err != nil {
		return err
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func AddReplica(sk *ecdsa.PrivateKey, rc types.ReplicaCore, pf []byte) error {
	logger.Debug("add replica: ", rc)
	rb, err := G1StringInSolidity(rc.Name)
	if err != nil {
		return err
	}
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()
	fi, err := NewPiece(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	pbyte, err := G1StringInSolidity(rc.Piece)
	if err != nil {
		return err
	}
	_pi, err := fi.GetPIndex(&bind.CallOpts{From: au.From}, pbyte)
	if err != nil {
		return err
	}

	if _pi == 0 {
		return fmt.Errorf("%s is not on chain", rc.Piece)
	}

	logger.Debug("add replica: ", _pi, rc)
	tx, err := fi.AddReplica(au, rb, _pi, rc.Index, pf)
	if err != nil {
		return err
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func UpdateStore(sk *ecdsa.PrivateKey, store common.Address) error {
	logger.Debug("update store: ", store)

	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	fi, err := NewPiece(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	tx, err := fi.CheckStore(au, store)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func ChallengeRS(sk *ecdsa.PrivateKey, _pn, _rn string, _pri uint8) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	ti, err := NewToken(ctx)
	if err != nil {
		return err
	}

	tx, err := ti.IncreaseAllowance(au, BankAddr, big.NewInt(int64(DefaultPenalty)))
	if err != nil {
		return err
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	rsp, err := NewRSProof(ctx)
	if err != nil {
		return err
	}
	pname, err := G1StringInSolidity(_pn)
	if err != nil {
		return err
	}

	rname, err := G1StringInSolidity(_rn)
	if err != nil {
		return err
	}

	logger.Debug("challenge rs proof: ", _rn, _pn, _pri)
	tx, err = rsp.Challenge(au, pname, rname, _pri)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func ProveRS(sk *ecdsa.PrivateKey, _pn, _rn string, _pri uint8, _pf []byte) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()
	rsp, err := NewRSProof(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	pname, err := G1StringInSolidity(_pn)
	if err != nil {
		return err
	}

	rname, err := G1StringInSolidity(_rn)
	if err != nil {
		return err
	}

	tx, err := rsp.Prove(au, pname, rname, _pri, _pf)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func CheckRSChallenge(sk *ecdsa.PrivateKey, _pn, _rn string, _pri uint8) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	rsp, err := NewRSProof(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	pname, err := G1StringInSolidity(_pn)
	if err != nil {
		return err
	}

	rname, err := G1StringInSolidity(_rn)
	if err != nil {
		return err
	}

	piece, err := NewPiece(ctx)
	if err != nil {
		return err
	}

	_pi, err := piece.GetPIndex(&bind.CallOpts{From: au.From}, pname)
	if err != nil {
		return err
	}

	_ri, err := piece.GetRIndex(&bind.CallOpts{From: au.From}, rname)
	if err != nil {
		return err
	}

	tx, err := rsp.Check(au, _pi, _ri, _pri)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func SubmitProof(sk *ecdsa.PrivateKey, _ep uint64, _pf bls.EpochProof) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	pi, err := NewEProof(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	_sum := G1InSolidity(_pf.Sum)
	_pfb := G1InSolidity(_pf.H)
	_frb := FrInSolidity(_pf.ClaimedValue)
	_pfb = append(_pfb, _frb...)

	logger.Debug("submit epoch proof: ", au.From, _ep)
	tx, err := pi.Submit(au, _ep, _sum, _pfb)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func ChallengeKZG(sk *ecdsa.PrivateKey, addr common.Address, _ep uint64) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	ti, err := NewToken(ctx)
	if err != nil {
		return err
	}

	tx, err := ti.IncreaseAllowance(au, BankAddr, big.NewInt(int64(DefaultPenalty)))
	if err != nil {
		return err
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	pi, err := NewEProof(ctx)
	if err != nil {
		return err
	}

	logger.Debug("challenge eproof: ", addr, _ep)
	tx, err = pi.ChalKZG(au, addr, _ep)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func ProveKZG(sk *ecdsa.PrivateKey, _ep uint64, _wroot []byte, _pf []byte) error {
	if len(_wroot) != 32 {
		return fmt.Errorf("invalid witness root length")
	}

	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()
	pi, err := NewEProof(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	var _wt [32]byte
	copy(_wt[:], _wroot)

	tx, err := pi.ProveKZG(au, _ep, _wt, _pf)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func ChallengeSum(sk *ecdsa.PrivateKey, addr common.Address, _ep uint64, _qIndex uint8, sum string) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}
	ti, err := NewToken(ctx)
	if err != nil {
		return err
	}

	if len(sum) > 0 {
		tx, err := ti.IncreaseAllowance(au, BankAddr, big.NewInt(int64(DefaultPenalty)))
		if err != nil {
			return err
		}
		err = CheckTx(DevChain, tx.Hash())
		if err != nil {
			return err
		}
	}

	pi, err := NewEProof(ctx)
	if err != nil {
		return err
	}

	if len(sum) > 0 {
		_sum, err := G1StringInSolidity(sum)
		if err != nil {
			return err
		}
		logger.Debug("challenge eproof sum0: ", addr, _ep)
		tx, err := pi.Challenge(au, addr, _ep, _sum)
		if err != nil {
			return err
		}
		err = CheckTx(DevChain, tx.Hash())
		if err != nil {
			return err
		}
	} else {
		logger.Debug("challenge eproof sum: ", addr, _ep, _qIndex)
		tx, err := pi.ChalCom(au, addr, _ep, _qIndex)
		if err != nil {
			return err
		}

		err = CheckTx(DevChain, tx.Hash())
		if err != nil {
			return err
		}
	}

	return nil
}

func ProveSum(sk *ecdsa.PrivateKey, _ep uint64, coms []bls.G1, _pf []byte) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	ti, err := NewToken(ctx)
	if err != nil {
		return err
	}

	tx, err := ti.IncreaseAllowance(au, BankAddr, big.NewInt(int64(DefaultPenalty)))
	if err != nil {
		return err
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	pi, err := NewEProof(ctx)
	if err != nil {
		return err
	}

	_coms := make([][]byte, len(coms))
	for i := 0; i < len(coms); i++ {
		_coms[i] = G1InSolidity(coms[i])
	}

	logger.Debug("prove eproof sum: ", au.From, _ep)
	tx, err = pi.ProveCom(au, _ep, _coms, _pf)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func ChallengeOne(sk *ecdsa.PrivateKey, addr common.Address, _ep uint64, _qIndex uint8) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	pi, err := NewEProof(ctx)
	if err != nil {
		return err
	}

	logger.Debug("challenge eproof one: ", addr, _ep, _qIndex)
	tx, err := pi.ChalOne(au, addr, _ep, _qIndex)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func ProveOne(sk *ecdsa.PrivateKey, _ep uint64, com bls.G1, _pf []byte) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	pi, err := NewEProof(ctx)
	if err != nil {
		return err
	}

	_com := G1InSolidity(com)
	logger.Debug("prove eproof one: ", au.From, _ep)
	tx, err := pi.ProveOne(au, _ep, _com, _pf)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func CheckEpochChallenge(sk *ecdsa.PrivateKey, addr common.Address, _ep uint64) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	ep, err := NewEProof(ctx)
	if err != nil {
		return err
	}

	logger.Debug("check eproof: ", addr, _ep)
	tx, err := ep.Check(au, addr, _ep)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func TestProveRS(rsn, rsk uint8, pub []*big.Int, _pf []byte) error {
	if len(pub) != 3 {
		return fmt.Errorf("invalid public length")
	}
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	rsp, err := NewRSProof(ctx)
	if err != nil {
		return err
	}

	vt, err := rsp.GetVKRoot(&bind.CallOpts{From: Base}, rsn, rsk)
	if err != nil {
		return err
	}

	if vt.Cmp(pub[0]) != 0 {
		return fmt.Errorf("unequal vkroot")
	}

	rsv, err := NewRSOne(ctx)
	if err != nil {
		return err
	}

	ok, err := rsv.Verify(&bind.CallOpts{From: Base}, _pf, pub)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("invalid rs proof")
	}

	return nil
}

func TestProveEpoch(_key string, pub []*big.Int, _pf []byte) error {
	if len(pub) != 2 {
		return fmt.Errorf("invalid public length")
	}
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	ev, err := NewEVerify(ctx)
	if err != nil {
		return err
	}

	vt := new(big.Int)
	switch _key {
	case "kzg":
		vt, err = ev.INKZGVK(&bind.CallOpts{From: Base})
	case "mul":
		vt, err = ev.INMULVK(&bind.CallOpts{From: Base})
	case "add":
		vt, err = ev.INADDVK(&bind.CallOpts{From: Base})
	default:
		return fmt.Errorf("unsupported inner circuit: %s", _key)
	}
	if err != nil {
		return err
	}

	if vt.Cmp(pub[0]) != 0 {
		return fmt.Errorf("unequal vkroot")
	}

	switch _key {
	case "kzg":
		pv, err := NewKZGPlonk(ctx)
		if err != nil {
			return err
		}
		ok, err := pv.Verify(&bind.CallOpts{From: Base}, _pf, pub)
		if err != nil {
			return err
		}
		if !ok {
			return fmt.Errorf("invalid rs proof")
		}
	case "add":
		pv, err := NewAddPlonk(ctx)
		if err != nil {
			return err
		}
		ok, err := pv.Verify(&bind.CallOpts{From: Base}, _pf, pub)
		if err != nil {
			return err
		}
		if !ok {
			return fmt.Errorf("invalid rs proof")
		}
	case "mul":
		pv, err := NewMulPlonk(ctx)
		if err != nil {
			return err
		}
		ok, err := pv.Verify(&bind.CallOpts{From: Base}, _pf, pub)
		if err != nil {
			return err
		}
		if !ok {
			return fmt.Errorf("invalid rs proof")
		}
	default:
		return fmt.Errorf("unsupported key")
	}

	return nil
}

func Settle(sk *ecdsa.PrivateKey, _money *big.Int) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	fi, err := NewPiece(ctx)
	if err != nil {
		return err
	}
	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	tx, err := fi.Settle(au, _money)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func WithdrawRevenue(sk *ecdsa.PrivateKey, _money *big.Int) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	fi, err := NewPiece(ctx)
	if err != nil {
		return err
	}
	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	tx, err := fi.Withdraw(au, _money)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func SettleReward(sk *ecdsa.PrivateKey, addr common.Address, _ep uint64) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	fi, err := NewReward(ctx)
	if err != nil {
		return err
	}
	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	tx, err := fi.Settle(au, addr, _ep)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func WithdrawReward(sk *ecdsa.PrivateKey, _money *big.Int) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	fi, err := NewReward(ctx)
	if err != nil {
		return err
	}
	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	tx, err := fi.Withdraw(au, _money)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func AddModel(sk *ecdsa.PrivateKey, mc types.ModelMeta) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	mi, err := NewModel(ctx)
	if err != nil {
		return err
	}
	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	rt, err := hex.DecodeString(mc.Hash)
	if err != nil {
		return err
	}

	var cnt [32]byte
	binary.BigEndian.PutUint64(cnt[24:32], uint64(mc.Count))

	tx, err := mi.Add(au, mc.Name, rt, cnt)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func AddGPU(sk *ecdsa.PrivateKey, _gn string) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	gi, err := NewGPU(ctx)
	if err != nil {
		return err
	}
	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	tx, err := gi.Add(au, _gn)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func AddSpace(sk *ecdsa.PrivateKey, msm types.SpaceMeta) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()
	si, err := NewSpace(ctx)
	if err != nil {
		return err
	}

	gi, err := NewGPU(ctx)
	if err != nil {
		return err
	}

	mi, err := NewModel(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	_mi, err := mi.GetIndex(&bind.CallOpts{From: au.From}, msm.Model)
	if err != nil {
		return err
	}

	_gi, err := gi.GetIndex(&bind.CallOpts{From: au.From}, msm.GPU)
	if err != nil {
		return err
	}

	ce, err := GetEpoch(sk)
	if err != nil {
		return err
	}

	val := big.NewInt(int64(DefaultSpacePrice))
	val.Mul(val, big.NewInt(int64(DefaultSpaceEpoch)))

	ti, err := NewToken(ctx)
	if err != nil {
		return err
	}

	tx, err := ti.IncreaseAllowance(au, BankAddr, val)
	if err != nil {
		return err
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	tx, err = si.Add(au, msm.Name, _mi, _gi, big.NewInt(int64(DefaultSpacePrice)), ce+uint64(DefaultSpaceEpoch))
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func ActivateSpace(sk *ecdsa.PrivateKey, sn, root string, pfbyte []byte) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()
	si, err := NewSpace(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	_ai, err := si.GetIndex(&bind.CallOpts{From: au.From}, sn)
	if err != nil {
		return err
	}

	_rt, err := hex.DecodeString(root)
	if err != nil {
		return err
	}

	tx, err := si.Activate(au, _ai, _rt, pfbyte)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func ShutdownSpace(sk *ecdsa.PrivateKey, _ai uint64) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()
	si, err := NewSpace(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	tx, err := si.Shutdown(au, _ai)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func GPUCheck(sk *ecdsa.PrivateKey) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()
	gi, err := NewGPU(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	tx, err := gi.Check(au)
	if err != nil {
		return err
	}

	return CheckTx(DevChain, tx.Hash())
}

func GPUMint(sk *ecdsa.PrivateKey, _gi uint64, _salt []byte) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()
	gi, err := NewGPU(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	tx, err := gi.Mint(au, _gi, _salt)
	if err != nil {
		return err
	}

	return CheckTx(DevChain, tx.Hash())
}
