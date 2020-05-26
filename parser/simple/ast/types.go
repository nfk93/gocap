package ast

import (
	"strings"

	"github.com/nfk93/gocap/utils"
)

type Typ interface {
	ToString() string
}

type IntType struct{}

func (t IntType) ToString() string {
	return "int"
}

type StringType struct{}

func (t StringType) ToString() string {
	return "string"
}

type NamedType struct {
	typeId string
}

func (t NamedType) ToString() string {
	return t.typeId
}

//TODO: not supported yet
type ImportedType struct {
	packageId string
	typeId    string
}

func (t ImportedType) ToString() string {
	return t.packageId + "." + t.typeId
}

type StructType struct {
	fields []StructField
}

func (t StructType) ToString() string {
	result := "struct { \n"
	for _, f := range t.fields {
		result += f.ToString()
	}
	return result + "}\n"
}

type StructField struct {
	id  string
	typ Typ
}

func (t StructField) ToString() string {
	return t.id + " " + t.typ.ToString() + "\n"
}

type PointerType struct {
	typ Typ
}

func (t PointerType) ToString() string {
	return "*" + t.typ.ToString()
}

type FunctionType struct {
	signature Signature
}

func (t FunctionType) ToString() string {
	return "func " + t.signature.ToString()
}

type InterfaceType struct {
	methods []InterfaceMethod
}

func (t InterfaceType) ToString() string {
	result := "interface { \n"
	for _, m := range t.methods {
		result += m.ToString()
	}
	return result + "}\n"
}

type InterfaceMethod struct {
	id        string
	signature Signature
}

func (t InterfaceMethod) ToString() string {
	return t.id + " " + t.signature.ToString()
}

type SliceType struct {
	typ Typ
}

func (t SliceType) ToString() string {
	return "[]" + t.typ.ToString()
}

type MapType struct {
	key  Typ
	elem Typ
}

func (t MapType) ToString() string {
	return "map[" + t.key.ToString() + "]" + t.elem.ToString()
}

type ChannelType struct {
	typ Typ
}

func (t ChannelType) ToString() string {
	return "chan " + t.typ.ToString()
}

type ROChannelType struct {
	typ Typ
}

func (t ROChannelType) ToString() string {
	return "<-chan " + t.typ.ToString()
}

type SOChannelType struct {
	typ Typ
}

func (t SOChannelType) ToString() string {
	return "chan<- " + t.typ.ToString()
}

type CapChannelType struct {
	typ Typ
}

var typeCapChannelTemplate = "capchan.Type_$TYPE"

func typeCapChannel(typeString string) string {
	typeString = utils.RemoveParentheses(typeString)
	result := strings.Replace(typeCapChannelTemplate, "$TYPE", typeString, -1)
	return result
}

func (t CapChannelType) ToString() string {
	return typeCapChannel(t.typ.ToString())
}

//TODO: doesn't support RO SO capchannels
type ROCapChannelType struct {
	typ Typ
}
type SOCapChannelType struct {
	typ Typ
}

func NewNamedType(id_ Attrib) (NamedType, error) {
	id := parseId(id_)
	return NamedType{id}, nil
}

func NewStructType(fields_ Attrib) (StructType, error) {
	fields := fields_.([]StructField)
	return StructType{fields}, nil
}

func MakeStructFields(idlist_, typ_ Attrib) ([]StructField, error) {
	idlist := idlist_.([]string)
	typ := typ_.(Typ)
	fields := make([]StructField, len(idlist))
	for i, id := range idlist {
		fields[i] = StructField{id, typ}
	}
	return fields, nil
}

func NewStructFieldList(field_ Attrib) ([]StructField, error) {
	field := field_.(StructField)
	list := make([]StructField, 1)
	list[0] = field
	return list, nil
}

func AppendStructFields(fielddecls1_, fielddecls2_ Attrib) ([]StructField, error) {
	fielddecls1 := fielddecls1_.([]StructField)
	fielddecls2 := fielddecls2_.([]StructField)
	return append(fielddecls1, fielddecls2...), nil
}

func NewTypeList(typ_ Attrib) ([]Typ, error) {
	typ := typ_.(Typ)
	list := make([]Typ, 1)
	list[0] = typ
	return list, nil
}

func AppendTypeList(typlist_, typ_ Attrib) ([]Typ, error) {
	typlist := typlist_.([]Typ)
	typ := typ_.(Typ)
	return append(typlist, typ), nil
}

func NewPointerType(baseType_ Attrib) (PointerType, error) {
	baseType := baseType_.(Typ)
	return PointerType{baseType}, nil
}

func NewFunctionType(signature_ Attrib) (FunctionType, error) {
	signature := signature_.(Signature)
	return FunctionType{signature}, nil
}

func NewInterfaceType(methods_ Attrib) (InterfaceType, error) {
	methods := methods_.([]InterfaceMethod)
	return InterfaceType{methods}, nil
}

func NewInterfaceMethod(id_, signature_ Attrib) (InterfaceMethod, error) {
	id := parseId(id_)
	signature := signature_.(Signature)
	return InterfaceMethod{id, signature}, nil
}

func NewInterfaceMethodList(method_ Attrib) ([]InterfaceMethod, error) {
	method := method_.(InterfaceMethod)
	list := make([]InterfaceMethod, 1)
	list[0] = method
	return list, nil
}

func AppendInterfaceMethodList(list_, method_ Attrib) ([]InterfaceMethod, error) {
	list := list_.([]InterfaceMethod)
	method := method_.(InterfaceMethod)
	return append(list, method), nil
}

func NewSliceType(typ_ Attrib) (SliceType, error) {
	typ := typ_.(Typ)
	return SliceType{typ}, nil
}

func NewMapType(keytyp_, elemtyp_ Attrib) (MapType, error) {
	key := keytyp_.(Typ)
	elem := elemtyp_.(Typ)
	return MapType{key, elem}, nil
}

func NewChannelType(typ_ Attrib) (ChannelType, error) {
	typ := typ_.(Typ)
	return ChannelType{typ}, nil
}

func NewROChannelType(typ_ Attrib) (ROChannelType, error) {
	typ := typ_.(Typ)
	return ROChannelType{typ}, nil
}

func NewSOChannelType(typ_ Attrib) (SOChannelType, error) {
	typ := typ_.(Typ)
	return SOChannelType{typ}, nil
}

func NewCapChanType(typ_ Attrib) (CapChannelType, error) {
	typ := typ_.(Typ)
	return CapChannelType{typ}, nil
}

func NewROCapChanType(typ_ Attrib) (ROCapChannelType, error) {
	typ := typ_.(Typ)
	return ROCapChannelType{typ}, nil
}

func NewSOCapChanType(typ_ Attrib) (SOCapChannelType, error) {
	typ := typ_.(Typ)
	return SOCapChannelType{typ}, nil
}
