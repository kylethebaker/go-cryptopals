package cryptopals

import (
  "crypto/aes"
  "errors"
)

/*--------------------------------------------------------

  xor two equal length byte slices

--------------------------------------------------------*/

func Xor(a, b []byte) []byte {
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

  AES-ECB Encrypt

--------------------------------------------------------*/

func EncryptECB(plaintext []byte, key []byte) []byte {

  // pad our plaintext if neccesary or create a copy
  pad_amount := len(plaintext) % 16

  var padded []byte
  if pad_amount != 0 {
    pad_amount = 16 - pad_amount
    padded = PadPkcs7(plaintext, len(plaintext) + pad_amount)
  } else {
    padded = plaintext
  }

  ciphertext := make([]byte, len(padded))
  bs := len(key)

  crypt, err := aes.NewCipher(key)
  if err != nil { panic(err) }

  for i := 0; i < len(padded); i += bs {
    crypt.Encrypt(ciphertext[i:i+bs], padded[i:i+bs])
  }

  return ciphertext
}

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

  AES-CBC Encrypt

--------------------------------------------------------*/

func EncryptCBC(plaintext, key, iv []byte) []byte {
  var ciphertext []byte

  // pad our plaintext if neccesary or create a copy
  pad_amount := len(plaintext) % 16

  var padded []byte
  if pad_amount != 0 {
    pad_amount = 16 - pad_amount
    padded = PadPkcs7(plaintext, len(plaintext) + pad_amount)
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
    round_block := Xor(prev_block, block)
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
    round_block := Xor(prev_block, cipher_block)
    plaintext = append(plaintext, round_block...)
    prev_block = block
  }

  return plaintext
}

/*--------------------------------------------------------

  pad a byte array using pkcs#7 padding

--------------------------------------------------------*/

func PadPkcs7(str []byte, pad int) []byte {

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

  check for ECB mode

--------------------------------------------------------*/

func IsECB(str []byte) bool {

  // split into 16 byte blocks
  blocks := make([]string, len(str)/16)

  for pos, i := 0, 0; i < len(str); i += 16 {
    blocks[pos] = string(str[i:i+16])
    pos++
  }

  // check for matching blocks
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

  check if two byte slices are equal

--------------------------------------------------------*/

func EqualBytes(a, b []byte) (bool, error) {

  if len(a) != len(b) {
    return false, errors.New("Unequal lengths.")
  }

  for i:=0; i<len(a); i++ {
    if a[i] != b[i] {
      return false, nil
    }
  }

  return true, nil
}

/*--------------------------------------------------------

  remove valid pkcs7 padding

--------------------------------------------------------*/

func RemovePKCS7(s []byte) ([]byte, error) {

  l := len(s)
  pad := s[l-1]
  padding := s[l-int(pad):]

  if int(pad) > 16 {
    return nil, errors.New("Last byte too high")
  }

  if int(pad) == 0 {
    return nil, errors.New("Invalid Padding")
  }

  // return error if invalid chars are found
  for _, c := range padding {
    if c != pad {
      return nil, errors.New("Invalid Padding")
    }
  }

  // return our slice sans padding
  output := s[:l-int(pad)]
  return output, nil
}
