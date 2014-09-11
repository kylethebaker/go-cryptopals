package main

import (
  "fmt"
  "io/ioutil"
  "encoding/base64"
  "crypto/aes"
)

/*--------------------------------------------------------

  xor two equal length byte slices

--------------------------------------------------------*/

func xor(a, b []byte) []byte {
  if len(a) != len(b) {
    panic("can't xor byte slices of unequal length")
  }

  prod := make([]byte, len(a))
  for i := range a {
    prod[i] = a[i] ^ b[i]
  }

  return prod
}

/*--------------------------------------------------------

  AES-ECB Encrypt a single block

--------------------------------------------------------*/

func EncryptBlockECB(block, key []byte) []byte {

  // make sure its a valid block size
  if len(block) != 16 {
    panic("blocks must be 16 bytes long")
  }

  ciphertext := make([]byte, len(block))

  crypt, err := aes.NewCipher(key)
  if err != nil { panic(err) }

  crypt.Encrypt(ciphertext, block)

  return ciphertext
}

/*--------------------------------------------------------

  AES-ECB Decrypt a single block

--------------------------------------------------------*/

func DecryptBlockECB(block, key []byte) []byte  {

  // make sure its a valid block size
  if len(block) != 16 {
    panic("blocks must be 16 bytes long")
  }

  plaintext := make([]byte, len(block))

  crypt, err := aes.NewCipher(key)
  if err != nil { panic(err) }

  crypt.Decrypt(plaintext, block)

  return plaintext
}

/*--------------------------------------------------------

  AES-CBC Encrypt

--------------------------------------------------------*/

func EncryptCBC(plaintext, key, iv []byte) []byte {
  var ciphertext []byte

  // pad our plaintext if neccesary or create a copy
  pad_amount := len(plaintext) % 16

  var padded []byte
  if pad_amount != 0 {
    pad_amount = 16 - pad_amount
    padded = pad_pkcs7(plaintext, len(plaintext) + pad_amount)
  } else {
    padded = plaintext
  }

  // split into 16 byte blocks
  blocks := make([][]byte, len(padded)/16)
  for i, pos := 0, 0; i < len(padded); i += 16 {
    blocks[pos] = padded[i:i+16]
    pos++
  }

  prev_block := iv

  for _, block := range blocks {
    round_block := xor(prev_block, block)
    cipher_block := EncryptBlockECB(round_block, key)
    ciphertext = append(ciphertext, cipher_block...)
    prev_block = cipher_block
  }

  return ciphertext

}

/*--------------------------------------------------------

  AES-CBC Decrypt

--------------------------------------------------------*/

func DecryptCBC(ciphertext, key, iv []byte) []byte  {
  var plaintext []byte

  // make sure our ciphertext is valid (16 byte multiple)
  if len(ciphertext) % 16 != 0 {
    panic("ciphertext length must be multiple of block length (16 byte)")
  }


  // split into 16 byte blocks
  blocks := make([][]byte, len(ciphertext)/16)
  for i, pos := 0, 0; i < len(ciphertext); i += 16 {
    blocks[pos] = ciphertext[i:i+16]
    pos++
  }

  prev_block := iv

  for _, block := range blocks {
    cipher_block := DecryptBlockECB(block, key)
    round_block := xor(prev_block, cipher_block)
    plaintext = append(plaintext, round_block...)
    prev_block = block
  }

  return plaintext
}

/*--------------------------------------------------------

  pad a byte array using pkcs#7 padding

--------------------------------------------------------*/

func pad_pkcs7(str []byte, pad int) []byte {

  if len(str) > pad {
    panic("string is longer than padding")
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

  // open our data file
  data, err := ioutil.ReadFile("ch10.dat")
  if err != nil { panic("error opening file") }

  ct := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
  base64.StdEncoding.Decode(ct, data)

  key := []byte("YELLOW SUBMARINE")
  iv := make([]byte, 16)

  fmt.Printf("Testing Decrypt(Encrypt())...\n")
  plaintext := []byte("this is a really long string just testing my shitty cbc mode")
  fmt.Printf("%q\n\n", DecryptCBC(EncryptCBC(plaintext, key, iv), key, iv))

  fmt.Printf("Testing Decrypt(ch10.dat))...\n")
  fmt.Printf("%q\n", DecryptCBC(ct, key, iv))
}
