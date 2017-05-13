// Generated from Nugget.g4 by ANTLR 4.7.

package parser // Nugget

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NuggetListener is a complete listener for a parse tree produced by NuggetParser.
type NuggetListener interface {
	antlr.ParseTreeListener

	// EnterNugget is called when entering the nugget production.
	EnterNugget(c *NuggetContext)

	// EnterLoad_stat is called when entering the load_stat production.
	EnterLoad_stat(c *Load_statContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterSourceidentifier is called when entering the sourceidentifier production.
	EnterSourceidentifier(c *SourceidentifierContext)

	// ExitNugget is called when exiting the nugget production.
	ExitNugget(c *NuggetContext)

	// ExitLoad_stat is called when exiting the load_stat production.
	ExitLoad_stat(c *Load_statContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitSourceidentifier is called when exiting the sourceidentifier production.
	ExitSourceidentifier(c *SourceidentifierContext)
}
