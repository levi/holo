package scanner

import (
	"fmt"
	"strconv"

	"github.com/levi/holo/token"
)

var keywords = map[string]string{
	"and":    token.AndToken,
	"class":  token.ClassToken,
	"else":   token.ElseToken,
	"false":  token.FalseToken,
	"for":    token.ForToken,
	"fn":     token.FnToken,
	"if":     token.IfToken,
	"nil":    token.NilToken,
	"or":     token.OrToken,
	"print":  token.PrintToken,
	"return": token.ReturnToken,
	"self":   token.SelfToken,
	"super":  token.SuperToken,
	"true":   token.TrueToken,
	"var":    token.VarToken,
	"while":  token.WhileToken,
}

// Scanner scans a source file for tokens
type Scanner struct {
	Source string
	Tokens []*token.Token
	Errors []ScannerError

	start  int
	cursor int
	line   int
}

// NewScanner allocates a scanner
func NewScanner(source string) *Scanner {
	s := new(Scanner)
	s.Source = source
	s.start = 0
	s.cursor = 0
	s.line = 1
	return s
}

// ScanTokens scans source string for tokens
func (s *Scanner) ScanTokens() []*token.Token {
	for !s.isAtEnd() {
		s.start = s.cursor
		s.scanToken()
	}

	s.Tokens = append(s.Tokens, token.NewToken(token.EOFToken, "", "", s.line))
	return s.Tokens
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case "(":
		s.addToken(token.LeftParenToken)
	case ")":
		s.addToken(token.RightParenToken)
	case "{":
		s.addToken(token.LeftBraceToken)
	case "}":
		s.addToken(token.RightBraceToken)
	case ",":
		s.addToken(token.CommaToken)
	case ".":
		s.addToken(token.DotToken)
	case "-":
		s.addToken(token.MinusToken)
	case "+":
		s.addToken(token.PlusToken)
	case ";":
		s.addToken(token.SemicolonToken)
	case "*":
		s.addToken(token.StarToken)
	case "!":
		if s.match("=") {
			s.addToken(token.BangEqualToken)
		} else {
			s.addToken(token.BangToken)
		}
	case "=":
		if s.match("=") {
			s.addToken(token.EqualEqualToken)
		} else {
			s.addToken(token.EqualToken)
		}
	case "<":
		if s.match("=") {
			s.addToken(token.LessEqualToken)
		} else {
			s.addToken(token.LessToken)
		}
	case ">":
		if s.match("=") {
			s.addToken(token.GreaterEqualToken)
		} else {
			s.addToken(token.GreaterToken)
		}
	case "/":
		if s.match("/") {
			for s.peek() != "\n" && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(token.SlashToken)
		}
	case "\"":
		s.string()
	case " ":
	case "\r":
	case "\t":
		// Ignore whitespace
		break
	case "\n":
		s.line++
		break
	default:
		if isDigit(c) {
			s.number()
		} else if isAlpha(c) {
			s.identifier()
		} else {
			s.raiseError(fmt.Sprintf("Unexpected character: \"%s\"", c))
		}
	}
}

func (s *Scanner) string() {
	for s.peek() != "\"" && !s.isAtEnd() {
		if s.peek() == "\n" {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		s.raiseError("Unterminated string")
		return
	}

	// consume the closing "
	s.advance()

	value := s.Source[s.start+1 : s.cursor-1]
	s.addTokenLiteral(token.StringToken, value)
}

func (s *Scanner) number() {
	for isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == "." && isDigit(s.peekNext()) {
		s.advance() // consume the .
		for isDigit(s.peek()) {
			s.advance()
		}
	}

	n, err := strconv.ParseFloat(s.Source[s.start:s.cursor], 64)
	if err != nil {
		s.raiseError("Failed to parse number literal")
		return
	}

	s.addTokenLiteral(token.NumberToken, n)
}

func (s *Scanner) identifier() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}

	value := s.Source[s.start:s.cursor]
	if t, ok := keywords[value]; ok {
		s.addToken(t)
	} else {
		s.addToken(token.IdentifierToken)
	}
}

func isAlphaNumeric(value string) bool {
	return isAlpha(value) || isDigit(value)
}

func isAlpha(value string) bool {
	for _, c := range value {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || ('_' == c) {
			return true
		}
	}
	return false
}

func isDigit(value string) bool {
	for _, c := range value {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}

func (s *Scanner) addToken(tokenType string) {
	s.addTokenLiteral(tokenType, nil)
}

func (s *Scanner) addTokenLiteral(tokenType string, literal interface{}) {
	text := s.Source[s.start:s.cursor]
	s.Tokens = append(s.Tokens, token.NewToken(tokenType, text, literal, s.line))
}

// raiseError appends an error with description at the current line to the Errors slice
func (s *Scanner) raiseError(description string) {
	s.Errors = append(s.Errors, ScannerError{description, s.line})
}

// Advance returns the current lexeme character and increments the cursor offset
func (s *Scanner) advance() string {
	s.cursor++
	return string(s.Source[s.cursor-1])
}

// Match identifies if the current character is the expected, advancing the cursor when match succeeds
func (s *Scanner) match(expected string) bool {
	if s.isAtEnd() {
		return false
	}
	if s.Source[s.cursor] != expected[0] {
		return false
	}

	s.cursor++
	return true
}

// Peek provides a lookahead of 1 of the current cursor position without advancing
func (s *Scanner) peek() string {
	if s.isAtEnd() {
		return "\x00"
	}
	return string(s.Source[s.cursor])
}

// peekNext provides a lookahead of 1 beyond the current cursor position without advancing
func (s *Scanner) peekNext() string {
	next := s.cursor + 1

	if next >= len(s.Source) {
		return "\x00"
	}

	return string(s.Source[next])
}

// isAtEnd determines if the cursor position is beyond the source's bounds
func (s *Scanner) isAtEnd() bool {
	return s.cursor >= len(s.Source)
}
