package advcommon

import (
  "bufio"
  "log"
  "os"
  "strconv"
  "strings"
)
//
// Function to grab the ../input.txt file and stuff it into a slice of Ints.
// 


// File2IntSlice This func must be Exported, Capitalized, and comment added.
func File2IntSlice() ([][]int, error) {
  var retSlice [][]int
  file, err := os.Open("../input.txt")

  fileScan := bufio.NewScanner(file)
  fileScan.Split(bufio.ScanLines)
  var lines []string

  for fileScan.Scan() {
    lines = append(lines, fileScan.Text())
  }

  err = file.Close()
  for _, line := range lines {
    tmpSlice := strings.Split(line, ",")
    nbrSlice := make([]int, len(tmpSlice))
    for i := range tmpSlice {
      nbrSlice[i], err = strconv.Atoi(tmpSlice[i])
      if err != nil {
        log.Fatalln(err)
      }
    }
    retSlice = append(retSlice, nbrSlice)
  }
  return retSlice, err
}
