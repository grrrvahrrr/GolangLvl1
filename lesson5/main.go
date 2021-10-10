//Lesson 5 Homework

package main

import (
	"fmt"
	"os"
)

func fibbRecursion(x uint) uint {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	return fibbRecursion(x-1) + fibbRecursion(x-2)
}

func fibbMapCalc(x int) int {
	fibbMap := make(map[int]int, x+1)

	fibbMap[0] = 0
	fibbMap[1] = 1

	for i := 2; i < x+1; i++ {
		fibbMap[i] = fibbMap[i-1] + fibbMap[i-2]
	}

	if x == 0 {
		return fibbMap[0]
	} else if x == 1 {
		return fibbMap[1]
	}

	return fibbMap[x]
}

func fibbMapRecursion(numMap map[int]int, x int) (map[int]int, int) {

	numMap[0] = 0
	numMap[1] = 1
	length := len(numMap)

	numMap[length] = numMap[length-1] + numMap[length-2]

	if len(numMap) == x+1 {
		return numMap, x
	}
	if x == 1 {
		return numMap, x
	}

	return fibbMapRecursion(numMap, x)
}

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

	fibbMap := make(map[int]int, num+1)

	fibbMapRecursion(fibbMap, num)

	fmt.Println(fibbMap)
	fmt.Println(fibbMap[num])
}
