package algorithms

// BirarnyTreeNode defines a struct for node in a singly linked list data structure
type BinaryTreeNode struct {
	item   interface{}
	left   *BinaryTreeNode
	right  *BinaryTreeNode
	parent *BinaryTreeNode
}

type BinaryTree struct {
	root *BinaryTreeNode
}

func (tree *BinaryTree) Search(needle interface{}) bool {
	node := tree.root
	for node != nil {
		if node.item == needle {
			return true
		}

	}
	return false

}
