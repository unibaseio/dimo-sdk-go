package types

import (
	"context"
	"fmt"
	"io"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fxamacker/cbor/v2"
)

type Policy struct {
	N uint8
	K uint8
}

func (p Policy) Check() error {
	switch {
	case p.N == 6 && p.K == 4:
	case p.N == 14 && p.K == 7:
	case p.N == 32 && p.K == 16:
	case p.N == 64 && p.K == 32:
	default:
		return fmt.Errorf("unsupported rs policy: %d %d", p.N, p.K)
	}
	return nil
}

type FileCore struct {
	Policy   Policy
	Name     string
	Hash     string
	Size     int64
	Owner    common.Address
	Creation time.Time
	Append   bool
}

func (frc *FileCore) Serialize() ([]byte, error) {
	return cbor.Marshal(frc)
}

func (frc *FileCore) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, frc)
}

type FileReceipt struct {
	FileCore
	Pieces []string
}

func (fr *FileReceipt) Serialize() ([]byte, error) {
	return cbor.Marshal(fr)
}

func (fr *FileReceipt) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, fr)
}

type FileFull struct {
	FileReceipt
	Proofs     [][]byte
	PieceSizes []int64
}

func (ff *FileFull) Serialize() ([]byte, error) {
	return cbor.Marshal(ff)
}

func (ff *FileFull) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, ff)
}

type PieceCore struct {
	Policy   Policy
	Name     string
	Serial   uint64
	Size     int64 // raw size
	Start    uint64
	Expire   uint64
	Price    *big.Int
	Owner    common.Address
	Streamer common.Address
}

func (pc *PieceCore) Serialize() ([]byte, error) {
	return cbor.Marshal(pc)
}

func (pc *PieceCore) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, pc)
}

type PieceReceipt struct {
	PieceCore
	Replicas []string
	StoredOn []common.Address
}

func (cr *PieceReceipt) Serialize() ([]byte, error) {
	return cbor.Marshal(cr)
}

func (cr *PieceReceipt) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, cr)
}

type PieceWitness struct {
	Choose  uint8
	Witness []byte
}

type ReplicaCore struct {
	Fake     bool
	Name     string // encoded
	Serial   uint64
	Size     int64  // stored size
	Piece    string // belongs to which piece
	Index    uint8
	StoredOn common.Address
}

func (rc *ReplicaCore) Serialize() ([]byte, error) {
	return cbor.Marshal(rc)
}

func (rc *ReplicaCore) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, rc)
}

type ReplicaWitness struct {
	Index uint64
	Proof []byte
}

func (rw *ReplicaWitness) Serialize() ([]byte, error) {
	return cbor.Marshal(rw)
}

func (rw *ReplicaWitness) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, rw)
}

type IFile interface {
	AddFile(context.Context, FileReceipt) error
	GetFile(context.Context, string, Options) (FileReceipt, error)
	GetPiece(context.Context, string, Options) (PieceReceipt, error)
	GetReplica(context.Context, string, io.Writer, Options) (ReplicaCore, error)

	ListFile(context.Context, common.Address, Options) ([]FileReceipt, error)
	ListPiece(context.Context, common.Address, Options) ([]PieceReceipt, error)
	ListReplica(context.Context, common.Address, Options) ([]ReplicaCore, error)
}

type IPieceStore interface {
	PutPiece(context.Context, PieceCore, []byte, bool) error
	GetPiece(context.Context, string, io.Writer, Options) (PieceReceipt, error)
	GetPieceBySerial(context.Context, uint64) (string, error)

	PutReplica(context.Context, ReplicaCore, []byte, bool) error
	GetReplica(context.Context, string, io.Writer, Options) (ReplicaCore, error)
	GetReplicaBySerial(context.Context, uint64) (string, error)

	DeleteData(ctx context.Context, name string) error

	PutReplicaWitness(context.Context, common.Address, ReplicaWitness) error
	GetReplicaWitness(context.Context, common.Address, uint64) (ReplicaWitness, error)

	PutFile(context.Context, FileCore, []string, bool) error
	GetFile(context.Context, string, Options) (FileReceipt, error)
}

type ICommitment interface {
	Add(ICommitment) error
	Serialize() []byte // 48
	Raw() []byte       // 96
}

type IProof interface {
	Type() int
	Add(IProof) error
	Serialize() []byte // 48+48
}

type IChallenge interface {
	Type() int
	Add(ICommitment) error
	Commitment() ICommitment
}

type IPublicKey interface {
	VerifyKey() IVerifyKey

	GenCommitments(int, io.Reader) ([]ICommitment, error)
	GenCommitment(int, []byte, int) (ICommitment, error)
	GenProofs(IChallenge, int, io.Reader) ([]IProof, error)
	GenProof(IChallenge, int, []byte) (IProof, error)

	Serialize() []byte
	Deserialize([]byte) error
}

type IVerifyKey interface {
	VerifyProof(IChallenge, IProof) error

	Serialize() []byte
	Deserialize([]byte) error
}
