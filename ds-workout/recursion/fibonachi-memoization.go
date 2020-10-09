package main

import "fmt"

var mem map[int]int

func main(){
  fmt.Println("Print nth fibonachi number")
  mem = make(map[int]int)
  n := 10
  fmt.Println(fib(n))
}

func fib(n int) int{  
  if n == 1 || n == 0 {
    return n
  }
  
  if _, ok := mem[n]; ok {
    fmt.Println("retrive from mem for",n,"value",mem[n])
    return mem[n]
  }
  fmt.Println("calculating for value",n)
  mem[n]= fib(n-1) + fib(n-2)  
  return mem[n]
}