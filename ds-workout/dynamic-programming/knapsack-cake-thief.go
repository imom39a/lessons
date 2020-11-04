package main

import "fmt"

var container map[Cake]int

func main() {

	fmt.Println("Cake Thief - Knapsack problem- with repetation allowed")

  bagCapacity := 20
  maxValueSoFar := 0 
  container = make(map[Cake]int)

	availableCakes := []Cake{Cake{7, 160}, Cake{3, 90}, Cake{2, 15}}
  availableCakesCount := len(availableCakes)

  fmt.Println("Bag Capacity", bagCapacity)	
  fmt.Println("Available Cakes", availableCakes)


  fmt.Println("----> Max value bag holds",maxValueSoFar)

}

type Cake [2]int
