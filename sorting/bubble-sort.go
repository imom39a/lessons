package main

import "fmt"


func main(){
  countSwaps([]int32{5,4,3,2,1,0})
}

func swap(i, j int32, a []int32) {
	tmp := a[j]
	a[j] = a[i]
	a[i] = tmp
}

// Complete the countSwaps function below.
func countSwaps(a []int32) {
	n := int32(len(a))
	countOfSwaps := 0
	for i := int32(0); i < n; i++ {
		for j := int32(0); j < n-1; j++ {
			if a[j] > a[j+1] {
				countOfSwaps++
				swap(j, j+1, a)
			}
		}
	}

	fmt.Println("Array is sorted in", countOfSwaps, "swaps.")
	fmt.Println("First Element:", a[0])
	fmt.Println("Last Element:", a[n-1])
}
