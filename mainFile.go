package main

import (
	"errors"
	"fmt"
	"io/ioutil"

	"example.com/m/StackArray.go"
)

func GetAll(path string, files []string) ([]string, error) {
	read, err := ioutil.ReadDir(path)
	if err != nil {
		return files, errors.New("文件夾不可讀取")
	}
	for _, fi := range read {
		if fi.IsDir() {
			if fi.Name() == ".git" {
				continue
			}
			fulldir := path + "/" + fi.Name()
			fmt.Println(fi.Name())
			files = append(files, fulldir)
			files, _ = GetAll(path, files)
		} else {
			fulldir := path + "/" + fi.Name()
			files = append(files, fulldir)
		}
	}
	return files, nil
}

// use recurrent to read files
func main1a() {
	path := "/home/lj/Pictures/"
	files := []string{}
	files, _ = GetAll(path, files)
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}

// use stack to read files
func main2a() {
	path := "/home/lj/GolandProjects"
	files := []string{}
	mystack := StackArray.NewStack()
	mystack.Push(path)
	for !mystack.IsEmpty() {
		getPath := mystack.Pop().(string) //取出後實例化成string
		files = append(files, getPath)
		read, _ := ioutil.ReadDir(getPath) // 讀取文件夾下所有路徑
		for _, fi := range read {
			if fi.IsDir() {
				fulldir := getPath + "/" + fi.Name() //構造新路徑
				//files = append(files, fulldir)
				mystack.Push(fulldir)
			} else {
				fulldir := getPath + "/" + fi.Name()
				files = append(files, fulldir)
			}
		}
	}

	//打印
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}
