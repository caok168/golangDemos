package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	list, err := getDirInfoList("/home/ck/go/src/github.com/lynxitech/ivs/volumes/snaps")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range list {
		fmt.Println(v)
	}

}

func getDirList(dirpath string) ([]string, error) {
	var dir_list []string
	dir_err := filepath.Walk(dirpath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				dir_list = append(dir_list, path)
				return nil
			}

			return nil
		})
	return dir_list, dir_err
}

func getDirInfoList(dirpath string) ([]FolderInfo, error) {
	var dir_list []FolderInfo
	dir_err := filepath.Walk(dirpath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				folderInfo := FolderInfo{
					Path: path,
					Name: f.Name(),
				}
				dir_list = append(dir_list, folderInfo)
				return nil
			}

			return nil
		})
	return dir_list, dir_err
}

type FolderInfo struct {
	Path string `json:"path"`
	Name string `json:"name"`
}
