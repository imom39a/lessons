package main

import "fmt"

type LinkedList struct {
	rootNode *LinkedListNode
}

func (ll *LinkedList) addItem(item interface{}) {
	currentItem := &LinkedListNode{
		item: item,
	}

	if ll.rootNode != nil {
		var lastItem *LinkedListNode = ll.rootNode

		for lastItem.next != nil {
			lastItem = lastItem.next
		}
		lastItem.next = currentItem
	} else {
		ll.rootNode = currentItem
	}
}

func (ll *LinkedList) deleteItem(item interface{}) {
	fmt.Println("Delete item", item)
	var lastItem *LinkedListNode = ll.rootNode
	for lastItem != nil {  
    if lastItem.next == nil {
      break
    } else if lastItem.next.item == item {
			fmt.Println("Found item to delte")
			lastItem.next = lastItem.next.next
			break
		} 
		lastItem = lastItem.next
	}
}

func (ll *LinkedList) dump() {
	var lastItem *LinkedListNode = ll.rootNode
	for {
		fmt.Println("item in list", lastItem.item)
		if lastItem.next != nil {
			lastItem = lastItem.next
			continue
		}
		break
	}
	fmt.Println("---------------")
}

type LinkedListNode struct {
	item interface{}
	next *LinkedListNode
}

func main() {

	var ll LinkedList

	ll.addItem(10)
	ll.addItem(20)
	ll.addItem(12)
	ll.addItem(32)
	ll.dump()
	ll.deleteItem(20)
	ll.dump()
  ll.deleteItem(220)
  ll.dump()

}
