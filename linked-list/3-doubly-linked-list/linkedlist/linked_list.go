package linkedlist

import "data-structures-algorithms/linked-list/3-doubly-linked-list/node"

type LinkedList struct {
	head *node.Node
	tail *node.Node
}

func NewLinkedList(node *node.Node) *LinkedList {
	linkedList := &LinkedList{}
	linkedList.initializeList(node)
	return linkedList
}

// O(1) - constant time - there is no increase of operation time, no matter the input
func (l *LinkedList) AddHead(data string) {
	newHead := node.NewNode(data)

	if l.isEmptyList() {
		l.initializeList(newHead)
		return
	}

	l.GetHead().SetPrevious(newHead)
	newHead.SetNext(l.GetHead())

	l.updateHead(newHead)
}

// Big O(1) - constant time - there is no increase of operation time, no matter the input
func (l *LinkedList) AddTail(data string) {
	newTail := node.NewNode(data)

	if l.isEmptyList() {
		l.initializeList(newTail)
		return
	}

	currentTail := l.GetTail()
	currentTail.SetNext(newTail)

	newTail.SetPrevious(currentTail)

	l.updateTail(newTail)
}

// Big O(n) - linear time - time increases linearly accordingly on how many items is linked
// In the worst case, the item we want to delete may be the last item so we have to iterate over all items linked.
func (l *LinkedList) DeleteFirstOcurrency(data string) {
	if l.GetHead().GetData() == data {
		l.updateHead(l.GetHead().GetNext())
		l.GetHead().SetPrevious(nil)
		return
	}

	if l.GetTail().GetData() == data {
		l.updateTail(l.GetTail().GetPrevious())
		l.GetTail().SetNext(nil)
		return
	}

	currentNode := l.GetHead().GetNext()

	for currentNode != nil {
		if currentNode.GetData() != data {
			currentNode = currentNode.GetNext()
			continue
		}

		currentNode.GetPrevious().SetNext(currentNode.GetNext())
		break
	}
}

// Big O(n) - linear time - time increases linearly accordingly on how many items is linked
func (l *LinkedList) GetNodesData() []string {
	var nodesData []string

	currentNode := l.GetHead()
	for currentNode != nil {
		nodesData = append(nodesData, currentNode.GetData())
		currentNode = currentNode.GetNext()
	}

	return nodesData
}

// Big O(1) - constant time - there is no increase of operation time
func (l *LinkedList) GetHead() *node.Node {
	return l.head
}

// Big O(1) - constant time - there is no increase of operation time
func (l *LinkedList) GetTail() *node.Node {
	return l.tail
}

// Big O(1) - constant time - there is no increase of operation time, no matter the input
func (l *LinkedList) initializeList(firstNode *node.Node) {
	l.updateHead(firstNode)
	l.updateTail(firstNode)
}

// Big O(1) - constant time - there is no increase of operation time, no matter the input
func (l *LinkedList) updateHead(newHead *node.Node) {
	l.head = newHead
}

// Big O(1) - constant time - there is no increase of operation time, no matter the input
func (l *LinkedList) updateTail(newTail *node.Node) {
	l.tail = newTail
}

// Big O(1) - constant time - there is no increase of operation time
func (l *LinkedList) isEmptyList() bool {
	return l.GetHead() == nil
}
