package main

import "fmt"
import "errors"

func findShortestPath(source, target string, input map[string][]string) error {
  var validationError error

  _, sourceFound := input[source]
  _, targetFound := input[target]
  if !sourceFound || !targetFound {
      validationError = errors.New("Source or target not found")
  }

  if validationError != nil {
    return validationError
  }
  
  // use BFS to see if the nodes are reachable 
  pathExist := false
  visitedNodes := make(map[string]struct{});
  pathTraveled := make(map[string]string)

  var queue Queue  
  queue.Enqueue(source)
  pathTraveled[source] = ""

  for !queue.IsEmpty() && !pathExist {
      currentPerson := queue.Dequeue()
      visitedNodes[currentPerson] = struct{}{}
      for i := range input[currentPerson] {
        currentPeer := input[currentPerson][i]
        _ , found := visitedNodes[currentPeer]                
        fmt.Println(currentPerson,"->",currentPeer)
        _, f := pathTraveled[currentPeer]
        if !f {
            pathTraveled[currentPeer] = currentPerson
        }
        
        if !found {
          if currentPeer == target {
            pathExist = true
            fmt.Println("Path Exists")
          }
          queue.Enqueue(currentPeer)
        }
      }      
  }
  
  k := target
  finalPath := make([]string,0)
  for {
    if (len(pathTraveled[k]) == 0){
      finalPath = append(finalPath, k)
      break      
    }
    finalPath = append(finalPath, k)
    k = pathTraveled[k]      
  }

  for i := range finalPath {
    lastIndex := len(finalPath) - 1 - i
    if i >= (lastIndex) {
      break
    }
    tmp := finalPath[lastIndex ]
    finalPath[lastIndex] = finalPath[i]
    finalPath[i] = tmp    
  }
  fmt.Println(finalPath)
  fmt.Println("-------------------")
  
  return nil

}

func main(){

  fmt.Println("Graph - shortest path")

  input := make(map[string][]string) 
  input["Min"]      = []string  { "William", "Jayden", "Omar" }
  input["William"]  = []string  { "Min", "Noam" }
  input["Jayden"]   = []string  { "Min", "Amelia", "Ren", "Noam" }
  input["Ren"]      = []string  { "Jayden", "Omar" }
  input["Amelia"]   = []string  { "Jayden", "Adam", "Miguel" }
  input["Adam"]     = []string  { "Amelia", "Miguel", "Sofia", "Lucas" }
  input["Miguel"]   = []string  { "Amelia", "Adam", "Liam", "Nathan" }
  input["Noam"]     = []string  { "Nathan", "Jayden", "William" }
  input["Omar"]     = []string  { "Ren", "Min", "Scott" }  

  findShortestPath("me", "you", input)
  findShortestPath("Min", "Noam", input)
  // findShortestPath("Min", "Noam", input)
  // findShortestPath("Miguel", "Noam", input)
  // findShortestPath("William", "Noam", input)
  // findShortestPath("Jayden", "Noam", input)
  // findShortestPath("Ren", "Noam", input)
  // findShortestPath("Amelia", "Noam", input)
  

}