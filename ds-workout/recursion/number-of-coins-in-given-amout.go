/*
Write a method that, given:

an amount of money
an array of coin denominations
computes the number of ways to make the amount of money with coins of the available denominations.

Example: for amount=4 (4¢) and denominations=[1,2,3] (1¢, 2¢ and 3¢), your program would output 4—the number of ways to make 4¢ with those denominations:

1¢, 1¢, 1¢, 1¢
1¢, 1¢, 2¢
1¢, 3¢
2¢, 2¢
*/

package main

import "fmt"
//import "time"

//var availableCoins []int
func main(){
  
  //solution1()
  solution2()
  
  
}

func solution2(){
  fmt.Println("Coin splitup function - solution 2")
  availableCoins := []int{1,2,3}
  //n := len(availableCoins)

  fmt.Println(coinSplitup2(4, 0, availableCoins))
  

}

func coinSplitup2(amout, denominationIndex int, availableCoins []int) int{
  
  fmt.Println("ways to make ",amout,"using coins", availableCoins[denominationIndex:])

  if amout == 0 {
    return 1
  }

  if amout < 0 {
    return 0
  }

  if denominationIndex == len(availableCoins) {
    return 0
  }

  //amout = coinSplitup(amout, denominationIndex-1, availableCoins)
  noOfPosiblitiesForACoin := 0 
  currentCoin := availableCoins[denominationIndex]
  for amout >= 0 {
    noOfPosiblitiesForACoin += coinSplitup2(amout, denominationIndex + 1, availableCoins)
    // if noOfPosiblitiesForACoin != 0 {
    //   fmt.Println("coin",currentCoin,"posiblities",noOfPosiblitiesForACoin)
    // }    
    amout -= currentCoin
  } 

  return noOfPosiblitiesForACoin
}

func solution1(){

  fmt.Println("Coin splitup function - solution 1")
  availableCoins := []int{3,2,1}
  n := len(availableCoins)

  for i:=0 ;i<n; i ++ {
      coinSplitup(4, n-i, availableCoins[i:])
      fmt.Println("--------------")
  }
}

func coinSplitup(amout, denominationIndex int, availableCoins []int) int{
  if denominationIndex < 0 {
    return amout
  }
  amout = coinSplitup(amout, denominationIndex-1, availableCoins)
  if amout != 0 {
    multiplier := amout/ availableCoins[denominationIndex]
    amout = amout % availableCoins[denominationIndex]
    if multiplier != 0 {
      fmt.Println(availableCoins[denominationIndex],"cent","*",multiplier,"=",availableCoins[denominationIndex] * multiplier)
    }
    return amout
  }
  return 0
}