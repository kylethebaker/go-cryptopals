package main

import (
  "fmt"
  "encoding/hex"
  "encoding/base64"
)

/*--------------------------------------------------------

  Convert hex string to base64 byte array

--------------------------------------------------------*/

func hex_to_b64(hexstr string) []byte {

  // convert our hex string to byte array
  rawstr, err := hex.DecodeString(hexstr)
  if err != nil {
    panic(err)
  }

  // allocate a byte array the size of our encoded value
  encoded := make([]byte, base64.StdEncoding.EncodedLen(len(rawstr)))

  // store the encoded value into our previous allocated array
  base64.StdEncoding.Encode(encoded, rawstr)

  return encoded
}

/*--------------------------------------------------------

  xor two equal length byte arrays (assumes equal length)

--------------------------------------------------------*/

func xor(str_a []byte, str_b []byte) []byte {

  // allocate a byte array for the size of our xored prod
  prod := make([]byte, len(str_a))

  // xor each byte from the arrays
  for i := 0; i < len(prod); i++ {
    prod[i] = str_a[i] ^ str_b[i]
  }

  return prod
}

/*--------------------------------------------------------

  main logic here

--------------------------------------------------------*/

func main() {
  s1 := "1c0111001f010100061a024b53535009181c"
  s2 := "686974207468652062756c6c277320657965"
  raws1, _ := hex.DecodeString(s1)
  raws2, _ := hex.DecodeString(s2)

  rawprod := xor(raws1, raws2)
  prod := hex.EncodeToString(rawprod)
  fmt.Println(prod)
}
