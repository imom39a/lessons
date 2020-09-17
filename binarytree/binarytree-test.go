package main

import "fmt"

func main(){
  
  tree := new(BinaryTree)
  tree.Insert(100)
  tree.Insert(110)
  tree.Insert(70)
  tree.Insert(90)
  tree.Insert(95)
    
  fmt.Println(tree.Search(70))
  fmt.Println(tree.Search(95))
  fmt.Println(tree.Search(25))

  tree.Walk()

}