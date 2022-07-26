package tbtree

func (b *BTree) Add(value float64) {

	if b.root == nil {
		b.root = &BTreeValue{value: value}
		return
	}

	newTree := &BTreeValue{value: value}

	add(b.root, newTree)
}

func add(currentTree *BTreeValue, newTree *BTreeValue) {
	cmp := compare(currentTree.value, newTree.value)

	if cmp > 0 {
		if currentTree.left == nil {
			currentTree.left = newTree
		} else {
			add(currentTree.left, newTree)
		}
	} else if cmp < 0 {
		if currentTree.right == nil {
			currentTree.right = newTree
		} else {
			add(currentTree.right, newTree)
		}
	}
}

func compare(v1 float64, v2 float64) float64 {
	return v1 - v2
}
