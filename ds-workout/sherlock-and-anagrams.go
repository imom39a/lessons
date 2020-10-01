package main

import "sort"

func sherlockAndAnagrams(s string) int32 {

    rune_of_s := []rune(s)
    map_of_substrings := make(map[string]int)
    countOfAnagramPairs := 0

    for i:=0; i<len(s); i++ {
        for j := i+1; j<= len(s); j++ {
            subStringRune := rune_of_s[i:j]
            subStringRuneCopy := make([]rune,len(subStringRune))
            copy(subStringRuneCopy, subStringRune)
            sort.Slice(subStringRuneCopy, func(a,b int) bool {
                return subStringRuneCopy[a] < subStringRuneCopy[b]
            })
            map_of_substrings[string(subStringRuneCopy)]++
        }
    }

    for i := range map_of_substrings {
        countOfAnagramPairs += map_of_substrings[i]*(map_of_substrings[i] - 1)/2
    }

    return int32(countOfAnagramPairs)
}
