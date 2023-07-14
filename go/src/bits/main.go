package main
import "fmt"

func main() {
	var flags byte = 10 // Binary 1010 (1x8,0x4,1x2,0x1)
	fmt.Printf("Flag #1:%v\n", (flags & 1) > 0)
	fmt.Printf("Flag #2:%v\n", (flags & 2) > 0)
	fmt.Printf("Flag #3:%v\n", (flags & 4) > 0)
	fmt.Printf("Flag #4:%v\n", (flags & 8) > 0)

	fmt.Printf("Flags value: %08b \t%v \n", flags, flags)
}