package main

import "fmt"

var MAX_POSIBLE_TRANSACTION_AMOUNT int32 = 201
var sortedArrayCounter []int32

func main(){
  fmt.Println(activityNotifications([]int32{2, 3, 4, 2, 3, 6, 8, 4, 5}, int32(5)))
}

func sortAndFindMedian(d int32) float32{
    
    subArrayCounter := int32(0)
    median := int32(0)
    if d % 2 == 0 {
        medianIndex1 := int32(0)
        medianIndex2 := int32(0)

        for i:=int32(0); i< MAX_POSIBLE_TRANSACTION_AMOUNT ; i++ {
            subArrayCounter += sortedArrayCounter[i]
            if medianIndex1 == 0 && subArrayCounter >= d/2 {
                medianIndex1 = i
            }
            if medianIndex2 == 0 && subArrayCounter >= d/2 +1 {
                medianIndex2 = i
                break
            }
        }
        return float32(medianIndex1+medianIndex2)/2
    } else {
        for i:=int32(0); i< MAX_POSIBLE_TRANSACTION_AMOUNT ; i++ {
            subArrayCounter += sortedArrayCounter[i]
            if subArrayCounter > int32(d/2) {
                median = i
                break
            }
        }
        return float32(median)
    }
}

// Complete the activityNotifications function below.
func activityNotifications(expenditure []int32, d int32) int32 {

    n := int32(len(expenditure))
    notifications := int32(0)

    sortedArrayCounter = make([]int32,MAX_POSIBLE_TRANSACTION_AMOUNT)

    for i:=int32(0); i< d; i++{
        sortedArrayCounter[expenditure[i]]++
    }
    
    for i:=d; i<n; i++ {
        currentMedian := sortAndFindMedian(d)
        if (float32(expenditure[i]) >= 2*currentMedian){
            notifications++
        }
        sortedArrayCounter[expenditure[i-d]]--
        sortedArrayCounter[expenditure[i]]++
    }
    

    return notifications
}