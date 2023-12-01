package usecase

import (
	"fmt"
	"log/slog"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/thijsvanbuuren/advent-of-golang/internal/domain"
)

const _filePathPattern = "./data/%v/%s"

var ExerciseList = []domain.Exercise{}

func AppendExercise(exercises domain.Exercise) {
	ExerciseList = append(ExerciseList, exercises)
}

func ExecuteExercises(logger *slog.Logger) {
	sort.SliceStable(ExerciseList, func(i, j int) bool {
		return ExerciseList[i].Day < ExerciseList[j].Day
	})
	for _, exercise := range ExerciseList {
		logger.Info("Starting", slog.Int("Year", exercise.Year), slog.Int("Day", exercise.Day))

		result1, dur1 := runPart(exercise.Part1, logger, exercise.Year, 1)
		logger.Info("P1:",
			slog.Duration("D", dur1),
			slog.Any("R", result1),
		)
		result2, dur2 := runPart(exercise.Part2, logger, exercise.Year, 2)
		logger.Info("P2:",
			slog.Duration("D", dur2),
			slog.Any("R", result2),
		)
	}

}

func runPart(part *domain.Part, logger *slog.Logger, year int, partName int) (any, time.Duration) {
	if part == nil {
		return nil, time.Duration(0)
	}
	var fileLines []string
	if len(part.Input) > 0 {
		fileLines = strings.Split(strings.TrimSpace(part.Input), "\n")
	}
	if len(part.File) > 0 {
		fl, err := loadFile(year, part, logger)
		fileLines = fl
		if err != nil {
			return "FILE NOT FOUND", time.Duration(0)
		}
	}

	now := time.Now()
	result := part.DoDay(fileLines, logger)
	return result, time.Since(now)

}

func loadFile(year int, part1 *domain.Part, logger *slog.Logger) ([]string, error) {
	file, err := os.ReadFile(fmt.Sprintf(_filePathPattern, year, part1.File))

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	fileLines := strings.Split(string(file), "\n")
	return fileLines, nil
}
