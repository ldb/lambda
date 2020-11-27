package repl

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ldb/lambda/lexer"
	"github.com/ldb/lambda/parser"
	"io"
	"strings"
)

const prompt = "λ > "

type Mode uint8

const (
	None         Mode = 0
	PrintLexemes Mode = 1 << iota
	PrintAST
)

func (m Mode) Set(mode Mode) Mode   { return m | mode }
func (m Mode) IsSet(mode Mode) bool { return m&mode != 0 }

func Start(in io.Reader, out io.Writer, mode Mode) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, prompt)
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

		if mode.IsSet(PrintLexemes) {
			fmt.Println("lex: not implemented")
		}

		if mode.IsSet(PrintAST) {
			b, err := json.Marshal(term)
			if err != nil {
				fmt.Fprintf(out, "ast: %s\n", err.Error())
			}
			fmt.Fprintf(out, "ast: %s\n", b)
		}

		fmt.Fprintln(out, term.String())
	}
}

func printParseErrors(out io.Writer, err error) {
	var pErr *parser.ParseError
	if errors.As(err, &pErr) {
		lp := strings.Repeat(" ", len(prompt)+pErr.Position-1)
		fmt.Fprint(out, lp+"^ ")
	}
	fmt.Fprintln(out, err.Error())
}
