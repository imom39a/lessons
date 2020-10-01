package main

import "fmt"

type Pair struct {
	vals [2]interface{}
}

func MakePair(k, v interface{}) Pair {
	return Pair{[2]interface{}{k, v}}
}

func (p Pair) Get(i int) interface{} {
	return p.vals[i]
}

func main() {
	p := MakePair("Hello", false)
	fmt.Println(p.Get(0), " ", p.Get(1))
}
