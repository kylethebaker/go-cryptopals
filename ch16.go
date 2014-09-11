package main

import (
  "./cryptopals"
  "crypto/rand"
  "fmt"
)

var KEY []byte

/*--------------------------------------------------------

  The oracle

--------------------------------------------------------*/

func EncryptionOracle(input []byte) []byte {

  pre := []byte("comment1=cooking%20MCs;userdata=")
  post := []byte(";comment2=%20like%20a%20pound%20of%20bacon")
}

func DecryptionOracle(input []byte) []byte {
}

func main() {

  KEY = make([]byte 16)
  _, _ = rand.Rand(KEY)
}
