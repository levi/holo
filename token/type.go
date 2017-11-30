package token

const (
	// Single character tokens
	LeftParenToken  = "LeftParen"
	RightParenToken = "RightParen"
	LeftBraceToken  = "LeftBrace"
	RightBraceToken = "RightBrace"
	CommaToken      = "Comma"
	DotToken        = "Dot"
	MinusToken      = "Minus"
	PlusToken       = "Plus"
	SemicolonToken  = "Semicolon"
	SlashToken      = "Slash"
	StarToken       = "Star"

	// One or two character tokens
	BangToken         = "Bang"
	BangEqualToken    = "BangEqual"
	EqualToken        = "Equal"
	EqualEqualToken   = "EqualEqual"
	GreaterToken      = "Greater"
	GreaterEqualToken = "GreaterEqual"
	LessToken         = "Less"
	LessEqualToken    = "LessEqual"

	// Literals
	IdentifierToken = "Identifier"
	StringToken     = "String"
	NumberToken     = "Number"

	// Keywords
	AndToken    = "And"
	ClassToken  = "Class"
	ElseToken   = "Else"
	FalseToken  = "False"
	FnToken     = "Fn"
	ForToken    = "For"
	IfToken     = "If"
	NilToken    = "Nil"
	OrToken     = "Or"
	PrintToken  = "Print"
	ReturnToken = "Return"
	SelfToken   = "Self"
	SuperToken  = "Super"
	TrueToken   = "True"
	VarToken    = "Var"
	WhileToken  = "While"

	EOFToken = "EOF"
)
