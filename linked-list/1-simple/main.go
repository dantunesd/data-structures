package main

import "fmt"

type Node struct {
	Name string
	Next *Node
}

func main() {
	tail := &Node{Name: "3"}
	mid := &Node{Name: "2", Next: tail}
	head := &Node{Name: "1", Next: mid}

	PrintNodes(head)
}

func PrintNodes(node *Node) {
	fmt.Println(node.Name)
	if node.Next != nil {
		PrintNodes(node.Next)
	}
}
