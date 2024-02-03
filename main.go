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
