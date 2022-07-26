package tbtree

/*
	The file contains models for various algorithms.
*/

// BTreeValue is the binary tree recursive node
type BTreeValue struct {
	value float64     // Node value
	left  *BTreeValue // Left node
	right *BTreeValue // Right node
}

// BTree is the root of binary tree
type BTree struct {
	root *BTreeValue // BTreeValue node
}
