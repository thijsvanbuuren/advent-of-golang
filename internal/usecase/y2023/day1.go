package y2023

import (
	"log/slog"
)

func Day1(input []string, logger *slog.Logger) any {
	logger.Debug("This is Day 1")
	return struct {
		test string
	}{test: "ok"}
}
