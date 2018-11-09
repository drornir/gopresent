package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("starting file server on 8080")
	dir := http.Dir("public")
	http.ListenAndServe(":8080", http.FileServer(dir))
}
