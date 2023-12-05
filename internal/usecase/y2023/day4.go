package y2023

import (
	"log/slog"
	"slices"
	"strings"

	"github.com/thijsvanbuuren/advent-of-golang/internal/domain"
	"github.com/thijsvanbuuren/advent-of-golang/internal/usecase"
)

func init() {
	usecase.AppendExercise(domain.Exercise{
		Year: 2023,
		Day:  4,
		Part1: &domain.Part{
			File: "day4.txt",
			Input: `
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
			`,
			DoDay: Day4Part1,
		},
		Part2: &domain.Part{
			File: "day4.txt",
			Input: `
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
			`,
			DoDay: Day4Part2,
		},
	})
}

func Day4Part1(input []string, logger *slog.Logger) any {
	count := 0
	for _, line := range input {
		data := strings.Split(line, ":")[1]
		sets := strings.Split(data, "|")
		winningNumbers := strings.Split(sets[0], " ")
		gottenNumbers := strings.Split(sets[1], " ")
		winningNumbers = slices.DeleteFunc(winningNumbers, func(e string) bool { return e == "" })
		gottenNumbers = slices.DeleteFunc(gottenNumbers, func(e string) bool { return e == "" })

		score := 0
		for _, nr := range gottenNumbers {
			if slices.Contains(winningNumbers, nr) {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
		count += score
	}
	return count
}

func Day4Part2(input []string, logger *slog.Logger) any {
	counted := make([]int, len(input))
	for i, line := range input {
		data := strings.Split(line, ":")[1]
		sets := strings.Split(data, "|")
		winningNumbers := strings.Split(sets[0], " ")
		gottenNumbers := strings.Split(sets[1], " ")
		winningNumbers = slices.DeleteFunc(winningNumbers, func(e string) bool { return e == "" })
		gottenNumbers = slices.DeleteFunc(gottenNumbers, func(e string) bool { return e == "" })

		count := 0
		for _, nr := range gottenNumbers {
			if slices.Contains(winningNumbers, nr) {
				count++
			}
		}
		counted[i] = count
	}
	// fmt.Println(counted)
	count := 0
	for i := len(counted) - 1; i >= 0; i-- {
		original := counted[i]
		counted[i] = 1 //itself
		for j := 1; j <= original; j++ {
			counted[i] += counted[i+j]
		}
		count += counted[i]
	}

	// fmt.Println(counted)
	return count
}
