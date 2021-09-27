package main

import (
	"fmt"
	"math"
	"os"
)

const (
	Addition       = "+"     //adding numbers
	Subtraction    = "-"     //subtracting numbers
	Multiplication = "*"     //multiplying numbers
	Division       = "/"     //divising numbers
	Squared        = "sqr"   //squared numbers
	SqrRoot        = "root"  //square root
	Logarithm      = "log"   //Logarithm of a number
	Prime          = "prime" //Get prime numbers upto this point
)

var number1, number2 float64
var err error

func twoNumberOperation() (float64, float64, error) {

	fmt.Println(`Input two numbers, please (divided by space)`)

	_, err = fmt.Scanf("%f %f", &number1, &number2)

	if err != nil {
		fmt.Println(`Try again, something went wrong..`)
		os.Exit(1)
	}

	return number1, number2, nil

}

func oneNumberOperation() (float64, error) {

	fmt.Println(`Input a number`)

	_, err = fmt.Scanf("%f", &number1)

	if err != nil {
		fmt.Println(`Try again, something went wrong..`)
		os.Exit(1)
	}

	return number1, nil
}

func main() {

	fmt.Println(`Choose operation: "+" for Addition, "-" for Subtraction, "*" for Multiplication, "/" for Division, "sqr" for Squared, "root" for Square root, "log" for Logarithm, "prime" to get prime numbers up to input`)

	var operation string

	_, err = fmt.Scan(&operation)

	if err != nil {
		fmt.Println(`Try again, something went wrong..`)
		os.Exit(1)
	}

	var result float64

	switch operation {
	case Addition:
		twoNumberOperation()

		result = number1 + number2
		fmt.Printf("Result: %.2f\n", result)

	case Subtraction:
		twoNumberOperation()

		result = number1 - number2
		fmt.Printf("Result: %.2f\n", result)

	case Multiplication:
		twoNumberOperation()

		result = number1 * number2
		fmt.Printf("Result: %.2f\n", result)

	case Division:
		twoNumberOperation()

		if number1 != 0 && number2 != 0 {
			result = number1 / number2
			fmt.Printf("Result: %.2f\n", result)
		} else {
			fmt.Println(`Can not use Zero`)
		}

	case Squared:
		oneNumberOperation()

		result = number1 * number1
		fmt.Printf("Result: %.2f\n", result)

	case SqrRoot:
		oneNumberOperation()

		result = math.Sqrt(number1)
		fmt.Printf("Result: %.2f\n", result)

	case Logarithm:
		oneNumberOperation()

		if number1 > 0 {
			result = math.Log(number1)
			fmt.Printf("Result: %.2f\n", result)
		} else {
			fmt.Println(`Can not use Zero or negative numbers`)
		}

	case Prime:

		var mbPrime int
		var isPrime bool
		var i int
		var j int

		fmt.Println(`Input a number`)

		_, err = fmt.Scanf("%d", &mbPrime)

		if err != nil {
			fmt.Println(`Try again, something went wrong..`)
			os.Exit(1)
		}

		for i = 1; i < mbPrime; i++ {
			isPrime = true

			for j = 2; j < i; j++ {
				if i%j == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				fmt.Printf("Prime Number: %d\n", i)
			}
		}

	default:
		fmt.Println(`Try again, something went wrong..`)
		os.Exit(1)
	}
}
