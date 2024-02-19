//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

const (
	Active = true
	Inactive = false
)

type Item struct {
	name string
	tag bool
}

func changeTagOfItem(item *Item, tagToChange bool) {
	item.tag = tagToChange
}

func checkout(items *[]Item) {
	for i := range *items {
		(*items)[i].tag = Inactive
	}
}

func main() {
	shoppingItems := []Item{
		{name:"towel", tag:Active},
		{name:"soap", tag:Active},
		{name:"shampoo", tag:Active},
		{name:"conditioner", tag:Active},
	}
	fmt.Println("shoppingItems",shoppingItems)

	changeTagOfItem(&shoppingItems[0], Inactive)
	fmt.Println("shoppingItems",shoppingItems)

	checkout(&shoppingItems)
	fmt.Println("shoppingItems",shoppingItems)
}
