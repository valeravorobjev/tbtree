package tbtree

import (
	"testing"
)

func TestBTree_Add(t *testing.T) {
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

func TestBTree_GetValues(t *testing.T) {
	btree := new(BTree)

	btree.Add(11)
	btree.Add(22)
	btree.Add(30)
	btree.Add(31)
	btree.Add(12)
	btree.Add(10)
	btree.Add(1)
	btree.Add(2)

	values := btree.GetValues()

	var tmp = values[0]
	for i := 1; i < len(values); i++ {
		if tmp > values[i] {
			t.Error("values can't order by acc because tmp > current value")
		}
	}
}
