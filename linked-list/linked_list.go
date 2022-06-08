package main

import "fmt"

type Node struct {
	Name string
	Next *Node
}

func main() {
	node := &Node{Name: "2"}

	InsertAfter(node, "3")
	Push(&node, "1")
	Append(node, "4")
	Append(node, "5")

	PrintNodes(node)

}

func Push(headNode **Node, name string) {
	newNodeBefore := &Node{Name: name, Next: *headNode}
	*headNode = newNodeBefore
}

func InsertAfter(previousNode *Node, name string) {
	newNodeAfter := &Node{Name: name, Next: previousNode.Next}
	previousNode.Next = newNodeAfter
}

func Append(headNode *Node, name string) {
	newNode := &Node{Name: name, Next: nil}

	last := headNode
	for last.Next != nil {
		last = last.Next
	}
	last.Next = newNode
}

func PrintNodes(node *Node) {
	fmt.Println(node.Name)
	if node.Next != nil {
		PrintNodes(node.Next)
	}
}
