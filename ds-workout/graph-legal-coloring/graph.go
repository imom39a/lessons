package main 

type GraphNode struct{

  label string
  neighbors map[string]*GraphNode
  color string

}

func (g *GraphNode) HasColor() bool{
  hasColor := false
  if len(g.color) > 0 {
    hasColor = true
  }
  return hasColor
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


func (g *GraphNode) AssignColors(colors []string) {  

  illegalColors := make(map[string]struct{})

  for i := range g.neighbors {
    if g.neighbors[i].HasColor() {
      illegalColors[g.neighbors[i].color] = struct{}{}
    }
  }

  nextLegalColors := ""
  for i := range colors {
    _, found := illegalColors[colors[i]]
    if !found {
        nextLegalColors = colors[i]
        break
    }
  }
  
  g.color = nextLegalColors
}