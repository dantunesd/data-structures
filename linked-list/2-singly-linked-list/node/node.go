package node

type Node struct {
	data int
	next *Node
}

func NewNode(data int) *Node {
	return &Node{data: data}
}

// Big O(1) - constant time - there is no increase of operation time
func (n *Node) GetData() int {
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
