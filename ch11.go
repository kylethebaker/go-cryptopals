package main

import (
  "fmt"
  "math/rand"
  "time"
  randbytes "crypto/rand"
  "./cryptopals"
)

/*--------------------------------------------------------

  ECB / CBC black box

--------------------------------------------------------*/

func oracle(input []byte) []byte {

  // seed our rand generator
  rand.Seed(time.Now().UnixNano())

  // make random key
  key := make([]byte, 16)
  _, _ = randbytes.Read(key)

  // make random prefix/suffix bytes
  prefix := make([]byte, rand.Intn(5) + 5)
  suffix := make([]byte, rand.Intn(5) + 5)

  _, _ = randbytes.Read(prefix)
  _, _ = randbytes.Read(suffix)

  pt := append(prefix, input...)
  pt = append(pt, suffix...)

  // ecb or cbc?
  coin := rand.Intn(2)

  var ciphertext []byte

  // do cbc
  if coin == 0 {

    // make iv
    iv := make([]byte, 16)
    _, err := randbytes.Read(iv)
    if err != nil { panic(err) }

    ciphertext = cryptopals.EncryptCBC(pt, key, iv)
    fmt.Println("Actual Mode: CBC")

  // do ecb
  } else {
    ciphertext = cryptopals.EncryptECB(pt, key)
    fmt.Println("Actual Mode: ECB")
  }

  return ciphertext

}

/*--------------------------------------------------------

  main logic here

--------------------------------------------------------*/

func main() {

  // trick it with our input
  input := []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

  // detect the mode of 10 different oracle passes
  for i:=0; i<=10; i++ {
    output := oracle(input)

    if cryptopals.IsECB(output) {
      fmt.Println("Detected Mode: ECB\n")
    } else {
      fmt.Println("Detected Mode: CBC\n")
    }
  }


}
