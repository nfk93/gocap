package test

import (
    "github.com/nfk93/gocap/parser/simple/lexer"
    "github.com/nfk93/gocap/parser/simple/parser"
    "io/ioutil"
    "testing"
    "fmt"
)

func TestUnsupportBufferedChannel(t *testing.T) {
  dat, err := ioutil.ReadFile("../tests/unsupport_buffered_chan.cgo")
  if err != nil{
    panic(err)
  }

  lex := lexer.NewLexer(dat)
  p := parser.NewParser()
  _, errParse := p.Parse(lex)
  if errParse == nil {
      t.Error("should have failed due to buffered channel")
  }
}


func TestWorld(t *testing.T) {
    dat, err := ioutil.ReadFile("../tests/test1.go")
    if err != nil{
      panic(err)
    }

    lex := lexer.NewLexer(dat)
    p := parser.NewParser()
    s, errParse := p.Parse(lex)
    if errParse != nil {
        panic(errParse)
    }

    fmt.Println(s)
}
