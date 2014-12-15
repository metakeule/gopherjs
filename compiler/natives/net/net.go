// +build js

package net

import (
	"gopkg.in/metakeule/gopherjs/js"
)

func Listen(net, laddr string) (Listener, error) {
	js.Global.Call("$notSupported", "net")
	panic("unreachable")
}
