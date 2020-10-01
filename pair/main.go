package main

import "fmt"

type Pair [2]interface{}

func main() {
	p := Pair{"Hello", false}
	fmt.Println(p[0], " ", p[1])
}