package main

import (
	"bytes"
	"fmt"
	"github.com/metakeule/gopherjs/jslib"

	"os"
)

func main() {

	var out bytes.Buffer

	// pass an existing package path here
	err := jslib.BuildPackage("pkg", &out, nil)

	if err != nil {

		fmt.Fprintf(os.Stderr, "compile error: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println(out.String())

}
