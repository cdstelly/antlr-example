// Generated from Nugget.g4 by ANTLR 4.7.

package parser // Nugget

import "github.com/antlr/antlr4/runtime/Go/antlr"
import "fmt"


// BaseNuggetListener is a complete listener for a parse tree produced by NuggetParser.
type BaseNuggetListener struct{}

var _ NuggetListener = &BaseNuggetListener{}
var History map[string]string

func init () {
	History = make(map[string]string)
}

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

// EnterInitextract is called when production initextract is entered.
func (s *BaseNuggetListener) EnterInitextract(ctx *InitextractContext) {
        id := ctx.Id().GetText()
	val := ctx.Target().GetText() + ":" + ctx.Subtype().GetText()
	History[id] = val
}

// ExitInitextract is called when production initextract is exited.
func (s *BaseNuggetListener) ExitInitextract(ctx *InitextractContext) {}

// EnterExecute is called when production execute is entered.
func (s *BaseNuggetListener) EnterExecute(ctx *ExecuteContext) {}

// ExitExecute is called when production execute is exited.
func (s *BaseNuggetListener) ExitExecute(ctx *ExecuteContext) {}

// EnterSubtype is called when production subtype is entered.
func (s *BaseNuggetListener) EnterSubtype(ctx *SubtypeContext) {}

// ExitSubtype is called when production subtype is exited.
func (s *BaseNuggetListener) ExitSubtype(ctx *SubtypeContext) {}

// EnterTask is called when production task is entered.
func (s *BaseNuggetListener) EnterTask(ctx *TaskContext) {}

// ExitTask is called when production task is exited.
func (s *BaseNuggetListener) ExitTask(ctx *TaskContext) {}

// EnterTarget is called when production target is entered.
func (s *BaseNuggetListener) EnterTarget(ctx *TargetContext) {}

// ExitTarget is called when production target is exited.
func (s *BaseNuggetListener) ExitTarget(ctx *TargetContext) {}

// EnterField is called when production field is entered.
func (s *BaseNuggetListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseNuggetListener) ExitField(ctx *FieldContext) {}

// EnterSourceidentifier is called when production sourceidentifier is entered.
func (s *BaseNuggetListener) EnterSourceidentifier(ctx *SourceidentifierContext) {}

// ExitSourceidentifier is called when production sourceidentifier is exited.
func (s *BaseNuggetListener) ExitSourceidentifier(ctx *SourceidentifierContext) {}

// EnterPrintId is called when production printId is entered.
func (s *BaseNuggetListener) EnterPrintId(ctx *PrintIdContext) {
	id := ctx.Id().GetText()
	_,ok := History[id]
	if ok {
		fmt.Println(History[id])
	} else {
		fmt.Println("Key: ", id, " has no history")
	}
}

// ExitPrintId is called when production printId is exited.
func (s *BaseNuggetListener) ExitPrintId(ctx *PrintIdContext) {}

// EnterSin is called when production sin is entered.
func (s *BaseNuggetListener) EnterSin(ctx *SinContext) {}

// ExitSin is called when production sin is exited.
func (s *BaseNuggetListener) ExitSin(ctx *SinContext) {}
