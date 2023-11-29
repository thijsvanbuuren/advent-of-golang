package usecase

import (
	"fmt"
	"log/slog"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/thijsvanbuuren/advent-of-golang/internal/domain"
	"github.com/thijsvanbuuren/advent-of-golang/internal/usecase/y2023"
)

const _filePathPattern = "./data/%s/%s"

var registry = []domain.Year{
	{
		Year: "2023",
		Exercises: []domain.Day{
			{File: "day1.txt", DoDay: y2023.Day1},
		},
	},
}

func ExecuteExercises(logger *slog.Logger) {
	for _, y := range registry {
		logger.Info("Running", slog.String("Year", y.Year))

		for _, day := range y.Exercises {
			nameList := strings.Split(runtime.FuncForPC(reflect.ValueOf(day.DoDay).Pointer()).Name(), ".")

			logger.Info("Starting", slog.String("Function", nameList[len(nameList)-1]))

			file, err := os.ReadFile(fmt.Sprintf(_filePathPattern, y.Year, day.File))

			if err != nil {
				logger.Error(err.Error())
				continue
			}
			fileLines := strings.Split(string(file), "\n")

			now := time.Now()
			result := day.DoDay(fileLines, logger)

			logger.Info("Result",
				slog.Duration("Time Taken", time.Since(now)),
				slog.Any("Result", result))
		}
	}
}
