package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func main() {
	filename := file_input()
	fset, parsed_res, err := get_parsed(filename)
	if err != nil {
		panic(err)
	}
	for _, c := range parsed_res.Comments {
		pos := c.Pos()
		fmt.Printf("L%d: %s", fset.Position(pos).Line, c.Text())
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
	fmt.Scanln(&filename)

	if !strings.Contains(filename, ".go") {
		panic("file have a .go extension")
	}
	return filename
}
