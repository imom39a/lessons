package main

import "fmt"

func main(){
  fmt.Println("Recursion - n!")
  n := 10
  fmt.Println("final product", factorial(n))
}

func factorial(n int) int{
  fmt.Println("call -> function(",n,")")
  if n == 1 {
    fmt.Println("function(",n,") return -> ", 1)
    return 1
  }
  prod := n * factorial(n-1)
  fmt.Println("function(",n,") return -> ",prod)
  return prod
}