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
		"'-'", "", "'.'", "'->'", "'struct'", "'map'", "'fn'", "'return'", "'case'",
		"'default'", "'open'", "'as'", "'if'", "'else'", "'for'", "'match'",
		"'break'", "", "'true'", "'false'", "", "", "", "", "':='",
	}
	staticData.SymbolicNames = []string{
		"", "LBrack", "RBrack", "LParen", "RParen", "LSquare", "RSquare", "Comma",
		"Col", "Semi", "Equals", "NotEq", "GreaterEq", "LessEq", "Greater",
		"Less", "Or", "And", "Pow", "Mul", "Div", "Add", "Sub", "Mod", "Dot",
		"To", "Struct", "Map", "Function", "Return", "Case", "Default", "Open",
		"As", "If", "Else", "For", "Match", "Break", "Continue", "True", "False",
		"IntegerLiteral", "NumberLiteral", "StringLiteral", "Not", "Assign",
		"Identifier", "Comment", "WS", "NewLine",
	}
	staticData.RuleNames = []string{
		"LBrack", "RBrack", "LParen", "RParen", "LSquare", "RSquare", "Comma",
		"Col", "Semi", "Equals", "NotEq", "GreaterEq", "LessEq", "Greater",
		"Less", "Or", "And", "Pow", "Mul", "Div", "Add", "Sub", "Mod", "Dot",
		"To", "Struct", "Map", "Function", "Return", "Case", "Default", "Open",
		"As", "If", "Else", "For", "Match", "Break", "Continue", "True", "False",
		"IntegerLiteral", "NumberLiteral", "StringLiteral", "Not", "Assign",
		"Identifier", "Comment", "WS", "NewLine",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 50, 333, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 3,
		15, 140, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 147, 8, 16, 1,
		17, 1, 17, 1, 17, 3, 17, 152, 8, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20,
		1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 166, 8, 22, 1,
		23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38,
		1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1,
		38, 1, 38, 1, 38, 3, 38, 255, 8, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39,
		1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 4, 41, 269, 8, 41, 11,
		41, 12, 41, 270, 1, 42, 4, 42, 274, 8, 42, 11, 42, 12, 42, 275, 1, 42,
		1, 42, 4, 42, 280, 8, 42, 11, 42, 12, 42, 281, 3, 42, 284, 8, 42, 1, 43,
		1, 43, 5, 43, 288, 8, 43, 10, 43, 12, 43, 291, 9, 43, 1, 43, 1, 43, 1,
		44, 1, 44, 1, 44, 1, 44, 3, 44, 299, 8, 44, 1, 45, 1, 45, 1, 45, 1, 46,
		1, 46, 5, 46, 306, 8, 46, 10, 46, 12, 46, 309, 9, 46, 1, 47, 1, 47, 5,
		47, 313, 8, 47, 10, 47, 12, 47, 316, 9, 47, 1, 47, 1, 47, 1, 48, 4, 48,
		321, 8, 48, 11, 48, 12, 48, 322, 1, 48, 1, 48, 1, 49, 3, 49, 328, 8, 49,
		1, 49, 1, 49, 3, 49, 332, 8, 49, 1, 289, 0, 50, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32,
		65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41,
		83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 95, 48, 97, 49, 99, 50,
		1, 0, 5, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65,
		90, 95, 95, 97, 122, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 12, 13, 32, 32,
		349, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0,
		0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1,
		0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23,
		1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0,
		31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0,
		0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0,
		0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0,
		0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1,
		0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69,
		1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0,
		77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0,
		0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0,
		0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0,
		0, 0, 1, 101, 1, 0, 0, 0, 3, 103, 1, 0, 0, 0, 5, 105, 1, 0, 0, 0, 7, 107,
		1, 0, 0, 0, 9, 109, 1, 0, 0, 0, 11, 111, 1, 0, 0, 0, 13, 113, 1, 0, 0,
		0, 15, 115, 1, 0, 0, 0, 17, 117, 1, 0, 0, 0, 19, 119, 1, 0, 0, 0, 21, 122,
		1, 0, 0, 0, 23, 125, 1, 0, 0, 0, 25, 128, 1, 0, 0, 0, 27, 131, 1, 0, 0,
		0, 29, 133, 1, 0, 0, 0, 31, 139, 1, 0, 0, 0, 33, 146, 1, 0, 0, 0, 35, 151,
		1, 0, 0, 0, 37, 153, 1, 0, 0, 0, 39, 155, 1, 0, 0, 0, 41, 157, 1, 0, 0,
		0, 43, 159, 1, 0, 0, 0, 45, 165, 1, 0, 0, 0, 47, 167, 1, 0, 0, 0, 49, 169,
		1, 0, 0, 0, 51, 172, 1, 0, 0, 0, 53, 179, 1, 0, 0, 0, 55, 183, 1, 0, 0,
		0, 57, 186, 1, 0, 0, 0, 59, 193, 1, 0, 0, 0, 61, 198, 1, 0, 0, 0, 63, 206,
		1, 0, 0, 0, 65, 211, 1, 0, 0, 0, 67, 214, 1, 0, 0, 0, 69, 217, 1, 0, 0,
		0, 71, 222, 1, 0, 0, 0, 73, 226, 1, 0, 0, 0, 75, 232, 1, 0, 0, 0, 77, 254,
		1, 0, 0, 0, 79, 256, 1, 0, 0, 0, 81, 261, 1, 0, 0, 0, 83, 268, 1, 0, 0,
		0, 85, 273, 1, 0, 0, 0, 87, 285, 1, 0, 0, 0, 89, 298, 1, 0, 0, 0, 91, 300,
		1, 0, 0, 0, 93, 303, 1, 0, 0, 0, 95, 310, 1, 0, 0, 0, 97, 320, 1, 0, 0,
		0, 99, 331, 1, 0, 0, 0, 101, 102, 5, 123, 0, 0, 102, 2, 1, 0, 0, 0, 103,
		104, 5, 125, 0, 0, 104, 4, 1, 0, 0, 0, 105, 106, 5, 40, 0, 0, 106, 6, 1,
		0, 0, 0, 107, 108, 5, 41, 0, 0, 108, 8, 1, 0, 0, 0, 109, 110, 5, 91, 0,
		0, 110, 10, 1, 0, 0, 0, 111, 112, 5, 93, 0, 0, 112, 12, 1, 0, 0, 0, 113,
		114, 5, 44, 0, 0, 114, 14, 1, 0, 0, 0, 115, 116, 5, 58, 0, 0, 116, 16,
		1, 0, 0, 0, 117, 118, 5, 59, 0, 0, 118, 18, 1, 0, 0, 0, 119, 120, 5, 61,
		0, 0, 120, 121, 5, 61, 0, 0, 121, 20, 1, 0, 0, 0, 122, 123, 5, 33, 0, 0,
		123, 124, 5, 61, 0, 0, 124, 22, 1, 0, 0, 0, 125, 126, 5, 62, 0, 0, 126,
		127, 5, 61, 0, 0, 127, 24, 1, 0, 0, 0, 128, 129, 5, 60, 0, 0, 129, 130,
		5, 61, 0, 0, 130, 26, 1, 0, 0, 0, 131, 132, 5, 62, 0, 0, 132, 28, 1, 0,
		0, 0, 133, 134, 5, 60, 0, 0, 134, 30, 1, 0, 0, 0, 135, 136, 5, 124, 0,
		0, 136, 140, 5, 124, 0, 0, 137, 138, 5, 111, 0, 0, 138, 140, 5, 114, 0,
		0, 139, 135, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 140, 32, 1, 0, 0, 0, 141,
		142, 5, 38, 0, 0, 142, 147, 5, 38, 0, 0, 143, 144, 5, 97, 0, 0, 144, 145,
		5, 110, 0, 0, 145, 147, 5, 100, 0, 0, 146, 141, 1, 0, 0, 0, 146, 143, 1,
		0, 0, 0, 147, 34, 1, 0, 0, 0, 148, 152, 5, 94, 0, 0, 149, 150, 5, 42, 0,
		0, 150, 152, 5, 42, 0, 0, 151, 148, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 152,
		36, 1, 0, 0, 0, 153, 154, 5, 42, 0, 0, 154, 38, 1, 0, 0, 0, 155, 156, 5,
		47, 0, 0, 156, 40, 1, 0, 0, 0, 157, 158, 5, 43, 0, 0, 158, 42, 1, 0, 0,
		0, 159, 160, 5, 45, 0, 0, 160, 44, 1, 0, 0, 0, 161, 166, 5, 37, 0, 0, 162,
		163, 5, 109, 0, 0, 163, 164, 5, 111, 0, 0, 164, 166, 5, 100, 0, 0, 165,
		161, 1, 0, 0, 0, 165, 162, 1, 0, 0, 0, 166, 46, 1, 0, 0, 0, 167, 168, 5,
		46, 0, 0, 168, 48, 1, 0, 0, 0, 169, 170, 5, 45, 0, 0, 170, 171, 5, 62,
		0, 0, 171, 50, 1, 0, 0, 0, 172, 173, 5, 115, 0, 0, 173, 174, 5, 116, 0,
		0, 174, 175, 5, 114, 0, 0, 175, 176, 5, 117, 0, 0, 176, 177, 5, 99, 0,
		0, 177, 178, 5, 116, 0, 0, 178, 52, 1, 0, 0, 0, 179, 180, 5, 109, 0, 0,
		180, 181, 5, 97, 0, 0, 181, 182, 5, 112, 0, 0, 182, 54, 1, 0, 0, 0, 183,
		184, 5, 102, 0, 0, 184, 185, 5, 110, 0, 0, 185, 56, 1, 0, 0, 0, 186, 187,
		5, 114, 0, 0, 187, 188, 5, 101, 0, 0, 188, 189, 5, 116, 0, 0, 189, 190,
		5, 117, 0, 0, 190, 191, 5, 114, 0, 0, 191, 192, 5, 110, 0, 0, 192, 58,
		1, 0, 0, 0, 193, 194, 5, 99, 0, 0, 194, 195, 5, 97, 0, 0, 195, 196, 5,
		115, 0, 0, 196, 197, 5, 101, 0, 0, 197, 60, 1, 0, 0, 0, 198, 199, 5, 100,
		0, 0, 199, 200, 5, 101, 0, 0, 200, 201, 5, 102, 0, 0, 201, 202, 5, 97,
		0, 0, 202, 203, 5, 117, 0, 0, 203, 204, 5, 108, 0, 0, 204, 205, 5, 116,
		0, 0, 205, 62, 1, 0, 0, 0, 206, 207, 5, 111, 0, 0, 207, 208, 5, 112, 0,
		0, 208, 209, 5, 101, 0, 0, 209, 210, 5, 110, 0, 0, 210, 64, 1, 0, 0, 0,
		211, 212, 5, 97, 0, 0, 212, 213, 5, 115, 0, 0, 213, 66, 1, 0, 0, 0, 214,
		215, 5, 105, 0, 0, 215, 216, 5, 102, 0, 0, 216, 68, 1, 0, 0, 0, 217, 218,
		5, 101, 0, 0, 218, 219, 5, 108, 0, 0, 219, 220, 5, 115, 0, 0, 220, 221,
		5, 101, 0, 0, 221, 70, 1, 0, 0, 0, 222, 223, 5, 102, 0, 0, 223, 224, 5,
		111, 0, 0, 224, 225, 5, 114, 0, 0, 225, 72, 1, 0, 0, 0, 226, 227, 5, 109,
		0, 0, 227, 228, 5, 97, 0, 0, 228, 229, 5, 116, 0, 0, 229, 230, 5, 99, 0,
		0, 230, 231, 5, 104, 0, 0, 231, 74, 1, 0, 0, 0, 232, 233, 5, 98, 0, 0,
		233, 234, 5, 114, 0, 0, 234, 235, 5, 101, 0, 0, 235, 236, 5, 97, 0, 0,
		236, 237, 5, 107, 0, 0, 237, 76, 1, 0, 0, 0, 238, 239, 5, 110, 0, 0, 239,
		240, 5, 101, 0, 0, 240, 241, 5, 120, 0, 0, 241, 255, 5, 116, 0, 0, 242,
		243, 5, 112, 0, 0, 243, 244, 5, 97, 0, 0, 244, 245, 5, 115, 0, 0, 245,
		255, 5, 115, 0, 0, 246, 247, 5, 99, 0, 0, 247, 248, 5, 111, 0, 0, 248,
		249, 5, 110, 0, 0, 249, 250, 5, 116, 0, 0, 250, 251, 5, 105, 0, 0, 251,
		252, 5, 110, 0, 0, 252, 253, 5, 117, 0, 0, 253, 255, 5, 101, 0, 0, 254,
		238, 1, 0, 0, 0, 254, 242, 1, 0, 0, 0, 254, 246, 1, 0, 0, 0, 255, 78, 1,
		0, 0, 0, 256, 257, 5, 116, 0, 0, 257, 258, 5, 114, 0, 0, 258, 259, 5, 117,
		0, 0, 259, 260, 5, 101, 0, 0, 260, 80, 1, 0, 0, 0, 261, 262, 5, 102, 0,
		0, 262, 263, 5, 97, 0, 0, 263, 264, 5, 108, 0, 0, 264, 265, 5, 115, 0,
		0, 265, 266, 5, 101, 0, 0, 266, 82, 1, 0, 0, 0, 267, 269, 7, 0, 0, 0, 268,
		267, 1, 0, 0, 0, 269, 270, 1, 0, 0, 0, 270, 268, 1, 0, 0, 0, 270, 271,
		1, 0, 0, 0, 271, 84, 1, 0, 0, 0, 272, 274, 7, 0, 0, 0, 273, 272, 1, 0,
		0, 0, 274, 275, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 275, 276, 1, 0, 0, 0,
		276, 283, 1, 0, 0, 0, 277, 279, 5, 46, 0, 0, 278, 280, 7, 0, 0, 0, 279,
		278, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0, 281, 282,
		1, 0, 0, 0, 282, 284, 1, 0, 0, 0, 283, 277, 1, 0, 0, 0, 283, 284, 1, 0,
		0, 0, 284, 86, 1, 0, 0, 0, 285, 289, 5, 34, 0, 0, 286, 288, 9, 0, 0, 0,
		287, 286, 1, 0, 0, 0, 288, 291, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 289,
		287, 1, 0, 0, 0, 290, 292, 1, 0, 0, 0, 291, 289, 1, 0, 0, 0, 292, 293,
		5, 34, 0, 0, 293, 88, 1, 0, 0, 0, 294, 299, 5, 33, 0, 0, 295, 296, 5, 110,
		0, 0, 296, 297, 5, 111, 0, 0, 297, 299, 5, 116, 0, 0, 298, 294, 1, 0, 0,
		0, 298, 295, 1, 0, 0, 0, 299, 90, 1, 0, 0, 0, 300, 301, 5, 58, 0, 0, 301,
		302, 5, 61, 0, 0, 302, 92, 1, 0, 0, 0, 303, 307, 7, 1, 0, 0, 304, 306,
		7, 2, 0, 0, 305, 304, 1, 0, 0, 0, 306, 309, 1, 0, 0, 0, 307, 305, 1, 0,
		0, 0, 307, 308, 1, 0, 0, 0, 308, 94, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0,
		310, 314, 5, 35, 0, 0, 311, 313, 8, 3, 0, 0, 312, 311, 1, 0, 0, 0, 313,
		316, 1, 0, 0, 0, 314, 312, 1, 0, 0, 0, 314, 315, 1, 0, 0, 0, 315, 317,
		1, 0, 0, 0, 316, 314, 1, 0, 0, 0, 317, 318, 6, 47, 0, 0, 318, 96, 1, 0,
		0, 0, 319, 321, 7, 4, 0, 0, 320, 319, 1, 0, 0, 0, 321, 322, 1, 0, 0, 0,
		322, 320, 1, 0, 0, 0, 322, 323, 1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324,
		325, 6, 48, 1, 0, 325, 98, 1, 0, 0, 0, 326, 328, 5, 13, 0, 0, 327, 326,
		1, 0, 0, 0, 327, 328, 1, 0, 0, 0, 328, 329, 1, 0, 0, 0, 329, 332, 5, 10,
		0, 0, 330, 332, 5, 13, 0, 0, 331, 327, 1, 0, 0, 0, 331, 330, 1, 0, 0, 0,
		332, 100, 1, 0, 0, 0, 17, 0, 139, 146, 151, 165, 254, 270, 275, 281, 283,
		289, 298, 307, 314, 322, 327, 331, 2, 0, 2, 0, 6, 0, 0,
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
	V2LexerIf             = 34
	V2LexerElse           = 35
	V2LexerFor            = 36
	V2LexerMatch          = 37
	V2LexerBreak          = 38
	V2LexerContinue       = 39
	V2LexerTrue           = 40
	V2LexerFalse          = 41
	V2LexerIntegerLiteral = 42
	V2LexerNumberLiteral  = 43
	V2LexerStringLiteral  = 44
	V2LexerNot            = 45
	V2LexerAssign         = 46
	V2LexerIdentifier     = 47
	V2LexerComment        = 48
	V2LexerWS             = 49
	V2LexerNewLine        = 50
)
