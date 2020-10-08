/*
Solves the Towers of Hanoi problem on n discs. The discs are labeled in increasing order of size from 1 to n.
Print the stack snapshot for each move
*/

package main

import "fmt"

func main(){

  fmt.Println("Towers of Hanoi problem")

  source := []int{2,1}
  helper := make([]int,0)
  target := make([]int,0)

  hanoi(len(source), source, helper, target)
}

func hanoi(disk int, source, helper, target []int){
    fmt.Println(disk,source,helper,target)
    if disk == 0 {
      return
    } 
    hanoi(disk - 1, source[:disk], target, helper)
    move(source, target)
    fmt.Print("After move --> ")
    fmt.Println(disk,source,helper,target)
    hanoi(disk - 1, helper, source ,target)
}

func move(source, target []int) {
  fmt.Println("Move one disk from", source, "to", target)
  target = append(target, source[len(source)-1])
  source = source[:len(source)-1]
}