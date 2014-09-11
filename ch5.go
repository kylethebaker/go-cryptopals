package main

import (
  "fmt"
  "encoding/hex"
)


/*--------------------------------------------------------

  repeating key xor

--------------------------------------------------------*/

func xor(key []byte, str []byte) []byte {
  result := make([]byte, len(str))

  for i, _ := range str {
    result[i] = str[i] ^ key[i % len(key)]
  }

  return result

}

/*--------------------------------------------------------

  main logic here

--------------------------------------------------------*/

func main() {
  str := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
  key := "ICE"

  rawcipher := xor([]byte(key), []byte(str))

  fmt.Printf("%s\n", hex.EncodeToString(rawcipher))

}
