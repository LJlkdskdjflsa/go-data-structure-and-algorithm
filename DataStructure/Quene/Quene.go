package Quene

type MyQuene interface {
	Size() int
	Front() interface{} //第一個元素
	End() interface{}   //最後一個
	IsEmpty() bool
	EnQueue() interface{} //入隊列
	DeQueue() interface{} //出隊列
	Clear()
}
type Quene struct {
	dataStore []interface{} //隊列存儲
	TheSize   int           //隊列大小
}

func NewQuene() *Quene {
	myQuene := new(Quene)
	myQuene.dataStore = make([]interface{}, 0)
	myQuene.TheSize = 0
	return myQuene
}
func (myQuene *Quene) Size() int {
	return myQuene.TheSize
}
func (myQuene *Quene) Front() interface{} {
	if myQuene.Size() == 0 {
		return nil
	}
	return myQuene.dataStore[0]
} //第一個元素
func (myQuene *Quene) End() interface{} { //最後一個
	if myQuene.Size() == 0 {
		return nil
	}
	return myQuene.dataStore[myQuene.Size()-1]
}
func (myQuene *Quene) IsEmpty() bool {
	return myQuene.Size() == 0
}

//入隊列
func (myQuene *Quene) EnQueue(data interface{}) {

	myQuene.dataStore = append(myQuene.dataStore, data)
	myQuene.TheSize++

}

//出隊列
func (myQuene *Quene) DeQueue() (data interface{}) {
	if myQuene.Size() == 0 {
		return nil
	}
	data = myQuene.dataStore[0]
	if myQuene.Size() > 1 {
		myQuene.dataStore = myQuene.dataStore[1:]
	}
	myQuene.TheSize--
	return data
}
func (myQuene *Quene) Clear() {
	myQuene.dataStore = make([]interface{}, 0)
	myQuene.TheSize = 0
}
