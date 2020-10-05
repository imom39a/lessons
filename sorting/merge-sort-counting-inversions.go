package main

import "fmt"

func main(){
    c, a := mergeSortInversion([]int32{7,5,3,1})
    fmt.Println(c)
    fmt.Println(a)

}

func mergeSortInversion(arr []int32) ([]int32, int64) {

    // split to left and right 
    // merge sort left unitl no items left
    // merge sort right unitl no items left
    // calculate inversions when right item needs to move to left (len(right) - no of items already moved to left)

    n := len(arr)
    if n < 2 {
        return arr, 0
    }
    
    midIndex := n/2

    l := make([]int32, midIndex)
    r := make([]int32, n - midIndex)

    copy(l, arr[:midIndex])
    copy(r,arr[midIndex:])
    
    l, swapl := mergeSortInversion(l)
    r, swapr :=mergeSortInversion(r)    

    swaps := swapl+ swapr

    var i,j,k int
    subArray := make([]int32,0)
    for i < len(l) && j < len(r) {
        if l[i] <= r[j]{
            subArray = append(subArray, l[i])
            i++
        } else {
            subArray = append(subArray, r[j])
            j++
            swaps += int64(len(l) - i)
        }
        k++        
    }
    subArray = append(subArray, l[i:]...)
    subArray = append(subArray, r[j:]...)
    return subArray, swaps
}