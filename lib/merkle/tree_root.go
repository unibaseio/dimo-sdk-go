package merkle

import (
	"hash"
)

// only calculate merkle root
type RTree struct {
	hash      hash.Hash
	Count     uint64
	HeightSum [][]byte // sum at each height
}

func NewRTree(h hash.Hash) *RTree {
	rt := &RTree{
		hash:      h,
		HeightSum: make([][]byte, 0, 1),
	}
	rt.HeightSum = append(rt.HeightSum, make([]byte, 0, rt.hash.Size()))
	return rt
}

func (rt *RTree) Push(data []byte) {
	if len(data) < rt.hash.BlockSize() {
		padding := make([]byte, rt.hash.BlockSize()-len(data))
		if len(data) == 31 {
			padding[rt.hash.BlockSize()-32] = PadByte
		}
		data = append(padding, data...)
	}

	sum := leafSum(rt.hash, data)
	for i := 0; i < len(rt.HeightSum); i++ {
		if len(rt.HeightSum[i]) > 0 {
			// need one more height
			if i == len(rt.HeightSum)-1 {
				rt.HeightSum = append(rt.HeightSum, make([]byte, 0, rt.hash.Size()))
			}
			// clear current and cache for next
			sum = nodeSum(rt.hash, rt.HeightSum[i], sum)
			rt.HeightSum[i] = rt.HeightSum[i][:0]
		} else {
			rt.HeightSum[i] = sum
			break
		}
	}

	rt.Count++
}

func (rt *RTree) Root() []byte {
	var root []byte
	for i := 0; i < len(rt.HeightSum); i++ {
		if len(rt.HeightSum[i]) > 0 {
			if len(root) != 0 {
				root = nodeSum(rt.hash, rt.HeightSum[i], root)
			} else {
				if i < len(rt.HeightSum)-1 {
					root = nodeSum(rt.hash, rt.HeightSum[i], rt.HeightSum[i])
				} else {
					// top one
					root = rt.HeightSum[i]
				}
			}
		} else {
			if len(root) != 0 {
				root = nodeSum(rt.hash, root, root)
			}
		}
	}

	return root
}
