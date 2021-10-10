//Lesson 7 Homework

package main

import (
	"fmt"
	"math"
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

//Type Declaration
type TwoNums struct {
	num1, num2 float64
}

type OneNum struct {
	num1 float64
}

//Error handling
type Failure string

func (f Failure) Error() string {
	return string(f)
}

//Interfaces
type TwoNumInt interface {
	ScanNums() error
	Addition() float64
	Subtraction() float64
	Multiplication() float64
	Division() float64
}

type OneNumInt interface {
	ScanNum() error
	Squared() float64
	Sqrt() float64
	Log() float64
}

//Two Number Functions
func (nums *TwoNums) ScanNums() (err error) {
	fmt.Println(`Input two numbers, please (divided by space)`)
	_, err = fmt.Scanf("%f %f", &nums.num1, &nums.num2)
	return Failure(`Try again, something went wrong.. One of nums is invalid`)
}

func (nums *TwoNums) Addition() float64 {
	return nums.num1 + nums.num2
}

func (nums *TwoNums) Subtraction() float64 {
	return nums.num1 - nums.num2
}

func (nums *TwoNums) Multiplication() float64 {
	return nums.num1 * nums.num2
}

func (nums *TwoNums) Division() float64 {
	return nums.num1 / nums.num2
}

//One number functions
func (num *OneNum) ScanNum() (err error) {
	fmt.Println(`Input a number`)
	_, err = fmt.Scanf("%f", &num.num1)
	return Failure(`Try again, something went wrong.. Num is invalid`)
}

func (num *OneNum) Squared() float64 {
	return num.num1 * num.num1
}

func (num *OneNum) Sqrt() float64 {
	return math.Sqrt(num.num1)
}

func (num *OneNum) Log() float64 {
	return math.Log(num.num1)
}

func main() {
	var operation string

	fmt.Println(`Choose operation: "+" for Addition, "-" for Subtraction, "*" for Multiplication, "/" for Division, "sqr" for Squared, "root" for Square root, "log" for Logarithm`)

	_, err := fmt.Scan(&operation)
	if err != nil {
		err = Failure(`Try again, something went wrong.. Invalid Operation`)
		panic(err)
	}

	switch operation {
	case Addition:
		var i TwoNumInt = &TwoNums{}
		err = i.ScanNums()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Result: %.2f\n", i.Addition())

	case Subtraction:
		var i TwoNumInt = &TwoNums{}
		err = i.ScanNums()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Result: %.2f\n", i.Subtraction())

	case Multiplication:
		var i TwoNumInt = &TwoNums{}
		err = i.ScanNums()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Result: %.2f\n", i.Multiplication())

	case Division:
		var i TwoNumInt = &TwoNums{}
		err = i.ScanNums()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Result: %.2f\n", i.Division())

	case Squared:
		var i OneNumInt = &OneNum{}
		err = i.ScanNum()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Result: %.2f\n", i.Squared())

	case SqrRoot:
		var i OneNumInt = &OneNum{}
		err = i.ScanNum()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Result: %.2f\n", i.Sqrt())

	case Logarithm:
		var i OneNumInt = &OneNum{}
		err = i.ScanNum()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Result: %.2f\n", i.Log())

	default:
		err = Failure(`Try again, something went wrong.. Invalid Operation`)
		panic(err)
	}
}
