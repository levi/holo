package scanner

import "github.com/levi/holo/token"

// Scanner scans a source file for tokens
type Scanner struct {
	Source string
	Tokens []*token.Token

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
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.Source)
}
