package ArrayList

import (
	"errors"
	"fmt"
)

//api
type List interface {
	Size() int                                  //list size
	Get(index int) (interface{}, error)         //index
	Set(index int, newval interface{}) error    //
	Insert(index int, newval interface{}) error //
	Append(newval interface{})                  //
	Clear()                                     //empty
	Delete(index int) error                     //delete one
	String() string                             //return string
	Iterator() Iterator                         //API
}

//data struct
type ArrayList struct {
	dataStore []interface{} //store list
	TheSize   int           //list size
}

func NewArrayList() *ArrayList {
	list := new(ArrayList)                      // initialize
	list.dataStore = make([]interface{}, 0, 10) //make 10 space for store
	list.TheSize = 0
	return list
}

func (list *ArrayList) checkIsFull() {
	// check the volumn is full or not
	if list.TheSize == cap(list.dataStore) {
		newdataStore := make([]interface{}, 2*list.TheSize, 2*list.TheSize) //double the volumn
		// copy(newdataStore, list.dataStore) // can't use copy here because not specific interface
		for i := 0; i < (list.TheSize); i++ {
			newdataStore[i] = list.dataStore[i]
		}
		list.dataStore = newdataStore // give value back
	}

}

func (list *ArrayList) Size() int {
	return list.TheSize
}

func (list *ArrayList) Append(newval interface{}) {
	list.dataStore = append(list.dataStore, newval)
	list.TheSize++
}

func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.TheSize {
		return nil, errors.New("over bound")
	}
	return list.dataStore[index], nil
}

func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}

func (list *ArrayList) Set(index int, value interface{}) error { return nil }

func (list *ArrayList) Insert(index int, newval interface{}) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("over bound")
	}
	list.checkIsFull()                               // check the storage, auto increase when full
	list.dataStore = list.dataStore[:list.TheSize+1] //insert data ,move 1 storage for 1 unit

	for i := list.TheSize; i > index; i-- { //from back move to front
		list.dataStore[i] = list.dataStore[i-1]

	}

	list.dataStore[index] = newval // insert value
	list.TheSize++

	return nil
}
func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{}, 0, 10) //rearrange storage
	list.TheSize = 0
}

func (list *ArrayList) Delete(index int) error {
	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...)
	return nil
}
