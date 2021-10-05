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

func twoNumberOperation() (number1 float64, number2 float64, err error) {
	fmt.Println(`Input two numbers, please (divided by space)`)
	_, err = fmt.Scanf("%f %f", &number1, &number2)
	return
}

func oneNumberOperation() (number1 float64, err error) {
	fmt.Println(`Input a number`)
	_, err = fmt.Scanf("%f", &number1)
	return
}

func main() {

	var operation string
	var result float64
	var number1, number2 float64

	fmt.Println(`Choose operation: "+" for Addition, "-" for Subtraction, "*" for Multiplication, "/" for Division, "sqr" for Squared, "root" for Square root, "log" for Logarithm, "prime" to get prime numbers up to input`)

	_, err := fmt.Scan(&operation)
	if err != nil {
		fmt.Println(`Try again, something went wrong..`)
		os.Exit(1)
	}

	switch operation {

	case Addition:
		number1, number2, err = twoNumberOperation()
		if err != nil {
			fmt.Println(`Try again, something went wrong..`)
			os.Exit(1)
		}
		result = number1 + number2
		fmt.Printf("Result: %.2f\n", result)

	case Subtraction:
		number1, number2, err = twoNumberOperation()
		if err != nil {
			fmt.Println(`Try again, something went wrong..`)
			os.Exit(1)
		}
		result = number1 - number2
		fmt.Printf("Result: %.2f\n", result)

	case Multiplication:
		number1, number2, err = twoNumberOperation()
		if err != nil {
			fmt.Println(`Try again, something went wrong..`)
			os.Exit(1)
		}
		result = number1 * number2
		fmt.Printf("Result: %.2f\n", result)

	case Division:
		number1, number2, err = twoNumberOperation()
		if err != nil {
			fmt.Println(`Try again, something went wrong..`)
			os.Exit(1)
		}
		if number1 != 0 && number2 != 0 {
			result = number1 / number2
			fmt.Printf("Result: %.2f\n", result)
		} else {
			fmt.Println(`Can not use Zero`)
		}

	case Squared:
		number1, err = oneNumberOperation()
		if err != nil {
			fmt.Println(`Try again, something went wrong..`)
			os.Exit(1)
		}
		result = number1 * number1
		fmt.Printf("Result: %.2f\n", result)

	case SqrRoot:
		number1, err = oneNumberOperation()
		if err != nil {
			fmt.Println(`Try again, something went wrong..`)
			os.Exit(1)
		}
		result = math.Sqrt(number1)
		fmt.Printf("Result: %.2f\n", result)

	case Logarithm:
		number1, err = oneNumberOperation()
		if err != nil {
			fmt.Println(`Try again, something went wrong..`)
			os.Exit(1)
		}
		if number1 > 0 {
			result = math.Log(number1)
			fmt.Printf("Result: %.2f\n", result)
		} else {
			fmt.Println(`Can not use Zero or negative numbers`)
		}

	case Prime:
		var mbPrime int
		var isPrime bool

		fmt.Println(`Input a number`)

		_, err = fmt.Scanf("%d", &mbPrime)
		if err != nil {
			fmt.Println(`Try again, something went wrong..`)
			os.Exit(1)
		}

		fmt.Println("Prime Number: 1")
		fmt.Println("Prime Number: 2")

		num := 3
		for num <= mbPrime {
			isPrime = true
			for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
				if num%i == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				fmt.Printf("Prime Number: %d\n", num)
			}
			num++
		}

	default:
		fmt.Println(`Try again, something went wrong..`)
		os.Exit(1)
	}
}
