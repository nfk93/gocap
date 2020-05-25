// Code generated by gocc; DO NOT EDIT.

package parser

import "github.com/nfk93/gocap/parser/simple/ast"

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : SourceFile	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SourceFile : PackageClause ImportDecls TopLevelDecls	<< ast.NewSourceFile(X[0], X[1], X[2]) >>`,
		Id:         "SourceFile",
		NTType:     1,
		Index:      1,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSourceFile(X[0], X[1], X[2])
		},
	},
	ProdTabEntry{
		String: `Terminators : terminator Terminators	<< nil, nil >>`,
		Id:         "Terminators",
		NTType:     2,
		Index:      2,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `Terminators : empty	<< nil, nil >>`,
		Id:         "Terminators",
		NTType:     2,
		Index:      3,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `PackageClause : kw_package id	<< X[1], nil >>`,
		Id:         "PackageClause",
		NTType:     3,
		Index:      4,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `ImportDecls : ImportDecls ImportDecl	<< ast.AppendImportLists(X[0], X[1]) >>`,
		Id:         "ImportDecls",
		NTType:     4,
		Index:      5,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendImportLists(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `ImportDecls : empty	<< make([]ast.Import, 0), nil >>`,
		Id:         "ImportDecls",
		NTType:     4,
		Index:      6,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return make([]ast.Import, 0), nil
		},
	},
	ProdTabEntry{
		String: `ImportDecl : kw_import ImportSpec	<< X[1], nil >>`,
		Id:         "ImportDecl",
		NTType:     5,
		Index:      7,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `ImportDecl : kw_import lparen ImportSpecs rparen	<< X[2], nil >>`,
		Id:         "ImportDecl",
		NTType:     5,
		Index:      8,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[2], nil
		},
	},
	ProdTabEntry{
		String: `ImportSpecs : ImportSpecs ImportSpec	<< ast.AppendImportLists(X[0], X[1]) >>`,
		Id:         "ImportSpecs",
		NTType:     6,
		Index:      9,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendImportLists(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `ImportSpecs : empty	<< make([]ast.Import, 0), nil >>`,
		Id:         "ImportSpecs",
		NTType:     6,
		Index:      10,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return make([]ast.Import, 0), nil
		},
	},
	ProdTabEntry{
		String: `ImportSpec : dot ImportPath	<< ast.NewImport(X[1], true) >>`,
		Id:         "ImportSpec",
		NTType:     7,
		Index:      11,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewImport(X[1], true)
		},
	},
	ProdTabEntry{
		String: `ImportSpec : id ImportPath	<< ast.NewNamedImport(X[0], X[1]) >>`,
		Id:         "ImportSpec",
		NTType:     7,
		Index:      12,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNamedImport(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `ImportSpec : ImportPath	<< ast.NewImport(X[0], false) >>`,
		Id:         "ImportSpec",
		NTType:     7,
		Index:      13,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewImport(X[0], false)
		},
	},
	ProdTabEntry{
		String: `ImportPath : string_lit	<< X[0], nil >>`,
		Id:         "ImportPath",
		NTType:     8,
		Index:      14,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `TopLevelDecls : TopLevelDecls TopLevelDecl	<< ast.AppendCodeList(X[0], X[1]) >>`,
		Id:         "TopLevelDecls",
		NTType:     9,
		Index:      15,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendCodeList(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `TopLevelDecls : empty	<< (make([]ast.Code, 0)), nil >>`,
		Id:         "TopLevelDecls",
		NTType:     9,
		Index:      16,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return (make([]ast.Code, 0)), nil
		},
	},
	ProdTabEntry{
		String: `TopLevelDecl : Declaration	<< X[0], nil >>`,
		Id:         "TopLevelDecl",
		NTType:     10,
		Index:      17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `TopLevelDecl : FunctionDecl	<< X[0], nil >>`,
		Id:         "TopLevelDecl",
		NTType:     10,
		Index:      18,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `TopLevelDecl : MethodDecl	<< X[0], nil >>`,
		Id:         "TopLevelDecl",
		NTType:     10,
		Index:      19,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `FunctionDecl : kw_func id Signature FunctionBody	<< ast.NewFunctionDecl(X[1], X[2], X[3]) >>`,
		Id:         "FunctionDecl",
		NTType:     11,
		Index:      20,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFunctionDecl(X[1], X[2], X[3])
		},
	},
	ProdTabEntry{
		String: `MethodDecl : kw_func Receiver id Signature FunctionBody	<< ast.NewMethodDecl(X[1], X[2], X[3], X[4]) >>`,
		Id:         "MethodDecl",
		NTType:     12,
		Index:      21,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewMethodDecl(X[1], X[2], X[3], X[4])
		},
	},
	ProdTabEntry{
		String: `FunctionBody : Block	<< X[0], nil >>`,
		Id:         "FunctionBody",
		NTType:     13,
		Index:      22,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `FunctionBody : empty	<< ast.Unsupported("UNSUPPORTED: function declarations must be follow by a body") >>`,
		Id:         "FunctionBody",
		NTType:     13,
		Index:      23,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Unsupported("UNSUPPORTED: function declarations must be follow by a body")
		},
	},
	ProdTabEntry{
		String: `Signature : Parameters Result	<< ast.NewSignature(X[0], X[1]) >>`,
		Id:         "Signature",
		NTType:     14,
		Index:      24,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSignature(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Result : Type	<< X[0], nil >>`,
		Id:         "Result",
		NTType:     15,
		Index:      25,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Result : lparen TypeList rparen	<< X[1], nil >>`,
		Id:         "Result",
		NTType:     15,
		Index:      26,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Result : Parameters	<< ast.Unsupported("UNSUPPORTED: identifiers in result type not supported") >>`,
		Id:         "Result",
		NTType:     15,
		Index:      27,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Unsupported("UNSUPPORTED: identifiers in result type not supported")
		},
	},
	ProdTabEntry{
		String: `Result : empty	<< (make([]ast.Typ, 0)), nil >>`,
		Id:         "Result",
		NTType:     15,
		Index:      28,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return (make([]ast.Typ, 0)), nil
		},
	},
	ProdTabEntry{
		String: `Parameters : lparen ParameterList rparen	<< X[1], nil >>`,
		Id:         "Parameters",
		NTType:     16,
		Index:      29,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `ParameterList : ParameterList ParameterDecl	<< ast.ConcatParameterLists(X[0], X[1]) >>`,
		Id:         "ParameterList",
		NTType:     17,
		Index:      30,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.ConcatParameterLists(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `ParameterList : empty	<< (make([]ast.ParameterDecl, 0)), nil >>`,
		Id:         "ParameterList",
		NTType:     17,
		Index:      31,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return (make([]ast.ParameterDecl, 0)), nil
		},
	},
	ProdTabEntry{
		String: `ParameterDecl : IdentifierList Type	<< ast.MakeParameterDecl(X[0], X[1]) >>`,
		Id:         "ParameterDecl",
		NTType:     18,
		Index:      32,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.MakeParameterDecl(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `ParameterDecl : IdentifierList "..." Type	<< ast.Unsupported("UNSUPPORTED: variadic function arguments are not supported") >>`,
		Id:         "ParameterDecl",
		NTType:     18,
		Index:      33,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Unsupported("UNSUPPORTED: variadic function arguments are not supported")
		},
	},
	ProdTabEntry{
		String: `IdentifierList : IdentifierList comma id	<< ast.AppendIdList(X[0], X[2]) >>`,
		Id:         "IdentifierList",
		NTType:     19,
		Index:      34,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendIdList(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `IdentifierList : id	<< ast.NewIdList(X[0]) >>`,
		Id:         "IdentifierList",
		NTType:     19,
		Index:      35,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIdList(X[0])
		},
	},
	ProdTabEntry{
		String: `Receiver : lparen id ast id rparen	<< ast.NewPointerReceiver(X[1], X[3]) >>`,
		Id:         "Receiver",
		NTType:     20,
		Index:      36,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewPointerReceiver(X[1], X[3])
		},
	},
	ProdTabEntry{
		String: `Receiver : lparen id id rparen	<< ast.NewReceiver(X[1], X[2]) >>`,
		Id:         "Receiver",
		NTType:     20,
		Index:      37,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewReceiver(X[1], X[2])
		},
	},
	ProdTabEntry{
		String: `Declaration : empty	<< nil, nil >>`,
		Id:         "Declaration",
		NTType:     21,
		Index:      38,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `Block : lcurl BlockContents rcurl	<< ast.NewBlock(X[1]) >>`,
		Id:         "Block",
		NTType:     22,
		Index:      39,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBlock(X[1])
		},
	},
	ProdTabEntry{
		String: `BlockContents : BlockContents BlockContent	<< ast.AppendCodeList(X[0], X[1]) >>`,
		Id:         "BlockContents",
		NTType:     23,
		Index:      40,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendCodeList(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `BlockContents : BlockContent	<< ast.NewBlockContentList(X[0]) >>`,
		Id:         "BlockContents",
		NTType:     23,
		Index:      41,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBlockContentList(X[0])
		},
	},
	ProdTabEntry{
		String: `BlockContent : Block	<< X[0], nil >>`,
		Id:         "BlockContent",
		NTType:     24,
		Index:      42,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BlockContent : ChannelExpr	<< X[0], nil >>`,
		Id:         "BlockContent",
		NTType:     24,
		Index:      43,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BlockContent : Skip	<< X[0], nil >>`,
		Id:         "BlockContent",
		NTType:     24,
		Index:      44,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ChannelExpr : ChannelMake	<<  >>`,
		Id:         "ChannelExpr",
		NTType:     25,
		Index:      45,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ChannelExpr : ChannelReceive	<<  >>`,
		Id:         "ChannelExpr",
		NTType:     25,
		Index:      46,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ChannelExpr : ChannelSend	<< X[0], nil >>`,
		Id:         "ChannelExpr",
		NTType:     25,
		Index:      47,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ChannelMake : id assign kw_make lparen kw_capchan Type rparen	<< ast.NewCapChanMake(X[0], X[5]) >>`,
		Id:         "ChannelMake",
		NTType:     26,
		Index:      48,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCapChanMake(X[0], X[5])
		},
	},
	ProdTabEntry{
		String: `ChannelMake : id assign kw_make lparen kw_chan Type rparen	<< ast.NewChanMake(X[0], X[5]) >>`,
		Id:         "ChannelMake",
		NTType:     26,
		Index:      49,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewChanMake(X[0], X[5])
		},
	},
	ProdTabEntry{
		String: `ChannelMake : id assign kw_make lparen kw_chan Type comma	<< ast.Unsupported("UNSUPPORTED: buffered channels not supported") >>`,
		Id:         "ChannelMake",
		NTType:     26,
		Index:      50,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Unsupported("UNSUPPORTED: buffered channels not supported")
		},
	},
	ProdTabEntry{
		String: `ChannelReceive : id assign llarrow id	<< ast.NewCapChanReceive(X[0], X[3]) >>`,
		Id:         "ChannelReceive",
		NTType:     27,
		Index:      51,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCapChanReceive(X[0], X[3])
		},
	},
	ProdTabEntry{
		String: `ChannelSend : id llarrow id	<< ast.NewCapChanSend(X[0], X[2]) >>`,
		Id:         "ChannelSend",
		NTType:     28,
		Index:      52,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCapChanSend(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `TypeList : TypeList comma Type	<< ast.AppendTypeList(X[0], X[2]) >>`,
		Id:         "TypeList",
		NTType:     29,
		Index:      53,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendTypeList(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `TypeList : Type	<< ast.NewTypeList(X[0]) >>`,
		Id:         "TypeList",
		NTType:     29,
		Index:      54,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewTypeList(X[0])
		},
	},
	ProdTabEntry{
		String: `Type : lparen Type rparen	<< X[1], nil >>`,
		Id:         "Type",
		NTType:     30,
		Index:      55,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Type : StructType	<< X[0], nil >>`,
		Id:         "Type",
		NTType:     30,
		Index:      56,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : PointerType	<< X[0], nil >>`,
		Id:         "Type",
		NTType:     30,
		Index:      57,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : FunctionType	<< X[0], nil >>`,
		Id:         "Type",
		NTType:     30,
		Index:      58,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : InterfaceType	<< X[0], nil >>`,
		Id:         "Type",
		NTType:     30,
		Index:      59,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : SliceType	<< X[0], nil >>`,
		Id:         "Type",
		NTType:     30,
		Index:      60,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : ChannelType	<< X[0], nil >>`,
		Id:         "Type",
		NTType:     30,
		Index:      61,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : CapChannelType	<< X[0], nil >>`,
		Id:         "Type",
		NTType:     30,
		Index:      62,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : MapType	<< X[0], nil >>`,
		Id:         "Type",
		NTType:     30,
		Index:      63,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : kw_int	<< ast.IntType{}, nil >>`,
		Id:         "Type",
		NTType:     30,
		Index:      64,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.IntType{}, nil
		},
	},
	ProdTabEntry{
		String: `Type : kw_string	<< ast.StringType{}, nil >>`,
		Id:         "Type",
		NTType:     30,
		Index:      65,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.StringType{}, nil
		},
	},
	ProdTabEntry{
		String: `StructType : kw_struct lcurl FieldDecls rcurl	<< ast.NewStructType(X[2]) >>`,
		Id:         "StructType",
		NTType:     31,
		Index:      66,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStructType(X[2])
		},
	},
	ProdTabEntry{
		String: `FieldDecls : FieldDecls FieldDecl	<< ast.AppendStructFields(X[0], X[1]) >>`,
		Id:         "FieldDecls",
		NTType:     32,
		Index:      67,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendStructFields(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `FieldDecls : empty	<< (make([]ast.StructField, 0)), nil >>`,
		Id:         "FieldDecls",
		NTType:     32,
		Index:      68,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return (make([]ast.StructField, 0)), nil
		},
	},
	ProdTabEntry{
		String: `FieldDecl : IdentifierList Type	<< ast.MakeStructFields(X[0], X[1]) >>`,
		Id:         "FieldDecl",
		NTType:     33,
		Index:      69,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.MakeStructFields(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `PointerType : ast Type	<< ast.NewPointerType(X[1]) >>`,
		Id:         "PointerType",
		NTType:     34,
		Index:      70,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewPointerType(X[1])
		},
	},
	ProdTabEntry{
		String: `FunctionType : kw_func Signature	<< ast.NewFunctionType(X[1]) >>`,
		Id:         "FunctionType",
		NTType:     35,
		Index:      71,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFunctionType(X[1])
		},
	},
	ProdTabEntry{
		String: `InterfaceType : kw_interface lcurl InterfaceMethods rcurl	<< ast.NewInterfaceType(X[2]) >>`,
		Id:         "InterfaceType",
		NTType:     36,
		Index:      72,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewInterfaceType(X[2])
		},
	},
	ProdTabEntry{
		String: `InterfaceMethods : InterfaceMethods InterfaceMethod	<< ast.AppendInterfaceMethodList(X[0], X[1]) >>`,
		Id:         "InterfaceMethods",
		NTType:     37,
		Index:      73,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendInterfaceMethodList(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `InterfaceMethods : empty	<< make([]ast.InterfaceMethod, 0), nil >>`,
		Id:         "InterfaceMethods",
		NTType:     37,
		Index:      74,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return make([]ast.InterfaceMethod, 0), nil
		},
	},
	ProdTabEntry{
		String: `InterfaceMethod : id Signature	<< ast.NewInterfaceMethod(X[0], X[1]) >>`,
		Id:         "InterfaceMethod",
		NTType:     38,
		Index:      75,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewInterfaceMethod(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `SliceType : lbrack rbrack Type	<< ast.NewSliceType(X[2]) >>`,
		Id:         "SliceType",
		NTType:     39,
		Index:      76,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSliceType(X[2])
		},
	},
	ProdTabEntry{
		String: `MapType : kw_map lbrack Type rbrack Type	<< ast.NewMapType(X[2], X[4]) >>`,
		Id:         "MapType",
		NTType:     40,
		Index:      77,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewMapType(X[2], X[4])
		},
	},
	ProdTabEntry{
		String: `ChannelType : kw_chan Type	<< ast.NewChannelType(X[1]) >>`,
		Id:         "ChannelType",
		NTType:     41,
		Index:      78,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewChannelType(X[1])
		},
	},
	ProdTabEntry{
		String: `ChannelType : kw_chan larrow Type	<< ast.NewSOChannelType(X[2]) >>`,
		Id:         "ChannelType",
		NTType:     41,
		Index:      79,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSOChannelType(X[2])
		},
	},
	ProdTabEntry{
		String: `ChannelType : larrow kw_chan Type	<< ast.NewROChannelType(X[2]) >>`,
		Id:         "ChannelType",
		NTType:     41,
		Index:      80,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewROChannelType(X[2])
		},
	},
	ProdTabEntry{
		String: `CapChannelType : kw_capchan Type	<< ast.NewCapChanType(X[1]) >>`,
		Id:         "CapChannelType",
		NTType:     42,
		Index:      81,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCapChanType(X[1])
		},
	},
	ProdTabEntry{
		String: `CapChannelType : kw_capchan llarrow Type	<< ast.NewSOCapChanType(X[2]) >>`,
		Id:         "CapChannelType",
		NTType:     42,
		Index:      82,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSOCapChanType(X[2])
		},
	},
	ProdTabEntry{
		String: `CapChannelType : llarrow kw_capchan Type	<< ast.NewROCapChanType(X[2]) >>`,
		Id:         "CapChannelType",
		NTType:     42,
		Index:      83,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewROCapChanType(X[2])
		},
	},
	ProdTabEntry{
		String: `Skip : kw_capchan	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      84,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : kw_chan	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      85,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : kw_const	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      86,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : kw_func	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      87,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : kw_import	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      88,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : kw_interface	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      89,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : kw_map	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      90,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : kw_make	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      91,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : kw_package	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      92,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : kw_struct	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      93,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : kw_type	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      94,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : kw_var	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      95,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : lparen	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      96,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : rparen	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      97,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : dot	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      98,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : id	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      99,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : string_lit	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      100,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : ignored	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      101,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : terminator	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      102,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : comma	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      103,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : kw_int	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      104,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : kw_string	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      105,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : ast	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      106,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : lbrack	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      107,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : rbrack	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      108,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : lcurl	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      109,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : rcurl	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      110,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : larrow	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      111,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Skip : llarrow	<< ast.SkipToken(X[0]) >>`,
		Id:         "Skip",
		NTType:     43,
		Index:      112,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.SkipToken(X[0])
		},
	},
}
