package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/wawoon/go-monkey/lexer"
	"github.com/wawoon/go-monkey/token"
)

const PROMPT = ">> "

// Start launches a read–eval–print loop.
//
// All input is read from the given io.Reader and any output – the prompt and
// the inspected tokens – is written to the supplied io.Writer.  Separating the
// writer from os.Stdout makes the REPL easier to test and to embed in other
// programs.
func Start(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)

    for {
        fmt.Fprint(out, PROMPT)

        if !scanner.Scan() {
            // io.EOF or other error -> quit silently; callers can check
            // scanner.Err() if they need the exact reason.
            return
        }

        line := scanner.Text()
        l := lexer.New(line)

        for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
            fmt.Fprintf(out, "%+v\n", tok)
        }
    }
}
