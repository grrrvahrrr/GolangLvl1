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
type Operation interface {
	Do() (float64, error)
}

func (a *Addition) Do() float64 {
	return a.num1 + a.num2
}

func (s *Subtraction) Do() float64 {
	return s.num1 - s.num2
}

func (m *Multiplication) Do() float64 {
	return m.num1 * m.num2
}

func (s *Square) Do() float64 {
	return s.num1 * s.num1
}

func main() {
	fmt.Println(`Choose operation: "+" for Addition, "-" for Subtraction, "*" for Multiplication, "sqr" for squared value`)

	var operation string
	_, err := fmt.Scan(&operation)

	switch operation {
	case Add:
		var a Addition
		a.num1, a.num2, err = Num2Scan()
		if err != nil {
			err = Failure(`Try again, something went wrong..`)
			panic(err)
		}
		fmt.Printf("Result: %.2f\n", a.Do())

	case Sub:
		var s Subtraction
		s.num1, s.num2, err = Num2Scan()
		if err != nil {
			err = Failure(`Try again, something went wrong..`)
			panic(err)
		}
		fmt.Printf("Result: %.2f\n", s.Do())
	case Multi:
		var m Multiplication
		m.num1, m.num2, err = Num2Scan()
		if err != nil {
			err = Failure(`Try again, something went wrong..`)
			panic(err)
		}
		fmt.Printf("Result: %.2f\n", m.Do())

	case Squared:
		var s Square
		s.num1, err = Num1Scan()
		if err != nil {
			err = Failure(`Try again, something went wrong..`)
			panic(err)
		}
		fmt.Printf("Result: %.2f\n", s.Do())

	default:
		err = Failure(`Try again, something went wrong.. Invalid Operation`)
		panic(err)
	}

}
