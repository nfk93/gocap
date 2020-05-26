package main

import (
	"fmt"
	"io/ioutil"

	"github.com/nfk93/gocap/generator"
	"github.com/nfk93/gocap/utils"

	"github.com/nfk93/gocap/parser/simple/ast"
	"github.com/nfk93/gocap/parser/simple/lexer"
	"github.com/nfk93/gocap/parser/simple/parser"
)

func main() {
	utils.PackagePath = "github.com/nfk93/gocap/output"

	testfile := "test3.cgo"
	outputfile := "test3.go"
	outputPath := "../../output"
	dat, err := ioutil.ReadFile("../tests/" + testfile)
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
	fmt.Println(astree.ToString())
	generator.CreateFile(astree.ToString(), outputPath+"/"+outputfile)
	utils.IfPrintPackages = false
	generator.GenerateCapChannelPackage("../../generator", outputPath)
}
