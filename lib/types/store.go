package types

import (
	"bytes"
	"reflect"
	"strconv"
)

type DsType = byte

const (
	DsAccount DsType = 100 + iota
	DsParam
	DsStat
	DsProxy
	DsBalance
	DsEdge
	DsFile
	DsPiece
	DsReplica
	DsEpoch
	DsRSProof
	DSEProof
	DsModel
	DsSpace
	DsGPU
	DsRepo
	DsBucket
	DsObject
	DsNeedle
)

const (
	// KeyDelimiter seps kv key
	KeyDelimiter = "/"
)

// key is human readable
// type/key/...
func NewKey(vs ...interface{}) []byte {
	var b bytes.Buffer
	vlen := len(vs)
	for i, v := range vs {
		typ := reflect.TypeOf(v).Kind()
		switch typ {
		case reflect.Slice:
			val := v.([]byte)
			b.Write(val)
		case reflect.String:
			val := v.(string)
			b.Write([]byte(val))
		case reflect.Int8:
			val := v.(int8)
			b.Write([]byte(strconv.FormatInt(int64(val), 10)))
		case reflect.Uint8:
			val := v.(uint8)
			b.Write([]byte(strconv.FormatInt(int64(val), 10)))
		case reflect.Int:
			val := v.(int)
			b.Write([]byte(strconv.FormatInt(int64(val), 10)))
		case reflect.Int32:
			val := v.(int32)
			b.Write([]byte(strconv.FormatInt(int64(val), 10)))
		case reflect.Int64:
			val := v.(int64)
			b.Write([]byte(strconv.FormatInt(val, 10)))
		case reflect.Uint:
			val := v.(uint)
			b.Write([]byte(strconv.FormatUint(uint64(val), 10)))
		case reflect.Uint32:
			val := v.(uint32)
			b.Write([]byte(strconv.FormatUint(uint64(val), 10)))
		case reflect.Uint64:
			val := v.(uint64)
			b.Write([]byte(strconv.FormatUint(val, 10)))
		default:
			continue
		}

		if i < vlen-1 {
			b.Write([]byte(KeyDelimiter))
		}
	}

	return b.Bytes()
}

type DiskStats struct {
	Path  string `json:"path"`
	Total uint64 `json:"all"`
	Used  uint64 `json:"used"`
	Free  uint64 `json:"free"`
}
type IStore interface {
	Put(key, value []byte) error
	Get(key []byte) ([]byte, error)
	Has(key []byte) (bool, error)
	Delete(key []byte) error
	Size() DiskStats
	Close() error
}

type IKVStore interface {
	IStore

	GetNext(key []byte, bandwidth int) (uint64, error)
	Iter(prefix []byte, fn func(k, v []byte) error) int64
	IterKeys(prefix []byte, fn func(k []byte) error) int64

	Sync() error

	NewTxnStore(bool) (ITxnStore, error)
}

type ITxnStore interface {
	IStore
	Commit() error
	Discard()
}

type IFileStore interface {
	Put(key, value []byte) error
	Get(key []byte, opts ...int) ([]byte, error)
	Has(key []byte) (bool, error)
	Delete(key []byte) error
	Size() DiskStats
	Close() error
}
