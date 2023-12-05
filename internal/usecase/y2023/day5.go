package y2023

import (
	"fmt"
	"log/slog"
	"math"
	"slices"
	"strings"

	"github.com/thijsvanbuuren/advent-of-golang/internal/domain"
	"github.com/thijsvanbuuren/advent-of-golang/internal/usecase"
	"github.com/thijsvanbuuren/advent-of-golang/internal/util"
)

func init() {
	usecase.AppendExercise(domain.Exercise{
		Year: 2023,
		Day:  5,
		Part1: &domain.Part{
			File: "day5.txt",
			Input: `
seeds: 79 14 55 13

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
56 93 4
			`,
			DoDay: Day5Part1,
		},
		Part2: &domain.Part{
			File:  "day5.txt",
			DoDay: Day5Part2,
		},
	})
}

type CMap struct {
	SrcStart  int
	DestStart int
	Range     int
}

func (m CMap) Next(seed int) int {
	if seed < m.SrcStart || seed > (m.SrcStart+(m.Range-1)) {
		return -1
	}
	return seed - m.SrcStart + m.DestStart
}

func Day5Part1(input []string, logger *slog.Logger) any {
	seeds, maps := parseSeeds(input)

	chain := [][]CMap{
		maps["seed-to-soil"],
		maps["soil-to-fertilizer"],
		maps["fertilizer-to-water"],
		maps["water-to-light"],
		maps["light-to-temperature"],
		maps["temperature-to-humidity"],
		maps["humidity-to-location"],
	}
	lowest := math.MaxInt
	for _, seed := range seeds {
		loc := util.ToInt(seed)
		for _, cMaps := range chain {
			for _, cMap := range cMaps {
				tLoc := cMap.Next(loc)
				if tLoc != -1 {
					loc = tLoc
					break
				}
			}
		}
		fmt.Println(seed + " - " + fmt.Sprintf("%v", loc))
		if loc < lowest {
			lowest = loc
		}
	}
	return lowest
}
func Day5Part2(input []string, logger *slog.Logger) any {
	seeds, maps := parseSeeds(input)

	chain := [][]CMap{
		maps["seed-to-soil"],
		maps["soil-to-fertilizer"],
		maps["fertilizer-to-water"],
		maps["water-to-light"],
		maps["light-to-temperature"],
		maps["temperature-to-humidity"],
		maps["humidity-to-location"],
	}
	lowest := math.MaxInt
	fmt.Println(seeds)

	for i := 0; i < len(seeds); i += 2 {
		start := util.ToInt(seeds[i])
		till := start + util.ToInt(seeds[i+1]) - 1
		for seed := start; seed <= till; seed++ {
			loc := seed
			for _, cMaps := range chain {
				for _, cMap := range cMaps {
					tLoc := cMap.Next(loc)
					if tLoc != -1 {
						loc = tLoc
						break
					}
				}
			}
			// fmt.Println(seed + " - " + fmt.Sprintf("%v", loc))
			if loc < lowest {
				lowest = loc
			}
		}
	}
	return lowest
}
func parseSeeds(input []string) ([]string, map[string][]CMap) {
	seeds := strings.Split(strings.Split(input[0], ":")[1], " ")
	seeds = slices.DeleteFunc(seeds, func(e string) bool { return e == "" })

	result := make(map[string][]CMap)
	mapName := ""
	for i := 1; i < len(input); i++ {
		if input[i] == "" {
			mapName = strings.Split(input[i+1], " ")[0]
			result[mapName] = []CMap{}
			i++
			continue
		}

		mapData := strings.Split(input[i], " ")
		result[mapName] = append(result[mapName], CMap{DestStart: util.ToInt(mapData[0]), SrcStart: util.ToInt(mapData[1]), Range: util.ToInt(mapData[2])})
	}
	return seeds, result
}
