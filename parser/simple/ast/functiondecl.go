package ast

import "strings"

type ParameterDecl struct {
	id  string
	typ Typ
}

func (p *ParameterDecl) ToString() string {
	return p.id + " " + p.typ.ToString()
}

type Signature struct {
	params     []ParameterDecl
	returnType []Typ
}

func (s *Signature) ToString() string {
	var paraStringArray []string
	for _, p := range s.params {
		paraStringArray = append(paraStringArray, p.ToString())
	}
	paraStringReturn := "(" + strings.Join(paraStringArray, ",") + ") "
	if len(s.returnType) == 0 {
		return paraStringReturn
	} else if len(s.returnType) == 1 {
		return paraStringReturn + s.returnType[0].ToString() + " "
	}
	var returnStringArray []string
	for _, r := range s.returnType {
		returnStringArray = append(returnStringArray, r.ToString())
	}
	return paraStringReturn + "(" + strings.Join(returnStringArray, ",") + ")"
}

type Receiver struct {
  Id string
  Typ Typ
}

func NewReceiver(id_, typName_ Attrib) (Receiver, error) {
  id := parseId(id_)
  typname := parseId(typName_)
  return Receiver{id, NamedType{typname}}, nil
}

func NewPointerReceiver(id_, typName_ Attrib) (Receiver, error) {
  id := parseId(id_)
  typname := parseId(typName_)
  return Receiver{id, PointerType{NamedType{typname}}}, nil
}

type FunctionDecl struct {
	id        string
	signature Signature
	body      Code
	//TODO: should be a Block instead of Code
}

func (f *FunctionDecl) ToString() string {
	return "func " + f.id + f.signature.ToString() + "{\n" + f.body.ToString() + "}\n"
}

func NewFunctionDecl(id_, sign_, body_ Attrib) (Code, error) {
	id := parseId(id_)
	sign := sign_.(Signature)
	body := body_.(Code)
	return &FunctionDecl{id, sign, body}, nil
}

type MethodDecl struct {
	receiver  Receiver
	id        string
	signature Signature
	body      Code
	//TODO: should be a Block instead of Code
}

func (m *MethodDecl) ToString() string {
	//return "func (" + m.receiver.ToString() + ") " + f.id + f.signature.ToString() + "{\n" + f.body.ToString() + "}\n"
	//TODO
	return "TODO"
}

func NewMethodDecl(receiver_, id_, sign_, body_ Attrib) (Code, error) {
	receiver := receiver_.(Receiver)
	id := parseId(id_)
	sign := sign_.(Signature)
	body := body_.(Code)
	return &MethodDecl{receiver, id, sign, body}, nil
}

func NewIdList(id_ Attrib) ([]string, error) {
	id := parseId(id_)
	list := make([]string, 1)
	list[0] = id
	return list, nil
}

func AppendIdList(idList_, id_ Attrib) ([]string, error) {
	idList := idList_.([]string)
	id := parseId(id_)
	return append(idList, id), nil
}

func MakeParameterDecl(idlist_, typ_ Attrib) ([]ParameterDecl, error) {
	idlist := idlist_.([]string)
	typ := typ_.(Typ)
	decls := make([]ParameterDecl, len(idlist))
	for i, id := range idlist {
		decls[i] = ParameterDecl{id, typ}
	}
	return decls, nil
}

func ConcatParameterLists(paramlist1_, paramlist2_ Attrib) ([]ParameterDecl, error) {
	paramlist1 := paramlist1_.([]ParameterDecl)
	paramlist2 := paramlist2_.([]ParameterDecl)
	return append(paramlist1, paramlist2...), nil
}

func NewSignature(params_, result_ Attrib) (Signature, error) {
	params := params_.([]ParameterDecl)
	result := result_.([]Typ)
	return Signature{params, result}, nil
}