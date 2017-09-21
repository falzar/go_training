package main

import (
	"fmt"
)

func main() {

	numbers := newNumericArray(0, 10)

	for _, number := range numbers {

		if number%2 == 0 {
			fmt.Printf("%v is even\n", number)
		} else {
			fmt.Printf("%v is odd\n", number)
		}
	}

}

func newNumericArray(ini int, end int) []int {
	var numericArray []int

	for i := ini; i <= end; i++ {
		numericArray = append(numericArray, i)
	}

	return numericArray
}
