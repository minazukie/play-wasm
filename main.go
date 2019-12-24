package main

import "fmt"

import "syscall/js"

func main() {
	fmt.Println("Start")
	app := js.Global().Get("document").Call("getElementById", "app")
	app.Set("innerHTML", "<h1>Hello!</h1>")
}
