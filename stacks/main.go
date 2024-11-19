package main

import (
	"fmt"

	d "github.com/aadarsh10c/stacks/dynamicStack"
)

func main() {
	myStack := d.CreateDynamicStack[int]()

	//expect epty stack
	_, err := myStack.Display()
	if err != nil {
		fmt.Println(err.Error())
	}
	//add two values
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)

	//expect 3 values
	list, _ := myStack.Display()
	fmt.Printf("Values %v \n", list)

	//get the top value
	top, _ := myStack.Peek()
	// fmt.Printf("%s\n", err.Error())
	fmt.Printf("Top values is %d \n", top)

	//pop two values
	myStack.Pop()
	myStack.Pop()
	topValue, _ := myStack.Peek()
	fmt.Printf("Top Values is %d \n", topValue)
	myStack.Pop()
	topValue, err = myStack.Peek()
	if err != nil {
		fmt.Printf("%s \n", err.Error())
	}
	fmt.Printf("Top Values is %d \n", topValue)
}
