package main

import (
	"fmt"
)

type Stuff struct {
	values []int
}

func (s *Stuff) Get(index int) (int, error) {
	if index > len(s.values) - 1 {
		// errors.New() will also work
		return 0, fmt.Errorf("index %d out of bounds", index)
		} else {
			return s.values[index], nil
		}

}

func main() {
		stuff := Stuff{
			values: []int{1, 2, 3},
		}
		value, err := stuff.Get(2)


		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(value)
		}
}
