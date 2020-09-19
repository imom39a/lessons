package main

import "fmt"

type Stack struct {
  items []*BinaryTreeNode
}

func (s *Stack) Push(node *BinaryTreeNode){
  s.items = append(s.items, node)
}

func (s *Stack) Pop() *BinaryTreeNode {  
  if s.IsEmpty() {
    fmt.Println("Stack is Empty")
    return nil
  } 
  popedItem := s.items [len(s.items) - 1 ]
  s.items = s.items[:len(s.items) - 1 ]
  return popedItem
}

func (s *Stack) IsEmpty() bool{
  return s.items == nil || len(s.items) == 0
}