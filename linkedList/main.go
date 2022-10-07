package main

import (
	"fmt"
	"linkedList/List"
)

func main() {
	linkedList := List.CreateLinkedList()
	node0 := List.CreateNode(3)
	node1 := List.CreateNode(1)
	node2 := List.CreateNode(2)
	node3 := List.CreateNode(3)
	node4 := List.CreateNode(3)
	linkedList.Insert(node0)
	linkedList.Insert(node1)
	linkedList.Insert(node2)
	linkedList.Insert(node3)
	linkedList.Insert(node4)
	fmt.Println("The size of the linked list is: ", linkedList.Size)
	linkedList.Traverse()

	linkedList.Remove(3)
	linkedList.Traverse()

}
