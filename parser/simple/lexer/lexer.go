// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"io/ioutil"
	"unicode/utf8"

	"github.com/nfk93/gocap/parser/simple/token"
)

const (
	NoState    = -1
	NumStates  = 91
	NumSymbols = 118
)

type Lexer struct {
	src    []byte
	pos    int
	line   int
	column int
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
	}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewLexer(src), nil
}

func (l *Lexer) Scan() (tok *token.Token) {
	tok = new(token.Token)
	if l.pos >= len(l.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = l.pos, l.line, l.column
		return
	}
	start, startLine, startColumn, end := l.pos, l.line, l.column, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {
		if l.pos >= len(l.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(l.src[l.pos:])
			l.pos += size
		}

		nextState := -1
		if rune1 != -1 {
			nextState = TransTab[state](rune1)
		}
		state = nextState

		if state != -1 {

			switch rune1 {
			case '\n':
				l.line++
				l.column = 1
			case '\r':
				l.column = 1
			case '\t':
				l.column += 4
			default:
				l.column++
			}

			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				end = l.pos
			case ActTab[state].Ignore != "":
				start, startLine, startColumn = l.pos, l.line, l.column
				state = 0
				if start >= len(l.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = l.pos
			}
		}
	}
	if end > start {
		l.pos = end
		tok.Lit = l.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = start, startLine, startColumn

	return
}

func (l *Lexer) Reset() {
	l.pos = 0
}

/*
Lexer symbols:
0: '\n'
1: ';'
2: 'c'
3: 'a'
4: 'p'
5: 'c'
6: 'h'
7: 'a'
8: 'n'
9: 'c'
10: 'h'
11: 'a'
12: 'n'
13: 'c'
14: 'o'
15: 'n'
16: 's'
17: 't'
18: 'f'
19: 'u'
20: 'n'
21: 'c'
22: 'i'
23: 'm'
24: 'p'
25: 'o'
26: 'r'
27: 't'
28: 'i'
29: 'n'
30: 't'
31: 'e'
32: 'r'
33: 'f'
34: 'a'
35: 'c'
36: 'e'
37: 'i'
38: 'n'
39: 't'
40: 'm'
41: 'a'
42: 'p'
43: 'm'
44: 'a'
45: 'k'
46: 'e'
47: 'p'
48: 'a'
49: 'c'
50: 'k'
51: 'a'
52: 'g'
53: 'e'
54: 's'
55: 'r'
56: 'i'
57: 'n'
58: 'g'
59: 's'
60: 't'
61: 'r'
62: 'u'
63: 'c'
64: 't'
65: 't'
66: 'y'
67: 'p'
68: 'e'
69: 'v'
70: 'a'
71: 'r'
72: '('
73: ')'
74: '['
75: ']'
76: '{'
77: '}'
78: '.'
79: ','
80: '*'
81: '<'
82: '-'
83: '<'
84: '-'
85: '-'
86: ':'
87: '='
88: '='
89: '_'
90: '.'
91: '.'
92: '.'
93: '_'
94: '\'
95: 'a'
96: 'b'
97: 'f'
98: 'n'
99: 'r'
100: 't'
101: 'v'
102: '\'
103: '''
104: '`'
105: '`'
106: '`'
107: '"'
108: '"'
109: ' '
110: '\t'
111: '\r'
112: 'a'-'z'
113: 'A'-'Z'
114: 'a'-'z'
115: 'A'-'Z'
116: '0'-'9'
117: .
*/
