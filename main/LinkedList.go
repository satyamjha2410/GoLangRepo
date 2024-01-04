package main

import "fmt"

// Node represents a node in the linked list
type Node struct {
    data int
    next *Node
}

// LinkedList represents the linked list
type LinkedList struct {
    head *Node
}

// Insert adds a new node with the given data at the end of the linked list
func (ll *LinkedList) Insert(data int) {
    newNode := &Node{data: data, next: nil}

    if ll.head == nil {
        ll.head = newNode
        return
    }

    current := ll.head
    for current.next != nil {
        current = current.next
    }

    current.next = newNode
}

// Display prints the elements of the linked list
func (ll *LinkedList) Display() {
    current := ll.head
    for current != nil {
        fmt.Printf("%d -> ", current.data)
        current = current.next
    }
    fmt.Println("nil")
}

func main() {
    // Create a linked list
    linkedList := LinkedList{}

    // Insert elements into the linked list
    linkedList.Insert(1)
    linkedList.Insert(2)
    linkedList.Insert(3)
    linkedList.Insert(4)

    // Display the linked list
    fmt.Println("Linked List:")
    linkedList.Display()
}
