package main

import "fmt"

func main() {
	n := make([]int, 11)

	for i := range n {
		n[i] = i
		if i%2 == 0 {
			fmt.Printf("%v is even", i)
		} else {
			fmt.Printf("%v is odd", i)
		}
		fmt.Println("")
	}
}
