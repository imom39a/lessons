package main

import "fmt"

var mem map[int]int

func main(){
  fmt.Println("Print nth fibonachi number")
  mem = make(map[int]int)
  n := 5
  fmt.Println(fib(n))
}


func fib(n int) int{  
  if n ==0 || n == 1 {
    return 1
  }
  prevprev := 0
  prev := 1 
  current := 0
  for i:=2 ; i<=n; i++{
    current = prevprev + prev
    prevprev = prev
    prev = current
  }
  return current
}