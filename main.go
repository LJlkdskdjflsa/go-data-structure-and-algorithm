package main

import (
	"fmt"

	"example.com/m/ArrayList"
)

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
func main() {
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
