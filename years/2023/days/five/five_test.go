package five_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jayseejay/advent-of-code/years/2023/days/five"
)

var example string = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

type Test = struct {
	name  string
	key   string
	input int
	want  int
}

func makeTest(key string, input int, want int) Test {
	mapElements := strings.Split(key, "-")
	return Test{
		name:  fmt.Sprintf("%s %d should be %s %d", mapElements[0], input, mapElements[2], want),
		key:   key,
		input: input,
		want:  want,
	}
}

func TestMapSeed(t *testing.T) {
	lines := strings.Split(example, "\n")
	maps := five.PopulateMaps(lines)

	var tests = []struct {
		name  string
		input int
		want  int
	}{
		{name: "Seed 79 -> Location 82", input: 79, want: 82},
		{name: "Seed 14 -> Location 43", input: 14, want: 43},
		{name: "Seed 55 -> Location 86", input: 55, want: 86},
		{name: "Seed 13 -> Location 35", input: 13, want: 35},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := five.MapSeed(test.input, maps)
			if got != test.want {
				t.Errorf("MapSeed(%d) = %d, want %d",
					test.input,
					got,
					test.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	lines := strings.Split(example, "\n")
	maps := five.PopulateMaps(lines)

	var tests = []Test{
		makeTest("seed-to-soil", 79, 81),
		makeTest("soil-to-fertilizer", 81, 81),
		makeTest("fertilizer-to-water", 81, 81),
		makeTest("water-to-light", 81, 74),
		makeTest("light-to-temperature", 74, 78),
		makeTest("temperature-to-humidity", 78, 78),
		makeTest("humidity-to-location", 78, 82),

		makeTest("seed-to-soil", 14, 14),
		makeTest("soil-to-fertilizer", 14, 53),
		makeTest("fertilizer-to-water", 53, 49),
		makeTest("water-to-light", 49, 42),
		makeTest("light-to-temperature", 42, 42),
		makeTest("temperature-to-humidity", 42, 43),
		makeTest("humidity-to-location", 43, 43),

		makeTest("seed-to-soil", 55, 57),
		makeTest("soil-to-fertilizer", 57, 57),
		makeTest("fertilizer-to-water", 57, 53),
		makeTest("water-to-light", 53, 46),
		makeTest("light-to-temperature", 46, 82),
		makeTest("temperature-to-humidity", 82, 82),
		makeTest("humidity-to-location", 82, 86),

		makeTest("seed-to-soil", 13, 13),
		makeTest("soil-to-fertilizer", 13, 52),
		makeTest("fertilizer-to-water", 52, 41),
		makeTest("water-to-light", 41, 34),
		makeTest("light-to-temperature", 34, 34),
		makeTest("temperature-to-humidity", 34, 35),
		makeTest("humidity-to-location", 35, 35),
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := five.Map(test.input, test.key, maps)
			if got != test.want {
				t.Errorf("Map(%d) = %d, want %d",
					test.input,
					got,
					test.want)
			}
		})
	}
}

func TestNearestLocation(t *testing.T) {
	lines := strings.Split(example, "\n")
	seeds := five.ExtractSeeds(lines[0])
	maps := five.PopulateMaps(lines)

	got := five.NearestLocation(seeds, maps)

	if got != 35 {
		t.Errorf("NearestLocation([82, 43, 86, 35]) = %v, want 35", got)
	}
}

func TestExtractSeeds(t *testing.T) {
	got := five.ExtractSeeds("seeds: 79 14 55 13")

	seeds := []int{79, 14, 55, 13}

	if len(got) != len(seeds) {
		t.Errorf("ExtractSeeds(seeds: 79 14...) = %v, want [79, 14, 55, 13]", got)
		return
	}

	for i, item := range got {
		if seeds[i] != item {
			t.Errorf("ExtractSeeds(seeds: 79 14...) = %v, want [79, 14, 55, 13]", got)
		}
	}
}

func TestExtractSeedsPart2(t *testing.T) {
	got := five.ExtractSeedsPart2("seeds: 79 14 55 13")

	seeds := []int{79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67}

	if len(got) != len(seeds) {
		t.Errorf("ExtractSeedsPart2(seeds: 79 14...) = %v, want [79, ..., 92, 55, ..., 67]", got)
		return
	}

	for i, item := range got {
		if seeds[i] != item {
			t.Errorf("ExtractSeedsPart2(seeds: 79 14...) = %v, want [79, ..., 92, 55, ..., 67]", got)
		}
	}
}

func TestRunPartOne(t *testing.T) {
	got := five.RunPartOne(example)

	if got != 35 {
		t.Errorf("RunPartOne(seeds: 79 14...) = %d, want 35", got)
	}
}

func TestRunPartTwo(t *testing.T) {
	got := five.RunPartTwo(example)

	if got != 46 {
		t.Errorf("RunPartTwo(seeds: 79 14...) = %d, want 46", got)
	}
}
