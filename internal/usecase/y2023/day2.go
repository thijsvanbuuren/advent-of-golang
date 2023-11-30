package y2023

import (
	"log/slog"

	"github.com/thijsvanbuuren/advent-of-golang/internal/domain"
	"github.com/thijsvanbuuren/advent-of-golang/internal/usecase"
)

func init() {
	usecase.AppendExercise(domain.Exercise{
		Year: 2023,
		Day:  2,
		Part1: &domain.Part{
			Input: `
			afsaf
			asdfsaf
			`,
			DoDay: Day2Part1,
		},
	})
}

func Day2Part1(input []string, logger *slog.Logger) any {
	return 234
}
