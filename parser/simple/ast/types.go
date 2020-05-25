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
