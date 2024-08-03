// Code generated from ./antlr4/V2Lexer.g4 by ANTLR 4.13.1. DO NOT EDIT.

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
		"'-'", "", "'.'", "'?'", "'++'", "'--'", "", "'struct'", "'map'", "",
		"'return'", "'case'", "'default'", "", "'as'", "'try'", "'catch'", "'if'",
		"'else'", "'for'", "'match'", "'break'", "'continue'", "'true'", "'false'",
		"'nil'", "", "", "", "", "", "'&'",
	}
	staticData.SymbolicNames = []string{
		"", "LBrack", "RBrack", "LParen", "RParen", "LSquare", "RSquare", "Comma",
		"Col", "Semi", "Equals", "NotEq", "GreaterEq", "LessEq", "Greater",
		"Less", "Or", "And", "Pow", "Mul", "Div", "Add", "Sub", "Mod", "Dot",
		"Question", "AddAdd", "SubSub", "To", "Struct", "Map", "Function", "Return",
		"Case", "Default", "Open", "As", "Try", "Catch", "If", "Else", "For",
		"Match", "Break", "Continue", "True", "False", "Nil", "IntegerLiteral",
		"NumberLiteral", "StringLiteral", "Not", "Assign", "Ref", "Identifier",
		"Comment", "BlkComment", "WS", "NewLine",
	}
	staticData.RuleNames = []string{
		"LBrack", "RBrack", "LParen", "RParen", "LSquare", "RSquare", "Comma",
		"Col", "Semi", "Equals", "NotEq", "GreaterEq", "LessEq", "Greater",
		"Less", "Or", "And", "Pow", "Mul", "Div", "Add", "Sub", "Mod", "Dot",
		"Question", "AddAdd", "SubSub", "To", "Struct", "Map", "Function", "Return",
		"Case", "Default", "Open", "As", "Try", "Catch", "If", "Else", "For",
		"Match", "Break", "Continue", "True", "False", "Nil", "IntegerLiteral",
		"NumberLiteral", "StringLiteral", "Not", "Assign", "Ref", "Identifier",
		"Comment", "BlkComment", "WS", "NewLine",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 58, 423, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 156, 8, 15, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 3, 16, 163, 8, 16, 1, 17, 1, 17, 1, 17, 3, 17, 168, 8, 17,
		1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1,
		22, 1, 22, 3, 22, 182, 8, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 25, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 198, 8,
		27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29,
		1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 217, 8, 30, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 249,
		8, 34, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39,
		1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1,
		41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43,
		1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1,
		45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 47,
		4, 47, 313, 8, 47, 11, 47, 12, 47, 314, 1, 48, 5, 48, 318, 8, 48, 10, 48,
		12, 48, 321, 9, 48, 1, 48, 1, 48, 4, 48, 325, 8, 48, 11, 48, 12, 48, 326,
		1, 49, 3, 49, 330, 8, 49, 1, 49, 1, 49, 1, 49, 1, 49, 5, 49, 336, 8, 49,
		10, 49, 12, 49, 339, 9, 49, 1, 49, 1, 49, 3, 49, 343, 8, 49, 1, 49, 1,
		49, 1, 49, 1, 49, 5, 49, 349, 8, 49, 10, 49, 12, 49, 352, 9, 49, 1, 49,
		3, 49, 355, 8, 49, 1, 50, 1, 50, 1, 50, 1, 50, 3, 50, 361, 8, 50, 1, 51,
		1, 51, 1, 51, 1, 51, 1, 51, 3, 51, 368, 8, 51, 1, 52, 1, 52, 1, 53, 1,
		53, 5, 53, 374, 8, 53, 10, 53, 12, 53, 377, 9, 53, 1, 54, 1, 54, 1, 54,
		3, 54, 382, 8, 54, 1, 54, 5, 54, 385, 8, 54, 10, 54, 12, 54, 388, 9, 54,
		1, 54, 1, 54, 1, 55, 1, 55, 1, 55, 1, 55, 5, 55, 396, 8, 55, 10, 55, 12,
		55, 399, 9, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 56, 4, 56, 407, 8,
		56, 11, 56, 12, 56, 408, 1, 56, 1, 56, 1, 57, 3, 57, 414, 8, 57, 1, 57,
		1, 57, 4, 57, 418, 8, 57, 11, 57, 12, 57, 419, 1, 57, 1, 57, 1, 397, 0,
		58, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21,
		11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39,
		20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57,
		29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75,
		38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93,
		47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 107, 54, 109, 55,
		111, 56, 113, 57, 115, 58, 1, 0, 8, 1, 0, 48, 57, 2, 0, 36, 36, 64, 64,
		1, 0, 34, 34, 1, 0, 39, 39, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57,
		65, 90, 95, 95, 97, 122, 2, 0, 10, 10, 13, 13, 3, 0, 9, 9, 12, 12, 32,
		32, 450, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0,
		0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1,
		0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61,
		1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0,
		69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0,
		0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0,
		0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0,
		0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1,
		0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0,
		107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1, 0,
		0, 0, 0, 115, 1, 0, 0, 0, 1, 117, 1, 0, 0, 0, 3, 119, 1, 0, 0, 0, 5, 121,
		1, 0, 0, 0, 7, 123, 1, 0, 0, 0, 9, 125, 1, 0, 0, 0, 11, 127, 1, 0, 0, 0,
		13, 129, 1, 0, 0, 0, 15, 131, 1, 0, 0, 0, 17, 133, 1, 0, 0, 0, 19, 135,
		1, 0, 0, 0, 21, 138, 1, 0, 0, 0, 23, 141, 1, 0, 0, 0, 25, 144, 1, 0, 0,
		0, 27, 147, 1, 0, 0, 0, 29, 149, 1, 0, 0, 0, 31, 155, 1, 0, 0, 0, 33, 162,
		1, 0, 0, 0, 35, 167, 1, 0, 0, 0, 37, 169, 1, 0, 0, 0, 39, 171, 1, 0, 0,
		0, 41, 173, 1, 0, 0, 0, 43, 175, 1, 0, 0, 0, 45, 181, 1, 0, 0, 0, 47, 183,
		1, 0, 0, 0, 49, 185, 1, 0, 0, 0, 51, 187, 1, 0, 0, 0, 53, 190, 1, 0, 0,
		0, 55, 197, 1, 0, 0, 0, 57, 199, 1, 0, 0, 0, 59, 206, 1, 0, 0, 0, 61, 216,
		1, 0, 0, 0, 63, 218, 1, 0, 0, 0, 65, 225, 1, 0, 0, 0, 67, 230, 1, 0, 0,
		0, 69, 248, 1, 0, 0, 0, 71, 250, 1, 0, 0, 0, 73, 253, 1, 0, 0, 0, 75, 257,
		1, 0, 0, 0, 77, 263, 1, 0, 0, 0, 79, 266, 1, 0, 0, 0, 81, 271, 1, 0, 0,
		0, 83, 275, 1, 0, 0, 0, 85, 281, 1, 0, 0, 0, 87, 287, 1, 0, 0, 0, 89, 296,
		1, 0, 0, 0, 91, 301, 1, 0, 0, 0, 93, 307, 1, 0, 0, 0, 95, 312, 1, 0, 0,
		0, 97, 319, 1, 0, 0, 0, 99, 354, 1, 0, 0, 0, 101, 360, 1, 0, 0, 0, 103,
		367, 1, 0, 0, 0, 105, 369, 1, 0, 0, 0, 107, 371, 1, 0, 0, 0, 109, 381,
		1, 0, 0, 0, 111, 391, 1, 0, 0, 0, 113, 406, 1, 0, 0, 0, 115, 417, 1, 0,
		0, 0, 117, 118, 5, 123, 0, 0, 118, 2, 1, 0, 0, 0, 119, 120, 5, 125, 0,
		0, 120, 4, 1, 0, 0, 0, 121, 122, 5, 40, 0, 0, 122, 6, 1, 0, 0, 0, 123,
		124, 5, 41, 0, 0, 124, 8, 1, 0, 0, 0, 125, 126, 5, 91, 0, 0, 126, 10, 1,
		0, 0, 0, 127, 128, 5, 93, 0, 0, 128, 12, 1, 0, 0, 0, 129, 130, 5, 44, 0,
		0, 130, 14, 1, 0, 0, 0, 131, 132, 5, 58, 0, 0, 132, 16, 1, 0, 0, 0, 133,
		134, 5, 59, 0, 0, 134, 18, 1, 0, 0, 0, 135, 136, 5, 61, 0, 0, 136, 137,
		5, 61, 0, 0, 137, 20, 1, 0, 0, 0, 138, 139, 5, 33, 0, 0, 139, 140, 5, 61,
		0, 0, 140, 22, 1, 0, 0, 0, 141, 142, 5, 62, 0, 0, 142, 143, 5, 61, 0, 0,
		143, 24, 1, 0, 0, 0, 144, 145, 5, 60, 0, 0, 145, 146, 5, 61, 0, 0, 146,
		26, 1, 0, 0, 0, 147, 148, 5, 62, 0, 0, 148, 28, 1, 0, 0, 0, 149, 150, 5,
		60, 0, 0, 150, 30, 1, 0, 0, 0, 151, 152, 5, 124, 0, 0, 152, 156, 5, 124,
		0, 0, 153, 154, 5, 111, 0, 0, 154, 156, 5, 114, 0, 0, 155, 151, 1, 0, 0,
		0, 155, 153, 1, 0, 0, 0, 156, 32, 1, 0, 0, 0, 157, 158, 5, 38, 0, 0, 158,
		163, 5, 38, 0, 0, 159, 160, 5, 97, 0, 0, 160, 161, 5, 110, 0, 0, 161, 163,
		5, 100, 0, 0, 162, 157, 1, 0, 0, 0, 162, 159, 1, 0, 0, 0, 163, 34, 1, 0,
		0, 0, 164, 168, 5, 94, 0, 0, 165, 166, 5, 42, 0, 0, 166, 168, 5, 42, 0,
		0, 167, 164, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 168, 36, 1, 0, 0, 0, 169,
		170, 5, 42, 0, 0, 170, 38, 1, 0, 0, 0, 171, 172, 5, 47, 0, 0, 172, 40,
		1, 0, 0, 0, 173, 174, 5, 43, 0, 0, 174, 42, 1, 0, 0, 0, 175, 176, 5, 45,
		0, 0, 176, 44, 1, 0, 0, 0, 177, 182, 5, 37, 0, 0, 178, 179, 5, 109, 0,
		0, 179, 180, 5, 111, 0, 0, 180, 182, 5, 100, 0, 0, 181, 177, 1, 0, 0, 0,
		181, 178, 1, 0, 0, 0, 182, 46, 1, 0, 0, 0, 183, 184, 5, 46, 0, 0, 184,
		48, 1, 0, 0, 0, 185, 186, 5, 63, 0, 0, 186, 50, 1, 0, 0, 0, 187, 188, 5,
		43, 0, 0, 188, 189, 5, 43, 0, 0, 189, 52, 1, 0, 0, 0, 190, 191, 5, 45,
		0, 0, 191, 192, 5, 45, 0, 0, 192, 54, 1, 0, 0, 0, 193, 194, 5, 45, 0, 0,
		194, 198, 5, 62, 0, 0, 195, 196, 5, 61, 0, 0, 196, 198, 5, 62, 0, 0, 197,
		193, 1, 0, 0, 0, 197, 195, 1, 0, 0, 0, 198, 56, 1, 0, 0, 0, 199, 200, 5,
		115, 0, 0, 200, 201, 5, 116, 0, 0, 201, 202, 5, 114, 0, 0, 202, 203, 5,
		117, 0, 0, 203, 204, 5, 99, 0, 0, 204, 205, 5, 116, 0, 0, 205, 58, 1, 0,
		0, 0, 206, 207, 5, 109, 0, 0, 207, 208, 5, 97, 0, 0, 208, 209, 5, 112,
		0, 0, 209, 60, 1, 0, 0, 0, 210, 211, 5, 102, 0, 0, 211, 217, 5, 110, 0,
		0, 212, 213, 5, 102, 0, 0, 213, 214, 5, 117, 0, 0, 214, 215, 5, 110, 0,
		0, 215, 217, 5, 99, 0, 0, 216, 210, 1, 0, 0, 0, 216, 212, 1, 0, 0, 0, 217,
		62, 1, 0, 0, 0, 218, 219, 5, 114, 0, 0, 219, 220, 5, 101, 0, 0, 220, 221,
		5, 116, 0, 0, 221, 222, 5, 117, 0, 0, 222, 223, 5, 114, 0, 0, 223, 224,
		5, 110, 0, 0, 224, 64, 1, 0, 0, 0, 225, 226, 5, 99, 0, 0, 226, 227, 5,
		97, 0, 0, 227, 228, 5, 115, 0, 0, 228, 229, 5, 101, 0, 0, 229, 66, 1, 0,
		0, 0, 230, 231, 5, 100, 0, 0, 231, 232, 5, 101, 0, 0, 232, 233, 5, 102,
		0, 0, 233, 234, 5, 97, 0, 0, 234, 235, 5, 117, 0, 0, 235, 236, 5, 108,
		0, 0, 236, 237, 5, 116, 0, 0, 237, 68, 1, 0, 0, 0, 238, 239, 5, 111, 0,
		0, 239, 240, 5, 112, 0, 0, 240, 241, 5, 101, 0, 0, 241, 249, 5, 110, 0,
		0, 242, 243, 5, 105, 0, 0, 243, 244, 5, 109, 0, 0, 244, 245, 5, 112, 0,
		0, 245, 246, 5, 111, 0, 0, 246, 247, 5, 114, 0, 0, 247, 249, 5, 116, 0,
		0, 248, 238, 1, 0, 0, 0, 248, 242, 1, 0, 0, 0, 249, 70, 1, 0, 0, 0, 250,
		251, 5, 97, 0, 0, 251, 252, 5, 115, 0, 0, 252, 72, 1, 0, 0, 0, 253, 254,
		5, 116, 0, 0, 254, 255, 5, 114, 0, 0, 255, 256, 5, 121, 0, 0, 256, 74,
		1, 0, 0, 0, 257, 258, 5, 99, 0, 0, 258, 259, 5, 97, 0, 0, 259, 260, 5,
		116, 0, 0, 260, 261, 5, 99, 0, 0, 261, 262, 5, 104, 0, 0, 262, 76, 1, 0,
		0, 0, 263, 264, 5, 105, 0, 0, 264, 265, 5, 102, 0, 0, 265, 78, 1, 0, 0,
		0, 266, 267, 5, 101, 0, 0, 267, 268, 5, 108, 0, 0, 268, 269, 5, 115, 0,
		0, 269, 270, 5, 101, 0, 0, 270, 80, 1, 0, 0, 0, 271, 272, 5, 102, 0, 0,
		272, 273, 5, 111, 0, 0, 273, 274, 5, 114, 0, 0, 274, 82, 1, 0, 0, 0, 275,
		276, 5, 109, 0, 0, 276, 277, 5, 97, 0, 0, 277, 278, 5, 116, 0, 0, 278,
		279, 5, 99, 0, 0, 279, 280, 5, 104, 0, 0, 280, 84, 1, 0, 0, 0, 281, 282,
		5, 98, 0, 0, 282, 283, 5, 114, 0, 0, 283, 284, 5, 101, 0, 0, 284, 285,
		5, 97, 0, 0, 285, 286, 5, 107, 0, 0, 286, 86, 1, 0, 0, 0, 287, 288, 5,
		99, 0, 0, 288, 289, 5, 111, 0, 0, 289, 290, 5, 110, 0, 0, 290, 291, 5,
		116, 0, 0, 291, 292, 5, 105, 0, 0, 292, 293, 5, 110, 0, 0, 293, 294, 5,
		117, 0, 0, 294, 295, 5, 101, 0, 0, 295, 88, 1, 0, 0, 0, 296, 297, 5, 116,
		0, 0, 297, 298, 5, 114, 0, 0, 298, 299, 5, 117, 0, 0, 299, 300, 5, 101,
		0, 0, 300, 90, 1, 0, 0, 0, 301, 302, 5, 102, 0, 0, 302, 303, 5, 97, 0,
		0, 303, 304, 5, 108, 0, 0, 304, 305, 5, 115, 0, 0, 305, 306, 5, 101, 0,
		0, 306, 92, 1, 0, 0, 0, 307, 308, 5, 110, 0, 0, 308, 309, 5, 105, 0, 0,
		309, 310, 5, 108, 0, 0, 310, 94, 1, 0, 0, 0, 311, 313, 7, 0, 0, 0, 312,
		311, 1, 0, 0, 0, 313, 314, 1, 0, 0, 0, 314, 312, 1, 0, 0, 0, 314, 315,
		1, 0, 0, 0, 315, 96, 1, 0, 0, 0, 316, 318, 7, 0, 0, 0, 317, 316, 1, 0,
		0, 0, 318, 321, 1, 0, 0, 0, 319, 317, 1, 0, 0, 0, 319, 320, 1, 0, 0, 0,
		320, 322, 1, 0, 0, 0, 321, 319, 1, 0, 0, 0, 322, 324, 5, 46, 0, 0, 323,
		325, 7, 0, 0, 0, 324, 323, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 324,
		1, 0, 0, 0, 326, 327, 1, 0, 0, 0, 327, 98, 1, 0, 0, 0, 328, 330, 7, 1,
		0, 0, 329, 328, 1, 0, 0, 0, 329, 330, 1, 0, 0, 0, 330, 331, 1, 0, 0, 0,
		331, 337, 5, 34, 0, 0, 332, 336, 8, 2, 0, 0, 333, 334, 5, 92, 0, 0, 334,
		336, 5, 34, 0, 0, 335, 332, 1, 0, 0, 0, 335, 333, 1, 0, 0, 0, 336, 339,
		1, 0, 0, 0, 337, 335, 1, 0, 0, 0, 337, 338, 1, 0, 0, 0, 338, 340, 1, 0,
		0, 0, 339, 337, 1, 0, 0, 0, 340, 355, 5, 34, 0, 0, 341, 343, 7, 1, 0, 0,
		342, 341, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343, 344, 1, 0, 0, 0, 344,
		350, 5, 39, 0, 0, 345, 349, 8, 3, 0, 0, 346, 347, 5, 92, 0, 0, 347, 349,
		5, 39, 0, 0, 348, 345, 1, 0, 0, 0, 348, 346, 1, 0, 0, 0, 349, 352, 1, 0,
		0, 0, 350, 348, 1, 0, 0, 0, 350, 351, 1, 0, 0, 0, 351, 353, 1, 0, 0, 0,
		352, 350, 1, 0, 0, 0, 353, 355, 5, 39, 0, 0, 354, 329, 1, 0, 0, 0, 354,
		342, 1, 0, 0, 0, 355, 100, 1, 0, 0, 0, 356, 361, 5, 33, 0, 0, 357, 358,
		5, 110, 0, 0, 358, 359, 5, 111, 0, 0, 359, 361, 5, 116, 0, 0, 360, 356,
		1, 0, 0, 0, 360, 357, 1, 0, 0, 0, 361, 102, 1, 0, 0, 0, 362, 363, 5, 58,
		0, 0, 363, 368, 5, 61, 0, 0, 364, 368, 5, 61, 0, 0, 365, 366, 5, 60, 0,
		0, 366, 368, 5, 45, 0, 0, 367, 362, 1, 0, 0, 0, 367, 364, 1, 0, 0, 0, 367,
		365, 1, 0, 0, 0, 368, 104, 1, 0, 0, 0, 369, 370, 5, 38, 0, 0, 370, 106,
		1, 0, 0, 0, 371, 375, 7, 4, 0, 0, 372, 374, 7, 5, 0, 0, 373, 372, 1, 0,
		0, 0, 374, 377, 1, 0, 0, 0, 375, 373, 1, 0, 0, 0, 375, 376, 1, 0, 0, 0,
		376, 108, 1, 0, 0, 0, 377, 375, 1, 0, 0, 0, 378, 382, 5, 35, 0, 0, 379,
		380, 5, 47, 0, 0, 380, 382, 5, 47, 0, 0, 381, 378, 1, 0, 0, 0, 381, 379,
		1, 0, 0, 0, 382, 386, 1, 0, 0, 0, 383, 385, 8, 6, 0, 0, 384, 383, 1, 0,
		0, 0, 385, 388, 1, 0, 0, 0, 386, 384, 1, 0, 0, 0, 386, 387, 1, 0, 0, 0,
		387, 389, 1, 0, 0, 0, 388, 386, 1, 0, 0, 0, 389, 390, 6, 54, 0, 0, 390,
		110, 1, 0, 0, 0, 391, 392, 5, 47, 0, 0, 392, 393, 5, 42, 0, 0, 393, 397,
		1, 0, 0, 0, 394, 396, 9, 0, 0, 0, 395, 394, 1, 0, 0, 0, 396, 399, 1, 0,
		0, 0, 397, 398, 1, 0, 0, 0, 397, 395, 1, 0, 0, 0, 398, 400, 1, 0, 0, 0,
		399, 397, 1, 0, 0, 0, 400, 401, 5, 42, 0, 0, 401, 402, 5, 47, 0, 0, 402,
		403, 1, 0, 0, 0, 403, 404, 6, 55, 0, 0, 404, 112, 1, 0, 0, 0, 405, 407,
		7, 7, 0, 0, 406, 405, 1, 0, 0, 0, 407, 408, 1, 0, 0, 0, 408, 406, 1, 0,
		0, 0, 408, 409, 1, 0, 0, 0, 409, 410, 1, 0, 0, 0, 410, 411, 6, 56, 1, 0,
		411, 114, 1, 0, 0, 0, 412, 414, 5, 13, 0, 0, 413, 412, 1, 0, 0, 0, 413,
		414, 1, 0, 0, 0, 414, 415, 1, 0, 0, 0, 415, 418, 5, 10, 0, 0, 416, 418,
		5, 13, 0, 0, 417, 413, 1, 0, 0, 0, 417, 416, 1, 0, 0, 0, 418, 419, 1, 0,
		0, 0, 419, 417, 1, 0, 0, 0, 419, 420, 1, 0, 0, 0, 420, 421, 1, 0, 0, 0,
		421, 422, 6, 57, 2, 0, 422, 116, 1, 0, 0, 0, 28, 0, 155, 162, 167, 181,
		197, 216, 248, 314, 319, 326, 329, 335, 337, 342, 348, 350, 354, 360, 367,
		375, 381, 386, 397, 408, 413, 417, 419, 3, 0, 2, 0, 6, 0, 0, 0, 1, 0,
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
	V2LexerAddAdd         = 26
	V2LexerSubSub         = 27
	V2LexerTo             = 28
	V2LexerStruct         = 29
	V2LexerMap            = 30
	V2LexerFunction       = 31
	V2LexerReturn         = 32
	V2LexerCase           = 33
	V2LexerDefault        = 34
	V2LexerOpen           = 35
	V2LexerAs             = 36
	V2LexerTry            = 37
	V2LexerCatch          = 38
	V2LexerIf             = 39
	V2LexerElse           = 40
	V2LexerFor            = 41
	V2LexerMatch          = 42
	V2LexerBreak          = 43
	V2LexerContinue       = 44
	V2LexerTrue           = 45
	V2LexerFalse          = 46
	V2LexerNil            = 47
	V2LexerIntegerLiteral = 48
	V2LexerNumberLiteral  = 49
	V2LexerStringLiteral  = 50
	V2LexerNot            = 51
	V2LexerAssign         = 52
	V2LexerRef            = 53
	V2LexerIdentifier     = 54
	V2LexerComment        = 55
	V2LexerBlkComment     = 56
	V2LexerWS             = 57
	V2LexerNewLine        = 58
)
