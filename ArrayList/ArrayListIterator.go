package ArrayList

import "errors"

type Iterator interface {
	HasNext() bool
	Next() (interface{}, error)
	Remove()
	GetIndex() int //得到索引
}

type Iterable interface {
	Iterator() Iterator //構造初始化接口
}

// 構造指針訪問數組
type ArrayListIterator struct {
	list         *ArrayList //包含 數組指針
	CurrentIndex int        //當前索引
}

func (list *ArrayList) Iterator() Iterator {
	it := new(ArrayListIterator) //構造迭代器
	it.CurrentIndex = 0
	it.list = list
	return it

}

func (it *ArrayListIterator) HasNext() bool {
	return it.CurrentIndex < it.list.TheSize
}

func (it *ArrayListIterator) Next() (interface{}, error) {
	if !it.HasNext() {
		return nil, errors.New("No next")

	}
	value, err := it.list.Get(it.CurrentIndex)
	it.CurrentIndex++
	return value, err
}
func (it *ArrayListIterator) Remove() {
	it.CurrentIndex--
	it.list.Delete(it.CurrentIndex)
}

func (it *ArrayListIterator) GetIndex() int {
	return it.CurrentIndeudo
}
