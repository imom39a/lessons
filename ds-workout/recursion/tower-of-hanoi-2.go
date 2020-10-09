/*
Solves the Towers of Hanoi problem on n discs. The discs are labeled in increasing order of size from 1 to n.
Print the stack snapshot for each move
*/

package main

import "fmt"

var mem map[string][]int 

func main(){

  fmt.Println("Towers of Hanoi problem")

  mem = make(map[string][]int)
  
  source := "Source"
  helper := "Helper"
  target := "Target"

  mem[source] = []int{5,4,3,2,1}  
  mem[helper] = make([]int,0)
  mem[target] = make([]int,0)

  hanoi(len(mem[source]), source, helper, target)
  printSnaphot()
}

func hanoi(disk int, source, helper, target string) {        
    if disk == 0 {
      return 
    } 
    hanoi(disk - 1, source,target,helper)    
    move(source, target)
    hanoi(disk - 1, helper,source,target)    
}

func move(source, target string) {
  printSnaphot()
  if(len(mem[source]) > 0) {
    mem[target] = append(mem[target], mem[source][len(mem[source])-1])
    mem[source] = mem[source][:len(mem[source])-1]
  }
}

func printSnaphot(){  
  fmt.Println("A -- ",mem["Source"])
  fmt.Println("B -- ",mem["Helper"])
  fmt.Println("C -- ",mem["Target"])
  fmt.Println("------------")
}