package advcommon

import (
	"bufio"
	"os"
	"strconv"
)

//
// Function to grab the ../input.txt file and stuff it into a slice of strings.
//

// File2Array This func must be Exported, Capitalized, and comment added.
func File2Array(day int) ([]string, error) {
	dayStr := ""
	genDirStr := "../../inputs/day"
	if day < 10 {
		dayStr = genDirStr + "0" + strconv.Itoa(day) + ".txt"
	} else {
		dayStr = genDirStr + strconv.Itoa(day) + ".txt"
	}
	file, err := os.Open(dayStr)

	fileScan := bufio.NewScanner(file)
	fileScan.Split(bufio.ScanLines)
	var lines []string

	for fileScan.Scan() {
		lines = append(lines, fileScan.Text())
	}

	err = file.Close()

	return lines, err
}
