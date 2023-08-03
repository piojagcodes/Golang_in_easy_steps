package main

import "fmt"

func main() {
	for counter := 1; counter <= 5; counter++ {
		fmt.Println("Loop iteration: ", counter)
	}

	for i := 1; i <= 3; i++ {
		fmt.Println("\nOuter Loop iteration: ", i)
	}

	for j := 1; j <= 3; j++ {
		fmt.Println("\nInner Loop iteration: ", j)
	}

}
