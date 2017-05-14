// Generated from Nugget.g4 by ANTLR 4.7.

package parser // Nugget

import "github.com/antlr/antlr4/runtime/Go/antlr"
import "math"
import "fmt"
import "strconv"

var (
	hist []float64
)

// BaseNuggetListener is a complete listener for a parse tree produced by NuggetParser.
type BaseNuggetListener struct{}

var _ NuggetListener = &BaseNuggetListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNuggetListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNuggetListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNuggetListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNuggetListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNugget is called when production nugget is entered.
func (s *BaseNuggetListener) EnterNugget(ctx *NuggetContext) {}

// ExitNugget is called when production nugget is exited.
func (s *BaseNuggetListener) ExitNugget(ctx *NuggetContext) {}

// EnterLoad_stat is called when production load_stat is entered.
func (s *BaseNuggetListener) EnterLoad_stat(ctx *Load_statContext) {}

// ExitLoad_stat is called when production load_stat is exited.
func (s *BaseNuggetListener) ExitLoad_stat(ctx *Load_statContext) {}

// EnterField is called when production field is entered.
func (s *BaseNuggetListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseNuggetListener) ExitField(ctx *FieldContext) {}

// EnterSourceidentifier is called when production sourceidentifier is entered.
func (s *BaseNuggetListener) EnterSourceidentifier(ctx *SourceidentifierContext) {}

// ExitSourceidentifier is called when production sourceidentifier is exited.
func (s *BaseNuggetListener) ExitSourceidentifier(ctx *SourceidentifierContext) {}

// EnterSin is called when production sin is entered.
func (s *BaseNuggetListener) EnterSin(ctx *SinContext) {}

// ExitSin is called when production sin is exited.
func (s *BaseNuggetListener) ExitSin(ctx *SinContext) float64 {
	//x := ctx.parser.GetCurrentToken().GetText()
	y := ctx.NUMBER().GetText()
	x,_ := strconv.ParseFloat(y, 64)
	hist = append(hist,x)
	fmt.Println("x := ", math.Sin(x))
	for _,h := range hist {
		fmt.Println("h := ", math.Sin(h))
	}
	return math.Sin(x)
}
