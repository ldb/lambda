package main

import (
	"flag"
	"fmt"
	"github.com/ldb/lambda/repl"
	"os"
)

func main() {
	mLex := flag.Bool("lex", false, "print lexemes")
	mAST := flag.Bool("ast", false, "print ast")
	flag.Parse()

	var m repl.Mode
	if *mLex {
		m = m.Set(repl.PrintLexemes)
	}
	if *mAST {
		m = m.Set(repl.PrintAST)
	}

	fmt.Println("Welcome to lambda, a tiny Lambda Calculus Interpreter")
	repl.Start(os.Stdin, os.Stdout, m)
}
