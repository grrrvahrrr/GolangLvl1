package main

import (
	"fmt"
)

const (
	Add     = "+"   //adding numbers
	Sub     = "-"   //subtracting numbers
	Multi   = "*"   //multiplying numbers
	Squared = "sqr" //squared numbers
)

//Error handling
type Failure string

func (f Failure) Error() string {
	return string(f)
}

//Structs
type Addition struct {
	num1, num2 float64
}

type Subtraction struct {
	num1, num2 float64
}

type Multiplication struct {
	num1, num2 float64
}

type Square struct {
	num1 float64
}

//Number Scanners
func Num2Scan() (a float64, b float64, err error) {
	fmt.Println(`Input two numbers, please (divided by space)`)
	_, err = fmt.Scanf("%f %f", &a, &b)
	return a, b, err
}

func Num1Scan() (a float64, err error) {
	fmt.Println(`Input a number`)
	_, err = fmt.Scanf("%f", &a)
	return a, err
}

//Do functions
type Operation2 interface {
	Do2(num1 float64, num2 float64) float64
}

type Operation1 interface {
	Do1(num1 float64) float64
}

func (a *Addition) Do2(num1 float64, num2 float64) float64 {
	return num1 + num2
}

func (s *Subtraction) Do2(num1 float64, num2 float64) float64 {
	return num1 - num2
}

func (m *Multiplication) Do2(num1 float64, num2 float64) float64 {
	return num1 * num2
}

func (s *Square) Do1(num1 float64) float64 {
	return num1 * num1
}

func main() {
	fmt.Println(`Choose operation: "+" for Addition, "-" for Subtraction, "*" for Multiplication, "sqr" for squared value`)

	var operation string
	_, err := fmt.Scan(&operation)

	switch operation {
	case Add:
		var i Operation2 = &Addition{}
		num1, num2, err := Num2Scan()
		if err != nil {
			err = Failure(`Try again, something went wrong..`)
			panic(err)
		}
		fmt.Printf("Result: %.2f\n", i.Do2(num1, num2))

	case Sub:
		var i Operation2 = &Subtraction{}
		num1, num2, err := Num2Scan()
		if err != nil {
			err = Failure(`Try again, something went wrong..`)
			panic(err)
		}
		fmt.Printf("Result: %.2f\n", i.Do2(num1, num2))

	case Multi:
		var i Operation2 = &Multiplication{}
		num1, num2, err := Num2Scan()
		if err != nil {
			err = Failure(`Try again, something went wrong..`)
			panic(err)
		}
		fmt.Printf("Result: %.2f\n", i.Do2(num1, num2))

	case Squared:
		var i Operation1 = &Square{}
		num1, err := Num1Scan()
		if err != nil {
			err = Failure(`Try again, something went wrong..`)
			panic(err)
		}
		fmt.Printf("Result: %.2f\n", i.Do1(num1))

	default:
		err = Failure(`Try again, something went wrong.. Invalid Operation`)
		panic(err)
	}

}
