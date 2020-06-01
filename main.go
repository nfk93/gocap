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
	if len(args) < 1 {
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
	//utils.PackagePath = args[0]
	cwd := getCWD()
	inputPath := args[0]
	directoryPath := path.Join(cwd, inputPath)

	allFiles := getAllfiles(directoryPath)

	generator.ExportedTypeMap = make(map[string]string)
	generator.CapChanTypeMap = make(map[string][]string)
	//generator.ImportPackage = make([]string, 0)

	packages := make(map[string]string)
	datas := make(map[string]string)
	for _, file_ := range allFiles {
		outputPath := file_[:len(file_)-4] + ".go"
		ast_ := handleOneFile(file_)
		datas[outputPath] = ast_.ToString()
		packages[outputPath] = ast_.Packag
	}
	for outputPath, pkg := range packages {
		generator.CreateFileCode(pkg, datas[outputPath], outputPath)
	}

	generator.GenerateCapChannelPackage(directoryPath)
}

func handleOneFile(filePath string) ast.SourceFile {

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
	return astree
}
