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
		"", "'{'", "'}'", "'('", "')'", "'['", "']'", "','", "':'", "';'", "'=='",
		"'!='", "'>='", "'<='", "'>'", "'<'", "", "", "", "'*'", "'/'", "'+'",
		"'-'", "", "'.'", "'?'", "", "'struct'", "'map'", "'fn'", "'return'",
		"'case'", "'default'", "'open'", "'as'", "'try'", "'catch'", "'if'",
		"'else'", "'for'", "'match'", "'break'", "'continue'", "'true'", "'false'",
		"'nil'", "", "", "", "", "", "'&'",
	}
	staticData.SymbolicNames = []string{
		"", "LBrack", "RBrack", "LParen", "RParen", "LSquare", "RSquare", "Comma",
		"Col", "Semi", "Equals", "NotEq", "GreaterEq", "LessEq", "Greater",
		"Less", "Or", "And", "Pow", "Mul", "Div", "Add", "Sub", "Mod", "Dot",
		"Question", "To", "Struct", "Map", "Function", "Return", "Case", "Default",
		"Open", "As", "Try", "Catch", "If", "Else", "For", "Match", "Break",
		"Continue", "True", "False", "Nil", "IntegerLiteral", "NumberLiteral",
		"StringLiteral", "Not", "Assign", "Ref", "Identifier", "Comment", "WS",
		"NewLine",
	}
	staticData.RuleNames = []string{
		"LBrack", "RBrack", "LParen", "RParen", "LSquare", "RSquare", "Comma",
		"Col", "Semi", "Equals", "NotEq", "GreaterEq", "LessEq", "Greater",
		"Less", "Or", "And", "Pow", "Mul", "Div", "Add", "Sub", "Mod", "Dot",
		"Question", "To", "Struct", "Map", "Function", "Return", "Case", "Default",
		"Open", "As", "Try", "Catch", "If", "Else", "For", "Match", "Break",
		"Continue", "True", "False", "Nil", "IntegerLiteral", "NumberLiteral",
		"StringLiteral", "Not", "Assign", "Ref", "Identifier", "Comment", "WS",
		"NewLine",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 55, 394, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15,
		150, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 157, 8, 16, 1, 17,
		1, 17, 1, 17, 3, 17, 162, 8, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 176, 8, 22, 1, 23,
		1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 186, 8, 25, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36,
		1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1,
		38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40,
		1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1,
		41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43,
		1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 4, 45, 289, 8, 45, 11, 45, 12,
		45, 290, 1, 46, 4, 46, 294, 8, 46, 11, 46, 12, 46, 295, 1, 46, 1, 46, 4,
		46, 300, 8, 46, 11, 46, 12, 46, 301, 3, 46, 304, 8, 46, 1, 46, 1, 46, 4,
		46, 308, 8, 46, 11, 46, 12, 46, 309, 3, 46, 312, 8, 46, 1, 47, 3, 47, 315,
		8, 47, 1, 47, 1, 47, 1, 47, 1, 47, 5, 47, 321, 8, 47, 10, 47, 12, 47, 324,
		9, 47, 1, 47, 1, 47, 3, 47, 328, 8, 47, 1, 47, 1, 47, 1, 47, 1, 47, 5,
		47, 334, 8, 47, 10, 47, 12, 47, 337, 9, 47, 1, 47, 3, 47, 340, 8, 47, 1,
		48, 1, 48, 1, 48, 1, 48, 3, 48, 346, 8, 48, 1, 49, 1, 49, 1, 49, 1, 49,
		1, 49, 3, 49, 353, 8, 49, 1, 50, 1, 50, 1, 51, 1, 51, 5, 51, 359, 8, 51,
		10, 51, 12, 51, 362, 9, 51, 1, 52, 1, 52, 1, 52, 3, 52, 367, 8, 52, 1,
		52, 5, 52, 370, 8, 52, 10, 52, 12, 52, 373, 9, 52, 1, 52, 1, 52, 1, 53,
		4, 53, 378, 8, 53, 11, 53, 12, 53, 379, 1, 53, 1, 53, 1, 54, 3, 54, 385,
		8, 54, 1, 54, 1, 54, 4, 54, 389, 8, 54, 11, 54, 12, 54, 390, 1, 54, 1,
		54, 0, 0, 55, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9,
		19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18,
		37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27,
		55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36,
		73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45,
		91, 46, 93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 107,
		54, 109, 55, 1, 0, 8, 1, 0, 48, 57, 2, 0, 36, 36, 64, 64, 1, 0, 34, 34,
		1, 0, 39, 39, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95,
		95, 97, 122, 2, 0, 10, 10, 13, 13, 3, 0, 9, 9, 12, 12, 32, 32, 421, 0,
		1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0,
		9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0,
		0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0,
		0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0,
		0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1,
		0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47,
		1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0,
		55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0,
		0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0,
		0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0,
		0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1,
		0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93,
		1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0,
		101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0,
		0, 0, 0, 109, 1, 0, 0, 0, 1, 111, 1, 0, 0, 0, 3, 113, 1, 0, 0, 0, 5, 115,
		1, 0, 0, 0, 7, 117, 1, 0, 0, 0, 9, 119, 1, 0, 0, 0, 11, 121, 1, 0, 0, 0,
		13, 123, 1, 0, 0, 0, 15, 125, 1, 0, 0, 0, 17, 127, 1, 0, 0, 0, 19, 129,
		1, 0, 0, 0, 21, 132, 1, 0, 0, 0, 23, 135, 1, 0, 0, 0, 25, 138, 1, 0, 0,
		0, 27, 141, 1, 0, 0, 0, 29, 143, 1, 0, 0, 0, 31, 149, 1, 0, 0, 0, 33, 156,
		1, 0, 0, 0, 35, 161, 1, 0, 0, 0, 37, 163, 1, 0, 0, 0, 39, 165, 1, 0, 0,
		0, 41, 167, 1, 0, 0, 0, 43, 169, 1, 0, 0, 0, 45, 175, 1, 0, 0, 0, 47, 177,
		1, 0, 0, 0, 49, 179, 1, 0, 0, 0, 51, 185, 1, 0, 0, 0, 53, 187, 1, 0, 0,
		0, 55, 194, 1, 0, 0, 0, 57, 198, 1, 0, 0, 0, 59, 201, 1, 0, 0, 0, 61, 208,
		1, 0, 0, 0, 63, 213, 1, 0, 0, 0, 65, 221, 1, 0, 0, 0, 67, 226, 1, 0, 0,
		0, 69, 229, 1, 0, 0, 0, 71, 233, 1, 0, 0, 0, 73, 239, 1, 0, 0, 0, 75, 242,
		1, 0, 0, 0, 77, 247, 1, 0, 0, 0, 79, 251, 1, 0, 0, 0, 81, 257, 1, 0, 0,
		0, 83, 263, 1, 0, 0, 0, 85, 272, 1, 0, 0, 0, 87, 277, 1, 0, 0, 0, 89, 283,
		1, 0, 0, 0, 91, 288, 1, 0, 0, 0, 93, 311, 1, 0, 0, 0, 95, 339, 1, 0, 0,
		0, 97, 345, 1, 0, 0, 0, 99, 352, 1, 0, 0, 0, 101, 354, 1, 0, 0, 0, 103,
		356, 1, 0, 0, 0, 105, 366, 1, 0, 0, 0, 107, 377, 1, 0, 0, 0, 109, 388,
		1, 0, 0, 0, 111, 112, 5, 123, 0, 0, 112, 2, 1, 0, 0, 0, 113, 114, 5, 125,
		0, 0, 114, 4, 1, 0, 0, 0, 115, 116, 5, 40, 0, 0, 116, 6, 1, 0, 0, 0, 117,
		118, 5, 41, 0, 0, 118, 8, 1, 0, 0, 0, 119, 120, 5, 91, 0, 0, 120, 10, 1,
		0, 0, 0, 121, 122, 5, 93, 0, 0, 122, 12, 1, 0, 0, 0, 123, 124, 5, 44, 0,
		0, 124, 14, 1, 0, 0, 0, 125, 126, 5, 58, 0, 0, 126, 16, 1, 0, 0, 0, 127,
		128, 5, 59, 0, 0, 128, 18, 1, 0, 0, 0, 129, 130, 5, 61, 0, 0, 130, 131,
		5, 61, 0, 0, 131, 20, 1, 0, 0, 0, 132, 133, 5, 33, 0, 0, 133, 134, 5, 61,
		0, 0, 134, 22, 1, 0, 0, 0, 135, 136, 5, 62, 0, 0, 136, 137, 5, 61, 0, 0,
		137, 24, 1, 0, 0, 0, 138, 139, 5, 60, 0, 0, 139, 140, 5, 61, 0, 0, 140,
		26, 1, 0, 0, 0, 141, 142, 5, 62, 0, 0, 142, 28, 1, 0, 0, 0, 143, 144, 5,
		60, 0, 0, 144, 30, 1, 0, 0, 0, 145, 146, 5, 124, 0, 0, 146, 150, 5, 124,
		0, 0, 147, 148, 5, 111, 0, 0, 148, 150, 5, 114, 0, 0, 149, 145, 1, 0, 0,
		0, 149, 147, 1, 0, 0, 0, 150, 32, 1, 0, 0, 0, 151, 152, 5, 38, 0, 0, 152,
		157, 5, 38, 0, 0, 153, 154, 5, 97, 0, 0, 154, 155, 5, 110, 0, 0, 155, 157,
		5, 100, 0, 0, 156, 151, 1, 0, 0, 0, 156, 153, 1, 0, 0, 0, 157, 34, 1, 0,
		0, 0, 158, 162, 5, 94, 0, 0, 159, 160, 5, 42, 0, 0, 160, 162, 5, 42, 0,
		0, 161, 158, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 162, 36, 1, 0, 0, 0, 163,
		164, 5, 42, 0, 0, 164, 38, 1, 0, 0, 0, 165, 166, 5, 47, 0, 0, 166, 40,
		1, 0, 0, 0, 167, 168, 5, 43, 0, 0, 168, 42, 1, 0, 0, 0, 169, 170, 5, 45,
		0, 0, 170, 44, 1, 0, 0, 0, 171, 176, 5, 37, 0, 0, 172, 173, 5, 109, 0,
		0, 173, 174, 5, 111, 0, 0, 174, 176, 5, 100, 0, 0, 175, 171, 1, 0, 0, 0,
		175, 172, 1, 0, 0, 0, 176, 46, 1, 0, 0, 0, 177, 178, 5, 46, 0, 0, 178,
		48, 1, 0, 0, 0, 179, 180, 5, 63, 0, 0, 180, 50, 1, 0, 0, 0, 181, 182, 5,
		45, 0, 0, 182, 186, 5, 62, 0, 0, 183, 184, 5, 61, 0, 0, 184, 186, 5, 62,
		0, 0, 185, 181, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 186, 52, 1, 0, 0, 0,
		187, 188, 5, 115, 0, 0, 188, 189, 5, 116, 0, 0, 189, 190, 5, 114, 0, 0,
		190, 191, 5, 117, 0, 0, 191, 192, 5, 99, 0, 0, 192, 193, 5, 116, 0, 0,
		193, 54, 1, 0, 0, 0, 194, 195, 5, 109, 0, 0, 195, 196, 5, 97, 0, 0, 196,
		197, 5, 112, 0, 0, 197, 56, 1, 0, 0, 0, 198, 199, 5, 102, 0, 0, 199, 200,
		5, 110, 0, 0, 200, 58, 1, 0, 0, 0, 201, 202, 5, 114, 0, 0, 202, 203, 5,
		101, 0, 0, 203, 204, 5, 116, 0, 0, 204, 205, 5, 117, 0, 0, 205, 206, 5,
		114, 0, 0, 206, 207, 5, 110, 0, 0, 207, 60, 1, 0, 0, 0, 208, 209, 5, 99,
		0, 0, 209, 210, 5, 97, 0, 0, 210, 211, 5, 115, 0, 0, 211, 212, 5, 101,
		0, 0, 212, 62, 1, 0, 0, 0, 213, 214, 5, 100, 0, 0, 214, 215, 5, 101, 0,
		0, 215, 216, 5, 102, 0, 0, 216, 217, 5, 97, 0, 0, 217, 218, 5, 117, 0,
		0, 218, 219, 5, 108, 0, 0, 219, 220, 5, 116, 0, 0, 220, 64, 1, 0, 0, 0,
		221, 222, 5, 111, 0, 0, 222, 223, 5, 112, 0, 0, 223, 224, 5, 101, 0, 0,
		224, 225, 5, 110, 0, 0, 225, 66, 1, 0, 0, 0, 226, 227, 5, 97, 0, 0, 227,
		228, 5, 115, 0, 0, 228, 68, 1, 0, 0, 0, 229, 230, 5, 116, 0, 0, 230, 231,
		5, 114, 0, 0, 231, 232, 5, 121, 0, 0, 232, 70, 1, 0, 0, 0, 233, 234, 5,
		99, 0, 0, 234, 235, 5, 97, 0, 0, 235, 236, 5, 116, 0, 0, 236, 237, 5, 99,
		0, 0, 237, 238, 5, 104, 0, 0, 238, 72, 1, 0, 0, 0, 239, 240, 5, 105, 0,
		0, 240, 241, 5, 102, 0, 0, 241, 74, 1, 0, 0, 0, 242, 243, 5, 101, 0, 0,
		243, 244, 5, 108, 0, 0, 244, 245, 5, 115, 0, 0, 245, 246, 5, 101, 0, 0,
		246, 76, 1, 0, 0, 0, 247, 248, 5, 102, 0, 0, 248, 249, 5, 111, 0, 0, 249,
		250, 5, 114, 0, 0, 250, 78, 1, 0, 0, 0, 251, 252, 5, 109, 0, 0, 252, 253,
		5, 97, 0, 0, 253, 254, 5, 116, 0, 0, 254, 255, 5, 99, 0, 0, 255, 256, 5,
		104, 0, 0, 256, 80, 1, 0, 0, 0, 257, 258, 5, 98, 0, 0, 258, 259, 5, 114,
		0, 0, 259, 260, 5, 101, 0, 0, 260, 261, 5, 97, 0, 0, 261, 262, 5, 107,
		0, 0, 262, 82, 1, 0, 0, 0, 263, 264, 5, 99, 0, 0, 264, 265, 5, 111, 0,
		0, 265, 266, 5, 110, 0, 0, 266, 267, 5, 116, 0, 0, 267, 268, 5, 105, 0,
		0, 268, 269, 5, 110, 0, 0, 269, 270, 5, 117, 0, 0, 270, 271, 5, 101, 0,
		0, 271, 84, 1, 0, 0, 0, 272, 273, 5, 116, 0, 0, 273, 274, 5, 114, 0, 0,
		274, 275, 5, 117, 0, 0, 275, 276, 5, 101, 0, 0, 276, 86, 1, 0, 0, 0, 277,
		278, 5, 102, 0, 0, 278, 279, 5, 97, 0, 0, 279, 280, 5, 108, 0, 0, 280,
		281, 5, 115, 0, 0, 281, 282, 5, 101, 0, 0, 282, 88, 1, 0, 0, 0, 283, 284,
		5, 110, 0, 0, 284, 285, 5, 105, 0, 0, 285, 286, 5, 108, 0, 0, 286, 90,
		1, 0, 0, 0, 287, 289, 7, 0, 0, 0, 288, 287, 1, 0, 0, 0, 289, 290, 1, 0,
		0, 0, 290, 288, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 92, 1, 0, 0, 0,
		292, 294, 7, 0, 0, 0, 293, 292, 1, 0, 0, 0, 294, 295, 1, 0, 0, 0, 295,
		293, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 303, 1, 0, 0, 0, 297, 299,
		5, 46, 0, 0, 298, 300, 7, 0, 0, 0, 299, 298, 1, 0, 0, 0, 300, 301, 1, 0,
		0, 0, 301, 299, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302, 304, 1, 0, 0, 0,
		303, 297, 1, 0, 0, 0, 303, 304, 1, 0, 0, 0, 304, 312, 1, 0, 0, 0, 305,
		307, 5, 46, 0, 0, 306, 308, 7, 0, 0, 0, 307, 306, 1, 0, 0, 0, 308, 309,
		1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 309, 310, 1, 0, 0, 0, 310, 312, 1, 0,
		0, 0, 311, 293, 1, 0, 0, 0, 311, 305, 1, 0, 0, 0, 312, 94, 1, 0, 0, 0,
		313, 315, 7, 1, 0, 0, 314, 313, 1, 0, 0, 0, 314, 315, 1, 0, 0, 0, 315,
		316, 1, 0, 0, 0, 316, 322, 5, 34, 0, 0, 317, 321, 8, 2, 0, 0, 318, 319,
		5, 92, 0, 0, 319, 321, 5, 34, 0, 0, 320, 317, 1, 0, 0, 0, 320, 318, 1,
		0, 0, 0, 321, 324, 1, 0, 0, 0, 322, 320, 1, 0, 0, 0, 322, 323, 1, 0, 0,
		0, 323, 325, 1, 0, 0, 0, 324, 322, 1, 0, 0, 0, 325, 340, 5, 34, 0, 0, 326,
		328, 7, 1, 0, 0, 327, 326, 1, 0, 0, 0, 327, 328, 1, 0, 0, 0, 328, 329,
		1, 0, 0, 0, 329, 335, 5, 39, 0, 0, 330, 334, 8, 3, 0, 0, 331, 332, 5, 92,
		0, 0, 332, 334, 5, 39, 0, 0, 333, 330, 1, 0, 0, 0, 333, 331, 1, 0, 0, 0,
		334, 337, 1, 0, 0, 0, 335, 333, 1, 0, 0, 0, 335, 336, 1, 0, 0, 0, 336,
		338, 1, 0, 0, 0, 337, 335, 1, 0, 0, 0, 338, 340, 5, 39, 0, 0, 339, 314,
		1, 0, 0, 0, 339, 327, 1, 0, 0, 0, 340, 96, 1, 0, 0, 0, 341, 346, 5, 33,
		0, 0, 342, 343, 5, 110, 0, 0, 343, 344, 5, 111, 0, 0, 344, 346, 5, 116,
		0, 0, 345, 341, 1, 0, 0, 0, 345, 342, 1, 0, 0, 0, 346, 98, 1, 0, 0, 0,
		347, 348, 5, 58, 0, 0, 348, 353, 5, 61, 0, 0, 349, 353, 5, 61, 0, 0, 350,
		351, 5, 60, 0, 0, 351, 353, 5, 45, 0, 0, 352, 347, 1, 0, 0, 0, 352, 349,
		1, 0, 0, 0, 352, 350, 1, 0, 0, 0, 353, 100, 1, 0, 0, 0, 354, 355, 5, 38,
		0, 0, 355, 102, 1, 0, 0, 0, 356, 360, 7, 4, 0, 0, 357, 359, 7, 5, 0, 0,
		358, 357, 1, 0, 0, 0, 359, 362, 1, 0, 0, 0, 360, 358, 1, 0, 0, 0, 360,
		361, 1, 0, 0, 0, 361, 104, 1, 0, 0, 0, 362, 360, 1, 0, 0, 0, 363, 367,
		5, 35, 0, 0, 364, 365, 5, 47, 0, 0, 365, 367, 5, 47, 0, 0, 366, 363, 1,
		0, 0, 0, 366, 364, 1, 0, 0, 0, 367, 371, 1, 0, 0, 0, 368, 370, 8, 6, 0,
		0, 369, 368, 1, 0, 0, 0, 370, 373, 1, 0, 0, 0, 371, 369, 1, 0, 0, 0, 371,
		372, 1, 0, 0, 0, 372, 374, 1, 0, 0, 0, 373, 371, 1, 0, 0, 0, 374, 375,
		6, 52, 0, 0, 375, 106, 1, 0, 0, 0, 376, 378, 7, 7, 0, 0, 377, 376, 1, 0,
		0, 0, 378, 379, 1, 0, 0, 0, 379, 377, 1, 0, 0, 0, 379, 380, 1, 0, 0, 0,
		380, 381, 1, 0, 0, 0, 381, 382, 6, 53, 1, 0, 382, 108, 1, 0, 0, 0, 383,
		385, 5, 13, 0, 0, 384, 383, 1, 0, 0, 0, 384, 385, 1, 0, 0, 0, 385, 386,
		1, 0, 0, 0, 386, 389, 5, 10, 0, 0, 387, 389, 5, 13, 0, 0, 388, 384, 1,
		0, 0, 0, 388, 387, 1, 0, 0, 0, 389, 390, 1, 0, 0, 0, 390, 388, 1, 0, 0,
		0, 390, 391, 1, 0, 0, 0, 391, 392, 1, 0, 0, 0, 392, 393, 6, 54, 2, 0, 393,
		110, 1, 0, 0, 0, 28, 0, 149, 156, 161, 175, 185, 290, 295, 301, 303, 309,
		311, 314, 320, 322, 327, 333, 335, 339, 345, 352, 360, 366, 371, 379, 384,
		388, 390, 3, 0, 2, 0, 6, 0, 0, 0, 1, 0,
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
	V2LexerEquals         = 10
	V2LexerNotEq          = 11
	V2LexerGreaterEq      = 12
	V2LexerLessEq         = 13
	V2LexerGreater        = 14
	V2LexerLess           = 15
	V2LexerOr             = 16
	V2LexerAnd            = 17
	V2LexerPow            = 18
	V2LexerMul            = 19
	V2LexerDiv            = 20
	V2LexerAdd            = 21
	V2LexerSub            = 22
	V2LexerMod            = 23
	V2LexerDot            = 24
	V2LexerQuestion       = 25
	V2LexerTo             = 26
	V2LexerStruct         = 27
	V2LexerMap            = 28
	V2LexerFunction       = 29
	V2LexerReturn         = 30
	V2LexerCase           = 31
	V2LexerDefault        = 32
	V2LexerOpen           = 33
	V2LexerAs             = 34
	V2LexerTry            = 35
	V2LexerCatch          = 36
	V2LexerIf             = 37
	V2LexerElse           = 38
	V2LexerFor            = 39
	V2LexerMatch          = 40
	V2LexerBreak          = 41
	V2LexerContinue       = 42
	V2LexerTrue           = 43
	V2LexerFalse          = 44
	V2LexerNil            = 45
	V2LexerIntegerLiteral = 46
	V2LexerNumberLiteral  = 47
	V2LexerStringLiteral  = 48
	V2LexerNot            = 49
	V2LexerAssign         = 50
	V2LexerRef            = 51
	V2LexerIdentifier     = 52
	V2LexerComment        = 53
	V2LexerWS             = 54
	V2LexerNewLine        = 55
)
