package magnalist

type magnaListNode struct {
}

type magnaList struct {
	Head   *magnaListNode
	Tail   *magnaListNode
	Length int
}

func (mList magnaList) Push() {
	mList.Length++
}

func (mList magnaList) Pop() *magnaListNode {
	return nil
}

func (mList magnaList) isEmpty() bool {
	return false
}
