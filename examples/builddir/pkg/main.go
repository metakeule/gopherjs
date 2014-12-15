package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	println(js.Global.Get("window"))
}
