package main

import "fmt"

func main(){

  var myStack Stack

  fmt.Println("Dumping stack")
  myStack.Dump()
  myStack.Push(10)
  myStack.Push(20)
  myStack.Push(30)
  fmt.Println("Dumping stack")
  myStack.Dump()
  fmt.Println("Popped item",myStack.Pop())
  fmt.Println("Popped item",myStack.Pop())
  fmt.Println("Popped item",myStack.Pop())
  fmt.Println("Popped item",myStack.Pop())

}