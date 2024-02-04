package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func getCalibrationValue(line string) int64 {
	var buffer []string
	var digits []string

	for _, char := range line {
		buffer = append(buffer, string(char))
		bufferStr := strings.Join(buffer, "")

		if unicode.IsDigit(char) {
			digits = append(digits, string(char))
		} else if strings.HasSuffix(bufferStr, "one") {
			digits = append(digits, "1")
		} else if strings.HasSuffix(bufferStr, "two") {
			digits = append(digits, "2")
		} else if strings.HasSuffix(bufferStr, "three") {
			digits = append(digits, "3")
		} else if strings.HasSuffix(bufferStr, "four") {
			digits = append(digits, "4")
		} else if strings.HasSuffix(bufferStr, "five") {
			digits = append(digits, "5")
		} else if strings.HasSuffix(bufferStr, "six") {
			digits = append(digits, "6")
		} else if strings.HasSuffix(bufferStr, "seven") {
			digits = append(digits, "7")
		} else if strings.HasSuffix(bufferStr, "eight") {
			digits = append(digits, "8")
		} else if strings.HasSuffix(bufferStr, "nine") {
			digits = append(digits, "9")
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
