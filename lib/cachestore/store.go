package cacheStore

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
	"sync"

	"github.com/MOSSV2/dimo-sdk-go/lib/bls"
	"github.com/MOSSV2/dimo-sdk-go/lib/lighthash"
	"github.com/MOSSV2/dimo-sdk-go/lib/log"
	"github.com/MOSSV2/dimo-sdk-go/lib/merkle"
	"github.com/MOSSV2/dimo-sdk-go/lib/types"

	"github.com/ethereum/go-ethereum/common"
)

var logger = log.Logger("cache")

const (
	hashSize = 32
	subShard = 1024
)

var _ types.IReplicaStore = &CacheStore{}

type CacheStore struct {
	sync.RWMutex
	typ   string
	local common.Address
	ds    types.IKVStore
	fs    types.IFileStore
}

func New(local common.Address, ds types.IKVStore, fs types.IFileStore) (*CacheStore, error) {
	logger.Info("create cache store")

	ks := &CacheStore{
		typ:   "cache",
		local: local,
		ds:    ds,
		fs:    fs,
	}

	return ks, nil
}

func (ks *CacheStore) Put(ctx context.Context, pc types.PieceCore, d []byte) (types.ReplicaCore, error) {
	logger.Debug(ks.typ, " store Put: ", pc)

	rc := types.ReplicaCore{
		Piece: pc.Name,
		Name:  pc.Name,
	}

	if int64(len(d)) != pc.Size {
		return rc, fmt.Errorf("com file size mismatch")
	}

	if len(d) > bls.MaxSize {
		return rc, fmt.Errorf("com file size too large")
	}

	metaKey := types.NewKey(types.DsReplica, rc.Name)
	has, err := ks.ds.Has(metaKey)
	if err == nil && has {
		return rc, nil
	}

	d = bls.Pad(d)
	root, aux, err := merkle.ReaderRootAux(bytes.NewReader(d), lighthash.New(), bls.PadSize, subShard)
	if err != nil {
		return rc, err
	}

	comhex := hex.EncodeToString(root)
	if pc.Name != comhex {
		return rc, fmt.Errorf("com file content mismatch, got %s need %s", comhex, pc.Name)
	}

	rc.Size = int64(len(d))
	rc.StoredOn = ks.local

	// save data
	dskey := types.NewKey(types.DsReplica, rc.Name)
	err = ks.fs.Put(dskey, d)
	if err != nil {
		return rc, err
	}

	if len(aux) > 0 {
		dskey = types.NewKey(types.DsReplica, types.DsReplica, rc.Name)
		err = ks.ds.Put(dskey, aux)
		if err != nil {
			return rc, err
		}
	}

	csb, err := rc.Serialize()
	if err != nil {
		return rc, err
	}
	metaKey = types.NewKey(types.DsReplica, rc.Name)
	err = ks.ds.Put(metaKey, csb)
	if err != nil {
		return rc, err
	}

	return rc, nil
}

func (ks *CacheStore) Get(ctx context.Context, name string, w io.Writer, opt types.Options) (types.ReplicaCore, error) {
	logger.Debug(ks.typ, " store Get: ", name)
	dskey := types.NewKey(types.DsReplica, name)
	cr := new(types.ReplicaCore)
	val, err := ks.ds.Get(dskey)
	if err != nil {
		return types.ReplicaCore{}, fmt.Errorf("no such replica %s", name)
	}
	err = cr.Deserialize(val)
	if err != nil {
		return types.ReplicaCore{}, err
	}

	if w != nil {
		dskey := types.NewKey(types.DsReplica, cr.Name)
		data, err := ks.fs.Get(dskey)
		if err != nil {
			return *cr, err
		}
		if opt.UserDefined != nil {
			if opt.UserDefined["unpad"] != "" {
				usize, err := strconv.Atoi(opt.UserDefined["unpad"])
				if err == nil {
					data, err = bls.Unpad(data, usize)
					if err != nil {
						return types.ReplicaCore{}, err
					}
				}
			}
			start := 0
			size := len(data)
			if opt.UserDefined["start"] != "" {
				start, _ = strconv.Atoi(opt.UserDefined["start"])
			}
			if opt.UserDefined["size"] != "" {
				size, _ = strconv.Atoi(opt.UserDefined["size"])
			}
			data = data[start : start+size]
		}

		_, err = w.Write(data)
		return *cr, err
	}

	return *cr, nil
}

func (ks *CacheStore) GenProof(ctx context.Context, name string, rnd []byte, typ int) ([]byte, error) {
	return nil, fmt.Errorf("unsupported method: GenProof")
}

func (ks *CacheStore) List(ctx context.Context, addr common.Address, opt types.Options) ([]types.ReplicaCore, error) {
	return nil, fmt.Errorf("unsupported method: list")
}
