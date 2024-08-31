package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResJson struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
	Error  string `json:"error"`
}

func MakeJsonRes(w http.ResponseWriter, dataBody any) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	resJson := &ResJson{
		Status: "ok",
		Data:   dataBody,
	}

	resByte, err := json.Marshal(resJson)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(resByte)
}

func MakeErrRes(w http.ResponseWriter, errMess string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	resJson := &ResJson{
		Status: "error",
		Error:  errMess,
	}

	resByte, err := json.Marshal(resJson)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write(resByte)
}

func ListFilesHandle(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")
	if path == "" {
		MakeErrRes(w, "path require")
		return
	}

	files, err := ListFiles(path)
	if err != nil {
		MakeErrRes(w, err.Error())
		return
	}

	MakeJsonRes(w, files)
}

func DeleteFileHandle(w http.ResponseWriter, r *http.Request) {
	
	log.Println("On call delete")
	
	path := r.URL.Query().Get("path")
	if path == "" {
		MakeErrRes(w, "path require")
		return
	}

	err := DeleteFile(path)
	if err != nil {
		MakeErrRes(w, err.Error())
		return
	}

	MakeJsonRes(w, nil)
}

func CreateDirHandle(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")
	if path == "" {
		MakeErrRes(w, "path require")
		return
	}

	err := CreateDir(path)
	if err != nil {
		MakeErrRes(w, err.Error())
		return
	}

	MakeJsonRes(w, nil)
}