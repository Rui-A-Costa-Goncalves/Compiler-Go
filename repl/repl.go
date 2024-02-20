package repl

import (
	"fmt"
	"goCompiler/token"
	"goCompiler/lexer"
	"io"
	"bufio"
)

func Start(in io.Reader, out io.Writer){
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprintf(out,">> ")
		scanned := scanner.Scan()
		if !scanned{
			return
		} else {
			input := scanner.Text()
			l := lexer.NewLexer(input)
			for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken(){
				fmt.Fprintf(out,"%+v\n",tok)
			}
		}
	}
}