package main

import (
  "fmt"
  "crypto/rand"
  "./cryptopals"
  "encoding/base64"
  "io/ioutil"
)

// our global key and secret
var KEY []byte
var SECRET []byte

/*--------------------------------------------------------

  ECB black box

--------------------------------------------------------*/

func oracle(input []byte) []byte {
  plaintext := append(input, SECRET...)
  return cryptopals.EncryptECB(plaintext, KEY)
}

/*--------------------------------------------------------

  main logic here

--------------------------------------------------------*/

func main() {

  // get our global random key
  KEY = make([]byte, 16)
  _, _  = rand.Read(KEY)

  // get our global secret
  data, err := ioutil.ReadFile("ch12.dat")
  if err != nil { panic(err) }

  SECRET = make([]byte, base64.StdEncoding.DecodedLen(len(data)))
  base64.StdEncoding.Decode(SECRET, data)

  // detect block size by calcing difference when our cipher len changes
  bs := 0
  for i:=1; i<33; i++ {

    str := make([]byte, i)
    for j:=0; j<len(str); j++ {
      str[j] = 'a'
    }

    oracle_len := len(oracle(str))
    if bs != 0 && oracle_len > bs {
      bs = oracle_len - bs
      break
    }
    bs = oracle_len
  }

  fmt.Printf("Block Size:\t%d\n", bs)

  // feed the oracle two blocks of known text to test for ecb
  test_pt := make([]byte, bs*2)

  for i:=0; i<len(test_pt); i++ {
    test_pt[i] = 'a'
  }

  is_ecb := cryptopals.IsECB(oracle(test_pt))
  fmt.Printf("Is it ECB?:\t%t\n", is_ecb)

  // get our ciphertext with just the secret
  raw_ct := oracle([]byte(""))
  secret_block_count := len(raw_ct) / bs

  // this will hold our solved secret
  known_chars := make([]byte, 0, len(raw_ct))

  // iterate over the blocks and solve each block one at a time
  for b:=1; b<secret_block_count+1; b++ {

    // iterate over the bytes of the current block
    for c:=1; c<bs+1; c++ {

      // populate our short block for this round
      short_block := make([]byte, (bs*b)-len(known_chars)-1)

      for i:=0; i<len(short_block); i++ {
        short_block[i] = 'a'
      }

      // get our ciphertext E(short_block || SECRET)
      ciphertext := oracle(short_block)

      // isolate the current block we're trying to solve
      round_cipher := ciphertext[0:bs*b]

      // create our control blocks and compare to our round block for equality
      // start at 10 for newline, then jump up to 32 for alpha range after first pass
      for i:=byte(10); i<127; i++ {

        // short_block || known_chars || unknown char
        control_block := make([]byte, 0)
        control_block = append(control_block, short_block...)
        control_block = append(control_block, known_chars...)
        control_block = append(control_block, i)

        // get the ciphertext of our control block
        control_cipher := oracle(control_block)
        control_cipher = control_cipher[0:bs*b]

        // if our control and round ciphers are equal, then the discovered char is char(i)
        equal, err := cryptopals.EqualBytes(round_cipher, control_cipher)
        if err != nil { panic(err) }

        if equal {
          // add the discovered char to our list, and break
          known_chars = append(known_chars, i)
          fmt.Printf("%c", i)
          break
        }

        // jump up to alpha range for the rest of the iterations
        if i == byte(10) {
          i=31
        }

      }
    }
  }

}
