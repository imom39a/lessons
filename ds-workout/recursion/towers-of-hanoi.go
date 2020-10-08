/*
Solves the Towers of Hanoi problem on n discs. The discs are labeled in increasing order of size from 1 to n.
Find number of moves
*/

package main

import "fmt"

var counter int 
func main(){

  fmt.Println("Towers of Hanoi problem")

  source := "Source"
  helper := "Helper"
  target := "Target"

  hanoi(3, source, helper, target)
  fmt.Println("total moves",counter)
}

func hanoi(disk int, source, helper, target string){
    if disk == 0 {
      return
    } 
    hanoi(disk - 1, source, target, helper)
    move(source, target)
    hanoi(disk - 1, helper, source ,target)
}

func move(source, target string) {
  fmt.Println("Move one disk from", source, "to", target)  
  counter++
}