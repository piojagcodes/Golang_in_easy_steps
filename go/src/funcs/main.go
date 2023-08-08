package main

import "fmt"

func first() {
	msg := "hello from first function"
	fmt.Println(msg)

}

func sqFive() int {
	return 5 * 5
}

func main() {
	first()
	fmt.Println("Result from 5 * 5 is:", sqFive())

}
