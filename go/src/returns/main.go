package main

import "fmt"

func cube(num int) (string, int, int) {
	return "Return:", num, (num * num * num)

}

func main() {
	_, b, c := cube(5)
	fmt.Println(b, "Cubed = ", c)

}
