// Generated from Nugget.g4 by ANTLR 4.7.

package parser // Nugget

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NuggetListener is a complete listener for a parse tree produced by NuggetParser.
type NuggetListener interface {
	antlr.ParseTreeListener

	// EnterNugget is called when entering the nugget production.
	EnterNugget(c *NuggetContext)

	// EnterInitextract is called when entering the initextract production.
	EnterInitextract(c *InitextractContext)

	// EnterExecute is called when entering the execute production.
	EnterExecute(c *ExecuteContext)

	// EnterSubtype is called when entering the subtype production.
	EnterSubtype(c *SubtypeContext)

	// EnterTask is called when entering the task production.
	EnterTask(c *TaskContext)

	// EnterTarget is called when entering the target production.
	EnterTarget(c *TargetContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterSourceidentifier is called when entering the sourceidentifier production.
	EnterSourceidentifier(c *SourceidentifierContext)

	// EnterPrintId is called when entering the printId production.
	EnterPrintId(c *PrintIdContext)

	// EnterSin is called when entering the sin production.
	EnterSin(c *SinContext)

	// ExitNugget is called when exiting the nugget production.
	ExitNugget(c *NuggetContext)

	// ExitInitextract is called when exiting the initextract production.
	ExitInitextract(c *InitextractContext)

	// ExitExecute is called when exiting the execute production.
	ExitExecute(c *ExecuteContext)

	// ExitSubtype is called when exiting the subtype production.
	ExitSubtype(c *SubtypeContext)

	// ExitTask is called when exiting the task production.
	ExitTask(c *TaskContext)

	// ExitTarget is called when exiting the target production.
	ExitTarget(c *TargetContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitSourceidentifier is called when exiting the sourceidentifier production.
	ExitSourceidentifier(c *SourceidentifierContext)

	// ExitPrintId is called when exiting the printId production.
	ExitPrintId(c *PrintIdContext)

	// ExitSin is called when exiting the sin production.
	ExitSin(c *SinContext)
}
