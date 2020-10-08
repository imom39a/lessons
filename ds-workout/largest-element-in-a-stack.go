package main

import "fmt"

func main() {

    var stackInput Stack
    stackInput.Push(5)
    stackInput.Push(8)
    stackInput.Push(15)
    stackInput.Push(20)
    stackInput.Push(25)
    stackInput.Push(10)
    fmt.Println("------ stack")
    stackInput.Dump()
    fmt.Println("------ max item")
    stackInput.DumpMax()
    fmt.Println("------ pop one")
    stackInput.Pop()
    fmt.Println("------ stack")
    stackInput.Dump()    
    fmt.Println("------ max item")
    stackInput.DumpMax()
    fmt.Println("------ pop one")
    stackInput.Pop()
    fmt.Println("------ stack")
    stackInput.Dump()    
    fmt.Println("------ max item")
    stackInput.DumpMax()    
    fmt.Println("------ pop one")
    stackInput.Pop()
    fmt.Println("------ stack")
    stackInput.Dump()    
    fmt.Println("------ max item")
    stackInput.DumpMax()    
    fmt.Println("------ pop one")
    stackInput.Pop()
    fmt.Println("------ stack")
    stackInput.Dump()    
    fmt.Println("------ max item")
    stackInput.DumpMax()        
    
}


type Stack struct {
  items []int
  maxItems []int
}

func (s *Stack) Push(item int) {
  s.items = append(s.items, item)
  
  var popedMaxItem int
  if s.maxItems != nil && len(s.maxItems) > 0 {
      popedMaxItem = s.maxItems[len(s.maxItems) - 1]
      if popedMaxItem > item {
          return
      }
  }

  s.maxItems = append(s.maxItems, item)
}

    
func (s *Stack) Pop() int {  
  var popedItem int

  if s.items != nil && len(s.items) > 0 {
    popedItem = s.items[len(s.items) - 1]
    s.items = s.items[:len(s.items) - 1]
    
    var popedMaxItem int
    if s.maxItems != nil && len(s.maxItems) > 0 {
      popedMaxItem = s.maxItems[len(s.maxItems) - 1]
      if popedMaxItem == popedItem {
          s.maxItems = s.maxItems[:len(s.maxItems) - 1]
      }
    }
  
  } else {
    fmt.Println("No Items in Stack")
  }

  return popedItem
}

func (s *Stack) DumpMax() {  
    if s.maxItems != nil {
      for i := range s.maxItems {
        fmt.Printf("%v\n",s.maxItems[i])
      }      
    } else {
      fmt.Println("No Max Items in Stack")
    }
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