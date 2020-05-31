package types

import "parser"

type k int
type (
	Point struct{ x, y string }
	polar Point
)

type TreeNode struct {
	left, right *TreeNode
	value *Comparable
}

type Block interface {
	BlockSize() int
	Encrypt(src, dst []byte)
	Decrypt(src, dst []byte)
}

type k parser.Mode

// unexported recursive type is fine
type recursive struct {
	child *recursive
}
