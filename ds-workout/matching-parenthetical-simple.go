/*
"Sometimes (when I nest them (my parentheticals) too much (like this (and this))) they get confusing."

Write a method that, given a sentence like the one above, along with the position of an opening parenthesis, finds the corresponding closing parenthesis.

Example: if the example string above is input with the number 10 (position of the first parenthesis), the output should be 79 (position of the last parenthesis).
*/

package main

import "fmt"

func myFunction(arg string, position int) {

    // Write the body of your function here
    bracketCounter := 0
    argRune := []rune(arg)
    for i:=position; i < len(argRune); i++ {
        if string(argRune[i]) == "("{
            bracketCounter++
        } else if string(argRune[i]) == ")"{
            bracketCounter--
            if bracketCounter == 0 {
                fmt.Println("closing bracket position is",i)
                return
            }
        }
        

    }
    
    fmt.Println("No closing bracket found")
    
}

func main() {

    // Run your function through some test cases here.
    // Remember: debuggin is half the battle!
    testString := "Sometimes (when I nest them (my parentheticals) too much (like this (and this))) they get confusing."
    myFunction(testString, 10)
}