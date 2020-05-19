package test

import (
    "lbs_proj/lexer"
    "lbs_proj/parser"
    "testing"
    "fmt"
)

func TestWorld(t *testing.T) {
    input := []byte(`hello gocc`)
    lex := lexer.NewLexer(input)
    p := parser.NewParser()
    st, err := p.Parse(lex)
    if err != nil {
        panic(err)
    }
    fmt.Println(st)
}
