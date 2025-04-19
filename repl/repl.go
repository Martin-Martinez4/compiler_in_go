package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Martin-Martinez4/compiler_in_go/evaluator"
	"github.com/Martin-Martinez4/compiler_in_go/lexer"
	"github.com/Martin-Martinez4/compiler_in_go/object"
	"github.com/Martin-Martinez4/compiler_in_go/parser"
)

const PROMPT = "-> "

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops! we ran into some monkey business here!")
	io.WriteString(out, " parser errors: \n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {

			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}
