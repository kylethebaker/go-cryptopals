package main

import (
  "fmt"
  "encoding/hex"
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

  hexstr := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
  rawstr, _ := hex.DecodeString(hexstr)

  winner_score := byte(0)
  winner_plain := make([]byte, len(hexstr) / 2)
  winner_char := byte(0)

  // try all possible chars in ascii range
  for c := byte(33); c < 127; c++ {
    pt := char_xor(rawstr, c)
    score := score_plaintext(pt)

    if score > winner_score {
      winner_score = score
      winner_plain = pt
      winner_char = c
    }
  }

  fmt.Printf("score:\t%d\n", winner_score)
  fmt.Printf("char:\t%c\n", winner_char)
  fmt.Printf("text:\t%s\n", winner_plain)

}
