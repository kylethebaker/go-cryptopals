package main

import (
  "fmt"
  "./cryptopals"
)

func main() {
  base := []byte("ICE ICE BABY!!!")

  gPad1 := []byte{byte(4), byte(4), byte(4), byte(4)}
  gPad2 := []byte{byte(1)}
  gPad3 := []byte{byte(8), byte(8), byte(8), byte(8), byte(8), byte(8), byte(8), byte(8)}

  bPad1 := []byte{byte(2)}
  bPad2 := []byte{byte(1), byte(2), byte(3), byte(4)}
  bPad3 := []byte{byte(5), byte(5), byte(5), byte(5)}

  gStr1 := make([]byte,0)
  gStr2 := make([]byte,0)
  gStr3 := make([]byte,0)
  bStr1 := make([]byte,0)
  bStr2 := make([]byte,0)
  bStr3 := make([]byte,0)

  gStr1 = append(gStr1, base...)
  gStr1 = append(gStr1, gPad1...)

  gStr2 = append(gStr2, base...)
  gStr2 = append(gStr2, gPad2...)

  gStr3 = append(gStr3, base...)
  gStr3 = append(gStr3, gPad3...)

  bStr1 = append(bStr1, base...)
  bStr1 = append(bStr1, bPad1...)

  bStr2 = append(bStr2, base...)
  bStr2 = append(bStr2, bPad2...)

  bStr3 = append(bStr3, base...)
  bStr3 = append(bStr3, bPad3...)


  fmt.Println("gStr1")
  gStr1_r, err := cryptopals.RemovePKCS7(gStr1)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Printf("good padding for gStr1\n")
    fmt.Printf("%v\n\n", gStr1_r)
  }
  fmt.Println("-------------------\n")

  fmt.Println("gStr2")
  gStr2_r, err := cryptopals.RemovePKCS7(gStr2)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Printf("good padding for gStr2\n")
    fmt.Printf("%v\n\n", gStr2_r)
  }
  fmt.Println("-------------------\n")

  fmt.Println("gStr3")
  gStr3_r, err := cryptopals.RemovePKCS7(gStr3)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Printf("good padding for gStr3\n")
    fmt.Printf("%v\n\n", gStr3_r)
  }
  fmt.Println("-------------------\n")

  fmt.Println("bStr1")
  bStr1_r, err := cryptopals.RemovePKCS7(bStr1)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Printf("good padding for bStr1\n")
    fmt.Printf("%v\n\n", bStr1_r)
  }
  fmt.Println("-------------------\n")

  fmt.Println("bStr2")
  bStr2_r, err := cryptopals.RemovePKCS7(bStr2)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Printf("good padding for bStr2\n")
    fmt.Printf("%v\n\n", bStr2_r)
  }
  fmt.Println("-------------------\n")

  fmt.Println("bStr3")
  bStr3_r, err := cryptopals.RemovePKCS7(bStr3)
  if bStr3 != nil {
    fmt.Println(err)
  } else {
    fmt.Printf("good padding for bStr3\n")
    fmt.Printf("%v\n\n", bStr3_r)
  }
  fmt.Println("-------------------\n")

}
