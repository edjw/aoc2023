package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")

	var allNumbers []int

	for _, line := range lines {

		onlyNumbers := regexp.MustCompile(`[a-z]`)
		result := onlyNumbers.ReplaceAllString(line, "")

		firstNumberString := string(result[0])
		lastNumberString := string(result[len(result)-1:])
		wholeNumberString := firstNumberString + lastNumberString

		wholeNumberInteger, err := strconv.Atoi(wholeNumberString)
		if err != nil {
			fmt.Println(err)
		}

		allNumbers = append(allNumbers, wholeNumberInteger)
	}

	var total int
	for _, number := range allNumbers {

		total += number
	}

	fmt.Println(total)

}
