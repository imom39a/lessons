package main

import "fmt"
import "math"

var nodeStack Stack

type BinaryTree struct{
  rootNode *BinaryTreeNode
}

type BinaryTreeNode struct{
  Data int
  Depth int
  left *BinaryTreeNode
  right *BinaryTreeNode
}

func (t *BinaryTree) IsSuperBalanced() bool {

  isSuperBalanced := true
  nodeStack.Push(t.rootNode)
  
  uniqueTreeDepth := make(map[int]struct{},0)
  
  for !nodeStack.IsEmpty(){
    // Depth first search

    n := nodeStack.Pop()

    if n.left == nil && n.right == nil {
      _, found := uniqueTreeDepth[n.Depth]
      if !found {
          uniqueTreeDepth[n.Depth] = struct{}{}
      }

      diff := 0
      for i := range uniqueTreeDepth {
        diff = i - diff 
      }

      if len(uniqueTreeDepth) > 2 || math.Abs(float64(diff)) > 1 {
        fmt.Println("Tree is not superbalanced")
        isSuperBalanced = false
        break
      }
    }

    
    if n.right != nil {
      n.right.Depth = n.Depth + 1 
      nodeStack.Push(n.right)
    } 

    if n.left != nil {
      n.left.Depth = n.Depth + 1 
      nodeStack.Push(n.left)
    }

    fmt.Println("Data",n.Data,"Depth",n.Depth)

  }

  return isSuperBalanced
}
