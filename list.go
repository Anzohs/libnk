package libnk

type Node struct {
	Data int
	Next *Node
}

func NewList(n int) *Node {
	l := Node{Data: n}
	l.Next = nil
	return &l
}

func (n *Node) ListLen() int {
	i := 0

	for n != nil {
		i++
		n = n.Next
	}
	return i
}

func (n *Node) PrintList() {

}
