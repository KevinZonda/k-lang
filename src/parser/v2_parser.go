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
		"program", "structBlock", "funcBlock", "codeBlock", "declareBlock",
		"funcSig", "funcReturnType", "funcSignArgs", "funcSignArgItem", "type",
		"var", "baseVar", "index", "indexes", "lambda", "unaryOper", "expr",
		"exprWithLambda", "funcCall", "funcCallArgs", "openStmt", "literal",
		"initializer", "arrayInitializer", "identifierPair", "mapPair", "structInitializer",
		"mapInitializer", "commaExpr", "stmt", "uOperStmt", "assignStmt", "vars",
		"typedIdentifiers", "typedIdentifier", "declareStmt", "ifStmt", "tryCatchSmt",
		"catchStmt", "loopStmt", "cStyleFor", "cStyleForSign", "iterFor", "iterForSign",
		"whileStyleFor", "matchStmt", "matchCase", "caseBlock", "jumpStmt",
		"sep",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 58, 676, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 105, 8,
		0, 1, 0, 5, 0, 108, 8, 0, 10, 0, 12, 0, 111, 9, 0, 5, 0, 113, 8, 0, 10,
		0, 12, 0, 116, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 3, 3, 3, 131, 8, 3, 1, 3, 5, 3, 134, 8, 3, 10,
		3, 12, 3, 137, 9, 3, 5, 3, 139, 8, 3, 10, 3, 12, 3, 142, 9, 3, 1, 3, 1,
		3, 1, 4, 1, 4, 5, 4, 148, 8, 4, 10, 4, 12, 4, 151, 9, 4, 1, 4, 1, 4, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 160, 8, 5, 1, 6, 3, 6, 163, 8, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 170, 8, 6, 10, 6, 12, 6, 173, 9, 6, 1, 6,
		1, 6, 3, 6, 177, 8, 6, 1, 7, 3, 7, 180, 8, 7, 1, 7, 1, 7, 5, 7, 184, 8,
		7, 10, 7, 12, 7, 187, 9, 7, 1, 8, 3, 8, 190, 8, 8, 1, 8, 3, 8, 193, 8,
		8, 1, 8, 1, 8, 3, 8, 197, 8, 8, 1, 8, 1, 8, 1, 8, 3, 8, 202, 8, 8, 3, 8,
		204, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 210, 8, 9, 1, 9, 1, 9, 1, 9, 3,
		9, 215, 8, 9, 1, 9, 3, 9, 218, 8, 9, 3, 9, 220, 8, 9, 1, 10, 1, 10, 1,
		10, 5, 10, 225, 8, 10, 10, 10, 12, 10, 228, 9, 10, 1, 11, 1, 11, 3, 11,
		232, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 4, 13, 239, 8, 13, 11, 13,
		12, 13, 240, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 248, 8, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 256, 8, 14, 1, 14, 1, 14, 1,
		14, 3, 14, 261, 8, 14, 3, 14, 263, 8, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 3, 16, 282, 8, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 3, 16, 312, 8, 16, 1, 16, 3, 16, 315, 8, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 3, 16, 322, 8, 16, 1, 16, 1, 16, 1, 16, 5, 16, 327, 8,
		16, 10, 16, 12, 16, 330, 9, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 345, 8, 17, 1,
		17, 1, 17, 1, 17, 3, 17, 350, 8, 17, 1, 17, 5, 17, 353, 8, 17, 10, 17,
		12, 17, 356, 9, 17, 1, 18, 1, 18, 1, 18, 3, 18, 361, 8, 18, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 19, 5, 19, 368, 8, 19, 10, 19, 12, 19, 371, 9, 19, 1,
		20, 1, 20, 1, 20, 1, 20, 3, 20, 377, 8, 20, 1, 21, 1, 21, 1, 22, 1, 22,
		3, 22, 383, 8, 22, 1, 23, 3, 23, 386, 8, 23, 1, 23, 1, 23, 1, 23, 3, 23,
		391, 8, 23, 5, 23, 393, 8, 23, 10, 23, 12, 23, 396, 9, 23, 1, 23, 1, 23,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 3, 25, 413, 8, 25, 1, 26, 1, 26, 3, 26, 417, 8, 26, 1,
		26, 1, 26, 1, 26, 3, 26, 422, 8, 26, 5, 26, 424, 8, 26, 10, 26, 12, 26,
		427, 9, 26, 1, 26, 1, 26, 1, 27, 3, 27, 432, 8, 27, 1, 27, 1, 27, 1, 27,
		5, 27, 437, 8, 27, 10, 27, 12, 27, 440, 9, 27, 5, 27, 442, 8, 27, 10, 27,
		12, 27, 445, 9, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 5, 28, 452, 8, 28,
		10, 28, 12, 28, 455, 9, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 468, 8, 29, 1, 30, 1, 30, 1, 30,
		1, 31, 3, 31, 474, 8, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 490, 8, 31,
		1, 32, 1, 32, 1, 32, 5, 32, 495, 8, 32, 10, 32, 12, 32, 498, 9, 32, 1,
		33, 1, 33, 1, 33, 1, 33, 5, 33, 504, 8, 33, 10, 33, 12, 33, 507, 9, 33,
		1, 33, 1, 33, 1, 33, 5, 33, 512, 8, 33, 10, 33, 12, 33, 515, 9, 33, 1,
		33, 1, 33, 3, 33, 519, 8, 33, 1, 34, 3, 34, 522, 8, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 3, 34, 528, 8, 34, 3, 34, 530, 8, 34, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 3, 35, 538, 8, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 3, 36, 546, 8, 36, 3, 36, 548, 8, 36, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 38, 1, 38, 1, 38, 3, 38, 557, 8, 38, 1, 38, 1, 38, 3, 38, 561, 8, 38,
		1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 3, 39, 568, 8, 39, 1, 40, 1, 40, 1,
		40, 1, 40, 1, 40, 1, 40, 3, 40, 576, 8, 40, 1, 40, 1, 40, 1, 41, 3, 41,
		581, 8, 41, 1, 41, 1, 41, 3, 41, 585, 8, 41, 1, 41, 1, 41, 3, 41, 589,
		8, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 3, 42, 597, 8, 42, 1,
		42, 1, 42, 1, 43, 3, 43, 602, 8, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 44,
		1, 44, 3, 44, 610, 8, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 5,
		45, 618, 8, 45, 10, 45, 12, 45, 621, 9, 45, 1, 45, 1, 45, 1, 46, 1, 46,
		1, 46, 1, 46, 5, 46, 629, 8, 46, 10, 46, 12, 46, 632, 9, 46, 1, 46, 1,
		46, 1, 46, 1, 46, 3, 46, 638, 8, 46, 1, 47, 1, 47, 1, 47, 1, 47, 3, 47,
		644, 8, 47, 1, 47, 5, 47, 647, 8, 47, 10, 47, 12, 47, 650, 9, 47, 5, 47,
		652, 8, 47, 10, 47, 12, 47, 655, 9, 47, 3, 47, 657, 8, 47, 1, 48, 1, 48,
		1, 48, 1, 48, 1, 48, 1, 48, 5, 48, 665, 8, 48, 10, 48, 12, 48, 668, 9,
		48, 3, 48, 670, 8, 48, 3, 48, 672, 8, 48, 1, 49, 1, 49, 1, 49, 0, 2, 32,
		34, 50, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68,
		70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 0, 10, 2, 0,
		8, 8, 28, 28, 2, 0, 21, 22, 51, 51, 2, 0, 19, 20, 23, 23, 1, 0, 21, 22,
		1, 0, 10, 15, 1, 0, 16, 17, 2, 0, 24, 24, 54, 54, 1, 0, 45, 50, 1, 0, 26,
		27, 2, 0, 9, 9, 58, 58, 744, 0, 114, 1, 0, 0, 0, 2, 119, 1, 0, 0, 0, 4,
		123, 1, 0, 0, 0, 6, 127, 1, 0, 0, 0, 8, 145, 1, 0, 0, 0, 10, 154, 1, 0,
		0, 0, 12, 162, 1, 0, 0, 0, 14, 179, 1, 0, 0, 0, 16, 203, 1, 0, 0, 0, 18,
		219, 1, 0, 0, 0, 20, 221, 1, 0, 0, 0, 22, 229, 1, 0, 0, 0, 24, 233, 1,
		0, 0, 0, 26, 238, 1, 0, 0, 0, 28, 262, 1, 0, 0, 0, 30, 264, 1, 0, 0, 0,
		32, 281, 1, 0, 0, 0, 34, 344, 1, 0, 0, 0, 36, 357, 1, 0, 0, 0, 38, 364,
		1, 0, 0, 0, 40, 372, 1, 0, 0, 0, 42, 378, 1, 0, 0, 0, 44, 382, 1, 0, 0,
		0, 46, 385, 1, 0, 0, 0, 48, 399, 1, 0, 0, 0, 50, 412, 1, 0, 0, 0, 52, 416,
		1, 0, 0, 0, 54, 431, 1, 0, 0, 0, 56, 448, 1, 0, 0, 0, 58, 467, 1, 0, 0,
		0, 60, 469, 1, 0, 0, 0, 62, 489, 1, 0, 0, 0, 64, 491, 1, 0, 0, 0, 66, 518,
		1, 0, 0, 0, 68, 529, 1, 0, 0, 0, 70, 537, 1, 0, 0, 0, 72, 539, 1, 0, 0,
		0, 74, 549, 1, 0, 0, 0, 76, 553, 1, 0, 0, 0, 78, 567, 1, 0, 0, 0, 80, 569,
		1, 0, 0, 0, 82, 580, 1, 0, 0, 0, 84, 590, 1, 0, 0, 0, 86, 601, 1, 0, 0,
		0, 88, 607, 1, 0, 0, 0, 90, 613, 1, 0, 0, 0, 92, 637, 1, 0, 0, 0, 94, 656,
		1, 0, 0, 0, 96, 671, 1, 0, 0, 0, 98, 673, 1, 0, 0, 0, 100, 105, 3, 2, 1,
		0, 101, 105, 3, 4, 2, 0, 102, 105, 3, 58, 29, 0, 103, 105, 3, 56, 28, 0,
		104, 100, 1, 0, 0, 0, 104, 101, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 104,
		103, 1, 0, 0, 0, 105, 109, 1, 0, 0, 0, 106, 108, 3, 98, 49, 0, 107, 106,
		1, 0, 0, 0, 108, 111, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 110, 1, 0,
		0, 0, 110, 113, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 112, 104, 1, 0, 0, 0,
		113, 116, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115,
		117, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 117, 118, 5, 0, 0, 1, 118, 1, 1,
		0, 0, 0, 119, 120, 5, 29, 0, 0, 120, 121, 5, 54, 0, 0, 121, 122, 3, 8,
		4, 0, 122, 3, 1, 0, 0, 0, 123, 124, 5, 31, 0, 0, 124, 125, 3, 10, 5, 0,
		125, 126, 3, 6, 3, 0, 126, 5, 1, 0, 0, 0, 127, 140, 5, 1, 0, 0, 128, 131,
		3, 58, 29, 0, 129, 131, 3, 32, 16, 0, 130, 128, 1, 0, 0, 0, 130, 129, 1,
		0, 0, 0, 131, 135, 1, 0, 0, 0, 132, 134, 3, 98, 49, 0, 133, 132, 1, 0,
		0, 0, 134, 137, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0,
		136, 139, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 138, 130, 1, 0, 0, 0, 139,
		142, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 143,
		1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 143, 144, 5, 2, 0, 0, 144, 7, 1, 0, 0,
		0, 145, 149, 5, 1, 0, 0, 146, 148, 3, 70, 35, 0, 147, 146, 1, 0, 0, 0,
		148, 151, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150,
		152, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 152, 153, 5, 2, 0, 0, 153, 9, 1,
		0, 0, 0, 154, 155, 5, 54, 0, 0, 155, 156, 5, 3, 0, 0, 156, 157, 3, 14,
		7, 0, 157, 159, 5, 4, 0, 0, 158, 160, 3, 12, 6, 0, 159, 158, 1, 0, 0, 0,
		159, 160, 1, 0, 0, 0, 160, 11, 1, 0, 0, 0, 161, 163, 7, 0, 0, 0, 162, 161,
		1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 176, 1, 0, 0, 0, 164, 177, 3, 18,
		9, 0, 165, 166, 5, 3, 0, 0, 166, 171, 3, 18, 9, 0, 167, 168, 5, 7, 0, 0,
		168, 170, 3, 18, 9, 0, 169, 167, 1, 0, 0, 0, 170, 173, 1, 0, 0, 0, 171,
		169, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 174, 1, 0, 0, 0, 173, 171,
		1, 0, 0, 0, 174, 175, 5, 4, 0, 0, 175, 177, 1, 0, 0, 0, 176, 164, 1, 0,
		0, 0, 176, 165, 1, 0, 0, 0, 177, 13, 1, 0, 0, 0, 178, 180, 3, 16, 8, 0,
		179, 178, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 185, 1, 0, 0, 0, 181,
		182, 5, 7, 0, 0, 182, 184, 3, 16, 8, 0, 183, 181, 1, 0, 0, 0, 184, 187,
		1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 15, 1, 0,
		0, 0, 187, 185, 1, 0, 0, 0, 188, 190, 3, 18, 9, 0, 189, 188, 1, 0, 0, 0,
		189, 190, 1, 0, 0, 0, 190, 192, 1, 0, 0, 0, 191, 193, 5, 53, 0, 0, 192,
		191, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 204,
		5, 54, 0, 0, 195, 197, 5, 53, 0, 0, 196, 195, 1, 0, 0, 0, 196, 197, 1,
		0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 201, 5, 54, 0, 0, 199, 200, 5, 8, 0,
		0, 200, 202, 3, 18, 9, 0, 201, 199, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202,
		204, 1, 0, 0, 0, 203, 189, 1, 0, 0, 0, 203, 196, 1, 0, 0, 0, 204, 17, 1,
		0, 0, 0, 205, 220, 5, 30, 0, 0, 206, 220, 5, 31, 0, 0, 207, 208, 5, 54,
		0, 0, 208, 210, 5, 24, 0, 0, 209, 207, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0,
		210, 211, 1, 0, 0, 0, 211, 214, 5, 54, 0, 0, 212, 213, 5, 5, 0, 0, 213,
		215, 5, 6, 0, 0, 214, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 217,
		1, 0, 0, 0, 216, 218, 5, 25, 0, 0, 217, 216, 1, 0, 0, 0, 217, 218, 1, 0,
		0, 0, 218, 220, 1, 0, 0, 0, 219, 205, 1, 0, 0, 0, 219, 206, 1, 0, 0, 0,
		219, 209, 1, 0, 0, 0, 220, 19, 1, 0, 0, 0, 221, 226, 3, 22, 11, 0, 222,
		223, 5, 24, 0, 0, 223, 225, 3, 22, 11, 0, 224, 222, 1, 0, 0, 0, 225, 228,
		1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 21, 1, 0,
		0, 0, 228, 226, 1, 0, 0, 0, 229, 231, 5, 54, 0, 0, 230, 232, 3, 26, 13,
		0, 231, 230, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 23, 1, 0, 0, 0, 233,
		234, 5, 5, 0, 0, 234, 235, 3, 32, 16, 0, 235, 236, 5, 6, 0, 0, 236, 25,
		1, 0, 0, 0, 237, 239, 3, 24, 12, 0, 238, 237, 1, 0, 0, 0, 239, 240, 1,
		0, 0, 0, 240, 238, 1, 0, 0, 0, 240, 241, 1, 0, 0, 0, 241, 27, 1, 0, 0,
		0, 242, 243, 5, 31, 0, 0, 243, 244, 5, 3, 0, 0, 244, 245, 3, 14, 7, 0,
		245, 247, 5, 4, 0, 0, 246, 248, 3, 12, 6, 0, 247, 246, 1, 0, 0, 0, 247,
		248, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 250, 3, 6, 3, 0, 250, 263,
		1, 0, 0, 0, 251, 252, 5, 3, 0, 0, 252, 253, 3, 14, 7, 0, 253, 255, 5, 4,
		0, 0, 254, 256, 3, 12, 6, 0, 255, 254, 1, 0, 0, 0, 255, 256, 1, 0, 0, 0,
		256, 257, 1, 0, 0, 0, 257, 260, 5, 28, 0, 0, 258, 261, 3, 6, 3, 0, 259,
		261, 3, 34, 17, 0, 260, 258, 1, 0, 0, 0, 260, 259, 1, 0, 0, 0, 261, 263,
		1, 0, 0, 0, 262, 242, 1, 0, 0, 0, 262, 251, 1, 0, 0, 0, 263, 29, 1, 0,
		0, 0, 264, 265, 7, 1, 0, 0, 265, 31, 1, 0, 0, 0, 266, 267, 6, 16, -1, 0,
		267, 268, 5, 53, 0, 0, 268, 282, 3, 32, 16, 18, 269, 282, 3, 36, 18, 0,
		270, 271, 3, 30, 15, 0, 271, 272, 3, 32, 16, 13, 272, 282, 1, 0, 0, 0,
		273, 274, 5, 3, 0, 0, 274, 275, 3, 32, 16, 0, 275, 276, 5, 4, 0, 0, 276,
		282, 1, 0, 0, 0, 277, 282, 3, 52, 26, 0, 278, 282, 5, 54, 0, 0, 279, 282,
		3, 42, 21, 0, 280, 282, 3, 44, 22, 0, 281, 266, 1, 0, 0, 0, 281, 269, 1,
		0, 0, 0, 281, 270, 1, 0, 0, 0, 281, 273, 1, 0, 0, 0, 281, 277, 1, 0, 0,
		0, 281, 278, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0, 281, 280, 1, 0, 0, 0, 282,
		328, 1, 0, 0, 0, 283, 284, 10, 17, 0, 0, 284, 285, 5, 24, 0, 0, 285, 327,
		3, 32, 16, 18, 286, 287, 10, 11, 0, 0, 287, 288, 5, 18, 0, 0, 288, 327,
		3, 32, 16, 12, 289, 290, 10, 10, 0, 0, 290, 291, 7, 2, 0, 0, 291, 327,
		3, 32, 16, 11, 292, 293, 10, 9, 0, 0, 293, 294, 7, 3, 0, 0, 294, 327, 3,
		32, 16, 10, 295, 296, 10, 8, 0, 0, 296, 297, 7, 4, 0, 0, 297, 327, 3, 32,
		16, 9, 298, 299, 10, 7, 0, 0, 299, 300, 7, 5, 0, 0, 300, 327, 3, 32, 16,
		8, 301, 302, 10, 1, 0, 0, 302, 303, 5, 25, 0, 0, 303, 304, 3, 32, 16, 0,
		304, 305, 5, 8, 0, 0, 305, 306, 3, 32, 16, 2, 306, 327, 1, 0, 0, 0, 307,
		308, 10, 16, 0, 0, 308, 309, 5, 5, 0, 0, 309, 311, 3, 32, 16, 0, 310, 312,
		5, 8, 0, 0, 311, 310, 1, 0, 0, 0, 311, 312, 1, 0, 0, 0, 312, 314, 1, 0,
		0, 0, 313, 315, 3, 32, 16, 0, 314, 313, 1, 0, 0, 0, 314, 315, 1, 0, 0,
		0, 315, 316, 1, 0, 0, 0, 316, 317, 5, 6, 0, 0, 317, 327, 1, 0, 0, 0, 318,
		319, 10, 14, 0, 0, 319, 321, 5, 3, 0, 0, 320, 322, 3, 38, 19, 0, 321, 320,
		1, 0, 0, 0, 321, 322, 1, 0, 0, 0, 322, 323, 1, 0, 0, 0, 323, 327, 5, 4,
		0, 0, 324, 325, 10, 2, 0, 0, 325, 327, 3, 26, 13, 0, 326, 283, 1, 0, 0,
		0, 326, 286, 1, 0, 0, 0, 326, 289, 1, 0, 0, 0, 326, 292, 1, 0, 0, 0, 326,
		295, 1, 0, 0, 0, 326, 298, 1, 0, 0, 0, 326, 301, 1, 0, 0, 0, 326, 307,
		1, 0, 0, 0, 326, 318, 1, 0, 0, 0, 326, 324, 1, 0, 0, 0, 327, 330, 1, 0,
		0, 0, 328, 326, 1, 0, 0, 0, 328, 329, 1, 0, 0, 0, 329, 33, 1, 0, 0, 0,
		330, 328, 1, 0, 0, 0, 331, 332, 6, 17, -1, 0, 332, 345, 3, 28, 14, 0, 333,
		345, 3, 32, 16, 0, 334, 335, 3, 32, 16, 0, 335, 336, 5, 25, 0, 0, 336,
		337, 3, 34, 17, 0, 337, 338, 5, 8, 0, 0, 338, 339, 3, 34, 17, 3, 339, 345,
		1, 0, 0, 0, 340, 341, 5, 3, 0, 0, 341, 342, 3, 34, 17, 0, 342, 343, 5,
		4, 0, 0, 343, 345, 1, 0, 0, 0, 344, 331, 1, 0, 0, 0, 344, 333, 1, 0, 0,
		0, 344, 334, 1, 0, 0, 0, 344, 340, 1, 0, 0, 0, 345, 354, 1, 0, 0, 0, 346,
		347, 10, 2, 0, 0, 347, 349, 5, 3, 0, 0, 348, 350, 3, 38, 19, 0, 349, 348,
		1, 0, 0, 0, 349, 350, 1, 0, 0, 0, 350, 351, 1, 0, 0, 0, 351, 353, 5, 4,
		0, 0, 352, 346, 1, 0, 0, 0, 353, 356, 1, 0, 0, 0, 354, 352, 1, 0, 0, 0,
		354, 355, 1, 0, 0, 0, 355, 35, 1, 0, 0, 0, 356, 354, 1, 0, 0, 0, 357, 358,
		5, 54, 0, 0, 358, 360, 5, 3, 0, 0, 359, 361, 3, 38, 19, 0, 360, 359, 1,
		0, 0, 0, 360, 361, 1, 0, 0, 0, 361, 362, 1, 0, 0, 0, 362, 363, 5, 4, 0,
		0, 363, 37, 1, 0, 0, 0, 364, 369, 3, 34, 17, 0, 365, 366, 5, 7, 0, 0, 366,
		368, 3, 34, 17, 0, 367, 365, 1, 0, 0, 0, 368, 371, 1, 0, 0, 0, 369, 367,
		1, 0, 0, 0, 369, 370, 1, 0, 0, 0, 370, 39, 1, 0, 0, 0, 371, 369, 1, 0,
		0, 0, 372, 373, 5, 35, 0, 0, 373, 376, 5, 50, 0, 0, 374, 375, 5, 36, 0,
		0, 375, 377, 7, 6, 0, 0, 376, 374, 1, 0, 0, 0, 376, 377, 1, 0, 0, 0, 377,
		41, 1, 0, 0, 0, 378, 379, 7, 7, 0, 0, 379, 43, 1, 0, 0, 0, 380, 383, 3,
		46, 23, 0, 381, 383, 3, 54, 27, 0, 382, 380, 1, 0, 0, 0, 382, 381, 1, 0,
		0, 0, 383, 45, 1, 0, 0, 0, 384, 386, 3, 18, 9, 0, 385, 384, 1, 0, 0, 0,
		385, 386, 1, 0, 0, 0, 386, 387, 1, 0, 0, 0, 387, 394, 5, 5, 0, 0, 388,
		390, 3, 32, 16, 0, 389, 391, 5, 7, 0, 0, 390, 389, 1, 0, 0, 0, 390, 391,
		1, 0, 0, 0, 391, 393, 1, 0, 0, 0, 392, 388, 1, 0, 0, 0, 393, 396, 1, 0,
		0, 0, 394, 392, 1, 0, 0, 0, 394, 395, 1, 0, 0, 0, 395, 397, 1, 0, 0, 0,
		396, 394, 1, 0, 0, 0, 397, 398, 5, 6, 0, 0, 398, 47, 1, 0, 0, 0, 399, 400,
		5, 54, 0, 0, 400, 401, 7, 0, 0, 0, 401, 402, 3, 34, 17, 0, 402, 49, 1,
		0, 0, 0, 403, 404, 3, 32, 16, 0, 404, 405, 7, 0, 0, 0, 405, 406, 3, 34,
		17, 0, 406, 413, 1, 0, 0, 0, 407, 408, 5, 3, 0, 0, 408, 409, 3, 50, 25,
		0, 409, 410, 5, 4, 0, 0, 410, 413, 1, 0, 0, 0, 411, 413, 3, 48, 24, 0,
		412, 403, 1, 0, 0, 0, 412, 407, 1, 0, 0, 0, 412, 411, 1, 0, 0, 0, 413,
		51, 1, 0, 0, 0, 414, 417, 3, 18, 9, 0, 415, 417, 5, 29, 0, 0, 416, 414,
		1, 0, 0, 0, 416, 415, 1, 0, 0, 0, 417, 418, 1, 0, 0, 0, 418, 425, 5, 1,
		0, 0, 419, 421, 3, 48, 24, 0, 420, 422, 5, 7, 0, 0, 421, 420, 1, 0, 0,
		0, 421, 422, 1, 0, 0, 0, 422, 424, 1, 0, 0, 0, 423, 419, 1, 0, 0, 0, 424,
		427, 1, 0, 0, 0, 425, 423, 1, 0, 0, 0, 425, 426, 1, 0, 0, 0, 426, 428,
		1, 0, 0, 0, 427, 425, 1, 0, 0, 0, 428, 429, 5, 2, 0, 0, 429, 53, 1, 0,
		0, 0, 430, 432, 5, 30, 0, 0, 431, 430, 1, 0, 0, 0, 431, 432, 1, 0, 0, 0,
		432, 433, 1, 0, 0, 0, 433, 443, 5, 1, 0, 0, 434, 438, 3, 50, 25, 0, 435,
		437, 5, 7, 0, 0, 436, 435, 1, 0, 0, 0, 437, 440, 1, 0, 0, 0, 438, 436,
		1, 0, 0, 0, 438, 439, 1, 0, 0, 0, 439, 442, 1, 0, 0, 0, 440, 438, 1, 0,
		0, 0, 441, 434, 1, 0, 0, 0, 442, 445, 1, 0, 0, 0, 443, 441, 1, 0, 0, 0,
		443, 444, 1, 0, 0, 0, 444, 446, 1, 0, 0, 0, 445, 443, 1, 0, 0, 0, 446,
		447, 5, 2, 0, 0, 447, 55, 1, 0, 0, 0, 448, 453, 3, 34, 17, 0, 449, 450,
		5, 7, 0, 0, 450, 452, 3, 34, 17, 0, 451, 449, 1, 0, 0, 0, 452, 455, 1,
		0, 0, 0, 453, 451, 1, 0, 0, 0, 453, 454, 1, 0, 0, 0, 454, 57, 1, 0, 0,
		0, 455, 453, 1, 0, 0, 0, 456, 468, 3, 62, 31, 0, 457, 468, 3, 70, 35, 0,
		458, 468, 3, 96, 48, 0, 459, 468, 3, 72, 36, 0, 460, 468, 3, 78, 39, 0,
		461, 468, 3, 90, 45, 0, 462, 468, 3, 60, 30, 0, 463, 468, 3, 6, 3, 0, 464,
		468, 3, 74, 37, 0, 465, 468, 3, 40, 20, 0, 466, 468, 3, 36, 18, 0, 467,
		456, 1, 0, 0, 0, 467, 457, 1, 0, 0, 0, 467, 458, 1, 0, 0, 0, 467, 459,
		1, 0, 0, 0, 467, 460, 1, 0, 0, 0, 467, 461, 1, 0, 0, 0, 467, 462, 1, 0,
		0, 0, 467, 463, 1, 0, 0, 0, 467, 464, 1, 0, 0, 0, 467, 465, 1, 0, 0, 0,
		467, 466, 1, 0, 0, 0, 468, 59, 1, 0, 0, 0, 469, 470, 5, 54, 0, 0, 470,
		471, 7, 8, 0, 0, 471, 61, 1, 0, 0, 0, 472, 474, 3, 18, 9, 0, 473, 472,
		1, 0, 0, 0, 473, 474, 1, 0, 0, 0, 474, 475, 1, 0, 0, 0, 475, 476, 3, 20,
		10, 0, 476, 477, 5, 52, 0, 0, 477, 478, 3, 34, 17, 0, 478, 490, 1, 0, 0,
		0, 479, 480, 3, 20, 10, 0, 480, 481, 5, 8, 0, 0, 481, 482, 3, 18, 9, 0,
		482, 483, 5, 52, 0, 0, 483, 484, 3, 34, 17, 0, 484, 490, 1, 0, 0, 0, 485,
		486, 3, 64, 32, 0, 486, 487, 5, 52, 0, 0, 487, 488, 3, 56, 28, 0, 488,
		490, 1, 0, 0, 0, 489, 473, 1, 0, 0, 0, 489, 479, 1, 0, 0, 0, 489, 485,
		1, 0, 0, 0, 490, 63, 1, 0, 0, 0, 491, 496, 3, 20, 10, 0, 492, 493, 5, 7,
		0, 0, 493, 495, 3, 20, 10, 0, 494, 492, 1, 0, 0, 0, 495, 498, 1, 0, 0,
		0, 496, 494, 1, 0, 0, 0, 496, 497, 1, 0, 0, 0, 497, 65, 1, 0, 0, 0, 498,
		496, 1, 0, 0, 0, 499, 500, 3, 18, 9, 0, 500, 505, 5, 54, 0, 0, 501, 502,
		5, 7, 0, 0, 502, 504, 5, 54, 0, 0, 503, 501, 1, 0, 0, 0, 504, 507, 1, 0,
		0, 0, 505, 503, 1, 0, 0, 0, 505, 506, 1, 0, 0, 0, 506, 519, 1, 0, 0, 0,
		507, 505, 1, 0, 0, 0, 508, 513, 5, 54, 0, 0, 509, 510, 5, 7, 0, 0, 510,
		512, 5, 54, 0, 0, 511, 509, 1, 0, 0, 0, 512, 515, 1, 0, 0, 0, 513, 511,
		1, 0, 0, 0, 513, 514, 1, 0, 0, 0, 514, 516, 1, 0, 0, 0, 515, 513, 1, 0,
		0, 0, 516, 517, 5, 8, 0, 0, 517, 519, 3, 18, 9, 0, 518, 499, 1, 0, 0, 0,
		518, 508, 1, 0, 0, 0, 519, 67, 1, 0, 0, 0, 520, 522, 3, 18, 9, 0, 521,
		520, 1, 0, 0, 0, 521, 522, 1, 0, 0, 0, 522, 523, 1, 0, 0, 0, 523, 530,
		5, 54, 0, 0, 524, 527, 5, 54, 0, 0, 525, 526, 5, 8, 0, 0, 526, 528, 3,
		18, 9, 0, 527, 525, 1, 0, 0, 0, 527, 528, 1, 0, 0, 0, 528, 530, 1, 0, 0,
		0, 529, 521, 1, 0, 0, 0, 529, 524, 1, 0, 0, 0, 530, 69, 1, 0, 0, 0, 531,
		538, 3, 66, 33, 0, 532, 533, 3, 68, 34, 0, 533, 534, 5, 52, 0, 0, 534,
		535, 3, 34, 17, 0, 535, 538, 1, 0, 0, 0, 536, 538, 3, 4, 2, 0, 537, 531,
		1, 0, 0, 0, 537, 532, 1, 0, 0, 0, 537, 536, 1, 0, 0, 0, 538, 71, 1, 0,
		0, 0, 539, 540, 5, 39, 0, 0, 540, 541, 3, 32, 16, 0, 541, 547, 3, 6, 3,
		0, 542, 545, 5, 40, 0, 0, 543, 546, 3, 6, 3, 0, 544, 546, 3, 72, 36, 0,
		545, 543, 1, 0, 0, 0, 545, 544, 1, 0, 0, 0, 546, 548, 1, 0, 0, 0, 547,
		542, 1, 0, 0, 0, 547, 548, 1, 0, 0, 0, 548, 73, 1, 0, 0, 0, 549, 550, 5,
		37, 0, 0, 550, 551, 3, 6, 3, 0, 551, 552, 3, 76, 38, 0, 552, 75, 1, 0,
		0, 0, 553, 560, 5, 38, 0, 0, 554, 556, 5, 3, 0, 0, 555, 557, 3, 18, 9,
		0, 556, 555, 1, 0, 0, 0, 556, 557, 1, 0, 0, 0, 557, 558, 1, 0, 0, 0, 558,
		559, 5, 54, 0, 0, 559, 561, 5, 4, 0, 0, 560, 554, 1, 0, 0, 0, 560, 561,
		1, 0, 0, 0, 561, 562, 1, 0, 0, 0, 562, 563, 3, 6, 3, 0, 563, 77, 1, 0,
		0, 0, 564, 568, 3, 84, 42, 0, 565, 568, 3, 80, 40, 0, 566, 568, 3, 88,
		44, 0, 567, 564, 1, 0, 0, 0, 567, 565, 1, 0, 0, 0, 567, 566, 1, 0, 0, 0,
		568, 79, 1, 0, 0, 0, 569, 575, 5, 41, 0, 0, 570, 571, 5, 3, 0, 0, 571,
		572, 3, 82, 41, 0, 572, 573, 5, 4, 0, 0, 573, 576, 1, 0, 0, 0, 574, 576,
		3, 82, 41, 0, 575, 570, 1, 0, 0, 0, 575, 574, 1, 0, 0, 0, 576, 577, 1,
		0, 0, 0, 577, 578, 3, 6, 3, 0, 578, 81, 1, 0, 0, 0, 579, 581, 3, 58, 29,
		0, 580, 579, 1, 0, 0, 0, 580, 581, 1, 0, 0, 0, 581, 582, 1, 0, 0, 0, 582,
		584, 5, 9, 0, 0, 583, 585, 3, 32, 16, 0, 584, 583, 1, 0, 0, 0, 584, 585,
		1, 0, 0, 0, 585, 586, 1, 0, 0, 0, 586, 588, 5, 9, 0, 0, 587, 589, 3, 58,
		29, 0, 588, 587, 1, 0, 0, 0, 588, 589, 1, 0, 0, 0, 589, 83, 1, 0, 0, 0,
		590, 596, 5, 41, 0, 0, 591, 592, 5, 3, 0, 0, 592, 593, 3, 86, 43, 0, 593,
		594, 5, 4, 0, 0, 594, 597, 1, 0, 0, 0, 595, 597, 3, 86, 43, 0, 596, 591,
		1, 0, 0, 0, 596, 595, 1, 0, 0, 0, 597, 598, 1, 0, 0, 0, 598, 599, 3, 6,
		3, 0, 599, 85, 1, 0, 0, 0, 600, 602, 3, 18, 9, 0, 601, 600, 1, 0, 0, 0,
		601, 602, 1, 0, 0, 0, 602, 603, 1, 0, 0, 0, 603, 604, 5, 54, 0, 0, 604,
		605, 5, 8, 0, 0, 605, 606, 3, 32, 16, 0, 606, 87, 1, 0, 0, 0, 607, 609,
		5, 41, 0, 0, 608, 610, 3, 32, 16, 0, 609, 608, 1, 0, 0, 0, 609, 610, 1,
		0, 0, 0, 610, 611, 1, 0, 0, 0, 611, 612, 3, 6, 3, 0, 612, 89, 1, 0, 0,
		0, 613, 614, 5, 42, 0, 0, 614, 615, 3, 32, 16, 0, 615, 619, 5, 1, 0, 0,
		616, 618, 3, 92, 46, 0, 617, 616, 1, 0, 0, 0, 618, 621, 1, 0, 0, 0, 619,
		617, 1, 0, 0, 0, 619, 620, 1, 0, 0, 0, 620, 622, 1, 0, 0, 0, 621, 619,
		1, 0, 0, 0, 622, 623, 5, 2, 0, 0, 623, 91, 1, 0, 0, 0, 624, 625, 5, 33,
		0, 0, 625, 630, 3, 32, 16, 0, 626, 627, 5, 7, 0, 0, 627, 629, 3, 32, 16,
		0, 628, 626, 1, 0, 0, 0, 629, 632, 1, 0, 0, 0, 630, 628, 1, 0, 0, 0, 630,
		631, 1, 0, 0, 0, 631, 633, 1, 0, 0, 0, 632, 630, 1, 0, 0, 0, 633, 634,
		3, 94, 47, 0, 634, 638, 1, 0, 0, 0, 635, 636, 5, 34, 0, 0, 636, 638, 3,
		94, 47, 0, 637, 624, 1, 0, 0, 0, 637, 635, 1, 0, 0, 0, 638, 93, 1, 0, 0,
		0, 639, 657, 3, 6, 3, 0, 640, 653, 5, 8, 0, 0, 641, 644, 3, 58, 29, 0,
		642, 644, 3, 32, 16, 0, 643, 641, 1, 0, 0, 0, 643, 642, 1, 0, 0, 0, 644,
		648, 1, 0, 0, 0, 645, 647, 3, 98, 49, 0, 646, 645, 1, 0, 0, 0, 647, 650,
		1, 0, 0, 0, 648, 646, 1, 0, 0, 0, 648, 649, 1, 0, 0, 0, 649, 652, 1, 0,
		0, 0, 650, 648, 1, 0, 0, 0, 651, 643, 1, 0, 0, 0, 652, 655, 1, 0, 0, 0,
		653, 651, 1, 0, 0, 0, 653, 654, 1, 0, 0, 0, 654, 657, 1, 0, 0, 0, 655,
		653, 1, 0, 0, 0, 656, 639, 1, 0, 0, 0, 656, 640, 1, 0, 0, 0, 657, 95, 1,
		0, 0, 0, 658, 672, 5, 44, 0, 0, 659, 672, 5, 43, 0, 0, 660, 669, 5, 32,
		0, 0, 661, 666, 3, 34, 17, 0, 662, 663, 5, 7, 0, 0, 663, 665, 3, 34, 17,
		0, 664, 662, 1, 0, 0, 0, 665, 668, 1, 0, 0, 0, 666, 664, 1, 0, 0, 0, 666,
		667, 1, 0, 0, 0, 667, 670, 1, 0, 0, 0, 668, 666, 1, 0, 0, 0, 669, 661,
		1, 0, 0, 0, 669, 670, 1, 0, 0, 0, 670, 672, 1, 0, 0, 0, 671, 658, 1, 0,
		0, 0, 671, 659, 1, 0, 0, 0, 671, 660, 1, 0, 0, 0, 672, 97, 1, 0, 0, 0,
		673, 674, 7, 9, 0, 0, 674, 99, 1, 0, 0, 0, 86, 104, 109, 114, 130, 135,
		140, 149, 159, 162, 171, 176, 179, 185, 189, 192, 196, 201, 203, 209, 214,
		217, 219, 226, 231, 240, 247, 255, 260, 262, 281, 311, 314, 321, 326, 328,
		344, 349, 354, 360, 369, 376, 382, 385, 390, 394, 412, 416, 421, 425, 431,
		438, 443, 453, 467, 473, 489, 496, 505, 513, 518, 521, 527, 529, 537, 545,
		547, 556, 560, 567, 575, 580, 584, 588, 596, 601, 609, 619, 630, 637, 643,
		648, 653, 656, 666, 669, 671,
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
	V2ParserQuestion       = 25
	V2ParserAddAdd         = 26
	V2ParserSubSub         = 27
	V2ParserTo             = 28
	V2ParserStruct         = 29
	V2ParserMap            = 30
	V2ParserFunction       = 31
	V2ParserReturn         = 32
	V2ParserCase           = 33
	V2ParserDefault        = 34
	V2ParserOpen           = 35
	V2ParserAs             = 36
	V2ParserTry            = 37
	V2ParserCatch          = 38
	V2ParserIf             = 39
	V2ParserElse           = 40
	V2ParserFor            = 41
	V2ParserMatch          = 42
	V2ParserBreak          = 43
	V2ParserContinue       = 44
	V2ParserTrue           = 45
	V2ParserFalse          = 46
	V2ParserNil            = 47
	V2ParserIntegerLiteral = 48
	V2ParserNumberLiteral  = 49
	V2ParserStringLiteral  = 50
	V2ParserNot            = 51
	V2ParserAssign         = 52
	V2ParserRef            = 53
	V2ParserIdentifier     = 54
	V2ParserComment        = 55
	V2ParserBlkComment     = 56
	V2ParserWS             = 57
	V2ParserNewLine        = 58
)

// V2Parser rules.
const (
	V2ParserRULE_program           = 0
	V2ParserRULE_structBlock       = 1
	V2ParserRULE_funcBlock         = 2
	V2ParserRULE_codeBlock         = 3
	V2ParserRULE_declareBlock      = 4
	V2ParserRULE_funcSig           = 5
	V2ParserRULE_funcReturnType    = 6
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
	V2ParserRULE_openStmt          = 20
	V2ParserRULE_literal           = 21
	V2ParserRULE_initializer       = 22
	V2ParserRULE_arrayInitializer  = 23
	V2ParserRULE_identifierPair    = 24
	V2ParserRULE_mapPair           = 25
	V2ParserRULE_structInitializer = 26
	V2ParserRULE_mapInitializer    = 27
	V2ParserRULE_commaExpr         = 28
	V2ParserRULE_stmt              = 29
	V2ParserRULE_uOperStmt         = 30
	V2ParserRULE_assignStmt        = 31
	V2ParserRULE_vars              = 32
	V2ParserRULE_typedIdentifiers  = 33
	V2ParserRULE_typedIdentifier   = 34
	V2ParserRULE_declareStmt       = 35
	V2ParserRULE_ifStmt            = 36
	V2ParserRULE_tryCatchSmt       = 37
	V2ParserRULE_catchStmt         = 38
	V2ParserRULE_loopStmt          = 39
	V2ParserRULE_cStyleFor         = 40
	V2ParserRULE_cStyleForSign     = 41
	V2ParserRULE_iterFor           = 42
	V2ParserRULE_iterForSign       = 43
	V2ParserRULE_whileStyleFor     = 44
	V2ParserRULE_matchStmt         = 45
	V2ParserRULE_matchCase         = 46
	V2ParserRULE_caseBlock         = 47
	V2ParserRULE_jumpStmt          = 48
	V2ParserRULE_sep               = 49
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStructBlock() []IStructBlockContext
	StructBlock(i int) IStructBlockContext
	AllFuncBlock() []IFuncBlockContext
	FuncBlock(i int) IFuncBlockContext
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext
	AllCommaExpr() []ICommaExprContext
	CommaExpr(i int) ICommaExprContext
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

func (s *ProgramContext) AllCommaExpr() []ICommaExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICommaExprContext); ok {
			len++
		}
	}

	tst := make([]ICommaExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICommaExprContext); ok {
			tst[i] = t.(ICommaExprContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) CommaExpr(i int) ICommaExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommaExprContext); ok {
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

	return t.(ICommaExprContext)
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
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&31523727982198826) != 0 {
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(100)
				p.StructBlock()
			}

		case 2:
			{
				p.SetState(101)
				p.FuncBlock()
			}

		case 3:
			{
				p.SetState(102)
				p.Stmt()
			}

		case 4:
			{
				p.SetState(103)
				p.CommaExpr()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == V2ParserSemi || _la == V2ParserNewLine {
			{
				p.SetState(106)
				p.Sep()
			}

			p.SetState(111)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(117)
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
	p.EnterRule(localctx, 2, V2ParserRULE_structBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(V2ParserStruct)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(120)
		p.Match(V2ParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
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
	p.EnterRule(localctx, 4, V2ParserRULE_funcBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Match(V2ParserFunction)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
		p.FuncSig()
	}
	{
		p.SetState(125)
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
	p.EnterRule(localctx, 6, V2ParserRULE_codeBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(V2ParserLBrack)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&31523727982198826) != 0 {
		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(128)
				p.Stmt()
			}

		case 2:
			{
				p.SetState(129)
				p.expr(0)
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == V2ParserSemi || _la == V2ParserNewLine {
			{
				p.SetState(132)
				p.Sep()
			}

			p.SetState(137)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(143)
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
	p.EnterRule(localctx, 8, V2ParserRULE_declareBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Match(V2ParserLBrack)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&18014401730707456) != 0 {
		{
			p.SetState(146)
			p.DeclareStmt()
		}

		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(152)
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
	FuncReturnType() IFuncReturnTypeContext

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

func (s *FuncSigContext) FuncReturnType() IFuncReturnTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncReturnTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncReturnTypeContext)
}

func (s *FuncSigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncSigContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
	p.EnterRule(localctx, 10, V2ParserRULE_funcSig)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.Match(V2ParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
		p.Match(V2ParserLParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.FuncSignArgs()
	}
	{
		p.SetState(157)
		p.Match(V2ParserRParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&18014401999143176) != 0 {
		{
			p.SetState(158)
			p.FuncReturnType()
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

// IFuncReturnTypeContext is an interface to support dynamic dispatch.
type IFuncReturnTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFlag returns the Flag token.
	GetFlag() antlr.Token

	// SetFlag sets the Flag token.
	SetFlag(antlr.Token)

	// Getter signatures
	AllType_() []ITypeContext
	Type_(i int) ITypeContext
	LParen() antlr.TerminalNode
	RParen() antlr.TerminalNode
	To() antlr.TerminalNode
	Col() antlr.TerminalNode
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsFuncReturnTypeContext differentiates from other interfaces.
	IsFuncReturnTypeContext()
}

type FuncReturnTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	Flag   antlr.Token
}

func NewEmptyFuncReturnTypeContext() *FuncReturnTypeContext {
	var p = new(FuncReturnTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_funcReturnType
	return p
}

func InitEmptyFuncReturnTypeContext(p *FuncReturnTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_funcReturnType
}

func (*FuncReturnTypeContext) IsFuncReturnTypeContext() {}

func NewFuncReturnTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncReturnTypeContext {
	var p = new(FuncReturnTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_funcReturnType

	return p
}

func (s *FuncReturnTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncReturnTypeContext) GetFlag() antlr.Token { return s.Flag }

func (s *FuncReturnTypeContext) SetFlag(v antlr.Token) { s.Flag = v }

func (s *FuncReturnTypeContext) AllType_() []ITypeContext {
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

func (s *FuncReturnTypeContext) Type_(i int) ITypeContext {
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

func (s *FuncReturnTypeContext) LParen() antlr.TerminalNode {
	return s.GetToken(V2ParserLParen, 0)
}

func (s *FuncReturnTypeContext) RParen() antlr.TerminalNode {
	return s.GetToken(V2ParserRParen, 0)
}

func (s *FuncReturnTypeContext) To() antlr.TerminalNode {
	return s.GetToken(V2ParserTo, 0)
}

func (s *FuncReturnTypeContext) Col() antlr.TerminalNode {
	return s.GetToken(V2ParserCol, 0)
}

func (s *FuncReturnTypeContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(V2ParserComma)
}

func (s *FuncReturnTypeContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(V2ParserComma, i)
}

func (s *FuncReturnTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncReturnTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncReturnTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitFuncReturnType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) FuncReturnType() (localctx IFuncReturnTypeContext) {
	localctx = NewFuncReturnTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, V2ParserRULE_funcReturnType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == V2ParserCol || _la == V2ParserTo {
		{
			p.SetState(161)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*FuncReturnTypeContext).Flag = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == V2ParserCol || _la == V2ParserTo) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*FuncReturnTypeContext).Flag = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case V2ParserMap, V2ParserFunction, V2ParserIdentifier:
		{
			p.SetState(164)
			p.Type_()
		}

	case V2ParserLParen:
		{
			p.SetState(165)
			p.Match(V2ParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)
			p.Type_()
		}
		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == V2ParserComma {
			{
				p.SetState(167)
				p.Match(V2ParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(168)
				p.Type_()
			}

			p.SetState(173)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(174)
			p.Match(V2ParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
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
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&27021600985448448) != 0 {
		{
			p.SetState(178)
			p.FuncSignArgItem()
		}

	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == V2ParserComma {
		{
			p.SetState(181)
			p.Match(V2ParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.FuncSignArgItem()
		}

		p.SetState(187)
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
	Ref() antlr.TerminalNode
	Col() antlr.TerminalNode

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

func (s *FuncSignArgItemContext) Ref() antlr.TerminalNode {
	return s.GetToken(V2ParserRef, 0)
}

func (s *FuncSignArgItemContext) Col() antlr.TerminalNode {
	return s.GetToken(V2ParserCol, 0)
}

func (s *FuncSignArgItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncSignArgItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
	var _la int

	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(189)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(188)
				p.Type_()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == V2ParserRef {
			{
				p.SetState(191)
				p.Match(V2ParserRef)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(194)
			p.Match(V2ParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == V2ParserRef {
			{
				p.SetState(195)
				p.Match(V2ParserRef)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(198)
			p.Match(V2ParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == V2ParserCol {
			{
				p.SetState(199)
				p.Match(V2ParserCol)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(200)
				p.Type_()
			}

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
	Question() antlr.TerminalNode

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

func (s *TypeContext) Question() antlr.TerminalNode {
	return s.GetToken(V2ParserQuestion, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
	var _la int

	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case V2ParserMap:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(205)
			p.Match(V2ParserMap)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case V2ParserFunction:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(206)
			p.Match(V2ParserFunction)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case V2ParserIdentifier:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(209)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(207)

				var _m = p.Match(V2ParserIdentifier)

				localctx.(*TypeContext).PackageName = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(208)
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
			p.SetState(211)

			var _m = p.Match(V2ParserIdentifier)

			localctx.(*TypeContext).TypeName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(214)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(212)
				p.Match(V2ParserLSquare)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(213)
				p.Match(V2ParserRSquare)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(217)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == V2ParserQuestion {
			{
				p.SetState(216)
				p.Match(V2ParserQuestion)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
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
		p.SetState(221)
		p.BaseVar()
	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == V2ParserDot {
		{
			p.SetState(222)
			p.Match(V2ParserDot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)
			p.BaseVar()
		}

		p.SetState(228)
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
		p.SetState(229)
		p.Match(V2ParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == V2ParserLSquare {
		{
			p.SetState(230)
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
		p.SetState(233)
		p.Match(V2ParserLSquare)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(234)
		p.expr(0)
	}
	{
		p.SetState(235)
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
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(237)
				p.Index()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
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
	FuncReturnType() IFuncReturnTypeContext
	To() antlr.TerminalNode
	ExprWithLambda() IExprWithLambdaContext

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

func (s *LambdaContext) FuncReturnType() IFuncReturnTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncReturnTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncReturnTypeContext)
}

func (s *LambdaContext) To() antlr.TerminalNode {
	return s.GetToken(V2ParserTo, 0)
}

func (s *LambdaContext) ExprWithLambda() IExprWithLambdaContext {
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

func (s *LambdaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case V2ParserFunction:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(242)
			p.Match(V2ParserFunction)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)
			p.Match(V2ParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(244)
			p.FuncSignArgs()
		}
		{
			p.SetState(245)
			p.Match(V2ParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&18014401999143176) != 0 {
			{
				p.SetState(246)
				p.FuncReturnType()
			}

		}
		{
			p.SetState(249)
			p.CodeBlock()
		}

	case V2ParserLParen:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(251)
			p.Match(V2ParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(252)
			p.FuncSignArgs()
		}
		{
			p.SetState(253)
			p.Match(V2ParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(255)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(254)
				p.FuncReturnType()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(257)
			p.Match(V2ParserTo)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(260)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(258)
				p.CodeBlock()
			}

		case 2:
			{
				p.SetState(259)
				p.exprWithLambda(0)
			}

		case antlr.ATNInvalidAltNumber:
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
		p.SetState(264)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2251799819976704) != 0) {
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

	// GetCallExpr returns the CallExpr rule contexts.
	GetCallExpr() IExprContext

	// GetTriCond returns the TriCond rule contexts.
	GetTriCond() IExprContext

	// GetRHS returns the RHS rule contexts.
	GetRHS() IExprContext

	// GetIfTrue returns the IfTrue rule contexts.
	GetIfTrue() IExprContext

	// GetIfFalse returns the IfFalse rule contexts.
	GetIfFalse() IExprContext

	// GetIndex returns the Index rule contexts.
	GetIndex() IExprContext

	// GetEndIndex returns the EndIndex rule contexts.
	GetEndIndex() IExprContext

	// SetLHS sets the LHS rule contexts.
	SetLHS(IExprContext)

	// SetCallExpr sets the CallExpr rule contexts.
	SetCallExpr(IExprContext)

	// SetTriCond sets the TriCond rule contexts.
	SetTriCond(IExprContext)

	// SetRHS sets the RHS rule contexts.
	SetRHS(IExprContext)

	// SetIfTrue sets the IfTrue rule contexts.
	SetIfTrue(IExprContext)

	// SetIfFalse sets the IfFalse rule contexts.
	SetIfFalse(IExprContext)

	// SetIndex sets the Index rule contexts.
	SetIndex(IExprContext)

	// SetEndIndex sets the EndIndex rule contexts.
	SetEndIndex(IExprContext)

	// Getter signatures
	Ref() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	FuncCall() IFuncCallContext
	UnaryOper() IUnaryOperContext
	LParen() antlr.TerminalNode
	RParen() antlr.TerminalNode
	StructInitializer() IStructInitializerContext
	Identifier() antlr.TerminalNode
	Literal() ILiteralContext
	Initializer() IInitializerContext
	Dot() antlr.TerminalNode
	Pow() antlr.TerminalNode
	Mul() antlr.TerminalNode
	Div() antlr.TerminalNode
	Mod() antlr.TerminalNode
	Add() antlr.TerminalNode
	Sub() antlr.TerminalNode
	Equals() antlr.TerminalNode
	NotEq() antlr.TerminalNode
	Greater() antlr.TerminalNode
	Less() antlr.TerminalNode
	GreaterEq() antlr.TerminalNode
	LessEq() antlr.TerminalNode
	Or() antlr.TerminalNode
	And() antlr.TerminalNode
	Question() antlr.TerminalNode
	Col() antlr.TerminalNode
	LSquare() antlr.TerminalNode
	RSquare() antlr.TerminalNode
	FuncCallArgs() IFuncCallArgsContext
	Indexes() IIndexesContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	LHS      IExprContext
	CallExpr IExprContext
	TriCond  IExprContext
	RHS      IExprContext
	OP       antlr.Token
	IfTrue   IExprContext
	IfFalse  IExprContext
	Index    IExprContext
	EndIndex IExprContext
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

func (s *ExprContext) GetCallExpr() IExprContext { return s.CallExpr }

func (s *ExprContext) GetTriCond() IExprContext { return s.TriCond }

func (s *ExprContext) GetRHS() IExprContext { return s.RHS }

func (s *ExprContext) GetIfTrue() IExprContext { return s.IfTrue }

func (s *ExprContext) GetIfFalse() IExprContext { return s.IfFalse }

func (s *ExprContext) GetIndex() IExprContext { return s.Index }

func (s *ExprContext) GetEndIndex() IExprContext { return s.EndIndex }

func (s *ExprContext) SetLHS(v IExprContext) { s.LHS = v }

func (s *ExprContext) SetCallExpr(v IExprContext) { s.CallExpr = v }

func (s *ExprContext) SetTriCond(v IExprContext) { s.TriCond = v }

func (s *ExprContext) SetRHS(v IExprContext) { s.RHS = v }

func (s *ExprContext) SetIfTrue(v IExprContext) { s.IfTrue = v }

func (s *ExprContext) SetIfFalse(v IExprContext) { s.IfFalse = v }

func (s *ExprContext) SetIndex(v IExprContext) { s.Index = v }

func (s *ExprContext) SetEndIndex(v IExprContext) { s.EndIndex = v }

func (s *ExprContext) Ref() antlr.TerminalNode {
	return s.GetToken(V2ParserRef, 0)
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

func (s *ExprContext) LParen() antlr.TerminalNode {
	return s.GetToken(V2ParserLParen, 0)
}

func (s *ExprContext) RParen() antlr.TerminalNode {
	return s.GetToken(V2ParserRParen, 0)
}

func (s *ExprContext) StructInitializer() IStructInitializerContext {
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

func (s *ExprContext) Dot() antlr.TerminalNode {
	return s.GetToken(V2ParserDot, 0)
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

func (s *ExprContext) Question() antlr.TerminalNode {
	return s.GetToken(V2ParserQuestion, 0)
}

func (s *ExprContext) Col() antlr.TerminalNode {
	return s.GetToken(V2ParserCol, 0)
}

func (s *ExprContext) LSquare() antlr.TerminalNode {
	return s.GetToken(V2ParserLSquare, 0)
}

func (s *ExprContext) RSquare() antlr.TerminalNode {
	return s.GetToken(V2ParserRSquare, 0)
}

func (s *ExprContext) FuncCallArgs() IFuncCallArgsContext {
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
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(267)
			p.Match(V2ParserRef)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(268)
			p.expr(18)
		}

	case 2:
		{
			p.SetState(269)
			p.FuncCall()
		}

	case 3:
		{
			p.SetState(270)
			p.UnaryOper()
		}
		{
			p.SetState(271)
			p.expr(13)
		}

	case 4:
		{
			p.SetState(273)
			p.Match(V2ParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(274)
			p.expr(0)
		}
		{
			p.SetState(275)
			p.Match(V2ParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		{
			p.SetState(277)
			p.StructInitializer()
		}

	case 6:
		{
			p.SetState(278)
			p.Match(V2ParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		{
			p.SetState(279)
			p.Literal()
		}

	case 8:
		{
			p.SetState(280)
			p.Initializer()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(328)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(326)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).LHS = _prevctx
				p.PushNewRecursionContext(localctx, _startState, V2ParserRULE_expr)
				p.SetState(283)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(284)
					p.Match(V2ParserDot)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(285)

					var _x = p.expr(18)

					localctx.(*ExprContext).RHS = _x
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).LHS = _prevctx
				p.PushNewRecursionContext(localctx, _startState, V2ParserRULE_expr)
				p.SetState(286)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(287)

					var _m = p.Match(V2ParserPow)

					localctx.(*ExprContext).OP = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(288)

					var _x = p.expr(12)

					localctx.(*ExprContext).RHS = _x
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).LHS = _prevctx
				p.PushNewRecursionContext(localctx, _startState, V2ParserRULE_expr)
				p.SetState(289)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(290)

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
					p.SetState(291)

					var _x = p.expr(11)

					localctx.(*ExprContext).RHS = _x
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).LHS = _prevctx
				p.PushNewRecursionContext(localctx, _startState, V2ParserRULE_expr)
				p.SetState(292)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(293)

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
					p.SetState(294)

					var _x = p.expr(10)

					localctx.(*ExprContext).RHS = _x
				}

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).LHS = _prevctx
				p.PushNewRecursionContext(localctx, _startState, V2ParserRULE_expr)
				p.SetState(295)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(296)

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
					p.SetState(297)

					var _x = p.expr(9)

					localctx.(*ExprContext).RHS = _x
				}

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).LHS = _prevctx
				p.PushNewRecursionContext(localctx, _startState, V2ParserRULE_expr)
				p.SetState(298)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(299)

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
					p.SetState(300)

					var _x = p.expr(8)

					localctx.(*ExprContext).RHS = _x
				}

			case 7:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).TriCond = _prevctx
				p.PushNewRecursionContext(localctx, _startState, V2ParserRULE_expr)
				p.SetState(301)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(302)
					p.Match(V2ParserQuestion)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(303)

					var _x = p.expr(0)

					localctx.(*ExprContext).IfTrue = _x
				}
				{
					p.SetState(304)
					p.Match(V2ParserCol)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(305)

					var _x = p.expr(2)

					localctx.(*ExprContext).IfFalse = _x
				}

			case 8:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).LHS = _prevctx
				p.PushNewRecursionContext(localctx, _startState, V2ParserRULE_expr)
				p.SetState(307)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(308)
					p.Match(V2ParserLSquare)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(309)

					var _x = p.expr(0)

					localctx.(*ExprContext).Index = _x
				}
				p.SetState(311)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == V2ParserCol {
					{
						p.SetState(310)
						p.Match(V2ParserCol)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				p.SetState(314)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&31490016783892522) != 0 {
					{
						p.SetState(313)

						var _x = p.expr(0)

						localctx.(*ExprContext).EndIndex = _x
					}

				}
				{
					p.SetState(316)
					p.Match(V2ParserRSquare)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 9:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).CallExpr = _prevctx
				p.PushNewRecursionContext(localctx, _startState, V2ParserRULE_expr)
				p.SetState(318)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(319)
					p.Match(V2ParserLParen)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(321)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&31490016783892522) != 0 {
					{
						p.SetState(320)
						p.FuncCallArgs()
					}

				}
				{
					p.SetState(323)
					p.Match(V2ParserRParen)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 10:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, V2ParserRULE_expr)
				p.SetState(324)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(325)
					p.Indexes()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(330)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
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

	// GetCallExpr returns the CallExpr rule contexts.
	GetCallExpr() IExprWithLambdaContext

	// GetTriCond returns the TriCond rule contexts.
	GetTriCond() IExprContext

	// GetIfTrue returns the IfTrue rule contexts.
	GetIfTrue() IExprWithLambdaContext

	// GetIfFalse returns the IfFalse rule contexts.
	GetIfFalse() IExprWithLambdaContext

	// GetParenExpr returns the ParenExpr rule contexts.
	GetParenExpr() IExprWithLambdaContext

	// SetCallExpr sets the CallExpr rule contexts.
	SetCallExpr(IExprWithLambdaContext)

	// SetTriCond sets the TriCond rule contexts.
	SetTriCond(IExprContext)

	// SetIfTrue sets the IfTrue rule contexts.
	SetIfTrue(IExprWithLambdaContext)

	// SetIfFalse sets the IfFalse rule contexts.
	SetIfFalse(IExprWithLambdaContext)

	// SetParenExpr sets the ParenExpr rule contexts.
	SetParenExpr(IExprWithLambdaContext)

	// Getter signatures
	Lambda() ILambdaContext
	Expr() IExprContext
	Question() antlr.TerminalNode
	Col() antlr.TerminalNode
	AllExprWithLambda() []IExprWithLambdaContext
	ExprWithLambda(i int) IExprWithLambdaContext
	LParen() antlr.TerminalNode
	RParen() antlr.TerminalNode
	FuncCallArgs() IFuncCallArgsContext

	// IsExprWithLambdaContext differentiates from other interfaces.
	IsExprWithLambdaContext()
}

type ExprWithLambdaContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	CallExpr  IExprWithLambdaContext
	TriCond   IExprContext
	IfTrue    IExprWithLambdaContext
	IfFalse   IExprWithLambdaContext
	ParenExpr IExprWithLambdaContext
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

func (s *ExprWithLambdaContext) GetCallExpr() IExprWithLambdaContext { return s.CallExpr }

func (s *ExprWithLambdaContext) GetTriCond() IExprContext { return s.TriCond }

func (s *ExprWithLambdaContext) GetIfTrue() IExprWithLambdaContext { return s.IfTrue }

func (s *ExprWithLambdaContext) GetIfFalse() IExprWithLambdaContext { return s.IfFalse }

func (s *ExprWithLambdaContext) GetParenExpr() IExprWithLambdaContext { return s.ParenExpr }

func (s *ExprWithLambdaContext) SetCallExpr(v IExprWithLambdaContext) { s.CallExpr = v }

func (s *ExprWithLambdaContext) SetTriCond(v IExprContext) { s.TriCond = v }

func (s *ExprWithLambdaContext) SetIfTrue(v IExprWithLambdaContext) { s.IfTrue = v }

func (s *ExprWithLambdaContext) SetIfFalse(v IExprWithLambdaContext) { s.IfFalse = v }

func (s *ExprWithLambdaContext) SetParenExpr(v IExprWithLambdaContext) { s.ParenExpr = v }

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

func (s *ExprWithLambdaContext) Question() antlr.TerminalNode {
	return s.GetToken(V2ParserQuestion, 0)
}

func (s *ExprWithLambdaContext) Col() antlr.TerminalNode {
	return s.GetToken(V2ParserCol, 0)
}

func (s *ExprWithLambdaContext) AllExprWithLambda() []IExprWithLambdaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprWithLambdaContext); ok {
			len++
		}
	}

	tst := make([]IExprWithLambdaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprWithLambdaContext); ok {
			tst[i] = t.(IExprWithLambdaContext)
			i++
		}
	}

	return tst
}

func (s *ExprWithLambdaContext) ExprWithLambda(i int) IExprWithLambdaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprWithLambdaContext); ok {
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

	return t.(IExprWithLambdaContext)
}

func (s *ExprWithLambdaContext) LParen() antlr.TerminalNode {
	return s.GetToken(V2ParserLParen, 0)
}

func (s *ExprWithLambdaContext) RParen() antlr.TerminalNode {
	return s.GetToken(V2ParserRParen, 0)
}

func (s *ExprWithLambdaContext) FuncCallArgs() IFuncCallArgsContext {
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

func (s *ExprWithLambdaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprWithLambdaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
	return p.exprWithLambda(0)
}

func (p *V2Parser) exprWithLambda(_p int) (localctx IExprWithLambdaContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprWithLambdaContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprWithLambdaContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, V2ParserRULE_exprWithLambda, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(332)
			p.Lambda()
		}

	case 2:
		{
			p.SetState(333)
			p.expr(0)
		}

	case 3:
		{
			p.SetState(334)

			var _x = p.expr(0)

			localctx.(*ExprWithLambdaContext).TriCond = _x
		}
		{
			p.SetState(335)
			p.Match(V2ParserQuestion)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(336)

			var _x = p.exprWithLambda(0)

			localctx.(*ExprWithLambdaContext).IfTrue = _x
		}
		{
			p.SetState(337)
			p.Match(V2ParserCol)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(338)

			var _x = p.exprWithLambda(3)

			localctx.(*ExprWithLambdaContext).IfFalse = _x
		}

	case 4:
		{
			p.SetState(340)
			p.Match(V2ParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(341)

			var _x = p.exprWithLambda(0)

			localctx.(*ExprWithLambdaContext).ParenExpr = _x
		}
		{
			p.SetState(342)
			p.Match(V2ParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExprWithLambdaContext(p, _parentctx, _parentState)
			localctx.(*ExprWithLambdaContext).CallExpr = _prevctx
			p.PushNewRecursionContext(localctx, _startState, V2ParserRULE_exprWithLambda)
			p.SetState(346)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(347)
				p.Match(V2ParserLParen)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(349)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&31490016783892522) != 0 {
				{
					p.SetState(348)
					p.FuncCallArgs()
				}

			}
			{
				p.SetState(351)
				p.Match(V2ParserRParen)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(356)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
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
		p.SetState(357)
		p.Match(V2ParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(358)
		p.Match(V2ParserLParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&31490016783892522) != 0 {
		{
			p.SetState(359)
			p.FuncCallArgs()
		}

	}
	{
		p.SetState(362)
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
	AllExprWithLambda() []IExprWithLambdaContext
	ExprWithLambda(i int) IExprWithLambdaContext
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

func (s *FuncCallArgsContext) AllExprWithLambda() []IExprWithLambdaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprWithLambdaContext); ok {
			len++
		}
	}

	tst := make([]IExprWithLambdaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprWithLambdaContext); ok {
			tst[i] = t.(IExprWithLambdaContext)
			i++
		}
	}

	return tst
}

func (s *FuncCallArgsContext) ExprWithLambda(i int) IExprWithLambdaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprWithLambdaContext); ok {
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

	return t.(IExprWithLambdaContext)
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
		p.SetState(364)
		p.exprWithLambda(0)
	}
	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == V2ParserComma {
		{
			p.SetState(365)
			p.Match(V2ParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(366)
			p.exprWithLambda(0)
		}

		p.SetState(371)
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

	// GetAS returns the AS token.
	GetAS() antlr.Token

	// SetAS sets the AS token.
	SetAS(antlr.Token)

	// Getter signatures
	Open() antlr.TerminalNode
	StringLiteral() antlr.TerminalNode
	As() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	Dot() antlr.TerminalNode

	// IsOpenStmtContext differentiates from other interfaces.
	IsOpenStmtContext()
}

type OpenStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	AS     antlr.Token
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

func (s *OpenStmtContext) GetAS() antlr.Token { return s.AS }

func (s *OpenStmtContext) SetAS(v antlr.Token) { s.AS = v }

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

func (s *OpenStmtContext) Dot() antlr.TerminalNode {
	return s.GetToken(V2ParserDot, 0)
}

func (s *OpenStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpenStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
	p.EnterRule(localctx, 40, V2ParserRULE_openStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(372)
		p.Match(V2ParserOpen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(373)
		p.Match(V2ParserStringLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == V2ParserAs {
		{
			p.SetState(374)
			p.Match(V2ParserAs)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(375)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*OpenStmtContext).AS = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == V2ParserDot || _la == V2ParserIdentifier) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*OpenStmtContext).AS = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
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
	Nil() antlr.TerminalNode
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

func (s *LiteralContext) Nil() antlr.TerminalNode {
	return s.GetToken(V2ParserNil, 0)
}

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
	p.EnterRule(localctx, 42, V2ParserRULE_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(378)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2216615441596416) != 0) {
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
	p.EnterRule(localctx, 44, V2ParserRULE_initializer)
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(380)
			p.ArrayInitializer()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(381)
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
	p.EnterRule(localctx, 46, V2ParserRULE_arrayInitializer)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&18014401730707456) != 0 {
		{
			p.SetState(384)
			p.Type_()
		}

	}
	{
		p.SetState(387)
		p.Match(V2ParserLSquare)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&31490016783892522) != 0 {
		{
			p.SetState(388)
			p.expr(0)
		}
		p.SetState(390)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == V2ParserComma {
			{
				p.SetState(389)
				p.Match(V2ParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

		p.SetState(396)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(397)
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
	p.EnterRule(localctx, 48, V2ParserRULE_identifierPair)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(399)

		var _m = p.Match(V2ParserIdentifier)

		localctx.(*IdentifierPairContext).LHS = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(400)
		_la = p.GetTokenStream().LA(1)

		if !(_la == V2ParserCol || _la == V2ParserTo) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(401)

		var _x = p.exprWithLambda(0)

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
	p.EnterRule(localctx, 50, V2ParserRULE_mapPair)
	var _la int

	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(403)

			var _x = p.expr(0)

			localctx.(*MapPairContext).LHS = _x
		}
		{
			p.SetState(404)
			_la = p.GetTokenStream().LA(1)

			if !(_la == V2ParserCol || _la == V2ParserTo) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(405)

			var _x = p.exprWithLambda(0)

			localctx.(*MapPairContext).RHS = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(407)
			p.Match(V2ParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(408)
			p.MapPair()
		}
		{
			p.SetState(409)
			p.Match(V2ParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(411)
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
	p.EnterRule(localctx, 52, V2ParserRULE_structInitializer)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(416)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case V2ParserMap, V2ParserFunction, V2ParserIdentifier:
		{
			p.SetState(414)
			p.Type_()
		}

	case V2ParserStruct:
		{
			p.SetState(415)
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
		p.SetState(418)
		p.Match(V2ParserLBrack)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(425)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == V2ParserIdentifier {
		{
			p.SetState(419)
			p.IdentifierPair()
		}
		p.SetState(421)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == V2ParserComma {
			{
				p.SetState(420)
				p.Match(V2ParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

		p.SetState(427)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(428)
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
	p.SetState(431)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == V2ParserMap {
		{
			p.SetState(430)
			p.Match(V2ParserMap)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(433)
		p.Match(V2ParserLBrack)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(443)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&31490016783892522) != 0 {
		{
			p.SetState(434)
			p.MapPair()
		}

		p.SetState(438)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == V2ParserComma {
			{
				p.SetState(435)
				p.Match(V2ParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(440)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(445)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(446)
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

// ICommaExprContext is an interface to support dynamic dispatch.
type ICommaExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExprWithLambda() []IExprWithLambdaContext
	ExprWithLambda(i int) IExprWithLambdaContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsCommaExprContext differentiates from other interfaces.
	IsCommaExprContext()
}

type CommaExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommaExprContext() *CommaExprContext {
	var p = new(CommaExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_commaExpr
	return p
}

func InitEmptyCommaExprContext(p *CommaExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_commaExpr
}

func (*CommaExprContext) IsCommaExprContext() {}

func NewCommaExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommaExprContext {
	var p = new(CommaExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_commaExpr

	return p
}

func (s *CommaExprContext) GetParser() antlr.Parser { return s.parser }

func (s *CommaExprContext) AllExprWithLambda() []IExprWithLambdaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprWithLambdaContext); ok {
			len++
		}
	}

	tst := make([]IExprWithLambdaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprWithLambdaContext); ok {
			tst[i] = t.(IExprWithLambdaContext)
			i++
		}
	}

	return tst
}

func (s *CommaExprContext) ExprWithLambda(i int) IExprWithLambdaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprWithLambdaContext); ok {
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

	return t.(IExprWithLambdaContext)
}

func (s *CommaExprContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(V2ParserComma)
}

func (s *CommaExprContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(V2ParserComma, i)
}

func (s *CommaExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommaExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommaExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitCommaExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) CommaExpr() (localctx ICommaExprContext) {
	localctx = NewCommaExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, V2ParserRULE_commaExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(448)
		p.exprWithLambda(0)
	}
	p.SetState(453)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == V2ParserComma {
		{
			p.SetState(449)
			p.Match(V2ParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(450)
			p.exprWithLambda(0)
		}

		p.SetState(455)
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
	UOperStmt() IUOperStmtContext
	CodeBlock() ICodeBlockContext
	TryCatchSmt() ITryCatchSmtContext
	OpenStmt() IOpenStmtContext
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

func (s *StmtContext) UOperStmt() IUOperStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUOperStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUOperStmtContext)
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

func (s *StmtContext) TryCatchSmt() ITryCatchSmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITryCatchSmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITryCatchSmtContext)
}

func (s *StmtContext) OpenStmt() IOpenStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpenStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpenStmtContext)
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
	p.SetState(467)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(456)
			p.AssignStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(457)
			p.DeclareStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(458)
			p.JumpStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(459)
			p.IfStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(460)
			p.LoopStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(461)
			p.MatchStmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(462)
			p.UOperStmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(463)
			p.CodeBlock()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(464)
			p.TryCatchSmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(465)
			p.OpenStmt()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(466)
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

// IUOperStmtContext is an interface to support dynamic dispatch.
type IUOperStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	AddAdd() antlr.TerminalNode
	SubSub() antlr.TerminalNode

	// IsUOperStmtContext differentiates from other interfaces.
	IsUOperStmtContext()
}

type UOperStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUOperStmtContext() *UOperStmtContext {
	var p = new(UOperStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_uOperStmt
	return p
}

func InitEmptyUOperStmtContext(p *UOperStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_uOperStmt
}

func (*UOperStmtContext) IsUOperStmtContext() {}

func NewUOperStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UOperStmtContext {
	var p = new(UOperStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_uOperStmt

	return p
}

func (s *UOperStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *UOperStmtContext) Identifier() antlr.TerminalNode {
	return s.GetToken(V2ParserIdentifier, 0)
}

func (s *UOperStmtContext) AddAdd() antlr.TerminalNode {
	return s.GetToken(V2ParserAddAdd, 0)
}

func (s *UOperStmtContext) SubSub() antlr.TerminalNode {
	return s.GetToken(V2ParserSubSub, 0)
}

func (s *UOperStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UOperStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UOperStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitUOperStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) UOperStmt() (localctx IUOperStmtContext) {
	localctx = NewUOperStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, V2ParserRULE_uOperStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(469)
		p.Match(V2ParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(470)
		_la = p.GetTokenStream().LA(1)

		if !(_la == V2ParserAddAdd || _la == V2ParserSubSub) {
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
	Col() antlr.TerminalNode
	Vars() IVarsContext
	CommaExpr() ICommaExprContext

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

func (s *AssignStmtContext) Col() antlr.TerminalNode {
	return s.GetToken(V2ParserCol, 0)
}

func (s *AssignStmtContext) Vars() IVarsContext {
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

func (s *AssignStmtContext) CommaExpr() ICommaExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommaExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommaExprContext)
}

func (s *AssignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
	p.EnterRule(localctx, 62, V2ParserRULE_assignStmt)
	p.SetState(489)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 55, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(473)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 54, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(472)
				p.Type_()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(475)
			p.Var_()
		}
		{
			p.SetState(476)
			p.Match(V2ParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(477)
			p.exprWithLambda(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(479)
			p.Var_()
		}
		{
			p.SetState(480)
			p.Match(V2ParserCol)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(481)
			p.Type_()
		}
		{
			p.SetState(482)
			p.Match(V2ParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(483)
			p.exprWithLambda(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(485)
			p.Vars()
		}
		{
			p.SetState(486)
			p.Match(V2ParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(487)
			p.CommaExpr()
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

// IVarsContext is an interface to support dynamic dispatch.
type IVarsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVar_() []IVarContext
	Var_(i int) IVarContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

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

func (s *VarsContext) AllVar_() []IVarContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarContext); ok {
			len++
		}
	}

	tst := make([]IVarContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarContext); ok {
			tst[i] = t.(IVarContext)
			i++
		}
	}

	return tst
}

func (s *VarsContext) Var_(i int) IVarContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarContext); ok {
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

	return t.(IVarContext)
}

func (s *VarsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(V2ParserComma)
}

func (s *VarsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(V2ParserComma, i)
}

func (s *VarsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
	p.EnterRule(localctx, 64, V2ParserRULE_vars)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(491)
		p.Var_()
	}
	p.SetState(496)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == V2ParserComma {
		{
			p.SetState(492)
			p.Match(V2ParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(493)
			p.Var_()
		}

		p.SetState(498)
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

// ITypedIdentifiersContext is an interface to support dynamic dispatch.
type ITypedIdentifiersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode
	Col() antlr.TerminalNode

	// IsTypedIdentifiersContext differentiates from other interfaces.
	IsTypedIdentifiersContext()
}

type TypedIdentifiersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypedIdentifiersContext() *TypedIdentifiersContext {
	var p = new(TypedIdentifiersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_typedIdentifiers
	return p
}

func InitEmptyTypedIdentifiersContext(p *TypedIdentifiersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_typedIdentifiers
}

func (*TypedIdentifiersContext) IsTypedIdentifiersContext() {}

func NewTypedIdentifiersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypedIdentifiersContext {
	var p = new(TypedIdentifiersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_typedIdentifiers

	return p
}

func (s *TypedIdentifiersContext) GetParser() antlr.Parser { return s.parser }

func (s *TypedIdentifiersContext) Type_() ITypeContext {
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

func (s *TypedIdentifiersContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(V2ParserIdentifier)
}

func (s *TypedIdentifiersContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(V2ParserIdentifier, i)
}

func (s *TypedIdentifiersContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(V2ParserComma)
}

func (s *TypedIdentifiersContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(V2ParserComma, i)
}

func (s *TypedIdentifiersContext) Col() antlr.TerminalNode {
	return s.GetToken(V2ParserCol, 0)
}

func (s *TypedIdentifiersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedIdentifiersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypedIdentifiersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitTypedIdentifiers(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) TypedIdentifiers() (localctx ITypedIdentifiersContext) {
	localctx = NewTypedIdentifiersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, V2ParserRULE_typedIdentifiers)
	var _la int

	p.SetState(518)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(499)
			p.Type_()
		}
		{
			p.SetState(500)
			p.Match(V2ParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(505)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == V2ParserComma {
			{
				p.SetState(501)
				p.Match(V2ParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(502)
				p.Match(V2ParserIdentifier)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(507)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(508)
			p.Match(V2ParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(513)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == V2ParserComma {
			{
				p.SetState(509)
				p.Match(V2ParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(510)
				p.Match(V2ParserIdentifier)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(515)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(516)
			p.Match(V2ParserCol)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(517)
			p.Type_()
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

// ITypedIdentifierContext is an interface to support dynamic dispatch.
type ITypedIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Type_() ITypeContext
	Col() antlr.TerminalNode

	// IsTypedIdentifierContext differentiates from other interfaces.
	IsTypedIdentifierContext()
}

type TypedIdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypedIdentifierContext() *TypedIdentifierContext {
	var p = new(TypedIdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_typedIdentifier
	return p
}

func InitEmptyTypedIdentifierContext(p *TypedIdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_typedIdentifier
}

func (*TypedIdentifierContext) IsTypedIdentifierContext() {}

func NewTypedIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypedIdentifierContext {
	var p = new(TypedIdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_typedIdentifier

	return p
}

func (s *TypedIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *TypedIdentifierContext) Identifier() antlr.TerminalNode {
	return s.GetToken(V2ParserIdentifier, 0)
}

func (s *TypedIdentifierContext) Type_() ITypeContext {
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

func (s *TypedIdentifierContext) Col() antlr.TerminalNode {
	return s.GetToken(V2ParserCol, 0)
}

func (s *TypedIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypedIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitTypedIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) TypedIdentifier() (localctx ITypedIdentifierContext) {
	localctx = NewTypedIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, V2ParserRULE_typedIdentifier)
	var _la int

	p.SetState(529)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 62, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(521)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 60, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(520)
				p.Type_()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(523)
			p.Match(V2ParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(524)
			p.Match(V2ParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(527)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == V2ParserCol {
			{
				p.SetState(525)
				p.Match(V2ParserCol)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(526)
				p.Type_()
			}

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
	TypedIdentifiers() ITypedIdentifiersContext
	TypedIdentifier() ITypedIdentifierContext
	Assign() antlr.TerminalNode
	ExprWithLambda() IExprWithLambdaContext
	FuncBlock() IFuncBlockContext

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

func (s *DeclareStmtContext) TypedIdentifiers() ITypedIdentifiersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypedIdentifiersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypedIdentifiersContext)
}

func (s *DeclareStmtContext) TypedIdentifier() ITypedIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypedIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypedIdentifierContext)
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

func (s *DeclareStmtContext) FuncBlock() IFuncBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncBlockContext)
}

func (s *DeclareStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclareStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
	p.EnterRule(localctx, 70, V2ParserRULE_declareStmt)
	p.SetState(537)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 63, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(531)
			p.TypedIdentifiers()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(532)
			p.TypedIdentifier()
		}
		{
			p.SetState(533)
			p.Match(V2ParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(534)
			p.exprWithLambda(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(536)
			p.FuncBlock()
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
	Expr() IExprContext
	AllCodeBlock() []ICodeBlockContext
	CodeBlock(i int) ICodeBlockContext
	Else() antlr.TerminalNode
	IfStmt() IIfStmtContext

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

func (s *IfStmtContext) IfStmt() IIfStmtContext {
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

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
	p.EnterRule(localctx, 72, V2ParserRULE_ifStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(539)
		p.Match(V2ParserIf)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(540)
		p.expr(0)
	}
	{
		p.SetState(541)
		p.CodeBlock()
	}
	p.SetState(547)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == V2ParserElse {
		{
			p.SetState(542)
			p.Match(V2ParserElse)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(545)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case V2ParserLBrack:
			{
				p.SetState(543)
				p.CodeBlock()
			}

		case V2ParserIf:
			{
				p.SetState(544)
				p.IfStmt()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// ITryCatchSmtContext is an interface to support dynamic dispatch.
type ITryCatchSmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Try() antlr.TerminalNode
	CodeBlock() ICodeBlockContext
	CatchStmt() ICatchStmtContext

	// IsTryCatchSmtContext differentiates from other interfaces.
	IsTryCatchSmtContext()
}

type TryCatchSmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTryCatchSmtContext() *TryCatchSmtContext {
	var p = new(TryCatchSmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_tryCatchSmt
	return p
}

func InitEmptyTryCatchSmtContext(p *TryCatchSmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_tryCatchSmt
}

func (*TryCatchSmtContext) IsTryCatchSmtContext() {}

func NewTryCatchSmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TryCatchSmtContext {
	var p = new(TryCatchSmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_tryCatchSmt

	return p
}

func (s *TryCatchSmtContext) GetParser() antlr.Parser { return s.parser }

func (s *TryCatchSmtContext) Try() antlr.TerminalNode {
	return s.GetToken(V2ParserTry, 0)
}

func (s *TryCatchSmtContext) CodeBlock() ICodeBlockContext {
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

func (s *TryCatchSmtContext) CatchStmt() ICatchStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICatchStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICatchStmtContext)
}

func (s *TryCatchSmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TryCatchSmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TryCatchSmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitTryCatchSmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) TryCatchSmt() (localctx ITryCatchSmtContext) {
	localctx = NewTryCatchSmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, V2ParserRULE_tryCatchSmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(549)
		p.Match(V2ParserTry)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(550)
		p.CodeBlock()
	}
	{
		p.SetState(551)
		p.CatchStmt()
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

// ICatchStmtContext is an interface to support dynamic dispatch.
type ICatchStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Catch() antlr.TerminalNode
	CodeBlock() ICodeBlockContext
	LParen() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	RParen() antlr.TerminalNode
	Type_() ITypeContext

	// IsCatchStmtContext differentiates from other interfaces.
	IsCatchStmtContext()
}

type CatchStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCatchStmtContext() *CatchStmtContext {
	var p = new(CatchStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_catchStmt
	return p
}

func InitEmptyCatchStmtContext(p *CatchStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_catchStmt
}

func (*CatchStmtContext) IsCatchStmtContext() {}

func NewCatchStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CatchStmtContext {
	var p = new(CatchStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_catchStmt

	return p
}

func (s *CatchStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *CatchStmtContext) Catch() antlr.TerminalNode {
	return s.GetToken(V2ParserCatch, 0)
}

func (s *CatchStmtContext) CodeBlock() ICodeBlockContext {
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

func (s *CatchStmtContext) LParen() antlr.TerminalNode {
	return s.GetToken(V2ParserLParen, 0)
}

func (s *CatchStmtContext) Identifier() antlr.TerminalNode {
	return s.GetToken(V2ParserIdentifier, 0)
}

func (s *CatchStmtContext) RParen() antlr.TerminalNode {
	return s.GetToken(V2ParserRParen, 0)
}

func (s *CatchStmtContext) Type_() ITypeContext {
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

func (s *CatchStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CatchStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CatchStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitCatchStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) CatchStmt() (localctx ICatchStmtContext) {
	localctx = NewCatchStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, V2ParserRULE_catchStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(553)
		p.Match(V2ParserCatch)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(560)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == V2ParserLParen {
		{
			p.SetState(554)
			p.Match(V2ParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(556)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 66, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(555)
				p.Type_()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(558)
			p.Match(V2ParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(559)
			p.Match(V2ParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(562)
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
	p.EnterRule(localctx, 78, V2ParserRULE_loopStmt)
	p.SetState(567)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 68, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(564)
			p.IterFor()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(565)
			p.CStyleFor()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(566)
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

	// Getter signatures
	For() antlr.TerminalNode
	CodeBlock() ICodeBlockContext
	CStyleForSign() ICStyleForSignContext
	LParen() antlr.TerminalNode
	RParen() antlr.TerminalNode

	// IsCStyleForContext differentiates from other interfaces.
	IsCStyleForContext()
}

type CStyleForContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *CStyleForContext) For() antlr.TerminalNode {
	return s.GetToken(V2ParserFor, 0)
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

func (s *CStyleForContext) CStyleForSign() ICStyleForSignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICStyleForSignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICStyleForSignContext)
}

func (s *CStyleForContext) LParen() antlr.TerminalNode {
	return s.GetToken(V2ParserLParen, 0)
}

func (s *CStyleForContext) RParen() antlr.TerminalNode {
	return s.GetToken(V2ParserRParen, 0)
}

func (s *CStyleForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CStyleForContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
	p.EnterRule(localctx, 80, V2ParserRULE_cStyleFor)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(569)
		p.Match(V2ParserFor)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(575)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case V2ParserLParen:
		{
			p.SetState(570)
			p.Match(V2ParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(571)
			p.CStyleForSign()
		}
		{
			p.SetState(572)
			p.Match(V2ParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case V2ParserLBrack, V2ParserSemi, V2ParserMap, V2ParserFunction, V2ParserReturn, V2ParserOpen, V2ParserTry, V2ParserIf, V2ParserFor, V2ParserMatch, V2ParserBreak, V2ParserContinue, V2ParserIdentifier:
		{
			p.SetState(574)
			p.CStyleForSign()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(577)
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

// ICStyleForSignContext is an interface to support dynamic dispatch.
type ICStyleForSignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOnInit returns the onInit rule contexts.
	GetOnInit() IStmtContext

	// GetOnCondition returns the onCondition rule contexts.
	GetOnCondition() IExprContext

	// GetOnEnd returns the onEnd rule contexts.
	GetOnEnd() IStmtContext

	// SetOnInit sets the onInit rule contexts.
	SetOnInit(IStmtContext)

	// SetOnCondition sets the onCondition rule contexts.
	SetOnCondition(IExprContext)

	// SetOnEnd sets the onEnd rule contexts.
	SetOnEnd(IStmtContext)

	// Getter signatures
	AllSemi() []antlr.TerminalNode
	Semi(i int) antlr.TerminalNode
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext
	Expr() IExprContext

	// IsCStyleForSignContext differentiates from other interfaces.
	IsCStyleForSignContext()
}

type CStyleForSignContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	onInit      IStmtContext
	onCondition IExprContext
	onEnd       IStmtContext
}

func NewEmptyCStyleForSignContext() *CStyleForSignContext {
	var p = new(CStyleForSignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_cStyleForSign
	return p
}

func InitEmptyCStyleForSignContext(p *CStyleForSignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_cStyleForSign
}

func (*CStyleForSignContext) IsCStyleForSignContext() {}

func NewCStyleForSignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CStyleForSignContext {
	var p = new(CStyleForSignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_cStyleForSign

	return p
}

func (s *CStyleForSignContext) GetParser() antlr.Parser { return s.parser }

func (s *CStyleForSignContext) GetOnInit() IStmtContext { return s.onInit }

func (s *CStyleForSignContext) GetOnCondition() IExprContext { return s.onCondition }

func (s *CStyleForSignContext) GetOnEnd() IStmtContext { return s.onEnd }

func (s *CStyleForSignContext) SetOnInit(v IStmtContext) { s.onInit = v }

func (s *CStyleForSignContext) SetOnCondition(v IExprContext) { s.onCondition = v }

func (s *CStyleForSignContext) SetOnEnd(v IStmtContext) { s.onEnd = v }

func (s *CStyleForSignContext) AllSemi() []antlr.TerminalNode {
	return s.GetTokens(V2ParserSemi)
}

func (s *CStyleForSignContext) Semi(i int) antlr.TerminalNode {
	return s.GetToken(V2ParserSemi, i)
}

func (s *CStyleForSignContext) AllStmt() []IStmtContext {
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

func (s *CStyleForSignContext) Stmt(i int) IStmtContext {
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

func (s *CStyleForSignContext) Expr() IExprContext {
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

func (s *CStyleForSignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CStyleForSignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CStyleForSignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitCStyleForSign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) CStyleForSign() (localctx ICStyleForSignContext) {
	localctx = NewCStyleForSignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, V2ParserRULE_cStyleForSign)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(580)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&18048112929013762) != 0 {
		{
			p.SetState(579)

			var _x = p.Stmt()

			localctx.(*CStyleForSignContext).onInit = _x
		}

	}
	{
		p.SetState(582)
		p.Match(V2ParserSemi)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(584)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&31490016783892522) != 0 {
		{
			p.SetState(583)

			var _x = p.expr(0)

			localctx.(*CStyleForSignContext).onCondition = _x
		}

	}
	{
		p.SetState(586)
		p.Match(V2ParserSemi)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(588)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 72, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(587)

			var _x = p.Stmt()

			localctx.(*CStyleForSignContext).onEnd = _x
		}

	} else if p.HasError() { // JIM
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

// IIterForContext is an interface to support dynamic dispatch.
type IIterForContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	For() antlr.TerminalNode
	CodeBlock() ICodeBlockContext
	IterForSign() IIterForSignContext
	LParen() antlr.TerminalNode
	RParen() antlr.TerminalNode

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

func (s *IterForContext) IterForSign() IIterForSignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIterForSignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIterForSignContext)
}

func (s *IterForContext) LParen() antlr.TerminalNode {
	return s.GetToken(V2ParserLParen, 0)
}

func (s *IterForContext) RParen() antlr.TerminalNode {
	return s.GetToken(V2ParserRParen, 0)
}

func (s *IterForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IterForContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
	p.EnterRule(localctx, 84, V2ParserRULE_iterFor)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(590)
		p.Match(V2ParserFor)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(596)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case V2ParserLParen:
		{
			p.SetState(591)
			p.Match(V2ParserLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(592)
			p.IterForSign()
		}
		{
			p.SetState(593)
			p.Match(V2ParserRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case V2ParserMap, V2ParserFunction, V2ParserIdentifier:
		{
			p.SetState(595)
			p.IterForSign()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(598)
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

// IIterForSignContext is an interface to support dynamic dispatch.
type IIterForSignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Col() antlr.TerminalNode
	Expr() IExprContext
	Type_() ITypeContext

	// IsIterForSignContext differentiates from other interfaces.
	IsIterForSignContext()
}

type IterForSignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIterForSignContext() *IterForSignContext {
	var p = new(IterForSignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_iterForSign
	return p
}

func InitEmptyIterForSignContext(p *IterForSignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_iterForSign
}

func (*IterForSignContext) IsIterForSignContext() {}

func NewIterForSignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IterForSignContext {
	var p = new(IterForSignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_iterForSign

	return p
}

func (s *IterForSignContext) GetParser() antlr.Parser { return s.parser }

func (s *IterForSignContext) Identifier() antlr.TerminalNode {
	return s.GetToken(V2ParserIdentifier, 0)
}

func (s *IterForSignContext) Col() antlr.TerminalNode {
	return s.GetToken(V2ParserCol, 0)
}

func (s *IterForSignContext) Expr() IExprContext {
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

func (s *IterForSignContext) Type_() ITypeContext {
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

func (s *IterForSignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IterForSignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IterForSignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitIterForSign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) IterForSign() (localctx IIterForSignContext) {
	localctx = NewIterForSignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, V2ParserRULE_iterForSign)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(601)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 74, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(600)
			p.Type_()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(603)
		p.Match(V2ParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(604)
		p.Match(V2ParserCol)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(605)
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

// IWhileStyleForContext is an interface to support dynamic dispatch.
type IWhileStyleForContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	For() antlr.TerminalNode
	CodeBlock() ICodeBlockContext
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
	p.EnterRule(localctx, 88, V2ParserRULE_whileStyleFor)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(607)
		p.Match(V2ParserFor)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(609)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 75, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(608)
			p.expr(0)
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(611)
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
	Expr() IExprContext
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
	p.EnterRule(localctx, 90, V2ParserRULE_matchStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(613)
		p.Match(V2ParserMatch)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(614)
		p.expr(0)
	}
	{
		p.SetState(615)
		p.Match(V2ParserLBrack)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(619)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == V2ParserCase || _la == V2ParserDefault {
		{
			p.SetState(616)
			p.MatchCase()
		}

		p.SetState(621)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(622)
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
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	CaseBlock() ICaseBlockContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode
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

func (s *MatchCaseContext) AllExpr() []IExprContext {
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

func (s *MatchCaseContext) Expr(i int) IExprContext {
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

func (s *MatchCaseContext) CaseBlock() ICaseBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaseBlockContext)
}

func (s *MatchCaseContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(V2ParserComma)
}

func (s *MatchCaseContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(V2ParserComma, i)
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
	p.EnterRule(localctx, 92, V2ParserRULE_matchCase)
	var _la int

	p.SetState(637)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case V2ParserCase:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(624)
			p.Match(V2ParserCase)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(625)
			p.expr(0)
		}
		p.SetState(630)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == V2ParserComma {
			{
				p.SetState(626)
				p.Match(V2ParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(627)
				p.expr(0)
			}

			p.SetState(632)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(633)
			p.CaseBlock()
		}

	case V2ParserDefault:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(635)
			p.Match(V2ParserDefault)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(636)
			p.CaseBlock()
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

// ICaseBlockContext is an interface to support dynamic dispatch.
type ICaseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CodeBlock() ICodeBlockContext
	Col() antlr.TerminalNode
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllSep() []ISepContext
	Sep(i int) ISepContext

	// IsCaseBlockContext differentiates from other interfaces.
	IsCaseBlockContext()
}

type CaseBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseBlockContext() *CaseBlockContext {
	var p = new(CaseBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_caseBlock
	return p
}

func InitEmptyCaseBlockContext(p *CaseBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = V2ParserRULE_caseBlock
}

func (*CaseBlockContext) IsCaseBlockContext() {}

func NewCaseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseBlockContext {
	var p = new(CaseBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserRULE_caseBlock

	return p
}

func (s *CaseBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseBlockContext) CodeBlock() ICodeBlockContext {
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

func (s *CaseBlockContext) Col() antlr.TerminalNode {
	return s.GetToken(V2ParserCol, 0)
}

func (s *CaseBlockContext) AllStmt() []IStmtContext {
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

func (s *CaseBlockContext) Stmt(i int) IStmtContext {
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

func (s *CaseBlockContext) AllExpr() []IExprContext {
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

func (s *CaseBlockContext) Expr(i int) IExprContext {
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

func (s *CaseBlockContext) AllSep() []ISepContext {
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

func (s *CaseBlockContext) Sep(i int) ISepContext {
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

func (s *CaseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitCaseBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2Parser) CaseBlock() (localctx ICaseBlockContext) {
	localctx = NewCaseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, V2ParserRULE_caseBlock)
	var _la int

	p.SetState(656)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case V2ParserLBrack:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(639)
			p.CodeBlock()
		}

	case V2ParserCol:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(640)
			p.Match(V2ParserCol)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(653)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&31523727982198826) != 0 {
			p.SetState(643)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 79, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(641)
					p.Stmt()
				}

			case 2:
				{
					p.SetState(642)
					p.expr(0)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}
			p.SetState(648)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == V2ParserSemi || _la == V2ParserNewLine {
				{
					p.SetState(645)
					p.Sep()
				}

				p.SetState(650)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

			p.SetState(655)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
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
	AllExprWithLambda() []IExprWithLambdaContext
	ExprWithLambda(i int) IExprWithLambdaContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

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

func (s *JumpStmtContext) AllExprWithLambda() []IExprWithLambdaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprWithLambdaContext); ok {
			len++
		}
	}

	tst := make([]IExprWithLambdaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprWithLambdaContext); ok {
			tst[i] = t.(IExprWithLambdaContext)
			i++
		}
	}

	return tst
}

func (s *JumpStmtContext) ExprWithLambda(i int) IExprWithLambdaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprWithLambdaContext); ok {
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

	return t.(IExprWithLambdaContext)
}

func (s *JumpStmtContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(V2ParserComma)
}

func (s *JumpStmtContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(V2ParserComma, i)
}

func (s *JumpStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JumpStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
	p.EnterRule(localctx, 96, V2ParserRULE_jumpStmt)
	var _la int

	p.SetState(671)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case V2ParserContinue:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(658)
			p.Match(V2ParserContinue)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case V2ParserBreak:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(659)
			p.Match(V2ParserBreak)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case V2ParserReturn:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(660)
			p.Match(V2ParserReturn)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(669)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 84, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(661)
				p.exprWithLambda(0)
			}
			p.SetState(666)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == V2ParserComma {
				{
					p.SetState(662)
					p.Match(V2ParserComma)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(663)
					p.exprWithLambda(0)
				}

				p.SetState(668)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 98, V2ParserRULE_sep)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(673)
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

	case 17:
		var t *ExprWithLambdaContext = nil
		if localctx != nil {
			t = localctx.(*ExprWithLambdaContext)
		}
		return p.ExprWithLambda_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *V2Parser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *V2Parser) ExprWithLambda_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 10:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
