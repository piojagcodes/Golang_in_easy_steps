package main

import "fmt"

func main() {
	counter := 1
	for counter <= 5 {
		fmt.Println("Loop Iteration", counter)
		counter++
	}
	// Statements to be inserted here.
	i := 10
	for {
		fmt.Println("\t\t\tCountdown", i)
		i--
	}

}
