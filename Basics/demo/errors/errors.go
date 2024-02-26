package main

import (
	"errors"
	"fmt"
)

type Stuff struct {
	values []int
}

func (s *Stuff) Get(index int) (int, error) {
	if index > len(s.values) - 1 {
		return 0, errors.New(fmt.Sprintf("index %d out of bounds", index))
		} else {
			return s.values[index], nil
		}

}

func main() {
		stuff := Stuff{
			values: []int{1, 2, 3},
		}
		value, err := stuff.Get(6)


		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(value)
		}
}
