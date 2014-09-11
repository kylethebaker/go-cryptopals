package main

import (
  "fmt"
  "encoding/hex"
  "bufio"
  "os"
)

/*--------------------------------------------------------

  scores the plaintext based on english chars and stuff 

--------------------------------------------------------*/

func score_plaintext(pt []byte) byte {
  score := byte(0)

  // iterate over our string and give points to nice chars
  for _, c := range pt {

    // if its a printable ascii char +1
    if c > 31 && c < 127 {
      score += 1

      // we love spaces and e's
      if c == 32 {
        score += 2
      } else if c == 69 || c == 101 {
        score += 1
      }
    }

  }

  return score
}

/*--------------------------------------------------------

  xor a byte array with a single char

--------------------------------------------------------*/

func char_xor(ct []byte, key byte) []byte {

  // allocate a byte array for our 'plaintext'
  pt := make([]byte, len(ct))

  // xor each byte of a ct with the key char
  for i := 0; i < len(ct); i++ {
    pt[i] = ct[i] ^ key
  }

  return pt
}

/*--------------------------------------------------------

  main logic here

--------------------------------------------------------*/

func main() {

  var best_score byte = 0
  var best_char byte
  var best_pt []byte
  var best_ct []byte

  // open our data file
  data, err := os.Open("ch4.dat")
  if err != nil {
    panic(err)
  }
  defer data.Close()

  // read each line in the file and find the most likely pt
  lines := bufio.NewScanner(data)

  for lines.Scan() {
    ct, _ := hex.DecodeString(lines.Text())

    // try all possible chars in ascii range
    for c:= byte(33); c < 127; c++ {
      pt := char_xor(ct, c)
      score := score_plaintext(pt)

      if score > best_score {
        best_score = score
        best_char = c
        best_ct = ct
        best_pt = pt
      }
    }
  }

  fmt.Printf("score:\t%d\n", best_score)
  fmt.Printf("char:\t%c\n", best_char)
  fmt.Printf("cipher:\t%s\n", best_ct)
  fmt.Printf("plain:\t%s\n", best_pt)

}
