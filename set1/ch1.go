package set1

import (
  "encoding/hex"
  "encoding/base64"
  "fmt"
)

func HexToBase64(){
  var input string

  fmt.Scan(&input)
  bytes, err := hex.DecodeString(input)

  if err != nil {
    fmt.Println("Error is:", err)
  }

  fmt.Println(base64.StdEncoding.EncodeToString(bytes))
}
