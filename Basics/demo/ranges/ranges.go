package main

import "fmt"

func main() {
	slice := []string{"hello", "world","!"}

	for i, element := range slice {
		fmt.Println(element, i, ":")
		for _, ch := range element {
			fmt.Printf("  %q\n", ch)
		}
	}
}
