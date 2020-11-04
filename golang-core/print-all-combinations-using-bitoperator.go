package main

import "fmt"

func main(){

  arr := []int{3,7,4,6,5}
  n := len(arr)

  for i:= 0 ; i < (1 << n); i ++ {
      m := 1
    	for j := 0; j < n; j++ {
        //fmt.Println("i & m",i & m)
        if i & m > 0 {
          fmt.Print(arr[j]," ")
        }
        m = m << 1
      }
      fmt.Println()
  }
}