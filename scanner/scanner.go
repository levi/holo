package scanner

import "github.com/levi/holo/token"

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
	default:
		s.Errors = append(s.Errors, ScannerError{"Unexpected character", s.line})
	}
}

func (s *Scanner) addToken(tokenType string) {
	text := s.Source[s.start:s.cursor]
	s.Tokens = append(s.Tokens, token.NewToken(tokenType, text, "", s.line))
}

// Advance returns the current lexeme character and increments the cursor offset
func (s *Scanner) advance() string {
	s.cursor++
	return string(s.Source[s.cursor-1])
}

// Match identifies if the current character is the expected, advancing the cursor when positive
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

func (s *Scanner) peek() string {
	if s.isAtEnd() {
		return "\x00"
	}
	return string(s.Source[s.cursor])
}

func (s *Scanner) isAtEnd() bool {
	return s.cursor >= len(s.Source)
}
