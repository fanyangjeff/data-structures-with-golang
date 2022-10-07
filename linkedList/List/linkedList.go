package List

import (
	"fmt"
)

type any = interface{}

type Node struct {
	Prev *Node
	Next *Node
	Val  any
}

func CreateNode(val any) *Node {
	return &Node{nil, nil, val}
}

type LinkedList struct {
	Head *Node
	Size uint32
}

func CreateLinkedList() *LinkedList {
	head := CreateNode(-1)
	list := &LinkedList{head, 0}
	return list
}

func (This *LinkedList) Insert(node *Node) {
	node.Next = This.Head.Next
	This.Head.Next = node
	This.Size++
}

func (This *LinkedList) Remove(val interface{}) {
	cur := This.Head.Next
	prev := This.Head
	for cur != nil {
		if cur.Val == val {
			for cur.Next != nil && cur.Next.Val == val {
				cur = cur.Next
			}
			prev.Next = cur.Next
		}

		prev = cur
		cur = cur.Next
	}
}

func (This *LinkedList) Traverse() {
	cur := This.Head.Next
	fmt.Print("The LinkedList is: ")
	for cur != nil {
		fmt.Print(cur.Val, " ")
		cur = cur.Next
	}
}

func (This *LinkedList) size() uint32 {
	return This.Size
}
