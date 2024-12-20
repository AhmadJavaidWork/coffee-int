package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/ahmadjavaidwork/coffee-int/evaluator"
	"github.com/ahmadjavaidwork/coffee-int/lexer"
	"github.com/ahmadjavaidwork/coffee-int/object"
	"github.com/ahmadjavaidwork/coffee-int/parser"
)

const PROMPT = ">> "
const COFFE_CUP = `
      )  (
     (   ) )
      ) ( (
 mrf_______)_
 .-'---------|  
( C|/\/\/\/\/|
 '-./\/\/\/\/|
   '_________'
    '-------'
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(PROMPT)
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

		evaulated := evaluator.Eval(program, env)
		if evaulated != nil {
			io.WriteString(out, evaulated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, COFFE_CUP)
	io.WriteString(out, "Woops! We ran into an error here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
