package main

import "fmt"

func main(){
  fmt.Println("Graph")

  a := &GraphNode {
    label : "a",
  }

  b := &GraphNode {
    label : "b",
  }

  c := &GraphNode {
    label : "c",
  }

  d := &GraphNode {
    label : "d",
  }

  e := &GraphNode {
    label : "e",
  }

  a.addNeighbor(b)
  a.addNeighbor(c)
  b.addNeighbor(a)
  b.addNeighbor(c)
  b.addNeighbor(d)
  c.addNeighbor(a)
  c.addNeighbor(b)
  c.addNeighbor(d)
  d.addNeighbor(b)
  d.addNeighbor(c)


  colors := []string{"R","B","G","Y"}
  graphNodes := []*GraphNode{a,b,c,d,e}


  for i := range graphNodes {
    graphNodes[i].AssignColors(colors)
    fmt.Println("Node",graphNodes[i].label, "has color",graphNodes[i].color)
  }    

}