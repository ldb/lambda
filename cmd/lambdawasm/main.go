// +build js,wasm

package main

import (
	"fmt"
	"github.com/ldb/lambda/repl"
	"io"
	"os"
	"syscall/js"
)

// https://pingcap.medium.com/how-we-compiled-a-golang-database-in-the-browser-using-webassembly-aba76119678b
// https://www.arp242.net/wasm-cli.html

func main() {
	r, w := io.Pipe()

	js.Global().Set("repl", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go func() {
			if len(args) != 1 {
				return
			}
			input := args[0]
			if input.Type() != js.TypeString {
				return
			}

			w.Write([]byte(input.String()))
		}()
		return nil
	}))
	fmt.Println("Welcome to lambda, a tiny Lambda Calculus Interpreter")
	repl.Start(r, os.Stdout, repl.WASM)
}
