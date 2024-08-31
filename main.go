package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// any path we not define will go to here
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("have request /")

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			return
		}

		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		fmt.Fprintf(w, "Server Work")
	})

	mux.HandleFunc("GET /files", ListFilesHandle)
	mux.HandleFunc("POST /files/delete", DeleteFileHandle)
	mux.HandleFunc("POST /files/create", CreateDirHandle)

	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
