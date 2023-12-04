package main

import (
	C "advent2023/advcommon"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

// The newly-improved calibration document consists of lines of text; each line originally
// contained a specific calibration value that the Elves now need to recover. On each line,
// the calibration value can be found by combining the first digit and the last digit
// (in that order) to form a single two-digit number.
//
// For example
//
// 1abc2
// pqr3stu8vwx
// a1b2c3d4e5f
// treb7uchet
// In this example, the calibration values of these four lines are 12, 38, 15, and 77.
// Adding these together produces 142.
//
// Consider your entire calibration document. What is the sum of all of the calibration values?
// Your calculation isn't quite right. It looks like some of the digits are actually spelled out
// with letters: one, two, three, four, five, six, seven, eight, and nine also count as valid
// "digits".
//
// Equipped with this new information, you now need to find the real first and last digit on
// each line. For example:
//
// two1nine
// eightwothree
// abcone2threexyz
// xtwone3four
// 4nineeightseven2
// zoneight234
// 7pqrstsixteen
// In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76. Adding
// these together produces 281.

func findStringDigits(line string) []int {
	leftIdx := len(line)
	leftValue := 0
	rightIdx := -1
	rightValue := 0

	numberMap := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for key := range numberMap {
		lidx := strings.Index(line, key)
		if lidx != -1 && lidx < leftIdx {
			leftIdx = lidx
			leftValue = numberMap[key]
		}
		ridx := strings.LastIndex(line, key)
		if ridx != -1 && ridx > rightIdx {
			rightIdx = ridx
			rightValue = numberMap[key]
		}
	}
	return []int{leftValue, leftIdx, rightValue, rightIdx}
}

func main() {
	lines, err := C.File2Array(1)
	if err != nil {
		log.Fatalln(err)
	}

	totalSum := 0
	for _, line := range lines {
		stringNumberOutput := findStringDigits(line)
		calibrationValue := 0

		leftIdx := len(line)
		leftValue := 0

		for i, char := range line {
			if unicode.IsDigit(char) {
				leftIdx = i
				val, err := strconv.Atoi(string(char))
				if err != nil {
					log.Fatalln(err)
				}

				leftValue = val
				break
			}
		}

		rightIdx := -1
		rightValue := 0

		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				rightIdx = i
				val, err := strconv.Atoi(string(rune(line[i])))
				if err != nil {
					log.Fatalln(err)
				}

				rightValue = val
				break
			}
		}

		// do we use the first string number or digit number?
		if leftIdx < stringNumberOutput[1] {
			calibrationValue += 10 * leftValue
		} else {
			calibrationValue += 10 * stringNumberOutput[0]
		}
		// do we use the last string number or digit number?
		if rightIdx > stringNumberOutput[3] {
			calibrationValue += rightValue
		} else {
			calibrationValue += stringNumberOutput[2]
		}

		totalSum += calibrationValue
	}
	fmt.Println(totalSum)
}
