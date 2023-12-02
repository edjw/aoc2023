package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFileIntoSlice(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(data), "\n")
}

func sumASlice(slice []int) int {
	total := 0
	for _, number := range slice {
		total += number
	}

	return total
}

type maxColours struct {
	maxRed   int
	maxGreen int
	maxBlue  int
}

type Game struct {
	ID int
	maxColours
}

func getGameID(line string) int {
	GameIDtext := strings.Split((line), ":")[0]

	GameIDStr := strings.ReplaceAll(GameIDtext, "Game ", "")

	GameID, err := strconv.Atoi(GameIDStr)

	if err != nil {
		fmt.Println(err)
	}

	return GameID

}

func getMaxColours(line string) maxColours {

	maxRed := 0
	maxGreen := 0
	maxBlue := 0

	coloursText := strings.Split((string(line)), ":")[1]

	replacedSemiColon := strings.ReplaceAll(coloursText, "; ", ", ")

	allNumberColoursSlice := strings.Split(replacedSemiColon, ",")

	for _, value := range allNumberColoursSlice {

		colourNumber := strings.Split((string(value)), " ")

		// colourNumber[0] is an empty string for some reason that I'm just ignoring here

		number, err := strconv.Atoi(colourNumber[1])
		if err != nil {
			fmt.Println(err)
		}

		colour := colourNumber[2]

		if colour == "red" && number > maxRed {
			maxRed = number
		}
		if colour == "green" && number > maxGreen {
			maxGreen = number
		}
		if colour == "blue" && number > maxBlue {
			maxBlue = number
		}

	}

	return maxColours{
		maxRed:   maxRed,
		maxGreen: maxGreen,
		maxBlue:  maxBlue,
	}

}

func main() {
	lines := readFileIntoSlice("input.txt")
	// lines := readFileIntoSlice("test.txt")

	var allGames []Game

	for _, value := range lines {

		var GameData = new(Game)

		GameData.ID = getGameID(value)
		maxColours := getMaxColours((value))
		GameData.maxRed = maxColours.maxRed
		GameData.maxGreen = maxColours.maxGreen
		GameData.maxBlue = maxColours.maxBlue

		allGames = append(allGames, *GameData)

	}

	var powers []int

	for _, game := range allGames {

		power := game.maxRed * game.maxGreen * game.maxBlue

		powers = append(powers, power)

	}

	fmt.Println(sumASlice(powers))

}
