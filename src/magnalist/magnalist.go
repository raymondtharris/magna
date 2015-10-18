package magnalist

type magnaListNode struct {
	Next *magnaListNode
}

type magnaList struct {
	Head   *magnaListNode
	Tail   *magnaListNode
	Length int
}

func (mList magnaList) Push(newListNode *magnaListNode) {
	if mList.Head == nil {
		mList.Head = newListNode
		mList.Tail = mList.Head
	} else {
		mList.Tail.Next = newListNode
		mList.Tail = newListNode
	}
	mList.Length++
}

func (mList magnaList) Pop() *magnaListNode {
	if mList.Tail == mList.Head {
		returnNode := mList.Head
		mList.Tail = nil
		mList.Head = nil
		mList.Length = 0
		return returnNode
	} else if mList.Tail != nil {
		currentNode := mList.Head
		var returnNode *magnaListNode
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

func (mList magnaList) isEmpty() bool {
	if mList.Head == nil {
		return true
	}
	return false

}
