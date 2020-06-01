package ast

import (
	"errors"
	"fmt"
	"strings"

	"github.com/nfk93/gocap/generator"
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
type ChanMake struct {
	VarId string
	Typ   Typ
}

func (m ChanMake) ToString() string {
	return m.VarId + " := make(chan " + m.Typ.ToString() + ")\n"
}

func NewChanMake(chanId_, typ_ Attrib) (Code, error) {
	chanId := string(chanId_.(*token.Token).Lit)
	typ := typ_.(Typ)
	return ChanMake{chanId, typ}, nil
}

type CapChanMake struct {
	VarId  string
	Typ    Typ
	userId string
}

func (c CapChanMake) ToString() string {
	switch captyp := c.Typ.(type) {
	case PointerType:
		switch captyp2 := captyp.Typ.(type) {
		case ImportedType:
			typename, ok := generator.ExportedTypeMap[captyp2.PackageId]
			if ok {
				return c.VarId + " := " + generator.MakeNewCapChannelTypeInline(captyp2.PackageId, "*"+typename, c.userId)
			}
		}
	case ImportedType:
		typename, ok := generator.ExportedTypeMap[captyp.PackageId]
		if ok {
			return c.VarId + " := " + generator.MakeNewCapChannelTypeInline(captyp.PackageId, typename, c.userId)
		}
	}
	//generator.ImportPackage = append(generator.ImportPackage, utils.TempPkg)
	return c.VarId + " := " + generator.MakeNewCapChannelType(c.Typ.ToString(), c.userId)
}

func NewCapChanMake(chanId_, typ_ Attrib) (Code, error) {
	chanId := string(chanId_.(*token.Token).Lit)
	typ := typ_.(Typ)
	return &CapChanMake{chanId, typ, ""}, nil
}

type CapChanReceive struct {
	receiverId string
	channelId  string
	userId     string
}

func (c CapChanReceive) ToString() string {
	return c.receiverId + " := " + generator.ReceiveCapChannel(c.channelId, c.userId)
}

func NewCapChanReceive(receiverId_, channelId_ Attrib) (Code, error) {
	receiverId := string(receiverId_.(*token.Token).Lit)
	channelId := string(channelId_.(*token.Token).Lit)
	return &CapChanReceive{receiverId, channelId, ""}, nil
}

type CapChanSend struct {
	channelId string
	sendId    string
	userId    string
}

func (c CapChanSend) ToString() string {
	return generator.SendCapChannel(c.channelId, c.sendId, c.userId)
}

func NewCapChanSend(channelId_, sendId_ Attrib) (Code, error) {
	channelId := string(channelId_.(*token.Token).Lit)
	sendId := string(sendId_.(*token.Token).Lit)
	return &CapChanSend{channelId, sendId, ""}, nil
}

type CapChanJoin struct {
	channelId string
	newuserId string
	userId    string
}

func (c *CapChanJoin) ToString() string {
	return generator.JoinCapChannel(c.channelId, c.newuserId, c.userId)
}

func NewCapChanJoin(channelId_, newuserId_ Attrib) (Code, error) {
	channelId := string(channelId_.(*token.Token).Lit)
	newuserId := string(newuserId_.(*token.Token).Lit)
	return &CapChanJoin{channelId, newuserId, ""}, nil
}

// Blocks
type Block struct {
	Code []Code
}

func (b Block) ToString() string {
	s := "{"
	for _, code := range b.Code {
		s += (code.ToString() + " ")
	}
	return s + "}"
}

func NewBlock(codelist_ Attrib) (Code, error) {
	codelist, success := codelist_.([]Code)
	if !success {
		return nil, errors.New("Can't create Block from non-list type")
	} else {
		block := Block{codelist}
		// fmt.Println(block.ToString())
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

func SkipTokens(a ...Attrib) (Code, error) {
	s := string(a[0].(*token.Token).Lit)
	for _, tok := range a[1:] {
		s += " " + string(tok.(*token.Token).Lit)
	}
	return &IgnoredCode{s}, nil
}

type IgnoredIdentifier struct {
	id string
}

func (i IgnoredIdentifier) ToString() string {
	return i.id
}

func SkipId(id_ Attrib) (IgnoredIdentifier, error) {
	id := parseId(id_)
	return IgnoredIdentifier{id}, nil
}

// Source File
type SourceFile struct {
	Packag        string
	imports       []Import
	TopLevelDecls []Code
}

func (s SourceFile) ToString() string {
	ret := "package " + s.Packag + "\n\n"
	//utils.TempPkg = s.Packag

	for _, import_ := range s.imports {
		ret += import_.ToString() + "\n"
	}

	ret += "\n"

	for _, decl := range s.TopLevelDecls {
		ret += decl.ToString() + "\n"
	}

	return ret
}

func NewSourceFile(package_, imports_, topLevelDecls_ Attrib) (SourceFile, error) {
	packag := parseId(package_)
	imports := imports_.([]Import)
	topLevelDecls := topLevelDecls_.([]Code)
	return SourceFile{packag, imports, topLevelDecls}, nil
}

// Imports
type Import struct {
	path  string
	alias string
	dot   bool
}

func (i Import) ToString() string {
	if i.dot {
		return "import " + ". " + "\"" + i.path + "\""
	} else {
		return "import " + i.alias + " " + "\"" + i.path + "\""
	}
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

func ConcatTokens(toks ...Attrib) ([]*token.Token, error) {
	var result []*token.Token

	switch t := toks[0].(type) {
	case []*token.Token:
		result = t
	case *token.Token:
		result = make([]*token.Token, 1)
		result[0] = t
	default:
		s := fmt.Sprint("first argument to ConcatTokens is neither tokenlist or token, but is: ", t)
		return nil, errors.New(s)
	}

	for _, tok := range toks[1:] {
		switch t := tok.(type) {
		case []*token.Token:
			result = append(result, t...)
		case *token.Token:
			result = append(result, t)
		default:
			return nil, errors.New("unrecognized tokens in ConcatTokens")
		}
	}

	return result, nil
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

// func parseTerminator(terminator Attrib) string {
// 	return string(terminator.(*token.Token).Lit)
// }

// func AddNewline(str Attrib) string {
// 	return str.(string) + "\n"
// }
