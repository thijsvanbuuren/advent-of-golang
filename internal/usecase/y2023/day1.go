package y2023

import (
	"log/slog"

	"github.com/thijsvanbuuren/advent-of-golang/internal/domain"
	"github.com/thijsvanbuuren/advent-of-golang/internal/usecase"
)

func init() {
	usecase.AppendExercise(domain.Exercise{
		Year: 2023,
		Day:  1,
		Part1: &domain.Part{
			// File: "day1.txt",
			Input: `
			dsfaf
			`,
			DoDay: Day1Part1,
		},
	})
}

func Day1Part1(input []string, logger *slog.Logger) any {
	return struct {
		test string
	}{test: "ok"}
}
