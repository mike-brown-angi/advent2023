package main

import (
	C "from-scratch/advcommon"
	"fmt"
	"log"
	"strconv"
	"strings"
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
	lines, err := C.File2Array()
	if err != nil {
		log.Fatalln(err)
	}

	var fieldArea [][]int
	fieldArea = make([][]int, 1000)
	for i := range fieldArea {
		fieldArea[i] = make([]int, 1000)
	}
	var fieldSpans []fieldSpan
	for _, line := range lines {
		tfs, err := newFieldSpan(line)
		if err != nil {
			log.Fatalln(err)
		}
		fieldSpans = append(fieldSpans, *tfs)
	}

	dangerCnt := 0
	for _, span := range fieldSpans {
		fmt.Printf("%d -> %d, %d -> %d\n", span.fromX, span.toX, span.fromY, span.toY)
		if span.lineType == "x" {
			for i := span.fromY; i <= span.toY; i++ {
				fieldArea[span.fromX][i]++
				if fieldArea[span.fromX][i] == 2 {
					dangerCnt++
				}
			}
		} else if span.lineType == "y" {
			for i := span.fromX; i <= span.toX; i++ {
				fieldArea[i][span.fromY]++
				if fieldArea[i][span.fromY] == 2 {
					dangerCnt++
				}
			}
		}
	}
	//fmt.Println(len(fieldSpans))
	fmt.Println(dangerCnt)
	// fmt.Println(fieldArea[500][500])
	// fmt.Println(fieldSpans)

}
