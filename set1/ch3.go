package set1

import (
	"encoding/hex"
	"fmt"
)

func SingleByteCipher(){
  input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
  decodedInput, err := hex.DecodeString(input)
  if err != nil{
    fmt.Println("Error is: ", err)
  }

  max := 0
  var res [][]byte
  var char []int
  for i := range(0xFF + 1){
    iter := make([]byte, len(decodedInput))
    total := 0
    for b := range(len(decodedInput)){
      iter[b] = decodedInput[b] ^ byte(i)
      if isChar(iter[b]){
        total++
      }
    }
    if total > max {
      max = total
      res = make([][]byte,0)
      res = append(res, iter)
      char = []int{i}
    }else if total == max{
      res = append(res, iter)
      char = append(char, i)
    }
  }
  for i, c := range char{
    fmt.Println(string(c),":", string(res[i]))

  }
}

func isChar(c byte) bool {
  return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}
