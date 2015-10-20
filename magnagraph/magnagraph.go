package magnagraph

import (
	//"fmt"
	//"magna/magnalist"
)
//Graph is comething
type Graph struct {
}

//Node is something
type Node struct {
	Index  int    //Inex of the node
	Value  string //String value for the particular node
	Weight int    // Weight of the node in relation to the others
}

//Node for a list
type node struct {
	Next *node
	//Payload *magnagraph.Node
}

//List data structure
type list struct {
	Head   *node
	Tail   *node
	Length int
}

//Push inserts a Node at the end of the List
func (mList list) Push(newNode *node) {
	if mList.Head == nil {
		mList.Head = newNode
		mList.Tail = mList.Head
	} else {
		mList.Tail.Next = newNode
		mList.Tail = newNode
	}
	mList.Length++
}

//Pop removes a Node from the end of the List
func (mList list) Pop() *node {
	if mList.Tail == mList.Head {
		returnNode := mList.Head
		mList.Tail = nil
		mList.Head = nil
		mList.Length = 0
		return returnNode
	} else if mList.Tail != nil {
		currentNode := mList.Head
		var returnNode *node
		for currentNode != mList.Tail {
			if currentNode.Next != mList.Tail {
				currentNode = currentNode.Next
			} else {
				returnNode = currentNode.Next
				currentNode.Next = nil
				mList.Tail = currentNode
			}
		}
		mList.Length--
		return returnNode
	}
	return nil
}

func (mList list) isEmpty() bool {
	if mList.Head == nil {
		return true
	}
	return false

}

//Search will look for a Node and return in from a list
func (mList list) Search(nodeToFind *Node) *node{
	return nil
}
