package domain

import (
	"log/slog"
)

type DoDay func(input []string, logger *slog.Logger) any

type Year struct {
	Year      string
	Exercises []Day
}
type Day struct {
	File string
	DoDay
}

type Exercise struct {
	Year  int
	Day   int
	Part1 *Part
	Part2 *Part
}

type Part struct {
	DoDay
	File  string
	Input string
}
