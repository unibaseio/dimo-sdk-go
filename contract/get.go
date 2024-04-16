package contract

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/contract/go/file"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/gpu"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/model"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/proof"
	"github.com/MOSSV2/dimo-sdk-go/contract/go/space"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func GetInstAddr(ctx context.Context, typ string) (common.Address, error) {
	bi, err := NewBank(ctx)
	if err != nil {
		return common.Address{}, err
	}

	return bi.Get(&bind.CallOpts{From: Base}, typ)
}

func GetOrder(count uint64) (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancle()

	ri, err := NewRound(ctx)
	if err != nil {
		return 0, err
	}
	return ri.GetOrder(&bind.CallOpts{From: Base}, count)
}

func Choose(addr common.Address, seed [32]byte, count uint64, i uint64) (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancle()
	ri, err := NewRound(ctx)
	if err != nil {
		return 0, err
	}
	return ri.Choose(&bind.CallOpts{From: Base}, addr, seed, count, i)
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

	return ni.Check(&bind.CallOpts{From: addr}, addr, _typ)
}

func GetFileIndex(_fn string) (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	fi, err := NewFile(ctx)
	if err != nil {
		return 0, err
	}
	_fi, err := fi.GetFileIndex(&bind.CallOpts{From: Base}, _fn)
	if err != nil {
		return 0, err
	}
	return _fi, nil
}

func IsFileAgent(_fi uint64, addr common.Address) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	fi, err := NewFile(ctx)
	if err != nil {
		return err
	}
	is, err := fi.IsAgent(&bind.CallOpts{From: Base}, _fi, addr)
	if err != nil {
		return err
	}
	if !is {
		return fmt.Errorf("not")
	}
	return nil
}

func GetFile(_fi uint64) (file.IFileFileInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	fi, err := NewFile(ctx)
	if err != nil {
		return file.IFileFileInfo{}, err
	}
	fb, err := fi.GetFile(&bind.CallOpts{From: Base}, _fi)
	if err != nil {
		return file.IFileFileInfo{}, err
	}
	return fb, nil
}

func GetPiece(_fi uint64, _pi uint64) (file.IFilePieceInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	fi, err := NewFile(ctx)
	if err != nil {
		return file.IFilePieceInfo{}, err
	}
	pb, err := fi.GetPiece(&bind.CallOpts{From: Base}, _fi, _pi)
	if err != nil {
		return file.IFilePieceInfo{}, err
	}
	return pb, nil
}

func GetReplica(_ri uint64) (file.IFileReplicaInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	fi, err := NewFile(ctx)
	if err != nil {
		return file.IFileReplicaInfo{}, err
	}
	pb, err := fi.GetReplica(&bind.CallOpts{From: Base}, _ri)
	if err != nil {
		return file.IFileReplicaInfo{}, err
	}

	return pb, nil
}

func GetStore(addr common.Address) (file.IFileStoreInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	fi, err := NewFile(ctx)
	if err != nil {
		return file.IFileStoreInfo{}, err
	}
	fss, err := fi.GetStore(&bind.CallOpts{From: addr}, addr)
	if err != nil {
		return file.IFileStoreInfo{}, err
	}
	return fss, nil
}

func GetStoreStat(addr common.Address, _epoch uint64) (file.IFileStoreStat, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	fi, err := NewFile(ctx)
	if err != nil {
		return file.IFileStoreStat{}, err
	}
	fss, err := fi.GetSStat(&bind.CallOpts{From: addr}, addr, _epoch)
	if err != nil {
		return file.IFileStoreStat{}, err
	}
	if fss.Count == 0 {
		return file.IFileStoreStat{}, fmt.Errorf("no storage")
	}
	return fss, nil
}

func GetStoreReplica(_a common.Address, _ri uint64) (string, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	fi, err := NewFile(ctx)
	if err != nil {
		return "", err
	}
	pb, err := fi.GetSReplica(&bind.CallOpts{From: Base}, _a, _ri)
	if err != nil {
		return "", err
	}

	return SolByteToString(pb)
}

func GetEpochProof(addr common.Address, _epoch uint64) (proof.IProofEProof, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	pi, err := NewProof(ctx)
	if err != nil {
		return proof.IProofEProof{}, err
	}
	return pi.GetEProof(&bind.CallOpts{From: addr}, addr, _epoch)
}

func GetRevenue(addr common.Address, typ string) (*big.Int, error) {
	res := big.NewInt(0)
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()

	switch typ {
	case "file":
		fi, err := NewFile(ctx)
		if err != nil {
			return res, err
		}
		return fi.BalanceOf(&bind.CallOpts{From: addr}, addr)
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
