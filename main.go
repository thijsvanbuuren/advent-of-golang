package main

import (
	"log/slog"

	"github.com/thijsvanbuuren/advent-of-golang/internal/usecase"
	"github.com/thijsvanbuuren/advent-of-golang/internal/usecase/y2023"
)

func main() {

	_ = y2023.NILL
	logger := slog.Default()
	usecase.ExecuteExercises(logger)
}
