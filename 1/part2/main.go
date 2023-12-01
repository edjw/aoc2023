package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func putNumberInsideANumberWord(line string) string {

	numberMap := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	processedString := line
	for key, value := range numberMap {
		processedString = strings.Replace(processedString, key, value, -1) // -1 says to replace all occurrences
	}

	return processedString
}

func removeNonNumberCharacters(str string) string {
	onlyNumbersRegex := regexp.MustCompile(`[a-z]`)
	return onlyNumbersRegex.ReplaceAllString(str, "")
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")

	var allNumbers []int

	for _, line := range lines {

		lineWithNumbersInsideNumberWords := putNumberInsideANumberWord(line)

		lineWithOnlyNumbers := removeNonNumberCharacters(lineWithNumbersInsideNumberWords)

		firstNumberString := string(lineWithOnlyNumbers[0])
		lastNumberString := string(lineWithOnlyNumbers[len(lineWithOnlyNumbers)-1])

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
