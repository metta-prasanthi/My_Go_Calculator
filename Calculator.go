package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers := make([]int, 0)
	fmt.Println("Pleae enter input followed by / : ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() != "/" {
			numbers = append(numbers, stringToInt(scanner.Text()))
		} else {
			break
		}
	}
	// Error handling
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	fmt.Println(numbers)
	fmt.Println("Addition of input numbers : ")
	fmt.Println(Add(numbers))
	fmt.Println("Subtraction of input numbers : ")
	fmt.Println(Subtract(numbers))
	fmt.Println("Multiplication of input numbers : ")
	fmt.Println(Multiply(numbers))
	fmt.Println("division of input numbers : ")
	fmt.Println(Divide(numbers))
}

func stringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return i
}
