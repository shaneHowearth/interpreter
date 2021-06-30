// Package repl -
package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/shanehowearth/interpreter/evaluator"
	"github.com/shanehowearth/interpreter/lexer"
	"github.com/shanehowearth/interpreter/object"
	"github.com/shanehowearth/interpreter/parser"
)

const prompt = ">> "

const monkeyFace = `
           __,__
  .--.  .-"     "-.  .--.
 / .. \/  .-. .-.  \/ .. \
| |  '|  /   Y   \  | ' | |
| \   \  \ 0 | 0 /  /   / |
 \ '- ,\.-"""""""-./, -' /
  ''-' /_   ^ ^   _\ '-''
       | \._   _./ |
       \  \ '~' /  /
       '._ '-=-' _.'
          '-----'
`

// Start -
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	for {
		fmt.Printf(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		if evaluated := evaluator.Eval(program, env); evaluated != nil {

			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, monkeyFace)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
