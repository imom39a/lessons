package main

import "fmt"

type BinaryTree struct {
  rootNode *BinaryTreeNode
}

func (t *BinaryTree) Insert(data int){
    if t.rootNode == nil {
        t.rootNode = &BinaryTreeNode{Data: data, Direction: "M"}
    } else {      
      t.rootNode.Insert(data)
    }    
}

func (t *BinaryTree) Walk(){
  if t.rootNode != nil {
    fmt.Println("Walking tree")
    t.rootNode.walk()
  } else {
    fmt.Println("Tree is empty")
  }
}

func (t *BinaryTree) findLargest() int{
  return t.rootNode.findLargest()
}

func (t *BinaryTree) findSecondLargest() int{  
    return t.rootNode.findSecondLargest()
}

type BinaryTreeNode struct {
  Data int
  Depth int
  Direction string 
  left *BinaryTreeNode
  right *BinaryTreeNode
}

func (n *BinaryTreeNode) Insert(data int){
  var nodeStack Stack
  nodeStack.Push(n)
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

    var nodeStack Stack
    nodeStack.Push(node)

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

func (node *BinaryTreeNode) findLargest() int {
    
    largestItem := 0
    var largestNodeStack Stack
    largestNodeStack.Push(node)

    for !largestNodeStack.isEmpty(){
        n := largestNodeStack.Pop()

        if n.right != nil {
          largestNodeStack.Push(n.right)
          continue
        }
        largestItem = n.Data
        break
    }

    return largestItem
}

func (node *BinaryTreeNode) findSecondLargest() int {         
    
    largestItem := 0

    var secondLargestNodeStack Stack
    secondLargestNodeStack.Push(node)

    for !secondLargestNodeStack.isEmpty(){
        n := secondLargestNodeStack.Pop() 
        
        if n.left != nil && n.right == nil  {
          largestItem =  n.left.findLargest()
          break
        }

        if n.right != nil && n.right.left == nil && n.right.right == nil {
          largestItem = n.Data
          break
        }

        secondLargestNodeStack.Push(n.right)

    }
    return largestItem
}