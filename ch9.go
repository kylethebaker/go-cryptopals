package main

import (
  "fmt"
)

/*--------------------------------------------------------

  pad a byte array using pkcs#7 padding

--------------------------------------------------------*/

func pad_pkcs7(str []byte, pad int) []byte {

  if len(str) > pad {
    panic("String is longer than padding")
  } else if len(str) == pad {
    return str
  }

  pad_slice := make([]byte, pad - len(str))
  for i := range pad_slice {
    pad_slice[i] = byte(pad - len(str))
  }

  padded := append(str, pad_slice...)

  return padded

}

/*--------------------------------------------------------

  main logic here

--------------------------------------------------------*/

func main() {
  str := []byte("YELLOW SUBMARINE")
  fmt.Printf("%q\n", pad_pkcs7(str, 20))
}
