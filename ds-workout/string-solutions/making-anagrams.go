/*
Print a single integer denoting the number of characters you must delete to make the two strings anagrams of each other.

Sample Input

cde
abc

Sample Output

4
*/

package main

import (
    "fmt"
)

func main(){
  makeAnagram("abc", "adb")
}
// Complete the makeAnagram function below.
func makeAnagram(a string, b string) int32 {
    aRunes := []rune(a)
    bRunes := []rune(b)
    charRunes := []rune{'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z'}

    aRuneCounter := make(map[rune]int)
    bRuneCounter := make(map[rune]int)
    oddCounter := int32(0)

    for _,aRune := range aRunes {
        aRuneCounter[aRune]++
    }

    for _,bRune := range bRunes {
        bRuneCounter[bRune]++
    }

    for _, charRune := range charRunes {
        oddCounter += int32(abs(aRuneCounter[charRune] - bRuneCounter[charRune]))
    }
    
    fmt.Println(oddCounter)

    return oddCounter
}

func abs(x int) int{
    if x < 0 {
        return -x 
    }
    return x
}