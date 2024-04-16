// Package merkle provides Merkle tree and proof following RFC 6962.
//
// From https://gitlab.com/NebulousLabs/merkletree
package merkle

import (
	"bytes"
	"errors"
	"hash"
)

const (
	PadByte = 0x01
)

type Tree struct {
	head *subTree
	hash hash.Hash

	currentIndex uint64
	proofIndex   uint64
	proofSet     [][]byte
	proofTree    bool
}

type subTree struct {
	next   *subTree
	height int // Int is okay because a height over 300 is physically unachievable.
	sum    []byte
}

// sum returns the hash of the input data using the specified algorithm.
func sum(h hash.Hash, data ...[]byte) []byte {
	h.Reset()
	for _, d := range data {
		// the Hash interface specifies that Write never returns an error
		_, err := h.Write(d)
		if err != nil {
			panic(err)
		}
	}
	return h.Sum(nil)
}

func leafSum(h hash.Hash, data []byte) []byte {
	return sum(h, data)
}

func nodeSum(h hash.Hash, a, b []byte) []byte {
	return sum(h, a, b)
}

// joinSubTrees combines two equal sized subTrees into a larger subTree.
func joinSubTrees(h hash.Hash, a, b *subTree) *subTree {
	return &subTree{
		next:   a.next,
		height: a.height + 1,
		sum:    nodeSum(h, a.sum, b.sum),
	}
}

func New(h hash.Hash) *Tree {
	return &Tree{
		hash: h,
	}
}

func (t *Tree) PushHash(data []byte) {
	if len(data) < t.hash.Size() {
		panic("data size not match")
	}

	node := make([]byte, t.hash.Size())
	copy(node, data)

	if t.currentIndex == t.proofIndex {
		t.proofSet = append(t.proofSet, node)
	}

	t.head = &subTree{
		next:   t.head,
		height: 0,
		sum:    node,
	}

	t.joinAllSubTrees()

	t.currentIndex++
}

func (t *Tree) Push(data []byte) {
	if len(data) < t.hash.BlockSize() {
		padding := make([]byte, t.hash.BlockSize()-len(data))
		if len(data) == 31 {
			padding[t.hash.BlockSize()-32] = PadByte
		}
		data = append(padding, data...)
	}

	node := make([]byte, len(data))
	copy(node, data)

	if t.currentIndex == t.proofIndex {
		t.proofSet = append(t.proofSet, node)
	}
	t.head = &subTree{
		next:   t.head,
		height: 0,
		sum:    leafSum(t.hash, data),
	}

	t.joinAllSubTrees()

	t.currentIndex++
}

func (t *Tree) SetIndex(i uint64) error {
	if t.head != nil {
		return errors.New("cannot call SetIndex on Tree if Tree has not been reset")
	}
	t.proofTree = true
	t.proofIndex = i
	return nil
}

func (t *Tree) joinAllSubTrees() {
	for t.head.next != nil && t.head.height == t.head.next.height {
		if t.head.height == len(t.proofSet)-1 {
			leaves := uint64(1 << uint(t.head.height))
			mid := (t.currentIndex / leaves) * leaves
			if t.proofIndex < mid {
				t.proofSet = append(t.proofSet, t.head.sum)
			} else {
				t.proofSet = append(t.proofSet, t.head.next.sum)
			}
		}

		t.head = joinSubTrees(t.hash, t.head.next, t.head)
	}
}

func (t *Tree) joinAndFillSubTrees(h hash.Hash, a, b *subTree, proofSet [][]byte) (*subTree, [][]byte) {
	nb := &subTree{
		height: b.height,
		sum:    make([]byte, len(b.sum)),
	}
	copy(nb.sum, b.sum)

	for nb.height < a.height {
		if nb.height == len(proofSet)-1 {
			proofSet = append(proofSet, nb.sum)
		}
		nb.sum = nodeSum(h, nb.sum, nb.sum) // fill right use left
		nb.height++
	}

	if nb.height == len(proofSet)-1 {
		leaves := uint64(1 << uint(nb.height))
		mid := (t.currentIndex / leaves) * leaves
		if t.proofIndex < mid {
			proofSet = append(proofSet, nb.sum)
		} else {
			proofSet = append(proofSet, a.sum)
		}
	}

	return &subTree{
		next:   a.next,
		height: a.height + 1,
		sum:    nodeSum(h, a.sum, nb.sum),
	}, proofSet
}

func (t *Tree) Root() []byte {
	if t.head == nil {
		return nil
	}
	current := t.head
	for current.next != nil {
		current, _ = t.joinAndFillSubTrees(t.hash, current.next, current, nil)
	}
	// Return a copy to prevent leaking a pointer to internal data.
	return append(current.sum[:0:0], current.sum...)
}

func (t *Tree) Prove() (merkleRoot []byte, proofSet [][]byte, proofIndex uint64, numLeaves uint64) {
	if !t.proofTree {
		panic("wrong usage: can't call prove on a tree if SetIndex wasn't called")
	}

	// Return nil if the Tree is empty, or if the proofIndex hasn't yet been
	// reached.
	if t.head == nil || len(t.proofSet) == 0 {
		return t.Root(), nil, t.proofIndex, t.currentIndex
	}
	proofSet = t.proofSet

	current := t.head
	for current.next != nil {
		current, proofSet = t.joinAndFillSubTrees(t.hash, current.next, current, proofSet)
	}

	return current.sum, proofSet, t.proofIndex, t.currentIndex
}

func VerifyProof(h hash.Hash, merkleRoot []byte, proofSet [][]byte, proofIndex uint64) bool {
	if len(proofSet) == 0 {
		return false
	}
	sum := leafSum(h, proofSet[0])
	for i := 1; i < len(proofSet); i++ {
		po := (proofIndex & (1 << (i - 1))) >> (i - 1)
		if po == 1 {
			sum = nodeSum(h, proofSet[i], sum)
		} else {
			sum = nodeSum(h, sum, proofSet[i])
		}
	}

	return bytes.Equal(sum, merkleRoot)
}
