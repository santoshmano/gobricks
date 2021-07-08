package singly_linked_list

// A single node of a singly linked list
type SinglyLinkedNode struct {
	// value to hold in the node
	Val interface{}
	// Next pointer
	Next *SinglyLinkedNode
}

// A singly linked list
type SinglyLinkedList struct {
	// pointer to the first node
	head *SinglyLinkedNode
	// length of the ll
	len int
}

func (l *SinglyLinkedList) init() *SinglyLinkedList {
	l.head = nil
	l.len = 0
	return l
}

// Create a new Singly Linked List and initialize it
func NewSinglyLinkedList() *SinglyLinkedList {
	return new(SinglyLinkedList).init()
}

// Return the length of the list
func (l *SinglyLinkedList) Len() int {
	return l.len
}

// Return the first node in the list
func (l *SinglyLinkedList) Head() *SinglyLinkedNode {
	return l.head
}

// InsertVal -  Insert a value into the head of the list, return the new node pointer
func (l *SinglyLinkedList) InsertVal(val interface{}) *SinglyLinkedNode {
	return l.InsertNode(&SinglyLinkedNode{val, nil})
}

// Insert a node to the head of the list, return back the same node pointer
func (l *SinglyLinkedList) InsertNode(node *SinglyLinkedNode) *SinglyLinkedNode {
	node.Next = l.head
	l.head = node
	l.len += 1

	return node
}

// Delete the first node with a value from the list
func (l *SinglyLinkedList) DeleteVal(val interface{}) {
	var prev *SinglyLinkedNode = nil

	for cur := l.head; cur != nil; cur = cur.Next {
		if cur.Val == val {
			if cur == l.head {
				l.head = cur.Next
			} else {
				prev.Next = cur.Next
			}
			cur.Next = nil
			cur.Val = nil
		}
		prev = cur
	}
}

// Find a value in the list, return back the value in the Node
func (l *SinglyLinkedList) FindVal(val interface{}) interface{} {
	for cur := l.head; cur != nil; cur = cur.Next {
		if cur.Val == val {
			return cur.Val
		}
	}
	return nil
}

func (l *SinglyLinkedList) HasCycle() bool {
	slow := l.head
	fast := l.head

	for (fast != nil) && (fast.Next != nil) {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}
