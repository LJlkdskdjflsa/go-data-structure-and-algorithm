package Link

type QueneLinkNode struct {
	rear  *Node
	front *Node
}

type QueneLink interface {
	length() int
	EnQueue() interface{}
	DeQueue() interface{}
}

func NewQueneLink() *QueneLinkNode {
	return &QueneLinkNode{}
}

func (n *QueneLinkNode) EnQueue(value interface{}) {
	//頭部插入
	newNode := &Node{value, nil}
	if n.front == nil {
		n.front = newNode //插入一個節點
		n.rear = newNode
	} else {
		n.rear.PNext = newNode
		n.rear = n.rear.PNext
	}
}

func (n *QueneLinkNode) DeQueue() interface{} {
	if n.front == nil {
		return nil
	}
	newNode := n.front     //紀錄頭部位置
	if n.front == n.rear { //只存在一個

		n.front = nil
		n.rear = nil
	} else {
		n.front = n.front.PNext
	}

	return newNode.Data
}

func (n *QueneLinkNode) length() int {
	pnext := n.front
	length := 0
	for pnext.PNext != nil {
		pnext = pnext.PNext
		length++
	}
	return length
}
