package node

type Node struct {
	data     string
	previous *Node
	next     *Node
}

func NewNode(data string) *Node {
	return &Node{data: data}
}

// Big O(1) - constant time - there is no increase of operation time
func (n *Node) GetData() string {
	return n.data
}

// Big O(1) - constant time - there is no increase of operation time
func (n *Node) GetNext() *Node {
	return n.next
}

// Big O(1) - constant time - there is no increase of operation time, no matter the input
func (n *Node) SetNext(next *Node) {
	n.next = next
}

// Big O(1) - constant time - there is no increase of operation time
func (n *Node) GetPrevious() *Node {
	return n.previous
}

// Big O(1) - constant time - there is no increase of operation time, no matter the input
func (n *Node) SetPrevious(previous *Node) {
	n.previous = previous
}
