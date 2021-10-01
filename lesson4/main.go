//Lesson 4 Homework

package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println(`Please, enter number of positive integers to sort (more than 1)`)

	var length int

	_, err := fmt.Scan(&length)

	if err != nil || length <= 1 {
		fmt.Println(`Try again, something went wrong..`)
		os.Exit(1)
	}

	fmt.Println(`Please, input integers divided by "space"`)

	toSort := make([]int, length)

	if err != nil {
		fmt.Println(`Try again, something went wrong..`)
		os.Exit(1)
	}

	for i := 0; i < length; i++ {
		_, err = fmt.Scan(&toSort[i])

		if err != nil {
			fmt.Println(`Try again, something went wrong..`)
			os.Exit(1)
		}

	}

	for i := 1; i < length; i++ {

		j := i

		for j > 0 {
			if toSort[j-1] > toSort[j] {
				toSort[j-1], toSort[j] = toSort[j], toSort[j-1]
			}
			j -= 1
		}
	}
	/*
		val := toSort[i]

		if i == 0 && toSort[i] > toSort[i+1] {
			toSort = append(toSort[:0], toSort[i+1:]...)
			toSort = append(toSort, val)

		} else if i == 0 && toSort[i] < toSort[i+1] {
			continue

		} else if i == length-1 && toSort[i] < toSort[i-1] {
			toSort = toSort[:length-1]
			toSort = append([]int{val}, toSort...)

		} else if i == length-1 && toSort[i] > toSort[i-1] {
			continue

		} else if toSort[i] < toSort[i-1] || toSort[i] > toSort[i+1] {
			toSort = append(toSort[:i], toSort[i+1:]...)
			toSort = append(toSort, val)
		}*/

	fmt.Println(`Your numbers are sorted: `, toSort)

}
