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

	fmt.Println(result)
}

func CalibrationValues(input string) (int, error) {
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

func openFile() []string {
	input, _ := os.Open("input.txt")
	defer input.Close()

	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines
}
