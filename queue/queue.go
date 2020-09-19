package main

import "fmt"

type Item interface {}

type Queue struct {
  items []Item
}

func (q *Queue) Enqueue(item Item){
  q.items = append(q.items, item)
}

func (q *Queue) Dequeue() Item {
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