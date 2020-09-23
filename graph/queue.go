package main

import "fmt"

type Queue struct {
  items []*GraphNode
}

func (q *Queue) Enqueue(item *GraphNode){
  q.items = append(q.items, item)
}

func (q *Queue) Dequeue() *GraphNode {
    if q.IsEmpty(){
      fmt.Println("Queue is empty")
      return nil
    } 
    itemDequeud := q.items[0]
    q.items = q.items[1:]
    return itemDequeud
}

func (q *Queue) IsEmpty() bool {
  return q.items == nil || len(q.items) == 0
}