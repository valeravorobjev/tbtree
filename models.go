package tbtree

type BTreeValue struct {
	value float64
	left  *BTreeValue
	right *BTreeValue
}

type BTree struct {
	root *BTreeValue
}
