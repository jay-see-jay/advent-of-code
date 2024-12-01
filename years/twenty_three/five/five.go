package five

import (
	"fmt"
	"strconv"
	"strings"
)

// For each seed, map:
// - seed to soil
// - soil to fertiliser
// - fertiliser to water
// - water to light
// - light to temperature
// - temperature to humidity
// - humidity to location

// For each line of each map:
// - If input number >= second number and < (second number + third number):
// -- second number to (second number + (second number + third number - 1))
// -- maps to...
// -- first number to (first number + (first number + third number - 1))
// - Else:
// -- return input number

// A map should take an input number:
// - check if it falls within any of the ranges
// - if yes apply modifier (second number - first number)

// map of maps:

type MapRange struct {
	sourceStart int
	sourceEnd   int
	modifier    int
}

type MapRanges = map[string][]MapRange

func extractSeedList(input string) []string {
	colonIndex := strings.Index(input, ":") + 1
	seeds := strings.TrimSpace(input[colonIndex:])
	return strings.Split(seeds, " ")
}

func ExtractSeeds(input string) []int {
	seedList := extractSeedList(input)

	var seedIntegers []int

	for _, seedStr := range seedList {
		seedInteger, err := strconv.Atoi(seedStr)
		if err != nil {
			fmt.Printf("Failed to convert seed to integer")
			continue
		}
		seedIntegers = append(seedIntegers, seedInteger)
	}

	return seedIntegers
}

func ExtractSeedsPart2(input string) []int {
	seedList := extractSeedList(input)

	var seedIntegers []int

	for i := 0; i < len(seedList); i += 2 {
		start, err := strconv.Atoi(seedList[i])
		if err != nil {
			continue
		}
		count, err := strconv.Atoi(seedList[i+1])
		if err != nil {
			continue
		}

		for count > 0 {
			seedIntegers = append(seedIntegers, start)
			start++
			count--
		}
	}

	return seedIntegers
}

func PopulateMaps(lines []string) MapRanges {
	var maps = make(MapRanges)
	var key string
	for _, line := range lines[1:] {
		if len(line) == 0 {
			key = ""
			continue
		}

		colonIndex := strings.Index(line, ":")
		if colonIndex >= 0 {
			key = line[:colonIndex]
			key = strings.Replace(key, " map", "", 1)
			continue
		}

		rangeNumbers := strings.Split(line, " ")
		destinationRangeStart, _ := strconv.Atoi(rangeNumbers[0])
		sourceRangeStart, _ := strconv.Atoi(rangeNumbers[1])
		rangeLength, _ := strconv.Atoi(rangeNumbers[2])
		sourceRangeEnd := sourceRangeStart + rangeLength
		modifier := destinationRangeStart - sourceRangeStart

		var mapRange = MapRange{
			sourceStart: sourceRangeStart,
			sourceEnd:   sourceRangeEnd,
			modifier:    modifier,
		}

		maps[key] = append(maps[key], mapRange)
	}

	return maps
}

func Map(input int, key string, maps MapRanges) int {
	var mapRanges = maps[key]
	for _, mapRange := range mapRanges {
		if input >= mapRange.sourceStart && input < mapRange.sourceEnd {
			return input + mapRange.modifier
		}
	}
	return input
}

func MapSeed(seed int, maps MapRanges) int {
	soil := Map(seed, "seed-to-soil", maps)
	fertilizer := Map(soil, "soil-to-fertilizer", maps)
	water := Map(fertilizer, "fertilizer-to-water", maps)
	light := Map(water, "water-to-light", maps)
	temperature := Map(light, "light-to-temperature", maps)
	humidity := Map(temperature, "temperature-to-humidity", maps)
	return Map(humidity, "humidity-to-location", maps)
}

func NearestLocation(seeds []int, maps MapRanges) int {
	var minLocation int = MapSeed(seeds[0], maps)
	for _, seed := range seeds[1:] {
		location := MapSeed(seed, maps)
		if location < minLocation {
			minLocation = location
		}
	}
	return minLocation
}

func RunPartOne(input string) int {
	lines := strings.Split(input, "\n")

	seeds := ExtractSeeds(lines[0])
	maps := PopulateMaps(lines)

	return NearestLocation(seeds, maps)
}

func RunPartTwo(input string) int {
	lines := strings.Split(input, "\n")
	seeds := ExtractSeedsPart2(lines[0])
	maps := PopulateMaps(lines)

	return NearestLocation(seeds, maps)
}
