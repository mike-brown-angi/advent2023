package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	constraints := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	file, err := os.Open("../../inputs/day02.txt")
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
	keepTotal := 0
	for _, line := range lines {
		fmt.Println(line)
		gameNbr := 0
		gameSlice := strings.Split(line, ":")
		gameNbr, err = strconv.Atoi(strings.Split(gameSlice[0], " ")[1])
		if err != nil {
			log.Fatalln(err)
		}
		drawSlice := strings.Split(gameSlice[1], ";")

		fmt.Println(gameNbr)
		fmt.Println(drawSlice)
		possible := true
		for _, drawString := range drawSlice {
			// fmt.Println(drawString)
			drawString = strings.Trim(drawString, " ")
			// var draw = make(map[string]int)
			for _, i := range strings.Split(drawString, ",") {
				i = strings.Trim(i, " ")
				tmpSlice := strings.Split(i, " ")
				tmpCount, err := strconv.Atoi(tmpSlice[0])
				if err != nil {
					log.Fatalln(err)
				}
				if tmpCount > constraints[tmpSlice[1]] {
					possible = false
				}
				fmt.Println(i)
			}
		}
		if possible {
			fmt.Println(gameNbr)
			keepTotal += gameNbr
		}
	}

	fmt.Println(constraints)
	fmt.Println(keepTotal)
}
