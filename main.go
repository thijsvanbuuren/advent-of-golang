package main

import (
	"log/slog"

	"github.com/thijsvanbuuren/advent-of-golang/internal/usecase"
)

func main() {

	logger := slog.Default()
	usecase.ExecuteExercises(logger)
}
