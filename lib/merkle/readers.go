package merkle

import (
	"bytes"
	"errors"
	"fmt"
	"hash"
	"io"
)

func (t *Tree) ReadAll(r io.Reader, segmentSize int) error {
	for {
		segment := make([]byte, segmentSize)
		n, readErr := io.ReadFull(r, segment)
		if readErr == io.EOF {
			break
		} else if readErr == io.ErrUnexpectedEOF {
			segment = segment[:n]
		} else if readErr != nil {
			return readErr
		}
		t.Push(segment)
	}
	return nil
}

func ReaderRoot(r io.Reader, h hash.Hash, segmentSize int) ([]byte, error) {
	tree := New(h)
	err := tree.ReadAll(r, segmentSize)
	if err != nil {
		return nil, err
	}
	return tree.Root(), nil
}

func BuildReaderProof(r io.Reader, h hash.Hash, segmentSize int, index uint64) (root []byte, proofSet [][]byte, numLeaves uint64, err error) {
	tree := New(h)
	err = tree.SetIndex(index)
	if err != nil {
		// This code should be unreachable - SetIndex will only return an error
		// if the tree is not empty, and yet the tree should be empty at this
		// point.
		panic(err)
	}
	err = tree.ReadAll(r, segmentSize)
	if err != nil {
		return
	}
	root, proofSet, _, numLeaves = tree.Prove()
	if len(proofSet) == 0 {
		err = errors.New("index was not reached while creating proof")
		return
	}
	return
}

func BuildDataAuxProof(d, aux []byte, h hash.Hash, segmentSize, shard int, index uint64) (root []byte, proofSet [][]byte, err error) {
	subaux := index / uint64(shard)
	sub := index % uint64(shard)

	if len(aux)/h.Size() <= int(subaux) {
		err = fmt.Errorf("unmatch data and aux length")
		return
	}

	tree := New(h)
	err = tree.SetIndex(sub)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(d)/segmentSize; i++ {
		tree.Push(d[i*segmentSize : (i+1)*segmentSize])
	}
	root, proofSet, _, _ = tree.Prove()
	tryup := 0
	for !bytes.Equal(root, aux[subaux*uint64(h.Size()):(subaux+1)*uint64(h.Size())]) {
		if tryup > 32 {
			err = fmt.Errorf("unmatch data and aux")
			return
		}
		proofSet = append(proofSet, root)
		root = nodeSum(h, root, root)
		tryup++
	}

	tree = New(h)
	err = tree.SetIndex(subaux)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(aux)/h.Size(); i++ {
		tree.PushHash(aux[i*h.Size() : (i+1)*h.Size()])
	}
	auxroot, auxproof, _, _ := tree.Prove()

	root = auxroot
	proofSet = append(proofSet, auxproof[1:]...)

	return
}

func BuildAuxProof(aux []byte, h hash.Hash, index uint64) (root []byte, proofSet [][]byte, err error) {
	tree := New(h)
	err = tree.SetIndex(index)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(aux)/h.Size(); i++ {
		tree.PushHash(aux[i*h.Size() : (i+1)*h.Size()])
	}
	root, proofSet, _, _ = tree.Prove()

	return
}

func BuildAuxRoot(aux []byte, h hash.Hash) ([]byte, error) {
	tree := New(h)
	for i := 0; i < len(aux)/h.Size(); i++ {
		tree.PushHash(aux[i*h.Size() : (i+1)*h.Size()])
	}
	return tree.Root(), nil
}

func ReaderRootAux(r io.Reader, h hash.Hash, segmentSize, shard int) ([]byte, []byte, error) {
	tree := New(h)
	subtree := New(h)

	aux := make([]byte, 0, h.Size()*shard)
	segment := make([]byte, segmentSize)
	for {
		n, readErr := io.ReadFull(r, segment)
		if readErr == io.EOF {
			break
		} else if readErr == io.ErrUnexpectedEOF {
			segment = segment[:n]
		} else if readErr != nil {
			return nil, nil, readErr
		}
		subtree.Push(segment)
		if subtree.currentIndex == uint64(shard) {
			auxtmp := subtree.Root()
			aux = append(aux, auxtmp...)
			tree.PushHash(auxtmp)
			subtree = New(h)
		}
	}
	if subtree.currentIndex < uint64(shard) && subtree.currentIndex > 0 {
		auxtmp := subtree.Root()

		// move up
		height := subtree.head.height
		head := subtree.head
		for head.next != nil {
			head = head.next
			height = head.height + 1
		}
		for i := (1 << height); i < shard; i *= 2 {
			auxtmp = nodeSum(h, auxtmp, auxtmp)
		}

		aux = append(aux, auxtmp...)
		tree.PushHash(auxtmp)
	}

	return tree.Root(), aux, nil
}
