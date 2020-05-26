package ast

import (
  "github.com/nfk93/gocap/parser/simple/token"
)

type ConstDeclBlock struct {
	decls []ConstDecl
}

func (c ConstDeclBlock) ToString() string {
	s := ""
	for _, decl := range c.decls {
		s += decl.ToString() + "\n"
	}
	return s
}

func NewConstDeclBlock(constDecls Attrib) (ConstDeclBlock, error) {
	return ConstDeclBlock{constDecls.([]ConstDecl)}, nil
}

type ConstDecl struct {
	id string
	untyped bool
	typ Typ
	uninitialized bool
	constinit []*token.Token
}

func (c ConstDecl) ToString() string {
	if c.uninitialized {
		if c.untyped {
      return "const " + c.id + "\n"
		} else {
      return "const " + c.id + " " + c.typ.ToString() + "\n"
		}
	} else {
		init := string(c.constinit[0].Lit)
		for _, tok := range c.constinit[1:] {
			init += " " + string(tok.Lit)
		}

		if c.untyped {
			return "const " + c.id + " = " + init + "\n"
		} else {
			return "const " + c.id + " " + c.typ.ToString() + " = " + init + "\n"
		}
	}
}

func NewUninitializedConsts(idlist_ Attrib) ([]ConstDecl, error) {
  idlist := idlist_.([]string)
  list := make([]ConstDecl, len(idlist))
  for i, id := range idlist {
    list[i] = ConstDecl{id, true, nil, true, nil}
  }
  return list, nil
}

func NewUntypedConsts(idlist_, init_ Attrib) ([]ConstDecl, error) {
  idlist := idlist_.([]string)
  init := init_.([]*token.Token)
  list := make([]ConstDecl, len(idlist))
  for i, id := range idlist {
    list[i] = ConstDecl{id, true, nil, false, init}
  }
  return list, nil
}

func NewConsts(idlist_, typ_, init_ Attrib) ([]ConstDecl, error) {
  idlist := idlist_.([]string)
  typ := typ_.(Typ)
  init := init_.([]*token.Token)
  list := make([]ConstDecl, len(idlist))
  for i, id := range idlist {
    list[i] = ConstDecl{id, false, typ, false, init}
  }
  return list, nil
}

func ConcatConstDecls(constdecls1_, constdecls2_ Attrib) ([]ConstDecl, error) {
  constdecls1 := constdecls1_.([]ConstDecl)
  constdecls2 := constdecls2_.([]ConstDecl)
  return append(constdecls1, constdecls2...), nil
}
