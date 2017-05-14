package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"./nug"
	"os"
	"fmt"
	"bufio"
)

type TreeShapeListener struct {
	*parser.BaseNuggetListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	// fmt.Println(ctx.GetText())
}

func main() {
	//fi, _ := antlr.NewFileStream(os.Args[1])
	stdin := bufio.NewReader(os.NewFile(0, "stdin"))
	for {
		fmt.Printf("nugget> ")
		var statement string
		var ok bool
		if statement, ok = readline(stdin); ok {

			fi := antlr.NewInputStream(statement)

			lexer := parser.NewNuggetLexer(fi)
			stream := antlr.NewCommonTokenStream(lexer,0)
			p := parser.NewNuggetParser(stream)
			p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
			p.BuildParseTrees = true	//necessary?
			tree := p.Nugget()
			antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
		} else {
			break
		}
	}
}

func readline(fi *bufio.Reader) (string, bool) {
	s, err := fi.ReadString('\n')
	if err != nil {
		return "", false
	}
	return s, true
}
