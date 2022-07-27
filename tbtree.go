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

// GetValues return all values from tree
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

// Min return minimum value from the tree
func (b *BTree) Min() float64 {
	node := min(b.root)
	return node.value
}

// Max return maximum value from the tree
func (b *BTree) Max() float64 {
	node := max(b.root)
	return node.value
}

// Remove delete value from the tree
func (b *BTree) Remove(value float64) {
	remove(b.root, value)
}

// add is a private function for recursive fill tree.
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

// getValues is the recursive function for return values from tree
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

// min is the recursive function for get min value
func min(currentTree *BTreeValue) *BTreeValue {
	if currentTree.left == nil {
		return currentTree
	}
	return min(currentTree.left)
}

// max is the recursive function for get max value
func max(currentTree *BTreeValue) *BTreeValue {
	if currentTree.right == nil {
		return currentTree
	}

	return max(currentTree.right)
}

// remove method for remove element from the tree
func remove(currentTree *BTreeValue, value float64) *BTreeValue {
	if currentTree == nil {
		return currentTree
	}

	if value < currentTree.value {
		currentTree.left = remove(currentTree.left, value)
	} else if value > currentTree.value {
		currentTree.right = remove(currentTree.right, value)
	} else if currentTree.left != nil && currentTree.right != nil {
		minNode := min(currentTree.right)

		currentTree.value = minNode.value
		currentTree.right = remove(currentTree.right, currentTree.value)
	} else if currentTree.left != nil {
		currentTree = currentTree.left
	} else if currentTree.right != nil {
		currentTree = currentTree.right
	} else {
		currentTree = nil
	}

	return currentTree
}

// compare is a private function for compare two numbers
func compare(v1 float64, v2 float64) float64 {
	return v1 - v2
}
