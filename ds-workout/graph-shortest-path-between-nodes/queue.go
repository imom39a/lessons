package main

import "fmt"

type Queue struct {
  items []string
}

func (q *Queue) Enqueue(item string){
  q.items = append(q.items, item)
}

func (q *Queue) Dequeue() string {
    if q.IsEmpty(){
      fmt.Println("Queue is empty")
      return ""
    } 
    itemDequeud := q.items[0]
    q.items = q.items[1:]
    return itemDequeud
}

func (q *Queue) IsEmpty() bool {
  return q.items == nil || len(q.items) == 0
}