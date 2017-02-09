package repl

import (
	"bufio"
	"fmt"
	"io"
	"github.com/adamdecaf/gointerp/lexer"
	"github.com/adamdecaf/gointerp/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		// Exit if the user asks us to.
		if line == "exit" || line == "quit" {
			return
		}

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
