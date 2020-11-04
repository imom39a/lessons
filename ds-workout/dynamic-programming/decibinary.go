package main

import (    
    "fmt"
    "strconv"
    "strings"
    "math"
    
)

var mem[17][10]int
const maxPositions = 54

var countmem [21]int
var lookup[100][55]int

func populateDecibinarySumLookup(){
    for i:= 0 ; i <17; i ++ {
        for j:= 0 ; j <10; j++ {
            mem[i][j] = j * int(math.Pow(2.0, float64(i)))
        }
    }
}
// Complete the decibinaryNumbers function below.
func decibinaryNumbers(x int64) int64 {
  fmt.Println("decibinaryNumbers",x)

  //fmt.Println("start",time.Now())

  // populate mem[i][j] to find position values 
  populateDecibinarySumLookup()    

  for i :=0; i < 10; i ++ {
    lookup[i][1] = 1
  }
    
  //decBinMax := int64(math.Pow(10.0,16.0))
  decimalMax := 10
  
  for i:= 0 ; i < decimalMax; i++ {
    countmem[i] = noOfOccurancesPossible(i)
  }

  // countmem[2] = noOfOccurancesPossible(2)
  // countmem[3] = noOfOccurancesPossible(3)
   // countmem[4] = noOfOccurancesPossible(4)
  // countmem[5] = noOfOccurancesPossible(5)
  // countmem[6] = noOfOccurancesPossible(6)

  for i := range countmem {
    fmt.Println(i,"=",countmem[i])
  }

  // for i := range lookup {
  //   fmt.Println(i,"=",lookup[i])
  // }
//    fmt.Println("end",time.Now())
    
    fmt.Println("------")
    return 0
}

func noOfOccurancesPossible(decimalNo int) int{  
  fmt.Println("noOfOccurancesPossible",decimalNo)
  totalPosiblities := 0 
  for i:=1 ; i <= maxPositions ; i++ {
    pmax := int(math.Floor(float64(decimalNo / int(math.Pow(2.0, float64(i))))))

    if pmax > 9 {
      pmax = 9 
    }
    
    val := 0 
    fmt.Println("pmax",pmax)
    if pmax != 0 {
      for p := 0; p <= pmax; p++ {
        r := decimalNo - p * int(math.Pow(2.0, float64(i)))              
        if r != 0 {
          fmt.Println("lookup[",r,"][",i,"]", lookup[r][i])          
          val += lookup[r][i]
        }
      }
      lookup[decimalNo][i+1] = val      
    } else {
      break
    }
    
    totalPosiblities += val  
  }
  return totalPosiblities
}

func decimal(x int) int{
    
    intArray := converIntToArray(int(x))
    digits := len(intArray)
    maxPow := digits - 1 
    decibinary := 0
    for j:= maxPow; j >= 0 ; j-- {
        decibinary += mem[maxPow - j][intArray[j]]
    }
    //fmt.Println("decibinary decimal",x,"is",decibinary)

    return decibinary
}

func converIntToArray(x int) []int {
    intArray := make([]int,0)
    for _, v := range strings.Split(strconv.Itoa(x),""){
        val, _ := strconv.Atoi(v)
        intArray = append(intArray, val)
    }
    return intArray
}

func main(){
  decibinaryNumbers(10)
}