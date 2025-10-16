package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Go Calculator!")
	cmd := readLine("Enter command: [a]dd, [s]ubtract, [m]ultiply, [d]ivide: ")
	// num1 := readLine("Enter first number: ")
	// num2 := readLine("Enter second number: ")
	// fmt.Printf("You entered: %s %s %s\n", cmd, num1, num2)
	if cmd == "a" {
		num1, num2 := getNumbers()
		result := num1 + num2
		fmt.Println("Result:", result)
	} else if cmd == "s" {
		num1, num2 := getNumbers()
		result := num1 - num2
		fmt.Println("Result:", result)
	} else if cmd == "m" {
		num1, num2 := getNumbers()
		result := num1 * num2
		fmt.Println("Result:", result)
	} else if cmd == "d" {
		num1, num2 := getNumbers()
		result := float32(num1) / float32(num2)
		fmt.Println("Result:", result)
	} else {
		fmt.Println("Unknown command")
	}
}

func readLine(message string) string {
	fmt.Print(message)
	var input string
	fmt.Scanln(&input)
	return input
}

func getNumbers() (int, int) {
	num1Str := readLine("Enter first number: ")
	num1, err := strconv.Atoi(num1Str)
	if err != nil {
		fmt.Println("Invalid input. Please enter valid numbers.")
	}
	num2Str := readLine("Enter second number: ")
	num2, err := strconv.Atoi(num2Str)
	if err != nil {
		fmt.Println("Invalid input. Please enter valid numbers.")
	}
	return num1, num2
}
