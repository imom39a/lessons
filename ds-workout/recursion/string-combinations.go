package main

import "fmt"

var combinations map[string]struct{}

func myFunction(arg string) map[string]struct{} {
    
    str := make([]rune,0,len(arg))
    str = append(str,[]rune(arg)...)
    
    combinations = make(map[string]struct{})
    findCombinations(0, 0, str)

    return combinations
}

func findCombinations(i, j int, str []rune) {
    //fmt.Println("->",i,j,string(str))
    if i > len(str) - 1 || j > len(str) - 1 {
        combinations[string(str)] = struct{}{}
        return 
    }
    
    findCombinations(i, j+1, str)
    swap(i,j,str)
    combinations[string(str)] = struct{}{}
    findCombinations(j, i+1, str)
}

func swap(i, j int, str []rune) {
    str[i], str[j] = str[j], str[i] 
}

func main() {

    // Run your function through some test cases here.
    // Remember: debuggin is half the battle!
    fmt.Println(myFunction("abcsd"))
}
