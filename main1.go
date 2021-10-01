package main

import (
	"fmt"

	"example.com/m/ArrayList"
	"example.com/m/StackArray.go"
)

// ArrayList
func main1() {
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append(1)
	list.Append(3)
	list.Append(2)
	fmt.Println(list)
}
func main2() {
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append(1)
	list.Append(3)
	list.Append(2)
	for i := 0; i < 30; i++ {

		list.Insert(1, 5)

		fmt.Println(list)
	}

}
func main3() {
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append("1")
	list.Append("2")
	list.Append("3")
	list.Append("4")
	list.Append("5")
	list.Append("6")
	for it := list.Iterator(); it.HasNext(); {
		item, _ := it.Next()
		if item == "4" {
			it.Remove()
			break
		}
		fmt.Println(item)
	}
	fmt.Println(list)
}

//StackArray
func main4() {
	mystack := StackArray.NewStack()
	mystack.Push(1)
	mystack.Push(2)
	mystack.Push(3)
	mystack.Push(4)
	mystack.Push(5)
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
}

// ArrayListStack
func main5() {
	mystack := ArrayList.NewArrayListStack()
	mystack.Push(1)
	mystack.Push(2)
	mystack.Push(3)
	mystack.Push(4)
	mystack.Push(5)
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
}

// ArrayListStackIterator
func main6() {
	mystack := ArrayList.NewArrayListStackX()
	mystack.Push(1)
	mystack.Push(2)
	mystack.Push(3)
	mystack.Push(4)
	mystack.Push(5)

	for it := mystack.Myit; it.HasNext(); {
		item, _ := it.Next()
		if item == 4 {
			fmt.Println(item)
			break
		}
		//fmt.Println(item)
	}

}
