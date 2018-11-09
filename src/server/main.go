package main

import (
	"fmt"
	"net/http"
	"path/filepath"
)

func main() {
	fmt.Println("starting file server on 8080")
	fmt.Println(filepath.Abs("."))
	dir := http.Dir("public")
	http.ListenAndServe(":8080", http.FileServer(dir))
}
