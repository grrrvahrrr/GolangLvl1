//Lesson 4 Homework

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sorting(sliceToSort []int) {
	for i := 1; i < len(sliceToSort); i++ {
		j := i
		for j > 0 {
			if sliceToSort[j-1] > sliceToSort[j] {
				sliceToSort[j-1], sliceToSort[j] = sliceToSort[j], sliceToSort[j-1]
			}
			j -= 1
		}
	}
	fmt.Println(`Your numbers are sorted: `, sliceToSort)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(`Enter integer numbers divided by "Space": `)

		scanner.Scan()

		numString := scanner.Text()
		if len(numString) != 0 {
			numStringSlice := strings.Split(numString, " ")
			length := len(numStringSlice)
			var toSort []int

			for i := 0; i < length; i++ {
				numToAdd, err := strconv.Atoi(numStringSlice[i])
				toSort = append(toSort, numToAdd)

				if err != nil {
					fmt.Println(`Try again, something went wrong..`)
					os.Exit(1)
				}
			}
			sorting(toSort)
			break
		} else {
			break
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}
}
