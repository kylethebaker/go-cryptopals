package main

import (
  "fmt"
  "io/ioutil"
  "encoding/base64"
  "crypto/aes"
)

/*--------------------------------------------------------

  AES-ECB Decrypt

--------------------------------------------------------*/

func DecryptECB(ciphertext []byte, key []byte) []byte  {
  plaintext := make([]byte, len(ciphertext))
  bs := len(key)

  crypt, err := aes.NewCipher(key)
  if err != nil { panic(err) }

  for i := 0; i < len(ciphertext); i += bs {
    crypt.Decrypt(plaintext[i:i+bs], ciphertext[i:i+bs])
  }

  return plaintext
}

/*--------------------------------------------------------

  AES-ECB Encrypt

--------------------------------------------------------*/

func EncryptECB(plaintext []byte, key []byte) []byte {
  ciphertext := make([]byte, len(plaintext))
  bs := len(key)

  crypt, err := aes.NewCipher(key)
  if err != nil { panic(err) }

  for i := 0; i < len(plaintext); i += bs {
    crypt.Encrypt(ciphertext[i:i+bs], plaintext[i:i+bs])
  }

  return ciphertext
}

/*--------------------------------------------------------

  main logic here

--------------------------------------------------------*/

func main() {

  // open our data file
  data, err := ioutil.ReadFile("ch7.dat")
  if err != nil { panic(err) }

  ciphertext := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
  base64.StdEncoding.Decode(ciphertext, data)

  key := []byte("YELLOW SUBMARINE")
  fmt.Printf("%q\n", DecryptECB(ciphertext, key))

  // test our encrypt/decrypt by going back and forth

  fmt.Printf("\n\n\n")
  fmt.Printf("%q\n", DecryptECB(EncryptECB(DecryptECB(ciphertext, key), key), key))

}
