package structs

import "fmt"

type Node struct {
	val  int
	next *Node
}

func CreateLinkedList(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}

	var head *Node
	var temp *Node
	for _, val := range arr {
		newNode := new(Node)
		newNode.val = val
		// newNode := &Node{val: val}

		if head == nil {
			head = newNode
			temp = newNode
		} else {
			temp.next = newNode
			temp = temp.next
		}
	}

	return head
}

func PrintLinkedList(head *Node) {
	temp := head

	for temp != nil {
		fmt.Print(temp.val, " -> ")
		temp = temp.next
	}
	fmt.Println("nil")
}
