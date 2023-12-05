package days

import (
	"fmt"
	"jayseejay/advent-of-code-23/days/one"
	"os"
)

func One() {
	b, err := os.ReadFile("./days/one/input.txt")
	if err != nil {
		panic("Failed to open file")
	}
	sum := one.Run(string(b))
	fmt.Printf("Sum: %d\n", sum)
	return
}
