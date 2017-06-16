package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"./nug"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var (
	pathToInput string
	parrotInput bool
)

func init() {
	flag.StringVar(&pathToInput, "input", "", "Path to input")
	flag.BoolVar(&parrotInput, "parrotInput", false, "Repeat parser input")
	flag.Parse()
}

type TreeShapeListener struct {
	*parser.BaseNuggetListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	if parrotInput {
		fmt.Println("Entered a rule, result: " + ctx.GetText())
	}
}

func (this *TreeShapeListener) EnterPrintId(ctx *parser.PrintIdContext) {
	fmt.Println("printid: " + ctx.GetText())
}

func (this *TreeShapeListener) EnterInitextract(ctx *parser.InitextractContext) {
	fmt.Println("extrating from filesystem..")
}

func (this *TreeShapeListener) EnterAssign(ctx *parser.AssignContext) {
	fmt.Println("assigning to variable: ", ctx.StrLit(0).GetText(), " the value : ", ctx.StrLit(1).GetText())
}

func main() {
	var stdin *bufio.Reader

	if pathToInput == "" {
		stdin = bufio.NewReader(os.NewFile(0, "stdin"))
		for {
			fmt.Printf("nugget> ")
			var statement string
			var ok bool
			if statement, ok = readline(stdin); ok {
				fi := antlr.NewInputStream(statement)
				lexer := parser.NewNuggetLexer(fi)
				stream := antlr.NewCommonTokenStream(lexer, 0)
				p := parser.NewNuggetParser(stream)
				p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
				p.BuildParseTrees = true
				tree := p.Nugget()
				antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
			} else {
				break
			}
		}
		// Todo: file input causes an error on newlines
	} else {
		file, err := os.Open(pathToInput)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			input := antlr.NewInputStream(scanner.Text())
			lexer := parser.NewNuggetLexer(input)
			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := parser.NewNuggetParser(stream)
			p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
			p.BuildParseTrees = true
			tree := p.Nugget()
			antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
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
