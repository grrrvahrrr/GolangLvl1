//Lesson 10 Homework

package main

import (
	"GolangLvl1/lesson10/fibonacci"
	"fmt"
	"os"
)

func main() {
	var num int

	fmt.Println(`Please input a number greater than "0"`)

	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println(`Try again, something went wrong..`)
		os.Exit(1)
	}

	if num <= 0 {
		fmt.Println(`Try again, something went wrong..`)
		os.Exit(1)
	}

	fibbMap := make(map[int]int, num)
	fibbMap[0] = num
	fibbMap[1] = 1
	fibbMap[2] = 1
	fibonacci.FibbMapRecursion2(fibbMap)
	fmt.Println(fibbMap[num])

	fibbMap2 := make(map[int]int, num)
	fibonacci.FibbMapRecursion(fibbMap2, num)
	fmt.Println(fibbMap2[num])

	fmt.Println(fibonacci.FibbMapCalc(num))

	fmt.Println(fibonacci.FibbRecursion(num))

}
