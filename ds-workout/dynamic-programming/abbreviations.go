package main

import "unicode"
import "fmt"
import "time"

func main(){

fmt.Println(time.Now())
abbreviation("ssssssssssssssssssssiisssissstsssssssssssssssssssssssssstsssslssssgsissstsssssssssslssssssssssssisssssssssisssssitssssssstslsssssssssssssssssssssstsssssssssssssbssssstsssssspssssssssssssssssssssssssssspssssssssssssssssssspssssssssssitsslissssssssssssssssssssssssssssssssssssssssssssssisssssslsslsssstsssssssssssslsssssssisssssssssssstsssssisssssssssssslsssssssssssssssssssssssssssssssssssssssssssssssssslstsssssssssssissssssssisssssssssspsssssssssssssssssssssssspsssssissssssssissssssssstspsssssstssssssssssslssslspssssssssssssssssisssssssssssssssssisssspssssssssssisssssssssssssssssstsssssssissssssssssssssssspslsssssssssstssssspsssssnssssslsssssssssssssssssssssissssssssssssssstsslssssssssssspsssssssssssisssssssssssssssstssssssssstsssslssspsssssssssssssspississspssssssssssstsssslpssssssssissssssssssssssssssssssstssssssssssssisssssssssssssslsssssssssssstsssssssssssisssssssssssssssssssssssssssstssssissssssssssssssssssssssssssspslsssssssissssssissssssssssssssspssssssssssssssssssssssssssissssssls",
"SSSSSSSSSSSSSSSSSIISSSSSTSSSSSSSSSSSSSSSSSSSSSSSSSSSLSSSSSISSTSSSSSSSSLSSSSSSSSSISSSSSSSSSSSITSSSSSTSLSSSSSSSSSSSSSSSSSSTSSSSSSSSSSSSSTSSSSSPSSSSSSSSSSSSSSSSSSSSSSSPSSSSSSSSSSSSSSSSSSSSITSSISSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSISSSSSLSSLSSSTSSSSSSSSSLSSSSSSISSSSSSSSSSSTSSSSISSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSLSTSSSSSSSSSSISSSSSSSSSSSSSSSPSSSSSSSSSSSSSSSSSSSPSSSSSSSSISSSSSSSSSPSSSSSSSSSSSSLSSSSPSSSSSSSSSSSSSSSSISSSSSSSSSSSSSSSISSSPSSSSSSSSISSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSPSLSSSSSSSTSSSPSSSSSSSLSSSSSSSSSSSSSSSSISSSSSSSSSSSSSSTSLSSSSSSSSSSSPSSSSSSSISSSSSSSSSSSSSTSSSSSSSSTSSSSSSSSSSSSSSSPISSISSSPSSSSSSSSSSTSSSSLPSSSSISSSSSSSSSSSSSSSSSSSSSSTSSSSSSSSSSSISSSSSSSSLSSSSSSSSTSSSSSSSSSSSSSSSSSSSSSSSSSSSSSTSSSSSSSSSSSSSSSSSSSSSSPLSSSSSSISSSSSISSSSSSSSSSPSSSSSSSSSSSSSSSSSSSSSISSSS")
fmt.Println(time.Now())
fmt.Println(recTimes)
}

type Pair [2]string

var isMatchFound bool
var mem map[Pair]struct{}
var recTimes int

// Complete the abbreviation function below.
func abbreviation(a string, b string) string {

    mem = make(map[Pair]struct{})
    isMatchFound = false
    isPresent := "NO"    

    rec([]rune(a), []rune(b))

    if isMatchFound {
        isPresent = "YES"    
    }

    fmt.Println(isPresent)

    return isPresent
}

func rec(a, b []rune) {
    //fmt.Println(">",string(a), string(b))
    recTimes++
    
    if isMatchFound == true {
        return
    }

    if _, ok := mem[Pair{string(a),string(b)}]; ok {
        fmt.Println("returning from memory",string(a), string(b))
        return
    }

    if len(a) < len(b) {
        //fmt.Println("->",string(a), string(b))
        isMatchFound = false 
        return 
    }
    if len(b) == 0 {
        for _, r := range a {
            if unicode.IsUpper(r){
                isMatchFound = false 
                return 
            }
        }
        isMatchFound = true 
        return
    }

    if unicode.IsLower(a[0]) {
        rec(a[1:], b)
        if isMatchFound == false {
            newStr := string(unicode.ToUpper(a[0])) + string(a[1:])
            rec([]rune(newStr), b)
        }
    } else if a[0] == b[0] {
        rec(a[1:], b[1:])
    }

   mem[Pair{string(a),string(b)}] = struct{}{}
}
