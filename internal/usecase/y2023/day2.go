package y2023

import (
	"log/slog"
	"strconv"
	"strings"

	"github.com/thijsvanbuuren/advent-of-golang/internal/domain"
	"github.com/thijsvanbuuren/advent-of-golang/internal/usecase"
	"github.com/thijsvanbuuren/advent-of-golang/internal/util"
)

func init() {
	usecase.AppendExercise(domain.Exercise{
		Year: 2023,
		Day:  2,
		Part1: &domain.Part{
			File: "day2.txt",
			Input: `
			Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
			Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
			Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
			Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
			Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
			`,
			DoDay: Day2Part1,
		},
		Part2: &domain.Part{
			File: "day2.txt",
			Input: `
			Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
			Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
			Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
			Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
			Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
			`,
			DoDay: Day2Part2,
		},
	})
}

type CubeGame struct {
	Nr       int
	maxBlue  int
	maxGreen int
	maxRed   int
	grabs    []Grab
}
type Grab struct {
	add   int
	color string
}

func parse(line string) CubeGame {
	gameSep := strings.Split(line, ":")
	gameNr, _ := strconv.Atoi(strings.Split(gameSep[0], " ")[1])
	rounds := strings.Split(gameSep[1], ";")
	game := CubeGame{Nr: gameNr}

	for _, r := range rounds {
		r = strings.TrimSpace(r)
		for _, grab := range strings.Split(r, ", ") {
			grabSplit := strings.Split(grab, " ")
			add, _ := strconv.Atoi(grabSplit[0])
			game.grabs = append(game.grabs, Grab{color: grabSplit[1], add: add})
			switch grabSplit[1] {
			case "blue":
				game.maxBlue = util.Max(game.maxBlue, add)
			case "green":
				game.maxGreen = util.Max(game.maxGreen, add)
			case "red":
				game.maxRed = util.Max(game.maxRed, add)
			}
		}
	}
	return game
}

func (g *CubeGame) isValid(redLimit, greenLimit, blueLimit int) bool {
	for _, grab := range g.grabs {
		if grab.color == "red" && grab.add > redLimit {
			return false
		}
		if grab.color == "green" && grab.add > greenLimit {
			return false
		}
		if grab.color == "blue" && grab.add > blueLimit {
			return false
		}
	}
	return true
}

func Day2Part1(input []string, logger *slog.Logger) any {
	count := 0
	for _, line := range input {
		game := parse(line)
		if !game.isValid(12, 13, 14) {
			continue
		}
		count += game.Nr
	}
	return count
}

func Day2Part2(input []string, logger *slog.Logger) any {
	count := 0
	for _, line := range input {
		game := parse(line)

		power := game.maxBlue * game.maxGreen * game.maxRed
		count += power
		// fmt.Println(power)

	}
	return count
}
