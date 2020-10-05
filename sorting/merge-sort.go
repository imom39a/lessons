package main

import "fmt"

func main(){
  arr := []int32{4,3,6,7,4,2,3}
  mergeSort(arr)
  fmt.Println(arr)
}

func mergeSort(arr []int32) {

    // split to left and right 
    // merge sort left unitl no items left
    // merge sort right unitl no items left
    // merge

    n := len(arr)

    if n < 2 {
        return
    }
    
    midIndex := n/2

    leftArray := make([]int32, midIndex)
    rightArray := make([]int32, n - midIndex)

    copy(leftArray, arr[:midIndex])
    copy(rightArray,arr[midIndex:])

    mergeSort(leftArray)
    mergeSort(rightArray)    

    merge(arr, leftArray, rightArray)
}

func merge(arr, l, r []int32){
    var i,j,k int
    for i < len(l) && j < len(r) {
        if l[i] <= r[j]{
            arr[k] = l[i]
            i++
        }else {
            arr[k] = r[j]
            j++
        }
        k++        
    }
    for i < len(l) {
        arr[k] = l[i]
        i++
        k++
    }
    for j < len(r) {
        arr[k] = r[j]
        j++
        k++
    }
}
