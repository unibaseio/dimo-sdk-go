package types

import "github.com/fxamacker/cbor/v2"

type MerkleProof struct {
	Leaf uint64
	Path [][]byte
}

func (pmp *MerkleProof) Serialize() ([]byte, error) {
	return cbor.Marshal(pmp)
}

func (pmp *MerkleProof) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, pmp)
}

type SpaceProof struct {
	Leaf        uint64
	PiecePath   [][]byte
	ReplicaPath [][]byte
	C1Proof     []byte
	Aux         []byte
}

func (msp *SpaceProof) Serialize() ([]byte, error) {
	return cbor.Marshal(msp)
}

func (msp *SpaceProof) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, msp)
}

type SageMerkleProof struct {
	Leaf            uint64
	PieceRoot       []byte
	AuxRoot         []byte
	ReplicaRoot     []byte
	PiecePath       [][]byte
	AuxPrevPath     [][]byte
	AuxPath         [][]byte
	ReplicaPrevPath [][]byte
	ReplicaPath     [][]byte
	ReplicaNextPath [][]byte
}

func (smp *SageMerkleProof) Serialize() ([]byte, error) {
	return cbor.Marshal(smp)
}

func (smp *SageMerkleProof) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, smp)
}
