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
func Do2nums(i interface{}, num1 float64, num2 float64) {
	switch v := i.(type) {
	case *Addition:
		fmt.Printf("Result: %.2f\n", v.Do2(num1, num2))
	case *Subtraction:
		fmt.Printf("Result: %.2f\n", v.Do2(num1, num2))
	case *Multiplication:
		fmt.Printf("Result: %.2f\n", v.Do2(num1, num2))
	}
}

func Do1num(i interface{}, num1 float64) {
	switch v := i.(type) {
	case *Square:
		fmt.Printf("Result: %.2f\n", v.Do1(num1))
	}
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
		num1, num2, err := Num2Scan()
		if err != nil {
			err = Failure(`Try again, something went wrong..`)
			panic(err)
		}
		Do2nums(&Addition{}, num1, num2)

	case Sub:
		num1, num2, err := Num2Scan()
		if err != nil {
			err = Failure(`Try again, something went wrong..`)
			panic(err)
		}
		Do2nums(&Subtraction{}, num1, num2)

	case Multi:
		num1, num2, err := Num2Scan()
		if err != nil {
			err = Failure(`Try again, something went wrong..`)
			panic(err)
		}
		Do2nums(&Multiplication{}, num1, num2)

	case Squared:
		num1, err := Num1Scan()
		if err != nil {
			err = Failure(`Try again, something went wrong..`)
			panic(err)
		}
		Do1num(&Square{}, num1)

	default:
		err = Failure(`Try again, something went wrong.. Invalid Operation`)
		panic(err)
	}

}
