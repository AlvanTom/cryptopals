package set1

import (
  "encoding/hex"
  "fmt"
)

func XOR(){
  var first, second string

  fmt.Println("first")
  fmt.Scan(&first)

  fmt.Println("second")
  fmt.Scan(&second)

  if(len(first) != len(second)){
    panic("first length does not match the second")
  }

  decodedFirst, err := hex.DecodeString(first)
  decodedSecond, err := hex.DecodeString(second)

  if err != nil {
    fmt.Println("Error is:", err)
  }

  res := make([]byte, len(decodedFirst))

  for i, value := range(decodedFirst) {
    res[i] = value ^ decodedSecond[i]
  }
  fmt.Println(hex.EncodeToString(res))

}
