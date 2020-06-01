package analysis

import (
	"errors"
	"fmt"
	"reflect"
	"unicode"

	im "github.com/benbjohnson/immutable"
	"github.com/nfk93/gocap/generator"
	. "github.com/nfk93/gocap/parser/simple/ast"
)

func AnalyzeTypes(s SourceFile) error {
	m := make(map[string]Typ)

	// name of the exported type:
	var exportedTypeName string = ""

	addTypeDecl := func(decl TypeDecl) error {
		if _, exists := m[decl.Id]; exists {
			return errors.New("Type " + decl.Id + " is declared twice")
		}
		if isExported(decl.Id) {
			if exportedTypeName != "" {
				return errors.New("Found two exported types. Only one exported type is allowed:\n" +
					"\t" + exportedTypeName + "\n" +
					"\t" + decl.Id)
			}
			exportedTypeName = decl.Id
		}
		m[decl.Id] = decl.Typ
		return nil
	}

	addTypeAlias := func(decl TypeAlias) error {
		if _, exists := m[decl.Id]; exists {
			return errors.New("Type " + decl.Id + " is declared twice")
		} else {
			if isExported(decl.Id) {
				return errors.New("type aliases can't be exported, but found type alias " + decl.Id)
			}
			m[decl.Id] = decl.Typ
			return nil
		}
	}

	// Add all type names to the type map
	for _, decl_ := range s.TopLevelDecls {
		switch decl := decl_.(type) {
		case TypeDecl:
			err := addTypeDecl(decl)
			if err != nil {
				return err
			}
		case TypeAlias:
			err := addTypeAlias(decl)
			if err != nil {
				return err
			}
		case TypeDeclBlock:
			for _, ds := range decl.Decls {
				switch d := ds.(type) {
				case TypeDecl:
					err := addTypeDecl(d)
					if err != nil {
						return err
					}
				case TypeAlias:
					err := addTypeAlias(d)
					if err != nil {
						return err
					}
				default:
					errmsg := fmt.Sprint("TypeDeclBlock expected to hold TypeAlias or TypeDecl but found: ", d)
					return errors.New(errmsg)
				}
			}
		default:
			continue
		}
	}

	// Calculate the base type of all named types
	baseTypeMap, err := getBaseTypeMap(m)
	if err != nil {
		return err
	}

	// Check that if there is an exported type that it is a struct with all
	// fields unexported.
	if exportedTypeName != "" {
		generator.ExportedTypeMap[s.Packag] = exportedTypeName
		exportedTyp_, _ := baseTypeMap[exportedTypeName]
		switch exportedTyp := exportedTyp_.(type) {
		case StructType:
			for _, field := range exportedTyp.Fields {
				if isExported(field.Id) {
					return errors.New("struct fields in exported struct type is ")
				}
			}
		default:
			return errors.New("Exported type must be a struct declared in this package")
		}
	}

	err = checkFunctionAndMethodDecls(s.TopLevelDecls, baseTypeMap, exportedTypeName)
	if err != nil {
		return err
	}

	for k, v := range baseTypeMap {
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
		if _, alreadyVisited := visited.Get(typ.TypeId); alreadyVisited {
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
	case FunctionType, InterfaceType, SliceType, MapType, CapChannelType, ImportedType:
		// We don't need to get basetypes of these functions
		return typ, nil
	default:
		return nil, errors.New(fmt.Sprint("unrecognized type in getBaseType: ", typ))
	}
}

//eliminate NamedType
func getBaseTypeForNamedType(typ_ Typ, typeMap map[string]Typ) Typ {
	switch typ := typ_.(type) {
	case NamedType:
		baseType, _ := typeMap[typ.TypeId]
		return baseType
	default:
		return typ
	}
}

func isCapability(typ_ Typ, typeMap map[string]Typ) bool {
	switch getBaseTypeForNamedType(typ_, typeMap).(type) {
	case IntType, StringType:
		return false
	case StructType:
		// result := false
		// for _, field := range typ.Fields {
		//   result = result || TypeIsCapability(field.Typ, typeMap)
		// }
		// return result
		return true
	case PointerType, FunctionType, ChannelType, InterfaceType, SliceType, MapType:
		return true
	case CapChannelType:
		return true
	default:
		return true
	}
}

func checkFunctionAndMethodDecls(decls []Code, typeMap map[string]Typ, exportedTypeName string) error {
	// Add all type names to the type map

	//collect all function identifiers
	var allFunctionIds []string
	for _, decl_ := range decls {
		switch decl := decl_.(type) {
		case FunctionDecl:
			allFunctionIds = append(allFunctionIds, decl.Id)
		default:
			// donothing
		}
	}

	for _, decl_ := range decls {
		switch decl := decl_.(type) {
		case FunctionDecl:
			exported := isExported(decl.Id)
			var err error
			if exported {
				if exportedTypeName != "" && decl.Id == "New"+exportedTypeName {
					err = checkFunctionReturn(decl.Signature.ReturnType, typeMap, exportedTypeName)
					if err != nil {
						return err
					}
				} else {
					return errors.New("unexpected exported function name:" + decl.Id)
				}
			}
			err = findAndCheckChannelMakes(decl.Body.Code, typeMap)
			if err != nil {
				return err
			}

		case MethodDecl:
			err := checkMethodSig(decl.Id, decl.Signature, typeMap)
			if err != nil {
				return err
			}
			err = checkMethodBody(decl.Id, decl.Body, typeMap, allFunctionIds)
			if err != nil {
				return err
			}

		default:
			// donothing
		}
	}
	return nil
}

//we assume Imported types are struct types
func isStructType(typ Typ, typeMap map[string]Typ) bool {
	switch getBaseTypeForNamedType(typ, typeMap).(type) {
	case StructType:
		return true
	case ImportedType:
		return true
	default:
		return false
	}
}

func checkMethodSig(id string, sig Signature, typeMap map[string]Typ) error {
	var err error
	//parameter type can be
	// (a) Non-capability types
	// (b) Pointers to non-capability types
	// (c) Pointers to structs
	// (d) Slices of non-capability types
	// (e) Channels of non-capability types
	// (f) Capability Channels
	for _, parameterDecl := range sig.Params {
		paraType_ := parameterDecl.Typ
		if !isCapability(paraType_, typeMap) { //Non-capability types
			continue
		} else {
			err = errors.New(fmt.Sprint("method ", id, " has invalid paramater type: ", paraType_.ToString()))
			switch paraType := getBaseTypeForNamedType(paraType_, typeMap).(type) {
			case PointerType:
				if !isCapability(paraType.Typ, typeMap) { //Pointers to non-capability types
					continue
				} else if isStructType(paraType.Typ, typeMap) { //Pointers to structs
					continue
				} else {
					return err
				}
			case SliceType: //Slices of non-capability types
				if isCapability(paraType.Typ, typeMap) {
					return err
				}
			case ChannelType: //Channels of non-capability types
				if isCapability(paraType.Typ, typeMap) {
					return err
				}
			case MapType:
				if isCapability(paraType.Elem, typeMap) {
					return err
				}
				continue
			case CapChannelType:
				continue //Capability Channels
			default:
				return err
			}
		}
	}

	// return type of method must be non-capability type
	for _, typ := range sig.ReturnType {
		if isCapability(typ, typeMap) {
			return errors.New(fmt.Sprint("method ", id, " has invalid return type. Return type can't be a capability type"))
		}
	}
	return nil
}

func checkMethodBody(id string, body Block, typeMap map[string]Typ, allFunctionIds []string) error {
	err := findAndCheckChannelMakes(body.Code, typeMap)
	if err != nil {
		return err
	}

	//check if no function call in method body
	for _, code := range body.Code {
		for _, funcId := range allFunctionIds {
			if code.ToString() == funcId {
				return errors.New(fmt.Sprint("method ", id, " has potiential function call to ", funcId))
			}
		}
	}
	return nil
}

func isExported(id string) bool {
	return unicode.IsUpper(rune(id[0]))
}

// return type of exported function can only be non-capability type or Pointers to structs
func checkFunctionReturn(returnTypes []Typ, typeMap map[string]Typ, exportedTypeName string) error {
	for _, returnType_ := range returnTypes {
		if !isCapability(returnType_, typeMap) { //Non-Capability type
			continue
		} else {
			switch returnType := getBaseTypeForNamedType(returnType_, typeMap).(type) {
			case PointerType:
				baseType := getBaseTypeForNamedType(returnType.Typ, typeMap)
				if reflect.DeepEqual(baseType, typeMap[exportedTypeName]) { //Pointers to exported structs
					continue
				} else {
					return errors.New("exported function cannot return " + returnType_.ToString())
				}
			default:
				return errors.New("exported function cannot return " + returnType_.ToString())
			}
		}
	}
	return nil
}

func findAndCheckChannelMakes(blockelements []Code, typeMap map[string]Typ) error {
	for _, code_ := range blockelements {
		switch code := code_.(type) {
		case ChanMake:
			if isCapability(code.Typ, typeMap) {
				return errors.New("can't make channel with type: " + code.Typ.ToString())
			}
		}
	}
	return nil
}
