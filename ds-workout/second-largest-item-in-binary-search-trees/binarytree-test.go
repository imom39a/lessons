package main

import "fmt"

func main(){

  tree := new(BinaryTree)
  tree.Insert(100)
  tree.Insert(110)
  tree.Insert(105)
  tree.Insert(115)
  tree.Insert(120)
  tree.Insert(130)
  tree.Insert(123)
  tree.Insert(124)
  tree.Insert(128)
  tree.Insert(121)
  tree.Insert(70)
  tree.Insert(90)
  tree.Insert(89)
  tree.Insert(95)

  tree.Walk()
  fmt.Println("Largest",tree.findLargest())
  fmt.Println("Second Largest",tree.findSecondLargest())
}

