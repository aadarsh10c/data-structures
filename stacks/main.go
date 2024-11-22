package main

import (
	"fmt"

	d "github.com/aadarsh10c/stacks/dynamicStack"
)

func main() {
	list := []int{1, 2, 3, 4, 5}
	reverse, _ := d.ReverseList(list)
	fmt.Printf("Original %v , Reversed %v\n", list, reverse)
}
