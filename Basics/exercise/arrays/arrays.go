//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name string
	price float64
}

func productDetails(products [4]Product) {
	var totalCost float64
	var totalItems int

	for i := 0; i < len(products); i++ {
		currentItem := products[i]

		if currentItem.name == "" {
			break
		}

		totalCost += currentItem.price
		totalItems += 1
	}
	fmt.Println(totalCost, "TotalCost",totalItems,
	"total Items",products[totalItems - 1].name,"lastItem" ,"here")
}

func main() {
	productList := [4]Product{
		{name: "towel", price: 50},
		{name: "shampoo", price: 10},
		{name: "soap", price: 10},
	}

	productDetails([4]Product(productList))
	
	productList[3] =  Product{name: "conditioner", price: 50}
	
	productDetails([4]Product(productList))
}

