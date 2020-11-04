/* https://www.hackerrank.com/challenges/minimum-absolute-difference-in-an-array/problem?h_l=interview&playlist_slugs%5B%5D%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D%5B%5D=greedy-algorithms&isFullScreen=true
*/

package main

import (    
    "fmt"
    "math"
    "sort"
)

func main() {
  fmt.Println(minimumAbsoluteDifference([]int32{-59,-36, -13, 1, -53, -92, -2, -96, -54, 75}))
  fmt.Println(minimumAbsoluteDifference([]int32{1,-3, 71, 68, 17}))
}

// Complete the minimumAbsoluteDifference function below.
func minimumAbsoluteDifference(arr []int32) int32 {
    
    minNoSoFar := int32(math.MaxInt32)
    mem := make(map[Pair]struct{})
    
    sort.Slice(arr, func (i,j int) bool{
        if arr[i] < arr[j] {
            return true
        }
        return false
    })

    for i:=0 ; i < len(arr)-1; i ++ {        
        pair := Pair{arr[i], arr[i+1]}
        if _, ok := mem[pair] ; !ok {
            currentDiff := abs(arr[i] - arr[i+1])
            if currentDiff < minNoSoFar{
                minNoSoFar = currentDiff                    
            }                                
            mem[pair] = struct{}{}
        }
    }

    return minNoSoFar
}

type Pair [2]int32 

func abs(x int32) int32 {
    if x < 0 {
        return -x
    }
    return x
}

