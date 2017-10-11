package main

// Token is a scanned token
type Token struct{}

// Scanner represents a scanned source string
type Scanner struct{}

// NewScanner creates a scanner from a source string
func NewScanner(source string) *Scanner {
	s := new(Scanner)
	return s
}

func (s *Scanner) scanTokens() []Token {
	return make([]Token, 0)
}
