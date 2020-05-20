package test

import (
    "github.com/nfk93/gocap/lexer"
    "github.com/nfk93/gocap/parser"
    "io/ioutil"
    "testing"
    "fmt"
)

func TestWorld(t *testing.T) {
    dat, err := ioutil.ReadFile("test/test1.go")
    if err != nil{
      panic(err)
    }

    lex := lexer.NewLexer(dat)
    p := parser.NewParser()
    _, errParse := p.Parse(lex)
    if errParse != nil {
        panic(errParse)
    }

    fmt.Println("ok")
}
