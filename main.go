package main

import "fmt"

func main() {

	fmt.Println("====Starting the function done by Tejaswi which will calculate the factorial of a number====")
	var num int
	fmt.Println("Enter a number:")
	fmt.Scanln(&num)

	fact := factorial(num)
	fmt.Println("Factorial of", num, "is", fact)
	fmt.Println("")
}

// Created by Tejaswi Cheripally - 500229934
// This function will calculate the factorial of a number

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
