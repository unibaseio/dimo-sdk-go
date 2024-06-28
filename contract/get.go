package contract

import (
	"context"
	"encoding/binary"
	"fmt"
	"math/big"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/contract/go/eproof"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/everify"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/gpu"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/model"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/node"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/piece"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/reward"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/rsproof"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/space"
	"github.com/MOSSV2/dimo-sdk-go/lib/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetInstAddr(ctx context.Context, typ string) (common.Address, error) {
	bi, err := NewBank(ctx)
	if err != nil {
		return common.Address{}, err
	}

	return bi.Get(&bind.CallOpts{From: Base}, typ)
}

func getOrder(count, pcnt uint64) (uint8, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()

	ri, err := NewEVerify(ctx)
	if err != nil {
		return 0, err
	}
	return ri.GetOrder(&bind.CallOpts{From: Base}, count, pcnt)
}

func GetOrder(count, pcnt uint64) (uint8, uint64) {
	if count == 0 {
		return 0, 0
	}
	total := uint64(8)
	order := uint8(1)
	for total < count {
		total *= 8
		order += 1
	}
	if order > 4 {
		total /= 8
		order -= 1
	}

	for order > 6 {
		total /= 8
		order -= 1
	}

	for total < count-pcnt {
		total *= 8
		order += 1
	}

	return order, total
}

func choose(addr common.Address, seed [32]byte, count, pcnt uint64, i uint64) (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	ri, err := NewEVerify(ctx)
	if err != nil {
		return 0, err
	}
	return ri.Choose(&bind.CallOpts{From: Base}, addr, seed, count, pcnt, i)
}

func Choose(addr common.Address, seed [32]byte, count, pcnt uint64, index uint64) uint64 {
	if pcnt+index < count {
		return pcnt + index
	}
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, index)
	if pcnt > 0 {
		buf = crypto.Keccak256(seed[:], addr.Bytes(), buf)
		res := new(big.Int).SetBytes(buf)
		res.Mod(res, big.NewInt(int64(pcnt)))
		return res.Uint64()
	}
	return pcnt
}

func GetEpochInfo(_epoch uint64) (*big.Int, [32]byte, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()

	ei, err := NewEpoch(ctx)
	if err != nil {
		return nil, [32]byte{}, err
	}

	return ei.GetEpoch(&bind.CallOpts{From: Base}, _epoch)
}

func CheckNode(addr common.Address, _typ uint8) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()

	ni, err := NewNode(ctx)
	if err != nil {
		return err
	}

	_, err = ni.Check(&bind.CallOpts{From: addr}, addr, _typ)
	return err
}

func GetPieceSerial(_pn string) (uint64, error) {
	pnb, err := G1StringInSolidity(_pn)
	if err != nil {
		return 0, err
	}
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	gi, err := NewPiece(ctx)
	if err != nil {
		return 0, err
	}
	return gi.GetPIndex(&bind.CallOpts{From: Base}, pnb)
}

func GetReplicaSerial(_pn string) (uint64, error) {
	pnb, err := G1StringInSolidity(_pn)
	if err != nil {
		return 0, err
	}
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	gi, err := NewPiece(ctx)
	if err != nil {
		return 0, err
	}
	return gi.GetRIndex(&bind.CallOpts{From: Base}, pnb)
}

func GetPiece(_pi uint64) (piece.IPiecePieceInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	fi, err := NewPiece(ctx)
	if err != nil {
		return piece.IPiecePieceInfo{}, err
	}
	pb, err := fi.GetPiece(&bind.CallOpts{From: Base}, _pi)
	if err != nil {
		return piece.IPiecePieceInfo{}, err
	}
	return pb, nil
}

func GetReplica(_ri uint64) (piece.IPieceReplicaInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	fi, err := NewPiece(ctx)
	if err != nil {
		return piece.IPieceReplicaInfo{}, err
	}
	pb, err := fi.GetReplica(&bind.CallOpts{From: Base}, _ri)
	if err != nil {
		return piece.IPieceReplicaInfo{}, err
	}

	return pb, nil
}

func GetMinPledge(_typ uint8) (*big.Int, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	ni, err := NewNode(ctx)
	if err != nil {
		return nil, err
	}
	return ni.GetMinPledge(&bind.CallOpts{From: Base}, _typ)
}

func GetPledgeInfo(addr common.Address, _typ uint8) (node.INodePledgeInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	ni, err := NewNode(ctx)
	if err != nil {
		return node.INodePledgeInfo{}, err
	}
	return ni.GetPledge(&bind.CallOpts{From: addr}, addr, _typ)
}

func GetStore(addr common.Address) (piece.IPieceStoreInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	fi, err := NewPiece(ctx)
	if err != nil {
		return piece.IPieceStoreInfo{}, err
	}
	fss, err := fi.GetStore(&bind.CallOpts{From: addr}, addr)
	if err != nil {
		return piece.IPieceStoreInfo{}, err
	}
	return fss, nil
}

func GetStoreStat(addr common.Address, _epoch uint64) (piece.IPieceStoreStat, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	fi, err := NewPiece(ctx)
	if err != nil {
		return piece.IPieceStoreStat{}, err
	}
	fss, err := fi.GetSStat(&bind.CallOpts{From: addr}, addr, _epoch)
	if err != nil {
		return piece.IPieceStoreStat{}, err
	}
	return fss, nil
}

func GetStoreReplica(_a common.Address, _ri uint64) (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	fi, err := NewPiece(ctx)
	if err != nil {
		return 0, err
	}
	return fi.GetSRAt(&bind.CallOpts{From: Base}, _a, _ri)
}

func GetRSChalInfo(_pi uint64, _pri uint8) (rsproof.RSProofProofInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	rsp, err := NewRSProof(ctx)
	if err != nil {
		return rsproof.RSProofProofInfo{}, err
	}

	return rsp.GetProof(&bind.CallOpts{From: Base}, _pi, _pri)
}

func GetRSMinTime() (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	rsp, err := NewRSProof(ctx)
	if err != nil {
		return 0, err
	}

	t, err := rsp.MinProveTime(&bind.CallOpts{From: Base})
	if err != nil {
		return 0, err
	}

	return t.Uint64(), nil
}

func GetEpochChalInfo(_a common.Address, _ep uint64) (eproof.IEproofProofInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	ep, err := NewEProof(ctx)
	if err != nil {
		return eproof.IEproofProofInfo{}, err
	}

	return ep.GetEProof(&bind.CallOpts{From: Base}, _a, _ep)
}

func GetEpochChalDetail(_a common.Address, _ep uint64) (everify.IEVerifyCInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	ep, err := NewEVerify(ctx)
	if err != nil {
		return everify.IEVerifyCInfo{}, err
	}

	return ep.GetCInfo(&bind.CallOpts{From: Base}, _a, _ep)
}

func GetEProofMinTime() (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	rsp, err := NewEProof(ctx)
	if err != nil {
		return 0, err
	}

	t, err := rsp.MinProveTime(&bind.CallOpts{From: Base})
	if err != nil {
		return 0, err
	}

	return t.Uint64(), nil
}

func GetReward(addr common.Address) (reward.IRewardRewardInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	gi, err := NewReward(ctx)
	if err != nil {
		return reward.IRewardRewardInfo{}, err
	}
	return gi.GetSReward(&bind.CallOpts{From: addr}, addr)

}

func GetRevenue(addr common.Address, typ string) (*big.Int, error) {
	res := big.NewInt(0)
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()

	switch typ {
	case types.StoreType:
		gi, err := NewPiece(ctx)
		if err != nil {
			return res, err
		}
		si, err := gi.GetStore(&bind.CallOpts{From: addr}, addr)
		if err != nil {
			return res, err
		}

		res.Set(si.Revenue)
		return res, nil
	case types.StreamType:
		gi, err := NewPiece(ctx)
		if err != nil {
			return res, err
		}
		val, err := gi.GetRevenue(&bind.CallOpts{From: addr}, addr)
		if err != nil {
			return res, err
		}

		res.Set(val)
		return res, nil
	case "reward":
		gi, err := NewReward(ctx)
		if err != nil {
			return res, err
		}
		si, err := gi.GetSReward(&bind.CallOpts{From: addr}, addr)
		if err != nil {
			return res, err
		}

		res.Set(si.Avail)
		return res, nil
	case "gpu":
		gi, err := NewGPU(ctx)
		if err != nil {
			return res, err
		}
		return gi.BalanceOf(&bind.CallOpts{From: addr}, addr)
	case "space":
		si, err := NewSpace(ctx)
		if err != nil {
			return res, err
		}
		return si.BalanceOf(&bind.CallOpts{From: addr}, addr)
	default:
		return res, fmt.Errorf("unsupported")
	}
}

func GetGPUIndex(_gn string) (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	gi, err := NewGPU(ctx)
	if err != nil {
		return 0, err
	}
	return gi.GetIndex(&bind.CallOpts{From: Base}, _gn)
}

func GetGPUInfo(_gi uint64) (gpu.IGPUInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	gi, err := NewGPU(ctx)
	if err != nil {
		return gpu.IGPUInfo{}, err
	}
	return gi.GetGPU(&bind.CallOpts{From: Base}, _gi)
}

func GetGPUOwner(_gi uint64) (common.Address, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	gi, err := NewGPU(ctx)
	if err != nil {
		return common.Address{}, err
	}
	return gi.GetOwner(&bind.CallOpts{From: Base}, _gi)
}

func GetGPUSetting() (*big.Int, [32]byte, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	gi, err := NewGPU(ctx)
	if err != nil {
		return nil, [32]byte{}, err
	}

	dif, err := gi.Difficulty(&bind.CallOpts{From: Base})
	if err != nil {
		return nil, [32]byte{}, err
	}

	seed, err := gi.Seed(&bind.CallOpts{From: Base})
	if err != nil {
		return nil, [32]byte{}, err
	}
	return dif, seed, nil
}

func GetModelIndex(_mn string) (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	mi, err := NewModel(ctx)
	if err != nil {
		return 0, err
	}
	return mi.GetIndex(&bind.CallOpts{From: Base}, _mn)
}

func GetModelInfo(_mi uint64) (model.IModelInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	mi, err := NewModel(ctx)
	if err != nil {
		return model.IModelInfo{}, err
	}
	return mi.GetModel(&bind.CallOpts{From: Base}, _mi)
}

func GetSpaceIndex(_mn string) (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	si, err := NewSpace(ctx)
	if err != nil {
		return 0, err
	}
	return si.GetIndex(&bind.CallOpts{From: Base}, _mn)
}

func GetSpaceInfo(_mi uint64) (space.ISpaceInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	si, err := NewSpace(ctx)
	if err != nil {
		return space.ISpaceInfo{}, err
	}
	return si.GetSpace(&bind.CallOpts{From: Base}, _mi)
}
