package main

import (
	C "advent2023/advcommon"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

// The newly-improved calibration document consists of lines of text; each line originally
// contained a specific calibration value that the Elves now need to recover. On each line,
// the calibration value can be found by combining the first digit and the last digit
// (in that order) to form a single two-digit number.
//
// For example:
//
// 1abc2
// pqr3stu8vwx
// a1b2c3d4e5f
// treb7uchet
// In this example, the calibration values of these four lines are 12, 38, 15, and 77.
// Adding these together produces 142.
//
// Consider your entire calibration document. What is the sum of all of the calibration values?

func main() {
	lines, err := C.File2Array(1)
	if err != nil {
		log.Fatalln(err)
	}
	totalSum := 0
	for _, line := range lines {
		calibrationValue := 0
		println(line)
		re := regexp.MustCompile("[0-9]")
		valueArray := re.FindAllString(line, -1)
		// fmt.Println(valueArray)
		calibrationString := valueArray[0] + valueArray[len(valueArray)-1]
		calibrationValue, err = strconv.Atoi(calibrationString)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(calibrationValue)
		totalSum = totalSum + calibrationValue
	}
	fmt.Println(totalSum)
}
