package main

import (
	"fmt"
	"io/ioutil"

	// "github.com/nfk93/gocap/utils"

	"github.com/nfk93/gocap/parser/simple/ast"
	"github.com/nfk93/gocap/parser/simple/lexer"
	"github.com/nfk93/gocap/parser/simple/parser"
)

func main() {
	dat, err := ioutil.ReadFile("parser/tests/test1.cgo")
	if err != nil {
		panic(err)
	}

	lex := lexer.NewLexer(dat)
	p := parser.NewParser()
	s, errParse := p.Parse(lex)

	if errParse != nil {
		panic(errParse)
	}
	astree := s.(ast.SourceFile)
	fmt.Println(astree)
}
