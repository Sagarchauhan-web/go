package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Distance int32
type Velocity float64

type Number interface {
	constraints.Float | constraints.Integer
}

func clamp[T Number](value, min, max T) T {
	if value > max {
		return max
	} else if value < min {
		return min
	} else {
		return value
	}
}

func textClampInt8() {
	var (
		min int8 = -10
		max int8 = 10
	)

	if clamp(-30, min, max) != min {
		panic("failed for int8")
	}
}

func textClampInt32() {
	var (
		min uint32 = 0
		max uint32 = 10
	)

	if clamp(20, min, max) != min {
		panic("failed for uint32")
	}
}

func textClampFloat32() {
	var (
		min float32 = 5.5
		max float32 = 9.9
	)

	if clamp(0, min, max) != min {
		panic("failed for float32")
	}
}

func main() {
	textClampInt8()
	textClampInt32()
	textClampFloat32()

	fmt.Println("Exercise complete")
}