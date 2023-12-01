package advcommon

import (
  "bufio"
  "os"
)
//
// Function to grab the ../input.txt file and stuff it into a slice of strings.
// 


// File2Array This func must be Exported, Capitalized, and comment added.
func File2Array() ([]string, error) {
  file, err := os.Open("../input.txt")

  fileScan := bufio.NewScanner(file)
  fileScan.Split(bufio.ScanLines)
  var lines []string

  for fileScan.Scan() {
    lines = append(lines, fileScan.Text())
  }

  err = file.Close()

  return lines, err
}
