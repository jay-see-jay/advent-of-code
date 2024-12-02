package years

import (
	"fmt"
	"os"
)

func getDayName(day int) string {
	dayMap := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}

	return dayMap[day]
}

func getYearName(year int) string {
	switch year {
	case 2024:
		return "twenty_four"
	case 2023:
		return "twenty_three"
	default:
		panic("Not that year!")
	}

}

func OpenFile(year int, day int) string {
	dayName := getDayName(day)
	yearName := getYearName(year)
	var path string = fmt.Sprintf("./years/%s/%s/input.txt", yearName, dayName)
	b, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		panic("Failed to open file")
	}
	return string(b)
}

func PrintResult(day int, part int, result int) {
	fmt.Printf("Day %d, Part %d\nResult: %d\n", day, part, result)
}
