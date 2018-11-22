package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/*
 * 1、获取文件夹下所有文件
 * 2、读取文件所有内容
 * 3、逐行读取文件内容
 */
func main() {
	filepath := "files/"
	files, err := GetAllFile(filepath)
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	fmt.Println("一次性读取文件内容")
	for _, file := range files {
		content, _ := ReadAllFile(filepath + file)
		fmt.Printf("filename:%s,content:\n%s\n", file, content)
	}
	fmt.Println("逐行读取文件内容")
	for _, file := range files {
		lines, _ := ReadFile(filepath + file)
		fmt.Printf("filename:%s\n", file)
		for _, line := range lines {
			fmt.Println(line)
		}
	}

}

// 读取文件目录下所有文件名称
func GetAllFile(pathname string) (files []string, err error) {

	files = make([]string, 0)
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Printf("[%s]\n", pathname+"\\"+fi.Name())
			GetAllFile(pathname + fi.Name() + "\\")
		} else {
			fmt.Println(fi.Name())
			files = append(files, fi.Name())
		}
	}
	return files, err
}

// 一次性读取文件所有内容
func ReadAllFile(filename string) (content string, err error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("read file error ,", err)
		return
	}
	content = string(b)

	return content, err
}

// 逐行读取文件内容
func ReadFile(filename string) (lines []string, err error) {
	lines = make([]string, 0)
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error :%s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		line, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}

		lines = append(lines, string(line))
	}

	return lines, err
}
