package main

import (
	"fmt"
	"os"
)

func main() {
	//createFile()
	//mkdir()
	//remove()
	//getWd()
}

// 创建一个文件
func createFile() {
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f.Name())
	}

}

// 创建目录
func mkdir() {
	//第二个参数是权限os.ModePerm代表最高权限777
	err := os.Mkdir("test", os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	//创建多个目录
	err2 := os.MkdirAll("test/a/b", os.ModePerm)
	if err2 != nil {
		return
	}

}

// 删除目录
func remove() {
	err := os.Remove("a.txt")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("test")
}

// 获得工作目录
func getWd() {
	dir, _ := os.Getwd()
	fmt.Println(dir)
	//改变目录
	os.Chdir("d:/")
}
