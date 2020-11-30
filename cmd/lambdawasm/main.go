package main

import (
	"fmt"
	"io"
	"os"
)

// https://pingcap.medium.com/how-we-compiled-a-golang-database-in-the-browser-using-webassembly-aba76119678b
// https://www.arp242.net/wasm-cli.html

func main() {
	fmt.Println("Hello, WebAssembly!")

	io.Copy(os.Stdout, os.Stdin)
}
