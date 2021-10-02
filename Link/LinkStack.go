package Link

type Node struct {
	Data  interface{}
	PNext *Node
}
type LinkStack interface {
	IsEmpty() bool
	Push(data interface{})
	Pop() interface{}
	Length() int
}

func NewStack() *Node {
	return &Node{}
}

func (n *Node) IsEmpty() bool {
	if n.PNext == nil {
		return true
	} else {
		return false
	}
}
func (n *Node) Push(data interface{}) {
	newNode := &Node{data, nil}
	newNode.PNext = n.PNext
	n.PNext = newNode
}
func (n *Node) Pop() interface{} {
	if n.IsEmpty() {
		return nil
	}
	value := n.PNext.Data
	n.PNext = n.PNext.PNext //刪除
	return value
}
func (n *Node) Length() int {
	pnext := n
	length := 0
	for pnext.PNext != nil {
		pnext = pnext.PNext
		length++
	}
	return length
}
