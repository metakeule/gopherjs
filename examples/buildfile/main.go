package main

import (
	"bytes"
	"fmt"
	"github.com/metakeule/gopherjs/jslib"
	"os"
	"strings"
)

func main() {
	var out bytes.Buffer

	goCode := strings.NewReader(`
package main

import (
   "github.com/gopherjs/gopherjs/js"
)

func main() {
	println(js.Global.Get("window"))
	js.Global.Get("window").Get("alert").Invoke("huhu")
}

`)

	err := jslib.Build(goCode, &out, nil)

	if err != nil {

		fmt.Fprintf(os.Stderr, "compile error: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println(out.String())
}
