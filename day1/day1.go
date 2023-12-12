package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileLines := openFile()
	var result int

	for _, line := range fileLines {
		c, _ := CalibrationValues(line)
		result += c
	}

	fmt.Printf("final result: %v", result)
}

func CalibrationValues(input string) (int, error) {
	input = replaceStrings(input)

	split := strings.Split(input, "")
	var n []int
	for _, v := range split {
		if converted, err := strconv.Atoi(v); err == nil {
			n = append(n, converted)
		}
	}

	firstDigit := strconv.Itoa(n[0])
	lastDigit := strconv.Itoa(n[len(n)-1])

	result, err := strconv.Atoi(firstDigit + lastDigit)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func replaceStrings(input string) string {
	r := strings.NewReplacer(
		"oneight", "oneeight",
		"twone", "twoone",
		"threeight", "threeeight",
		"fiveight", "fiveeight",
		"sevenine", "sevennine",
		"eightwo", "eighttwo",
		"nineight", "nineeight",
	)
	input = r.Replace(input)

	replacer := strings.NewReplacer(
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9")

	input = replacer.Replace(input)
	return input
}

func openFile() []string {
	input, _ := os.Open("part2Data.txt")
	defer input.Close()

	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines
}
