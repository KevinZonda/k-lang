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
		"'break'", "", "'true'", "'false'",
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
		"program", "openBlock", "structBlock", "funcBlock", "codeBlock", "declareBlock",
		"funcSig", "funcSignArgs", "funcSignArgItem", "type", "var", "baseVar",
		"index", "indexes", "lambda", "unaryOper", "expr", "exprWithLambda",
		"funcCall", "funcCallArgs", "stmtWithSep", "openStmt", "literal", "initializer",
		"arrayInitializer", "identifierPair", "mapPair", "structInitializer",
		"mapInitializer", "stmt", "assignStmt", "declareStmt", "ifStmt", "loopStmt",
		"cStyleFor", "iterFor", "whileStyleFor", "matchStmt", "matchCase", "jumpStmt",
		"sep",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 50, 482, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 3, 0, 88, 8, 0, 1, 0, 5, 0, 91, 8, 0, 10, 0, 12, 0, 94,
		9, 0, 5, 0, 96, 8, 0, 10, 0, 12, 0, 99, 9, 0, 1, 0, 1, 0, 1, 1, 5, 1, 104,
		8, 1, 10, 1, 12, 1, 107, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 4, 3, 4, 120, 8, 4, 1, 4, 5, 4, 123, 8, 4, 10, 4,
		12, 4, 126, 9, 4, 5, 4, 128, 8, 4, 10, 4, 12, 4, 131, 9, 4, 1, 4, 1, 4,
		1, 5, 1, 5, 5, 5, 137, 8, 5, 10, 5, 12, 5, 140, 9, 5, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 149, 8, 6, 1, 7, 3, 7, 152, 8, 7, 1, 7, 1,
		7, 5, 7, 156, 8, 7, 10, 7, 12, 7, 159, 9, 7, 1, 8, 3, 8, 162, 8, 8, 1,
		8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 170, 8, 9, 1, 9, 1, 9, 1, 9, 3,
		9, 175, 8, 9, 3, 9, 177, 8, 9, 1, 10, 1, 10, 1, 10, 5, 10, 182, 8, 10,
		10, 10, 12, 10, 185, 9, 10, 1, 11, 1, 11, 3, 11, 189, 8, 11, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 13, 4, 13, 196, 8, 13, 11, 13, 12, 13, 197, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 205, 8, 14, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 3, 16, 224, 8, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5,
		16, 251, 8, 16, 10, 16, 12, 16, 254, 9, 16, 1, 17, 1, 17, 3, 17, 258, 8,
		17, 1, 18, 1, 18, 1, 18, 3, 18, 263, 8, 18, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 19, 5, 19, 270, 8, 19, 10, 19, 12, 19, 273, 9, 19, 1, 20, 1, 20, 5,
		20, 277, 8, 20, 10, 20, 12, 20, 280, 9, 20, 1, 21, 1, 21, 1, 21, 1, 21,
		3, 21, 286, 8, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 3, 23, 293, 8, 23,
		1, 24, 3, 24, 296, 8, 24, 1, 24, 1, 24, 1, 24, 3, 24, 301, 8, 24, 5, 24,
		303, 8, 24, 10, 24, 12, 24, 306, 9, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		3, 26, 323, 8, 26, 1, 27, 1, 27, 3, 27, 327, 8, 27, 1, 27, 1, 27, 1, 27,
		3, 27, 332, 8, 27, 5, 27, 334, 8, 27, 10, 27, 12, 27, 337, 9, 27, 1, 27,
		1, 27, 1, 28, 3, 28, 342, 8, 28, 1, 28, 1, 28, 1, 28, 5, 28, 347, 8, 28,
		10, 28, 12, 28, 350, 9, 28, 5, 28, 352, 8, 28, 10, 28, 12, 28, 355, 9,
		28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29,
		3, 29, 367, 8, 29, 1, 30, 3, 30, 370, 8, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 31, 1, 31, 1, 31, 1, 31, 5, 31, 380, 8, 31, 10, 31, 12, 31, 383, 9,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 391, 8, 31, 1, 31,
		1, 31, 1, 31, 3, 31, 396, 8, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 3, 32, 405, 8, 32, 1, 33, 1, 33, 1, 33, 3, 33, 410, 8, 33, 1,
		34, 1, 34, 1, 34, 3, 34, 415, 8, 34, 1, 34, 1, 34, 3, 34, 419, 8, 34, 1,
		34, 1, 34, 3, 34, 423, 8, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35,
		3, 35, 431, 8, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1,
		36, 1, 36, 3, 36, 442, 8, 36, 1, 36, 3, 36, 445, 8, 36, 1, 36, 1, 36, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 5, 37, 455, 8, 37, 10, 37, 12, 37,
		458, 9, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1,
		38, 1, 38, 3, 38, 470, 8, 38, 1, 39, 1, 39, 1, 39, 1, 39, 3, 39, 476, 8,
		39, 3, 39, 478, 8, 39, 1, 40, 1, 40, 1, 40, 0, 1, 32, 41, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
		46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80,
		0, 8, 2, 0, 21, 22, 45, 45, 1, 0, 10, 15, 1, 0, 16, 17, 2, 0, 19, 20, 23,
		23, 1, 0, 21, 22, 1, 0, 40, 44, 2, 0, 8, 8, 25, 25, 2, 0, 9, 9, 50, 50,
		520, 0, 82, 1, 0, 0, 0, 2, 105, 1, 0, 0, 0, 4, 108, 1, 0, 0, 0, 6, 112,
		1, 0, 0, 0, 8, 116, 1, 0, 0, 0, 10, 134, 1, 0, 0, 0, 12, 143, 1, 0, 0,
		0, 14, 151, 1, 0, 0, 0, 16, 161, 1, 0, 0, 0, 18, 176, 1, 0, 0, 0, 20, 178,
		1, 0, 0, 0, 22, 186, 1, 0, 0, 0, 24, 190, 1, 0, 0, 0, 26, 195, 1, 0, 0,
		0, 28, 199, 1, 0, 0, 0, 30, 208, 1, 0, 0, 0, 32, 223, 1, 0, 0, 0, 34, 257,
		1, 0, 0, 0, 36, 259, 1, 0, 0, 0, 38, 266, 1, 0, 0, 0, 40, 274, 1, 0, 0,
		0, 42, 281, 1, 0, 0, 0, 44, 287, 1, 0, 0, 0, 46, 292, 1, 0, 0, 0, 48, 295,
		1, 0, 0, 0, 50, 309, 1, 0, 0, 0, 52, 322, 1, 0, 0, 0, 54, 326, 1, 0, 0,
		0, 56, 341, 1, 0, 0, 0, 58, 366, 1, 0, 0, 0, 60, 369, 1, 0, 0, 0, 62, 395,
		1, 0, 0, 0, 64, 397, 1, 0, 0, 0, 66, 409, 1, 0, 0, 0, 68, 411, 1, 0, 0,
		0, 70, 427, 1, 0, 0, 0, 72, 438, 1, 0, 0, 0, 74, 448, 1, 0, 0, 0, 76, 469,
		1, 0, 0, 0, 78, 477, 1, 0, 0, 0, 80, 479, 1, 0, 0, 0, 82, 97, 3, 2, 1,
		0, 83, 88, 3, 4, 2, 0, 84, 88, 3, 6, 3, 0, 85, 88, 3, 58, 29, 0, 86, 88,
		3, 32, 16, 0, 87, 83, 1, 0, 0, 0, 87, 84, 1, 0, 0, 0, 87, 85, 1, 0, 0,
		0, 87, 86, 1, 0, 0, 0, 88, 92, 1, 0, 0, 0, 89, 91, 3, 80, 40, 0, 90, 89,
		1, 0, 0, 0, 91, 94, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0,
		93, 96, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 95, 87, 1, 0, 0, 0, 96, 99, 1,
		0, 0, 0, 97, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 100, 1, 0, 0, 0, 99,
		97, 1, 0, 0, 0, 100, 101, 5, 0, 0, 1, 101, 1, 1, 0, 0, 0, 102, 104, 3,
		42, 21, 0, 103, 102, 1, 0, 0, 0, 104, 107, 1, 0, 0, 0, 105, 103, 1, 0,
		0, 0, 105, 106, 1, 0, 0, 0, 106, 3, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 108,
		109, 5, 26, 0, 0, 109, 110, 5, 47, 0, 0, 110, 111, 3, 10, 5, 0, 111, 5,
		1, 0, 0, 0, 112, 113, 5, 28, 0, 0, 113, 114, 3, 12, 6, 0, 114, 115, 3,
		8, 4, 0, 115, 7, 1, 0, 0, 0, 116, 129, 5, 1, 0, 0, 117, 120, 3, 58, 29,
		0, 118, 120, 3, 32, 16, 0, 119, 117, 1, 0, 0, 0, 119, 118, 1, 0, 0, 0,
		120, 124, 1, 0, 0, 0, 121, 123, 3, 80, 40, 0, 122, 121, 1, 0, 0, 0, 123,
		126, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 128,
		1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 127, 119, 1, 0, 0, 0, 128, 131, 1, 0,
		0, 0, 129, 127, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 132, 1, 0, 0, 0,
		131, 129, 1, 0, 0, 0, 132, 133, 5, 2, 0, 0, 133, 9, 1, 0, 0, 0, 134, 138,
		5, 1, 0, 0, 135, 137, 3, 62, 31, 0, 136, 135, 1, 0, 0, 0, 137, 140, 1,
		0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 141, 1, 0, 0,
		0, 140, 138, 1, 0, 0, 0, 141, 142, 5, 2, 0, 0, 142, 11, 1, 0, 0, 0, 143,
		144, 5, 47, 0, 0, 144, 145, 5, 3, 0, 0, 145, 146, 3, 14, 7, 0, 146, 148,
		5, 4, 0, 0, 147, 149, 3, 18, 9, 0, 148, 147, 1, 0, 0, 0, 148, 149, 1, 0,
		0, 0, 149, 13, 1, 0, 0, 0, 150, 152, 3, 16, 8, 0, 151, 150, 1, 0, 0, 0,
		151, 152, 1, 0, 0, 0, 152, 157, 1, 0, 0, 0, 153, 154, 5, 7, 0, 0, 154,
		156, 3, 16, 8, 0, 155, 153, 1, 0, 0, 0, 156, 159, 1, 0, 0, 0, 157, 155,
		1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 15, 1, 0, 0, 0, 159, 157, 1, 0,
		0, 0, 160, 162, 3, 18, 9, 0, 161, 160, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0,
		162, 163, 1, 0, 0, 0, 163, 164, 5, 47, 0, 0, 164, 17, 1, 0, 0, 0, 165,
		177, 5, 27, 0, 0, 166, 177, 5, 28, 0, 0, 167, 168, 5, 47, 0, 0, 168, 170,
		5, 24, 0, 0, 169, 167, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 171, 1, 0,
		0, 0, 171, 174, 5, 47, 0, 0, 172, 173, 5, 5, 0, 0, 173, 175, 5, 6, 0, 0,
		174, 172, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 177, 1, 0, 0, 0, 176,
		165, 1, 0, 0, 0, 176, 166, 1, 0, 0, 0, 176, 169, 1, 0, 0, 0, 177, 19, 1,
		0, 0, 0, 178, 183, 3, 22, 11, 0, 179, 180, 5, 24, 0, 0, 180, 182, 3, 22,
		11, 0, 181, 179, 1, 0, 0, 0, 182, 185, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0,
		183, 184, 1, 0, 0, 0, 184, 21, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 186, 188,
		5, 47, 0, 0, 187, 189, 3, 26, 13, 0, 188, 187, 1, 0, 0, 0, 188, 189, 1,
		0, 0, 0, 189, 23, 1, 0, 0, 0, 190, 191, 5, 5, 0, 0, 191, 192, 3, 32, 16,
		0, 192, 193, 5, 6, 0, 0, 193, 25, 1, 0, 0, 0, 194, 196, 3, 24, 12, 0, 195,
		194, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 195, 1, 0, 0, 0, 197, 198,
		1, 0, 0, 0, 198, 27, 1, 0, 0, 0, 199, 200, 5, 28, 0, 0, 200, 201, 5, 3,
		0, 0, 201, 202, 3, 14, 7, 0, 202, 204, 5, 4, 0, 0, 203, 205, 3, 18, 9,
		0, 204, 203, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206,
		207, 3, 8, 4, 0, 207, 29, 1, 0, 0, 0, 208, 209, 7, 0, 0, 0, 209, 31, 1,
		0, 0, 0, 210, 211, 6, 16, -1, 0, 211, 224, 3, 36, 18, 0, 212, 213, 3, 30,
		15, 0, 213, 214, 3, 32, 16, 12, 214, 224, 1, 0, 0, 0, 215, 216, 5, 3, 0,
		0, 216, 217, 3, 32, 16, 0, 217, 218, 5, 4, 0, 0, 218, 224, 1, 0, 0, 0,
		219, 224, 5, 47, 0, 0, 220, 224, 3, 44, 22, 0, 221, 224, 3, 46, 23, 0,
		222, 224, 3, 60, 30, 0, 223, 210, 1, 0, 0, 0, 223, 212, 1, 0, 0, 0, 223,
		215, 1, 0, 0, 0, 223, 219, 1, 0, 0, 0, 223, 220, 1, 0, 0, 0, 223, 221,
		1, 0, 0, 0, 223, 222, 1, 0, 0, 0, 224, 252, 1, 0, 0, 0, 225, 226, 10, 15,
		0, 0, 226, 227, 5, 24, 0, 0, 227, 251, 3, 32, 16, 16, 228, 229, 10, 10,
		0, 0, 229, 230, 7, 1, 0, 0, 230, 251, 3, 32, 16, 11, 231, 232, 10, 9, 0,
		0, 232, 233, 7, 2, 0, 0, 233, 251, 3, 32, 16, 10, 234, 235, 10, 8, 0, 0,
		235, 236, 5, 18, 0, 0, 236, 251, 3, 32, 16, 9, 237, 238, 10, 7, 0, 0, 238,
		239, 7, 3, 0, 0, 239, 251, 3, 32, 16, 8, 240, 241, 10, 6, 0, 0, 241, 242,
		7, 4, 0, 0, 242, 251, 3, 32, 16, 7, 243, 244, 10, 14, 0, 0, 244, 245, 5,
		5, 0, 0, 245, 246, 3, 32, 16, 0, 246, 247, 5, 6, 0, 0, 247, 251, 1, 0,
		0, 0, 248, 249, 10, 2, 0, 0, 249, 251, 3, 26, 13, 0, 250, 225, 1, 0, 0,
		0, 250, 228, 1, 0, 0, 0, 250, 231, 1, 0, 0, 0, 250, 234, 1, 0, 0, 0, 250,
		237, 1, 0, 0, 0, 250, 240, 1, 0, 0, 0, 250, 243, 1, 0, 0, 0, 250, 248,
		1, 0, 0, 0, 251, 254, 1, 0, 0, 0, 252, 250, 1, 0, 0, 0, 252, 253, 1, 0,
		0, 0, 253, 33, 1, 0, 0, 0, 254, 252, 1, 0, 0, 0, 255, 258, 3, 28, 14, 0,
		256, 258, 3, 32, 16, 0, 257, 255, 1, 0, 0, 0, 257, 256, 1, 0, 0, 0, 258,
		35, 1, 0, 0, 0, 259, 260, 5, 47, 0, 0, 260, 262, 5, 3, 0, 0, 261, 263,
		3, 38, 19, 0, 262, 261, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 264, 1,
		0, 0, 0, 264, 265, 5, 4, 0, 0, 265, 37, 1, 0, 0, 0, 266, 271, 3, 32, 16,
		0, 267, 268, 5, 7, 0, 0, 268, 270, 3, 32, 16, 0, 269, 267, 1, 0, 0, 0,
		270, 273, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272,
		39, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 274, 278, 3, 58, 29, 0, 275, 277,
		3, 80, 40, 0, 276, 275, 1, 0, 0, 0, 277, 280, 1, 0, 0, 0, 278, 276, 1,
		0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 41, 1, 0, 0, 0, 280, 278, 1, 0, 0,
		0, 281, 282, 5, 32, 0, 0, 282, 285, 5, 44, 0, 0, 283, 284, 5, 33, 0, 0,
		284, 286, 5, 47, 0, 0, 285, 283, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286,
		43, 1, 0, 0, 0, 287, 288, 7, 5, 0, 0, 288, 45, 1, 0, 0, 0, 289, 293, 3,
		48, 24, 0, 290, 293, 3, 54, 27, 0, 291, 293, 3, 56, 28, 0, 292, 289, 1,
		0, 0, 0, 292, 290, 1, 0, 0, 0, 292, 291, 1, 0, 0, 0, 293, 47, 1, 0, 0,
		0, 294, 296, 3, 18, 9, 0, 295, 294, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296,
		297, 1, 0, 0, 0, 297, 304, 5, 5, 0, 0, 298, 300, 3, 32, 16, 0, 299, 301,
		5, 7, 0, 0, 300, 299, 1, 0, 0, 0, 300, 301, 1, 0, 0, 0, 301, 303, 1, 0,
		0, 0, 302, 298, 1, 0, 0, 0, 303, 306, 1, 0, 0, 0, 304, 302, 1, 0, 0, 0,
		304, 305, 1, 0, 0, 0, 305, 307, 1, 0, 0, 0, 306, 304, 1, 0, 0, 0, 307,
		308, 5, 6, 0, 0, 308, 49, 1, 0, 0, 0, 309, 310, 5, 47, 0, 0, 310, 311,
		7, 6, 0, 0, 311, 312, 3, 34, 17, 0, 312, 51, 1, 0, 0, 0, 313, 314, 3, 32,
		16, 0, 314, 315, 7, 6, 0, 0, 315, 316, 3, 34, 17, 0, 316, 323, 1, 0, 0,
		0, 317, 318, 5, 3, 0, 0, 318, 319, 3, 52, 26, 0, 319, 320, 5, 4, 0, 0,
		320, 323, 1, 0, 0, 0, 321, 323, 3, 50, 25, 0, 322, 313, 1, 0, 0, 0, 322,
		317, 1, 0, 0, 0, 322, 321, 1, 0, 0, 0, 323, 53, 1, 0, 0, 0, 324, 327, 3,
		18, 9, 0, 325, 327, 5, 26, 0, 0, 326, 324, 1, 0, 0, 0, 326, 325, 1, 0,
		0, 0, 327, 328, 1, 0, 0, 0, 328, 335, 5, 1, 0, 0, 329, 331, 3, 50, 25,
		0, 330, 332, 5, 7, 0, 0, 331, 330, 1, 0, 0, 0, 331, 332, 1, 0, 0, 0, 332,
		334, 1, 0, 0, 0, 333, 329, 1, 0, 0, 0, 334, 337, 1, 0, 0, 0, 335, 333,
		1, 0, 0, 0, 335, 336, 1, 0, 0, 0, 336, 338, 1, 0, 0, 0, 337, 335, 1, 0,
		0, 0, 338, 339, 5, 2, 0, 0, 339, 55, 1, 0, 0, 0, 340, 342, 5, 27, 0, 0,
		341, 340, 1, 0, 0, 0, 341, 342, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343,
		353, 5, 1, 0, 0, 344, 348, 3, 52, 26, 0, 345, 347, 5, 7, 0, 0, 346, 345,
		1, 0, 0, 0, 347, 350, 1, 0, 0, 0, 348, 346, 1, 0, 0, 0, 348, 349, 1, 0,
		0, 0, 349, 352, 1, 0, 0, 0, 350, 348, 1, 0, 0, 0, 351, 344, 1, 0, 0, 0,
		352, 355, 1, 0, 0, 0, 353, 351, 1, 0, 0, 0, 353, 354, 1, 0, 0, 0, 354,
		356, 1, 0, 0, 0, 355, 353, 1, 0, 0, 0, 356, 357, 5, 2, 0, 0, 357, 57, 1,
		0, 0, 0, 358, 367, 3, 60, 30, 0, 359, 367, 3, 62, 31, 0, 360, 367, 3, 78,
		39, 0, 361, 367, 3, 64, 32, 0, 362, 367, 3, 66, 33, 0, 363, 367, 3, 74,
		37, 0, 364, 367, 3, 8, 4, 0, 365, 367, 3, 36, 18, 0, 366, 358, 1, 0, 0,
		0, 366, 359, 1, 0, 0, 0, 366, 360, 1, 0, 0, 0, 366, 361, 1, 0, 0, 0, 366,
		362, 1, 0, 0, 0, 366, 363, 1, 0, 0, 0, 366, 364, 1, 0, 0, 0, 366, 365,
		1, 0, 0, 0, 367, 59, 1, 0, 0, 0, 368, 370, 3, 18, 9, 0, 369, 368, 1, 0,
		0, 0, 369, 370, 1, 0, 0, 0, 370, 371, 1, 0, 0, 0, 371, 372, 3, 20, 10,
		0, 372, 373, 5, 46, 0, 0, 373, 374, 3, 34, 17, 0, 374, 61, 1, 0, 0, 0,
		375, 376, 3, 18, 9, 0, 376, 381, 5, 47, 0, 0, 377, 378, 5, 7, 0, 0, 378,
		380, 5, 47, 0, 0, 379, 377, 1, 0, 0, 0, 380, 383, 1, 0, 0, 0, 381, 379,
		1, 0, 0, 0, 381, 382, 1, 0, 0, 0, 382, 396, 1, 0, 0, 0, 383, 381, 1, 0,
		0, 0, 384, 385, 3, 18, 9, 0, 385, 386, 5, 47, 0, 0, 386, 387, 5, 46, 0,
		0, 387, 388, 3, 34, 17, 0, 388, 396, 1, 0, 0, 0, 389, 391, 3, 18, 9, 0,
		390, 389, 1, 0, 0, 0, 390, 391, 1, 0, 0, 0, 391, 392, 1, 0, 0, 0, 392,
		393, 5, 47, 0, 0, 393, 394, 5, 46, 0, 0, 394, 396, 3, 34, 17, 0, 395, 375,
		1, 0, 0, 0, 395, 384, 1, 0, 0, 0, 395, 390, 1, 0, 0, 0, 396, 63, 1, 0,
		0, 0, 397, 398, 5, 34, 0, 0, 398, 399, 5, 3, 0, 0, 399, 400, 3, 32, 16,
		0, 400, 401, 5, 4, 0, 0, 401, 404, 3, 8, 4, 0, 402, 403, 5, 35, 0, 0, 403,
		405, 3, 8, 4, 0, 404, 402, 1, 0, 0, 0, 404, 405, 1, 0, 0, 0, 405, 65, 1,
		0, 0, 0, 406, 410, 3, 70, 35, 0, 407, 410, 3, 68, 34, 0, 408, 410, 3, 72,
		36, 0, 409, 406, 1, 0, 0, 0, 409, 407, 1, 0, 0, 0, 409, 408, 1, 0, 0, 0,
		410, 67, 1, 0, 0, 0, 411, 412, 5, 36, 0, 0, 412, 414, 5, 3, 0, 0, 413,
		415, 3, 32, 16, 0, 414, 413, 1, 0, 0, 0, 414, 415, 1, 0, 0, 0, 415, 416,
		1, 0, 0, 0, 416, 418, 5, 9, 0, 0, 417, 419, 3, 32, 16, 0, 418, 417, 1,
		0, 0, 0, 418, 419, 1, 0, 0, 0, 419, 420, 1, 0, 0, 0, 420, 422, 5, 9, 0,
		0, 421, 423, 3, 32, 16, 0, 422, 421, 1, 0, 0, 0, 422, 423, 1, 0, 0, 0,
		423, 424, 1, 0, 0, 0, 424, 425, 5, 4, 0, 0, 425, 426, 3, 8, 4, 0, 426,
		69, 1, 0, 0, 0, 427, 428, 5, 36, 0, 0, 428, 430, 5, 3, 0, 0, 429, 431,
		3, 18, 9, 0, 430, 429, 1, 0, 0, 0, 430, 431, 1, 0, 0, 0, 431, 432, 1, 0,
		0, 0, 432, 433, 5, 47, 0, 0, 433, 434, 5, 8, 0, 0, 434, 435, 3, 32, 16,
		0, 435, 436, 5, 4, 0, 0, 436, 437, 3, 8, 4, 0, 437, 71, 1, 0, 0, 0, 438,
		444, 5, 36, 0, 0, 439, 441, 5, 3, 0, 0, 440, 442, 3, 32, 16, 0, 441, 440,
		1, 0, 0, 0, 441, 442, 1, 0, 0, 0, 442, 443, 1, 0, 0, 0, 443, 445, 5, 4,
		0, 0, 444, 439, 1, 0, 0, 0, 444, 445, 1, 0, 0, 0, 445, 446, 1, 0, 0, 0,
		446, 447, 3, 8, 4, 0, 447, 73, 1, 0, 0, 0, 448, 449, 5, 37, 0, 0, 449,
		450, 5, 3, 0, 0, 450, 451, 3, 32, 16, 0, 451, 452, 5, 4, 0, 0, 452, 456,
		5, 1, 0, 0, 453, 455, 3, 76, 38, 0, 454, 453, 1, 0, 0, 0, 455, 458, 1,
		0, 0, 0, 456, 454, 1, 0, 0, 0, 456, 457, 1, 0, 0, 0, 457, 459, 1, 0, 0,
		0, 458, 456, 1, 0, 0, 0, 459, 460, 5, 2, 0, 0, 460, 75, 1, 0, 0, 0, 461,
		462, 5, 30, 0, 0, 462, 463, 3, 32, 16, 0, 463, 464, 5, 8, 0, 0, 464, 465,
		3, 8, 4, 0, 465, 470, 1, 0, 0, 0, 466, 467, 5, 31, 0, 0, 467, 468, 5, 8,
		0, 0, 468, 470, 3, 8, 4, 0, 469, 461, 1, 0, 0, 0, 469, 466, 1, 0, 0, 0,
		470, 77, 1, 0, 0, 0, 471, 478, 5, 39, 0, 0, 472, 478, 5, 38, 0, 0, 473,
		475, 5, 29, 0, 0, 474, 476, 3, 34, 17, 0, 475, 474, 1, 0, 0, 0, 475, 476,
		1, 0, 0, 0, 476, 478, 1, 0, 0, 0, 477, 471, 1, 0, 0, 0, 477, 472, 1, 0,
		0, 0, 477, 473, 1, 0, 0, 0, 478, 79, 1, 0, 0, 0, 479, 480, 7, 7, 0, 0,
		480, 81, 1, 0, 0, 0, 55, 87, 92, 97, 105, 119, 124, 129, 138, 148, 151,
		157, 161, 169, 174, 176, 183, 188, 197, 204, 223, 250, 252, 257, 262, 271,
		278, 285, 292, 295, 300, 304, 322, 326, 331, 335, 341, 348, 353, 366, 369,
		381, 390, 395, 404, 409, 414, 418, 422, 430, 441, 444, 456, 469, 475, 477,
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
	V2ParserComment        = 48
	V2ParserWS             = 49
	V2ParserNewLine        = 50
)

// V2Parser rules.
const (
	V2ParserRULE_program           = 0
	V2ParserRULE_openBlock         = 1
	V2ParserRULE_structBlock       = 2
	V2ParserRULE_funcBlock         = 3
	V2ParserRULE_codeBlock         = 4
	V2ParserRULE_declareBlock      = 5
	V2ParserRULE_funcSig           = 6
	V2ParserRULE_funcSignArgs      = 7
	V2ParserRULE_funcSignArgItem   = 8
	V2ParserRULE_type              = 9
	V2ParserRULE_var               = 10
	V2ParserRULE_baseVar           = 11
	V2ParserRULE_index             = 12
	V2ParserRULE_indexes           = 13
	V2ParserRULE_lambda            = 14
	V2ParserRULE_unaryOper         = 15
	V2ParserRULE_expr              = 16
	V2ParserRULE_exprWithLambda    = 17
	V2ParserRULE_funcCall          = 18
	V2ParserRULE_funcCallArgs      = 19
	V2ParserRULE_stmtWithSep       = 20
	V2ParserRULE_openStmt          = 21
	V2ParserRULE_literal           = 22
	V2ParserRULE_initializer       = 23
	V2ParserRULE_arrayInitializer  = 24
	V2ParserRULE_identifierPair    = 25
	V2ParserRULE_mapPair           = 26
	V2ParserRULE_structInitializer = 27
	V2ParserRULE_mapInitializer    = 28
	V2ParserRULE_stmt              = 29
	V2ParserRULE_assignStmt        = 30
	V2ParserRULE_declareStmt       = 31
	V2ParserRULE_ifStmt            = 32
	V2ParserRULE_loopStmt          = 33
	V2ParserRULE_cStyleFor         = 34
	V2ParserRULE_iterFor           = 35
	V2ParserRULE_whileStyleFor     = 36
	V2ParserRULE_matchStmt         = 37
	V2ParserRULE_matchCase         = 38
	V2ParserRULE_jumpStmt          = 39
	V2ParserRULE_sep               = 40
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
		p.SetState(82)
		p.OpenBlock()
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&211055705849898) != 0 {
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(83)
				p.StructBlock()
			}

		case 2:
			{
				p.SetState(84)
				p.FuncBlock()
			}

		case 3:
			{
				p.SetState(85)
				p.Stmt()
			}

		case 4:
			{
				p.SetState(86)
				p.expr(0)
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == V2ParserSemi || _la == V2ParserNewLine {
			{
				p.SetState(89)
				p.Sep()
			}

			p.SetState(94)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(100)
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
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == V2ParserOpen {
		{
			p.SetState(102)
			p.OpenStmt()
		}

		p.SetState(107)
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
		p.SetState(108)
		p.Match(V2ParserStruct)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Match(V2ParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)
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
		p.SetState(112)
		p.Match(V2ParserFunction)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.FuncSig()
	}
	{
		p.SetState(114)
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
	AllExpr() []IExprContext
	Expr(i int) IExprContext
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

func (s *CodeBlockContext) AllExpr() []IExprContext {
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

func (s *CodeBlockContext) Expr(i int) IExprContext {
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
		p.SetState(116)
		p.Match(V2ParserLBrack)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&211055705849898) != 0 {
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(117)
				p.Stmt()
			}

		case 2:
			{
				p.SetState(118)
				p.expr(0)
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == V2ParserSemi || _la == V2ParserNewLine {
			{
				p.SetState(121)
				p.Sep()
			}

			p.SetState(126)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(132)
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
		p.SetState(134)
		p.Match(V2ParserLBrack)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140737891008512) != 0 {
		{
			p.SetState(135)
			p.DeclareStmt()
		}

		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(141)
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
	FuncSignArgs() IFuncSignArgsContext
	RParen() antlr.TerminalNode
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

func (s *FuncSigContext) RParen() antlr.TerminalNode {
	return s.GetToken(V2ParserRParen, 0)
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
		p.SetState(143)
		p.Match(V2ParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Match(V2ParserLParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.FuncSignArgs()
	}
	{
		p.SetState(146)
		p.Match(V2ParserRParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140737891008512) != 0 {
		{
			p.SetState(147)
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
	AllFuncSignArgItem() []IFuncSignArgItemContext
	FuncSignArgItem(i int) IFuncSignArgItemContext
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

func (s *FuncSignArgsContext) AllFuncSignArgItem() []IFuncSignArgItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncSignArgItemContext); ok {
			len++
		}
	}

	tst := make([]IFuncSignArgItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncSignArgItemContext); ok {
			tst[i] = t.(IFuncSignArgItemContext)
			i++
		}
	}

	return tst
}

func (s *FuncSignArgsContext) FuncSignArgItem(i int) IFuncSignArgItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncSignArgItemContext); ok {
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

	return t.(IFuncSignArgItemContext)
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
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140737891008512) != 0 {
		{
			p.SetState(150)
			p.FuncSignArgItem()
		}

	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == V2ParserComma {
		{
			p.SetState(153)
			p.Match(V2ParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.FuncSignArgItem()
		}

		p.SetState(159)
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

// IFuncSignArgItemContext is an interface to support dynamic dispatch.
type IFuncSignArgItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Type_() ITypeContext

	// IsFuncSignArgItemContext differentiates from other interfaces.
	IsFuncSignArgItemContext()
}

type FuncSignArgItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncSignArgItemContext() *FuncSignArgItemContext {
	var p = new(FuncSignArgItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_funcSignArgItem
	return p
}

func InitEmptyFuncSignArgItemContext(p *FuncSignArgItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_funcSignArgItem
}

func (*FuncSignArgItemContext) IsFuncSignArgItemContext() {}

func NewFuncSignArgItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncSignArgItemContext {
	var p = new(FuncSignArgItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_funcSignArgItem

	return p
}

func (s *FuncSignArgItemContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncSignArgItemContext) Identifier() antlr.TerminalNode {
	return s.GetToken(V2ParserIdentifier, 0)
}

func (s *FuncSignArgItemContext) Type_() ITypeContext {
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

func (s *FuncSignArgItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncSignArgItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncSignArgItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterFuncSignArgItem(s)
	}
}

func (s *FuncSignArgItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitFuncSignArgItem(s)
	}
}

func (s *FuncSignArgItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitFuncSignArgItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) FuncSignArgItem() (localctx IFuncSignArgItemContext) {
	localctx = NewFuncSignArgItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, V2ParserRULE_funcSignArgItem)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(161)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(160)
			p.Type_()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(163)
		p.Match(V2ParserIdentifier)
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

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPackageName returns the PackageName token.
	GetPackageName() antlr.Token

	// GetTypeName returns the TypeName token.
	GetTypeName() antlr.Token

	// SetPackageName sets the PackageName token.
	SetPackageName(antlr.Token)

	// SetTypeName sets the TypeName token.
	SetTypeName(antlr.Token)

	// Getter signatures
	Map() antlr.TerminalNode
	Function() antlr.TerminalNode
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode
	Dot() antlr.TerminalNode
	LSquare() antlr.TerminalNode
	RSquare() antlr.TerminalNode

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	PackageName antlr.Token
	TypeName    antlr.Token
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

func (s *TypeContext) GetPackageName() antlr.Token { return s.PackageName }

func (s *TypeContext) GetTypeName() antlr.Token { return s.TypeName }

func (s *TypeContext) SetPackageName(v antlr.Token) { s.PackageName = v }

func (s *TypeContext) SetTypeName(v antlr.Token) { s.TypeName = v }

func (s *TypeContext) Map() antlr.TerminalNode {
	return s.GetToken(V2ParserMap, 0)
}

func (s *TypeContext) Function() antlr.TerminalNode {
	return s.GetToken(V2ParserFunction, 0)
}

func (s *TypeContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(V2ParserIdentifier)
}

func (s *TypeContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(V2ParserIdentifier, i)
}

func (s *TypeContext) Dot() antlr.TerminalNode {
	return s.GetToken(V2ParserDot, 0)
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
	p.EnterRule(localctx, 18, V2ParserRULE_type)
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case V2ParserMap:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(165)
			p.Match(V2ParserMap)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case V2ParserFunction:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(166)
			p.Match(V2ParserFunction)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case V2ParserIdentifier:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(169)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(167)

				var _m = p.Match(V2ParserIdentifier)

				localctx.(*TypeContext).PackageName = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(168)
				p.Match(V2ParserDot)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(171)

			var _m = p.Match(V2ParserIdentifier)

			localctx.(*TypeContext).TypeName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(172)
				p.Match(V2ParserLSquare)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(173)
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
	AllBaseVar() []IBaseVarContext
	BaseVar(i int) IBaseVarContext
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

func (s *VarContext) AllBaseVar() []IBaseVarContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBaseVarContext); ok {
			len++
		}
	}

	tst := make([]IBaseVarContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBaseVarContext); ok {
			tst[i] = t.(IBaseVarContext)
			i++
		}
	}

	return tst
}

func (s *VarContext) BaseVar(i int) IBaseVarContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBaseVarContext); ok {
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

	return t.(IBaseVarContext)
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
	p.EnterRule(localctx, 20, V2ParserRULE_var)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.BaseVar()
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == V2ParserDot {
		{
			p.SetState(179)
			p.Match(V2ParserDot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)
			p.BaseVar()
		}

		p.SetState(185)
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

// IBaseVarContext is an interface to support dynamic dispatch.
type IBaseVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Indexes() IIndexesContext

	// IsBaseVarContext differentiates from other interfaces.
	IsBaseVarContext()
}

type BaseVarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBaseVarContext() *BaseVarContext {
	var p = new(BaseVarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_baseVar
	return p
}

func InitEmptyBaseVarContext(p *BaseVarContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_baseVar
}

func (*BaseVarContext) IsBaseVarContext() {}

func NewBaseVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BaseVarContext {
	var p = new(BaseVarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_baseVar

	return p
}

func (s *BaseVarContext) GetParser() antlr.Parser { return s.parser }

func (s *BaseVarContext) Identifier() antlr.TerminalNode {
	return s.GetToken(V2ParserIdentifier, 0)
}

func (s *BaseVarContext) Indexes() IIndexesContext {
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

func (s *BaseVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseVarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BaseVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterBaseVar(s)
	}
}

func (s *BaseVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitBaseVar(s)
	}
}

func (s *BaseVarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitBaseVar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) BaseVar() (localctx IBaseVarContext) {
	localctx = NewBaseVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, V2ParserRULE_baseVar)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		p.Match(V2ParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == V2ParserLSquare {
		{
			p.SetState(187)
			p.Indexes()
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
		p.SetState(190)
		p.Match(V2ParserLSquare)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(191)
		p.expr(0)
	}
	{
		p.SetState(192)
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
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(194)
				p.Index()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
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

// ILambdaContext is an interface to support dynamic dispatch.
type ILambdaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Function() antlr.TerminalNode
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

func (s *LambdaContext) Function() antlr.TerminalNode {
	return s.GetToken(V2ParserFunction, 0)
}

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
		p.SetState(199)
		p.Match(V2ParserFunction)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(200)
		p.Match(V2ParserLParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(201)
		p.FuncSignArgs()
	}
	{
		p.SetState(202)
		p.Match(V2ParserRParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140737891008512) != 0 {
		{
			p.SetState(203)
			p.Type_()
		}

	}
	{
		p.SetState(206)
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
	p.EnterRule(localctx, 30, V2ParserRULE_unaryOper)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
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

	// GetOP returns the OP token.
	GetOP() antlr.Token

	// SetOP sets the OP token.
	SetOP(antlr.Token)

	// GetLHS returns the LHS rule contexts.
	GetLHS() IExprContext

	// GetRHS returns the RHS rule contexts.
	GetRHS() IExprContext

	// GetIndex returns the Index rule contexts.
	GetIndex() IExprContext

	// SetLHS sets the LHS rule contexts.
	SetLHS(IExprContext)

	// SetRHS sets the RHS rule contexts.
	SetRHS(IExprContext)

	// SetIndex sets the Index rule contexts.
	SetIndex(IExprContext)

	// Getter signatures
	FuncCall() IFuncCallContext
	UnaryOper() IUnaryOperContext
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	LParen() antlr.TerminalNode
	RParen() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	Literal() ILiteralContext
	Initializer() IInitializerContext
	AssignStmt() IAssignStmtContext
	Dot() antlr.TerminalNode
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
	LSquare() antlr.TerminalNode
	RSquare() antlr.TerminalNode
	Indexes() IIndexesContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	LHS    IExprContext
	RHS    IExprContext
	OP     antlr.Token
	Index  IExprContext
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

func (s *ExprContext) GetOP() antlr.Token { return s.OP }

func (s *ExprContext) SetOP(v antlr.Token) { s.OP = v }

func (s *ExprContext) GetLHS() IExprContext { return s.LHS }

func (s *ExprContext) GetRHS() IExprContext { return s.RHS }

func (s *ExprContext) GetIndex() IExprContext { return s.Index }

func (s *ExprContext) SetLHS(v IExprContext) { s.LHS = v }

func (s *ExprContext) SetRHS(v IExprContext) { s.RHS = v }

func (s *ExprContext) SetIndex(v IExprContext) { s.Index = v }

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

func (s *ExprContext) LParen() antlr.TerminalNode {
	return s.GetToken(V2ParserLParen, 0)
}

func (s *ExprContext) RParen() antlr.TerminalNode {
	return s.GetToken(V2ParserRParen, 0)
}

func (s *ExprContext) Identifier() antlr.TerminalNode {
	return s.GetToken(V2ParserIdentifier, 0)
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

func (s *ExprContext) Initializer() IInitializerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInitializerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInitializerContext)
}

func (s *ExprContext) AssignStmt() IAssignStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignStmtContext)
}

func (s *ExprContext) Dot() antlr.TerminalNode {
	return s.GetToken(V2ParserDot, 0)
}

func (s *ExprContext) Equals() antlr.TerminalNode {
	return s.GetToken(V2ParserEquals, 0)
}

func (s *ExprContext) NotEq() antlr.TerminalNode {
	return s.GetToken(V2ParserNotEq, 0)
}

func (s *ExprContext) Greater() antlr.TerminalNode {
	return s.GetToken(V2ParserGreater, 0)
}

func (s *ExprContext) Less() antlr.TerminalNode {
	return s.GetToken(V2ParserLess, 0)
}

func (s *ExprContext) GreaterEq() antlr.TerminalNode {
	return s.GetToken(V2ParserGreaterEq, 0)
}

func (s *ExprContext) LessEq() antlr.TerminalNode {
	return s.GetToken(V2ParserLessEq, 0)
}

func (s *ExprContext) Or() antlr.TerminalNode {
	return s.GetToken(V2ParserOr, 0)
}

func (s *ExprContext) And() antlr.TerminalNode {
	return s.GetToken(V2ParserAnd, 0)
}

func (s *ExprContext) Pow() antlr.TerminalNode {
	return s.GetToken(V2ParserPow, 0)
}

func (s *ExprContext) Mul() antlr.TerminalNode {
	return s.GetToken(V2ParserMul, 0)
}

func (s *ExprContext) Div() antlr.TerminalNode {
	return s.GetToken(V2ParserDiv, 0)
}

func (s *ExprContext) Mod() antlr.TerminalNode {
	return s.GetToken(V2ParserMod, 0)
}

func (s *ExprContext) Add() antlr.TerminalNode {
	return s.GetToken(V2ParserAdd, 0)
}

func (s *ExprContext) Sub() antlr.TerminalNode {
	return s.GetToken(V2ParserSub, 0)
}

func (s *ExprContext) LSquare() antlr.TerminalNode {
	return s.GetToken(V2ParserLSquare, 0)
}

func (s *ExprContext) RSquare() antlr.TerminalNode {
	return s.GetToken(V2ParserRSquare, 0)
}

func (s *ExprContext) Indexes() IIndexesContext {
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
	_startState := 32
	p.EnterRecursionRule(localctx, 32, V2ParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(211)
			p.FuncCall()
		}

	case 2:
		{
			p.SetState(212)
			p.UnaryOper()
		}
		{
			p.SetState(213)
			p.expr(12)
		}

	case 3:
		{
			p.SetState(215)
			p.Match(V2ParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)
			p.expr(0)
		}
		{
			p.SetState(217)
			p.Match(V2ParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(219)
			p.Match(V2ParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		{
			p.SetState(220)
			p.Literal()
		}

	case 6:
		{
			p.SetState(221)
			p.Initializer()
		}

	case 7:
		{
			p.SetState(222)
			p.AssignStmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(250)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).LHS = _prevctx
				p.PushNewRecursionContext(localctx, _startState, V2ParserRULE_expr)
				p.SetState(225)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(226)
					p.Match(V2ParserDot)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(227)

					var _x = p.expr(16)

					localctx.(*ExprContext).RHS = _x
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).LHS = _prevctx
				p.PushNewRecursionContext(localctx, _startState, V2ParserRULE_expr)
				p.SetState(228)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(229)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).OP = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&64512) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).OP = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(230)

					var _x = p.expr(11)

					localctx.(*ExprContext).RHS = _x
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).LHS = _prevctx
				p.PushNewRecursionContext(localctx, _startState, V2ParserRULE_expr)
				p.SetState(231)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(232)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).OP = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == V2ParserOr || _la == V2ParserAnd) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).OP = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(233)

					var _x = p.expr(10)

					localctx.(*ExprContext).RHS = _x
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).LHS = _prevctx
				p.PushNewRecursionContext(localctx, _startState, V2ParserRULE_expr)
				p.SetState(234)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(235)

					var _m = p.Match(V2ParserPow)

					localctx.(*ExprContext).OP = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(236)

					var _x = p.expr(9)

					localctx.(*ExprContext).RHS = _x
				}

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).LHS = _prevctx
				p.PushNewRecursionContext(localctx, _startState, V2ParserRULE_expr)
				p.SetState(237)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(238)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).OP = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9961472) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).OP = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(239)

					var _x = p.expr(8)

					localctx.(*ExprContext).RHS = _x
				}

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).LHS = _prevctx
				p.PushNewRecursionContext(localctx, _startState, V2ParserRULE_expr)
				p.SetState(240)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(241)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).OP = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == V2ParserAdd || _la == V2ParserSub) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).OP = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(242)

					var _x = p.expr(7)

					localctx.(*ExprContext).RHS = _x
				}

			case 7:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).LHS = _prevctx
				p.PushNewRecursionContext(localctx, _startState, V2ParserRULE_expr)
				p.SetState(243)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(244)
					p.Match(V2ParserLSquare)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(245)

					var _x = p.expr(0)

					localctx.(*ExprContext).Index = _x
				}
				{
					p.SetState(246)
					p.Match(V2ParserRSquare)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 8:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, V2ParserRULE_expr)
				p.SetState(248)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(249)
					p.Indexes()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 34, V2ParserRULE_exprWithLambda)
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(255)
			p.Lambda()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(256)
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

// IFuncCallContext is an interface to support dynamic dispatch.
type IFuncCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
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

func (s *FuncCallContext) Identifier() antlr.TerminalNode {
	return s.GetToken(V2ParserIdentifier, 0)
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
	p.EnterRule(localctx, 36, V2ParserRULE_funcCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.Match(V2ParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(260)
		p.Match(V2ParserLParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&210007196958762) != 0 {
		{
			p.SetState(261)
			p.FuncCallArgs()
		}

	}
	{
		p.SetState(264)
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
	p.EnterRule(localctx, 38, V2ParserRULE_funcCallArgs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.expr(0)
	}
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == V2ParserComma {
		{
			p.SetState(267)
			p.Match(V2ParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(268)
			p.expr(0)
		}

		p.SetState(273)
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
	p.EnterRule(localctx, 40, V2ParserRULE_stmtWithSep)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(274)
		p.Stmt()
	}
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == V2ParserSemi || _la == V2ParserNewLine {
		{
			p.SetState(275)
			p.Sep()
		}

		p.SetState(280)
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
	p.EnterRule(localctx, 42, V2ParserRULE_openStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(281)
		p.Match(V2ParserOpen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(282)
		p.Match(V2ParserStringLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == V2ParserAs {
		{
			p.SetState(283)
			p.Match(V2ParserAs)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(284)
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
	p.EnterRule(localctx, 44, V2ParserRULE_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(287)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&34084860461056) != 0) {
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

// IInitializerContext is an interface to support dynamic dispatch.
type IInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ArrayInitializer() IArrayInitializerContext
	StructInitializer() IStructInitializerContext
	MapInitializer() IMapInitializerContext

	// IsInitializerContext differentiates from other interfaces.
	IsInitializerContext()
}

type InitializerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitializerContext() *InitializerContext {
	var p = new(InitializerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_initializer
	return p
}

func InitEmptyInitializerContext(p *InitializerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_initializer
}

func (*InitializerContext) IsInitializerContext() {}

func NewInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitializerContext {
	var p = new(InitializerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_initializer

	return p
}

func (s *InitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *InitializerContext) ArrayInitializer() IArrayInitializerContext {
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

func (s *InitializerContext) StructInitializer() IStructInitializerContext {
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

func (s *InitializerContext) MapInitializer() IMapInitializerContext {
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

func (s *InitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitializerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterInitializer(s)
	}
}

func (s *InitializerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitInitializer(s)
	}
}

func (s *InitializerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitInitializer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) Initializer() (localctx IInitializerContext) {
	localctx = NewInitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, V2ParserRULE_initializer)
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(289)
			p.ArrayInitializer()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(290)
			p.StructInitializer()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(291)
			p.MapInitializer()
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
	p.EnterRule(localctx, 48, V2ParserRULE_arrayInitializer)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140737891008512) != 0 {
		{
			p.SetState(294)
			p.Type_()
		}

	}
	{
		p.SetState(297)
		p.Match(V2ParserLSquare)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&210007196958762) != 0 {
		{
			p.SetState(298)
			p.expr(0)
		}
		p.SetState(300)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == V2ParserComma {
			{
				p.SetState(299)
				p.Match(V2ParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(307)
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

// IIdentifierPairContext is an interface to support dynamic dispatch.
type IIdentifierPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLHS returns the LHS token.
	GetLHS() antlr.Token

	// SetLHS sets the LHS token.
	SetLHS(antlr.Token)

	// GetRHS returns the RHS rule contexts.
	GetRHS() IExprWithLambdaContext

	// SetRHS sets the RHS rule contexts.
	SetRHS(IExprWithLambdaContext)

	// Getter signatures
	Identifier() antlr.TerminalNode
	Col() antlr.TerminalNode
	To() antlr.TerminalNode
	ExprWithLambda() IExprWithLambdaContext

	// IsIdentifierPairContext differentiates from other interfaces.
	IsIdentifierPairContext()
}

type IdentifierPairContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	LHS    antlr.Token
	RHS    IExprWithLambdaContext
}

func NewEmptyIdentifierPairContext() *IdentifierPairContext {
	var p = new(IdentifierPairContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_identifierPair
	return p
}

func InitEmptyIdentifierPairContext(p *IdentifierPairContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_identifierPair
}

func (*IdentifierPairContext) IsIdentifierPairContext() {}

func NewIdentifierPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierPairContext {
	var p = new(IdentifierPairContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_identifierPair

	return p
}

func (s *IdentifierPairContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierPairContext) GetLHS() antlr.Token { return s.LHS }

func (s *IdentifierPairContext) SetLHS(v antlr.Token) { s.LHS = v }

func (s *IdentifierPairContext) GetRHS() IExprWithLambdaContext { return s.RHS }

func (s *IdentifierPairContext) SetRHS(v IExprWithLambdaContext) { s.RHS = v }

func (s *IdentifierPairContext) Identifier() antlr.TerminalNode {
	return s.GetToken(V2ParserIdentifier, 0)
}

func (s *IdentifierPairContext) Col() antlr.TerminalNode {
	return s.GetToken(V2ParserCol, 0)
}

func (s *IdentifierPairContext) To() antlr.TerminalNode {
	return s.GetToken(V2ParserTo, 0)
}

func (s *IdentifierPairContext) ExprWithLambda() IExprWithLambdaContext {
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

func (s *IdentifierPairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierPairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierPairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterIdentifierPair(s)
	}
}

func (s *IdentifierPairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitIdentifierPair(s)
	}
}

func (s *IdentifierPairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitIdentifierPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) IdentifierPair() (localctx IIdentifierPairContext) {
	localctx = NewIdentifierPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, V2ParserRULE_identifierPair)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(309)

		var _m = p.Match(V2ParserIdentifier)

		localctx.(*IdentifierPairContext).LHS = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(310)
		_la = p.GetTokenStream().LA(1)

		if !(_la == V2ParserCol || _la == V2ParserTo) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(311)

		var _x = p.ExprWithLambda()

		localctx.(*IdentifierPairContext).RHS = _x
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

	// GetLHS returns the LHS rule contexts.
	GetLHS() IExprContext

	// GetRHS returns the RHS rule contexts.
	GetRHS() IExprWithLambdaContext

	// SetLHS sets the LHS rule contexts.
	SetLHS(IExprContext)

	// SetRHS sets the RHS rule contexts.
	SetRHS(IExprWithLambdaContext)

	// Getter signatures
	Expr() IExprContext
	Col() antlr.TerminalNode
	To() antlr.TerminalNode
	ExprWithLambda() IExprWithLambdaContext
	LParen() antlr.TerminalNode
	MapPair() IMapPairContext
	RParen() antlr.TerminalNode
	IdentifierPair() IIdentifierPairContext

	// IsMapPairContext differentiates from other interfaces.
	IsMapPairContext()
}

type MapPairContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	LHS    IExprContext
	RHS    IExprWithLambdaContext
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

func (s *MapPairContext) GetLHS() IExprContext { return s.LHS }

func (s *MapPairContext) GetRHS() IExprWithLambdaContext { return s.RHS }

func (s *MapPairContext) SetLHS(v IExprContext) { s.LHS = v }

func (s *MapPairContext) SetRHS(v IExprWithLambdaContext) { s.RHS = v }

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

func (s *MapPairContext) Col() antlr.TerminalNode {
	return s.GetToken(V2ParserCol, 0)
}

func (s *MapPairContext) To() antlr.TerminalNode {
	return s.GetToken(V2ParserTo, 0)
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

func (s *MapPairContext) LParen() antlr.TerminalNode {
	return s.GetToken(V2ParserLParen, 0)
}

func (s *MapPairContext) MapPair() IMapPairContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapPairContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapPairContext)
}

func (s *MapPairContext) RParen() antlr.TerminalNode {
	return s.GetToken(V2ParserRParen, 0)
}

func (s *MapPairContext) IdentifierPair() IIdentifierPairContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierPairContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierPairContext)
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
	p.EnterRule(localctx, 52, V2ParserRULE_mapPair)
	var _la int

	p.SetState(322)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(313)

			var _x = p.expr(0)

			localctx.(*MapPairContext).LHS = _x
		}
		{
			p.SetState(314)
			_la = p.GetTokenStream().LA(1)

			if !(_la == V2ParserCol || _la == V2ParserTo) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(315)

			var _x = p.ExprWithLambda()

			localctx.(*MapPairContext).RHS = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(317)
			p.Match(V2ParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(318)
			p.MapPair()
		}
		{
			p.SetState(319)
			p.Match(V2ParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(321)
			p.IdentifierPair()
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
	LBrack() antlr.TerminalNode
	RBrack() antlr.TerminalNode
	Type_() ITypeContext
	Struct() antlr.TerminalNode
	AllIdentifierPair() []IIdentifierPairContext
	IdentifierPair(i int) IIdentifierPairContext
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

func (s *StructInitializerContext) LBrack() antlr.TerminalNode {
	return s.GetToken(V2ParserLBrack, 0)
}

func (s *StructInitializerContext) RBrack() antlr.TerminalNode {
	return s.GetToken(V2ParserRBrack, 0)
}

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

func (s *StructInitializerContext) Struct() antlr.TerminalNode {
	return s.GetToken(V2ParserStruct, 0)
}

func (s *StructInitializerContext) AllIdentifierPair() []IIdentifierPairContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierPairContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierPairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierPairContext); ok {
			tst[i] = t.(IIdentifierPairContext)
			i++
		}
	}

	return tst
}

func (s *StructInitializerContext) IdentifierPair(i int) IIdentifierPairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierPairContext); ok {
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

	return t.(IIdentifierPairContext)
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
	p.EnterRule(localctx, 54, V2ParserRULE_structInitializer)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case V2ParserMap, V2ParserFunction, V2ParserIdentifier:
		{
			p.SetState(324)
			p.Type_()
		}

	case V2ParserStruct:
		{
			p.SetState(325)
			p.Match(V2ParserStruct)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(328)
		p.Match(V2ParserLBrack)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == V2ParserIdentifier {
		{
			p.SetState(329)
			p.IdentifierPair()
		}
		p.SetState(331)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == V2ParserComma {
			{
				p.SetState(330)
				p.Match(V2ParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

		p.SetState(337)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(338)
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

// IMapInitializerContext is an interface to support dynamic dispatch.
type IMapInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBrack() antlr.TerminalNode
	RBrack() antlr.TerminalNode
	Map() antlr.TerminalNode
	AllMapPair() []IMapPairContext
	MapPair(i int) IMapPairContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

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

func (s *MapInitializerContext) LBrack() antlr.TerminalNode {
	return s.GetToken(V2ParserLBrack, 0)
}

func (s *MapInitializerContext) RBrack() antlr.TerminalNode {
	return s.GetToken(V2ParserRBrack, 0)
}

func (s *MapInitializerContext) Map() antlr.TerminalNode {
	return s.GetToken(V2ParserMap, 0)
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

func (s *MapInitializerContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(V2ParserComma)
}

func (s *MapInitializerContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(V2ParserComma, i)
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
	p.EnterRule(localctx, 56, V2ParserRULE_mapInitializer)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == V2ParserMap {
		{
			p.SetState(340)
			p.Match(V2ParserMap)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(343)
		p.Match(V2ParserLBrack)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&210007196958762) != 0 {
		{
			p.SetState(344)
			p.MapPair()
		}

		p.SetState(348)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == V2ParserComma {
			{
				p.SetState(345)
				p.Match(V2ParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(350)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(355)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(356)
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

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AssignStmt() IAssignStmtContext
	DeclareStmt() IDeclareStmtContext
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

func (s *StmtContext) AssignStmt() IAssignStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignStmtContext)
}

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
	p.EnterRule(localctx, 58, V2ParserRULE_stmt)
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(358)
			p.AssignStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(359)
			p.DeclareStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(360)
			p.JumpStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(361)
			p.IfStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(362)
			p.LoopStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(363)
			p.MatchStmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(364)
			p.CodeBlock()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(365)
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

// IAssignStmtContext is an interface to support dynamic dispatch.
type IAssignStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_() IVarContext
	Assign() antlr.TerminalNode
	ExprWithLambda() IExprWithLambdaContext
	Type_() ITypeContext

	// IsAssignStmtContext differentiates from other interfaces.
	IsAssignStmtContext()
}

type AssignStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignStmtContext() *AssignStmtContext {
	var p = new(AssignStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_assignStmt
	return p
}

func InitEmptyAssignStmtContext(p *AssignStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_assignStmt
}

func (*AssignStmtContext) IsAssignStmtContext() {}

func NewAssignStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignStmtContext {
	var p = new(AssignStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_assignStmt

	return p
}

func (s *AssignStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignStmtContext) Var_() IVarContext {
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

func (s *AssignStmtContext) Assign() antlr.TerminalNode {
	return s.GetToken(V2ParserAssign, 0)
}

func (s *AssignStmtContext) ExprWithLambda() IExprWithLambdaContext {
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

func (s *AssignStmtContext) Type_() ITypeContext {
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

func (s *AssignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterAssignStmt(s)
	}
}

func (s *AssignStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitAssignStmt(s)
	}
}

func (s *AssignStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitAssignStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) AssignStmt() (localctx IAssignStmtContext) {
	localctx = NewAssignStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, V2ParserRULE_assignStmt)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(369)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(368)
			p.Type_()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(371)
		p.Var_()
	}
	{
		p.SetState(372)
		p.Match(V2ParserAssign)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(373)
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

// IDeclareStmtContext is an interface to support dynamic dispatch.
type IDeclareStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode
	Assign() antlr.TerminalNode
	ExprWithLambda() IExprWithLambdaContext

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

func (s *DeclareStmtContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(V2ParserIdentifier)
}

func (s *DeclareStmtContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(V2ParserIdentifier, i)
}

func (s *DeclareStmtContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(V2ParserComma)
}

func (s *DeclareStmtContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(V2ParserComma, i)
}

func (s *DeclareStmtContext) Assign() antlr.TerminalNode {
	return s.GetToken(V2ParserAssign, 0)
}

func (s *DeclareStmtContext) ExprWithLambda() IExprWithLambdaContext {
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
	p.EnterRule(localctx, 62, V2ParserRULE_declareStmt)
	var _la int

	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(375)
			p.Type_()
		}
		{
			p.SetState(376)
			p.Match(V2ParserIdentifier)
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

		for _la == V2ParserComma {
			{
				p.SetState(377)
				p.Match(V2ParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(378)
				p.Match(V2ParserIdentifier)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(383)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(384)
			p.Type_()
		}
		{
			p.SetState(385)
			p.Match(V2ParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(386)
			p.Match(V2ParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(387)
			p.ExprWithLambda()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(390)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(389)
				p.Type_()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(392)
			p.Match(V2ParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(393)
			p.Match(V2ParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(394)
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
	p.EnterRule(localctx, 64, V2ParserRULE_ifStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(397)
		p.Match(V2ParserIf)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(398)
		p.Match(V2ParserLParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(399)
		p.expr(0)
	}
	{
		p.SetState(400)
		p.Match(V2ParserRParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(401)
		p.CodeBlock()
	}
	p.SetState(404)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == V2ParserElse {
		{
			p.SetState(402)
			p.Match(V2ParserElse)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(403)
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
	IterFor() IIterForContext
	CStyleFor() ICStyleForContext
	WhileStyleFor() IWhileStyleForContext

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

func (s *LoopStmtContext) IterFor() IIterForContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIterForContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIterForContext)
}

func (s *LoopStmtContext) CStyleFor() ICStyleForContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICStyleForContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICStyleForContext)
}

func (s *LoopStmtContext) WhileStyleFor() IWhileStyleForContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStyleForContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStyleForContext)
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
	p.EnterRule(localctx, 66, V2ParserRULE_loopStmt)
	p.SetState(409)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(406)
			p.IterFor()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(407)
			p.CStyleFor()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(408)
			p.WhileStyleFor()
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

// ICStyleForContext is an interface to support dynamic dispatch.
type ICStyleForContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOnInit returns the onInit rule contexts.
	GetOnInit() IExprContext

	// GetOnCondition returns the onCondition rule contexts.
	GetOnCondition() IExprContext

	// GetOnEnd returns the onEnd rule contexts.
	GetOnEnd() IExprContext

	// SetOnInit sets the onInit rule contexts.
	SetOnInit(IExprContext)

	// SetOnCondition sets the onCondition rule contexts.
	SetOnCondition(IExprContext)

	// SetOnEnd sets the onEnd rule contexts.
	SetOnEnd(IExprContext)

	// Getter signatures
	For() antlr.TerminalNode
	LParen() antlr.TerminalNode
	AllSemi() []antlr.TerminalNode
	Semi(i int) antlr.TerminalNode
	RParen() antlr.TerminalNode
	CodeBlock() ICodeBlockContext
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsCStyleForContext differentiates from other interfaces.
	IsCStyleForContext()
}

type CStyleForContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	onInit      IExprContext
	onCondition IExprContext
	onEnd       IExprContext
}

func NewEmptyCStyleForContext() *CStyleForContext {
	var p = new(CStyleForContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_cStyleFor
	return p
}

func InitEmptyCStyleForContext(p *CStyleForContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_cStyleFor
}

func (*CStyleForContext) IsCStyleForContext() {}

func NewCStyleForContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CStyleForContext {
	var p = new(CStyleForContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_cStyleFor

	return p
}

func (s *CStyleForContext) GetParser() antlr.Parser { return s.parser }

func (s *CStyleForContext) GetOnInit() IExprContext { return s.onInit }

func (s *CStyleForContext) GetOnCondition() IExprContext { return s.onCondition }

func (s *CStyleForContext) GetOnEnd() IExprContext { return s.onEnd }

func (s *CStyleForContext) SetOnInit(v IExprContext) { s.onInit = v }

func (s *CStyleForContext) SetOnCondition(v IExprContext) { s.onCondition = v }

func (s *CStyleForContext) SetOnEnd(v IExprContext) { s.onEnd = v }

func (s *CStyleForContext) For() antlr.TerminalNode {
	return s.GetToken(V2ParserFor, 0)
}

func (s *CStyleForContext) LParen() antlr.TerminalNode {
	return s.GetToken(V2ParserLParen, 0)
}

func (s *CStyleForContext) AllSemi() []antlr.TerminalNode {
	return s.GetTokens(V2ParserSemi)
}

func (s *CStyleForContext) Semi(i int) antlr.TerminalNode {
	return s.GetToken(V2ParserSemi, i)
}

func (s *CStyleForContext) RParen() antlr.TerminalNode {
	return s.GetToken(V2ParserRParen, 0)
}

func (s *CStyleForContext) CodeBlock() ICodeBlockContext {
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

func (s *CStyleForContext) AllExpr() []IExprContext {
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

func (s *CStyleForContext) Expr(i int) IExprContext {
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

func (s *CStyleForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CStyleForContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CStyleForContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterCStyleFor(s)
	}
}

func (s *CStyleForContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitCStyleFor(s)
	}
}

func (s *CStyleForContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitCStyleFor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) CStyleFor() (localctx ICStyleForContext) {
	localctx = NewCStyleForContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, V2ParserRULE_cStyleFor)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(411)
		p.Match(V2ParserFor)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(412)
		p.Match(V2ParserLParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&210007196958762) != 0 {
		{
			p.SetState(413)

			var _x = p.expr(0)

			localctx.(*CStyleForContext).onInit = _x
		}

	}
	{
		p.SetState(416)
		p.Match(V2ParserSemi)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(418)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&210007196958762) != 0 {
		{
			p.SetState(417)

			var _x = p.expr(0)

			localctx.(*CStyleForContext).onCondition = _x
		}

	}
	{
		p.SetState(420)
		p.Match(V2ParserSemi)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(422)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&210007196958762) != 0 {
		{
			p.SetState(421)

			var _x = p.expr(0)

			localctx.(*CStyleForContext).onEnd = _x
		}

	}
	{
		p.SetState(424)
		p.Match(V2ParserRParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(425)
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

// IIterForContext is an interface to support dynamic dispatch.
type IIterForContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	For() antlr.TerminalNode
	LParen() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	Col() antlr.TerminalNode
	Expr() IExprContext
	RParen() antlr.TerminalNode
	CodeBlock() ICodeBlockContext
	Type_() ITypeContext

	// IsIterForContext differentiates from other interfaces.
	IsIterForContext()
}

type IterForContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIterForContext() *IterForContext {
	var p = new(IterForContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_iterFor
	return p
}

func InitEmptyIterForContext(p *IterForContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_iterFor
}

func (*IterForContext) IsIterForContext() {}

func NewIterForContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IterForContext {
	var p = new(IterForContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_iterFor

	return p
}

func (s *IterForContext) GetParser() antlr.Parser { return s.parser }

func (s *IterForContext) For() antlr.TerminalNode {
	return s.GetToken(V2ParserFor, 0)
}

func (s *IterForContext) LParen() antlr.TerminalNode {
	return s.GetToken(V2ParserLParen, 0)
}

func (s *IterForContext) Identifier() antlr.TerminalNode {
	return s.GetToken(V2ParserIdentifier, 0)
}

func (s *IterForContext) Col() antlr.TerminalNode {
	return s.GetToken(V2ParserCol, 0)
}

func (s *IterForContext) Expr() IExprContext {
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

func (s *IterForContext) RParen() antlr.TerminalNode {
	return s.GetToken(V2ParserRParen, 0)
}

func (s *IterForContext) CodeBlock() ICodeBlockContext {
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

func (s *IterForContext) Type_() ITypeContext {
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

func (s *IterForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IterForContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IterForContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterIterFor(s)
	}
}

func (s *IterForContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitIterFor(s)
	}
}

func (s *IterForContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitIterFor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) IterFor() (localctx IIterForContext) {
	localctx = NewIterForContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, V2ParserRULE_iterFor)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(427)
		p.Match(V2ParserFor)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(428)
		p.Match(V2ParserLParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(430)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(429)
			p.Type_()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(432)
		p.Match(V2ParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(433)
		p.Match(V2ParserCol)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(434)
		p.expr(0)
	}
	{
		p.SetState(435)
		p.Match(V2ParserRParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(436)
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

// IWhileStyleForContext is an interface to support dynamic dispatch.
type IWhileStyleForContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	For() antlr.TerminalNode
	CodeBlock() ICodeBlockContext
	LParen() antlr.TerminalNode
	RParen() antlr.TerminalNode
	Expr() IExprContext

	// IsWhileStyleForContext differentiates from other interfaces.
	IsWhileStyleForContext()
}

type WhileStyleForContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStyleForContext() *WhileStyleForContext {
	var p = new(WhileStyleForContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_whileStyleFor
	return p
}

func InitEmptyWhileStyleForContext(p *WhileStyleForContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_whileStyleFor
}

func (*WhileStyleForContext) IsWhileStyleForContext() {}

func NewWhileStyleForContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStyleForContext {
	var p = new(WhileStyleForContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_whileStyleFor

	return p
}

func (s *WhileStyleForContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStyleForContext) For() antlr.TerminalNode {
	return s.GetToken(V2ParserFor, 0)
}

func (s *WhileStyleForContext) CodeBlock() ICodeBlockContext {
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

func (s *WhileStyleForContext) LParen() antlr.TerminalNode {
	return s.GetToken(V2ParserLParen, 0)
}

func (s *WhileStyleForContext) RParen() antlr.TerminalNode {
	return s.GetToken(V2ParserRParen, 0)
}

func (s *WhileStyleForContext) Expr() IExprContext {
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

func (s *WhileStyleForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStyleForContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStyleForContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.EnterWhileStyleFor(s)
	}
}

func (s *WhileStyleForContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(V2ParserListener); ok {
		listenerT.ExitWhileStyleFor(s)
	}
}

func (s *WhileStyleForContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitWhileStyleFor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) WhileStyleFor() (localctx IWhileStyleForContext) {
	localctx = NewWhileStyleForContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, V2ParserRULE_whileStyleFor)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(438)
		p.Match(V2ParserFor)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(444)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == V2ParserLParen {
		{
			p.SetState(439)
			p.Match(V2ParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(441)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&210007196958762) != 0 {
			{
				p.SetState(440)
				p.expr(0)
			}

		}
		{
			p.SetState(443)
			p.Match(V2ParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(446)
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
	p.EnterRule(localctx, 74, V2ParserRULE_matchStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(448)
		p.Match(V2ParserMatch)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(449)
		p.Match(V2ParserLParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(450)
		p.expr(0)
	}
	{
		p.SetState(451)
		p.Match(V2ParserRParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(452)
		p.Match(V2ParserLBrack)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(456)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == V2ParserCase || _la == V2ParserDefault {
		{
			p.SetState(453)
			p.MatchCase()
		}

		p.SetState(458)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(459)
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
	p.EnterRule(localctx, 76, V2ParserRULE_matchCase)
	p.SetState(469)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case V2ParserCase:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(461)
			p.Match(V2ParserCase)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(462)
			p.expr(0)
		}
		{
			p.SetState(463)
			p.Match(V2ParserCol)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(464)
			p.CodeBlock()
		}

	case V2ParserDefault:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(466)
			p.Match(V2ParserDefault)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(467)
			p.Match(V2ParserCol)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(468)
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
	ExprWithLambda() IExprWithLambdaContext

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

func (s *JumpStmtContext) ExprWithLambda() IExprWithLambdaContext {
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
	p.EnterRule(localctx, 78, V2ParserRULE_jumpStmt)
	p.SetState(477)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case V2ParserContinue:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(471)
			p.Match(V2ParserContinue)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case V2ParserBreak:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(472)
			p.Match(V2ParserBreak)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case V2ParserReturn:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(473)
			p.Match(V2ParserReturn)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(475)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(474)
				p.ExprWithLambda()
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
	p.EnterRule(localctx, 80, V2ParserRULE_sep)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(479)
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
	case 16:
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
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
