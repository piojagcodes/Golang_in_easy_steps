package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		// Inner loop can be inserted here
		for j := 1; j <= 3; j++ {
			// Continue statement
			if i == 3 && j == 2 {
				fmt.Println("Continues when i=", i, "Continues when j=", j)
				continue
			}
			// Statements to be inserted here.
			if i == 2 && j == 2 {
				fmt.Println("Breaks when i=", i, "Breaks when j=", j)
			}
			fmt.Println("Running i=", i, "j=", j)
		}
	}

}
