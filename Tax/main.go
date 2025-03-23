package main

import (
	"fmt"
)

func main() {
	var income float64
	fmt.Print("Enter your income-salary in past year: ")
	fmt.Scanf("%f", &income)
	var tax float64

	if income > 0 {

		tier := min(income, 100)
		tax += tier * 0.05
		income -= tier

	}

	if income > 0 {

		tier := min(income, 400)
		tax += tier * 0.10
		income -= tier
	}

	if income > 0 {

		tier := min(income, 500)
		tax += tier * 0.15
		income -= tier

	}

	if income > 0 {

		tax += income * 0.20

	}

	final := int(tax)
	fmt.Printf("Your total tax is: %d\n", final)
}

func min(a, b float64) float64 {

	if a > b {
		return b
	}
	return a
}
