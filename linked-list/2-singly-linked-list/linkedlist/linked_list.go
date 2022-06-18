package linkedlist

import "data-structures-algorithms/linked-list/2-singly-linked-list/node"

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
	currentHead := l.getHead()

	newHead.SetNext(currentHead)
	l.setHead(newHead)
}

// Big O(n) - This function executes in a linear time since we have to traverse the entire list (in the worst case) to put the new node.
func (l *LinkedList) Append(data int) {
	newNode := node.NewNode(data)

	if l.isEmptyList() {
		l.setHead(newNode)
		return
	}

	currentHead := l.getHead()
	for currentHead.GetNext() != nil {
		currentHead = currentHead.GetNext()
	}

	currentHead.SetNext(newNode)
}

// Big O(n) - This function executes in a linear time since we have to traverse the entire list (in the worst case) to find the new node.
func (l *LinkedList) DeleteFirstOcurrency(data int) {
	if l.isEmptyList() {
		return
	}

	currentHead := l.getHead()
	if currentHead.GetData() == data {
		l.setHead(currentHead.GetNext())
		return
	}

	current := l.getHead()
	previous := current

	for current != nil {
		if current.GetData() == data {
			next := current.GetNext()
			previous.SetNext(next)
			return
		}

		previous = current
		current = current.GetNext()
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
