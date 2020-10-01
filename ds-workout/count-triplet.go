package main

func countTriplets(arr []int64, r int64) int64 {

    // geometic progress is a, ar, ar2
    // eqvalent to a/r,a,ar
    itemsToVisit := make(map[int64]int)
    visitedItems := make(map[int64]int)
    countOfOccurances := int64(0)

    for i:= range arr {
        itemsToVisit[arr[i]]++
    }

    for i:= len(arr) - 1; i>=0; i--{
        index1 := arr[i]
        itemsToVisit[index1]--
        if index1 % r == 0 {
            index2 := index1*r
            index0 := index1/r
            countOfOccurances += int64(visitedItems[index2] * 
            itemsToVisit[index0])
        }
        visitedItems[index1]++
    }

    return countOfOccurances
}
