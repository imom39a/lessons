/*
https://www.hackerrank.com/challenges/special-palindrome-again/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=strings
*/

package main

import "fmt"

func main() {

  fmt.Println(substrCount("asasdasasd"))
  fmt.Println(substrCount("abcbaba"))
  fmt.Println(substrCount("aaaa"))
}
// Complete the substrCount function below.
func substrCount(s string) int64 {

    subStrings := make([]string,0)

    for i := range s {
        subStrings = append(subStrings, string(s[i]))
        for j := i+1; j < len(s); j ++ {
            //fmt.Println(s[i],s[j])
            if s[i] != s[j] {                
                if 2*j-i < len(s) && s[j+1:(2*j-i + 1)] == s[i:j]{
                    subStrings = append(subStrings, string(s[i:(2*j-i + 1)]))
                    break
                } else {
                    break
                }
            } else {
                subStrings = append(subStrings, string(s[i:j+1]))
            }
        }
    }
    //fmt.Println(subStrings)
    return int64(len(subStrings))
}

