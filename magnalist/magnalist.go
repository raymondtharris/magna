package magnalist

//Node for a list
type Node struct {
	Next *Node
}

//List data structure
type List struct {
	Head   *Node
	Tail   *Node
	Length int
}

//Push inserts a Node at the end of the List
func (mList List) Push(newNode *Node) {
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
func (mList List) Pop() *Node {
	if mList.Tail == mList.Head {
		returnNode := mList.Head
		mList.Tail = nil
		mList.Head = nil
		mList.Length = 0
		return returnNode
	} else if mList.Tail != nil {
		currentNode := mList.Head
		var returnNode *Node
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

func (mList List) isEmpty() bool {
	if mList.Head == nil {
		return true
	}
	return false

}
