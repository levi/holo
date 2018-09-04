package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: generate_ast <output_directory>")
		return
	}

	outputDir := os.Args[1]
	err := defineAst(outputDir, "Expr", []string{
		"Binary		: left Expr, operation token.Token, right Expr",
		"Grouping	: expression Expr",
		"Literal	: value interface{}",
		"Unary		: operation token.Token, right Expr",
	})

	if err != nil {
		panic(err)
	}
}

func defineAst(outputDir, baseName string, types []string) error {
	snakeName := strcase.ToSnake(baseName)
	filename := fmt.Sprintf("%s/%s.go", outputDir, snakeName)
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString("package holo\n")
	f.WriteString("\n")

	f.WriteString("import (\n")
	f.WriteString("    \"github.com/levi/holo/token\"\n")
	f.WriteString(")\n")
	f.WriteString("\n")

	f.WriteString(fmt.Sprintf("type %s interface {}\n", baseName))
	f.WriteString("\n")

	for _, langType := range types {
		production := strings.Split(langType, ":")
		className := strings.TrimSpace(production[0])
		fields := strings.TrimSpace(production[1])
		defineType(f, baseName, className, fields)
	}

	f.Sync()

	return nil
}

func defineType(f *os.File, baseName, className, fields string) {
	f.WriteString(fmt.Sprintf("type %s struct {\n", className))
	fieldList := strings.Split(fields, ", ")
	for _, field := range fieldList {
		name, argType := fieldParts(field, baseName)
		f.WriteString(fmt.Sprintf("    %s %s\n", name, argType))
	}
	f.WriteString("}\n")
	f.WriteString("\n")

	f.WriteString(fmt.Sprintf("func New%s(", className))

	arguments := []string{}
	for _, field := range fieldList {
		name, argType := fieldParts(field, baseName)
		name = strcase.ToLowerCamel(name)
		arguments = append(arguments, fmt.Sprintf("%s %s", name, argType))
	}
	f.WriteString(strings.Join(arguments, ", "))
	f.WriteString(fmt.Sprintf(") *%s {\n", className))

	f.WriteString(fmt.Sprintf("    n = new(%s)\n", className))
	for _, field := range fieldList {
		name, _ := fieldParts(field, baseName)
		f.WriteString(fmt.Sprintf("    n.%s = %s\n", name, strcase.ToLowerCamel(name)))
	}
	f.WriteString("    return n\n")

	f.WriteString("}\n")
	f.WriteString("\n")
}

func fieldParts(field, baseName string) (string, string) {
	parts := strings.Split(field, " ")

	argType := parts[1]
	if argType == baseName {
		argType = fmt.Sprintf("%s{}", baseName)
	}

	return parts[0], argType
}
