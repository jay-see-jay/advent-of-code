package years

import (
	"fmt"
	"os"
)

func OpenFile(day int) string {
	var path string = fmt.Sprintf("./years/2023/days/%d/input.txt", day)
	b, err := os.ReadFile(path)
	if err != nil {
		panic("Failed to open file")
	}
	return string(b)
}

func PrintResult(day int, part int, result int) {
	fmt.Printf("Day %d, Part %d\nResult: %d\n", day, part, result)
}
