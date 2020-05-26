package test

import (
	"fmt"
	"io/ioutil"
	"testing"
  // "os"
  // "strings"

  "github.com/nfk93/gocap/parser/simple/ast"
	"github.com/nfk93/gocap/parser/simple/lexer"
	"github.com/nfk93/gocap/parser/simple/parser"
)

func TestAll(t *testing.T) {
  // var files []string
  // gopath := strings.Split(os.Getenv("GOPATH"), string(os.PathListSeparator))
  // root := gopath + "/src/github.com/nfk93/gocap/parser/tests"
  root := "../tests"
  items, _ := ioutil.ReadDir(root)
  for _, file := range items {
    t.Run(file.Name(), func(t *testing.T) {
      dat, err := ioutil.ReadFile(root + "/" + file.Name())
      if err != nil {
        panic(err)
      }

      lex := lexer.NewLexer(dat)
      p := parser.NewParser()
      s, errParse := p.Parse(lex)
      if errParse != nil {
        panic(errParse)
      }
      fmt.Println(s.(ast.Code).ToString())
    })
  }
}

// func TestUnsupportBufferedChannel(t *testing.T) {
// 	dat, err := ioutil.ReadFile("../tests/unsupport_buffered_chan.cgo")
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	lex := lexer.NewLexer(dat)
// 	p := parser.NewParser()
// 	_, errParse := p.Parse(lex)
// 	if errParse == nil {
// 		t.Error("should have failed due to buffered channel")
// 	}
// }
//
// func TestWorld(t *testing.T) {
// 	dat, err := ioutil.ReadFile("../tests/test1.go")
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	lex := lexer.NewLexer(dat)
// 	p := parser.NewParser()
// 	s, errParse := p.Parse(lex)
// 	fmt.Println(s)
// 	fmt.Println(123)
// 	if errParse != nil {
// 		panic(errParse)
// 	}
//
// }
