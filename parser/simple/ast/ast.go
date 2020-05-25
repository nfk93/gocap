package ast

import (
	"errors"
	"fmt"
  "strings"
	"github.com/nfk93/gocap/parser/simple/token"
)

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
	VarId string
	Typ   string
}

func (m MakeChanExpr) ToString() string { return m.VarId + m.Typ }

func NewChanMake(chanId, typ Attrib) (Code, error) {
	return MakeChanExpr{"", ""}, nil
}

type CapChanMake struct {
	VarId string
	Typ   string
}

func (c CapChanMake) ToString() string { return c.VarId + c.Typ }

func NewCapChanMake(a1, a2 Attrib) (Code, error) {
	//TODO: need to convert typ to string
	return CapChanMake{"", ""}, nil
}

type CapChanReceive struct {
	ReceiverId string
	ChannelId  string
}

func (c CapChanReceive) ToString() string { return fmt.Sprintf("%+v\n", c) }

func NewCapChanReceive(receiverId_, channelId_ Attrib) (Code, error) {
	receiverId := string(receiverId_.(*token.Token).Lit)
	channelId := string(channelId_.(*token.Token).Lit)
	return CapChanReceive{receiverId, channelId}, nil
}

type CapChanSend struct {
	ChannelId string
	SendId    string
}

func (c CapChanSend) ToString() string { return fmt.Sprintf("%+v\n", c) }

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
	for _, code := range b.code {
		s += code.ToString()
	}
	return s
}
func NewBlock(codelist_ Attrib) (Code, error) {
	codelist, success := codelist_.([]Code)
	if !success {
		return nil, errors.New("Can't create Block from non-list type")
	} else {
		block := Block{codelist}
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

// Source File
type SourceFile struct {
  packag string
  imports []Import
  topLevelDecls []Code
}

func NewSourceFile(package_, imports_, topLevelDecls_ Attrib) (SourceFile, error) {
  packag := parseId(package_)
  imports := imports_.([]Import)
  topLevelDecls := topLevelDecls_.([]Code)
  return SourceFile{packag, imports, topLevelDecls}, nil
}

// Imports
type Import struct{
  path string
  alias string
  dot bool
}

func NewImport(path Attrib, dot bool) (Import, error) {
  p := parseString(path)
  alias := packageId(p)
  return Import{p, alias, dot}, nil
}

func NewNamedImport(path_, alias_ Attrib) (Import, error) {
  path := parseString(path_)
  alias := parseId(alias_)
  return Import{path, alias, false}, nil
}

func AppendImportLists(list1_, list2_ Attrib) ([]Import, error) {
  list1 := list1_.([]Import)
  switch list2 := list2_.(type) {
  case []Import:
    return append(list1, list2...), nil
  case Import:
    return append(list1, list2), nil
  default:
    return nil, errors.New("Unrecognized import, can't append import lists")
  }
}

// Unsupported, throws error
func Unsupported(err string) (interface{}, error) {
	return nil, errors.New(err)
}

// Utility functions
func parseId(id Attrib) string {
	return string(id.(*token.Token).Lit)
}

func parseString(str Attrib) string {
	s := string(str.(*token.Token).Lit)[1:]
	s = s[:(len(s) - 1)]
	return s
}

func packageId(path string) string {
  id := path[strings.LastIndex(path, "/")+1:]
  return id
}
