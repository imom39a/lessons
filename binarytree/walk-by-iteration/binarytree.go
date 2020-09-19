package main

import "fmt"

var nodeStack Stack

type BinaryTree struct {
  rootNode *BinaryTreeNode
}

func (t *BinaryTree) Insert(data int){
    if t.rootNode == nil {
        t.rootNode = &BinaryTreeNode{Data: data, Direction: "M"}
    } else {
      nodeStack.Push(t.rootNode)
      t.rootNode.Insert(data)
    }    
}

func (t *BinaryTree) Walk(){
  if t.rootNode != nil {
    fmt.Println("Walking tree")
    nodeStack.Push(t.rootNode)
    t.rootNode.walk()
  } else {
    fmt.Println("Tree is empty")
  }
}

type BinaryTreeNode struct {
  Data int
  Depth int
  Direction string 
  left *BinaryTreeNode
  right *BinaryTreeNode
}

func (n *BinaryTreeNode) Insert(data int){

  for !nodeStack.isEmpty(){    
    nsn := nodeStack.Pop()
    
    if nsn.Data > data {
      if nsn.left == nil {
        nsn.left = &BinaryTreeNode{Data: data, Direction: "L"}
        nsn.left.Depth = nsn.Depth+1
      } else {
        nodeStack.Push(nsn.left)
      }      
    } else {      
      if nsn.right == nil {
        nsn.right = &BinaryTreeNode{Data: data, Direction: "R"}
        nsn.right.Depth = nsn.Depth+1
      } else {
        nodeStack.Push(nsn.right)
      }
    }
  }

}

func (node *BinaryTreeNode) walk(){
    for !nodeStack.isEmpty() {

      n := nodeStack.Pop()

      if n.right != nil {
        nodeStack.Push(n.right)
      }      

      if n.left != nil {
        nodeStack.Push(n.left)
      }

      fmt.Println("Depth,", n.Depth, "Direction", n.Direction, "Data",n.Data)      
    }
}