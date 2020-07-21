package algorithms

import (
	"testing"
)

func TestBinaryTree(t *testing.T) {
	root := BinaryTreeNode{
		item: 20,
	}

	tree := BinaryTree{
		root: &root,
	}

	t.Run("BinaryTree Check", func(t *testing.T) {
		if tree.root.item != 20 {
			t.Errorf("got %v, want %v", tree.root.item, 20)
		}

	})

}
