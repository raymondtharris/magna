package magnagraph

import (
	//"fmt"
)
//Graph is comething
type Graph struct {
	NumberOfVerticies int
	NumberOfEdges int
	ADJList []Node
}

//Node is something
type Node struct {
	Index  int    //Inex of the node
	Value  string //String value for the particular node
	Weight int    // Weight of the node in relation to the others
}

//Node for a queue
type node struct {
	Next *node
	Payload *Node
}

//Queue data structure
type queue struct {
	Head   *node
	Tail   *node
	Length int
}

//Push inserts a Node at the end of the List
func (mQueue queue) Enqueue(newNode *node) {
	if mQueue.Head == nil {
		mQueue.Head = newNode
		mQueue.Tail = mQueue.Head
	} else {
		mQueue.Tail.Next = newNode
		mQueue.Tail = newNode
	}
	mQueue.Length++
}

//Pop removes a Node from the end of the queue
func (mQueue queue) Dequque() *node {
	if mQueue.Head != nil {
		nodeToReturn := mQueue.Head
		mQueue.Head = mQueue.Head.Next
		nodeToReturn.Next = nil
		mQueue.Lenth--
		return nodeToReturn
	}
	return nil
}

func (mQueue queue) isEmpty() bool {
	if mQueue.Head == nil {
		return true
	}
	return false

}

//Search will look for a Node and return in from a queue
func (mQueue queue) Search(nodeToFind *Node) *node{
	currentNode := mQueue.Head
	for currentNode != nil {
		if currentNode.Payload == nodeToFind.Payload{
			return currentNode
		} else{
			currentNode = currentNode.Next
		}
	}
	return nil
}
