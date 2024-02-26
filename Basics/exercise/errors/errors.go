//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)


type Time struct {
	hour, minute, second int
}

type TimeParseError struct {
	msg string
	input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.msg, t.input)
}

// Example time string: 14:07:33
func ParseTime(input string) (Time, error) {
	components := strings.Split(input, ":")
	if len(components) != 3 {
		return Time{}, &TimeParseError{"invalid number of components", input}
	} else {
		hour, err := strconv.Atoi(components[0])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing hour: %v", hour), input}
		}
		minute, err := strconv.Atoi(components[1])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing hour: %v", minute), input}
		}
		second, err := strconv.Atoi(components[2])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing hour: %v", second), input}
		}

		if hour > 23 || hour < 0 {
			return Time{}, &TimeParseError{fmt.Sprintf("hour out of range: 0 <= hour <= 23, got: %v", hour), input}
		}

		if minute > 59 || minute < 0 {
			return Time{}, &TimeParseError{fmt.Sprintf("hour out of range: 0 <= minute <= 59, got: %v", minute), input}
		}
		
		if second > 59 || minute < 0 {
			return Time{}, &TimeParseError{fmt.Sprintf("hour out of range: 0 <= second <= 59, got: %v", second), input}
		}

		return Time{hour, minute, second}, nil
	}
}