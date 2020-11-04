// knapsack with no duplicate selection
// not sure if below approach works all the time for all scenarios 
// need more tests 
package main

import "fmt"

var container map[Cake]int

func main() {

	fmt.Println("Cake Thief - Knapsack problem- with repetation allowed")

  bagCapacity := 20
	//bagCapacity := 9
  maxValueSoFar := 0 
  container = make(map[Cake]int)

	availableCakes := []Cake{Cake{7, 160, 160/7}, Cake{3, 90, 90/3}, Cake{2, 15, 15/2}}
  //availableCakes := []Cake{Cake{3, 40, 40/3}, Cake{5, 70, 70/5}}
  availableCakesCount := len(availableCakes)

  fmt.Println("Bag Capacity", bagCapacity)	
  fmt.Println("Available Cakes", availableCakes)
	
  // buble sort and terminate early by skipping last items
  // n^2 
  for i:= 0 ; i < availableCakesCount - 1; i ++ {                           
    for j:= 0 ; j < availableCakesCount - 1 - i; j ++ {                     
      if availableCakes[j+1][2] > availableCakes[j][2] {
          availableCakes[j], availableCakes[j+1] = availableCakes[j+1], availableCakes[j]
      }
    }
  }

  maxValueSoFar = fillBag(bagCapacity, maxValueSoFar, availableCakes)
  fmt.Println("----> Max value bag holds",maxValueSoFar)


}

func fillBag(bagCapacity, maxValueSoFar int, availableCakes []Cake) int{

  if len(availableCakes) == 0 {
    return maxValueSoFar
  }

  totalCapacitySoFar :=0 
  remainingCacacity := 0 
  multiplier := 0 
  
  v := availableCakes[0]
  multiplier = bagCapacity / v[0]
  totalCapacitySoFar = multiplier * v[0]
  maxValueSoFar +=  (multiplier) * v[1]
  remainingCacacity = bagCapacity % v[0]

  fmt.Println()
  fmt.Println("Bag capcaity filled so far", totalCapacitySoFar)
  fmt.Println("Remaining Bag capcaity", remainingCacacity)
	fmt.Println("Max value bag holds", maxValueSoFar)

  maxValueSoFar = fillBag(remainingCacacity, maxValueSoFar, availableCakes[1:])

  return maxValueSoFar

}


type Cake [3]int
