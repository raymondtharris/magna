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
	mList.Length--
	return nil
}

func (mList magnaList) isEmpty() bool {
	if mList.Head == nil {
		return true
	} else {
		return false
	}
}
