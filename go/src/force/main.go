package main
import "fmt"

func main() {
	sum := 2 *3 + 4 - 5
	fmt.Printf("Default order: %v \n", sum)

	sum = 2 * ((3 + 4) - 5)
	fmt.Printf("Force order: %v \n\n", sum)

	sum = 7 % 3 * 2
	fmt.Printf("Default order: %v \n", sum)

	sum = 7 % (3 * 2)
	fmt.Printf("Forced order: %v \n", sum)
}