package dynamicstack

import "errors"

//Create a dynamic stack
//What is a dynamic stack
// stack which has a variable length
//It is based on linked list

type Node[K comparable] struct {
	data K
	next *Node[K]
}

type DynamicStack[K comparable] struct {
	top *Node[K]
}

func CreateDynamicStack[K comparable]() DynamicStack[K] {
	return DynamicStack[K]{}
}

func (d *DynamicStack[K]) Push(value K) {
	//create node
	ptr := Node[K]{
		data: value,
	}

	if d.top == nil {
		d.top = &ptr
	} else {
		ptr.next = d.top
		d.top = &ptr
	}
}

func (d *DynamicStack[K]) Pop() (value K, err error) {
	if d.top == nil {
		return value, errors.New("stack is empty")
	} else {
		value = d.top.data
		ptr := d.top.next
		d.top = ptr
		ptr = nil
		return value, nil
	}
}

func (d *DynamicStack[K]) Peek() (value K, err error) {
	if d.top == nil {
		return value, errors.New("stack is empty")
	} else {
		return d.top.data, nil
	}
}

func (d *DynamicStack[K]) Display() ([]K, error) {
	list := make([]K, 0)
	if d.top == nil {
		return list, errors.New("stack is empty")
	} else {
		topPtr := d.top
		for topPtr != nil {
			list = append(list, topPtr.data)
			topPtr = topPtr.next
		}
	}
	return list, nil
}
