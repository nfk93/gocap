package typeanalysis

import (
	// "fmt"
	"io/ioutil"
	"testing"
  // "os"
  // "strings"

  "github.com/nfk93/gocap/parser/simple/ast"
	"github.com/nfk93/gocap/parser/simple/lexer"
	"github.com/nfk93/gocap/parser/simple/parser"
)

func TestAllGood(t *testing.T) {
  root := "tests/success"
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

      typeErr := AnalyzeTypes(s.(ast.SourceFile))

    	if typeErr != nil {
    		t.Error(typeErr)
    	}
    })
  }
}

func TestAllBad( t *testing.T) {
  root := "tests/bad"
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
        t.Error(errParse)
      }

      typeErr := AnalyzeTypes(s.(ast.SourceFile))

    	if typeErr == nil {
    		t.Error("should have found error in typeanalysis but didn't")
    	}
    })
  }
}
