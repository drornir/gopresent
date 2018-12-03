package main

import "github.com/dennwc/dom/js"

func writeToBody(s string) {
	document := js.Get("document")
	body := document.Call("getElementById", "app")
	body.Set("innerHTML", s)
}

func main() {
	writeToBody("<b>Hi Wasm!</b>")
}
