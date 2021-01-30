// Package repl -
package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/shanehowearth/interpreter/lexer"
	"github.com/shanehowearth/interpreter/token"
)

const prompt = ">> "

// Start -
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
