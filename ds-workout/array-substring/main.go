package main

import "fmt"

func main() {

  s := "jumanjojojiii"
  x := "jojoo"

  runeOf_s := []rune(s)
  lengthOf_s := len(runeOf_s)
  runeOf_x := []rune(x)
  lengthOf_x := len(runeOf_x)

  wordCheckerSet := make(map[string]struct{})
  wordCheckerSet[string(runeOf_x)] = struct{}{}
  indexOfFound := -1

  if lengthOf_x > lengthOf_x {
    indexOfFound = -1
    print(s,x,indexOfFound)
    return
  }

  for i:= 0; i< lengthOf_s; i++ {
    if i + lengthOf_x > lengthOf_s {
      indexOfFound =  - 1
      break
    }
    _, found := wordCheckerSet[string(runeOf_s[i:i + lengthOf_x])]
    if (found){
      indexOfFound = i
      break
    }
  }

  print(s,x,indexOfFound)
  
  
}

func print(s, x string, indexOfFound int) {
  fmt.Println(s)
  fmt.Println(x)
  fmt.Println("Index Found: ",indexOfFound)
}
