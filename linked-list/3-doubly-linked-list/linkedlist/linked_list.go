package linkedlist

import "data-structures-algorithms/linked-list/3-doubly-linked-list/node"

type LinkedList struct {
	head *node.Node
}

func NewLinkedList(node *node.Node) *LinkedList {
	return &LinkedList{
		head: node,
	}
}

// Big O(n) - This function executes in a linear time since we have to traverse the entire list to get all data.
func (l *LinkedList) GetNodesData() []int {
	var nodesData []int

	currentHead := l.getHead()
	for currentHead != nil {
		nodesData = append(nodesData, currentHead.GetData())
		currentHead = currentHead.GetNext()
	}

	return nodesData
}

// Big O(1) - This function executes in a constant time since we don't have to traverse over any input
func (l *LinkedList) Push(data int) {
	newHead := node.NewNode(data)

	if l.isEmptyList() {
		l.setHead(newHead)
		return
	}

	currentHead := l.getHead()
	currentHead.SetPrevious(newHead)

	newHead.SetNext(currentHead)
	l.setHead(newHead)
}

// Big O(n) - This function executes in a linear time since we have to traverse the entire list (in the worst case) to put the new node.
func (l *LinkedList) Append(data int) {
	newTail := node.NewNode(data)

	if l.isEmptyList() {
		l.setHead(newTail)
		return
	}

	currentHead := l.getHead()
	for currentHead.GetNext() != nil {
		currentHead = currentHead.GetNext()
	}

	newTail.SetPrevious(currentHead)
	currentHead.SetNext(newTail)
}

// Big O(n) - This function executes in a linear time since we have to traverse the entire list (in the worst case) to find the node.
func (l *LinkedList) DeleteFirstOcurrency(data int) {
	if l.isEmptyList() {
		return
	}

	currentHead := l.getHead()
	if currentHead.GetData() == data {
		newHead := currentHead.GetNext()

		if newHead != nil {
			newHead.SetPrevious(nil)
		}

		l.setHead(newHead)
		return
	}

	current := l.getHead()

	for current != nil {

		if current.GetData() == data {
			previous := current.GetPrevious()
			next := current.GetNext()

			if previous != nil {
				previous.SetNext(next)
			}

			if next != nil {
				next.SetPrevious(previous)
			}

			return
		}

		current = current.GetNext()
	}
}

// Big O(n) - This function executes in a linear time since we have to traverse the entire list to swap one node to another.
func (l *LinkedList) Reverse() {
	if l.isEmptyList() {
		return
	}

	current := l.getHead()
	for current != nil {
		next := current.GetNext()
		previous := current.GetPrevious()

		current.SetPrevious(next)
		current.SetNext(previous)

		l.setHead(current)
		current = next
	}
}

// Big O(1)
func (l *LinkedList) getHead() *node.Node {
	return l.head
}

// Big O(1)
func (l *LinkedList) setHead(node *node.Node) {
	l.head = node
}

// Big O(1)
func (l *LinkedList) isEmptyList() bool {
	return l.head == nil
}
