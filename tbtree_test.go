package tbtree

import (
	"testing"
)

func initTree() *BTree {
	btree := new(BTree)

	btree.Add(11)
	btree.Add(22)
	btree.Add(30)
	btree.Add(31)
	btree.Add(12)
	btree.Add(10)
	btree.Add(1)
	btree.Add(2)

	return btree
}

func TestBTree_Add(t *testing.T) {
	btree := initTree()

	if btree.root == nil || btree.root.value != 11 {
		t.Error("btree.root is nil or empty")
	}

	if btree.root.right == nil || btree.root.right.value != 22 {
		t.Error("btree.root.right is nil or empty")
	}

	if btree.root.right.right == nil || btree.root.right.right.value != 30 {
		t.Error("btree.root.right.right is nil or empty")
	}

	if btree.root.right.right.right == nil || btree.root.right.right.right.value != 31 {
		t.Error("btree.root.right.right.right is nil or empty")
	}

	if btree.root.right.left == nil || btree.root.right.left.value != 12 {
		t.Error("btree.root.right.left is nil or empty")
	}

	if btree.root.left == nil || btree.root.left.value != 10 {
		t.Error("btree.root.left is nil or empty")
	}
}

func TestBTree_GetValues(t *testing.T) {
	btree := initTree()

	values := btree.GetValues()

	var tmp = values[0]
	for i := 1; i < len(values); i++ {
		if tmp > values[i] {
			t.Error("values can't order by acc because tmp > current value")
		}
	}
}

func TestBTree_Contains(t *testing.T) {
	btree := initTree()

	ext := btree.Contains(9)
	if ext {
		t.Error("btree isn't contains the value")
	}

	ext = btree.Contains(12)

	if !ext {
		t.Error("btree can't find value")
	}
}

func TestBTree_Max(t *testing.T) {
	btree := initTree()

	max := btree.Max()

	if max != 31 {
		t.Error("max is not 31")
	}
}

func TestBTree_Min(t *testing.T) {
	btree := initTree()

	min := btree.Min()
	if min != 1 {
		t.Error("min is not 1")
	}
}

func TestBTree_Remove(t *testing.T) {
	btree := initTree()

	btree.Remove(12)

	if btree.Contains(12) {
		t.Error("value 12 has not been removed from the tree")
	}
}
