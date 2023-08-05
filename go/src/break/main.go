package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		// Inner loop can be inserted here
		for j := 1; j <= 3; j++ {
			fmt.Println("Running i=", i, "j=", j)
		}
	}

}
