// duplicate selection allowed 

package main

import "fmt"

func main() {

	//availableCakes := []Cake{Cake{7, 160}, Cake{3, 90}, Cake{2, 15}}
	availableCakes := []Cake{Cake{5, 70}, Cake{3, 40}}
	weightCapacity := 9

  fmt.Println(availableCakes)
	maxValueAtCapacity := maxDuffelBagValue(availableCakes, weightCapacity)
	fmt.Println(maxValueAtCapacity)

}

func maxDuffelBagValue(availableCakes []Cake, weightCapacity int) int {

	maxValuesAtCapacities := make([]int, weightCapacity+1)

	for currentCapacity := 0; currentCapacity <= weightCapacity; currentCapacity++ {
    maxValueAtCapacity := 0
		for _, cakeType := range availableCakes {
			if cakeType[0] <= currentCapacity {
        possibleMaxValueUsingCurrentCakes :=  maxValuesAtCapacities[currentCapacity - cakeType[0]] + cakeType[1]
        maxValueAtCapacity = max(maxValueAtCapacity, possibleMaxValueUsingCurrentCakes)
			}      
		}
    maxValuesAtCapacities[currentCapacity] = maxValueAtCapacity   
    fmt.Println("currentCapacity",currentCapacity,"max value",maxValueAtCapacity)
	}
  fmt.Println(maxValuesAtCapacities)
	return maxValuesAtCapacities[weightCapacity-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Cake [2]int
