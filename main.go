package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("====Starting the function done by Tejaswi which will calculate the factorial of a number====")
	var num int
	fmt.Println("Enter a number:")
	fmt.Scanln(&num)

	fact := factorial(num)
	fmt.Println("Factorial of", num, "is", fact)
	fmt.Println("")

	fmt.Println("====Starting the function done by Pamodi which will calculate the Area and Perimeter of a Rectangle====")
	// Area and Perimeter of Rectangle
	// Declare the length and width variables
	var length, width float64
	// Prompt user to enter the value of length and scan it
	fmt.Println("Enter the length of the rectangle: ")
	fmt.Scanln(&length)
	// Prompt user to enter the value of width and scan it
	fmt.Println("Enter the width of the rectangle: ")
	fmt.Scanln(&width)
	// Print user entered values
	fmt.Println("Your length is: ", length)
	fmt.Println("Your width is: ", width)
	// Call function to calculate area and perimeter
	area, perimeter := calcRectangleValues(length, width)
	// Print the values returned from the function
	fmt.Println("Area of the rectangle is: ", area)
	fmt.Println("perimeter of the rectangle is: ", perimeter)
	fmt.Println("")

	fmt.Println("====Starting the function done by Samhitha which will Print Fibonacci Series====")

	//Fibonacci Series
	// Set the limit for the Fibonacci series
	limit := 50
	// Generate and print the Fibonacci series
	result := fibonacci(limit)
	fmt.Printf("Fibonacci series up to %d: %v\n", limit, result)
	fmt.Println("")

	fmt.Println("====Starting the function done by Abhisheik Yadla which will print Hello world====")

	var year int

	fmt.Print("Enter a year: ")
	fmt.Scan(&year)

	if isLeapYear(year) {
		fmt.Println(year, "is a leap year.")
	} else {
		fmt.Println(year, "is not a leap year.")
	}

	fmt.Println("====Starting the function done by Mohammed Abdul which Reverse the String====")
	//Reverse a String
	// Getting input from the user
	var input string
	fmt.Println("Enter a string: ")
	fmt.Scanln(&input)
	// Calling the reverseString function to get the reversed string
	reversed := reverseString(input)
	// Printing the reversed string
	fmt.Printf("Reversed string: %s\n", reversed)
	fmt.Println("")

	fmt.Println("====Starting the function done by Syed Abdul Khadeer which will Calculate Power of number using Recursion and Binary to Decimal Conversion====")

	//Power of Exponent and Binary to Decimal conversion
	// Collect user input for the power calculation.
	var base, exponent int
	fmt.Println("Enter the base for power calculation: ")
	fmt.Scanln(&base)
	fmt.Println("Enter the exponent for power calculation: ")
	fmt.Scanln(&exponent)

	resultPower := power(base, exponent)
	fmt.Printf("%d^%d = %d\n", base, exponent, resultPower)

	// Get user input for binary-to-decimal conversion
	var binaryNumber string
	fmt.Println("Enter a binary number: ")
	fmt.Scanln(&binaryNumber)

	resultDecimal, err := binaryToDecimal(binaryNumber)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Binary: %s, Decimal: %d\n", binaryNumber, resultDecimal)
	}
	fmt.Println("")

	fmt.Println("====Starting the function done by Shubham Bathla which is accepting principal amount, rate of Interest and time of loan during execution and output Simple Interest====")
	calculateAndDisplaySimpleInterest()

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

// Created by Balangoda Pamodi - 500229522
// This function will accept length and width of a rectangle and return it's area and perimeter
func calcRectangleValues(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2

	return area, perimeter
}

// Created by Samhitha Dubbaka - 500225971
// fibonacci generates a Fibonacci series up to the given limit
func fibonacci(limit int) []int {
	// Initialize the series with the first two Fibonacci numbers
	series := []int{0, 1}
	// Generate Fibonacci numbers until reaching the limit
	for i := 2; ; i++ {
		nextNum := series[i-1] + series[i-2]
		// Break if the next number exceeds the limit
		if nextNum > limit {
			break
		}
		series = append(series, nextNum)
	}
	return series
}

// Created by Abhisheik Yadla - 500219580
// This function will check whether the given year is leap year or not
func isLeapYear(year int) bool {
	// Leap year conditions: divisible by 4, not divisible by 100 unless divisible by 400
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)

}

// Created by Mohammed Abdul - 500225922
// This function will reverse the string
func reverseString(input string) string {
	// To make reversing easier, convert the string to a rune slice.
	characters := []rune(input)

	// Determining the length of the slice
	length := len(characters)

	// Iterate through half of the loop and swaping the characters
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		// Swaping the characters at positions i and j
		characters[i], characters[j] = characters[j], characters[i]
	}

	// Converting the rune slice back to a string and return
	return string(characters)
}

// funtion Nine
// created by SYED ABDUL QADEER - 500228186
// program calculates the power of a number using recursion
// BINARY TO DECIMAL CONVERTOR
// Recursion is used in power to determine a number's power.
func power(base, exponent int) int {
	if exponent == 0 {
		return 1
	}
	return base * power(base, exponent-1)
}

// binaryToDecimal is a function that changes a binary number from a string representation to a decimal.
func binaryToDecimal(binary string) (decimal int, err error) {
	for i := len(binary) - 1; i >= 0; i-- {
		bit := binary[i]
		if bit != '0' && bit != '1' {
			return 0, fmt.Errorf("invalid binary digit: %c", bit)
		}
		decimal += int(bit-'0') * int(math.Pow(2, float64(len(binary)-1-i)))
	}
	return decimal, nil
}

// Function
// Created by Shubham Bathla - 500232317
// This function is accepting principal amount, rate of Interest and time of loan during execution and output Simple Interest
func calculateAndDisplaySimpleInterest() {
	var principal, rate, time float64

	// Getting user input
	fmt.Print("Enter the principal amount: ")
	fmt.Scanln(&principal)

	fmt.Print("Enter the rate of interest (in percent): ")
	fmt.Scanln(&rate)

	fmt.Print("Enter the time (in years): ")
	fmt.Scanln(&time)

	// Calculating simple interest
	interest := (principal * rate * time) / 100

	// Displaying the result
	fmt.Printf("The simple interest is: %f\n", interest)
}
