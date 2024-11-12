package single

type LinkedList[T comparable] struct {
	start *Node[T]
}

// structure for a linked list
type Node[K comparable] struct {
	data K
	next *Node[K]
}

//List of methods of Single Linked List

// Traverse the list ad return the searched node
func (l *LinkedList[T]) GetNode(val T) (*Node[T], bool) {
	ptr := l.start
	found := false
	lllen := l.Len()
	for i := 0; i < lllen; i++ {
		if found = (ptr.data == val); found {
			return ptr, found
		}
		ptr = ptr.next
	}
	return nil, found
}

// returns length of the linkd list
func (l *LinkedList[T]) Len() int {
	length := 0
	if l.start == nil {
		return length
	} else {
		//point to first node
		ptr := l.start
		length++
		for ptr.next != nil {
			length++
			ptr = ptr.next
		}
		return length
	}

}

// Add value at the start of the linked list
func (l *LinkedList[T]) AddAtStart(value T) {
	//create a new node
	newNode := new(Node[T])
	newNode.data = value

	if l.start != nil {
		newNode.next = l.start
	}
	l.start = newNode
}

// Add value at the end of linked list
func (l *LinkedList[T]) AddAtEnd(value T) {
	//if linked list is empty
	if l.start == nil {
		l.AddAtStart(value)
		return
	}
	//create a new Node
	newNode := new(Node[T])
	newNode.data = value

	ptr := l.start

	//fetch last node on the linked list
	for ptr.next != nil {
		ptr = ptr.next
	}

	ptr.next = newNode
}

// Add the value after a node
func (l *LinkedList[T]) AddAfter(valueAfter, newValue T) bool {
	//if linked list is empty
	if l.start == nil {
		l.AddAtStart(newValue)
		return true
	}

	//create a new Node
	newNode := new(Node[T])
	newNode.data = newValue

	//fetch the valueafter node and add the new Node after it
	if ptr, found := l.GetNode(valueAfter); found {
		if ptr.next != nil {
			newNode.next = ptr.next
			ptr.next = newNode
		} else {
			//valueAfter is a last value
			ptr.next = newNode
		}
		return true
	} else {
		return false
	}
}

// Add the Value before the target value
func (l *LinkedList[T]) AddBefore(targetValue, newValue T) bool {
	//if linked list is empty
	if l.start == nil {
		l.AddAtStart(newValue)
		return true
	}

	//create a new Node
	newNode := new(Node[T])
	newNode.data = newValue
	ptr := l.start
	prevPtr := ptr

	for prevPtr.next != nil {
		if ptr.data == targetValue {
			newNode.next = prevPtr.next
			prevPtr.next = newNode
			return true
		}
		prevPtr.next = ptr
		ptr = ptr.next
	}
	return false
}

// Return all the values returns a slice of values present in the Linked List
func (l *LinkedList[T]) Values() []T {
	//return an empty slice if no values added
	if l.Len() == 0 {
		return []T{}
	}

	//traverse the Linked list add values to the result slice
	lllen := l.Len()
	result := make([]T, lllen)

	//point the first value
	ptr := l.start
	for i := 0; i < lllen; i++ {
		result[i] = ptr.data
		ptr = ptr.next
	}
	return result
}

// creates a linked list and initialises first node with the given value
func Create[T comparable](value T) *LinkedList[T] {
	l := new(LinkedList[T])
	newNode := new(Node[T])
	newNode.data = value
	l.start = newNode
	return l
}
