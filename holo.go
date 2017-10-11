package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
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
	scanner := NewScanner(source)
	tokens := scanner.scanTokens()

	for _, token := range tokens {
		fmt.Println(token)
	}
}

func reportError(line int, message string) {
	report(line, "", message)
}

func report(line int, where, message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error %s: %s\n", line, where, message)
	hadError = true
}
