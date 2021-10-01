package main

import (
	"errors"
	"fmt"
	"io/ioutil"
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
func main() {
	path := "/home/lj/Pictures/"
	files := []string{}
	files, _ = GetAll(path, files)
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}
