package set1

func XOR(first,second []byte) []byte{

  if(len(first) != len(second)){
    panic("first length does not match the second")
  }

  res := make([]byte, len(first))

  for i, value := range(first) {
    res[i] = value ^ second[i]
  }
  return res
  // fmt.Println(hex.EncodeToString(res))
}
