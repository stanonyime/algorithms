package algorithms

// LinkedListNode defines a struct for node in a singly linked list data structure
type LinkedListNode struct {
	item interface{}
	next *LinkedListNode
}

// LinkedList defines a struct for a singly linked List data structure
type LinkedList struct {
	//items map[int]LinkedListNode // holds the items in the map
	head *LinkedListNode
}

// Search finds the node with a given value
func (list *LinkedList) Search(x interface{}) *LinkedListNode {
	node := list.head

	for node != nil {
		if node.item == x {
			return node
		} else {
			node = node.next
		}
	}
	return nil

}

//Insert adds a new node to a linked list and resets the head
func (list *LinkedList) Insert(x interface{}) {
	head := list.head

	node := LinkedListNode{
		item: x,
		next: head,
	}
	list.head = &node
	//list.items[x] = node

}

// Delete removes a node from a linked list
func (list *LinkedList) Delete(x interface{}) {
	node := list.Search(x)
	next := node.next
	head := list.head

	//delete(list.items, x.item)
	if head == node {
		list.head = next

	} else {
		prev := list.prevNode(x)
		prev.next = next
	}

}

func (list *LinkedList) prevNode(x interface{}) *LinkedListNode {
	node := list.head

	for node != nil {
		if node.next.item == x {
			return node
		} else {
			node = node.next
		}
	}
	return nil
}
