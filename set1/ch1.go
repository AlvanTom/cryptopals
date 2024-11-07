package main

import (
  "encoding/hex"
  "encoding/base64"
  "fmt"
)

func main(){
  var input string
  fmt.Println("hello world")
  fmt.Scan(&input)
  bytes, err := hex.DecodeString(input)
  if err != nil {
    fmt.Println("Error is:", err)
  }
  fmt.Println(base64.StdEncoding.EncodeToString(bytes))
}
