package main

import "fmt"

func main(){
  solution2() 
}

func solution2(){
  fmt.Println("Coin splitup function - solution 2")
  availableCoins := []int{1,2,3}
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

  noOfPosiblitiesForACoin := 0 
  currentCoin := availableCoins[denominationIndex]
  for amout >= 0 {
    noOfPosiblitiesForACoin += coinSplitup2(amout, denominationIndex + 1, availableCoins)
    amout -= currentCoin
  } 

  return noOfPosiblitiesForACoin
}