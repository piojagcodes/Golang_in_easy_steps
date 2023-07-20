package main

import "fmt"

func main() {
	num := 2
	char := 'B'

	switch num {
	case 1:
		fmt.Println("\nNumber is One")
	case 2:
		fmt.Println("\nNumber is Two")
	case 3:
		fmt.Println("\nNumber is Three")
	default:
		fmt.Println("\nNumber is Unknown")
	}

	switch char {
	case 'A':
		fmt.Println("\nLetter is A")
	case 'B':
		fmt.Println("\nLetter is B")
	case 'C':
		fmt.Println("\nLetter is C")
	default:
		fmt.Println("\nLetter is Unknown")
	}
}
