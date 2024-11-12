package main

import (
	"fmt"

	"github.com/aadarsh10c/linkedLists/single"
)

func main() {
	//create a linked list
	trains := single.Create("Train1")
	//Add value at begining
	trains.AddAtStart("Train2")
	node, found := trains.GetNode("Train2")
	fmt.Printf("%v , %t\n", node, found)
	trains.AddAtStart("Train3")
	//add at end
	trains.AddAtEnd("Train4")
	done := trains.AddAfter("Train3", "Train101")
	fmt.Printf("Done %t\n", done)
	fmt.Printf("list length %d\n", trains.Len())
	fmt.Printf("returned Train linked list : %v\n", trains.Values())
}
