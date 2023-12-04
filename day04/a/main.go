package main

import (
	C "advent2023/advcommon"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("../../inputs/day04.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fileScan := bufio.NewScanner(file)
	fileScan.Split(bufio.ScanLines)
	var lines []string

	for fileScan.Scan() {
		lines = append(lines, fileScan.Text())
	}

	err = file.Close()
	totalValue := 0
	for _, line := range lines {
		fmt.Println(line)
		lineSs := strings.Replace(line, "  ", " ", -1)
		linePart := strings.Split(lineSs, ":")
		CardID := strings.Replace(linePart[0], "Card ", "", -1)
		Cards := strings.Split(linePart[1], "|")
		winnerSlice := strings.Split(Cards[0], " ")

		playerSlice := strings.Split(Cards[1], " ")

		s := C.SimpleGeneric(playerSlice, winnerSlice)
		powerValue := len(s) - 1
		tmpValue := 0
		for i := 1; i < len(s); i++ {
			if i == 1 {
				tmpValue = 1
			} else {
				tmpValue = 2 * tmpValue
			}
		}

		fmt.Println("Card: ", CardID, " ", s, tmpValue, powerValue)
		totalValue += tmpValue
	}

	fmt.Println("Total Value: ", totalValue)

}
