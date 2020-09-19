package main

func main(){

  tree := new(BinaryTree)
  tree.Insert(100)
  tree.Insert(110)
  tree.Insert(70)
  tree.Insert(90)
  tree.Insert(89)
  tree.Insert(95)

  tree.Walk()

}