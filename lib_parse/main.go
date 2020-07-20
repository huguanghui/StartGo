package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 1. 轮训当前目录下的所有库, 排序, 静态库在前, 动态库在后
func getFileListInfoByPath(path string, file *[]string, targetType *[]string, ignoreFile *[]string, ignorePath *[]string, ignoreType *[]string) (err error) {
	l, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	sep := string(os.PathSeparator)
	for _, f := range l {
		tmp := string(path + sep + f.Name())
		fmt.Println(tmp)
	}
	return nil
}

// 2. 获取每个库的md5信息和大小(KB)表示
// 3. 将结果保存到文件中

func main() {
	targetType := []string{".so", ".a"}
	ignoreFile := []string{}
	ignorePath := []string{}
	ignoreType := []string{}

	var files []string
	// 获取当前目录的绝对路径
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = getFileListInfoByPath(path, &files, &targetType, &ignoreFile, &ignorePath, &ignoreType)
	if err != nil {
		fmt.Printf(err.Error() + "\n")
	}
}
