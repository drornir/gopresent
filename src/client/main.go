package main

import (
	"fmt"

	"github.com/dennwc/dom/js"
)

func main() {
	fmt.Println("hi wasm")
	writeToBody(`
	<div id='app'>
		Hi Wasm!
	</div>
	`)
}

func writeToBody(s string) {
	document := js.Get("document")

	body := document.Call("getElementById", "body")
	body.Set("innerHTML", s)
}
