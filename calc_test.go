package main

import (
	"errors"
	"testing"
)

func TestAdd(t *testing.T) {
	var err error
	result := 0.0

	data := []struct {
		op       string
		input    []int
		expected float64
	}{
		{"Add", []int{5, 2, 3}, 10.0},
		{"Subtract", []int{7, 2, 3}, 2.0},
		{"Multiply", []int{2, 2, 2}, 8.0},
		{"Divide", []int{18, 2, 3}, 3.0},
	}
	for _, value := range data {
		switch value.op {
		case "Add":
			result = float64(Add(value.input))
		case "Subtract":
			result = float64(Subtract(value.input))
		case "Multiply":
			result = float64(Multiply(value.input))
		case "Divide":
			result = Divide(value.input)
		default:
			err = errors.New("Unknown operator '" + value.op + "'")
		}
		if nil == err {
			if value.expected != result {
				t.Errorf("Expected value %f is not equal to result %f for %s", value.expected, result, value.op)
			}
		}
	}
}
