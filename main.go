package main

import (
	"fmt"

	"example.com/m/ArrayList"
	"example.com/m/CircleQuene"
	"example.com/m/Link"
	"example.com/m/Quene"
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

// Quene 隊列
//隊列是廣度遍歷
//棧 Stack,遞迴 是深度遍歷
func main7() {
	myQuene := Quene.NewQuene()
	myQuene.EnQueue(1)
	myQuene.EnQueue(2)
	myQuene.EnQueue(3)
	myQuene.EnQueue(4)
	fmt.Println(myQuene.DeQueue())
	fmt.Println(myQuene.DeQueue())
	fmt.Println(myQuene.DeQueue())
	fmt.Println(myQuene.DeQueue())
	fmt.Println(myQuene.DeQueue())
}

//CircleQuene
/*
循環隊列,環狀結構
不需要數組移動,效率較高
*/
func main8() {
	var myCircleQuene CircleQuene.CircleQuene
	CircleQuene.InitQueue(&myCircleQuene)
	CircleQuene.EnQueue(&myCircleQuene, 1)
	CircleQuene.EnQueue(&myCircleQuene, 2)
	CircleQuene.EnQueue(&myCircleQuene, 3)
	CircleQuene.EnQueue(&myCircleQuene, 4)
	fmt.Println(CircleQuene.DeQueue(&myCircleQuene))
	fmt.Println(CircleQuene.DeQueue(&myCircleQuene))
	fmt.Println(CircleQuene.DeQueue(&myCircleQuene))
	fmt.Println(CircleQuene.DeQueue(&myCircleQuene))
	fmt.Println(CircleQuene.DeQueue(&myCircleQuene))
}

//Link

func main9() {
	node1 := new(Link.Node)
	node2 := new(Link.Node)
	node3 := new(Link.Node)
	node4 := new(Link.Node)
	node5 := new(Link.Node)

	node1.Data = 1
	node1.PNext = node2
	node2.Data = 2
	node2.PNext = node3
	node3.Data = 3
	node3.PNext = node4
	node4.Data = 4
	node4.PNext = node5
	node5.Data = 5
	fmt.Println(node1.Data)
}

//Link Stack
func main() {
	mystack := Link.NewStack()
	for i := 0; i < 1000; i++ {
		mystack.Push(i)
	}
	for data := mystack.Pop(); data != nil; data = mystack.Pop() {
		fmt.Println(data)
	}
}
