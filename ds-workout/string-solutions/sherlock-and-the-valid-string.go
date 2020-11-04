package main

import "fmt"

func main(){

  fmt.Println(isValid("abcdefghhgfedecba"))
  fmt.Println(isValid("aaaabbcc"))

}
func isValid(s string) string {

    returnString := "NO"

    sRunes := []rune(s)
    charRunes := []rune{'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z'}
    counter := make(map[rune]int)

    for _,sRune := range sRunes {
        counter[sRune]++
    }

    occurances := make(map[int]int)
    for _,charRune := range charRunes{
        if counter[charRune] != 0 {
            occurances[counter[charRune]]++
        }        
    }
    
    if len(occurances) > 2 {
        returnString = "NO"
    } else if len(occurances) == 1 {
        returnString = "YES"
    } else {
        
        sortedVal := make([]int,2)
        i := 0;
        for v:= range occurances {
            sortedVal[i] = v
            i++
        }
        if sortedVal[0] > sortedVal[1] {
            sortedVal[0],sortedVal[1] = sortedVal[1],sortedVal[0]
        }

        if (sortedVal[0] == 1 && occurances[sortedVal[0]] == 1) || 
        (abs(sortedVal[1] - sortedVal[0]) == 1 &&  occurances[sortedVal[1]] == 1) {
            returnString = "YES"
        } 
    }

    return returnString
}


func abs(x int) int {
    if x < 0 {
        return -x 
    }
    return x
}
