package main

import "fmt"

var str1 []rune = []rune("asddfg")
var str2 []rune = []rune("addf")
var mem [][]int;

func main(){

  fmt.Println("Longest common subsequence string count")
  mem = make([][]int,len(str1)+1)
  for i := range mem {
    mem[i] = make([]int, len(str2)+1)
  }
  
  fmt.Println(lcs(0,0))
  // for _, v := range mem {
  //   fmt.Println(v)
  // }

}

func lcs(i, j int) int {

  if i == len(str1) || j == len(str2) {
    return 0
  }

  if str1[i] == str2[j] {
    mem[i][j] = 1 + lcs(i+1, j+1)
    return mem[i][j]
  } else {
    if (mem[i][j+1] == 0){
      mem[i][j+1] = lcs(i, j+1)
    }
    if (mem[i+1][j] == 0){
      mem[i+1][j] = lcs(i+1, j)
    }
    
    if mem[i][j+1] > mem[i+1][j] {
      return mem[i][j+1]
    } else {
      return mem[i+1][j]
    }
  }

}