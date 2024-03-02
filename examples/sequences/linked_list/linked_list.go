package linked_list

type Node struct {
	  value int
	  next *Node
}

type LinkedList struct {
	  head *Node
	  size int
}

func NewLinkedList() *LinkedList {
	  return &LinkedList{}
}

func (l *LinkedList) InsertFront(value int) {
	newNode := Node{value: value, next: l.head}
	l.head = &newNode
	l.size++
}

func (l *LinkedList) InsertEnd(value int){
	newNode := Node{value: value, next: nil}
	if l.head == nil{
		l.head = &newNode
	}else{
		var current Node = *l.head
		for current.next != nil{
			current = *current.next
		}

		current.next = &newNode
	}
}

