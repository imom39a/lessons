package main

import "fmt"

func main(){

  //  valid tree

  tree := new(BinaryTree)
  tree.Insert(100)
  tree.Insert(110)
  tree.Insert(70)  
  tree.Insert(90)
  tree.Insert(89)
  tree.Insert(95)
  tree.Insert(22)
  tree.Insert(2)
  tree.Walk()
  fmt.Println("Is a valid tree ?",tree.isValidSearchTree())

  //  not a valid tree -- should have only distinct keys 
  tree = new(BinaryTree)
  tree.Insert(100)
  tree.Insert(110)
  tree.Insert(70)
  tree.Insert(120)
  tree.Insert(90)
  tree.Insert(89)
  tree.Insert(89)
  tree.Insert(95)
  tree.Insert(22)
  tree.Insert(2)

  tree.Walk()
  fmt.Println("Is a valid tree ?",tree.isValidSearchTree())

  //  not a valid tree - has value 220 which is greater then max bound 

  tree = new(BinaryTree)
  tree.Insert(100)
  tree.Insert(220)
  tree.Insert(110)
  tree.Insert(70)  
  tree.Insert(90)
  tree.Insert(89)
  tree.Insert(95)
  tree.Insert(22)
  tree.Insert(2)
  tree.Walk()
  fmt.Println("Is a valid tree ?",tree.isValidSearchTree())

}
