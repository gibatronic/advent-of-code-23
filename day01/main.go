package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func getCalibrationValue(line string) int64 {
	var digits []string

	for _, char := range line {
		if unicode.IsDigit(char) {
			digits = append(digits, string(char))
		}
	}

	first := digits[0]
	last := digits[len(digits)-1]
	calibrationValue, err := strconv.ParseInt(first+last, 10, 64)

	if err != nil {
		log.Panic(err)
	}

	return calibrationValue
}

func main() {
	input, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()
	scanner := bufio.NewScanner(input)
	var sum int64

	for scanner.Scan() {
		line := scanner.Text()
		sum += getCalibrationValue(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}
