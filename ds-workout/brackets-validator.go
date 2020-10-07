package main

import "fmt"

func myFunction(arg string) bool {

    isValid := true
    
    argRune := []rune(arg)
    
    openBrackets := make(map[string]struct{})
    openBrackets["("] = struct{}{}
    openBrackets["{"] = struct{}{}
    openBrackets["["] = struct{}{}
    
    bracketClosePair := make(map[string]string)
    bracketClosePair[")"] = "("
    bracketClosePair["}"] = "{"
    bracketClosePair["]"] = "["
    
    var openBracketsStack Stack
    
    for _, v := range argRune {
        //openBrackets.Dump()
        currentItem := string(v)
        if _,found := openBrackets[currentItem]; found {
            openBracketsStack.Push(string(v))
        } else {
            popedItem := openBracketsStack.Pop()
            if bracketClosePair[currentItem] != popedItem {
                return false
            }
        }
    }
    

    return isValid
}

func main() {

    // Run your function through some test cases here.
    // Remember: debuggin is half the battle!

    test1 := "{[]()}"
    test2 := "{[(])}"
    test3 := "{[}"

    fmt.Println(test1,"is valid",myFunction(test1))
    fmt.Println(test2,"is valid",myFunction(test2))
    fmt.Println(test3,"is valid",myFunction(test3))
}

type Item interface{}

type Stack struct {
  items []Item
}


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