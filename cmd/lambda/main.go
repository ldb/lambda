package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/ldb/lambda/lexer"
	"github.com/ldb/lambda/parser"
	"io"
	"os"
)

const prompt = "λ > "

type mode uint8

const (
	none         mode = 0
	printLexemes mode = 1 << iota
	printAST
)

func (m mode) Set(mode mode) mode   { return m | mode }
func (m mode) IsSet(mode mode) bool { return m&mode != 0 }

func main() {
	mLex := flag.Bool("lex", false, "print lexemes")
	mAST := flag.Bool("ast", false, "print ast")
	flag.Parse()

	var m mode
	if *mLex {
		m = m.Set(printLexemes)
	}
	if *mAST {
		m = m.Set(printAST)
	}

	fmt.Println("Welcome to Leo's tiny Lambda Calculus Interpreter")
	startREPL(os.Stdin, os.Stdout, m)
}

func startREPL(in io.Reader, out io.Writer, mode mode) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		term := p.ParseLambdaTerm()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		if mode.IsSet(printLexemes) {
			fmt.Fprintf(out, "lex: not implemented\n")
		}

		if mode.IsSet(printAST) {
			fmt.Fprintf(out, "ast: %+t\n", term)
		}

		io.WriteString(out, term.String())
		io.WriteString(out, "\n")
	}
}

func printParseErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
