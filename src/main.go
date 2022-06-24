package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

func main() {
	filename := file_input()
	fset, parsed_res, err := get_parsed(filename)
	if err != nil {
		panic(err)
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	file_by_line := strings.Split(string(file), "\n")

	for _, c := range parsed_res.Comments {
		if strings.TrimSpace(c.Text()) != "add values a and b" {
			continue
		}

		line_number := fset.Position(c.Pos()).Line
		file_by_line = append(
			file_by_line[:line_number+1],
			file_by_line[line_number:]...,
		)
		file_by_line[line_number] = "\tfmt.Println(\"yay!\")"
		file_string := strings.Join(file_by_line, "\n")
		os.WriteFile(filename, []byte(file_string), 0644)
	}
}

func get_parsed(filename string) (*token.FileSet, *ast.File, error) {
	fset := token.NewFileSet()
	tokens, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	return fset, tokens, err
}

func file_input() string {
	fmt.Print("filename :: ")
	var filename string
	fmt.Scan(&filename)

	if !strings.Contains(filename, ".go") {
		panic("file have a .go extension")
	}
	return filename
}
