// Code generated from ./antlr4/V2Parser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // V2Parser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type V2Parser struct {
	*antlr.BaseParser
}

var V2ParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func v2parserParserInit() {
	staticData := &V2ParserParserStaticData
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
		"Identifier", "WS", "NewLine",
	}
	staticData.RuleNames = []string{
		"program", "openBlock", "structBlock", "funcBlock", "codeBlock", "declareBlock",
		"funcSig", "funcSignArgs", "type", "var", "vars", "varWithIdx", "index",
		"indexes", "lambda", "binaryOper", "unaryOper", "expr", "exprWithLambda",
		"exprOrAssign", "funcCall", "funcCallArgs", "stmtWithSep", "openStmt",
		"literal", "literalWithLambda", "arrayInitializer", "mapInitializer",
		"mapPair", "structInitializer", "structElementInitializer", "stmt",
		"declareStmt", "assgnStmt", "ifStmt", "loopStmt", "matchStmt", "matchCase",
		"jumpStmt", "sep",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 49, 432, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		3, 0, 86, 8, 0, 1, 0, 5, 0, 89, 8, 0, 10, 0, 12, 0, 92, 9, 0, 5, 0, 94,
		8, 0, 10, 0, 12, 0, 97, 9, 0, 1, 0, 1, 0, 1, 1, 5, 1, 102, 8, 1, 10, 1,
		12, 1, 105, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 5, 4, 118, 8, 4, 10, 4, 12, 4, 121, 9, 4, 5, 4, 123, 8, 4,
		10, 4, 12, 4, 126, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 5, 5, 132, 8, 5, 10, 5,
		12, 5, 135, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 3, 6, 142, 8, 6, 1, 6,
		1, 6, 3, 6, 146, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 154, 8,
		7, 10, 7, 12, 7, 157, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 164, 8,
		8, 3, 8, 166, 8, 8, 1, 9, 1, 9, 1, 9, 5, 9, 171, 8, 9, 10, 9, 12, 9, 174,
		9, 9, 1, 10, 1, 10, 1, 10, 5, 10, 179, 8, 10, 10, 10, 12, 10, 182, 9, 10,
		1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 5, 13, 192, 8,
		13, 10, 13, 12, 13, 195, 9, 13, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 201,
		8, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 3, 15, 216, 8, 15, 1, 16, 1, 16, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 231,
		8, 17, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 237, 8, 17, 10, 17, 12, 17, 240,
		9, 17, 1, 18, 1, 18, 3, 18, 244, 8, 18, 1, 19, 1, 19, 3, 19, 248, 8, 19,
		1, 20, 1, 20, 1, 20, 3, 20, 253, 8, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1,
		21, 5, 21, 260, 8, 21, 10, 21, 12, 21, 263, 9, 21, 1, 22, 1, 22, 5, 22,
		267, 8, 22, 10, 22, 12, 22, 270, 9, 22, 1, 23, 1, 23, 1, 23, 1, 23, 3,
		23, 276, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		3, 24, 286, 8, 24, 1, 25, 1, 25, 3, 25, 290, 8, 25, 1, 26, 3, 26, 293,
		8, 26, 1, 26, 1, 26, 1, 26, 3, 26, 298, 8, 26, 5, 26, 300, 8, 26, 10, 26,
		12, 26, 303, 9, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 5, 27, 310, 8, 27,
		10, 27, 12, 27, 313, 9, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 327, 8, 28, 1, 29, 1, 29,
		1, 29, 1, 29, 3, 29, 333, 8, 29, 5, 29, 335, 8, 29, 10, 29, 12, 29, 338,
		9, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 354, 8, 31, 1, 32, 1, 32, 1, 32,
		1, 33, 3, 33, 360, 8, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 373, 8, 34, 1, 35, 1, 35, 1, 35,
		3, 35, 378, 8, 35, 1, 35, 1, 35, 3, 35, 382, 8, 35, 1, 35, 1, 35, 3, 35,
		386, 8, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 3, 35, 393, 8, 35, 1, 35,
		1, 35, 3, 35, 397, 8, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 5,
		36, 405, 8, 36, 10, 36, 12, 36, 408, 9, 36, 1, 36, 1, 36, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 3, 37, 420, 8, 37, 1, 38, 1,
		38, 1, 38, 1, 38, 3, 38, 426, 8, 38, 3, 38, 428, 8, 38, 1, 39, 1, 39, 1,
		39, 0, 1, 34, 40, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64,
		66, 68, 70, 72, 74, 76, 78, 0, 5, 1, 0, 16, 17, 1, 0, 19, 20, 1, 0, 21,
		22, 2, 0, 21, 22, 45, 45, 2, 0, 9, 9, 49, 49, 465, 0, 80, 1, 0, 0, 0, 2,
		103, 1, 0, 0, 0, 4, 106, 1, 0, 0, 0, 6, 110, 1, 0, 0, 0, 8, 114, 1, 0,
		0, 0, 10, 129, 1, 0, 0, 0, 12, 138, 1, 0, 0, 0, 14, 147, 1, 0, 0, 0, 16,
		165, 1, 0, 0, 0, 18, 167, 1, 0, 0, 0, 20, 175, 1, 0, 0, 0, 22, 183, 1,
		0, 0, 0, 24, 186, 1, 0, 0, 0, 26, 193, 1, 0, 0, 0, 28, 196, 1, 0, 0, 0,
		30, 215, 1, 0, 0, 0, 32, 217, 1, 0, 0, 0, 34, 230, 1, 0, 0, 0, 36, 243,
		1, 0, 0, 0, 38, 247, 1, 0, 0, 0, 40, 249, 1, 0, 0, 0, 42, 256, 1, 0, 0,
		0, 44, 264, 1, 0, 0, 0, 46, 271, 1, 0, 0, 0, 48, 285, 1, 0, 0, 0, 50, 289,
		1, 0, 0, 0, 52, 292, 1, 0, 0, 0, 54, 306, 1, 0, 0, 0, 56, 326, 1, 0, 0,
		0, 58, 328, 1, 0, 0, 0, 60, 341, 1, 0, 0, 0, 62, 353, 1, 0, 0, 0, 64, 355,
		1, 0, 0, 0, 66, 359, 1, 0, 0, 0, 68, 365, 1, 0, 0, 0, 70, 396, 1, 0, 0,
		0, 72, 398, 1, 0, 0, 0, 74, 419, 1, 0, 0, 0, 76, 427, 1, 0, 0, 0, 78, 429,
		1, 0, 0, 0, 80, 95, 3, 2, 1, 0, 81, 86, 3, 4, 2, 0, 82, 86, 3, 6, 3, 0,
		83, 86, 3, 62, 31, 0, 84, 86, 3, 34, 17, 0, 85, 81, 1, 0, 0, 0, 85, 82,
		1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 84, 1, 0, 0, 0, 86, 90, 1, 0, 0, 0,
		87, 89, 3, 78, 39, 0, 88, 87, 1, 0, 0, 0, 89, 92, 1, 0, 0, 0, 90, 88, 1,
		0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 94, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 93,
		85, 1, 0, 0, 0, 94, 97, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0,
		0, 96, 98, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 98, 99, 5, 0, 0, 1, 99, 1, 1,
		0, 0, 0, 100, 102, 3, 46, 23, 0, 101, 100, 1, 0, 0, 0, 102, 105, 1, 0,
		0, 0, 103, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 3, 1, 0, 0, 0, 105,
		103, 1, 0, 0, 0, 106, 107, 5, 26, 0, 0, 107, 108, 5, 47, 0, 0, 108, 109,
		3, 10, 5, 0, 109, 5, 1, 0, 0, 0, 110, 111, 5, 28, 0, 0, 111, 112, 3, 12,
		6, 0, 112, 113, 3, 8, 4, 0, 113, 7, 1, 0, 0, 0, 114, 124, 5, 1, 0, 0, 115,
		119, 3, 62, 31, 0, 116, 118, 3, 78, 39, 0, 117, 116, 1, 0, 0, 0, 118, 121,
		1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 123, 1, 0,
		0, 0, 121, 119, 1, 0, 0, 0, 122, 115, 1, 0, 0, 0, 123, 126, 1, 0, 0, 0,
		124, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 127, 1, 0, 0, 0, 126,
		124, 1, 0, 0, 0, 127, 128, 5, 2, 0, 0, 128, 9, 1, 0, 0, 0, 129, 133, 5,
		1, 0, 0, 130, 132, 3, 64, 32, 0, 131, 130, 1, 0, 0, 0, 132, 135, 1, 0,
		0, 0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 136, 1, 0, 0, 0,
		135, 133, 1, 0, 0, 0, 136, 137, 5, 2, 0, 0, 137, 11, 1, 0, 0, 0, 138, 139,
		5, 47, 0, 0, 139, 141, 5, 3, 0, 0, 140, 142, 3, 14, 7, 0, 141, 140, 1,
		0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 145, 5, 4, 0,
		0, 144, 146, 3, 16, 8, 0, 145, 144, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146,
		13, 1, 0, 0, 0, 147, 148, 3, 16, 8, 0, 148, 155, 5, 47, 0, 0, 149, 150,
		5, 7, 0, 0, 150, 151, 3, 16, 8, 0, 151, 152, 5, 47, 0, 0, 152, 154, 1,
		0, 0, 0, 153, 149, 1, 0, 0, 0, 154, 157, 1, 0, 0, 0, 155, 153, 1, 0, 0,
		0, 155, 156, 1, 0, 0, 0, 156, 15, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 158,
		166, 5, 27, 0, 0, 159, 166, 5, 28, 0, 0, 160, 163, 5, 47, 0, 0, 161, 162,
		5, 5, 0, 0, 162, 164, 5, 6, 0, 0, 163, 161, 1, 0, 0, 0, 163, 164, 1, 0,
		0, 0, 164, 166, 1, 0, 0, 0, 165, 158, 1, 0, 0, 0, 165, 159, 1, 0, 0, 0,
		165, 160, 1, 0, 0, 0, 166, 17, 1, 0, 0, 0, 167, 172, 5, 47, 0, 0, 168,
		169, 5, 24, 0, 0, 169, 171, 5, 47, 0, 0, 170, 168, 1, 0, 0, 0, 171, 174,
		1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 19, 1, 0,
		0, 0, 174, 172, 1, 0, 0, 0, 175, 180, 3, 18, 9, 0, 176, 177, 5, 7, 0, 0,
		177, 179, 3, 20, 10, 0, 178, 176, 1, 0, 0, 0, 179, 182, 1, 0, 0, 0, 180,
		178, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 21, 1, 0, 0, 0, 182, 180, 1,
		0, 0, 0, 183, 184, 3, 18, 9, 0, 184, 185, 3, 26, 13, 0, 185, 23, 1, 0,
		0, 0, 186, 187, 5, 5, 0, 0, 187, 188, 3, 34, 17, 0, 188, 189, 5, 6, 0,
		0, 189, 25, 1, 0, 0, 0, 190, 192, 3, 24, 12, 0, 191, 190, 1, 0, 0, 0, 192,
		195, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 27, 1,
		0, 0, 0, 195, 193, 1, 0, 0, 0, 196, 197, 5, 3, 0, 0, 197, 198, 3, 14, 7,
		0, 198, 200, 5, 4, 0, 0, 199, 201, 3, 16, 8, 0, 200, 199, 1, 0, 0, 0, 200,
		201, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 203, 3, 8, 4, 0, 203, 29, 1,
		0, 0, 0, 204, 216, 5, 10, 0, 0, 205, 216, 5, 11, 0, 0, 206, 216, 5, 14,
		0, 0, 207, 216, 5, 15, 0, 0, 208, 216, 5, 12, 0, 0, 209, 216, 5, 13, 0,
		0, 210, 216, 7, 0, 0, 0, 211, 216, 5, 18, 0, 0, 212, 216, 7, 1, 0, 0, 213,
		216, 5, 23, 0, 0, 214, 216, 7, 2, 0, 0, 215, 204, 1, 0, 0, 0, 215, 205,
		1, 0, 0, 0, 215, 206, 1, 0, 0, 0, 215, 207, 1, 0, 0, 0, 215, 208, 1, 0,
		0, 0, 215, 209, 1, 0, 0, 0, 215, 210, 1, 0, 0, 0, 215, 211, 1, 0, 0, 0,
		215, 212, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 215, 214, 1, 0, 0, 0, 216,
		31, 1, 0, 0, 0, 217, 218, 7, 3, 0, 0, 218, 33, 1, 0, 0, 0, 219, 220, 6,
		17, -1, 0, 220, 231, 3, 40, 20, 0, 221, 222, 3, 32, 16, 0, 222, 223, 3,
		34, 17, 5, 223, 231, 1, 0, 0, 0, 224, 231, 3, 48, 24, 0, 225, 226, 5, 3,
		0, 0, 226, 227, 3, 34, 17, 0, 227, 228, 5, 4, 0, 0, 228, 231, 1, 0, 0,
		0, 229, 231, 5, 47, 0, 0, 230, 219, 1, 0, 0, 0, 230, 221, 1, 0, 0, 0, 230,
		224, 1, 0, 0, 0, 230, 225, 1, 0, 0, 0, 230, 229, 1, 0, 0, 0, 231, 238,
		1, 0, 0, 0, 232, 233, 10, 4, 0, 0, 233, 234, 3, 30, 15, 0, 234, 235, 3,
		34, 17, 5, 235, 237, 1, 0, 0, 0, 236, 232, 1, 0, 0, 0, 237, 240, 1, 0,
		0, 0, 238, 236, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 35, 1, 0, 0, 0,
		240, 238, 1, 0, 0, 0, 241, 244, 3, 28, 14, 0, 242, 244, 3, 34, 17, 0, 243,
		241, 1, 0, 0, 0, 243, 242, 1, 0, 0, 0, 244, 37, 1, 0, 0, 0, 245, 248, 3,
		34, 17, 0, 246, 248, 3, 66, 33, 0, 247, 245, 1, 0, 0, 0, 247, 246, 1, 0,
		0, 0, 248, 39, 1, 0, 0, 0, 249, 250, 3, 18, 9, 0, 250, 252, 5, 3, 0, 0,
		251, 253, 3, 42, 21, 0, 252, 251, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253,
		254, 1, 0, 0, 0, 254, 255, 5, 4, 0, 0, 255, 41, 1, 0, 0, 0, 256, 261, 3,
		34, 17, 0, 257, 258, 5, 7, 0, 0, 258, 260, 3, 34, 17, 0, 259, 257, 1, 0,
		0, 0, 260, 263, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0,
		262, 43, 1, 0, 0, 0, 263, 261, 1, 0, 0, 0, 264, 268, 3, 62, 31, 0, 265,
		267, 3, 78, 39, 0, 266, 265, 1, 0, 0, 0, 267, 270, 1, 0, 0, 0, 268, 266,
		1, 0, 0, 0, 268, 269, 1, 0, 0, 0, 269, 45, 1, 0, 0, 0, 270, 268, 1, 0,
		0, 0, 271, 272, 5, 32, 0, 0, 272, 275, 5, 44, 0, 0, 273, 274, 5, 33, 0,
		0, 274, 276, 5, 47, 0, 0, 275, 273, 1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276,
		47, 1, 0, 0, 0, 277, 286, 5, 40, 0, 0, 278, 286, 5, 41, 0, 0, 279, 286,
		5, 42, 0, 0, 280, 286, 5, 43, 0, 0, 281, 286, 5, 44, 0, 0, 282, 286, 3,
		52, 26, 0, 283, 286, 3, 54, 27, 0, 284, 286, 3, 58, 29, 0, 285, 277, 1,
		0, 0, 0, 285, 278, 1, 0, 0, 0, 285, 279, 1, 0, 0, 0, 285, 280, 1, 0, 0,
		0, 285, 281, 1, 0, 0, 0, 285, 282, 1, 0, 0, 0, 285, 283, 1, 0, 0, 0, 285,
		284, 1, 0, 0, 0, 286, 49, 1, 0, 0, 0, 287, 290, 3, 48, 24, 0, 288, 290,
		3, 28, 14, 0, 289, 287, 1, 0, 0, 0, 289, 288, 1, 0, 0, 0, 290, 51, 1, 0,
		0, 0, 291, 293, 3, 16, 8, 0, 292, 291, 1, 0, 0, 0, 292, 293, 1, 0, 0, 0,
		293, 294, 1, 0, 0, 0, 294, 301, 5, 5, 0, 0, 295, 297, 3, 34, 17, 0, 296,
		298, 5, 7, 0, 0, 297, 296, 1, 0, 0, 0, 297, 298, 1, 0, 0, 0, 298, 300,
		1, 0, 0, 0, 299, 295, 1, 0, 0, 0, 300, 303, 1, 0, 0, 0, 301, 299, 1, 0,
		0, 0, 301, 302, 1, 0, 0, 0, 302, 304, 1, 0, 0, 0, 303, 301, 1, 0, 0, 0,
		304, 305, 5, 6, 0, 0, 305, 53, 1, 0, 0, 0, 306, 307, 5, 27, 0, 0, 307,
		311, 5, 1, 0, 0, 308, 310, 3, 56, 28, 0, 309, 308, 1, 0, 0, 0, 310, 313,
		1, 0, 0, 0, 311, 309, 1, 0, 0, 0, 311, 312, 1, 0, 0, 0, 312, 314, 1, 0,
		0, 0, 313, 311, 1, 0, 0, 0, 314, 315, 5, 2, 0, 0, 315, 55, 1, 0, 0, 0,
		316, 317, 5, 3, 0, 0, 317, 318, 3, 34, 17, 0, 318, 319, 5, 7, 0, 0, 319,
		320, 3, 36, 18, 0, 320, 321, 5, 4, 0, 0, 321, 327, 1, 0, 0, 0, 322, 323,
		3, 34, 17, 0, 323, 324, 5, 25, 0, 0, 324, 325, 3, 36, 18, 0, 325, 327,
		1, 0, 0, 0, 326, 316, 1, 0, 0, 0, 326, 322, 1, 0, 0, 0, 327, 57, 1, 0,
		0, 0, 328, 329, 3, 16, 8, 0, 329, 336, 5, 1, 0, 0, 330, 332, 3, 60, 30,
		0, 331, 333, 5, 7, 0, 0, 332, 331, 1, 0, 0, 0, 332, 333, 1, 0, 0, 0, 333,
		335, 1, 0, 0, 0, 334, 330, 1, 0, 0, 0, 335, 338, 1, 0, 0, 0, 336, 334,
		1, 0, 0, 0, 336, 337, 1, 0, 0, 0, 337, 339, 1, 0, 0, 0, 338, 336, 1, 0,
		0, 0, 339, 340, 5, 2, 0, 0, 340, 59, 1, 0, 0, 0, 341, 342, 5, 47, 0, 0,
		342, 343, 5, 7, 0, 0, 343, 344, 3, 34, 17, 0, 344, 61, 1, 0, 0, 0, 345,
		354, 3, 64, 32, 0, 346, 354, 3, 66, 33, 0, 347, 354, 3, 76, 38, 0, 348,
		354, 3, 68, 34, 0, 349, 354, 3, 70, 35, 0, 350, 354, 3, 72, 36, 0, 351,
		354, 3, 8, 4, 0, 352, 354, 3, 40, 20, 0, 353, 345, 1, 0, 0, 0, 353, 346,
		1, 0, 0, 0, 353, 347, 1, 0, 0, 0, 353, 348, 1, 0, 0, 0, 353, 349, 1, 0,
		0, 0, 353, 350, 1, 0, 0, 0, 353, 351, 1, 0, 0, 0, 353, 352, 1, 0, 0, 0,
		354, 63, 1, 0, 0, 0, 355, 356, 3, 16, 8, 0, 356, 357, 3, 20, 10, 0, 357,
		65, 1, 0, 0, 0, 358, 360, 3, 16, 8, 0, 359, 358, 1, 0, 0, 0, 359, 360,
		1, 0, 0, 0, 360, 361, 1, 0, 0, 0, 361, 362, 3, 22, 11, 0, 362, 363, 5,
		46, 0, 0, 363, 364, 3, 36, 18, 0, 364, 67, 1, 0, 0, 0, 365, 366, 5, 34,
		0, 0, 366, 367, 5, 3, 0, 0, 367, 368, 3, 34, 17, 0, 368, 369, 5, 4, 0,
		0, 369, 372, 3, 8, 4, 0, 370, 371, 5, 35, 0, 0, 371, 373, 3, 8, 4, 0, 372,
		370, 1, 0, 0, 0, 372, 373, 1, 0, 0, 0, 373, 69, 1, 0, 0, 0, 374, 375, 5,
		36, 0, 0, 375, 377, 5, 3, 0, 0, 376, 378, 3, 38, 19, 0, 377, 376, 1, 0,
		0, 0, 377, 378, 1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379, 381, 5, 9, 0, 0,
		380, 382, 3, 34, 17, 0, 381, 380, 1, 0, 0, 0, 381, 382, 1, 0, 0, 0, 382,
		383, 1, 0, 0, 0, 383, 385, 5, 9, 0, 0, 384, 386, 3, 34, 17, 0, 385, 384,
		1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386, 387, 1, 0, 0, 0, 387, 388, 5, 4,
		0, 0, 388, 397, 3, 8, 4, 0, 389, 390, 5, 36, 0, 0, 390, 392, 5, 3, 0, 0,
		391, 393, 3, 34, 17, 0, 392, 391, 1, 0, 0, 0, 392, 393, 1, 0, 0, 0, 393,
		394, 1, 0, 0, 0, 394, 395, 5, 4, 0, 0, 395, 397, 3, 8, 4, 0, 396, 374,
		1, 0, 0, 0, 396, 389, 1, 0, 0, 0, 397, 71, 1, 0, 0, 0, 398, 399, 5, 37,
		0, 0, 399, 400, 5, 3, 0, 0, 400, 401, 3, 34, 17, 0, 401, 402, 5, 4, 0,
		0, 402, 406, 5, 1, 0, 0, 403, 405, 3, 74, 37, 0, 404, 403, 1, 0, 0, 0,
		405, 408, 1, 0, 0, 0, 406, 404, 1, 0, 0, 0, 406, 407, 1, 0, 0, 0, 407,
		409, 1, 0, 0, 0, 408, 406, 1, 0, 0, 0, 409, 410, 5, 2, 0, 0, 410, 73, 1,
		0, 0, 0, 411, 412, 5, 30, 0, 0, 412, 413, 3, 34, 17, 0, 413, 414, 5, 8,
		0, 0, 414, 415, 3, 8, 4, 0, 415, 420, 1, 0, 0, 0, 416, 417, 5, 31, 0, 0,
		417, 418, 5, 8, 0, 0, 418, 420, 3, 8, 4, 0, 419, 411, 1, 0, 0, 0, 419,
		416, 1, 0, 0, 0, 420, 75, 1, 0, 0, 0, 421, 428, 5, 39, 0, 0, 422, 428,
		5, 38, 0, 0, 423, 425, 5, 29, 0, 0, 424, 426, 3, 34, 17, 0, 425, 424, 1,
		0, 0, 0, 425, 426, 1, 0, 0, 0, 426, 428, 1, 0, 0, 0, 427, 421, 1, 0, 0,
		0, 427, 422, 1, 0, 0, 0, 427, 423, 1, 0, 0, 0, 428, 77, 1, 0, 0, 0, 429,
		430, 7, 4, 0, 0, 430, 79, 1, 0, 0, 0, 46, 85, 90, 95, 103, 119, 124, 133,
		141, 145, 155, 163, 165, 172, 180, 193, 200, 215, 230, 238, 243, 247, 252,
		261, 268, 275, 285, 289, 292, 297, 301, 311, 326, 332, 336, 353, 359, 372,
		377, 381, 385, 392, 396, 406, 419, 425, 427,
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

// V2ParserInit initializes any static state used to implement V2Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewV2Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func V2ParserInit() {
	staticData := &V2ParserParserStaticData
	staticData.once.Do(v2parserParserInit)
}

// NewV2Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewV2Parser(input antlr.TokenStream) *V2Parser {
	V2ParserInit()
	this := new(V2Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &V2ParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "V2Parser.g4"

	return this
}

// V2Parser tokens.
const (
	V2ParserEOF            = antlr.TokenEOF
	V2ParserLBrack         = 1
	V2ParserRBrack         = 2
	V2ParserLParen         = 3
	V2ParserRParen         = 4
	V2ParserLSquare        = 5
	V2ParserRSquare        = 6
	V2ParserComma          = 7
	V2ParserCol            = 8
	V2ParserSemi           = 9
	V2ParserEquals         = 10
	V2ParserNotEq          = 11
	V2ParserGreaterEq      = 12
	V2ParserLessEq         = 13
	V2ParserGreater        = 14
	V2ParserLess           = 15
	V2ParserOr             = 16
	V2ParserAnd            = 17
	V2ParserPow            = 18
	V2ParserMul            = 19
	V2ParserDiv            = 20
	V2ParserAdd            = 21
	V2ParserSub            = 22
	V2ParserMod            = 23
	V2ParserDot            = 24
	V2ParserTo             = 25
	V2ParserStruct         = 26
	V2ParserMap            = 27
	V2ParserFunction       = 28
	V2ParserReturn         = 29
	V2ParserCase           = 30
	V2ParserDefault        = 31
	V2ParserOpen           = 32
	V2ParserAs             = 33
	V2ParserIf             = 34
	V2ParserElse           = 35
	V2ParserFor            = 36
	V2ParserMatch          = 37
	V2ParserBreak          = 38
	V2ParserContinue       = 39
	V2ParserTrue           = 40
	V2ParserFalse          = 41
	V2ParserIntegerLiteral = 42
	V2ParserNumberLiteral  = 43
	V2ParserStringLiteral  = 44
	V2ParserNot            = 45
	V2ParserAssign         = 46
	V2ParserIdentifier     = 47
	V2ParserWS             = 48
	V2ParserNewLine        = 49
)

// V2Parser rules.
const (
	V2ParserRULE_program                  = 0
	V2ParserRULE_openBlock                = 1
	V2ParserRULE_structBlock              = 2
	V2ParserRULE_funcBlock                = 3
	V2ParserRULE_codeBlock                = 4
	V2ParserRULE_declareBlock             = 5
	V2ParserRULE_funcSig                  = 6
	V2ParserRULE_funcSignArgs             = 7
	V2ParserRULE_type                     = 8
	V2ParserRULE_var                      = 9
	V2ParserRULE_vars                     = 10
	V2ParserRULE_varWithIdx               = 11
	V2ParserRULE_index                    = 12
	V2ParserRULE_indexes                  = 13
	V2ParserRULE_lambda                   = 14
	V2ParserRULE_binaryOper               = 15
	V2ParserRULE_unaryOper                = 16
	V2ParserRULE_expr                     = 17
	V2ParserRULE_exprWithLambda           = 18
	V2ParserRULE_exprOrAssign             = 19
	V2ParserRULE_funcCall                 = 20
	V2ParserRULE_funcCallArgs             = 21
	V2ParserRULE_stmtWithSep              = 22
	V2ParserRULE_openStmt                 = 23
	V2ParserRULE_literal                  = 24
	V2ParserRULE_literalWithLambda        = 25
	V2ParserRULE_arrayInitializer         = 26
	V2ParserRULE_mapInitializer           = 27
	V2ParserRULE_mapPair                  = 28
	V2ParserRULE_structInitializer        = 29
	V2ParserRULE_structElementInitializer = 30
	V2ParserRULE_stmt                     = 31
	V2ParserRULE_declareStmt              = 32
	V2ParserRULE_assgnStmt                = 33
	V2ParserRULE_ifStmt                   = 34
	V2ParserRULE_loopStmt                 = 35
	V2ParserRULE_matchStmt                = 36
	V2ParserRULE_matchCase                = 37
	V2ParserRULE_jumpStmt                 = 38
	V2ParserRULE_sep                      = 39
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpenBlock() IOpenBlockContext
	EOF() antlr.TerminalNode
	AllStructBlock() []IStructBlockContext
	StructBlock(i int) IStructBlockContext
	AllFuncBlock() []IFuncBlockContext
	FuncBlock(i int) IFuncBlockContext
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllSep() []ISepContext
	Sep(i int) ISepContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) OpenBlock() IOpenBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpenBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpenBlockContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(V2ParserEOF, 0)
}

func (s *ProgramContext) AllStructBlock() []IStructBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructBlockContext); ok {
			len++
		}
	}

	tst := make([]IStructBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructBlockContext); ok {
			tst[i] = t.(IStructBlockContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) StructBlock(i int) IStructBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructBlockContext)
}

func (s *ProgramContext) AllFuncBlock() []IFuncBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncBlockContext); ok {
			len++
		}
	}

	tst := make([]IFuncBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncBlockContext); ok {
			tst[i] = t.(IFuncBlockContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) FuncBlock(i int) IFuncBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncBlockContext)
}

func (s *ProgramContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ProgramContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ProgramContext) AllSep() []ISepContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISepContext); ok {
			len++
		}
	}

	tst := make([]ISepContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISepContext); ok {
			tst[i] = t.(ISepContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Sep(i int) ISepContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISepContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISepContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, V2ParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.OpenBlock()
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&211055705849898) != 0 {
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(81)
				p.StructBlock()
			}

		case 2:
			{
				p.SetState(82)
				p.FuncBlock()
			}

		case 3:
			{
				p.SetState(83)
				p.Stmt()
			}

		case 4:
			{
				p.SetState(84)
				p.expr(0)
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == V2ParserSemi || _la == V2ParserNewLine {
			{
				p.SetState(87)
				p.Sep()
			}

			p.SetState(92)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(98)
		p.Match(V2ParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOpenBlockContext is an interface to support dynamic dispatch.
type IOpenBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllOpenStmt() []IOpenStmtContext
	OpenStmt(i int) IOpenStmtContext

	// IsOpenBlockContext differentiates from other interfaces.
	IsOpenBlockContext()
}

type OpenBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpenBlockContext() *OpenBlockContext {
	var p = new(OpenBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_openBlock
	return p
}

func InitEmptyOpenBlockContext(p *OpenBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_openBlock
}

func (*OpenBlockContext) IsOpenBlockContext() {}

func NewOpenBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpenBlockContext {
	var p = new(OpenBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_openBlock

	return p
}

func (s *OpenBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *OpenBlockContext) AllOpenStmt() []IOpenStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOpenStmtContext); ok {
			len++
		}
	}

	tst := make([]IOpenStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOpenStmtContext); ok {
			tst[i] = t.(IOpenStmtContext)
			i++
		}
	}

	return tst
}

func (s *OpenBlockContext) OpenStmt(i int) IOpenStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpenStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpenStmtContext)
}

func (s *OpenBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpenBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpenBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterOpenBlock(s)
	}
}

func (s *OpenBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitOpenBlock(s)
	}
}

func (s *OpenBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitOpenBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) OpenBlock() (localctx IOpenBlockContext) {
	localctx = NewOpenBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, V2ParserRULE_openBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == V2ParserOpen {
		{
			p.SetState(100)
			p.OpenStmt()
		}

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructBlockContext is an interface to support dynamic dispatch.
type IStructBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Struct() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	DeclareBlock() IDeclareBlockContext

	// IsStructBlockContext differentiates from other interfaces.
	IsStructBlockContext()
}

type StructBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructBlockContext() *StructBlockContext {
	var p = new(StructBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_structBlock
	return p
}

func InitEmptyStructBlockContext(p *StructBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_structBlock
}

func (*StructBlockContext) IsStructBlockContext() {}

func NewStructBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructBlockContext {
	var p = new(StructBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_structBlock

	return p
}

func (s *StructBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *StructBlockContext) Struct() antlr.TerminalNode {
	return s.GetToken(V2ParserStruct, 0)
}

func (s *StructBlockContext) Identifier() antlr.TerminalNode {
	return s.GetToken(V2ParserIdentifier, 0)
}

func (s *StructBlockContext) DeclareBlock() IDeclareBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclareBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclareBlockContext)
}

func (s *StructBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterStructBlock(s)
	}
}

func (s *StructBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitStructBlock(s)
	}
}

func (s *StructBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitStructBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) StructBlock() (localctx IStructBlockContext) {
	localctx = NewStructBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, V2ParserRULE_structBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Match(V2ParserStruct)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.Match(V2ParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(108)
		p.DeclareBlock()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncBlockContext is an interface to support dynamic dispatch.
type IFuncBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Function() antlr.TerminalNode
	FuncSig() IFuncSigContext
	CodeBlock() ICodeBlockContext

	// IsFuncBlockContext differentiates from other interfaces.
	IsFuncBlockContext()
}

type FuncBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncBlockContext() *FuncBlockContext {
	var p = new(FuncBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_funcBlock
	return p
}

func InitEmptyFuncBlockContext(p *FuncBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_funcBlock
}

func (*FuncBlockContext) IsFuncBlockContext() {}

func NewFuncBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncBlockContext {
	var p = new(FuncBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_funcBlock

	return p
}

func (s *FuncBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncBlockContext) Function() antlr.TerminalNode {
	return s.GetToken(V2ParserFunction, 0)
}

func (s *FuncBlockContext) FuncSig() IFuncSigContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncSigContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncSigContext)
}

func (s *FuncBlockContext) CodeBlock() ICodeBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICodeBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICodeBlockContext)
}

func (s *FuncBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterFuncBlock(s)
	}
}

func (s *FuncBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitFuncBlock(s)
	}
}

func (s *FuncBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitFuncBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) FuncBlock() (localctx IFuncBlockContext) {
	localctx = NewFuncBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, V2ParserRULE_funcBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.Match(V2ParserFunction)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		p.FuncSig()
	}
	{
		p.SetState(112)
		p.CodeBlock()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICodeBlockContext is an interface to support dynamic dispatch.
type ICodeBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBrack() antlr.TerminalNode
	RBrack() antlr.TerminalNode
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext
	AllSep() []ISepContext
	Sep(i int) ISepContext

	// IsCodeBlockContext differentiates from other interfaces.
	IsCodeBlockContext()
}

type CodeBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCodeBlockContext() *CodeBlockContext {
	var p = new(CodeBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_codeBlock
	return p
}

func InitEmptyCodeBlockContext(p *CodeBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_codeBlock
}

func (*CodeBlockContext) IsCodeBlockContext() {}

func NewCodeBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CodeBlockContext {
	var p = new(CodeBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_codeBlock

	return p
}

func (s *CodeBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *CodeBlockContext) LBrack() antlr.TerminalNode {
	return s.GetToken(V2ParserLBrack, 0)
}

func (s *CodeBlockContext) RBrack() antlr.TerminalNode {
	return s.GetToken(V2ParserRBrack, 0)
}

func (s *CodeBlockContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *CodeBlockContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *CodeBlockContext) AllSep() []ISepContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISepContext); ok {
			len++
		}
	}

	tst := make([]ISepContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISepContext); ok {
			tst[i] = t.(ISepContext)
			i++
		}
	}

	return tst
}

func (s *CodeBlockContext) Sep(i int) ISepContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISepContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISepContext)
}

func (s *CodeBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CodeBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CodeBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterCodeBlock(s)
	}
}

func (s *CodeBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitCodeBlock(s)
	}
}

func (s *CodeBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitCodeBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) CodeBlock() (localctx ICodeBlockContext) {
	localctx = NewCodeBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, V2ParserRULE_codeBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(V2ParserLBrack)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&141786399899650) != 0 {
		{
			p.SetState(115)
			p.Stmt()
		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == V2ParserSemi || _la == V2ParserNewLine {
			{
				p.SetState(116)
				p.Sep()
			}

			p.SetState(121)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(127)
		p.Match(V2ParserRBrack)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclareBlockContext is an interface to support dynamic dispatch.
type IDeclareBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBrack() antlr.TerminalNode
	RBrack() antlr.TerminalNode
	AllDeclareStmt() []IDeclareStmtContext
	DeclareStmt(i int) IDeclareStmtContext

	// IsDeclareBlockContext differentiates from other interfaces.
	IsDeclareBlockContext()
}

type DeclareBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclareBlockContext() *DeclareBlockContext {
	var p = new(DeclareBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_declareBlock
	return p
}

func InitEmptyDeclareBlockContext(p *DeclareBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_declareBlock
}

func (*DeclareBlockContext) IsDeclareBlockContext() {}

func NewDeclareBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclareBlockContext {
	var p = new(DeclareBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_declareBlock

	return p
}

func (s *DeclareBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclareBlockContext) LBrack() antlr.TerminalNode {
	return s.GetToken(V2ParserLBrack, 0)
}

func (s *DeclareBlockContext) RBrack() antlr.TerminalNode {
	return s.GetToken(V2ParserRBrack, 0)
}

func (s *DeclareBlockContext) AllDeclareStmt() []IDeclareStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclareStmtContext); ok {
			len++
		}
	}

	tst := make([]IDeclareStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclareStmtContext); ok {
			tst[i] = t.(IDeclareStmtContext)
			i++
		}
	}

	return tst
}

func (s *DeclareBlockContext) DeclareStmt(i int) IDeclareStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclareStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclareStmtContext)
}

func (s *DeclareBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclareBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclareBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterDeclareBlock(s)
	}
}

func (s *DeclareBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitDeclareBlock(s)
	}
}

func (s *DeclareBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitDeclareBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) DeclareBlock() (localctx IDeclareBlockContext) {
	localctx = NewDeclareBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, V2ParserRULE_declareBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.Match(V2ParserLBrack)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140737891008512) != 0 {
		{
			p.SetState(130)
			p.DeclareStmt()
		}

		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(136)
		p.Match(V2ParserRBrack)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncSigContext is an interface to support dynamic dispatch.
type IFuncSigContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	LParen() antlr.TerminalNode
	RParen() antlr.TerminalNode
	FuncSignArgs() IFuncSignArgsContext
	Type_() ITypeContext

	// IsFuncSigContext differentiates from other interfaces.
	IsFuncSigContext()
}

type FuncSigContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncSigContext() *FuncSigContext {
	var p = new(FuncSigContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_funcSig
	return p
}

func InitEmptyFuncSigContext(p *FuncSigContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_funcSig
}

func (*FuncSigContext) IsFuncSigContext() {}

func NewFuncSigContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncSigContext {
	var p = new(FuncSigContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_funcSig

	return p
}

func (s *FuncSigContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncSigContext) Identifier() antlr.TerminalNode {
	return s.GetToken(V2ParserIdentifier, 0)
}

func (s *FuncSigContext) LParen() antlr.TerminalNode {
	return s.GetToken(V2ParserLParen, 0)
}

func (s *FuncSigContext) RParen() antlr.TerminalNode {
	return s.GetToken(V2ParserRParen, 0)
}

func (s *FuncSigContext) FuncSignArgs() IFuncSignArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncSignArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncSignArgsContext)
}

func (s *FuncSigContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FuncSigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncSigContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncSigContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterFuncSig(s)
	}
}

func (s *FuncSigContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitFuncSig(s)
	}
}

func (s *FuncSigContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitFuncSig(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) FuncSig() (localctx IFuncSigContext) {
	localctx = NewFuncSigContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, V2ParserRULE_funcSig)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(V2ParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.Match(V2ParserLParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140737891008512) != 0 {
		{
			p.SetState(140)
			p.FuncSignArgs()
		}

	}
	{
		p.SetState(143)
		p.Match(V2ParserRParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140737891008512) != 0 {
		{
			p.SetState(144)
			p.Type_()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncSignArgsContext is an interface to support dynamic dispatch.
type IFuncSignArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllType_() []ITypeContext
	Type_(i int) ITypeContext
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsFuncSignArgsContext differentiates from other interfaces.
	IsFuncSignArgsContext()
}

type FuncSignArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncSignArgsContext() *FuncSignArgsContext {
	var p = new(FuncSignArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_funcSignArgs
	return p
}

func InitEmptyFuncSignArgsContext(p *FuncSignArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_funcSignArgs
}

func (*FuncSignArgsContext) IsFuncSignArgsContext() {}

func NewFuncSignArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncSignArgsContext {
	var p = new(FuncSignArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_funcSignArgs

	return p
}

func (s *FuncSignArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncSignArgsContext) AllType_() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *FuncSignArgsContext) Type_(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FuncSignArgsContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(V2ParserIdentifier)
}

func (s *FuncSignArgsContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(V2ParserIdentifier, i)
}

func (s *FuncSignArgsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(V2ParserComma)
}

func (s *FuncSignArgsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(V2ParserComma, i)
}

func (s *FuncSignArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncSignArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncSignArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterFuncSignArgs(s)
	}
}

func (s *FuncSignArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitFuncSignArgs(s)
	}
}

func (s *FuncSignArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitFuncSignArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) FuncSignArgs() (localctx IFuncSignArgsContext) {
	localctx = NewFuncSignArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, V2ParserRULE_funcSignArgs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Type_()
	}
	{
		p.SetState(148)
		p.Match(V2ParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == V2ParserComma {
		{
			p.SetState(149)
			p.Match(V2ParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)
			p.Type_()
		}
		{
			p.SetState(151)
			p.Match(V2ParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Map() antlr.TerminalNode
	Function() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	LSquare() antlr.TerminalNode
	RSquare() antlr.TerminalNode

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) Map() antlr.TerminalNode {
	return s.GetToken(V2ParserMap, 0)
}

func (s *TypeContext) Function() antlr.TerminalNode {
	return s.GetToken(V2ParserFunction, 0)
}

func (s *TypeContext) Identifier() antlr.TerminalNode {
	return s.GetToken(V2ParserIdentifier, 0)
}

func (s *TypeContext) LSquare() antlr.TerminalNode {
	return s.GetToken(V2ParserLSquare, 0)
}

func (s *TypeContext) RSquare() antlr.TerminalNode {
	return s.GetToken(V2ParserRSquare, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitType(s)
	}
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, V2ParserRULE_type)
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case V2ParserMap:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)
			p.Match(V2ParserMap)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case V2ParserFunction:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(159)
			p.Match(V2ParserFunction)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case V2ParserIdentifier:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(160)
			p.Match(V2ParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(161)
				p.Match(V2ParserLSquare)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(162)
				p.Match(V2ParserRSquare)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarContext is an interface to support dynamic dispatch.
type IVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode
	AllDot() []antlr.TerminalNode
	Dot(i int) antlr.TerminalNode

	// IsVarContext differentiates from other interfaces.
	IsVarContext()
}

type VarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarContext() *VarContext {
	var p = new(VarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_var
	return p
}

func InitEmptyVarContext(p *VarContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_var
}

func (*VarContext) IsVarContext() {}

func NewVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarContext {
	var p = new(VarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_var

	return p
}

func (s *VarContext) GetParser() antlr.Parser { return s.parser }

func (s *VarContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(V2ParserIdentifier)
}

func (s *VarContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(V2ParserIdentifier, i)
}

func (s *VarContext) AllDot() []antlr.TerminalNode {
	return s.GetTokens(V2ParserDot)
}

func (s *VarContext) Dot(i int) antlr.TerminalNode {
	return s.GetToken(V2ParserDot, i)
}

func (s *VarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterVar(s)
	}
}

func (s *VarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitVar(s)
	}
}

func (s *VarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitVar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) Var_() (localctx IVarContext) {
	localctx = NewVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, V2ParserRULE_var)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(V2ParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == V2ParserDot {
		{
			p.SetState(168)
			p.Match(V2ParserDot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.Match(V2ParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarsContext is an interface to support dynamic dispatch.
type IVarsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_() IVarContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode
	AllVars() []IVarsContext
	Vars(i int) IVarsContext

	// IsVarsContext differentiates from other interfaces.
	IsVarsContext()
}

type VarsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarsContext() *VarsContext {
	var p = new(VarsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_vars
	return p
}

func InitEmptyVarsContext(p *VarsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_vars
}

func (*VarsContext) IsVarsContext() {}

func NewVarsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsContext {
	var p = new(VarsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_vars

	return p
}

func (s *VarsContext) GetParser() antlr.Parser { return s.parser }

func (s *VarsContext) Var_() IVarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarContext)
}

func (s *VarsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(V2ParserComma)
}

func (s *VarsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(V2ParserComma, i)
}

func (s *VarsContext) AllVars() []IVarsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarsContext); ok {
			len++
		}
	}

	tst := make([]IVarsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarsContext); ok {
			tst[i] = t.(IVarsContext)
			i++
		}
	}

	return tst
}

func (s *VarsContext) Vars(i int) IVarsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarsContext)
}

func (s *VarsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterVars(s)
	}
}

func (s *VarsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitVars(s)
	}
}

func (s *VarsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitVars(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) Vars() (localctx IVarsContext) {
	localctx = NewVarsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, V2ParserRULE_vars)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Var_()
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(176)
				p.Match(V2ParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(177)
				p.Vars()
			}

		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarWithIdxContext is an interface to support dynamic dispatch.
type IVarWithIdxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_() IVarContext
	Indexes() IIndexesContext

	// IsVarWithIdxContext differentiates from other interfaces.
	IsVarWithIdxContext()
}

type VarWithIdxContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarWithIdxContext() *VarWithIdxContext {
	var p = new(VarWithIdxContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_varWithIdx
	return p
}

func InitEmptyVarWithIdxContext(p *VarWithIdxContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_varWithIdx
}

func (*VarWithIdxContext) IsVarWithIdxContext() {}

func NewVarWithIdxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarWithIdxContext {
	var p = new(VarWithIdxContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_varWithIdx

	return p
}

func (s *VarWithIdxContext) GetParser() antlr.Parser { return s.parser }

func (s *VarWithIdxContext) Var_() IVarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarContext)
}

func (s *VarWithIdxContext) Indexes() IIndexesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexesContext)
}

func (s *VarWithIdxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarWithIdxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarWithIdxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterVarWithIdx(s)
	}
}

func (s *VarWithIdxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitVarWithIdx(s)
	}
}

func (s *VarWithIdxContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitVarWithIdx(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) VarWithIdx() (localctx IVarWithIdxContext) {
	localctx = NewVarWithIdxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, V2ParserRULE_varWithIdx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Var_()
	}
	{
		p.SetState(184)
		p.Indexes()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIndexContext is an interface to support dynamic dispatch.
type IIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LSquare() antlr.TerminalNode
	Expr() IExprContext
	RSquare() antlr.TerminalNode

	// IsIndexContext differentiates from other interfaces.
	IsIndexContext()
}

type IndexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexContext() *IndexContext {
	var p = new(IndexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_index
	return p
}

func InitEmptyIndexContext(p *IndexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_index
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) LSquare() antlr.TerminalNode {
	return s.GetToken(V2ParserLSquare, 0)
}

func (s *IndexContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IndexContext) RSquare() antlr.TerminalNode {
	return s.GetToken(V2ParserRSquare, 0)
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterIndex(s)
	}
}

func (s *IndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitIndex(s)
	}
}

func (s *IndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) Index() (localctx IIndexContext) {
	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, V2ParserRULE_index)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		p.Match(V2ParserLSquare)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(187)
		p.expr(0)
	}
	{
		p.SetState(188)
		p.Match(V2ParserRSquare)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIndexesContext is an interface to support dynamic dispatch.
type IIndexesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIndex() []IIndexContext
	Index(i int) IIndexContext

	// IsIndexesContext differentiates from other interfaces.
	IsIndexesContext()
}

type IndexesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexesContext() *IndexesContext {
	var p = new(IndexesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_indexes
	return p
}

func InitEmptyIndexesContext(p *IndexesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_indexes
}

func (*IndexesContext) IsIndexesContext() {}

func NewIndexesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexesContext {
	var p = new(IndexesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_indexes

	return p
}

func (s *IndexesContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexesContext) AllIndex() []IIndexContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIndexContext); ok {
			len++
		}
	}

	tst := make([]IIndexContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIndexContext); ok {
			tst[i] = t.(IIndexContext)
			i++
		}
	}

	return tst
}

func (s *IndexesContext) Index(i int) IIndexContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *IndexesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterIndexes(s)
	}
}

func (s *IndexesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitIndexes(s)
	}
}

func (s *IndexesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitIndexes(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) Indexes() (localctx IIndexesContext) {
	localctx = NewIndexesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, V2ParserRULE_indexes)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == V2ParserLSquare {
		{
			p.SetState(190)
			p.Index()
		}

		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILambdaContext is an interface to support dynamic dispatch.
type ILambdaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LParen() antlr.TerminalNode
	FuncSignArgs() IFuncSignArgsContext
	RParen() antlr.TerminalNode
	CodeBlock() ICodeBlockContext
	Type_() ITypeContext

	// IsLambdaContext differentiates from other interfaces.
	IsLambdaContext()
}

type LambdaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLambdaContext() *LambdaContext {
	var p = new(LambdaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_lambda
	return p
}

func InitEmptyLambdaContext(p *LambdaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_lambda
}

func (*LambdaContext) IsLambdaContext() {}

func NewLambdaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LambdaContext {
	var p = new(LambdaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_lambda

	return p
}

func (s *LambdaContext) GetParser() antlr.Parser { return s.parser }

func (s *LambdaContext) LParen() antlr.TerminalNode {
	return s.GetToken(V2ParserLParen, 0)
}

func (s *LambdaContext) FuncSignArgs() IFuncSignArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncSignArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncSignArgsContext)
}

func (s *LambdaContext) RParen() antlr.TerminalNode {
	return s.GetToken(V2ParserRParen, 0)
}

func (s *LambdaContext) CodeBlock() ICodeBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICodeBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICodeBlockContext)
}

func (s *LambdaContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *LambdaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LambdaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterLambda(s)
	}
}

func (s *LambdaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitLambda(s)
	}
}

func (s *LambdaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitLambda(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) Lambda() (localctx ILambdaContext) {
	localctx = NewLambdaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, V2ParserRULE_lambda)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(V2ParserLParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
		p.FuncSignArgs()
	}
	{
		p.SetState(198)
		p.Match(V2ParserRParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140737891008512) != 0 {
		{
			p.SetState(199)
			p.Type_()
		}

	}
	{
		p.SetState(202)
		p.CodeBlock()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBinaryOperContext is an interface to support dynamic dispatch.
type IBinaryOperContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Equals() antlr.TerminalNode
	NotEq() antlr.TerminalNode
	Greater() antlr.TerminalNode
	Less() antlr.TerminalNode
	GreaterEq() antlr.TerminalNode
	LessEq() antlr.TerminalNode
	Or() antlr.TerminalNode
	And() antlr.TerminalNode
	Pow() antlr.TerminalNode
	Mul() antlr.TerminalNode
	Div() antlr.TerminalNode
	Mod() antlr.TerminalNode
	Add() antlr.TerminalNode
	Sub() antlr.TerminalNode

	// IsBinaryOperContext differentiates from other interfaces.
	IsBinaryOperContext()
}

type BinaryOperContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryOperContext() *BinaryOperContext {
	var p = new(BinaryOperContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_binaryOper
	return p
}

func InitEmptyBinaryOperContext(p *BinaryOperContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_binaryOper
}

func (*BinaryOperContext) IsBinaryOperContext() {}

func NewBinaryOperContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryOperContext {
	var p = new(BinaryOperContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_binaryOper

	return p
}

func (s *BinaryOperContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryOperContext) Equals() antlr.TerminalNode {
	return s.GetToken(V2ParserEquals, 0)
}

func (s *BinaryOperContext) NotEq() antlr.TerminalNode {
	return s.GetToken(V2ParserNotEq, 0)
}

func (s *BinaryOperContext) Greater() antlr.TerminalNode {
	return s.GetToken(V2ParserGreater, 0)
}

func (s *BinaryOperContext) Less() antlr.TerminalNode {
	return s.GetToken(V2ParserLess, 0)
}

func (s *BinaryOperContext) GreaterEq() antlr.TerminalNode {
	return s.GetToken(V2ParserGreaterEq, 0)
}

func (s *BinaryOperContext) LessEq() antlr.TerminalNode {
	return s.GetToken(V2ParserLessEq, 0)
}

func (s *BinaryOperContext) Or() antlr.TerminalNode {
	return s.GetToken(V2ParserOr, 0)
}

func (s *BinaryOperContext) And() antlr.TerminalNode {
	return s.GetToken(V2ParserAnd, 0)
}

func (s *BinaryOperContext) Pow() antlr.TerminalNode {
	return s.GetToken(V2ParserPow, 0)
}

func (s *BinaryOperContext) Mul() antlr.TerminalNode {
	return s.GetToken(V2ParserMul, 0)
}

func (s *BinaryOperContext) Div() antlr.TerminalNode {
	return s.GetToken(V2ParserDiv, 0)
}

func (s *BinaryOperContext) Mod() antlr.TerminalNode {
	return s.GetToken(V2ParserMod, 0)
}

func (s *BinaryOperContext) Add() antlr.TerminalNode {
	return s.GetToken(V2ParserAdd, 0)
}

func (s *BinaryOperContext) Sub() antlr.TerminalNode {
	return s.GetToken(V2ParserSub, 0)
}

func (s *BinaryOperContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryOperContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryOperContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterBinaryOper(s)
	}
}

func (s *BinaryOperContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitBinaryOper(s)
	}
}

func (s *BinaryOperContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitBinaryOper(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) BinaryOper() (localctx IBinaryOperContext) {
	localctx = NewBinaryOperContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, V2ParserRULE_binaryOper)
	var _la int

	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case V2ParserEquals:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(204)
			p.Match(V2ParserEquals)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case V2ParserNotEq:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(205)
			p.Match(V2ParserNotEq)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case V2ParserGreater:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(206)
			p.Match(V2ParserGreater)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case V2ParserLess:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(207)
			p.Match(V2ParserLess)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case V2ParserGreaterEq:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(208)
			p.Match(V2ParserGreaterEq)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case V2ParserLessEq:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(209)
			p.Match(V2ParserLessEq)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case V2ParserOr, V2ParserAnd:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(210)
			_la = p.GetTokenStream().LA(1)

			if !(_la == V2ParserOr || _la == V2ParserAnd) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case V2ParserPow:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(211)
			p.Match(V2ParserPow)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case V2ParserMul, V2ParserDiv:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(212)
			_la = p.GetTokenStream().LA(1)

			if !(_la == V2ParserMul || _la == V2ParserDiv) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case V2ParserMod:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(213)
			p.Match(V2ParserMod)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case V2ParserAdd, V2ParserSub:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(214)
			_la = p.GetTokenStream().LA(1)

			if !(_la == V2ParserAdd || _la == V2ParserSub) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryOperContext is an interface to support dynamic dispatch.
type IUnaryOperContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Add() antlr.TerminalNode
	Sub() antlr.TerminalNode
	Not() antlr.TerminalNode

	// IsUnaryOperContext differentiates from other interfaces.
	IsUnaryOperContext()
}

type UnaryOperContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOperContext() *UnaryOperContext {
	var p = new(UnaryOperContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_unaryOper
	return p
}

func InitEmptyUnaryOperContext(p *UnaryOperContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_unaryOper
}

func (*UnaryOperContext) IsUnaryOperContext() {}

func NewUnaryOperContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOperContext {
	var p = new(UnaryOperContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_unaryOper

	return p
}

func (s *UnaryOperContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOperContext) Add() antlr.TerminalNode {
	return s.GetToken(V2ParserAdd, 0)
}

func (s *UnaryOperContext) Sub() antlr.TerminalNode {
	return s.GetToken(V2ParserSub, 0)
}

func (s *UnaryOperContext) Not() antlr.TerminalNode {
	return s.GetToken(V2ParserNot, 0)
}

func (s *UnaryOperContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOperContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOperContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterUnaryOper(s)
	}
}

func (s *UnaryOperContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitUnaryOper(s)
	}
}

func (s *UnaryOperContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitUnaryOper(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) UnaryOper() (localctx IUnaryOperContext) {
	localctx = NewUnaryOperContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, V2ParserRULE_unaryOper)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(217)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&35184378380288) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FuncCall() IFuncCallContext
	UnaryOper() IUnaryOperContext
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	Literal() ILiteralContext
	LParen() antlr.TerminalNode
	RParen() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	BinaryOper() IBinaryOperContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) FuncCall() IFuncCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallContext)
}

func (s *ExprContext) UnaryOper() IUnaryOperContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryOperContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryOperContext)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ExprContext) LParen() antlr.TerminalNode {
	return s.GetToken(V2ParserLParen, 0)
}

func (s *ExprContext) RParen() antlr.TerminalNode {
	return s.GetToken(V2ParserRParen, 0)
}

func (s *ExprContext) Identifier() antlr.TerminalNode {
	return s.GetToken(V2ParserIdentifier, 0)
}

func (s *ExprContext) BinaryOper() IBinaryOperContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryOperContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryOperContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *V2Parser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, V2ParserRULE_expr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(220)
			p.FuncCall()
		}

	case 2:
		{
			p.SetState(221)
			p.UnaryOper()
		}
		{
			p.SetState(222)
			p.expr(5)
		}

	case 3:
		{
			p.SetState(224)
			p.Literal()
		}

	case 4:
		{
			p.SetState(225)
			p.Match(V2ParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(226)
			p.expr(0)
		}
		{
			p.SetState(227)
			p.Match(V2ParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		{
			p.SetState(229)
			p.Match(V2ParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, V2ParserRULE_expr)
			p.SetState(232)

			if !(p.Precpred(p.GetParserRuleContext(), 4)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				goto errorExit
			}
			{
				p.SetState(233)
				p.BinaryOper()
			}
			{
				p.SetState(234)
				p.expr(5)
			}

		}
		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprWithLambdaContext is an interface to support dynamic dispatch.
type IExprWithLambdaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Lambda() ILambdaContext
	Expr() IExprContext

	// IsExprWithLambdaContext differentiates from other interfaces.
	IsExprWithLambdaContext()
}

type ExprWithLambdaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprWithLambdaContext() *ExprWithLambdaContext {
	var p = new(ExprWithLambdaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_exprWithLambda
	return p
}

func InitEmptyExprWithLambdaContext(p *ExprWithLambdaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_exprWithLambda
}

func (*ExprWithLambdaContext) IsExprWithLambdaContext() {}

func NewExprWithLambdaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprWithLambdaContext {
	var p = new(ExprWithLambdaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_exprWithLambda

	return p
}

func (s *ExprWithLambdaContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprWithLambdaContext) Lambda() ILambdaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILambdaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILambdaContext)
}

func (s *ExprWithLambdaContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprWithLambdaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprWithLambdaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprWithLambdaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterExprWithLambda(s)
	}
}

func (s *ExprWithLambdaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitExprWithLambda(s)
	}
}

func (s *ExprWithLambdaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitExprWithLambda(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) ExprWithLambda() (localctx IExprWithLambdaContext) {
	localctx = NewExprWithLambdaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, V2ParserRULE_exprWithLambda)
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(241)
			p.Lambda()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(242)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprOrAssignContext is an interface to support dynamic dispatch.
type IExprOrAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	AssgnStmt() IAssgnStmtContext

	// IsExprOrAssignContext differentiates from other interfaces.
	IsExprOrAssignContext()
}

type ExprOrAssignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprOrAssignContext() *ExprOrAssignContext {
	var p = new(ExprOrAssignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_exprOrAssign
	return p
}

func InitEmptyExprOrAssignContext(p *ExprOrAssignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_exprOrAssign
}

func (*ExprOrAssignContext) IsExprOrAssignContext() {}

func NewExprOrAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprOrAssignContext {
	var p = new(ExprOrAssignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_exprOrAssign

	return p
}

func (s *ExprOrAssignContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprOrAssignContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprOrAssignContext) AssgnStmt() IAssgnStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssgnStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssgnStmtContext)
}

func (s *ExprOrAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprOrAssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprOrAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterExprOrAssign(s)
	}
}

func (s *ExprOrAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitExprOrAssign(s)
	}
}

func (s *ExprOrAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitExprOrAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) ExprOrAssign() (localctx IExprOrAssignContext) {
	localctx = NewExprOrAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, V2ParserRULE_exprOrAssign)
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(245)
			p.expr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(246)
			p.AssgnStmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncCallContext is an interface to support dynamic dispatch.
type IFuncCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_() IVarContext
	LParen() antlr.TerminalNode
	RParen() antlr.TerminalNode
	FuncCallArgs() IFuncCallArgsContext

	// IsFuncCallContext differentiates from other interfaces.
	IsFuncCallContext()
}

type FuncCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncCallContext() *FuncCallContext {
	var p = new(FuncCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_funcCall
	return p
}

func InitEmptyFuncCallContext(p *FuncCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_funcCall
}

func (*FuncCallContext) IsFuncCallContext() {}

func NewFuncCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncCallContext {
	var p = new(FuncCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_funcCall

	return p
}

func (s *FuncCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncCallContext) Var_() IVarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarContext)
}

func (s *FuncCallContext) LParen() antlr.TerminalNode {
	return s.GetToken(V2ParserLParen, 0)
}

func (s *FuncCallContext) RParen() antlr.TerminalNode {
	return s.GetToken(V2ParserRParen, 0)
}

func (s *FuncCallContext) FuncCallArgs() IFuncCallArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallArgsContext)
}

func (s *FuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterFuncCall(s)
	}
}

func (s *FuncCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitFuncCall(s)
	}
}

func (s *FuncCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitFuncCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) FuncCall() (localctx IFuncCallContext) {
	localctx = NewFuncCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, V2ParserRULE_funcCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		p.Var_()
	}
	{
		p.SetState(250)
		p.Match(V2ParserLParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&210007129849896) != 0 {
		{
			p.SetState(251)
			p.FuncCallArgs()
		}

	}
	{
		p.SetState(254)
		p.Match(V2ParserRParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncCallArgsContext is an interface to support dynamic dispatch.
type IFuncCallArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsFuncCallArgsContext differentiates from other interfaces.
	IsFuncCallArgsContext()
}

type FuncCallArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncCallArgsContext() *FuncCallArgsContext {
	var p = new(FuncCallArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_funcCallArgs
	return p
}

func InitEmptyFuncCallArgsContext(p *FuncCallArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_funcCallArgs
}

func (*FuncCallArgsContext) IsFuncCallArgsContext() {}

func NewFuncCallArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncCallArgsContext {
	var p = new(FuncCallArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_funcCallArgs

	return p
}

func (s *FuncCallArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncCallArgsContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *FuncCallArgsContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FuncCallArgsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(V2ParserComma)
}

func (s *FuncCallArgsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(V2ParserComma, i)
}

func (s *FuncCallArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncCallArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterFuncCallArgs(s)
	}
}

func (s *FuncCallArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitFuncCallArgs(s)
	}
}

func (s *FuncCallArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitFuncCallArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) FuncCallArgs() (localctx IFuncCallArgsContext) {
	localctx = NewFuncCallArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, V2ParserRULE_funcCallArgs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.expr(0)
	}
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == V2ParserComma {
		{
			p.SetState(257)
			p.Match(V2ParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(258)
			p.expr(0)
		}

		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtWithSepContext is an interface to support dynamic dispatch.
type IStmtWithSepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Stmt() IStmtContext
	AllSep() []ISepContext
	Sep(i int) ISepContext

	// IsStmtWithSepContext differentiates from other interfaces.
	IsStmtWithSepContext()
}

type StmtWithSepContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtWithSepContext() *StmtWithSepContext {
	var p = new(StmtWithSepContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_stmtWithSep
	return p
}

func InitEmptyStmtWithSepContext(p *StmtWithSepContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_stmtWithSep
}

func (*StmtWithSepContext) IsStmtWithSepContext() {}

func NewStmtWithSepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtWithSepContext {
	var p = new(StmtWithSepContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_stmtWithSep

	return p
}

func (s *StmtWithSepContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtWithSepContext) Stmt() IStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *StmtWithSepContext) AllSep() []ISepContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISepContext); ok {
			len++
		}
	}

	tst := make([]ISepContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISepContext); ok {
			tst[i] = t.(ISepContext)
			i++
		}
	}

	return tst
}

func (s *StmtWithSepContext) Sep(i int) ISepContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISepContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISepContext)
}

func (s *StmtWithSepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtWithSepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtWithSepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterStmtWithSep(s)
	}
}

func (s *StmtWithSepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitStmtWithSep(s)
	}
}

func (s *StmtWithSepContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitStmtWithSep(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) StmtWithSep() (localctx IStmtWithSepContext) {
	localctx = NewStmtWithSepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, V2ParserRULE_stmtWithSep)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
		p.Stmt()
	}
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == V2ParserSemi || _la == V2ParserNewLine {
		{
			p.SetState(265)
			p.Sep()
		}

		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOpenStmtContext is an interface to support dynamic dispatch.
type IOpenStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Open() antlr.TerminalNode
	StringLiteral() antlr.TerminalNode
	As() antlr.TerminalNode
	Identifier() antlr.TerminalNode

	// IsOpenStmtContext differentiates from other interfaces.
	IsOpenStmtContext()
}

type OpenStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpenStmtContext() *OpenStmtContext {
	var p = new(OpenStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_openStmt
	return p
}

func InitEmptyOpenStmtContext(p *OpenStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_openStmt
}

func (*OpenStmtContext) IsOpenStmtContext() {}

func NewOpenStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpenStmtContext {
	var p = new(OpenStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_openStmt

	return p
}

func (s *OpenStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *OpenStmtContext) Open() antlr.TerminalNode {
	return s.GetToken(V2ParserOpen, 0)
}

func (s *OpenStmtContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(V2ParserStringLiteral, 0)
}

func (s *OpenStmtContext) As() antlr.TerminalNode {
	return s.GetToken(V2ParserAs, 0)
}

func (s *OpenStmtContext) Identifier() antlr.TerminalNode {
	return s.GetToken(V2ParserIdentifier, 0)
}

func (s *OpenStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpenStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpenStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterOpenStmt(s)
	}
}

func (s *OpenStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitOpenStmt(s)
	}
}

func (s *OpenStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitOpenStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) OpenStmt() (localctx IOpenStmtContext) {
	localctx = NewOpenStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, V2ParserRULE_openStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(271)
		p.Match(V2ParserOpen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(272)
		p.Match(V2ParserStringLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == V2ParserAs {
		{
			p.SetState(273)
			p.Match(V2ParserAs)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(274)
			p.Match(V2ParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	True() antlr.TerminalNode
	False() antlr.TerminalNode
	IntegerLiteral() antlr.TerminalNode
	NumberLiteral() antlr.TerminalNode
	StringLiteral() antlr.TerminalNode
	ArrayInitializer() IArrayInitializerContext
	MapInitializer() IMapInitializerContext
	StructInitializer() IStructInitializerContext

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) True() antlr.TerminalNode {
	return s.GetToken(V2ParserTrue, 0)
}

func (s *LiteralContext) False() antlr.TerminalNode {
	return s.GetToken(V2ParserFalse, 0)
}

func (s *LiteralContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(V2ParserIntegerLiteral, 0)
}

func (s *LiteralContext) NumberLiteral() antlr.TerminalNode {
	return s.GetToken(V2ParserNumberLiteral, 0)
}

func (s *LiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(V2ParserStringLiteral, 0)
}

func (s *LiteralContext) ArrayInitializer() IArrayInitializerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayInitializerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayInitializerContext)
}

func (s *LiteralContext) MapInitializer() IMapInitializerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapInitializerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapInitializerContext)
}

func (s *LiteralContext) StructInitializer() IStructInitializerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructInitializerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructInitializerContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, V2ParserRULE_literal)
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(277)
			p.Match(V2ParserTrue)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(278)
			p.Match(V2ParserFalse)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(279)
			p.Match(V2ParserIntegerLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(280)
			p.Match(V2ParserNumberLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(281)
			p.Match(V2ParserStringLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(282)
			p.ArrayInitializer()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(283)
			p.MapInitializer()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(284)
			p.StructInitializer()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralWithLambdaContext is an interface to support dynamic dispatch.
type ILiteralWithLambdaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal() ILiteralContext
	Lambda() ILambdaContext

	// IsLiteralWithLambdaContext differentiates from other interfaces.
	IsLiteralWithLambdaContext()
}

type LiteralWithLambdaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralWithLambdaContext() *LiteralWithLambdaContext {
	var p = new(LiteralWithLambdaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_literalWithLambda
	return p
}

func InitEmptyLiteralWithLambdaContext(p *LiteralWithLambdaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_literalWithLambda
}

func (*LiteralWithLambdaContext) IsLiteralWithLambdaContext() {}

func NewLiteralWithLambdaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralWithLambdaContext {
	var p = new(LiteralWithLambdaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_literalWithLambda

	return p
}

func (s *LiteralWithLambdaContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralWithLambdaContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralWithLambdaContext) Lambda() ILambdaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILambdaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILambdaContext)
}

func (s *LiteralWithLambdaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralWithLambdaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralWithLambdaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterLiteralWithLambda(s)
	}
}

func (s *LiteralWithLambdaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitLiteralWithLambda(s)
	}
}

func (s *LiteralWithLambdaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitLiteralWithLambda(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) LiteralWithLambda() (localctx ILiteralWithLambdaContext) {
	localctx = NewLiteralWithLambdaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, V2ParserRULE_literalWithLambda)
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case V2ParserLSquare, V2ParserMap, V2ParserFunction, V2ParserTrue, V2ParserFalse, V2ParserIntegerLiteral, V2ParserNumberLiteral, V2ParserStringLiteral, V2ParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(287)
			p.Literal()
		}

	case V2ParserLParen:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(288)
			p.Lambda()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayInitializerContext is an interface to support dynamic dispatch.
type IArrayInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LSquare() antlr.TerminalNode
	RSquare() antlr.TerminalNode
	Type_() ITypeContext
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsArrayInitializerContext differentiates from other interfaces.
	IsArrayInitializerContext()
}

type ArrayInitializerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayInitializerContext() *ArrayInitializerContext {
	var p = new(ArrayInitializerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_arrayInitializer
	return p
}

func InitEmptyArrayInitializerContext(p *ArrayInitializerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_arrayInitializer
}

func (*ArrayInitializerContext) IsArrayInitializerContext() {}

func NewArrayInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayInitializerContext {
	var p = new(ArrayInitializerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_arrayInitializer

	return p
}

func (s *ArrayInitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayInitializerContext) LSquare() antlr.TerminalNode {
	return s.GetToken(V2ParserLSquare, 0)
}

func (s *ArrayInitializerContext) RSquare() antlr.TerminalNode {
	return s.GetToken(V2ParserRSquare, 0)
}

func (s *ArrayInitializerContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ArrayInitializerContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ArrayInitializerContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArrayInitializerContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(V2ParserComma)
}

func (s *ArrayInitializerContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(V2ParserComma, i)
}

func (s *ArrayInitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayInitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayInitializerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterArrayInitializer(s)
	}
}

func (s *ArrayInitializerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitArrayInitializer(s)
	}
}

func (s *ArrayInitializerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitArrayInitializer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) ArrayInitializer() (localctx IArrayInitializerContext) {
	localctx = NewArrayInitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, V2ParserRULE_arrayInitializer)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140737891008512) != 0 {
		{
			p.SetState(291)
			p.Type_()
		}

	}
	{
		p.SetState(294)
		p.Match(V2ParserLSquare)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&210007129849896) != 0 {
		{
			p.SetState(295)
			p.expr(0)
		}
		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == V2ParserComma {
			{
				p.SetState(296)
				p.Match(V2ParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(304)
		p.Match(V2ParserRSquare)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMapInitializerContext is an interface to support dynamic dispatch.
type IMapInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Map() antlr.TerminalNode
	LBrack() antlr.TerminalNode
	RBrack() antlr.TerminalNode
	AllMapPair() []IMapPairContext
	MapPair(i int) IMapPairContext

	// IsMapInitializerContext differentiates from other interfaces.
	IsMapInitializerContext()
}

type MapInitializerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapInitializerContext() *MapInitializerContext {
	var p = new(MapInitializerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_mapInitializer
	return p
}

func InitEmptyMapInitializerContext(p *MapInitializerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_mapInitializer
}

func (*MapInitializerContext) IsMapInitializerContext() {}

func NewMapInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapInitializerContext {
	var p = new(MapInitializerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_mapInitializer

	return p
}

func (s *MapInitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *MapInitializerContext) Map() antlr.TerminalNode {
	return s.GetToken(V2ParserMap, 0)
}

func (s *MapInitializerContext) LBrack() antlr.TerminalNode {
	return s.GetToken(V2ParserLBrack, 0)
}

func (s *MapInitializerContext) RBrack() antlr.TerminalNode {
	return s.GetToken(V2ParserRBrack, 0)
}

func (s *MapInitializerContext) AllMapPair() []IMapPairContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMapPairContext); ok {
			len++
		}
	}

	tst := make([]IMapPairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMapPairContext); ok {
			tst[i] = t.(IMapPairContext)
			i++
		}
	}

	return tst
}

func (s *MapInitializerContext) MapPair(i int) IMapPairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapPairContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapPairContext)
}

func (s *MapInitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapInitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapInitializerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterMapInitializer(s)
	}
}

func (s *MapInitializerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitMapInitializer(s)
	}
}

func (s *MapInitializerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitMapInitializer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) MapInitializer() (localctx IMapInitializerContext) {
	localctx = NewMapInitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, V2ParserRULE_mapInitializer)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(306)
		p.Match(V2ParserMap)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(307)
		p.Match(V2ParserLBrack)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&210007129849896) != 0 {
		{
			p.SetState(308)
			p.MapPair()
		}

		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(314)
		p.Match(V2ParserRBrack)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMapPairContext is an interface to support dynamic dispatch.
type IMapPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LParen() antlr.TerminalNode
	Expr() IExprContext
	Comma() antlr.TerminalNode
	ExprWithLambda() IExprWithLambdaContext
	RParen() antlr.TerminalNode
	To() antlr.TerminalNode

	// IsMapPairContext differentiates from other interfaces.
	IsMapPairContext()
}

type MapPairContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapPairContext() *MapPairContext {
	var p = new(MapPairContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_mapPair
	return p
}

func InitEmptyMapPairContext(p *MapPairContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_mapPair
}

func (*MapPairContext) IsMapPairContext() {}

func NewMapPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapPairContext {
	var p = new(MapPairContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_mapPair

	return p
}

func (s *MapPairContext) GetParser() antlr.Parser { return s.parser }

func (s *MapPairContext) LParen() antlr.TerminalNode {
	return s.GetToken(V2ParserLParen, 0)
}

func (s *MapPairContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MapPairContext) Comma() antlr.TerminalNode {
	return s.GetToken(V2ParserComma, 0)
}

func (s *MapPairContext) ExprWithLambda() IExprWithLambdaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprWithLambdaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprWithLambdaContext)
}

func (s *MapPairContext) RParen() antlr.TerminalNode {
	return s.GetToken(V2ParserRParen, 0)
}

func (s *MapPairContext) To() antlr.TerminalNode {
	return s.GetToken(V2ParserTo, 0)
}

func (s *MapPairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapPairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapPairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterMapPair(s)
	}
}

func (s *MapPairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitMapPair(s)
	}
}

func (s *MapPairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitMapPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) MapPair() (localctx IMapPairContext) {
	localctx = NewMapPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, V2ParserRULE_mapPair)
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(316)
			p.Match(V2ParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(317)
			p.expr(0)
		}
		{
			p.SetState(318)
			p.Match(V2ParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(319)
			p.ExprWithLambda()
		}
		{
			p.SetState(320)
			p.Match(V2ParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(322)
			p.expr(0)
		}
		{
			p.SetState(323)
			p.Match(V2ParserTo)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(324)
			p.ExprWithLambda()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructInitializerContext is an interface to support dynamic dispatch.
type IStructInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	LBrack() antlr.TerminalNode
	RBrack() antlr.TerminalNode
	AllStructElementInitializer() []IStructElementInitializerContext
	StructElementInitializer(i int) IStructElementInitializerContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsStructInitializerContext differentiates from other interfaces.
	IsStructInitializerContext()
}

type StructInitializerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructInitializerContext() *StructInitializerContext {
	var p = new(StructInitializerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_structInitializer
	return p
}

func InitEmptyStructInitializerContext(p *StructInitializerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_structInitializer
}

func (*StructInitializerContext) IsStructInitializerContext() {}

func NewStructInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructInitializerContext {
	var p = new(StructInitializerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_structInitializer

	return p
}

func (s *StructInitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *StructInitializerContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *StructInitializerContext) LBrack() antlr.TerminalNode {
	return s.GetToken(V2ParserLBrack, 0)
}

func (s *StructInitializerContext) RBrack() antlr.TerminalNode {
	return s.GetToken(V2ParserRBrack, 0)
}

func (s *StructInitializerContext) AllStructElementInitializer() []IStructElementInitializerContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructElementInitializerContext); ok {
			len++
		}
	}

	tst := make([]IStructElementInitializerContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructElementInitializerContext); ok {
			tst[i] = t.(IStructElementInitializerContext)
			i++
		}
	}

	return tst
}

func (s *StructInitializerContext) StructElementInitializer(i int) IStructElementInitializerContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructElementInitializerContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructElementInitializerContext)
}

func (s *StructInitializerContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(V2ParserComma)
}

func (s *StructInitializerContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(V2ParserComma, i)
}

func (s *StructInitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructInitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructInitializerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterStructInitializer(s)
	}
}

func (s *StructInitializerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitStructInitializer(s)
	}
}

func (s *StructInitializerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitStructInitializer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) StructInitializer() (localctx IStructInitializerContext) {
	localctx = NewStructInitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, V2ParserRULE_structInitializer)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(328)
		p.Type_()
	}
	{
		p.SetState(329)
		p.Match(V2ParserLBrack)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == V2ParserIdentifier {
		{
			p.SetState(330)
			p.StructElementInitializer()
		}
		p.SetState(332)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == V2ParserComma {
			{
				p.SetState(331)
				p.Match(V2ParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

		p.SetState(338)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(339)
		p.Match(V2ParserRBrack)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructElementInitializerContext is an interface to support dynamic dispatch.
type IStructElementInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Comma() antlr.TerminalNode
	Expr() IExprContext

	// IsStructElementInitializerContext differentiates from other interfaces.
	IsStructElementInitializerContext()
}

type StructElementInitializerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructElementInitializerContext() *StructElementInitializerContext {
	var p = new(StructElementInitializerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_structElementInitializer
	return p
}

func InitEmptyStructElementInitializerContext(p *StructElementInitializerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_structElementInitializer
}

func (*StructElementInitializerContext) IsStructElementInitializerContext() {}

func NewStructElementInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructElementInitializerContext {
	var p = new(StructElementInitializerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_structElementInitializer

	return p
}

func (s *StructElementInitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *StructElementInitializerContext) Identifier() antlr.TerminalNode {
	return s.GetToken(V2ParserIdentifier, 0)
}

func (s *StructElementInitializerContext) Comma() antlr.TerminalNode {
	return s.GetToken(V2ParserComma, 0)
}

func (s *StructElementInitializerContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StructElementInitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructElementInitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructElementInitializerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterStructElementInitializer(s)
	}
}

func (s *StructElementInitializerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitStructElementInitializer(s)
	}
}

func (s *StructElementInitializerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitStructElementInitializer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) StructElementInitializer() (localctx IStructElementInitializerContext) {
	localctx = NewStructElementInitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, V2ParserRULE_structElementInitializer)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(341)
		p.Match(V2ParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(342)
		p.Match(V2ParserComma)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(343)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DeclareStmt() IDeclareStmtContext
	AssgnStmt() IAssgnStmtContext
	JumpStmt() IJumpStmtContext
	IfStmt() IIfStmtContext
	LoopStmt() ILoopStmtContext
	MatchStmt() IMatchStmtContext
	CodeBlock() ICodeBlockContext
	FuncCall() IFuncCallContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) DeclareStmt() IDeclareStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclareStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclareStmtContext)
}

func (s *StmtContext) AssgnStmt() IAssgnStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssgnStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssgnStmtContext)
}

func (s *StmtContext) JumpStmt() IJumpStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJumpStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJumpStmtContext)
}

func (s *StmtContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *StmtContext) LoopStmt() ILoopStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopStmtContext)
}

func (s *StmtContext) MatchStmt() IMatchStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatchStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatchStmtContext)
}

func (s *StmtContext) CodeBlock() ICodeBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICodeBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICodeBlockContext)
}

func (s *StmtContext) FuncCall() IFuncCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, V2ParserRULE_stmt)
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(345)
			p.DeclareStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(346)
			p.AssgnStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(347)
			p.JumpStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(348)
			p.IfStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(349)
			p.LoopStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(350)
			p.MatchStmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(351)
			p.CodeBlock()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(352)
			p.FuncCall()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclareStmtContext is an interface to support dynamic dispatch.
type IDeclareStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	Vars() IVarsContext

	// IsDeclareStmtContext differentiates from other interfaces.
	IsDeclareStmtContext()
}

type DeclareStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclareStmtContext() *DeclareStmtContext {
	var p = new(DeclareStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_declareStmt
	return p
}

func InitEmptyDeclareStmtContext(p *DeclareStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_declareStmt
}

func (*DeclareStmtContext) IsDeclareStmtContext() {}

func NewDeclareStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclareStmtContext {
	var p = new(DeclareStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_declareStmt

	return p
}

func (s *DeclareStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclareStmtContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeclareStmtContext) Vars() IVarsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarsContext)
}

func (s *DeclareStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclareStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclareStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterDeclareStmt(s)
	}
}

func (s *DeclareStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitDeclareStmt(s)
	}
}

func (s *DeclareStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitDeclareStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) DeclareStmt() (localctx IDeclareStmtContext) {
	localctx = NewDeclareStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, V2ParserRULE_declareStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		p.Type_()
	}
	{
		p.SetState(356)
		p.Vars()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssgnStmtContext is an interface to support dynamic dispatch.
type IAssgnStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarWithIdx() IVarWithIdxContext
	Assign() antlr.TerminalNode
	ExprWithLambda() IExprWithLambdaContext
	Type_() ITypeContext

	// IsAssgnStmtContext differentiates from other interfaces.
	IsAssgnStmtContext()
}

type AssgnStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssgnStmtContext() *AssgnStmtContext {
	var p = new(AssgnStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_assgnStmt
	return p
}

func InitEmptyAssgnStmtContext(p *AssgnStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_assgnStmt
}

func (*AssgnStmtContext) IsAssgnStmtContext() {}

func NewAssgnStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssgnStmtContext {
	var p = new(AssgnStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_assgnStmt

	return p
}

func (s *AssgnStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AssgnStmtContext) VarWithIdx() IVarWithIdxContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarWithIdxContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarWithIdxContext)
}

func (s *AssgnStmtContext) Assign() antlr.TerminalNode {
	return s.GetToken(V2ParserAssign, 0)
}

func (s *AssgnStmtContext) ExprWithLambda() IExprWithLambdaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprWithLambdaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprWithLambdaContext)
}

func (s *AssgnStmtContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *AssgnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssgnStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssgnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterAssgnStmt(s)
	}
}

func (s *AssgnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitAssgnStmt(s)
	}
}

func (s *AssgnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitAssgnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) AssgnStmt() (localctx IAssgnStmtContext) {
	localctx = NewAssgnStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, V2ParserRULE_assgnStmt)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(359)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(358)
			p.Type_()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(361)
		p.VarWithIdx()
	}
	{
		p.SetState(362)
		p.Match(V2ParserAssign)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(363)
		p.ExprWithLambda()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	If() antlr.TerminalNode
	LParen() antlr.TerminalNode
	Expr() IExprContext
	RParen() antlr.TerminalNode
	AllCodeBlock() []ICodeBlockContext
	CodeBlock(i int) ICodeBlockContext
	Else() antlr.TerminalNode

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_ifStmt
	return p
}

func InitEmptyIfStmtContext(p *IfStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_ifStmt
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) If() antlr.TerminalNode {
	return s.GetToken(V2ParserIf, 0)
}

func (s *IfStmtContext) LParen() antlr.TerminalNode {
	return s.GetToken(V2ParserLParen, 0)
}

func (s *IfStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfStmtContext) RParen() antlr.TerminalNode {
	return s.GetToken(V2ParserRParen, 0)
}

func (s *IfStmtContext) AllCodeBlock() []ICodeBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICodeBlockContext); ok {
			len++
		}
	}

	tst := make([]ICodeBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICodeBlockContext); ok {
			tst[i] = t.(ICodeBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) CodeBlock(i int) ICodeBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICodeBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICodeBlockContext)
}

func (s *IfStmtContext) Else() antlr.TerminalNode {
	return s.GetToken(V2ParserElse, 0)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, V2ParserRULE_ifStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(365)
		p.Match(V2ParserIf)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(366)
		p.Match(V2ParserLParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(367)
		p.expr(0)
	}
	{
		p.SetState(368)
		p.Match(V2ParserRParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(369)
		p.CodeBlock()
	}
	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == V2ParserElse {
		{
			p.SetState(370)
			p.Match(V2ParserElse)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(371)
			p.CodeBlock()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILoopStmtContext is an interface to support dynamic dispatch.
type ILoopStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	For() antlr.TerminalNode
	LParen() antlr.TerminalNode
	AllSemi() []antlr.TerminalNode
	Semi(i int) antlr.TerminalNode
	RParen() antlr.TerminalNode
	CodeBlock() ICodeBlockContext
	ExprOrAssign() IExprOrAssignContext
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsLoopStmtContext differentiates from other interfaces.
	IsLoopStmtContext()
}

type LoopStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopStmtContext() *LoopStmtContext {
	var p = new(LoopStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_loopStmt
	return p
}

func InitEmptyLoopStmtContext(p *LoopStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_loopStmt
}

func (*LoopStmtContext) IsLoopStmtContext() {}

func NewLoopStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopStmtContext {
	var p = new(LoopStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_loopStmt

	return p
}

func (s *LoopStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopStmtContext) For() antlr.TerminalNode {
	return s.GetToken(V2ParserFor, 0)
}

func (s *LoopStmtContext) LParen() antlr.TerminalNode {
	return s.GetToken(V2ParserLParen, 0)
}

func (s *LoopStmtContext) AllSemi() []antlr.TerminalNode {
	return s.GetTokens(V2ParserSemi)
}

func (s *LoopStmtContext) Semi(i int) antlr.TerminalNode {
	return s.GetToken(V2ParserSemi, i)
}

func (s *LoopStmtContext) RParen() antlr.TerminalNode {
	return s.GetToken(V2ParserRParen, 0)
}

func (s *LoopStmtContext) CodeBlock() ICodeBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICodeBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICodeBlockContext)
}

func (s *LoopStmtContext) ExprOrAssign() IExprOrAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprOrAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprOrAssignContext)
}

func (s *LoopStmtContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *LoopStmtContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LoopStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterLoopStmt(s)
	}
}

func (s *LoopStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitLoopStmt(s)
	}
}

func (s *LoopStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitLoopStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) LoopStmt() (localctx ILoopStmtContext) {
	localctx = NewLoopStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, V2ParserRULE_loopStmt)
	var _la int

	p.SetState(396)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(374)
			p.Match(V2ParserFor)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(375)
			p.Match(V2ParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(377)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&210007129849896) != 0 {
			{
				p.SetState(376)
				p.ExprOrAssign()
			}

		}
		{
			p.SetState(379)
			p.Match(V2ParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(381)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&210007129849896) != 0 {
			{
				p.SetState(380)
				p.expr(0)
			}

		}
		{
			p.SetState(383)
			p.Match(V2ParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(385)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&210007129849896) != 0 {
			{
				p.SetState(384)
				p.expr(0)
			}

		}
		{
			p.SetState(387)
			p.Match(V2ParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(388)
			p.CodeBlock()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(389)
			p.Match(V2ParserFor)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(390)
			p.Match(V2ParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(392)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&210007129849896) != 0 {
			{
				p.SetState(391)
				p.expr(0)
			}

		}
		{
			p.SetState(394)
			p.Match(V2ParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(395)
			p.CodeBlock()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMatchStmtContext is an interface to support dynamic dispatch.
type IMatchStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Match() antlr.TerminalNode
	LParen() antlr.TerminalNode
	Expr() IExprContext
	RParen() antlr.TerminalNode
	LBrack() antlr.TerminalNode
	RBrack() antlr.TerminalNode
	AllMatchCase() []IMatchCaseContext
	MatchCase(i int) IMatchCaseContext

	// IsMatchStmtContext differentiates from other interfaces.
	IsMatchStmtContext()
}

type MatchStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatchStmtContext() *MatchStmtContext {
	var p = new(MatchStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_matchStmt
	return p
}

func InitEmptyMatchStmtContext(p *MatchStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_matchStmt
}

func (*MatchStmtContext) IsMatchStmtContext() {}

func NewMatchStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchStmtContext {
	var p = new(MatchStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_matchStmt

	return p
}

func (s *MatchStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchStmtContext) Match() antlr.TerminalNode {
	return s.GetToken(V2ParserMatch, 0)
}

func (s *MatchStmtContext) LParen() antlr.TerminalNode {
	return s.GetToken(V2ParserLParen, 0)
}

func (s *MatchStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MatchStmtContext) RParen() antlr.TerminalNode {
	return s.GetToken(V2ParserRParen, 0)
}

func (s *MatchStmtContext) LBrack() antlr.TerminalNode {
	return s.GetToken(V2ParserLBrack, 0)
}

func (s *MatchStmtContext) RBrack() antlr.TerminalNode {
	return s.GetToken(V2ParserRBrack, 0)
}

func (s *MatchStmtContext) AllMatchCase() []IMatchCaseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMatchCaseContext); ok {
			len++
		}
	}

	tst := make([]IMatchCaseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMatchCaseContext); ok {
			tst[i] = t.(IMatchCaseContext)
			i++
		}
	}

	return tst
}

func (s *MatchStmtContext) MatchCase(i int) IMatchCaseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatchCaseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatchCaseContext)
}

func (s *MatchStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatchStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterMatchStmt(s)
	}
}

func (s *MatchStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitMatchStmt(s)
	}
}

func (s *MatchStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitMatchStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) MatchStmt() (localctx IMatchStmtContext) {
	localctx = NewMatchStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, V2ParserRULE_matchStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(398)
		p.Match(V2ParserMatch)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(399)
		p.Match(V2ParserLParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(400)
		p.expr(0)
	}
	{
		p.SetState(401)
		p.Match(V2ParserRParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(402)
		p.Match(V2ParserLBrack)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(406)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == V2ParserCase || _la == V2ParserDefault {
		{
			p.SetState(403)
			p.MatchCase()
		}

		p.SetState(408)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(409)
		p.Match(V2ParserRBrack)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMatchCaseContext is an interface to support dynamic dispatch.
type IMatchCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Case() antlr.TerminalNode
	Expr() IExprContext
	Col() antlr.TerminalNode
	CodeBlock() ICodeBlockContext
	Default() antlr.TerminalNode

	// IsMatchCaseContext differentiates from other interfaces.
	IsMatchCaseContext()
}

type MatchCaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatchCaseContext() *MatchCaseContext {
	var p = new(MatchCaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_matchCase
	return p
}

func InitEmptyMatchCaseContext(p *MatchCaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_matchCase
}

func (*MatchCaseContext) IsMatchCaseContext() {}

func NewMatchCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchCaseContext {
	var p = new(MatchCaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_matchCase

	return p
}

func (s *MatchCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchCaseContext) Case() antlr.TerminalNode {
	return s.GetToken(V2ParserCase, 0)
}

func (s *MatchCaseContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MatchCaseContext) Col() antlr.TerminalNode {
	return s.GetToken(V2ParserCol, 0)
}

func (s *MatchCaseContext) CodeBlock() ICodeBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICodeBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICodeBlockContext)
}

func (s *MatchCaseContext) Default() antlr.TerminalNode {
	return s.GetToken(V2ParserDefault, 0)
}

func (s *MatchCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatchCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterMatchCase(s)
	}
}

func (s *MatchCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitMatchCase(s)
	}
}

func (s *MatchCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitMatchCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) MatchCase() (localctx IMatchCaseContext) {
	localctx = NewMatchCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, V2ParserRULE_matchCase)
	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case V2ParserCase:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(411)
			p.Match(V2ParserCase)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(412)
			p.expr(0)
		}
		{
			p.SetState(413)
			p.Match(V2ParserCol)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(414)
			p.CodeBlock()
		}

	case V2ParserDefault:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(416)
			p.Match(V2ParserDefault)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(417)
			p.Match(V2ParserCol)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(418)
			p.CodeBlock()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IJumpStmtContext is an interface to support dynamic dispatch.
type IJumpStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Continue() antlr.TerminalNode
	Break() antlr.TerminalNode
	Return() antlr.TerminalNode
	Expr() IExprContext

	// IsJumpStmtContext differentiates from other interfaces.
	IsJumpStmtContext()
}

type JumpStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJumpStmtContext() *JumpStmtContext {
	var p = new(JumpStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_jumpStmt
	return p
}

func InitEmptyJumpStmtContext(p *JumpStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_jumpStmt
}

func (*JumpStmtContext) IsJumpStmtContext() {}

func NewJumpStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JumpStmtContext {
	var p = new(JumpStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_jumpStmt

	return p
}

func (s *JumpStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *JumpStmtContext) Continue() antlr.TerminalNode {
	return s.GetToken(V2ParserContinue, 0)
}

func (s *JumpStmtContext) Break() antlr.TerminalNode {
	return s.GetToken(V2ParserBreak, 0)
}

func (s *JumpStmtContext) Return() antlr.TerminalNode {
	return s.GetToken(V2ParserReturn, 0)
}

func (s *JumpStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *JumpStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JumpStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JumpStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterJumpStmt(s)
	}
}

func (s *JumpStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitJumpStmt(s)
	}
}

func (s *JumpStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitJumpStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) JumpStmt() (localctx IJumpStmtContext) {
	localctx = NewJumpStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, V2ParserRULE_jumpStmt)
	p.SetState(427)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case V2ParserContinue:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(421)
			p.Match(V2ParserContinue)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case V2ParserBreak:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(422)
			p.Match(V2ParserBreak)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case V2ParserReturn:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(423)
			p.Match(V2ParserReturn)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(425)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(424)
				p.expr(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISepContext is an interface to support dynamic dispatch.
type ISepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NewLine() antlr.TerminalNode
	Semi() antlr.TerminalNode

	// IsSepContext differentiates from other interfaces.
	IsSepContext()
}

type SepContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySepContext() *SepContext {
	var p = new(SepContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_sep
	return p
}

func InitEmptySepContext(p *SepContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_sep
}

func (*SepContext) IsSepContext() {}

func NewSepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SepContext {
	var p = new(SepContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_sep

	return p
}

func (s *SepContext) GetParser() antlr.Parser { return s.parser }

func (s *SepContext) NewLine() antlr.TerminalNode {
	return s.GetToken(V2ParserNewLine, 0)
}

func (s *SepContext) Semi() antlr.TerminalNode {
	return s.GetToken(V2ParserSemi, 0)
}

func (s *SepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterSep(s)
	}
}

func (s *SepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitSep(s)
	}
}

func (s *SepContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitSep(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) Sep() (localctx ISepContext) {
	localctx = NewSepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, V2ParserRULE_sep)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(429)
		_la = p.GetTokenStream().LA(1)

		if !(_la == V2ParserSemi || _la == V2ParserNewLine) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *V2Parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 17:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *V2Parser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
