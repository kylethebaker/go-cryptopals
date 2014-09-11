package main

import (
  "fmt"
  "strconv"
  "strings"
  "crypto/rand"
  "./cryptopals"
)

// global key
var KEY []byte

// our uid counter
var NEXT_UID int = 10

// our user object
type UserObject struct {
  email string
  uid int
  role string
}

/*--------------------------------------------------------

  our UserObject methods

--------------------------------------------------------*/

func NewUser(user_email string) *UserObject {
  if strings.ContainsAny(user_email, "&%=") {
    panic("email cannot contains encoding metachars")
  }
  p := new(UserObject)
  p.email = user_email
  p.uid = NEXT_UID
  p.role = "user"
  NEXT_UID += 5
  return p
}

func (u *UserObject) Encoded() string {
  return "email="+u.email+"&uid="+strconv.Itoa(u.uid)+"&role="+u.role
}

func EncryptQuery(query string) (cipher []byte) {
  cipher = cryptopals.EncryptECB([]byte(query), KEY)
  return
}

func DecryptQuery(cipher []byte) (query string) {
  query_bytes := cryptopals.DecryptECB(cipher, KEY)

  query = string(query_bytes[:])
  return
}

/*--------------------------------------------------------

  main logic here

--------------------------------------------------------*/

func main() {

  // fill our key with random bytes
  KEY = make([]byte, 16)
  _, _ = rand.Read(KEY)

  // this will give us only the users role isolated in the last block
  role_block := NewUser("khb@gmail.com")
  role_cipher := EncryptQuery(role_block.Encoded())
  role_isolate := role_cipher[32:]

  // this will give us a block (2nd) with just the word admin in it
  admin_block := NewUser("asd@gmail.admin           ")
  admin_cipher := EncryptQuery(admin_block.Encoded())
  admin_isolate := admin_cipher[16:32]

  for i := 0; i<16; i++ {
    role_isolate[i] = admin_isolate[i]
  }

  fmt.Printf("%v\n\n", DecryptQuery(role_cipher))

}
