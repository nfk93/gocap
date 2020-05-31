package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/nfk93/gocap/generator"
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

func getAllfiles(rootPath string) []string {
	var files []string

	err := filepath.Walk(rootPath, func(p string, info os.FileInfo, err error) error {
		if path.Ext(p) == ".cgo" {
			files = append(files, p)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}
	return files
}

/*
first arugment: package path
second argument: souece code file ending with .cgo
*/

func main() {
	args := getArgument()
	utils.PackagePath = args[0]
	cwd := getCWD()
	inputPath := args[1]
	directoryPath := path.Join(cwd, inputPath)

	allFiles := getAllfiles(directoryPath)

	for _, file_ := range allFiles {
		handleOneFile(file_)
	}

	generator.GenerateCapChannelPackage(directoryPath)
}

func handleOneFile(filePath string) {

	outputPath := filePath[:len(filePath)-4] + ".go"

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

	generator.CreateFile(astree.ToString(), outputPath)
}
