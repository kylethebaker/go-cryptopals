package main

import (
  "fmt"
  "encoding/hex"
  "bufio"
  "os"
)

/*--------------------------------------------------------

  split into 16 byte blocks

--------------------------------------------------------*/

func split_into_blocks(str []byte) []string {
  blocks := make([]string, len(str)/16)

  for pos, i := 0, 0; i < len(str); i += 16 {
    blocks[pos] = string(str[i:i+16])
    pos++
  }

  return blocks
}

/*--------------------------------------------------------

  check for matching blocks

--------------------------------------------------------*/

func has_matching_blocks(blocks []string) bool {
  for i, b1 := range blocks {
    for j, b2 := range blocks {
      if i != j && b1 == b2 {
        return true
      }
    }
  }
  return false
}

/*--------------------------------------------------------

  main logic here

--------------------------------------------------------*/

func main() {

  // open our data file
  data, err := os.Open("ch8.dat")
  if err != nil { panic(err) }
  defer data.Close()

  // read each line in the file and find the most likely pt
  lines := bufio.NewScanner(data)

  for lines.Scan() {
    ct, _ := hex.DecodeString(lines.Text())
    if has_matching_blocks(split_into_blocks(ct)) {
      ct_hex := hex.EncodeToString(ct)
      fmt.Printf("%s\n", ct_hex)
    }
  }
}
