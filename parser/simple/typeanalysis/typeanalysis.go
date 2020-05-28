package typeanalysis

import (
  "errors"
  "fmt"

  im "github.com/benbjohnson/immutable"
	. "github.com/nfk93/gocap/parser/simple/ast"
)

func AnalyzeTypes(s SourceFile) error {
  m := make(map[string]Typ)

  addTypeDecl := func(decl TypeDecl) error {
    if _, exists := m[decl.Id]; exists {
      return errors.New("Type " + decl.Id + " is declared twice")
    } else {
      m[decl.Id] = decl.Typ
      return nil
    }
  }

  addTypeAlias := func(decl TypeAlias) error {
    if _, exists := m[decl.Id]; exists {
      return errors.New("Type " + decl.Id + " is declared twice")
    } else {
      m[decl.Id] = decl.Typ
      return nil
    }
  }

  // Add all type names to the type map
  for _, decl_ := range s.TopLevelDecls {
    switch decl := decl_.(type) {
    case TypeDecl:
      err := addTypeDecl(decl)
      if err != nil { return err }
    case TypeAlias:
      err := addTypeAlias(decl)
      if err != nil { return err }
    case TypeDeclBlock:
      for _, ds := range decl.Decls {
        switch d := ds.(type) {
        case TypeDecl:
          err := addTypeDecl(d)
          if err != nil { return err }
        case TypeAlias:
          err := addTypeAlias(d)
          if err != nil { return err }
        default:
          errmsg := fmt.Sprint("TypeDeclBlock expected to hold TypeAlias or TypeDecl but found: ", d)
          return errors.New(errmsg)
        }
      }
    default:
      continue
    }
  }

  baseMap, err := getBaseTypeMap(m)
  if err != nil {
    return err
  }
  for k, v := range baseMap {
    fmt.Println(k, v.ToString())
  }
  return nil
}

func getBaseTypeMap(typeMap map[string]Typ) (map[string]Typ, error) {
  // m := make(map[string]Typ)

  for id, typ := range typeMap {
    baseType, err := getBaseType(typ, im.NewMap(nil), typeMap)
    if err != nil {
      return nil, err
    }
    typeMap[id] = baseType
  }

  return typeMap, nil
}

func getBaseType(typ_ Typ, visited *im.Map, typeMap map[string]Typ) (Typ, error) {
  switch typ := typ_.(type) {
  case IntType, StringType:
    return typ, nil
  case NamedType:
    if _, alreadyVisited := visited.Get(typ.TypeId) ; alreadyVisited {
      return nil, errors.New("error: found recursive type " + typ.TypeId + ". These are not allowed in cgo")
    }
    nonBaseType, ok := typeMap[typ.TypeId]
    if !ok {
      return nil, errors.New("undefined type: " + typ.TypeId)
    }
    visited = visited.Set(typ.TypeId, true)
    return getBaseType(nonBaseType, visited, typeMap)
  case StructType:
    fields := make([]StructField, len(typ.Fields))
    for i, field := range typ.Fields {
      baseTyp, err := getBaseType(field.Typ, visited, typeMap)
      if err != nil {
        return nil, err
      }
      fields[i] = StructField{field.Id, baseTyp}
    }
    return StructType{fields}, nil
  case PointerType:
    baseTyp, err := getBaseType(typ.Typ, visited, typeMap)
    if err != nil {
      return nil, err
    }
    return PointerType{baseTyp}, nil
  case ChannelType:
    baseTyp, err := getBaseType(typ.Typ, visited, typeMap)
    if err != nil {
      return nil, err
    }
    return ChannelType{baseTyp}, nil
  case FunctionType, InterfaceType, SliceType, MapType, CapChannelType:
    return typ, nil
  default:
    return nil, errors.New(fmt.Sprint("unrecognized type in getBaseType: ", typ))
  }
}

func TypeIsCapability(typ_ Typ, typeMap map[string]Typ) bool {
  switch typ := typ_.(type) {
  case IntType, StringType:
    return false
  case NamedType:
    baseType, _ := typeMap[typ.TypeId]
    return TypeIsCapability(baseType, typeMap)
  case StructType:
    result := false
    for _, field := range typ.Fields {
      result = result || TypeIsCapability(field.Typ, typeMap)
    }
    return result
  case PointerType, FunctionType, ChannelType, InterfaceType, SliceType, MapType:
    return true
  case CapChannelType:
    return true
  default:
    return true
  }
}
