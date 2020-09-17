package main

import "fmt"

type BinaryTree struct {
    rootNode *BinaryTreeNode
}

func (t *BinaryTree) Insert(value int) *BinaryTreeNode {
    if t.rootNode == nil {
        t.rootNode = &BinaryTreeNode{Value: value}
    } else {
        t.rootNode.Insert(value)
    }
    return t.rootNode
}

func (t *BinaryTree) Search(x int) bool{
  found := false
  if(t.rootNode != nil) {
      found = t.searchNodes(t.rootNode, x)
  }
  return found
}

func (t *BinaryTree) searchNodes(node *BinaryTreeNode, x int) bool{
  
  found := false     
  //fmt.Println(node.Value)
  
  if x == node.Value {
      //fmt.Println("Found")
      found = true
      return found
  }

  if node.left != nil && !found {
    found = t.searchNodes(node.left, x)
  }

  if node.right != nil && !found {
    found = t.searchNodes(node.right, x)
  }            
  return found    
}

func (t *BinaryTree) Walk() {
    if t.rootNode != nil {
      walkNodes(t.rootNode)
    }    
}
// Postorder traversal
func walkNodes(node *BinaryTreeNode) {
    
    // post order
    //fmt.Printf("%+v\n", node)
    
    if node.left != nil {
        walkNodes(node.left)
    }
    
    if node.right != nil {
        walkNodes(node.right)
    }
    
    // pre order 
    fmt.Println(node.Value)

}

type BinaryTreeNode struct {
    Value int
    left *BinaryTreeNode
    right *BinaryTreeNode
}

func (n *BinaryTreeNode) Insert(value int) {
    
    if value <= n.Value {
        if n.left == nil {
            n.left = n.InsertLeft(value)
        } else {
            n.left.Insert(value)
        }
    } else {
        if n.right == nil {
            n.right = n.InsertRight(value)
        } else {
            n.right.Insert(value)
        }
    }
    
}

func (n *BinaryTreeNode) InsertLeft(leftValue int) *BinaryTreeNode{
    n.left = &BinaryTreeNode{Value: leftValue }
    return n.left
}

func (n *BinaryTreeNode) InsertRight(rightValue int) *BinaryTreeNode{
    n.right = &BinaryTreeNode{Value: rightValue }
    return n.right
}