package main

import "fmt"

func main(){

  var graph *Graph = &Graph{}  
  graph.nodes = make(map[string]*GraphNode)

  graph.addEdge("a", "b")
  graph.addEdge("b", "c")
  graph.addEdge("c", "d")
  graph.addEdge("d", "f")
  graph.addEdge("f", "a")
  graph.addEdge("d", "b")
  graph.addNode("z")
  graph.addEdge("a", "z")

  fmt.Println("BFS Walk with initial state")
  //graph.BFSWalk()  
  // fmt.Println("---------------")
  graph.nodes["a"].BFSWalk(make(map[string]struct{}))
  // graph.nodes["f"].BFSWalk(make(map[string]struct{}))
  // graph.nodes["c"].BFSWalk(make(map[string]struct{}))
  // graph.nodes["z"].BFSWalk(make(map[string]struct{}))

  fmt.Println("DFS Walk with initial state")
  //graph.DFSWalk()  
  // fmt.Println("---------------")
  graph.nodes["a"].DFSWalk(make(map[string]struct{}))
  // graph.nodes["f"].DFSWalk(make(map[string]struct{}))
  // graph.nodes["c"].DFSWalk(make(map[string]struct{}))
  // graph.nodes["z"].DFSWalk(make(map[string]struct{}))

}