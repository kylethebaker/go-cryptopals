package main

import (
  "encoding/base64"
  "fmt"
  "io/ioutil"
)

/*--------------------------------------------------------

  find hamming distance

--------------------------------------------------------*/

func hamming(a []byte, b []byte) int {
  prod := xor(a, b)
  distance := 0

  for _, b := range prod {
    for b != 0 {
      b &= b - 1
      distance++
    }
  }

  return distance
}

/*--------------------------------------------------------

  expand a key to a certain length

--------------------------------------------------------*/

func expand_key(key []byte, length int) []byte {
  x_key := make([]byte, length)

  for i := range x_key {
    x_key[i] = key[i % len(key)]
  }

  return x_key
}

/*--------------------------------------------------------

  xor two byte slices

--------------------------------------------------------*/

func xor(a []byte, b []byte) []byte {
  prod := make([]byte, len(a))

  for i := range prod {
    prod[i] = a[i] ^ b[i]
  }

  return prod
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

  get two sequential keysize length blocks from an offset

--------------------------------------------------------*/

func get_keysize_blocks(str []byte, ksize int, offset int) ([]byte, []byte) {
  a := make([]byte, ksize)
  b := make([]byte, ksize)

  a = str[offset:offset+ksize]
  b = str[offset+ksize:offset+ksize+ksize]

  return a, b
}

/*--------------------------------------------------------

  scores the plaintext based on english chars and stuff 

--------------------------------------------------------*/

func score_plaintext(pt []byte) int {
  score := 0

  // iterate over our string and give points to nice chars
  for _, c := range pt {

    // printable ascii
    if c > 31 && c < 127 {
      score += 10
    }

    switch {

    // space
    case c == 32:
      score += 40

    // e's
    case c == 69 || c == 101:
      score += 20

    // lowercase alpha
    case c > 96 && c < 123:
      score += 10

    // uppercase alpha
    case c > 64 && c < 91:
      score += 5
    }

  }

  return score
}

/*--------------------------------------------------------

  main logic here

--------------------------------------------------------*/

func main() {

  // open our data file
  data, err := ioutil.ReadFile("ch6.dat")
  if err != nil { panic(err) }

  rawstr := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
  base64.StdEncoding.Decode(rawstr, data)

  // try different keysizes to find the closest approx
  ksize_weight := float32(100)
  ksize_guess := 0

  for ksize := 2; ksize < 40; ksize++ {

    ham_avg := float32(0)

    // take 20 different hamming weights of various sequential blocks
    for offset := 0; offset <= 60; offset += 2 {
      a, b := get_keysize_blocks(rawstr, ksize, offset)
      ham_avg += float32(hamming(a, b))
    }

    // get the average hamming weight, and store the keysize if its the best
    ham_avg /= float32(20)
    ham_avg /= float32(ksize)

    if ham_avg < ksize_weight {
      ksize_weight = ham_avg
      ksize_guess = ksize
    }
  }

  fmt.Printf("Key Size %d with weight %f\n", ksize_guess, ksize_weight)

  // take 1st, 2nd, 3rd, etc, byte of every keylength size blocked
  keysize_blocks := make([][]byte, ksize_guess)

  for i := range keysize_blocks {
    keysize_blocks[i] = make([]byte, len(rawstr)/ksize_guess)

    for j := 0; j < len(keysize_blocks[i]); j++ {
      keysize_blocks[i][j] = rawstr[i+(ksize_guess * j)]
    }
  }

  // at this point we might (hopefully) have our keysize blocks
  // lets try to single char xor each one and get the char
  // then put them together and get the key

  key := make([]byte, ksize_guess)

  for pos, block := range keysize_blocks {
    var best_score int
    var best_char byte

    // try all possible chars in ascii range
    for c:= byte(32); c < 127; c++ {
      pt := char_xor(block, c)
      score := score_plaintext(pt)

      if score > best_score {
        best_score = score
        best_char = c
      }
    }

    key[pos] = best_char
  }

  fmt.Printf("%s\n", key)

  plaintext := xor(rawstr, expand_key(key, len(rawstr)))

  fmt.Printf("%s\n", plaintext)

}
