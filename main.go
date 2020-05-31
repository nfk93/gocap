package main

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/nfk93/gocap/parser/simple/analysis"
	"github.com/nfk93/gocap/parser/simple/ast"
	"github.com/nfk93/gocap/parser/simple/lexer"
	"github.com/nfk93/gocap/parser/simple/parser"
	"github.com/nfk93/gocap/utils"
)

//path to executable
// func getPath() string {
// 	ex, err := os.Executable()
// 	if err != nil {
// 		panic(err)
// 	}
// 	exPath := filepath.Dir(ex)
// 	fmt.Println(exPath)
// 	return exPath
// }

//path to codefile
// func getPath2() string {
// 	_, filename, _, _ := runtime.Caller(1)
// 	exPath := filepath.Dir(filename)
// 	fmt.Println(exPath)
// 	return exPath
// }

func getArgument() []string {
	args := os.Args[1:]
	if len(args) < 2 {
		panic("Need more arguments")
	}
	return args
}

func getCWD() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return path
}

/*
first arugment: package path
second argument: souece code file ending with .cgo
*/

func main() {
	args := getArgument()
	utils.PackagePath = args[0]
	cwd := getCWD()
	fileName := args[1]
	filePath := path.Join(cwd, fileName)

	if path.Ext(fileName) != ".cgo" {
		panic("target file doesn't end with .cgo")
	}
	//baseName := path.Base(fileName)
	//outputName := baseName[:len(baseName)-4] + ".go"
	//outputFilePath := path.Join(cwd, outputName)

	// fmt.Println(filePath)
	// fmt.Println(outputFilePath)

	dat, err := ioutil.ReadFile(filePath)
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
	err = analysis.AnalyzeTypes(astree)
	if err != nil {
		panic(err)
	}

	// generator.CreateFile(astree.ToString(), outputFilePath)
	// generator.GenerateCapChannelPackage(cwd)
}
