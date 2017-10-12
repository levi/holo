package scanner

import "github.com/levi/holo/token"

type ScannerError struct {
	s    string
	Line int
}

func (e *ScannerError) Error() string {
	return e.s
}

// Scanner scans a source file for tokens
type Scanner struct {
	Source string
	Tokens []*token.Token
	Errors []ScannerError

	start   int
	current int
	line    int
}

// NewScanner allocates a scanner
func NewScanner(source string) *Scanner {
	s := new(Scanner)
	s.Source = source
	s.start = 0
	s.current = 0
	s.line = 1
	return s
}

// ScanTokens scans source string for tokens
func (s *Scanner) ScanTokens() []*token.Token {
	for !s.isAtEnd() {
		s.start = s.current
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
	default:
		s.Errors = append(s.Errors, ScannerError{"Unexpected character", s.line})
	}
}

func (s *Scanner) addToken(tokenType string) {
	text := s.Source[s.start:s.current]
	s.Tokens = append(s.Tokens, token.NewToken(tokenType, text, "", s.line))
}

// Advance returns the current lexeme character and increments the current offset
func (s *Scanner) advance() string {
	s.current++
	return string(s.Source[s.current-1])
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.Source)
}
