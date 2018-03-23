package main

import (
	"fmt"
)

//Node is a singly linked list node
type Node struct {
	value interface{}
	next  *Node
}

func next(node *Node) *Node {
	return node.next
}

//LinkedList struct
type LinkedList struct {
	head   *Node
	length int
}

func (n *LinkedList) init() {
	n.head = nil
	n.length = 0
}

func (n *LinkedList) push(val int) {
	nd := new(Node)
	nd.value = val
	nd.next = n.head
	n.head = nd
}

func (n *LinkedList) add(val int) {
	nd := new(Node)
	nd.value = val
	nd.next = nil
	temp := n.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = nd
}

func (n *LinkedList) printList() {
	nd := n.head
	for nd != nil {
		fmt.Println(nd.value)
		nd = nd.next
	}
}

func main() {
	listNode := new(LinkedList)
	listNode.init()
	listNode.push(12)
	listNode.push(3)
	listNode.add(15)
	listNode.printList()
}
