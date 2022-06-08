package main

import (
	"data-structures-algorithms/linked-list/2-structured/linkedlist"
	"data-structures-algorithms/linked-list/2-structured/node"
	"fmt"
)

func main() {
	linkedList := linkedlist.NewLinkedList(
		node.NewNode("0"),
	)

	linkedList.AddHead("1")
	linkedList.AddHead("2")
	linkedList.AddHead("3")

	linkedList.AddTail("4")
	linkedList.AddTail("5")

	linkedList.DeleteFirstOcurrency("3")
	linkedList.DeleteFirstOcurrency("5")

	fmt.Println(linkedList.GetNodesData())
	fmt.Println(linkedList.GetHead().GetData())
	fmt.Println(linkedList.GetTail().GetData())
}
