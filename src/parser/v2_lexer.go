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
		"'-'", "", "'.'", "'->'", "'struct'", "'map'", "'fn'", "'return'", "'case'",
		"'default'", "'open'", "'as'", "'try'", "'catch'", "'if'", "'else'",
		"'for'", "'match'", "'break'", "", "'true'", "'false'", "", "", "",
		"", "", "'&'",
	}
	staticData.SymbolicNames = []string{
		"", "LBrack", "RBrack", "LParen", "RParen", "LSquare", "RSquare", "Comma",
		"Col", "Semi", "Equals", "NotEq", "GreaterEq", "LessEq", "Greater",
		"Less", "Or", "And", "Pow", "Mul", "Div", "Add", "Sub", "Mod", "Dot",
		"To", "Struct", "Map", "Function", "Return", "Case", "Default", "Open",
		"As", "Try", "Catch", "If", "Else", "For", "Match", "Break", "Continue",
		"True", "False", "IntegerLiteral", "NumberLiteral", "StringLiteral",
		"Not", "Assign", "Ref", "Identifier", "Comment", "WS", "NewLine",
	}
	staticData.RuleNames = []string{
		"LBrack", "RBrack", "LParen", "RParen", "LSquare", "RSquare", "Comma",
		"Col", "Semi", "Equals", "NotEq", "GreaterEq", "LessEq", "Greater",
		"Less", "Or", "And", "Pow", "Mul", "Div", "Add", "Sub", "Mod", "Dot",
		"To", "Struct", "Map", "Function", "Return", "Case", "Default", "Open",
		"As", "Try", "Catch", "If", "Else", "For", "Match", "Break", "Continue",
		"True", "False", "IntegerLiteral", "NumberLiteral", "StringLiteral",
		"Not", "Assign", "Ref", "Identifier", "Comment", "WS", "NewLine",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 53, 360, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 146, 8, 15, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 3, 16, 153, 8, 16, 1, 17, 1, 17, 1, 17, 3, 17, 158,
		8, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1,
		22, 1, 22, 1, 22, 3, 22, 172, 8, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40,
		1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 3, 40, 271, 8,
		40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42,
		1, 42, 1, 43, 4, 43, 285, 8, 43, 11, 43, 12, 43, 286, 1, 44, 4, 44, 290,
		8, 44, 11, 44, 12, 44, 291, 1, 44, 1, 44, 4, 44, 296, 8, 44, 11, 44, 12,
		44, 297, 3, 44, 300, 8, 44, 1, 45, 3, 45, 303, 8, 45, 1, 45, 1, 45, 1,
		45, 1, 45, 5, 45, 309, 8, 45, 10, 45, 12, 45, 312, 9, 45, 1, 45, 1, 45,
		1, 46, 1, 46, 1, 46, 1, 46, 3, 46, 320, 8, 46, 1, 47, 1, 47, 1, 47, 1,
		47, 1, 47, 3, 47, 327, 8, 47, 1, 48, 1, 48, 1, 49, 1, 49, 5, 49, 333, 8,
		49, 10, 49, 12, 49, 336, 9, 49, 1, 50, 1, 50, 5, 50, 340, 8, 50, 10, 50,
		12, 50, 343, 9, 50, 1, 50, 1, 50, 1, 51, 4, 51, 348, 8, 51, 11, 51, 12,
		51, 349, 1, 51, 1, 51, 1, 52, 3, 52, 355, 8, 52, 1, 52, 1, 52, 3, 52, 359,
		8, 52, 0, 0, 53, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17,
		9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35,
		18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53,
		27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71,
		36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89,
		45, 91, 46, 93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53,
		1, 0, 7, 1, 0, 48, 57, 2, 0, 36, 36, 64, 64, 1, 0, 34, 34, 3, 0, 65, 90,
		95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 10, 10, 13,
		13, 3, 0, 9, 10, 12, 13, 32, 32, 380, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1,
		0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35,
		1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0,
		43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0,
		0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0,
		0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0,
		0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1,
		0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81,
		1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0,
		89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0,
		0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0,
		0, 0, 0, 105, 1, 0, 0, 0, 1, 107, 1, 0, 0, 0, 3, 109, 1, 0, 0, 0, 5, 111,
		1, 0, 0, 0, 7, 113, 1, 0, 0, 0, 9, 115, 1, 0, 0, 0, 11, 117, 1, 0, 0, 0,
		13, 119, 1, 0, 0, 0, 15, 121, 1, 0, 0, 0, 17, 123, 1, 0, 0, 0, 19, 125,
		1, 0, 0, 0, 21, 128, 1, 0, 0, 0, 23, 131, 1, 0, 0, 0, 25, 134, 1, 0, 0,
		0, 27, 137, 1, 0, 0, 0, 29, 139, 1, 0, 0, 0, 31, 145, 1, 0, 0, 0, 33, 152,
		1, 0, 0, 0, 35, 157, 1, 0, 0, 0, 37, 159, 1, 0, 0, 0, 39, 161, 1, 0, 0,
		0, 41, 163, 1, 0, 0, 0, 43, 165, 1, 0, 0, 0, 45, 171, 1, 0, 0, 0, 47, 173,
		1, 0, 0, 0, 49, 175, 1, 0, 0, 0, 51, 178, 1, 0, 0, 0, 53, 185, 1, 0, 0,
		0, 55, 189, 1, 0, 0, 0, 57, 192, 1, 0, 0, 0, 59, 199, 1, 0, 0, 0, 61, 204,
		1, 0, 0, 0, 63, 212, 1, 0, 0, 0, 65, 217, 1, 0, 0, 0, 67, 220, 1, 0, 0,
		0, 69, 224, 1, 0, 0, 0, 71, 230, 1, 0, 0, 0, 73, 233, 1, 0, 0, 0, 75, 238,
		1, 0, 0, 0, 77, 242, 1, 0, 0, 0, 79, 248, 1, 0, 0, 0, 81, 270, 1, 0, 0,
		0, 83, 272, 1, 0, 0, 0, 85, 277, 1, 0, 0, 0, 87, 284, 1, 0, 0, 0, 89, 289,
		1, 0, 0, 0, 91, 302, 1, 0, 0, 0, 93, 319, 1, 0, 0, 0, 95, 326, 1, 0, 0,
		0, 97, 328, 1, 0, 0, 0, 99, 330, 1, 0, 0, 0, 101, 337, 1, 0, 0, 0, 103,
		347, 1, 0, 0, 0, 105, 358, 1, 0, 0, 0, 107, 108, 5, 123, 0, 0, 108, 2,
		1, 0, 0, 0, 109, 110, 5, 125, 0, 0, 110, 4, 1, 0, 0, 0, 111, 112, 5, 40,
		0, 0, 112, 6, 1, 0, 0, 0, 113, 114, 5, 41, 0, 0, 114, 8, 1, 0, 0, 0, 115,
		116, 5, 91, 0, 0, 116, 10, 1, 0, 0, 0, 117, 118, 5, 93, 0, 0, 118, 12,
		1, 0, 0, 0, 119, 120, 5, 44, 0, 0, 120, 14, 1, 0, 0, 0, 121, 122, 5, 58,
		0, 0, 122, 16, 1, 0, 0, 0, 123, 124, 5, 59, 0, 0, 124, 18, 1, 0, 0, 0,
		125, 126, 5, 61, 0, 0, 126, 127, 5, 61, 0, 0, 127, 20, 1, 0, 0, 0, 128,
		129, 5, 33, 0, 0, 129, 130, 5, 61, 0, 0, 130, 22, 1, 0, 0, 0, 131, 132,
		5, 62, 0, 0, 132, 133, 5, 61, 0, 0, 133, 24, 1, 0, 0, 0, 134, 135, 5, 60,
		0, 0, 135, 136, 5, 61, 0, 0, 136, 26, 1, 0, 0, 0, 137, 138, 5, 62, 0, 0,
		138, 28, 1, 0, 0, 0, 139, 140, 5, 60, 0, 0, 140, 30, 1, 0, 0, 0, 141, 142,
		5, 124, 0, 0, 142, 146, 5, 124, 0, 0, 143, 144, 5, 111, 0, 0, 144, 146,
		5, 114, 0, 0, 145, 141, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 146, 32, 1, 0,
		0, 0, 147, 148, 5, 38, 0, 0, 148, 153, 5, 38, 0, 0, 149, 150, 5, 97, 0,
		0, 150, 151, 5, 110, 0, 0, 151, 153, 5, 100, 0, 0, 152, 147, 1, 0, 0, 0,
		152, 149, 1, 0, 0, 0, 153, 34, 1, 0, 0, 0, 154, 158, 5, 94, 0, 0, 155,
		156, 5, 42, 0, 0, 156, 158, 5, 42, 0, 0, 157, 154, 1, 0, 0, 0, 157, 155,
		1, 0, 0, 0, 158, 36, 1, 0, 0, 0, 159, 160, 5, 42, 0, 0, 160, 38, 1, 0,
		0, 0, 161, 162, 5, 47, 0, 0, 162, 40, 1, 0, 0, 0, 163, 164, 5, 43, 0, 0,
		164, 42, 1, 0, 0, 0, 165, 166, 5, 45, 0, 0, 166, 44, 1, 0, 0, 0, 167, 172,
		5, 37, 0, 0, 168, 169, 5, 109, 0, 0, 169, 170, 5, 111, 0, 0, 170, 172,
		5, 100, 0, 0, 171, 167, 1, 0, 0, 0, 171, 168, 1, 0, 0, 0, 172, 46, 1, 0,
		0, 0, 173, 174, 5, 46, 0, 0, 174, 48, 1, 0, 0, 0, 175, 176, 5, 45, 0, 0,
		176, 177, 5, 62, 0, 0, 177, 50, 1, 0, 0, 0, 178, 179, 5, 115, 0, 0, 179,
		180, 5, 116, 0, 0, 180, 181, 5, 114, 0, 0, 181, 182, 5, 117, 0, 0, 182,
		183, 5, 99, 0, 0, 183, 184, 5, 116, 0, 0, 184, 52, 1, 0, 0, 0, 185, 186,
		5, 109, 0, 0, 186, 187, 5, 97, 0, 0, 187, 188, 5, 112, 0, 0, 188, 54, 1,
		0, 0, 0, 189, 190, 5, 102, 0, 0, 190, 191, 5, 110, 0, 0, 191, 56, 1, 0,
		0, 0, 192, 193, 5, 114, 0, 0, 193, 194, 5, 101, 0, 0, 194, 195, 5, 116,
		0, 0, 195, 196, 5, 117, 0, 0, 196, 197, 5, 114, 0, 0, 197, 198, 5, 110,
		0, 0, 198, 58, 1, 0, 0, 0, 199, 200, 5, 99, 0, 0, 200, 201, 5, 97, 0, 0,
		201, 202, 5, 115, 0, 0, 202, 203, 5, 101, 0, 0, 203, 60, 1, 0, 0, 0, 204,
		205, 5, 100, 0, 0, 205, 206, 5, 101, 0, 0, 206, 207, 5, 102, 0, 0, 207,
		208, 5, 97, 0, 0, 208, 209, 5, 117, 0, 0, 209, 210, 5, 108, 0, 0, 210,
		211, 5, 116, 0, 0, 211, 62, 1, 0, 0, 0, 212, 213, 5, 111, 0, 0, 213, 214,
		5, 112, 0, 0, 214, 215, 5, 101, 0, 0, 215, 216, 5, 110, 0, 0, 216, 64,
		1, 0, 0, 0, 217, 218, 5, 97, 0, 0, 218, 219, 5, 115, 0, 0, 219, 66, 1,
		0, 0, 0, 220, 221, 5, 116, 0, 0, 221, 222, 5, 114, 0, 0, 222, 223, 5, 121,
		0, 0, 223, 68, 1, 0, 0, 0, 224, 225, 5, 99, 0, 0, 225, 226, 5, 97, 0, 0,
		226, 227, 5, 116, 0, 0, 227, 228, 5, 99, 0, 0, 228, 229, 5, 104, 0, 0,
		229, 70, 1, 0, 0, 0, 230, 231, 5, 105, 0, 0, 231, 232, 5, 102, 0, 0, 232,
		72, 1, 0, 0, 0, 233, 234, 5, 101, 0, 0, 234, 235, 5, 108, 0, 0, 235, 236,
		5, 115, 0, 0, 236, 237, 5, 101, 0, 0, 237, 74, 1, 0, 0, 0, 238, 239, 5,
		102, 0, 0, 239, 240, 5, 111, 0, 0, 240, 241, 5, 114, 0, 0, 241, 76, 1,
		0, 0, 0, 242, 243, 5, 109, 0, 0, 243, 244, 5, 97, 0, 0, 244, 245, 5, 116,
		0, 0, 245, 246, 5, 99, 0, 0, 246, 247, 5, 104, 0, 0, 247, 78, 1, 0, 0,
		0, 248, 249, 5, 98, 0, 0, 249, 250, 5, 114, 0, 0, 250, 251, 5, 101, 0,
		0, 251, 252, 5, 97, 0, 0, 252, 253, 5, 107, 0, 0, 253, 80, 1, 0, 0, 0,
		254, 255, 5, 110, 0, 0, 255, 256, 5, 101, 0, 0, 256, 257, 5, 120, 0, 0,
		257, 271, 5, 116, 0, 0, 258, 259, 5, 112, 0, 0, 259, 260, 5, 97, 0, 0,
		260, 261, 5, 115, 0, 0, 261, 271, 5, 115, 0, 0, 262, 263, 5, 99, 0, 0,
		263, 264, 5, 111, 0, 0, 264, 265, 5, 110, 0, 0, 265, 266, 5, 116, 0, 0,
		266, 267, 5, 105, 0, 0, 267, 268, 5, 110, 0, 0, 268, 269, 5, 117, 0, 0,
		269, 271, 5, 101, 0, 0, 270, 254, 1, 0, 0, 0, 270, 258, 1, 0, 0, 0, 270,
		262, 1, 0, 0, 0, 271, 82, 1, 0, 0, 0, 272, 273, 5, 116, 0, 0, 273, 274,
		5, 114, 0, 0, 274, 275, 5, 117, 0, 0, 275, 276, 5, 101, 0, 0, 276, 84,
		1, 0, 0, 0, 277, 278, 5, 102, 0, 0, 278, 279, 5, 97, 0, 0, 279, 280, 5,
		108, 0, 0, 280, 281, 5, 115, 0, 0, 281, 282, 5, 101, 0, 0, 282, 86, 1,
		0, 0, 0, 283, 285, 7, 0, 0, 0, 284, 283, 1, 0, 0, 0, 285, 286, 1, 0, 0,
		0, 286, 284, 1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 88, 1, 0, 0, 0, 288,
		290, 7, 0, 0, 0, 289, 288, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 289,
		1, 0, 0, 0, 291, 292, 1, 0, 0, 0, 292, 299, 1, 0, 0, 0, 293, 295, 5, 46,
		0, 0, 294, 296, 7, 0, 0, 0, 295, 294, 1, 0, 0, 0, 296, 297, 1, 0, 0, 0,
		297, 295, 1, 0, 0, 0, 297, 298, 1, 0, 0, 0, 298, 300, 1, 0, 0, 0, 299,
		293, 1, 0, 0, 0, 299, 300, 1, 0, 0, 0, 300, 90, 1, 0, 0, 0, 301, 303, 7,
		1, 0, 0, 302, 301, 1, 0, 0, 0, 302, 303, 1, 0, 0, 0, 303, 304, 1, 0, 0,
		0, 304, 310, 5, 34, 0, 0, 305, 309, 8, 2, 0, 0, 306, 307, 5, 92, 0, 0,
		307, 309, 5, 34, 0, 0, 308, 305, 1, 0, 0, 0, 308, 306, 1, 0, 0, 0, 309,
		312, 1, 0, 0, 0, 310, 308, 1, 0, 0, 0, 310, 311, 1, 0, 0, 0, 311, 313,
		1, 0, 0, 0, 312, 310, 1, 0, 0, 0, 313, 314, 5, 34, 0, 0, 314, 92, 1, 0,
		0, 0, 315, 320, 5, 33, 0, 0, 316, 317, 5, 110, 0, 0, 317, 318, 5, 111,
		0, 0, 318, 320, 5, 116, 0, 0, 319, 315, 1, 0, 0, 0, 319, 316, 1, 0, 0,
		0, 320, 94, 1, 0, 0, 0, 321, 322, 5, 58, 0, 0, 322, 327, 5, 61, 0, 0, 323,
		327, 5, 61, 0, 0, 324, 325, 5, 60, 0, 0, 325, 327, 5, 45, 0, 0, 326, 321,
		1, 0, 0, 0, 326, 323, 1, 0, 0, 0, 326, 324, 1, 0, 0, 0, 327, 96, 1, 0,
		0, 0, 328, 329, 5, 38, 0, 0, 329, 98, 1, 0, 0, 0, 330, 334, 7, 3, 0, 0,
		331, 333, 7, 4, 0, 0, 332, 331, 1, 0, 0, 0, 333, 336, 1, 0, 0, 0, 334,
		332, 1, 0, 0, 0, 334, 335, 1, 0, 0, 0, 335, 100, 1, 0, 0, 0, 336, 334,
		1, 0, 0, 0, 337, 341, 5, 35, 0, 0, 338, 340, 8, 5, 0, 0, 339, 338, 1, 0,
		0, 0, 340, 343, 1, 0, 0, 0, 341, 339, 1, 0, 0, 0, 341, 342, 1, 0, 0, 0,
		342, 344, 1, 0, 0, 0, 343, 341, 1, 0, 0, 0, 344, 345, 6, 50, 0, 0, 345,
		102, 1, 0, 0, 0, 346, 348, 7, 6, 0, 0, 347, 346, 1, 0, 0, 0, 348, 349,
		1, 0, 0, 0, 349, 347, 1, 0, 0, 0, 349, 350, 1, 0, 0, 0, 350, 351, 1, 0,
		0, 0, 351, 352, 6, 51, 1, 0, 352, 104, 1, 0, 0, 0, 353, 355, 5, 13, 0,
		0, 354, 353, 1, 0, 0, 0, 354, 355, 1, 0, 0, 0, 355, 356, 1, 0, 0, 0, 356,
		359, 5, 10, 0, 0, 357, 359, 5, 13, 0, 0, 358, 354, 1, 0, 0, 0, 358, 357,
		1, 0, 0, 0, 359, 106, 1, 0, 0, 0, 20, 0, 145, 152, 157, 171, 270, 286,
		291, 297, 299, 302, 308, 310, 319, 326, 334, 341, 349, 354, 358, 2, 0,
		2, 0, 6, 0, 0,
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
	V2LexerTo             = 25
	V2LexerStruct         = 26
	V2LexerMap            = 27
	V2LexerFunction       = 28
	V2LexerReturn         = 29
	V2LexerCase           = 30
	V2LexerDefault        = 31
	V2LexerOpen           = 32
	V2LexerAs             = 33
	V2LexerTry            = 34
	V2LexerCatch          = 35
	V2LexerIf             = 36
	V2LexerElse           = 37
	V2LexerFor            = 38
	V2LexerMatch          = 39
	V2LexerBreak          = 40
	V2LexerContinue       = 41
	V2LexerTrue           = 42
	V2LexerFalse          = 43
	V2LexerIntegerLiteral = 44
	V2LexerNumberLiteral  = 45
	V2LexerStringLiteral  = 46
	V2LexerNot            = 47
	V2LexerAssign         = 48
	V2LexerRef            = 49
	V2LexerIdentifier     = 50
	V2LexerComment        = 51
	V2LexerWS             = 52
	V2LexerNewLine        = 53
)
