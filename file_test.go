package main

import "testing"

func TestListFiles(t *testing.T) {
	files, err := ListFiles("/")

	if err != nil {
		t.Fatal(err)
	}

	t.Log("Length :", len(files))
	for _, file := range files {
		t.Log("file ", file)
	}
}

func TestDeleteFile(t *testing.T) {
	err := DeleteFile("/Users/ro/godriver/abc")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateDir(t *testing.T) {
	err := CreateDir("/Users/ro/godriver/abc")
	if err != nil {
		t.Fatal(err)
	}
}