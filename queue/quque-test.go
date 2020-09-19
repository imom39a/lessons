package main

import "fmt"

func main(){

  fmt.Println("Queue Test")
  
  var queue Queue

  queue.Enqueue(1)
  queue.Enqueue(2)
  queue.Enqueue(3)
  fmt.Println("item",queue.Dequeue())
  fmt.Println("item",queue.Dequeue())
  fmt.Println("item",queue.Dequeue())

}