package y2023

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/thijsvanbuuren/advent-of-golang/internal/domain"
	"github.com/thijsvanbuuren/advent-of-golang/internal/usecase"
)

func init() {
	usecase.AppendExercise(domain.Exercise{
		Year: 2023,
		Day:  3,
		Part1: &domain.Part{
			File: "day3.txt",
			Input: `
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
			`,
			DoDay: Day3Part1,
		},
	})
}

var numbers = "1234567890"

func Day3Part1(input []string, logger *slog.Logger) any {
	isSymbolGrid := make([][]bool, len(input))
	for y, line := range input {
		isSymbolGrid[y] = make([]bool, len(line))
		for x := range line {
			isSymbolGrid[y][x] = isSymbol(input, x, y)
		}
	}
	// fmt.Println(isSymbolGrid)

	var partNumbers []string
	isPart := false
	nr := ""
	foundTag := ""
	var tagged map[string][]string = make(map[string][]string)

	for y, line := range input {
		if isPart {
			partNumbers = append(partNumbers, nr)
			parts, ok := tagged[foundTag]
			if ok {
				tagged[foundTag] = append(parts, nr)
			} else {
				tagged[foundTag] = []string{nr}
			}
		}
		isPart = false
		nr = ""
		for x, c := range line {

			if strings.Contains(numbers, string(c)) {
				nr += string(c)
				isSymbol, fX, fY := checkSurround(isSymbolGrid, x, y)
				isPart = isPart || isSymbol
				if isSymbol {
					foundTag = fmt.Sprintf("%v:%v", fX, fY)
				}
			} else {
				if isPart {
					partNumbers = append(partNumbers, nr)
					parts, ok := tagged[foundTag]
					if ok {
						tagged[foundTag] = append(parts, nr)
					} else {
						tagged[foundTag] = []string{nr}
					}
				}
				isPart = false
				nr = ""
				continue
			}
		}
	}
	//	fmt.Println(tagged)

	// part 1
	count := 0
	for _, partN := range partNumbers {
		partInt, _ := strconv.Atoi(partN)
		count += partInt
	}
	// part 2
	var gearCount int64 = 0
	for k, val := range tagged {
		if len(tagged[k]) == 2 {
			fmt.Println(k)
			fmt.Println(val)
			v1, _ := strconv.Atoi(val[0])
			v2, _ := strconv.Atoi(val[1])
			gearCount += (int64(v1) * int64(v2))
		}

	}
	return struct {
		part1 int
		part2 int64
	}{
		part1: count,
		part2: gearCount,
	}
}

func checkSymbolGrid(input [][]bool, x, y int) bool {
	if y >= len(input) || y < 0 {
		return false
	}
	line := input[y]
	if x >= len(line) || x < 0 {
		return false
	}
	return input[y][x]
}
func isSymbol(input []string, x, y int) bool {
	if y >= len(input) || y < 0 {
		return false
	}
	line := input[y]
	if x >= len(line) || x < 0 {
		return false
	}
	c := line[x]
	if c == '.' {
		return false
	}
	if strings.Contains(numbers, string(c)) {
		return false
	}
	return true
}

// input is a [y][x] grid
func checkSurround(input [][]bool, x, y int) (bool, int, int) {
	baseX := x - 1
	baseY := y - 1
	for xOff := 0; xOff < 3; xOff++ {
		for yOff := 0; yOff < 3; yOff++ {
			xR := baseX + xOff
			yR := baseY + yOff
			if xR == x && yR == y {
				continue
			}
			if checkSymbolGrid(input, xR, yR) {
				return true, xR, yR
			}
		}
	}
	return false, -1, -1
}
