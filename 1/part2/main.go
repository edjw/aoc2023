package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
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
		processedString = strings.ReplaceAll(processedString, key, value)
	}

	return processedString
}

func removeNonNumberCharacters(str string) string {
	onlyNumbersRegex := regexp.MustCompile(`[a-z]`)
	return onlyNumbersRegex.ReplaceAllString(str, "")
}

func main() {

<<<<<<< HEAD
	lines := readFileIntoSlice("input.txt")
=======
	lines := readFileIntoSplice("input.txt")
>>>>>>> 8b4207b8038fd083a7cb166163617f8424daacc5

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

	total := sumASlice(allNumbers)

	fmt.Println(total)

}
