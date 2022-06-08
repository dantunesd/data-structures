package main

import (
	"data-structures-algorithms/linked-list/3-doubly-linked-list/linkedlist"
	"data-structures-algorithms/linked-list/3-doubly-linked-list/node"
	"fmt"
)

func main() {
	linkedList := linkedlist.NewLinkedList(node.NewNode("1"))

	linkedList.AddHead("0")
	linkedList.AddTail("2")
	linkedList.AddTail("3")

	linkedList.DeleteFirstOcurrency("3")

	fmt.Println(linkedList.GetNodesData())
	fmt.Println(linkedList.GetHead().GetData())
	fmt.Println(linkedList.GetTail().GetData())
}
