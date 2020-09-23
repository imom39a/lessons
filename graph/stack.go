package main

import "fmt"

type Stack struct {
  items []*GraphNode
}

func (s *Stack) Push(item *GraphNode) {
  s.items = append(s.items, item)  
}

func (s *Stack) Pop() *GraphNode {  
  var popedItem *GraphNode;

  if s.items != nil && len(s.items) > 0 {
    popedItem = s.items[len(s.items) - 1]
    s.items = s.items[:len(s.items) - 1]    
  } else {
    fmt.Println("No Items in Stack")
  }

  return popedItem
}

func (s *Stack) Dump(){
    if s.items != nil {
      for i := range s.items {
        fmt.Printf("%v\n",s.items[i])
      }      
    } else {
      fmt.Println("No Items in Stack")
    }
}

func (s *Stack) Peek() *GraphNode{
  return s.items[len(s.items) - 1]
}

func (s *Stack) isEmpty() bool{
    return s.items == nil || len(s.items) == 0
}