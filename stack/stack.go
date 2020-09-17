package main

import "fmt"

type Stack struct {
  items []Item
}

// func (s *Stack) Init(){
//   if s.items == nil {
//     s.items = make([]Item,0)
//   }  
// }

func (s *Stack) Push(item Item) {
  s.items = append(s.items, item)  
}

func (s *Stack) Pop() Item {  
  var popedItem Item = nil;

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