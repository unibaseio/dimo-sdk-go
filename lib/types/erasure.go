package types

type IEncode interface {
	Encode([][]byte, []int) ([][]byte, error)
}

type IErasure interface {
	IEncode
	Check([][]byte, [][]byte, []int) error
	NewReconst([]int) (IEncode, error)
}
