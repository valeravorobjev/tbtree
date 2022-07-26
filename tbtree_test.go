package tbtree

import "testing"

func TestBTree(t *testing.T) {
	btree := new(BTree)

	btree.Add(11)

	if btree.root == nil || btree.root.value != 11 {
		t.Error("btree.root is nil or empty")
	}

	btree.Add(22)

	if btree.root.right == nil || btree.root.right.value != 22 {
		t.Error("btree.root.right is nil or empty")
	}

	btree.Add(30)

	if btree.root.right.right == nil || btree.root.right.right.value != 30 {
		t.Error("btree.root.right.right is nil or empty")
	}

	btree.Add(31)

	if btree.root.right.right.right == nil || btree.root.right.right.right.value != 31 {
		t.Error("btree.root.right.right.right is nil or empty")
	}

	btree.Add(12)

	if btree.root.right.left == nil || btree.root.right.left.value != 12 {
		t.Error("btree.root.right.left is nil or empty")
	}

	btree.Add(10)

	if btree.root.left == nil || btree.root.left.value != 10 {
		t.Error("btree.root.left is nil or empty")
	}
}
