package main

// Add returns sum of input numbers
func Add(input []int) int {
	result := 0
	for _, value := range input {
		result = result + value
	}
	return result
}

// Subtract returns sum of input numbers
func Subtract(input []int) int {
	result := 0
	for _, value := range input {
		if result == 0 {
			result = value
		} else {
			result = result - value
		}
	}
	return result
}

// Multiply returns sum of input numbers
func Multiply(input []int) int {
	result := 1
	for _, value := range input {
		result = result * value
	}
	return result
}

// Devide returns sum of input numbers
func Divide(input []int) float64 {
	result := 1.0
	for _, value := range input {
		if result == 1.0 {
			result = float64(value)
		} else if value != 0 {
			result = result / float64(value)
		}
	}
	return result
}
