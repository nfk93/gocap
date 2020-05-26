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

func TestAllGood(t *testing.T) {
  root := "../tests/success"
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

func TestAllBad( t *testing.T) {
  root := "../tests/bad"
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
      if errParse == nil {
        t.Error(s.(ast.Code).ToString())
      }
    })
  }
}
