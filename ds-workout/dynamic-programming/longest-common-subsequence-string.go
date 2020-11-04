package main

import "fmt"

var str1 []rune = []rune("asddefg")
var str2 []rune = []rune("addef")
var mem [][]int

func main() {

	fmt.Println("Longest common subsequence string")
	mem = make([][]int, len(str1)+1)
	for i := range mem {
		mem[i] = make([]int, len(str2)+1)
	}

	for i := len(str1) - 1; i >= 0 ; i-- {
		for j := len(str2) - 1; j >= 0; j-- {
			if str1[i] == str2[j] {
				mem[i][j] = 1 + mem[i+1][j+1]
			} else {
				v1 := mem[i][j+1]
				v2 := mem[i+1][j]
				if v1 > v2 {
					mem[i][j] = v1
				} else {
					mem[i][j] = v2
				}
			}
		}
	}

  i := 0
  j := 0
  str := make([]rune,0)
  for  i < len(str1) && j < len(str2) {        
    if str1[i] == str2[j] {
      str = append(str, str1[i])
      i++
      j++
    } else if mem[i+1][j] >= mem[i][j+1]{
      i++
    } else {
      j++
    }    
  }

  for i := range mem {
    fmt.Println(mem[i])
  }

  fmt.Println("substring --> ", string(str))

}
