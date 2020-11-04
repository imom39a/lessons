package main

import "fmt"

func main(){
  luckBalance := luckBalance(3, [][]int32{
    []int32{3 ,1},
    []int32{13 ,1},
    []int32{21 ,1},
    []int32{31 ,1},
    []int32{17 ,1},
    []int32{20 ,1},
    []int32{8 ,1},
    []int32{12 ,0},
    []int32{14 ,0},
    []int32{21 ,1},
  })
  fmt.Println(luckBalance)
}

func luckBalance(k int32, contests [][]int32) int32 {
    
    maximumLuckBalance := int32(0)
    totalContests := len(contests)

    for i:=0 ; i < totalContests; i ++ {
        for j := 0; j < totalContests - i - 1; j ++ {
            if contests[j][0] > contests[j+1][0] {
                contests[j][0], contests[j+1][0] = contests[j+1][0], contests[j][0]
                contests[j][1], contests[j+1][1] = contests[j+1][1], contests[j][1]
            }
        }
//        fmt.Println("-->",contests)
        isImportantGame := contests[totalContests-i-1][1] == 1

        if isImportantGame {            
          canLooseMoreGames := k > 0 
            if canLooseMoreGames {
                k--
            } else {
                maximumLuckBalance -= contests[totalContests-i-1][0]
                continue
            }
        }
        maximumLuckBalance += contests[totalContests-i-1][0]
    }


    return maximumLuckBalance
}