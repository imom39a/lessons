package main

import "fmt"


func main() {

    head := &LinkedListNode{
        item: "head",
    }
    
    item1 := head.addNext("item1")
    item2 := item1.addNext("item2")
    item3 := item2.addNext("item3")
    
    fmt.Println("------- printing linked list")
    head.dump()
    fmt.Println("------- reversing linked list")
    head.reverse()
    item3.dump()
    
}


type LinkedListNode struct {
	item interface{}
	next *LinkedListNode
}

func (node *LinkedListNode) addNext(item interface{}) *LinkedListNode {
    
    newItem := &LinkedListNode{
        item: item,
    }
    
    node.next = newItem
    
    return newItem
}

func (node *LinkedListNode) dump(){
    var lastNode *LinkedListNode = node
    for lastNode != nil {
        fmt.Print("current node -> ",lastNode.item,"    |   ")
        if lastNode.next != nil {
            fmt.Println("next node -> ",lastNode.next.item)
            
        }else {
            fmt.Println()
        }
        
        lastNode = lastNode.next
    }
    
}

func (node *LinkedListNode) reverse(){
    var previousNode *LinkedListNode = nil
    var currentNode *LinkedListNode = node
    var nextNode *LinkedListNode = currentNode.next
    for currentNode != nil {
        fmt.Println("-->",currentNode.item)
        nextNode = currentNode.next
        currentNode.next = previousNode
        previousNode = currentNode
        currentNode = nextNode
    }
}