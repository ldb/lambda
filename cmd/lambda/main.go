package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/ldb/lambda/lexer"
	"github.com/ldb/lambda/parser"
	"io"
	"os"
	"strings"
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
		trimmed := strings.Trim(line, " \n\t")
		l := lexer.New(trimmed)
		p := parser.New(l)

		term := p.ParseLambdaTerm()
		if err := p.Error(); err != nil {
			printParseErrors(out, err)
			continue
		}

		if mode.IsSet(printLexemes) {
			fmt.Fprintf(out, "lex: not implemented\n")
		}

		if mode.IsSet(printAST) {
			b, err := json.Marshal(term)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(b))
		}

		io.WriteString(out, term.String())
		io.WriteString(out, "\n")
	}
}

func printParseErrors(out io.Writer, err error) {
	var pErr *parser.ParseError
	if errors.As(err, &pErr) {
		lp := strings.Repeat(" ", len(prompt)+pErr.Position-1)
		io.WriteString(out, lp+"^ ")
	}
	io.WriteString(out, err.Error()+"\n")
}
