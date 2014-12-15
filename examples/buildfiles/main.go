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

	builder := jslib.NewBuilder(&out, nil)

	file1 := strings.NewReader(`
package main

import (
   "github.com/gopherjs/gopherjs/js"
)

func a() {
	println(js.Global.Get("window"))
}

`)

	file2 := strings.NewReader(`
package main

import (
   "github.com/gopherjs/gopherjs/js"
)

func main() {
	a()
	println(js.Global.Get("document"))
}

`)

	builder.Add("a.go", file1).Add("b.go", file2)

	err := builder.Build()

	if err != nil {
		fmt.Fprintf(os.Stderr, "can't build: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println(out.String())
}
