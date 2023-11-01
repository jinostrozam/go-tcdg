package main

import "fmt"

func main() {
	intSlices := []int{1, 2, 3, 5, 6, 7, 8, 9, 10}

	for _, slice := range intSlices {
		if slice%2 == 0 {
			fmt.Printf("%v is even\n", slice)
		} else {
			fmt.Printf("%v is odd\n", slice)
		}
	}
}
