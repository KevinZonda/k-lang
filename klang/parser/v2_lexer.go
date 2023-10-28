// Code generated from .//antlr4//V2Lexer.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type V2Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var V2LexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func v2lexerLexerInit() {
	staticData := &V2LexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'{'", "'}'", "'('", "')'", "'['", "']'", "','", "':'", "';'", "",
		"'=='", "'!='", "'>='", "'<='", "'>'", "'<'", "", "", "'+'", "'-'",
		"'*'", "'/'", "", "'.'", "'->'", "'struct'", "'map'", "'fn'", "'return'",
		"'case'", "'default'", "'open'", "'as'", "'if'", "'else'", "'for'",
		"'match'", "'break'", "", "", "", "", "", "':='",
	}
	staticData.SymbolicNames = []string{
		"", "LBrack", "RBrack", "LParen", "RParen", "LSquare", "RSquare", "Comma",
		"Col", "Semi", "BinaryOper", "Equals", "NotEq", "GreaterEq", "LessEq",
		"Greater", "Less", "Or", "And", "Add", "Sub", "Mul", "Div", "Mod", "Dot",
		"To", "Struct", "Map", "Function", "Return", "Case", "Default", "Open",
		"As", "If", "Else", "For", "Match", "Break", "Continue", "IntegerLiteral",
		"NumberLiteral", "StringLiteral", "Not", "Assign", "Identity", "WS",
		"NewLine",
	}
	staticData.RuleNames = []string{
		"LBrack", "RBrack", "LParen", "RParen", "LSquare", "RSquare", "Comma",
		"Col", "Semi", "BinaryOper", "Equals", "NotEq", "GreaterEq", "LessEq",
		"Greater", "Less", "Or", "And", "Add", "Sub", "Mul", "Div", "Mod", "Dot",
		"To", "Struct", "Map", "Function", "Return", "Case", "Default", "Open",
		"As", "If", "Else", "For", "Match", "Break", "Continue", "IntegerLiteral",
		"NumberLiteral", "StringLiteral", "Not", "Assign", "Identity", "WS",
		"NewLine",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 47, 323, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 127, 8, 9, 1, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1,
		14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 149, 8, 16,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 156, 8, 17, 1, 18, 1, 18, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22,
		170, 8, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38,
		1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 3, 38, 259, 8, 38, 1, 39, 3, 39, 262,
		8, 39, 1, 39, 4, 39, 265, 8, 39, 11, 39, 12, 39, 266, 1, 40, 3, 40, 270,
		8, 40, 1, 40, 4, 40, 273, 8, 40, 11, 40, 12, 40, 274, 1, 40, 1, 40, 4,
		40, 279, 8, 40, 11, 40, 12, 40, 280, 3, 40, 283, 8, 40, 1, 41, 1, 41, 5,
		41, 287, 8, 41, 10, 41, 12, 41, 290, 9, 41, 1, 41, 1, 41, 1, 42, 1, 42,
		1, 42, 1, 42, 3, 42, 298, 8, 42, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 5,
		44, 305, 8, 44, 10, 44, 12, 44, 308, 9, 44, 1, 45, 4, 45, 311, 8, 45, 11,
		45, 12, 45, 312, 1, 45, 1, 45, 1, 46, 3, 46, 318, 8, 46, 1, 46, 1, 46,
		3, 46, 322, 8, 46, 1, 288, 0, 47, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6,
		13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31,
		16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49,
		25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67,
		34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85,
		43, 87, 44, 89, 45, 91, 46, 93, 47, 1, 0, 4, 1, 0, 48, 57, 3, 0, 65, 90,
		95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 12,
		13, 32, 32, 351, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0,
		0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0,
		0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0,
		0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1,
		0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37,
		1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0,
		45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0,
		0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0,
		0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0,
		0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1,
		0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83,
		1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0,
		91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 1, 95, 1, 0, 0, 0, 3, 97, 1, 0, 0, 0,
		5, 99, 1, 0, 0, 0, 7, 101, 1, 0, 0, 0, 9, 103, 1, 0, 0, 0, 11, 105, 1,
		0, 0, 0, 13, 107, 1, 0, 0, 0, 15, 109, 1, 0, 0, 0, 17, 111, 1, 0, 0, 0,
		19, 126, 1, 0, 0, 0, 21, 128, 1, 0, 0, 0, 23, 131, 1, 0, 0, 0, 25, 134,
		1, 0, 0, 0, 27, 137, 1, 0, 0, 0, 29, 140, 1, 0, 0, 0, 31, 142, 1, 0, 0,
		0, 33, 148, 1, 0, 0, 0, 35, 155, 1, 0, 0, 0, 37, 157, 1, 0, 0, 0, 39, 159,
		1, 0, 0, 0, 41, 161, 1, 0, 0, 0, 43, 163, 1, 0, 0, 0, 45, 169, 1, 0, 0,
		0, 47, 171, 1, 0, 0, 0, 49, 173, 1, 0, 0, 0, 51, 176, 1, 0, 0, 0, 53, 183,
		1, 0, 0, 0, 55, 187, 1, 0, 0, 0, 57, 190, 1, 0, 0, 0, 59, 197, 1, 0, 0,
		0, 61, 202, 1, 0, 0, 0, 63, 210, 1, 0, 0, 0, 65, 215, 1, 0, 0, 0, 67, 218,
		1, 0, 0, 0, 69, 221, 1, 0, 0, 0, 71, 226, 1, 0, 0, 0, 73, 230, 1, 0, 0,
		0, 75, 236, 1, 0, 0, 0, 77, 258, 1, 0, 0, 0, 79, 261, 1, 0, 0, 0, 81, 269,
		1, 0, 0, 0, 83, 284, 1, 0, 0, 0, 85, 297, 1, 0, 0, 0, 87, 299, 1, 0, 0,
		0, 89, 302, 1, 0, 0, 0, 91, 310, 1, 0, 0, 0, 93, 321, 1, 0, 0, 0, 95, 96,
		5, 123, 0, 0, 96, 2, 1, 0, 0, 0, 97, 98, 5, 125, 0, 0, 98, 4, 1, 0, 0,
		0, 99, 100, 5, 40, 0, 0, 100, 6, 1, 0, 0, 0, 101, 102, 5, 41, 0, 0, 102,
		8, 1, 0, 0, 0, 103, 104, 5, 91, 0, 0, 104, 10, 1, 0, 0, 0, 105, 106, 5,
		93, 0, 0, 106, 12, 1, 0, 0, 0, 107, 108, 5, 44, 0, 0, 108, 14, 1, 0, 0,
		0, 109, 110, 5, 58, 0, 0, 110, 16, 1, 0, 0, 0, 111, 112, 5, 59, 0, 0, 112,
		18, 1, 0, 0, 0, 113, 127, 3, 21, 10, 0, 114, 127, 3, 23, 11, 0, 115, 127,
		3, 29, 14, 0, 116, 127, 3, 31, 15, 0, 117, 127, 3, 25, 12, 0, 118, 127,
		3, 27, 13, 0, 119, 127, 3, 33, 16, 0, 120, 127, 3, 35, 17, 0, 121, 127,
		3, 37, 18, 0, 122, 127, 3, 39, 19, 0, 123, 127, 3, 41, 20, 0, 124, 127,
		3, 43, 21, 0, 125, 127, 3, 45, 22, 0, 126, 113, 1, 0, 0, 0, 126, 114, 1,
		0, 0, 0, 126, 115, 1, 0, 0, 0, 126, 116, 1, 0, 0, 0, 126, 117, 1, 0, 0,
		0, 126, 118, 1, 0, 0, 0, 126, 119, 1, 0, 0, 0, 126, 120, 1, 0, 0, 0, 126,
		121, 1, 0, 0, 0, 126, 122, 1, 0, 0, 0, 126, 123, 1, 0, 0, 0, 126, 124,
		1, 0, 0, 0, 126, 125, 1, 0, 0, 0, 127, 20, 1, 0, 0, 0, 128, 129, 5, 61,
		0, 0, 129, 130, 5, 61, 0, 0, 130, 22, 1, 0, 0, 0, 131, 132, 5, 33, 0, 0,
		132, 133, 5, 61, 0, 0, 133, 24, 1, 0, 0, 0, 134, 135, 5, 62, 0, 0, 135,
		136, 5, 61, 0, 0, 136, 26, 1, 0, 0, 0, 137, 138, 5, 60, 0, 0, 138, 139,
		5, 61, 0, 0, 139, 28, 1, 0, 0, 0, 140, 141, 5, 62, 0, 0, 141, 30, 1, 0,
		0, 0, 142, 143, 5, 60, 0, 0, 143, 32, 1, 0, 0, 0, 144, 145, 5, 124, 0,
		0, 145, 149, 5, 124, 0, 0, 146, 147, 5, 111, 0, 0, 147, 149, 5, 114, 0,
		0, 148, 144, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 149, 34, 1, 0, 0, 0, 150,
		151, 5, 38, 0, 0, 151, 156, 5, 38, 0, 0, 152, 153, 5, 97, 0, 0, 153, 154,
		5, 110, 0, 0, 154, 156, 5, 100, 0, 0, 155, 150, 1, 0, 0, 0, 155, 152, 1,
		0, 0, 0, 156, 36, 1, 0, 0, 0, 157, 158, 5, 43, 0, 0, 158, 38, 1, 0, 0,
		0, 159, 160, 5, 45, 0, 0, 160, 40, 1, 0, 0, 0, 161, 162, 5, 42, 0, 0, 162,
		42, 1, 0, 0, 0, 163, 164, 5, 47, 0, 0, 164, 44, 1, 0, 0, 0, 165, 170, 5,
		37, 0, 0, 166, 167, 5, 109, 0, 0, 167, 168, 5, 111, 0, 0, 168, 170, 5,
		100, 0, 0, 169, 165, 1, 0, 0, 0, 169, 166, 1, 0, 0, 0, 170, 46, 1, 0, 0,
		0, 171, 172, 5, 46, 0, 0, 172, 48, 1, 0, 0, 0, 173, 174, 5, 45, 0, 0, 174,
		175, 5, 62, 0, 0, 175, 50, 1, 0, 0, 0, 176, 177, 5, 115, 0, 0, 177, 178,
		5, 116, 0, 0, 178, 179, 5, 114, 0, 0, 179, 180, 5, 117, 0, 0, 180, 181,
		5, 99, 0, 0, 181, 182, 5, 116, 0, 0, 182, 52, 1, 0, 0, 0, 183, 184, 5,
		109, 0, 0, 184, 185, 5, 97, 0, 0, 185, 186, 5, 112, 0, 0, 186, 54, 1, 0,
		0, 0, 187, 188, 5, 102, 0, 0, 188, 189, 5, 110, 0, 0, 189, 56, 1, 0, 0,
		0, 190, 191, 5, 114, 0, 0, 191, 192, 5, 101, 0, 0, 192, 193, 5, 116, 0,
		0, 193, 194, 5, 117, 0, 0, 194, 195, 5, 114, 0, 0, 195, 196, 5, 110, 0,
		0, 196, 58, 1, 0, 0, 0, 197, 198, 5, 99, 0, 0, 198, 199, 5, 97, 0, 0, 199,
		200, 5, 115, 0, 0, 200, 201, 5, 101, 0, 0, 201, 60, 1, 0, 0, 0, 202, 203,
		5, 100, 0, 0, 203, 204, 5, 101, 0, 0, 204, 205, 5, 102, 0, 0, 205, 206,
		5, 97, 0, 0, 206, 207, 5, 117, 0, 0, 207, 208, 5, 108, 0, 0, 208, 209,
		5, 116, 0, 0, 209, 62, 1, 0, 0, 0, 210, 211, 5, 111, 0, 0, 211, 212, 5,
		112, 0, 0, 212, 213, 5, 101, 0, 0, 213, 214, 5, 110, 0, 0, 214, 64, 1,
		0, 0, 0, 215, 216, 5, 97, 0, 0, 216, 217, 5, 115, 0, 0, 217, 66, 1, 0,
		0, 0, 218, 219, 5, 105, 0, 0, 219, 220, 5, 102, 0, 0, 220, 68, 1, 0, 0,
		0, 221, 222, 5, 101, 0, 0, 222, 223, 5, 108, 0, 0, 223, 224, 5, 115, 0,
		0, 224, 225, 5, 101, 0, 0, 225, 70, 1, 0, 0, 0, 226, 227, 5, 102, 0, 0,
		227, 228, 5, 111, 0, 0, 228, 229, 5, 114, 0, 0, 229, 72, 1, 0, 0, 0, 230,
		231, 5, 109, 0, 0, 231, 232, 5, 97, 0, 0, 232, 233, 5, 116, 0, 0, 233,
		234, 5, 99, 0, 0, 234, 235, 5, 104, 0, 0, 235, 74, 1, 0, 0, 0, 236, 237,
		5, 98, 0, 0, 237, 238, 5, 114, 0, 0, 238, 239, 5, 101, 0, 0, 239, 240,
		5, 97, 0, 0, 240, 241, 5, 107, 0, 0, 241, 76, 1, 0, 0, 0, 242, 243, 5,
		110, 0, 0, 243, 244, 5, 101, 0, 0, 244, 245, 5, 120, 0, 0, 245, 259, 5,
		116, 0, 0, 246, 247, 5, 112, 0, 0, 247, 248, 5, 97, 0, 0, 248, 249, 5,
		115, 0, 0, 249, 259, 5, 115, 0, 0, 250, 251, 5, 99, 0, 0, 251, 252, 5,
		111, 0, 0, 252, 253, 5, 110, 0, 0, 253, 254, 5, 116, 0, 0, 254, 255, 5,
		105, 0, 0, 255, 256, 5, 110, 0, 0, 256, 257, 5, 117, 0, 0, 257, 259, 5,
		101, 0, 0, 258, 242, 1, 0, 0, 0, 258, 246, 1, 0, 0, 0, 258, 250, 1, 0,
		0, 0, 259, 78, 1, 0, 0, 0, 260, 262, 5, 45, 0, 0, 261, 260, 1, 0, 0, 0,
		261, 262, 1, 0, 0, 0, 262, 264, 1, 0, 0, 0, 263, 265, 7, 0, 0, 0, 264,
		263, 1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0, 266, 267,
		1, 0, 0, 0, 267, 80, 1, 0, 0, 0, 268, 270, 5, 45, 0, 0, 269, 268, 1, 0,
		0, 0, 269, 270, 1, 0, 0, 0, 270, 272, 1, 0, 0, 0, 271, 273, 7, 0, 0, 0,
		272, 271, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 272, 1, 0, 0, 0, 274,
		275, 1, 0, 0, 0, 275, 282, 1, 0, 0, 0, 276, 278, 5, 46, 0, 0, 277, 279,
		7, 0, 0, 0, 278, 277, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 278, 1, 0,
		0, 0, 280, 281, 1, 0, 0, 0, 281, 283, 1, 0, 0, 0, 282, 276, 1, 0, 0, 0,
		282, 283, 1, 0, 0, 0, 283, 82, 1, 0, 0, 0, 284, 288, 5, 34, 0, 0, 285,
		287, 9, 0, 0, 0, 286, 285, 1, 0, 0, 0, 287, 290, 1, 0, 0, 0, 288, 289,
		1, 0, 0, 0, 288, 286, 1, 0, 0, 0, 289, 291, 1, 0, 0, 0, 290, 288, 1, 0,
		0, 0, 291, 292, 5, 34, 0, 0, 292, 84, 1, 0, 0, 0, 293, 298, 5, 33, 0, 0,
		294, 295, 5, 110, 0, 0, 295, 296, 5, 111, 0, 0, 296, 298, 5, 116, 0, 0,
		297, 293, 1, 0, 0, 0, 297, 294, 1, 0, 0, 0, 298, 86, 1, 0, 0, 0, 299, 300,
		5, 58, 0, 0, 300, 301, 5, 61, 0, 0, 301, 88, 1, 0, 0, 0, 302, 306, 7, 1,
		0, 0, 303, 305, 7, 2, 0, 0, 304, 303, 1, 0, 0, 0, 305, 308, 1, 0, 0, 0,
		306, 304, 1, 0, 0, 0, 306, 307, 1, 0, 0, 0, 307, 90, 1, 0, 0, 0, 308, 306,
		1, 0, 0, 0, 309, 311, 7, 3, 0, 0, 310, 309, 1, 0, 0, 0, 311, 312, 1, 0,
		0, 0, 312, 310, 1, 0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 314, 1, 0, 0, 0,
		314, 315, 6, 45, 0, 0, 315, 92, 1, 0, 0, 0, 316, 318, 5, 13, 0, 0, 317,
		316, 1, 0, 0, 0, 317, 318, 1, 0, 0, 0, 318, 319, 1, 0, 0, 0, 319, 322,
		5, 10, 0, 0, 320, 322, 5, 13, 0, 0, 321, 317, 1, 0, 0, 0, 321, 320, 1,
		0, 0, 0, 322, 94, 1, 0, 0, 0, 18, 0, 126, 148, 155, 169, 258, 261, 266,
		269, 274, 280, 282, 288, 297, 306, 312, 317, 321, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// V2LexerInit initializes any static state used to implement V2Lexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewV2Lexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func V2LexerInit() {
	staticData := &V2LexerLexerStaticData
	staticData.once.Do(v2lexerLexerInit)
}

// NewV2Lexer produces a new lexer instance for the optional input antlr.CharStream.
func NewV2Lexer(input antlr.CharStream) *V2Lexer {
	V2LexerInit()
	l := new(V2Lexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &V2LexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "V2Lexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// V2Lexer tokens.
const (
	V2LexerLBrack         = 1
	V2LexerRBrack         = 2
	V2LexerLParen         = 3
	V2LexerRParen         = 4
	V2LexerLSquare        = 5
	V2LexerRSquare        = 6
	V2LexerComma          = 7
	V2LexerCol            = 8
	V2LexerSemi           = 9
	V2LexerBinaryOper     = 10
	V2LexerEquals         = 11
	V2LexerNotEq          = 12
	V2LexerGreaterEq      = 13
	V2LexerLessEq         = 14
	V2LexerGreater        = 15
	V2LexerLess           = 16
	V2LexerOr             = 17
	V2LexerAnd            = 18
	V2LexerAdd            = 19
	V2LexerSub            = 20
	V2LexerMul            = 21
	V2LexerDiv            = 22
	V2LexerMod            = 23
	V2LexerDot            = 24
	V2LexerTo             = 25
	V2LexerStruct         = 26
	V2LexerMap            = 27
	V2LexerFunction       = 28
	V2LexerReturn         = 29
	V2LexerCase           = 30
	V2LexerDefault        = 31
	V2LexerOpen           = 32
	V2LexerAs             = 33
	V2LexerIf             = 34
	V2LexerElse           = 35
	V2LexerFor            = 36
	V2LexerMatch          = 37
	V2LexerBreak          = 38
	V2LexerContinue       = 39
	V2LexerIntegerLiteral = 40
	V2LexerNumberLiteral  = 41
	V2LexerStringLiteral  = 42
	V2LexerNot            = 43
	V2LexerAssign         = 44
	V2LexerIdentity       = 45
	V2LexerWS             = 46
	V2LexerNewLine        = 47
)