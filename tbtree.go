package tbtree

/*
	The file contains methods and functions for implementing algorithms
*/

// Add value to the binary tree
func (b *BTree) Add(value float64) {

	if b.root == nil {
		b.root = &BTreeValue{value: value}
		return
	}

	newTree := &BTreeValue{value: value}

	add(b.root, newTree)
}

// GetValues return all values from b-tree
func (b *BTree) GetValues() []float64 {
	var values []float64
	getValues(b.root, &values)

	return values
}

// Contains method checks whether the specified value is contained in the tree
func (b *BTree) Contains(value float64) bool {
	ext := contains(b.root, value)
	return ext
}

// add is a private function for recursive fill b-tree.
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

// getValues is the recursive function for return values from b-tree
func getValues(currentTree *BTreeValue, values *[]float64) {
	if currentTree == nil {
		return
	}

	getValues(currentTree.left, values)
	*values = append(*values, currentTree.value)
	getValues(currentTree.right, values)
}

// contains is the recursive function for check exists value
func contains(currentTree *BTreeValue, value float64) bool {
	if currentTree == nil {
		return false
	}

	if value < currentTree.value {
		return contains(currentTree.left, value)
	}

	if value > currentTree.value {
		return contains(currentTree.right, value)
	}

	return true
}

// compare is a private function for compare two numbers
func compare(v1 float64, v2 float64) float64 {
	return v1 - v2
}
