package y2023

import (
	"log/slog"
	"strconv"
	"strings"

	"github.com/thijsvanbuuren/advent-of-golang/internal/domain"
	"github.com/thijsvanbuuren/advent-of-golang/internal/usecase"
)

func init() {
	usecase.AppendExercise(domain.Exercise{
		Year: 2023,
		Day:  1,
		Part1: &domain.Part{
			File: "day1.txt",
			Input: `
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
			`,
			DoDay: Day1Part1,
		},
		Part2: &domain.Part{
			File: "day1.txt",
			Input: `
			two1nine
			eightwothree
			abcone2threexyz
			xtwone3four
			4nineeightseven2
			zoneight234
			7pqrstsixteen
			`,
			DoDay: Day1Part2,
		},
	})
}

var numberNamesMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
}

func Day1Part1(input []string, logger *slog.Logger) any {
	const searchChars = "123456789"
	result := 0
	for _, line := range input {
		// fmt.Println(line)
		numbers := []string{}
		for _, char := range line {
			if strings.Contains(searchChars, string(char)) {
				numbers = append(numbers, string(char))
			}
		}
		combined := numbers[0] + numbers[len(numbers)-1]
		nr, _ := strconv.Atoi(combined)
		result += nr
		// fmt.Println(nr)
	}
	return result
}

func Day1Part2(input []string, logger *slog.Logger) any {
	// NOTE Replace does not work due to overlapping numbers
	// replaceList := make([]string, len(numberNamesMap)*2)
	// for k, v := range numberNamesMap {
	// 	replaceList = append(replaceList, k, v)
	// }
	// replacer := strings.NewReplacer(replaceList...)
	// for i := range input {

	// 	input[i] = replacer.Replace(input[i])
	// 	fmt.Println(input[i])
	// }
	searchTerms := make([]string, len(numberNamesMap))
	i := 0
	for k := range numberNamesMap {
		searchTerms[i] = k
		i++
	}
	// fmt.Println(searchTerms)

	result := 0
	for _, line := range input {
		// fmt.Println(line)
		first := numberNamesMap[find(line, searchTerms)]
		last := numberNamesMap[findLast(line, searchTerms)]

		combined := first + last
		nr, _ := strconv.Atoi(combined)
		result += nr
		// fmt.Println(combined)
	}
	return result
}

func find(str string, terms []string) string {
	var lowestIndex = 1000000
	var foundTerm = ""
	for _, term := range terms {
		// fmt.Println(term)
		index := strings.Index(str, term)
		if index > -1 && index < lowestIndex {
			lowestIndex = index
			foundTerm = term
		}
	}

	return foundTerm
}
func findLast(str string, terms []string) string {
	var highestIndex = -1
	var foundTerm = ""
	for _, term := range terms {
		index := strings.LastIndex(str, term)
		if index > highestIndex {
			highestIndex = index
			foundTerm = term
		}
	}

	return foundTerm
}
