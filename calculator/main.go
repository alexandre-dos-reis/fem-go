package main

import "fmt"

func main() {
	fmt.Println("CALCULATOR GO")
	fmt.Println("=============")
	fmt.Println("Which operation do you want to perform ?")
	fmt.Println("add, substract, multiply or divide ?")

	var operation string
	fmt.Scanf("%s", &operation)
	fmt.Println("Enter first number")

	// https://pkg.go.dev/fmt@go1.22.1#hdr-Printing
	var number1, number2 int
	fmt.Scanf("%d", &number1)
	fmt.Println("Enter second number")
	fmt.Scanf("%d", &number2)

	switch operation {
	case "add":
		fmt.Println(number1 + number2)
	case "substract":
		fmt.Println(number1 - number2)
	case "multiply":
		fmt.Println(number1 * number2)
	case "divide":
		fmt.Println(number1 / number2)
	}
}
