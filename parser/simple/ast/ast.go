package ast

import "fmt"
import "github.com/nfk93/gocap/parser/simple/token"
import "errors"

type Attrib interface{}

type Code interface {
  ToString() string
}

type IgnoredCode struct {
  code string
}
func (i IgnoredCode) ToString() string { return i.code }

// Channel Expressions
type MakeChanExpr struct {
  varId string
  typ string
}
func (m MakeChanExpr) ToString() string { return m.varId + m.typ }

func NewChanMake(chanId, typ Attrib) (Code, error) {
  return MakeChanExpr{"", ""}, nil
}

type CapChanMake struct {
  varId string
  typ string
}
func (c CapChanMake) ToString() string { return c.varId + c.typ }

func NewCapChanMake(a1, a2 Attrib) (Code, error) {
  return CapChanMake{"", ""}, nil
}

type CapChanReceive struct {
  receiverId string
  channelId string
}
func (c CapChanReceive) ToString() string { return fmt.Sprintf("%+v\n", c)}

func NewCapChanReceive(receiverId_, channelId_ Attrib) (Code, error) {
  receiverId := string(receiverId_.(*token.Token).Lit)
  channelId := string(channelId_.(*token.Token).Lit)
  return CapChanReceive{receiverId, channelId}, nil
}

type CapChanSend struct {
  channelId string
  sendId string
}
func (c CapChanSend) ToString() string { return fmt.Sprintf("%+v\n", c)}

func NewCapChanSend(channelId_, sendId_ Attrib) (Code, error) {
  channelId := string(channelId_.(*token.Token).Lit)
  sendId := string(sendId_.(*token.Token).Lit)
  return CapChanSend{channelId, sendId}, nil
}

// Blocks
type Block struct {
  code []Code
}
func (b Block) ToString() string {
  s := ""
  for _, code := range (b.code) {
    s += code.ToString()
  }
  return s
}
func NewBlock(codelist_ Attrib) (Code, error) {
  codelist, success := codelist_.([]Code)
  if !success {
    return nil, errors.New("Can't create Block from non-list type")
  } else {
    block :=  Block{codelist}
    fmt.Println(block.ToString())
    return block, nil
  }
}
func NewBlockContentList(a Attrib) ([]Code, error) {
  l := make([]Code, 1)
  l[0] = a.(Code)
  return l, nil
}

// Cast arguments as []Code and Code and appends the second argument
// to the first
func AppendCodeList(list, a Attrib) ([]Code, error) {
  codelist := list.([]Code)
  code := a.(Code)
  return append(codelist, code), nil
}

func SkipToken(a Attrib) (Code, error) {
  fmt.Println(string(a.(*token.Token).Lit))
  return &IgnoredCode{}, nil
}

func Unsupported(err string) (interface{}, error) {
  return nil, errors.New(err)
}

func PrintDebug(s Attrib) (string, error) {
  fmt.Println(s.(string))
  return "", nil
}

func parseId(id Attrib) string {
  return string(id.(*token.Token).Lit)
}
