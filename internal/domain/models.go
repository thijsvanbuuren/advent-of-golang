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
