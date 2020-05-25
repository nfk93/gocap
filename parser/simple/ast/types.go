package ast

type Typ interface{}

type IntType struct{}
type StringType struct{}
type NamedType struct{
  typeId string
}
type ImportedType struct{
  packageId string
  typeId string
}
type StructType struct{
  fields []StructField
}
type StructField struct{
  id string
  typ Typ
}
type PointerType struct {
  typ Typ
}
type FunctionType struct{
  signature Signature
}
type InterfaceType struct {
  methods []InterfaceMethod
}
type InterfaceMethod struct {
  id string
  signature Signature
}
type SliceType struct {
  typ Typ
}
type MapType struct {
  key Typ
  elem Typ
}
type ChannelType struct {
  typ Typ
}
type ROChannelType struct {
  typ Typ
}
type SOChannelType struct {
  typ Typ
}
type CapChannelType struct {
  typ Typ
}
type ROCapChannelType struct {
  typ Typ
}
type SOCapChannelType struct {
  typ Typ
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
