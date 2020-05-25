package ast

type ParameterDecl struct {
  id string
  typ Typ
}

type Signature struct {
  params []ParameterDecl
  returnType []Typ
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
  id string
  signature Signature
  body Code
}
func (f FunctionDecl) ToString() string { return "functiondecl" }

func NewFunctionDecl(id_, sign_, body_ Attrib) (Code, error) {
  id := parseId(id_)
  sign := sign_.(Signature)
  body := body_.(Code)
  return FunctionDecl{id, sign, body}, nil
}

type MethodDecl struct {
  receiver Receiver
  id string
  signature Signature
  body Code
}
func (m MethodDecl) ToString() string { return "methoddecl" }

func NewMethodDecl(receiver_, id_, sign_, body_ Attrib) (Code, error) {
  receiver := receiver_.(Receiver)
  id := parseId(id_)
  sign := sign_.(Signature)
  body := body_.(Code)
  return MethodDecl{receiver, id, sign, body}, nil
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
