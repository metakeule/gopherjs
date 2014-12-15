// +build js

package syscall

import (
	"bytes"
	"gopkg.in/metakeule/gopherjs/js"
)

var warningPrinted = false
var lineBuffer []byte

func printWarning() {
	if !warningPrinted {
		println("warning: system calls not available, see https://github.com/gopherjs/gopherjs/blob/master/doc/syscalls.md")
	}
	warningPrinted = true
}

func printToConsole(b []byte) {
	lineBuffer = append(lineBuffer, b...)
	for {
		i := bytes.IndexByte(lineBuffer, '\n')
		if i == -1 {
			break
		}
		js.Global.Get("console").Call("log", string(lineBuffer[:i])) // don't use println, since it does not externalize multibyte characters
		lineBuffer = lineBuffer[i+1:]
	}
}
