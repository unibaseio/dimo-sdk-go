package contract

import (
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/contract/go/proof"
	"github.com/MOSSV2/dimo-sdk-go/lib/log"
	"github.com/MOSSV2/dimo-sdk-go/lib/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var logger = log.Logger("contract")

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
	fmt.Println(_typ, " set success to: ", ca.String())
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

func RegisterNode(sk *ecdsa.PrivateKey, _typ uint8) error {
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

	err = ni.Check(&bind.CallOpts{From: au.From}, au.From, _typ)
	if err == nil {
		logger.Debug("already pledge")
		return nil
	}

	ti, err := NewToken(ctx)
	if err != nil {
		return err
	}

	pval, err := ni.GetMinPledge(&bind.CallOpts{From: au.From}, _typ)
	if err != nil {
		return err
	}
	pinfo, err := ni.GetPledge(&bind.CallOpts{From: au.From}, au.From, _typ)
	if err != nil {
		return err
	}
	pval.Sub(pval, pinfo.Value)

	tx, err := ti.IncreaseAllowance(au, BankAddr, pval)
	if err != nil {
		return err
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}
	tx, err = ni.Pledge(au, _typ, pval)
	if err != nil {
		return err
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}
	return ni.Check(&bind.CallOpts{From: au.From}, au.From, _typ)
}

func AddFileAndPiece(sk *ecdsa.PrivateKey, _fn string, fc types.FileCore, _proxy common.Address) error {
	_fi, err := AddFile(sk, _fn, fc)
	if err != nil {
		return err
	}

	err = AddPiece(sk, _fi, fc.Pieces)
	if err != nil {
		return err
	}

	if _proxy != common.HexToAddress("0x0") {
		err = SetFileAgent(sk, _fi, _proxy)
		if err != nil {
			return err
		}
		return PrePayFile(sk, _fi, int(fc.Policy.N))
	}

	return nil
}

func AddFile(sk *ecdsa.PrivateKey, _fn string, fc types.FileCore) (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	fi, err := NewFile(ctx)
	if err != nil {
		return 0, err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return 0, err
	}

	ce, err := GetEpoch(sk)
	if err != nil {
		return 0, err
	}

	logger.Debug("add file: ", _fn)
	_, err = fi.GetFileIndex(&bind.CallOpts{From: au.From}, _fn)
	if err == nil {
		logger.Debug("already has: %s", _fn)
		return 0, fmt.Errorf("already has: %s", _fn)
	}

	tx, err := fi.AddFile(au, _fn, big.NewInt(int64(DefaultReplicaPrice)), ce+uint64(DefaultStoreEpoch))
	if err != nil {
		return 0, err
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return 0, err
	}

	return fi.GetFileIndex(&bind.CallOpts{From: au.From}, _fn)
}

func AddPiece(sk *ecdsa.PrivateKey, _fi uint64, pieces []string) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	fi, err := NewFile(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	piecebytes := make([][]byte, 0, len(pieces))
	for i := 0; i < len(pieces); i++ {
		pb, err := StringToSolByte(pieces[i])
		if err != nil {
			return err
		}
		piecebytes = append(piecebytes, pb)
	}

	for i := 0; i < len(pieces); i += 128 {
		end := i + 128
		if end > len(pieces) {
			end = len(pieces)
		}
		logger.Debug("add piece to: ", _fi)
		tx, err := fi.AddPiece(au, _fi, piecebytes[i:end])
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

func SetFileAgent(sk *ecdsa.PrivateKey, _fi uint64, _proxy common.Address) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	fi, err := NewFile(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	logger.Debug(_fi, " set agent to: ", _proxy)
	tx, err := fi.SetAgent(au, _fi, _proxy)
	if err != nil {
		return err
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func PrePayFile(sk *ecdsa.PrivateKey, _fi uint64, _n int) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	logger.Debug("allowrance to: ", _fi)
	fb, err := GetFile(_fi)
	if err != nil {
		return err
	}

	val := new(big.Int).Mul(fb.Price, big.NewInt(int64(fb.Count)*int64(_n)))
	val.Mul(val, big.NewInt(int64(fb.Expire-fb.Start)+1))

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

	logger.Debug("prepay to: ", _fi)
	fi, err := NewFile(ctx)
	if err != nil {
		return err
	}

	tx, err = fi.PrePay(au, _fi, val)
	if err != nil {
		return err
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func AddReplica(sk *ecdsa.PrivateKey, _fi uint64, _pi uint64, replica string, _sign []byte) error {
	rb, err := StringToSolByte(replica)
	if err != nil {
		return err
	}
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Minute)
	defer cancle()
	fi, err := NewFile(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	fb, err := fi.GetFile(&bind.CallOpts{From: au.From}, _fi)
	if err != nil {
		return err
	}

	logger.Debug("add replica: ", replica)
	if fb.Owner == au.From {
		ce, err := GetEpoch(sk)
		if err != nil {
			return err
		}

		val := big.NewInt(int64(fb.Expire - ce))
		val.Mul(val, fb.Price)

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
	}

	tx, err := fi.AddReplica(au, _fi, _pi, rb, _sign)
	if err != nil {
		return err
	}
	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func SubmitProof(sk *ecdsa.PrivateKey, _kp proof.IProofKZGProof) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	pi, err := NewProof(ctx)
	if err != nil {
		return err
	}

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	tx, err := pi.Submit(au, _kp)
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func Settle(sk *ecdsa.PrivateKey, _epoch uint64) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	fi, err := NewFile(ctx)
	if err != nil {
		return err
	}
	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	tx, err := fi.Settle(au, _epoch)
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

	val := big.NewInt(int64(DefaultGPUPrice))
	val.Mul(val, big.NewInt(int64(DefaultGPUEpoch)))

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

	tx, err = si.Add(au, msm.Name, _mi, _gi, big.NewInt(int64(DefaultGPUPrice)), ce+uint64(DefaultGPUEpoch))
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

func Mint(sk *ecdsa.PrivateKey) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	gi, err := NewControl(ctx)
	if err != nil {
		return err
	}

	baddr, err := gi.Bank(&bind.CallOpts{From: Base})
	if err != nil {
		return err
	}

	fmt.Println(baddr)

	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		return err
	}

	tx, err := gi.Mint(au, au.From, big.NewInt(100))
	if err != nil {
		return err
	}

	err = CheckTx(DevChain, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}
