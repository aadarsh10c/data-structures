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
	trains.AddAtStart("Train3")
	done := trains.AddAfter("Train2", "Traiin100")
	//add at end
	trains.AddAtEnd("Train4")
	fmt.Printf("was it added %t", done)
	fmt.Printf("list length %d\n", trains.Len())
	fmt.Printf("returned Train linked list : %v\n", trains.Values())
}
