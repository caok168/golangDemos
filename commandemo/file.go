package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	commandLines, err := GetFileContent("files/", 3)
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	for _, line := range commandLines {
		fmt.Println(line)
	}
}

// 读取文件名大于num的命令行内容
func GetFileContent(pathname string, num int) (commandLines []string, err error) {
	commandLines = make([]string, 0)

	filenames, err := GetAllFiles(pathname)
	if err != nil {
		return nil, err
	}
	executeFiles := make([]string, 0)
	for _, filename := range filenames {
		filevalue, err := strconv.Atoi(filename)
		if err != nil {
			return nil, err
		}
		if filevalue > num {
			executeFiles = append(executeFiles, filename)
		}
	}

	for _, executeFile := range executeFiles {
		lines, err := ReadFileContent(pathname + executeFile)
		if err != nil {
			return nil, err
		}
		commandLines = append(commandLines, lines...)
	}

	return commandLines, nil
}

// 读取文件目录下所有文件名称
func GetAllFiles(pathname string) (files []string, err error) {

	files = make([]string, 0)
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			//fmt.Printf("[%s]\n", pathname+"\\"+fi.Name())
			GetAllFiles(pathname + fi.Name() + "\\")
		} else {
			//fmt.Println(fi.Name())
			files = append(files, fi.Name())
		}
	}
	return files, err
}

// 逐行读取文件内容
func ReadFileContent(filename string) (lines []string, err error) {
	lines = make([]string, 0)
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error :%s\n", err)
		return nil, err
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
