//Lesson 2 Homework

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	RectArea   = iota + 1 //Task 1 Rectengular area calculation
	CircleCalc            //Task 2 Circle diametr and length calculation
	NumberCalc            //Task 3 Digits printing
)

func main() {

	fmt.Println(`Choose task number: "1" for Rectengular area calculation, "2" for Circle diameter and length calculation, "3" for Digits printing.`)

	var task int

	fmt.Scan(&task)

	switch task {
	case RectArea:

		fmt.Println(`Input length, please`)

		var rectLength int
		_, err := fmt.Scan(&rectLength)

		if err != nil {
			fmt.Println(`Try again, something went wrong..`)
			os.Exit(1)
		}

		fmt.Println(`Input width, please`)

		var rectWidth int

		_, err = fmt.Scan(&rectWidth)

		if err != nil {
			fmt.Println(`Try again, something went wrong..`)
			os.Exit(1)
		}

		fmt.Println(rectLength * rectWidth)

		break

	case CircleCalc:

		fmt.Println(`Input Area of a Circle, please`)

		var circleArea float64
		_, err := fmt.Scan(&circleArea)

		if err != nil {
			fmt.Println(`Try again, something went wrong..`)
			os.Exit(1)
		}

		pi := math.Pi

		fmt.Println(`Circle Diameter is `, math.Sqrt(4*circleArea/pi))
		fmt.Println(`Circle Length is `, (math.Sqrt(4*circleArea/pi))*pi)

		break

	case NumberCalc:

		fmt.Println(`Input a three digit number, please`)

		var number int
		fmt.Scan(&number)

		numString := strconv.Itoa(number)

		if len(numString) != 3 {
			fmt.Println(`Try again, something went wrong..`)
			os.Exit(1)
		}

		hundreds, err := strconv.Atoi(string(numString[0]))
		tens, err := strconv.Atoi(string(numString[1]))
		singles, err := strconv.Atoi(string(numString[2]))

		fmt.Println(`Its hundreds:`, hundreds)
		fmt.Println(`Its tens:`, tens)
		fmt.Println(`Its singles:`, singles)

		if err != nil {
			fmt.Println(`Try again, something went wrong..`)
			os.Exit(1)
		}

		break

	default:
		fmt.Println(`Try again, something went wrong..`)
		os.Exit(1)
	}

}
