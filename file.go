package main

import (
	"log"
	"os"
)

type FileInfo struct {
	Name  string `json:"name"`
	IsDir bool   `json:"is_dir"`
}

func ListFiles(path string) ([]FileInfo, error) {
	var filesInfo []FileInfo

	files, err := os.ReadDir(path)

	if err != nil {
		return filesInfo, err
	}
	for _, file := range files {
		filesInfo = append(filesInfo, FileInfo{
			Name:  file.Name(),
			IsDir: file.IsDir(),
		})
	}

	log.Println("filesInfo ", filesInfo)

	return filesInfo, nil
}

func DeleteFile(path string) error {
	return os.RemoveAll(path)
}

func CreateDir(path string) error {
	return os.MkdirAll(path, 0777)
}