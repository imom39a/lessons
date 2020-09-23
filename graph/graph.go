package main 

import "fmt"

type Graph struct{
  nodes map[string]*GraphNode
}

func (g *Graph) addNode(nodeLabel string){
   _,foundNode := g.nodes[nodeLabel] 
   if !foundNode {
    g.nodes[nodeLabel] = &GraphNode{label: nodeLabel}
   }
}

func (g *Graph) addEdge(edge1,edge2 string ){
   _,foundEdge1 := g.nodes[edge1] 
   if !foundEdge1 {
    g.nodes[edge1] = &GraphNode{label: edge1}
   }

  _,foundEdge2 := g.nodes[edge2] 
   if !foundEdge2 {
    g.nodes[edge2] = &GraphNode{label: edge2}
   }

  g.nodes[edge1].addNeighbor(g.nodes[edge2])
  g.nodes[edge2].addNeighbor(g.nodes[edge1])

}

func (g *Graph) DFSWalk(){
    visitedNodes := make(map[string]struct{})    
    for i := range g.nodes {
      _, isCurrentNodeVisited := visitedNodes[g.nodes[i].label]
      if !isCurrentNodeVisited {
        g.nodes[i].DFSWalk(visitedNodes)
      } 
    }
}

func (g *Graph) BFSWalk(){
    visitedNodes := make(map[string]struct{})
    for i := range g.nodes {
      _, isCurrentNodeVisited := visitedNodes[g.nodes[i].label]
      if !isCurrentNodeVisited {
        g.nodes[i].BFSWalk(visitedNodes)
      } 
    }
}

type GraphNode struct{
  label string
  neighbors map[string]*GraphNode
}

func (g *GraphNode) DFSWalk(visitedNodes map[string]struct{}){
    fmt.Println("DFS Walk for",g.label)

    var stack Stack
    stack.Push(g)

    for !stack.isEmpty() {
      currentNode := stack.Pop()

      _, found := visitedNodes[currentNode.label] 
      if !found {
            visitedNodes[currentNode.label] = struct{}{}
      }

      if currentNode.hasNeighbor() {
         for i := range currentNode.neighbors {
           _, isNeighbouringNodeVisited := visitedNodes[currentNode.neighbors[i].label]
           if !isNeighbouringNodeVisited { 
             fmt.Println(currentNode.label,"->",currentNode.neighbors[i].label)
             stack.Push(currentNode.neighbors[i])
           }
         }
      }       
    }
    
}

func (n *GraphNode) BFSWalk(visitedNodes map[string]struct{}){
    fmt.Println("BFS Walk for",n.label)

    var queue Queue
    queue.Enqueue(n)

    for !queue.IsEmpty() {
      currentNode := queue.Dequeue()

      _, found := visitedNodes[currentNode.label] 
      if !found {
          visitedNodes[currentNode.label] = struct{}{}
      }

      if currentNode.hasNeighbor() {
        for i := range currentNode.neighbors {
           _, isNeighbouringNodeVisited := visitedNodes[currentNode.neighbors[i].label]
           if !isNeighbouringNodeVisited { 
             fmt.Println(currentNode.label,"->",currentNode.neighbors[i].label)
             queue.Enqueue(currentNode.neighbors[i])
           }
         }
      }

    }

}

func (g *GraphNode) hasNeighbor() bool{
  hasNeighbor := false
  if len(g.neighbors) > 0 {
    hasNeighbor = true
  }
  return hasNeighbor
}

func (g *GraphNode) addNeighbor(node *GraphNode){
  if g.neighbors == nil {
    g.neighbors = make(map[string]*GraphNode)
  }
  _, found :=  g.neighbors[node.label]
  if !found {
    g.neighbors[node.label] = node
  }  
}