//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type rectangle struct {
	length int
	width int
}

func calculateAreaAndPerimeter(rec rectangle) (int, int) {
	 area := rec.length * rec.width
	 perimeter := (2 * (rec.length + rec.width))
	return area,perimeter
}

func twoFoldRectangle(rec rectangle) (rectangle) {
	rec.length *= 2
	rec.width *= 2
	return rec
}

func main() {
	rec := rectangle{5,10};

	area, perimeter := calculateAreaAndPerimeter(rec)
	fmt.Println("Area: ",area, "Perimeter",perimeter)

	rec = twoFoldRectangle(rec)
	fmt.Println(rec.length,"length", rec.width,"width")

	area2, perimeter2 := calculateAreaAndPerimeter(rec)
	fmt.Println("Area: ",area2, "Perimeter",perimeter2)
}
