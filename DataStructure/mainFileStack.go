package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func GetAllX(path string, files []string, level int) ([]string, error) {
	levelstr := ""
	if level == 1 {
		levelstr = "+"
	} else {
		levelstr = "|" + strings.Repeat("--", level) + "+"
	}
	read, err := ioutil.ReadDir(path)

	if err != nil {
		return files, errors.New("文件夾不可讀取")
	}
	//fmt.Println(read)
	for _, fi := range read {
		if fi.IsDir() {
			fulldir := path + "/" + fi.Name()
			files = append(files, fulldir)
			fmt.Println(level)
			files, _ = GetAllX(fulldir, files, level+1) //文件遞歸
		} else {
			fulldir := path + "/" + fi.Name()
			files = append(files, levelstr+fulldir)
		}
	}
	return files, nil
}

// oringin
func main1b() {
	path := "/home/lj/Pictures"
	files := []string{}
	files, _ = GetAllX(path, files, 1)

	for i := 0; i < len(files); i++ {

		fmt.Println(files[i])
	}
}

//求深度
func main2b() {}
