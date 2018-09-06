package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/levi/holo/expr"
	"github.com/levi/holo/parser"
	"github.com/levi/holo/scanner"
	"github.com/levi/holo/token"
)

var hadError = false

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: holo [script]")
	} else if len(os.Args) == 2 {
		e := runFile(os.Args[0])
		if e != nil {
			panic(e)
		}
	} else {
		runPrompt()
	}
}

func runFile(path string) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	run(string(bytes))
	if hadError {
		return errors.New("Parsing failed")
	}
	return nil
}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if scanner.Scan() {
			run(scanner.Text())
			hadError = false
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
		}
	}
}

func run(source string) {
	s := scanner.NewScanner(source)
	tokens := s.ScanTokens()
	errors := s.Errors

	if len(errors) > 0 {
		for _, err := range errors {
			fmt.Fprintln(os.Stderr, "Scanner error:", err)
		}
	}

	for _, token := range tokens {
		fmt.Println(token.String())
	}

	p := parser.NewParser(tokens)
	expression, err := p.Parse()

	if err, ok := err.(*parser.ParseError); ok {
		reportParseError(*err)
		return
	}

	fmt.Println(expression.(expr.AstPrinter).ToString())
}

func reportError(line int, message string) {
	report(line, "", message)
}

func reportParseError(err parser.ParseError) {
	t := err.Token
	if t.TokenType == token.EOFToken {
		report(t.Line, " at end", err.Message)
	} else {
		report(t.Line, " at '"+t.Lexeme+"'", err.Message)
	}
}

func report(line int, where, message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error %s: %s\n", line, where, message)
	hadError = true
}
