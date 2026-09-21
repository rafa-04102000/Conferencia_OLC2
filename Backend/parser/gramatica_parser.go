// Code generated from Gramatica.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Gramatica

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

type GramaticaParser struct {
	*antlr.BaseParser
}

var GramaticaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func gramaticaParserInit() {
	staticData := &GramaticaParserStaticData
	staticData.LiteralNames = []string{
		"", "'('", "')'", "'{'", "'}'", "'='", "';'", "'[]'", "'struct'", "','",
		"':'", "'+='", "'-='", "'++'", "'--'", "'['", "']'", "'default'", "'in'",
		"'!'", "'-'", "'&'", "'*'", "'.'", "'/'", "'%'", "'+'", "'>='", "'>'",
		"'<='", "'<'", "'=='", "'!='", "'&&'", "'||'", "'main'", "'fn'", "'mut'",
		"'println'", "'if'", "'else'", "'switch'", "'case'", "'for'", "'break'",
		"'continue'", "'return'", "'int'", "'f64'", "'bool'", "'string'", "",
		"'indexOf'", "'join'", "'len'", "'append'", "'Atoi'", "'parseFloat'",
		"'typeOf'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "MAIN", "FUNCTION", "MUT", "PRINT", "IF", "ELSE", "SWTCH", "CASE",
		"FOR", "TRANSFER_STATEMENT_BREAK", "TRANSFER_STATEMENT_CONTINUE", "TRANSFER_STATEMENT_RETURN",
		"PRIMITIVE_INT", "PRIMITIVE_FLOAT", "PRIMITIVE_BOOL", "PRIMITIVE_STRING",
		"BOOL", "INDEXOF", "JOIN", "LENGTH", "APPEND", "ATOI", "PARSEFLOAT",
		"TYPEOF", "ID", "FLOAT", "INT", "STRING", "LINE_COMMENT", "BLOCK_COMMENT",
		"WS",
	}
	staticData.RuleNames = []string{
		"program", "block", "statement", "mainfunction", "declaration", "parameters",
		"parameter", "tiposSlice", "struct_field", "lstStruct", "listSlice",
		"assignment", "print", "printList", "if", "elseIfBlock", "elseBlock",
		"switch", "caseBlock", "defaultBlock", "for", "individualBlock", "transferStatemen",
		"call", "arguments", "primitiveType", "slicePrimitiveValueFunction",
		"expression",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 65, 492, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 61, 8, 1, 10, 1, 12,
		1, 64, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		3, 2, 76, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 3,
		4, 87, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 94, 8, 4, 1, 4, 3, 4,
		97, 8, 4, 1, 4, 1, 4, 1, 4, 3, 4, 102, 8, 4, 1, 4, 3, 4, 105, 8, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 115, 8, 4, 1, 4, 3,
		4, 118, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 124, 8, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 5, 4, 130, 8, 4, 10, 4, 12, 4, 133, 9, 4, 1, 4, 1, 4, 3, 4, 137,
		8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 146, 8, 4, 10, 4,
		12, 4, 149, 9, 4, 1, 4, 1, 4, 3, 4, 153, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		3, 4, 159, 8, 4, 1, 4, 1, 4, 1, 4, 3, 4, 164, 8, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 3, 4, 170, 8, 4, 3, 4, 172, 8, 4, 1, 5, 1, 5, 1, 5, 5, 5, 177, 8, 5,
		10, 5, 12, 5, 180, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 187, 8, 6,
		1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 3, 8, 195, 8, 8, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 10, 5, 10, 204, 8, 10, 10, 10, 12, 10, 207, 9, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 213, 8, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 3, 11, 219, 8, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 225, 8, 11, 1,
		11, 1, 11, 1, 11, 3, 11, 230, 8, 11, 1, 11, 1, 11, 1, 11, 3, 11, 235, 8,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 244, 8, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 250, 8, 11, 3, 11, 252, 8, 11, 1, 12,
		1, 12, 1, 12, 3, 12, 257, 8, 12, 1, 12, 1, 12, 3, 12, 261, 8, 12, 1, 13,
		1, 13, 1, 13, 5, 13, 266, 8, 13, 10, 13, 12, 13, 269, 9, 13, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 277, 8, 14, 10, 14, 12, 14, 280,
		9, 14, 1, 14, 3, 14, 283, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17,
		4, 17, 301, 8, 17, 11, 17, 12, 17, 302, 1, 17, 3, 17, 306, 8, 17, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 3, 20, 331, 8, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 347, 8, 20,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 3, 22, 355, 8, 22, 1, 22, 1,
		22, 3, 22, 359, 8, 22, 1, 22, 1, 22, 3, 22, 363, 8, 22, 1, 22, 3, 22, 366,
		8, 22, 3, 22, 368, 8, 22, 1, 23, 1, 23, 1, 23, 3, 23, 373, 8, 23, 1, 23,
		1, 23, 3, 23, 377, 8, 23, 1, 24, 1, 24, 1, 24, 5, 24, 382, 8, 24, 10, 24,
		12, 24, 385, 9, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 3, 26, 407, 8, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 3, 27, 438, 8, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 461, 8, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 5, 27, 487, 8, 27, 10, 27, 12, 27, 490, 9, 27, 1, 27, 0, 1, 54,
		28, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 0, 6, 1, 0, 47, 50, 2, 0, 22, 22,
		24, 25, 2, 0, 20, 20, 26, 26, 1, 0, 27, 28, 1, 0, 29, 30, 1, 0, 31, 32,
		558, 0, 56, 1, 0, 0, 0, 2, 62, 1, 0, 0, 0, 4, 75, 1, 0, 0, 0, 6, 77, 1,
		0, 0, 0, 8, 171, 1, 0, 0, 0, 10, 173, 1, 0, 0, 0, 12, 186, 1, 0, 0, 0,
		14, 188, 1, 0, 0, 0, 16, 191, 1, 0, 0, 0, 18, 196, 1, 0, 0, 0, 20, 200,
		1, 0, 0, 0, 22, 251, 1, 0, 0, 0, 24, 253, 1, 0, 0, 0, 26, 262, 1, 0, 0,
		0, 28, 270, 1, 0, 0, 0, 30, 284, 1, 0, 0, 0, 32, 291, 1, 0, 0, 0, 34, 296,
		1, 0, 0, 0, 36, 309, 1, 0, 0, 0, 38, 314, 1, 0, 0, 0, 40, 346, 1, 0, 0,
		0, 42, 348, 1, 0, 0, 0, 44, 367, 1, 0, 0, 0, 46, 369, 1, 0, 0, 0, 48, 378,
		1, 0, 0, 0, 50, 386, 1, 0, 0, 0, 52, 406, 1, 0, 0, 0, 54, 460, 1, 0, 0,
		0, 56, 57, 3, 2, 1, 0, 57, 58, 5, 0, 0, 1, 58, 1, 1, 0, 0, 0, 59, 61, 3,
		4, 2, 0, 60, 59, 1, 0, 0, 0, 61, 64, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 62,
		63, 1, 0, 0, 0, 63, 3, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 65, 76, 3, 6, 3,
		0, 66, 76, 3, 8, 4, 0, 67, 76, 3, 22, 11, 0, 68, 76, 3, 24, 12, 0, 69,
		76, 3, 28, 14, 0, 70, 76, 3, 34, 17, 0, 71, 76, 3, 40, 20, 0, 72, 76, 3,
		42, 21, 0, 73, 76, 3, 44, 22, 0, 74, 76, 3, 46, 23, 0, 75, 65, 1, 0, 0,
		0, 75, 66, 1, 0, 0, 0, 75, 67, 1, 0, 0, 0, 75, 68, 1, 0, 0, 0, 75, 69,
		1, 0, 0, 0, 75, 70, 1, 0, 0, 0, 75, 71, 1, 0, 0, 0, 75, 72, 1, 0, 0, 0,
		75, 73, 1, 0, 0, 0, 75, 74, 1, 0, 0, 0, 76, 5, 1, 0, 0, 0, 77, 78, 5, 36,
		0, 0, 78, 79, 5, 35, 0, 0, 79, 80, 5, 1, 0, 0, 80, 81, 5, 2, 0, 0, 81,
		82, 5, 3, 0, 0, 82, 83, 3, 2, 1, 0, 83, 84, 5, 4, 0, 0, 84, 7, 1, 0, 0,
		0, 85, 87, 5, 37, 0, 0, 86, 85, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 88,
		1, 0, 0, 0, 88, 89, 5, 59, 0, 0, 89, 90, 3, 50, 25, 0, 90, 91, 5, 5, 0,
		0, 91, 93, 3, 54, 27, 0, 92, 94, 5, 6, 0, 0, 93, 92, 1, 0, 0, 0, 93, 94,
		1, 0, 0, 0, 94, 172, 1, 0, 0, 0, 95, 97, 5, 37, 0, 0, 96, 95, 1, 0, 0,
		0, 96, 97, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 99, 5, 59, 0, 0, 99, 101,
		3, 50, 25, 0, 100, 102, 5, 6, 0, 0, 101, 100, 1, 0, 0, 0, 101, 102, 1,
		0, 0, 0, 102, 172, 1, 0, 0, 0, 103, 105, 5, 37, 0, 0, 104, 103, 1, 0, 0,
		0, 104, 105, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 107, 5, 59, 0, 0, 107,
		108, 5, 5, 0, 0, 108, 109, 5, 7, 0, 0, 109, 110, 3, 50, 25, 0, 110, 111,
		5, 3, 0, 0, 111, 112, 3, 20, 10, 0, 112, 114, 5, 4, 0, 0, 113, 115, 5,
		6, 0, 0, 114, 113, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 172, 1, 0, 0,
		0, 116, 118, 5, 37, 0, 0, 117, 116, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118,
		119, 1, 0, 0, 0, 119, 120, 5, 59, 0, 0, 120, 121, 5, 7, 0, 0, 121, 123,
		3, 50, 25, 0, 122, 124, 5, 6, 0, 0, 123, 122, 1, 0, 0, 0, 123, 124, 1,
		0, 0, 0, 124, 172, 1, 0, 0, 0, 125, 126, 5, 8, 0, 0, 126, 127, 5, 59, 0,
		0, 127, 131, 5, 3, 0, 0, 128, 130, 3, 16, 8, 0, 129, 128, 1, 0, 0, 0, 130,
		133, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 134,
		1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 134, 136, 5, 4, 0, 0, 135, 137, 5, 6,
		0, 0, 136, 135, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 172, 1, 0, 0, 0,
		138, 139, 5, 59, 0, 0, 139, 140, 5, 5, 0, 0, 140, 141, 5, 59, 0, 0, 141,
		142, 5, 3, 0, 0, 142, 147, 3, 18, 9, 0, 143, 144, 5, 9, 0, 0, 144, 146,
		3, 18, 9, 0, 145, 143, 1, 0, 0, 0, 146, 149, 1, 0, 0, 0, 147, 145, 1, 0,
		0, 0, 147, 148, 1, 0, 0, 0, 148, 150, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0,
		150, 152, 5, 4, 0, 0, 151, 153, 5, 6, 0, 0, 152, 151, 1, 0, 0, 0, 152,
		153, 1, 0, 0, 0, 153, 172, 1, 0, 0, 0, 154, 155, 5, 36, 0, 0, 155, 156,
		5, 59, 0, 0, 156, 158, 5, 1, 0, 0, 157, 159, 3, 10, 5, 0, 158, 157, 1,
		0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 163, 5, 2, 0,
		0, 161, 164, 3, 50, 25, 0, 162, 164, 3, 14, 7, 0, 163, 161, 1, 0, 0, 0,
		163, 162, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165,
		166, 5, 3, 0, 0, 166, 167, 3, 2, 1, 0, 167, 169, 5, 4, 0, 0, 168, 170,
		5, 6, 0, 0, 169, 168, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 172, 1, 0,
		0, 0, 171, 86, 1, 0, 0, 0, 171, 96, 1, 0, 0, 0, 171, 104, 1, 0, 0, 0, 171,
		117, 1, 0, 0, 0, 171, 125, 1, 0, 0, 0, 171, 138, 1, 0, 0, 0, 171, 154,
		1, 0, 0, 0, 172, 9, 1, 0, 0, 0, 173, 178, 3, 12, 6, 0, 174, 175, 5, 9,
		0, 0, 175, 177, 3, 12, 6, 0, 176, 174, 1, 0, 0, 0, 177, 180, 1, 0, 0, 0,
		178, 176, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 11, 1, 0, 0, 0, 180, 178,
		1, 0, 0, 0, 181, 182, 5, 59, 0, 0, 182, 187, 3, 50, 25, 0, 183, 184, 5,
		59, 0, 0, 184, 185, 5, 7, 0, 0, 185, 187, 3, 50, 25, 0, 186, 181, 1, 0,
		0, 0, 186, 183, 1, 0, 0, 0, 187, 13, 1, 0, 0, 0, 188, 189, 5, 7, 0, 0,
		189, 190, 3, 50, 25, 0, 190, 15, 1, 0, 0, 0, 191, 192, 3, 50, 25, 0, 192,
		194, 5, 59, 0, 0, 193, 195, 5, 6, 0, 0, 194, 193, 1, 0, 0, 0, 194, 195,
		1, 0, 0, 0, 195, 17, 1, 0, 0, 0, 196, 197, 5, 59, 0, 0, 197, 198, 5, 10,
		0, 0, 198, 199, 3, 54, 27, 0, 199, 19, 1, 0, 0, 0, 200, 205, 3, 54, 27,
		0, 201, 202, 5, 9, 0, 0, 202, 204, 3, 54, 27, 0, 203, 201, 1, 0, 0, 0,
		204, 207, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206,
		21, 1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 208, 209, 5, 59, 0, 0, 209, 210,
		5, 5, 0, 0, 210, 212, 3, 54, 27, 0, 211, 213, 5, 6, 0, 0, 212, 211, 1,
		0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 252, 1, 0, 0, 0, 214, 215, 5, 59, 0,
		0, 215, 216, 5, 11, 0, 0, 216, 218, 3, 54, 27, 0, 217, 219, 5, 6, 0, 0,
		218, 217, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 252, 1, 0, 0, 0, 220,
		221, 5, 59, 0, 0, 221, 222, 5, 12, 0, 0, 222, 224, 3, 54, 27, 0, 223, 225,
		5, 6, 0, 0, 224, 223, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 252, 1, 0,
		0, 0, 226, 227, 5, 59, 0, 0, 227, 229, 5, 13, 0, 0, 228, 230, 5, 6, 0,
		0, 229, 228, 1, 0, 0, 0, 229, 230, 1, 0, 0, 0, 230, 252, 1, 0, 0, 0, 231,
		232, 5, 59, 0, 0, 232, 234, 5, 14, 0, 0, 233, 235, 5, 6, 0, 0, 234, 233,
		1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 252, 1, 0, 0, 0, 236, 237, 5, 59,
		0, 0, 237, 238, 5, 15, 0, 0, 238, 239, 3, 54, 27, 0, 239, 240, 5, 16, 0,
		0, 240, 241, 5, 5, 0, 0, 241, 243, 3, 54, 27, 0, 242, 244, 5, 6, 0, 0,
		243, 242, 1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244, 252, 1, 0, 0, 0, 245,
		246, 3, 54, 27, 0, 246, 247, 5, 5, 0, 0, 247, 249, 3, 54, 27, 0, 248, 250,
		5, 6, 0, 0, 249, 248, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 252, 1, 0,
		0, 0, 251, 208, 1, 0, 0, 0, 251, 214, 1, 0, 0, 0, 251, 220, 1, 0, 0, 0,
		251, 226, 1, 0, 0, 0, 251, 231, 1, 0, 0, 0, 251, 236, 1, 0, 0, 0, 251,
		245, 1, 0, 0, 0, 252, 23, 1, 0, 0, 0, 253, 254, 5, 38, 0, 0, 254, 256,
		5, 1, 0, 0, 255, 257, 3, 26, 13, 0, 256, 255, 1, 0, 0, 0, 256, 257, 1,
		0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 260, 5, 2, 0, 0, 259, 261, 5, 6, 0,
		0, 260, 259, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261, 25, 1, 0, 0, 0, 262,
		267, 3, 54, 27, 0, 263, 264, 5, 9, 0, 0, 264, 266, 3, 54, 27, 0, 265, 263,
		1, 0, 0, 0, 266, 269, 1, 0, 0, 0, 267, 265, 1, 0, 0, 0, 267, 268, 1, 0,
		0, 0, 268, 27, 1, 0, 0, 0, 269, 267, 1, 0, 0, 0, 270, 271, 5, 39, 0, 0,
		271, 272, 3, 54, 27, 0, 272, 273, 5, 3, 0, 0, 273, 274, 3, 2, 1, 0, 274,
		278, 5, 4, 0, 0, 275, 277, 3, 30, 15, 0, 276, 275, 1, 0, 0, 0, 277, 280,
		1, 0, 0, 0, 278, 276, 1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 282, 1, 0,
		0, 0, 280, 278, 1, 0, 0, 0, 281, 283, 3, 32, 16, 0, 282, 281, 1, 0, 0,
		0, 282, 283, 1, 0, 0, 0, 283, 29, 1, 0, 0, 0, 284, 285, 5, 40, 0, 0, 285,
		286, 5, 39, 0, 0, 286, 287, 3, 54, 27, 0, 287, 288, 5, 3, 0, 0, 288, 289,
		3, 2, 1, 0, 289, 290, 5, 4, 0, 0, 290, 31, 1, 0, 0, 0, 291, 292, 5, 40,
		0, 0, 292, 293, 5, 3, 0, 0, 293, 294, 3, 2, 1, 0, 294, 295, 5, 4, 0, 0,
		295, 33, 1, 0, 0, 0, 296, 297, 5, 41, 0, 0, 297, 298, 3, 54, 27, 0, 298,
		300, 5, 3, 0, 0, 299, 301, 3, 36, 18, 0, 300, 299, 1, 0, 0, 0, 301, 302,
		1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 302, 303, 1, 0, 0, 0, 303, 305, 1, 0,
		0, 0, 304, 306, 3, 38, 19, 0, 305, 304, 1, 0, 0, 0, 305, 306, 1, 0, 0,
		0, 306, 307, 1, 0, 0, 0, 307, 308, 5, 4, 0, 0, 308, 35, 1, 0, 0, 0, 309,
		310, 5, 42, 0, 0, 310, 311, 3, 54, 27, 0, 311, 312, 5, 10, 0, 0, 312, 313,
		3, 2, 1, 0, 313, 37, 1, 0, 0, 0, 314, 315, 5, 17, 0, 0, 315, 316, 5, 10,
		0, 0, 316, 317, 3, 2, 1, 0, 317, 39, 1, 0, 0, 0, 318, 319, 5, 43, 0, 0,
		319, 320, 3, 54, 27, 0, 320, 321, 5, 3, 0, 0, 321, 322, 3, 2, 1, 0, 322,
		323, 5, 4, 0, 0, 323, 347, 1, 0, 0, 0, 324, 325, 5, 43, 0, 0, 325, 326,
		3, 8, 4, 0, 326, 327, 5, 6, 0, 0, 327, 328, 3, 54, 27, 0, 328, 330, 5,
		6, 0, 0, 329, 331, 3, 22, 11, 0, 330, 329, 1, 0, 0, 0, 330, 331, 1, 0,
		0, 0, 331, 332, 1, 0, 0, 0, 332, 333, 5, 3, 0, 0, 333, 334, 3, 2, 1, 0,
		334, 335, 5, 4, 0, 0, 335, 347, 1, 0, 0, 0, 336, 337, 5, 43, 0, 0, 337,
		338, 5, 59, 0, 0, 338, 339, 5, 9, 0, 0, 339, 340, 5, 59, 0, 0, 340, 341,
		5, 18, 0, 0, 341, 342, 3, 54, 27, 0, 342, 343, 5, 3, 0, 0, 343, 344, 3,
		2, 1, 0, 344, 345, 5, 4, 0, 0, 345, 347, 1, 0, 0, 0, 346, 318, 1, 0, 0,
		0, 346, 324, 1, 0, 0, 0, 346, 336, 1, 0, 0, 0, 347, 41, 1, 0, 0, 0, 348,
		349, 5, 3, 0, 0, 349, 350, 3, 2, 1, 0, 350, 351, 5, 4, 0, 0, 351, 43, 1,
		0, 0, 0, 352, 354, 5, 44, 0, 0, 353, 355, 5, 6, 0, 0, 354, 353, 1, 0, 0,
		0, 354, 355, 1, 0, 0, 0, 355, 368, 1, 0, 0, 0, 356, 358, 5, 45, 0, 0, 357,
		359, 5, 6, 0, 0, 358, 357, 1, 0, 0, 0, 358, 359, 1, 0, 0, 0, 359, 368,
		1, 0, 0, 0, 360, 362, 5, 46, 0, 0, 361, 363, 3, 54, 27, 0, 362, 361, 1,
		0, 0, 0, 362, 363, 1, 0, 0, 0, 363, 365, 1, 0, 0, 0, 364, 366, 5, 6, 0,
		0, 365, 364, 1, 0, 0, 0, 365, 366, 1, 0, 0, 0, 366, 368, 1, 0, 0, 0, 367,
		352, 1, 0, 0, 0, 367, 356, 1, 0, 0, 0, 367, 360, 1, 0, 0, 0, 368, 45, 1,
		0, 0, 0, 369, 370, 5, 59, 0, 0, 370, 372, 5, 1, 0, 0, 371, 373, 3, 48,
		24, 0, 372, 371, 1, 0, 0, 0, 372, 373, 1, 0, 0, 0, 373, 374, 1, 0, 0, 0,
		374, 376, 5, 2, 0, 0, 375, 377, 5, 6, 0, 0, 376, 375, 1, 0, 0, 0, 376,
		377, 1, 0, 0, 0, 377, 47, 1, 0, 0, 0, 378, 383, 3, 54, 27, 0, 379, 380,
		5, 9, 0, 0, 380, 382, 3, 54, 27, 0, 381, 379, 1, 0, 0, 0, 382, 385, 1,
		0, 0, 0, 383, 381, 1, 0, 0, 0, 383, 384, 1, 0, 0, 0, 384, 49, 1, 0, 0,
		0, 385, 383, 1, 0, 0, 0, 386, 387, 7, 0, 0, 0, 387, 51, 1, 0, 0, 0, 388,
		389, 5, 52, 0, 0, 389, 390, 5, 1, 0, 0, 390, 391, 5, 59, 0, 0, 391, 392,
		5, 9, 0, 0, 392, 393, 3, 54, 27, 0, 393, 394, 5, 2, 0, 0, 394, 407, 1,
		0, 0, 0, 395, 396, 5, 53, 0, 0, 396, 397, 5, 1, 0, 0, 397, 398, 5, 59,
		0, 0, 398, 399, 5, 9, 0, 0, 399, 400, 3, 54, 27, 0, 400, 401, 5, 2, 0,
		0, 401, 407, 1, 0, 0, 0, 402, 403, 5, 54, 0, 0, 403, 404, 5, 1, 0, 0, 404,
		405, 5, 59, 0, 0, 405, 407, 5, 2, 0, 0, 406, 388, 1, 0, 0, 0, 406, 395,
		1, 0, 0, 0, 406, 402, 1, 0, 0, 0, 407, 53, 1, 0, 0, 0, 408, 409, 6, 27,
		-1, 0, 409, 410, 5, 19, 0, 0, 410, 461, 3, 54, 27, 25, 411, 412, 5, 20,
		0, 0, 412, 461, 3, 54, 27, 24, 413, 414, 5, 21, 0, 0, 414, 461, 5, 59,
		0, 0, 415, 416, 5, 22, 0, 0, 416, 461, 5, 59, 0, 0, 417, 418, 5, 1, 0,
		0, 418, 419, 3, 54, 27, 0, 419, 420, 5, 2, 0, 0, 420, 461, 1, 0, 0, 0,
		421, 461, 3, 52, 26, 0, 422, 423, 5, 55, 0, 0, 423, 424, 5, 1, 0, 0, 424,
		425, 5, 59, 0, 0, 425, 426, 5, 9, 0, 0, 426, 427, 3, 54, 27, 0, 427, 428,
		5, 2, 0, 0, 428, 461, 1, 0, 0, 0, 429, 430, 5, 59, 0, 0, 430, 431, 5, 15,
		0, 0, 431, 432, 3, 54, 27, 0, 432, 433, 5, 16, 0, 0, 433, 461, 1, 0, 0,
		0, 434, 435, 5, 59, 0, 0, 435, 437, 5, 1, 0, 0, 436, 438, 3, 48, 24, 0,
		437, 436, 1, 0, 0, 0, 437, 438, 1, 0, 0, 0, 438, 439, 1, 0, 0, 0, 439,
		461, 5, 2, 0, 0, 440, 441, 5, 56, 0, 0, 441, 442, 5, 1, 0, 0, 442, 443,
		3, 54, 27, 0, 443, 444, 5, 2, 0, 0, 444, 461, 1, 0, 0, 0, 445, 446, 5,
		57, 0, 0, 446, 447, 5, 1, 0, 0, 447, 448, 3, 54, 27, 0, 448, 449, 5, 2,
		0, 0, 449, 461, 1, 0, 0, 0, 450, 451, 5, 58, 0, 0, 451, 452, 5, 1, 0, 0,
		452, 453, 3, 54, 27, 0, 453, 454, 5, 2, 0, 0, 454, 461, 1, 0, 0, 0, 455,
		461, 5, 59, 0, 0, 456, 461, 5, 61, 0, 0, 457, 461, 5, 60, 0, 0, 458, 461,
		5, 51, 0, 0, 459, 461, 5, 62, 0, 0, 460, 408, 1, 0, 0, 0, 460, 411, 1,
		0, 0, 0, 460, 413, 1, 0, 0, 0, 460, 415, 1, 0, 0, 0, 460, 417, 1, 0, 0,
		0, 460, 421, 1, 0, 0, 0, 460, 422, 1, 0, 0, 0, 460, 429, 1, 0, 0, 0, 460,
		434, 1, 0, 0, 0, 460, 440, 1, 0, 0, 0, 460, 445, 1, 0, 0, 0, 460, 450,
		1, 0, 0, 0, 460, 455, 1, 0, 0, 0, 460, 456, 1, 0, 0, 0, 460, 457, 1, 0,
		0, 0, 460, 458, 1, 0, 0, 0, 460, 459, 1, 0, 0, 0, 461, 488, 1, 0, 0, 0,
		462, 463, 10, 20, 0, 0, 463, 464, 7, 1, 0, 0, 464, 487, 3, 54, 27, 21,
		465, 466, 10, 19, 0, 0, 466, 467, 7, 2, 0, 0, 467, 487, 3, 54, 27, 20,
		468, 469, 10, 18, 0, 0, 469, 470, 7, 3, 0, 0, 470, 487, 3, 54, 27, 19,
		471, 472, 10, 17, 0, 0, 472, 473, 7, 4, 0, 0, 473, 487, 3, 54, 27, 18,
		474, 475, 10, 16, 0, 0, 475, 476, 7, 5, 0, 0, 476, 487, 3, 54, 27, 17,
		477, 478, 10, 15, 0, 0, 478, 479, 5, 33, 0, 0, 479, 487, 3, 54, 27, 16,
		480, 481, 10, 14, 0, 0, 481, 482, 5, 34, 0, 0, 482, 487, 3, 54, 27, 15,
		483, 484, 10, 21, 0, 0, 484, 485, 5, 23, 0, 0, 485, 487, 5, 59, 0, 0, 486,
		462, 1, 0, 0, 0, 486, 465, 1, 0, 0, 0, 486, 468, 1, 0, 0, 0, 486, 471,
		1, 0, 0, 0, 486, 474, 1, 0, 0, 0, 486, 477, 1, 0, 0, 0, 486, 480, 1, 0,
		0, 0, 486, 483, 1, 0, 0, 0, 487, 490, 1, 0, 0, 0, 488, 486, 1, 0, 0, 0,
		488, 489, 1, 0, 0, 0, 489, 55, 1, 0, 0, 0, 490, 488, 1, 0, 0, 0, 52, 62,
		75, 86, 93, 96, 101, 104, 114, 117, 123, 131, 136, 147, 152, 158, 163,
		169, 171, 178, 186, 194, 205, 212, 218, 224, 229, 234, 243, 249, 251, 256,
		260, 267, 278, 282, 302, 305, 330, 346, 354, 358, 362, 365, 367, 372, 376,
		383, 406, 437, 460, 486, 488,
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

// GramaticaParserInit initializes any static state used to implement GramaticaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGramaticaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GramaticaParserInit() {
	staticData := &GramaticaParserStaticData
	staticData.once.Do(gramaticaParserInit)
}

// NewGramaticaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGramaticaParser(input antlr.TokenStream) *GramaticaParser {
	GramaticaParserInit()
	this := new(GramaticaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GramaticaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Gramatica.g4"

	return this
}

// GramaticaParser tokens.
const (
	GramaticaParserEOF                         = antlr.TokenEOF
	GramaticaParserT__0                        = 1
	GramaticaParserT__1                        = 2
	GramaticaParserT__2                        = 3
	GramaticaParserT__3                        = 4
	GramaticaParserT__4                        = 5
	GramaticaParserT__5                        = 6
	GramaticaParserT__6                        = 7
	GramaticaParserT__7                        = 8
	GramaticaParserT__8                        = 9
	GramaticaParserT__9                        = 10
	GramaticaParserT__10                       = 11
	GramaticaParserT__11                       = 12
	GramaticaParserT__12                       = 13
	GramaticaParserT__13                       = 14
	GramaticaParserT__14                       = 15
	GramaticaParserT__15                       = 16
	GramaticaParserT__16                       = 17
	GramaticaParserT__17                       = 18
	GramaticaParserT__18                       = 19
	GramaticaParserT__19                       = 20
	GramaticaParserT__20                       = 21
	GramaticaParserT__21                       = 22
	GramaticaParserT__22                       = 23
	GramaticaParserT__23                       = 24
	GramaticaParserT__24                       = 25
	GramaticaParserT__25                       = 26
	GramaticaParserT__26                       = 27
	GramaticaParserT__27                       = 28
	GramaticaParserT__28                       = 29
	GramaticaParserT__29                       = 30
	GramaticaParserT__30                       = 31
	GramaticaParserT__31                       = 32
	GramaticaParserT__32                       = 33
	GramaticaParserT__33                       = 34
	GramaticaParserMAIN                        = 35
	GramaticaParserFUNCTION                    = 36
	GramaticaParserMUT                         = 37
	GramaticaParserPRINT                       = 38
	GramaticaParserIF                          = 39
	GramaticaParserELSE                        = 40
	GramaticaParserSWTCH                       = 41
	GramaticaParserCASE                        = 42
	GramaticaParserFOR                         = 43
	GramaticaParserTRANSFER_STATEMENT_BREAK    = 44
	GramaticaParserTRANSFER_STATEMENT_CONTINUE = 45
	GramaticaParserTRANSFER_STATEMENT_RETURN   = 46
	GramaticaParserPRIMITIVE_INT               = 47
	GramaticaParserPRIMITIVE_FLOAT             = 48
	GramaticaParserPRIMITIVE_BOOL              = 49
	GramaticaParserPRIMITIVE_STRING            = 50
	GramaticaParserBOOL                        = 51
	GramaticaParserINDEXOF                     = 52
	GramaticaParserJOIN                        = 53
	GramaticaParserLENGTH                      = 54
	GramaticaParserAPPEND                      = 55
	GramaticaParserATOI                        = 56
	GramaticaParserPARSEFLOAT                  = 57
	GramaticaParserTYPEOF                      = 58
	GramaticaParserID                          = 59
	GramaticaParserFLOAT                       = 60
	GramaticaParserINT                         = 61
	GramaticaParserSTRING                      = 62
	GramaticaParserLINE_COMMENT                = 63
	GramaticaParserBLOCK_COMMENT               = 64
	GramaticaParserWS                          = 65
)

// GramaticaParser rules.
const (
	GramaticaParserRULE_program                     = 0
	GramaticaParserRULE_block                       = 1
	GramaticaParserRULE_statement                   = 2
	GramaticaParserRULE_mainfunction                = 3
	GramaticaParserRULE_declaration                 = 4
	GramaticaParserRULE_parameters                  = 5
	GramaticaParserRULE_parameter                   = 6
	GramaticaParserRULE_tiposSlice                  = 7
	GramaticaParserRULE_struct_field                = 8
	GramaticaParserRULE_lstStruct                   = 9
	GramaticaParserRULE_listSlice                   = 10
	GramaticaParserRULE_assignment                  = 11
	GramaticaParserRULE_print                       = 12
	GramaticaParserRULE_printList                   = 13
	GramaticaParserRULE_if                          = 14
	GramaticaParserRULE_elseIfBlock                 = 15
	GramaticaParserRULE_elseBlock                   = 16
	GramaticaParserRULE_switch                      = 17
	GramaticaParserRULE_caseBlock                   = 18
	GramaticaParserRULE_defaultBlock                = 19
	GramaticaParserRULE_for                         = 20
	GramaticaParserRULE_individualBlock             = 21
	GramaticaParserRULE_transferStatemen            = 22
	GramaticaParserRULE_call                        = 23
	GramaticaParserRULE_arguments                   = 24
	GramaticaParserRULE_primitiveType               = 25
	GramaticaParserRULE_slicePrimitiveValueFunction = 26
	GramaticaParserRULE_expression                  = 27
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	EOF() antlr.TerminalNode

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
	p.RuleIndex = GramaticaParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(GramaticaParserEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GramaticaParserRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Block()
	}
	{
		p.SetState(57)
		p.Match(GramaticaParserEOF)
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GramaticaParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9221255408259694858) != 0 {
		{
			p.SetState(59)
			p.Statement()
		}

		p.SetState(64)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mainfunction() IMainfunctionContext
	Declaration() IDeclarationContext
	Assignment() IAssignmentContext
	Print_() IPrintContext
	If_() IIfContext
	Switch_() ISwitchContext
	For_() IForContext
	IndividualBlock() IIndividualBlockContext
	TransferStatemen() ITransferStatemenContext
	Call() ICallContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Mainfunction() IMainfunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMainfunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMainfunctionContext)
}

func (s *StatementContext) Declaration() IDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *StatementContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) Print_() IPrintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintContext)
}

func (s *StatementContext) If_() IIfContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfContext)
}

func (s *StatementContext) Switch_() ISwitchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchContext)
}

func (s *StatementContext) For_() IForContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForContext)
}

func (s *StatementContext) IndividualBlock() IIndividualBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndividualBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndividualBlockContext)
}

func (s *StatementContext) TransferStatemen() ITransferStatemenContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITransferStatemenContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITransferStatemenContext)
}

func (s *StatementContext) Call() ICallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GramaticaParserRULE_statement)
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.Mainfunction()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.Declaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(67)
			p.Assignment()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(68)
			p.Print_()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(69)
			p.If_()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(70)
			p.Switch_()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(71)
			p.For_()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(72)
			p.IndividualBlock()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(73)
			p.TransferStatemen()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(74)
			p.Call()
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

// IMainfunctionContext is an interface to support dynamic dispatch.
type IMainfunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNCTION() antlr.TerminalNode
	MAIN() antlr.TerminalNode
	Block() IBlockContext

	// IsMainfunctionContext differentiates from other interfaces.
	IsMainfunctionContext()
}

type MainfunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainfunctionContext() *MainfunctionContext {
	var p = new(MainfunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_mainfunction
	return p
}

func InitEmptyMainfunctionContext(p *MainfunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_mainfunction
}

func (*MainfunctionContext) IsMainfunctionContext() {}

func NewMainfunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainfunctionContext {
	var p = new(MainfunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_mainfunction

	return p
}

func (s *MainfunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *MainfunctionContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(GramaticaParserFUNCTION, 0)
}

func (s *MainfunctionContext) MAIN() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMAIN, 0)
}

func (s *MainfunctionContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *MainfunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainfunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainfunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterMainfunction(s)
	}
}

func (s *MainfunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitMainfunction(s)
	}
}

func (s *MainfunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitMainfunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) Mainfunction() (localctx IMainfunctionContext) {
	localctx = NewMainfunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GramaticaParserRULE_mainfunction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(GramaticaParserFUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(78)
		p.Match(GramaticaParserMAIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(79)
		p.Match(GramaticaParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(80)
		p.Match(GramaticaParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)
		p.Match(GramaticaParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.Block()
	}
	{
		p.SetState(83)
		p.Match(GramaticaParserT__3)
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

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_declaration
	return p
}

func InitEmptyDeclarationContext(p *DeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_declaration
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) CopyAll(ctx *DeclarationContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DeclarationImplicitSliceContext struct {
	DeclarationContext
}

func NewDeclarationImplicitSliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclarationImplicitSliceContext {
	var p = new(DeclarationImplicitSliceContext)

	InitEmptyDeclarationContext(&p.DeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclarationContext))

	return p
}

func (s *DeclarationImplicitSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationImplicitSliceContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *DeclarationImplicitSliceContext) PrimitiveType() IPrimitiveTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *DeclarationImplicitSliceContext) MUT() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMUT, 0)
}

func (s *DeclarationImplicitSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterDeclarationImplicitSlice(s)
	}
}

func (s *DeclarationImplicitSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitDeclarationImplicitSlice(s)
	}
}

func (s *DeclarationImplicitSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitDeclarationImplicitSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionDeclarationContext struct {
	DeclarationContext
}

func NewFunctionDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	InitEmptyDeclarationContext(&p.DeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclarationContext))

	return p
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(GramaticaParserFUNCTION, 0)
}

func (s *FunctionDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *FunctionDeclarationContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionDeclarationContext) Parameters() IParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *FunctionDeclarationContext) PrimitiveType() IPrimitiveTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *FunctionDeclarationContext) TiposSlice() ITiposSliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITiposSliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITiposSliceContext)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitFunctionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclarationExplicitVarContext struct {
	DeclarationContext
}

func NewDeclarationExplicitVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclarationExplicitVarContext {
	var p = new(DeclarationExplicitVarContext)

	InitEmptyDeclarationContext(&p.DeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclarationContext))

	return p
}

func (s *DeclarationExplicitVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationExplicitVarContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *DeclarationExplicitVarContext) PrimitiveType() IPrimitiveTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *DeclarationExplicitVarContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DeclarationExplicitVarContext) MUT() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMUT, 0)
}

func (s *DeclarationExplicitVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterDeclarationExplicitVar(s)
	}
}

func (s *DeclarationExplicitVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitDeclarationExplicitVar(s)
	}
}

func (s *DeclarationExplicitVarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitDeclarationExplicitVar(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructDeclarationContext struct {
	DeclarationContext
}

func NewStructDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructDeclarationContext {
	var p = new(StructDeclarationContext)

	InitEmptyDeclarationContext(&p.DeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclarationContext))

	return p
}

func (s *StructDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *StructDeclarationContext) AllStruct_field() []IStruct_fieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStruct_fieldContext); ok {
			len++
		}
	}

	tst := make([]IStruct_fieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStruct_fieldContext); ok {
			tst[i] = t.(IStruct_fieldContext)
			i++
		}
	}

	return tst
}

func (s *StructDeclarationContext) Struct_field(i int) IStruct_fieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_fieldContext); ok {
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

	return t.(IStruct_fieldContext)
}

func (s *StructDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterStructDeclaration(s)
	}
}

func (s *StructDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitStructDeclaration(s)
	}
}

func (s *StructDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitStructDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclarationExplicitSliceContext struct {
	DeclarationContext
}

func NewDeclarationExplicitSliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclarationExplicitSliceContext {
	var p = new(DeclarationExplicitSliceContext)

	InitEmptyDeclarationContext(&p.DeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclarationContext))

	return p
}

func (s *DeclarationExplicitSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationExplicitSliceContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *DeclarationExplicitSliceContext) PrimitiveType() IPrimitiveTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *DeclarationExplicitSliceContext) ListSlice() IListSliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListSliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListSliceContext)
}

func (s *DeclarationExplicitSliceContext) MUT() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMUT, 0)
}

func (s *DeclarationExplicitSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterDeclarationExplicitSlice(s)
	}
}

func (s *DeclarationExplicitSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitDeclarationExplicitSlice(s)
	}
}

func (s *DeclarationExplicitSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitDeclarationExplicitSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructInstanceDeclarationContext struct {
	DeclarationContext
}

func NewStructInstanceDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructInstanceDeclarationContext {
	var p = new(StructInstanceDeclarationContext)

	InitEmptyDeclarationContext(&p.DeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclarationContext))

	return p
}

func (s *StructInstanceDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructInstanceDeclarationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(GramaticaParserID)
}

func (s *StructInstanceDeclarationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, i)
}

func (s *StructInstanceDeclarationContext) AllLstStruct() []ILstStructContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILstStructContext); ok {
			len++
		}
	}

	tst := make([]ILstStructContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILstStructContext); ok {
			tst[i] = t.(ILstStructContext)
			i++
		}
	}

	return tst
}

func (s *StructInstanceDeclarationContext) LstStruct(i int) ILstStructContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILstStructContext); ok {
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

	return t.(ILstStructContext)
}

func (s *StructInstanceDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterStructInstanceDeclaration(s)
	}
}

func (s *StructInstanceDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitStructInstanceDeclaration(s)
	}
}

func (s *StructInstanceDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitStructInstanceDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclarationImplicitVarContext struct {
	DeclarationContext
}

func NewDeclarationImplicitVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclarationImplicitVarContext {
	var p = new(DeclarationImplicitVarContext)

	InitEmptyDeclarationContext(&p.DeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclarationContext))

	return p
}

func (s *DeclarationImplicitVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationImplicitVarContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *DeclarationImplicitVarContext) PrimitiveType() IPrimitiveTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *DeclarationImplicitVarContext) MUT() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMUT, 0)
}

func (s *DeclarationImplicitVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterDeclarationImplicitVar(s)
	}
}

func (s *DeclarationImplicitVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitDeclarationImplicitVar(s)
	}
}

func (s *DeclarationImplicitVarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitDeclarationImplicitVar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GramaticaParserRULE_declaration)
	var _la int

	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDeclarationExplicitVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramaticaParserMUT {
			{
				p.SetState(85)
				p.Match(GramaticaParserMUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(88)
			p.Match(GramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(89)
			p.PrimitiveType()
		}
		{
			p.SetState(90)
			p.Match(GramaticaParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)
			p.expression(0)
		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(92)
				p.Match(GramaticaParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 2:
		localctx = NewDeclarationImplicitVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramaticaParserMUT {
			{
				p.SetState(95)
				p.Match(GramaticaParserMUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(98)
			p.Match(GramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)
			p.PrimitiveType()
		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(100)
				p.Match(GramaticaParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 3:
		localctx = NewDeclarationExplicitSliceContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramaticaParserMUT {
			{
				p.SetState(103)
				p.Match(GramaticaParserMUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(106)
			p.Match(GramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.Match(GramaticaParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(108)
			p.Match(GramaticaParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)
			p.PrimitiveType()
		}
		{
			p.SetState(110)
			p.Match(GramaticaParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.ListSlice()
		}
		{
			p.SetState(112)
			p.Match(GramaticaParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(114)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(113)
				p.Match(GramaticaParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 4:
		localctx = NewDeclarationImplicitSliceContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramaticaParserMUT {
			{
				p.SetState(116)
				p.Match(GramaticaParserMUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(119)
			p.Match(GramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.Match(GramaticaParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.PrimitiveType()
		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(122)
				p.Match(GramaticaParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 5:
		localctx = NewStructDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(125)
			p.Match(GramaticaParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
			p.Match(GramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			p.Match(GramaticaParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2111062325329920) != 0 {
			{
				p.SetState(128)
				p.Struct_field()
			}

			p.SetState(133)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(134)
			p.Match(GramaticaParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(135)
				p.Match(GramaticaParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 6:
		localctx = NewStructInstanceDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(138)
			p.Match(GramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Match(GramaticaParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.Match(GramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.Match(GramaticaParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.LstStruct()
		}
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GramaticaParserT__8 {
			{
				p.SetState(143)
				p.Match(GramaticaParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(144)
				p.LstStruct()
			}

			p.SetState(149)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(150)
			p.Match(GramaticaParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(152)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(151)
				p.Match(GramaticaParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 7:
		localctx = NewFunctionDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(154)
			p.Match(GramaticaParserFUNCTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(155)
			p.Match(GramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)
			p.Match(GramaticaParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramaticaParserID {
			{
				p.SetState(157)
				p.Parameters()
			}

		}
		{
			p.SetState(160)
			p.Match(GramaticaParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		switch p.GetTokenStream().LA(1) {
		case GramaticaParserPRIMITIVE_INT, GramaticaParserPRIMITIVE_FLOAT, GramaticaParserPRIMITIVE_BOOL, GramaticaParserPRIMITIVE_STRING:
			{
				p.SetState(161)
				p.PrimitiveType()
			}

		case GramaticaParserT__6:
			{
				p.SetState(162)
				p.TiposSlice()
			}

		case GramaticaParserT__2:

		default:
		}
		{
			p.SetState(165)
			p.Match(GramaticaParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)
			p.Block()
		}
		{
			p.SetState(167)
			p.Match(GramaticaParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(168)
				p.Match(GramaticaParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
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

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_parameters
	return p
}

func InitEmptyParametersContext(p *ParametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_parameters
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) AllParameter() []IParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterContext); ok {
			len++
		}
	}

	tst := make([]IParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterContext); ok {
			tst[i] = t.(IParameterContext)
			i++
		}
	}

	return tst
}

func (s *ParametersContext) Parameter(i int) IParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
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

	return t.(IParameterContext)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterParameters(s)
	}
}

func (s *ParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitParameters(s)
	}
}

func (s *ParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GramaticaParserRULE_parameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
		p.Parameter()
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GramaticaParserT__8 {
		{
			p.SetState(174)
			p.Match(GramaticaParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(175)
			p.Parameter()
		}

		p.SetState(180)
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

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_parameter
	return p
}

func InitEmptyParameterContext(p *ParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_parameter
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) CopyAll(ctx *ParameterContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NormalTypeParameterContext struct {
	ParameterContext
}

func NewNormalTypeParameterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NormalTypeParameterContext {
	var p = new(NormalTypeParameterContext)

	InitEmptyParameterContext(&p.ParameterContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParameterContext))

	return p
}

func (s *NormalTypeParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NormalTypeParameterContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *NormalTypeParameterContext) PrimitiveType() IPrimitiveTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *NormalTypeParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterNormalTypeParameter(s)
	}
}

func (s *NormalTypeParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitNormalTypeParameter(s)
	}
}

func (s *NormalTypeParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitNormalTypeParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

type SliceTypeParameterContext struct {
	ParameterContext
}

func NewSliceTypeParameterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceTypeParameterContext {
	var p = new(SliceTypeParameterContext)

	InitEmptyParameterContext(&p.ParameterContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParameterContext))

	return p
}

func (s *SliceTypeParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceTypeParameterContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *SliceTypeParameterContext) PrimitiveType() IPrimitiveTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *SliceTypeParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterSliceTypeParameter(s)
	}
}

func (s *SliceTypeParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitSliceTypeParameter(s)
	}
}

func (s *SliceTypeParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitSliceTypeParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GramaticaParserRULE_parameter)
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNormalTypeParameterContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(181)
			p.Match(GramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.PrimitiveType()
		}

	case 2:
		localctx = NewSliceTypeParameterContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(183)
			p.Match(GramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(184)
			p.Match(GramaticaParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)
			p.PrimitiveType()
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

// ITiposSliceContext is an interface to support dynamic dispatch.
type ITiposSliceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimitiveType() IPrimitiveTypeContext

	// IsTiposSliceContext differentiates from other interfaces.
	IsTiposSliceContext()
}

type TiposSliceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTiposSliceContext() *TiposSliceContext {
	var p = new(TiposSliceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_tiposSlice
	return p
}

func InitEmptyTiposSliceContext(p *TiposSliceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_tiposSlice
}

func (*TiposSliceContext) IsTiposSliceContext() {}

func NewTiposSliceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TiposSliceContext {
	var p = new(TiposSliceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_tiposSlice

	return p
}

func (s *TiposSliceContext) GetParser() antlr.Parser { return s.parser }

func (s *TiposSliceContext) PrimitiveType() IPrimitiveTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *TiposSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TiposSliceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TiposSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterTiposSlice(s)
	}
}

func (s *TiposSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitTiposSlice(s)
	}
}

func (s *TiposSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitTiposSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) TiposSlice() (localctx ITiposSliceContext) {
	localctx = NewTiposSliceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GramaticaParserRULE_tiposSlice)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Match(GramaticaParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(189)
		p.PrimitiveType()
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

// IStruct_fieldContext is an interface to support dynamic dispatch.
type IStruct_fieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimitiveType() IPrimitiveTypeContext
	ID() antlr.TerminalNode

	// IsStruct_fieldContext differentiates from other interfaces.
	IsStruct_fieldContext()
}

type Struct_fieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_fieldContext() *Struct_fieldContext {
	var p = new(Struct_fieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_struct_field
	return p
}

func InitEmptyStruct_fieldContext(p *Struct_fieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_struct_field
}

func (*Struct_fieldContext) IsStruct_fieldContext() {}

func NewStruct_fieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_fieldContext {
	var p = new(Struct_fieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_struct_field

	return p
}

func (s *Struct_fieldContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_fieldContext) PrimitiveType() IPrimitiveTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *Struct_fieldContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *Struct_fieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_fieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Struct_fieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterStruct_field(s)
	}
}

func (s *Struct_fieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitStruct_field(s)
	}
}

func (s *Struct_fieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitStruct_field(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) Struct_field() (localctx IStruct_fieldContext) {
	localctx = NewStruct_fieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GramaticaParserRULE_struct_field)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		p.PrimitiveType()
	}
	{
		p.SetState(192)
		p.Match(GramaticaParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GramaticaParserT__5 {
		{
			p.SetState(193)
			p.Match(GramaticaParserT__5)
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

// ILstStructContext is an interface to support dynamic dispatch.
type ILstStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Expression() IExpressionContext

	// IsLstStructContext differentiates from other interfaces.
	IsLstStructContext()
}

type LstStructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLstStructContext() *LstStructContext {
	var p = new(LstStructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_lstStruct
	return p
}

func InitEmptyLstStructContext(p *LstStructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_lstStruct
}

func (*LstStructContext) IsLstStructContext() {}

func NewLstStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LstStructContext {
	var p = new(LstStructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_lstStruct

	return p
}

func (s *LstStructContext) GetParser() antlr.Parser { return s.parser }

func (s *LstStructContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *LstStructContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LstStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LstStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LstStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterLstStruct(s)
	}
}

func (s *LstStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitLstStruct(s)
	}
}

func (s *LstStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitLstStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) LstStruct() (localctx ILstStructContext) {
	localctx = NewLstStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GramaticaParserRULE_lstStruct)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(GramaticaParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
		p.Match(GramaticaParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(198)
		p.expression(0)
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

// IListSliceContext is an interface to support dynamic dispatch.
type IListSliceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsListSliceContext differentiates from other interfaces.
	IsListSliceContext()
}

type ListSliceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListSliceContext() *ListSliceContext {
	var p = new(ListSliceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_listSlice
	return p
}

func InitEmptyListSliceContext(p *ListSliceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_listSlice
}

func (*ListSliceContext) IsListSliceContext() {}

func NewListSliceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListSliceContext {
	var p = new(ListSliceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_listSlice

	return p
}

func (s *ListSliceContext) GetParser() antlr.Parser { return s.parser }

func (s *ListSliceContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ListSliceContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ListSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListSliceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterListSlice(s)
	}
}

func (s *ListSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitListSlice(s)
	}
}

func (s *ListSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitListSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) ListSlice() (localctx IListSliceContext) {
	localctx = NewListSliceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GramaticaParserRULE_listSlice)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.expression(0)
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GramaticaParserT__8 {
		{
			p.SetState(201)
			p.Match(GramaticaParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(202)
			p.expression(0)
		}

		p.SetState(207)
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

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) CopyAll(ctx *AssignmentContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AssignmentMinusExprContext struct {
	AssignmentContext
}

func NewAssignmentMinusExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentMinusExprContext {
	var p = new(AssignmentMinusExprContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignmentMinusExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentMinusExprContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *AssignmentMinusExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentMinusExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterAssignmentMinusExpr(s)
	}
}

func (s *AssignmentMinusExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitAssignmentMinusExpr(s)
	}
}

func (s *AssignmentMinusExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitAssignmentMinusExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentDecrementExprContext struct {
	AssignmentContext
}

func NewAssignmentDecrementExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentDecrementExprContext {
	var p = new(AssignmentDecrementExprContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignmentDecrementExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentDecrementExprContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *AssignmentDecrementExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterAssignmentDecrementExpr(s)
	}
}

func (s *AssignmentDecrementExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitAssignmentDecrementExpr(s)
	}
}

func (s *AssignmentDecrementExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitAssignmentDecrementExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentExprContext struct {
	AssignmentContext
}

func NewAssignmentExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentExprContext {
	var p = new(AssignmentExprContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignmentExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentExprContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *AssignmentExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterAssignmentExpr(s)
	}
}

func (s *AssignmentExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitAssignmentExpr(s)
	}
}

func (s *AssignmentExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitAssignmentExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentIncrementExprContext struct {
	AssignmentContext
}

func NewAssignmentIncrementExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentIncrementExprContext {
	var p = new(AssignmentIncrementExprContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignmentIncrementExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentIncrementExprContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *AssignmentIncrementExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterAssignmentIncrementExpr(s)
	}
}

func (s *AssignmentIncrementExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitAssignmentIncrementExpr(s)
	}
}

func (s *AssignmentIncrementExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitAssignmentIncrementExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SliceAssignmentExprContext struct {
	AssignmentContext
}

func NewSliceAssignmentExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceAssignmentExprContext {
	var p = new(SliceAssignmentExprContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *SliceAssignmentExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceAssignmentExprContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *SliceAssignmentExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *SliceAssignmentExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *SliceAssignmentExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterSliceAssignmentExpr(s)
	}
}

func (s *SliceAssignmentExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitSliceAssignmentExpr(s)
	}
}

func (s *SliceAssignmentExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitSliceAssignmentExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentPlusExprContext struct {
	AssignmentContext
}

func NewAssignmentPlusExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentPlusExprContext {
	var p = new(AssignmentPlusExprContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AssignmentPlusExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentPlusExprContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *AssignmentPlusExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentPlusExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterAssignmentPlusExpr(s)
	}
}

func (s *AssignmentPlusExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitAssignmentPlusExpr(s)
	}
}

func (s *AssignmentPlusExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitAssignmentPlusExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AttributeAssignmentExprContext struct {
	AssignmentContext
}

func NewAttributeAssignmentExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AttributeAssignmentExprContext {
	var p = new(AttributeAssignmentExprContext)

	InitEmptyAssignmentContext(&p.AssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentContext))

	return p
}

func (s *AttributeAssignmentExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeAssignmentExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AttributeAssignmentExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *AttributeAssignmentExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterAttributeAssignmentExpr(s)
	}
}

func (s *AttributeAssignmentExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitAttributeAssignmentExpr(s)
	}
}

func (s *AttributeAssignmentExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitAttributeAssignmentExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GramaticaParserRULE_assignment)
	var _la int

	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAssignmentExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(208)
			p.Match(GramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(209)
			p.Match(GramaticaParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(210)
			p.expression(0)
		}
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramaticaParserT__5 {
			{
				p.SetState(211)
				p.Match(GramaticaParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		localctx = NewAssignmentPlusExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(214)
			p.Match(GramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(215)
			p.Match(GramaticaParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)
			p.expression(0)
		}
		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramaticaParserT__5 {
			{
				p.SetState(217)
				p.Match(GramaticaParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 3:
		localctx = NewAssignmentMinusExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(220)
			p.Match(GramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(221)
			p.Match(GramaticaParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(222)
			p.expression(0)
		}
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramaticaParserT__5 {
			{
				p.SetState(223)
				p.Match(GramaticaParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 4:
		localctx = NewAssignmentIncrementExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(226)
			p.Match(GramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(227)
			p.Match(GramaticaParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(229)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramaticaParserT__5 {
			{
				p.SetState(228)
				p.Match(GramaticaParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 5:
		localctx = NewAssignmentDecrementExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(231)
			p.Match(GramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(232)
			p.Match(GramaticaParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramaticaParserT__5 {
			{
				p.SetState(233)
				p.Match(GramaticaParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 6:
		localctx = NewSliceAssignmentExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(236)
			p.Match(GramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(237)
			p.Match(GramaticaParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
			p.expression(0)
		}
		{
			p.SetState(239)
			p.Match(GramaticaParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(240)
			p.Match(GramaticaParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(241)
			p.expression(0)
		}
		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramaticaParserT__5 {
			{
				p.SetState(242)
				p.Match(GramaticaParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 7:
		localctx = NewAttributeAssignmentExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(245)
			p.expression(0)
		}
		{
			p.SetState(246)
			p.Match(GramaticaParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(247)
			p.expression(0)
		}
		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramaticaParserT__5 {
			{
				p.SetState(248)
				p.Match(GramaticaParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
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

// IPrintContext is an interface to support dynamic dispatch.
type IPrintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINT() antlr.TerminalNode
	PrintList() IPrintListContext

	// IsPrintContext differentiates from other interfaces.
	IsPrintContext()
}

type PrintContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintContext() *PrintContext {
	var p = new(PrintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_print
	return p
}

func InitEmptyPrintContext(p *PrintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_print
}

func (*PrintContext) IsPrintContext() {}

func NewPrintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintContext {
	var p = new(PrintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_print

	return p
}

func (s *PrintContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintContext) PRINT() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPRINT, 0)
}

func (s *PrintContext) PrintList() IPrintListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintListContext)
}

func (s *PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterPrint(s)
	}
}

func (s *PrintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitPrint(s)
	}
}

func (s *PrintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitPrint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) Print_() (localctx IPrintContext) {
	localctx = NewPrintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GramaticaParserRULE_print)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(253)
		p.Match(GramaticaParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(254)
		p.Match(GramaticaParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9221120237048954882) != 0 {
		{
			p.SetState(255)
			p.PrintList()
		}

	}
	{
		p.SetState(258)
		p.Match(GramaticaParserT__1)
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
	_la = p.GetTokenStream().LA(1)

	if _la == GramaticaParserT__5 {
		{
			p.SetState(259)
			p.Match(GramaticaParserT__5)
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

// IPrintListContext is an interface to support dynamic dispatch.
type IPrintListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsPrintListContext differentiates from other interfaces.
	IsPrintListContext()
}

type PrintListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintListContext() *PrintListContext {
	var p = new(PrintListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_printList
	return p
}

func InitEmptyPrintListContext(p *PrintListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_printList
}

func (*PrintListContext) IsPrintListContext() {}

func NewPrintListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintListContext {
	var p = new(PrintListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_printList

	return p
}

func (s *PrintListContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *PrintListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *PrintListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterPrintList(s)
	}
}

func (s *PrintListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitPrintList(s)
	}
}

func (s *PrintListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitPrintList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) PrintList() (localctx IPrintListContext) {
	localctx = NewPrintListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GramaticaParserRULE_printList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(262)
		p.expression(0)
	}
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GramaticaParserT__8 {
		{
			p.SetState(263)
			p.Match(GramaticaParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(264)
			p.expression(0)
		}

		p.SetState(269)
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

// IIfContext is an interface to support dynamic dispatch.
type IIfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	Expression() IExpressionContext
	Block() IBlockContext
	AllElseIfBlock() []IElseIfBlockContext
	ElseIfBlock(i int) IElseIfBlockContext
	ElseBlock() IElseBlockContext

	// IsIfContext differentiates from other interfaces.
	IsIfContext()
}

type IfContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfContext() *IfContext {
	var p = new(IfContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_if
	return p
}

func InitEmptyIfContext(p *IfContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_if
}

func (*IfContext) IsIfContext() {}

func NewIfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfContext {
	var p = new(IfContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_if

	return p
}

func (s *IfContext) GetParser() antlr.Parser { return s.parser }

func (s *IfContext) IF() antlr.TerminalNode {
	return s.GetToken(GramaticaParserIF, 0)
}

func (s *IfContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfContext) AllElseIfBlock() []IElseIfBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElseIfBlockContext); ok {
			len++
		}
	}

	tst := make([]IElseIfBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElseIfBlockContext); ok {
			tst[i] = t.(IElseIfBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfContext) ElseIfBlock(i int) IElseIfBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseIfBlockContext); ok {
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

	return t.(IElseIfBlockContext)
}

func (s *IfContext) ElseBlock() IElseBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseBlockContext)
}

func (s *IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterIf(s)
	}
}

func (s *IfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitIf(s)
	}
}

func (s *IfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitIf(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) If_() (localctx IIfContext) {
	localctx = NewIfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GramaticaParserRULE_if)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(270)
		p.Match(GramaticaParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(271)
		p.expression(0)
	}
	{
		p.SetState(272)
		p.Match(GramaticaParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(273)
		p.Block()
	}
	{
		p.SetState(274)
		p.Match(GramaticaParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(275)
				p.ElseIfBlock()
			}

		}
		p.SetState(280)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GramaticaParserELSE {
		{
			p.SetState(281)
			p.ElseBlock()
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

// IElseIfBlockContext is an interface to support dynamic dispatch.
type IElseIfBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ELSE() antlr.TerminalNode
	IF() antlr.TerminalNode
	Expression() IExpressionContext
	Block() IBlockContext

	// IsElseIfBlockContext differentiates from other interfaces.
	IsElseIfBlockContext()
}

type ElseIfBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseIfBlockContext() *ElseIfBlockContext {
	var p = new(ElseIfBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_elseIfBlock
	return p
}

func InitEmptyElseIfBlockContext(p *ElseIfBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_elseIfBlock
}

func (*ElseIfBlockContext) IsElseIfBlockContext() {}

func NewElseIfBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseIfBlockContext {
	var p = new(ElseIfBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_elseIfBlock

	return p
}

func (s *ElseIfBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseIfBlockContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GramaticaParserELSE, 0)
}

func (s *ElseIfBlockContext) IF() antlr.TerminalNode {
	return s.GetToken(GramaticaParserIF, 0)
}

func (s *ElseIfBlockContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ElseIfBlockContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseIfBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseIfBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterElseIfBlock(s)
	}
}

func (s *ElseIfBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitElseIfBlock(s)
	}
}

func (s *ElseIfBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitElseIfBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) ElseIfBlock() (localctx IElseIfBlockContext) {
	localctx = NewElseIfBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GramaticaParserRULE_elseIfBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(284)
		p.Match(GramaticaParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(285)
		p.Match(GramaticaParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(286)
		p.expression(0)
	}
	{
		p.SetState(287)
		p.Match(GramaticaParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(288)
		p.Block()
	}
	{
		p.SetState(289)
		p.Match(GramaticaParserT__3)
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

// IElseBlockContext is an interface to support dynamic dispatch.
type IElseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ELSE() antlr.TerminalNode
	Block() IBlockContext

	// IsElseBlockContext differentiates from other interfaces.
	IsElseBlockContext()
}

type ElseBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseBlockContext() *ElseBlockContext {
	var p = new(ElseBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_elseBlock
	return p
}

func InitEmptyElseBlockContext(p *ElseBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_elseBlock
}

func (*ElseBlockContext) IsElseBlockContext() {}

func NewElseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseBlockContext {
	var p = new(ElseBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_elseBlock

	return p
}

func (s *ElseBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseBlockContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GramaticaParserELSE, 0)
}

func (s *ElseBlockContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterElseBlock(s)
	}
}

func (s *ElseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitElseBlock(s)
	}
}

func (s *ElseBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitElseBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) ElseBlock() (localctx IElseBlockContext) {
	localctx = NewElseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GramaticaParserRULE_elseBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		p.Match(GramaticaParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(292)
		p.Match(GramaticaParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(293)
		p.Block()
	}
	{
		p.SetState(294)
		p.Match(GramaticaParserT__3)
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

// ISwitchContext is an interface to support dynamic dispatch.
type ISwitchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SWTCH() antlr.TerminalNode
	Expression() IExpressionContext
	AllCaseBlock() []ICaseBlockContext
	CaseBlock(i int) ICaseBlockContext
	DefaultBlock() IDefaultBlockContext

	// IsSwitchContext differentiates from other interfaces.
	IsSwitchContext()
}

type SwitchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchContext() *SwitchContext {
	var p = new(SwitchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_switch
	return p
}

func InitEmptySwitchContext(p *SwitchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_switch
}

func (*SwitchContext) IsSwitchContext() {}

func NewSwitchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchContext {
	var p = new(SwitchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_switch

	return p
}

func (s *SwitchContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchContext) SWTCH() antlr.TerminalNode {
	return s.GetToken(GramaticaParserSWTCH, 0)
}

func (s *SwitchContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SwitchContext) AllCaseBlock() []ICaseBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICaseBlockContext); ok {
			len++
		}
	}

	tst := make([]ICaseBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICaseBlockContext); ok {
			tst[i] = t.(ICaseBlockContext)
			i++
		}
	}

	return tst
}

func (s *SwitchContext) CaseBlock(i int) ICaseBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseBlockContext); ok {
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

	return t.(ICaseBlockContext)
}

func (s *SwitchContext) DefaultBlock() IDefaultBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultBlockContext)
}

func (s *SwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterSwitch(s)
	}
}

func (s *SwitchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitSwitch(s)
	}
}

func (s *SwitchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitSwitch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) Switch_() (localctx ISwitchContext) {
	localctx = NewSwitchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GramaticaParserRULE_switch)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(296)
		p.Match(GramaticaParserSWTCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(297)
		p.expression(0)
	}
	{
		p.SetState(298)
		p.Match(GramaticaParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GramaticaParserCASE {
		{
			p.SetState(299)
			p.CaseBlock()
		}

		p.SetState(302)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GramaticaParserT__16 {
		{
			p.SetState(304)
			p.DefaultBlock()
		}

	}
	{
		p.SetState(307)
		p.Match(GramaticaParserT__3)
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

// ICaseBlockContext is an interface to support dynamic dispatch.
type ICaseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASE() antlr.TerminalNode
	Expression() IExpressionContext
	Block() IBlockContext

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
	p.RuleIndex = GramaticaParserRULE_caseBlock
	return p
}

func InitEmptyCaseBlockContext(p *CaseBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_caseBlock
}

func (*CaseBlockContext) IsCaseBlockContext() {}

func NewCaseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseBlockContext {
	var p = new(CaseBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_caseBlock

	return p
}

func (s *CaseBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseBlockContext) CASE() antlr.TerminalNode {
	return s.GetToken(GramaticaParserCASE, 0)
}

func (s *CaseBlockContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CaseBlockContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *CaseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterCaseBlock(s)
	}
}

func (s *CaseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitCaseBlock(s)
	}
}

func (s *CaseBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitCaseBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) CaseBlock() (localctx ICaseBlockContext) {
	localctx = NewCaseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GramaticaParserRULE_caseBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(309)
		p.Match(GramaticaParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(310)
		p.expression(0)
	}
	{
		p.SetState(311)
		p.Match(GramaticaParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(312)
		p.Block()
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

// IDefaultBlockContext is an interface to support dynamic dispatch.
type IDefaultBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext

	// IsDefaultBlockContext differentiates from other interfaces.
	IsDefaultBlockContext()
}

type DefaultBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultBlockContext() *DefaultBlockContext {
	var p = new(DefaultBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_defaultBlock
	return p
}

func InitEmptyDefaultBlockContext(p *DefaultBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_defaultBlock
}

func (*DefaultBlockContext) IsDefaultBlockContext() {}

func NewDefaultBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultBlockContext {
	var p = new(DefaultBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_defaultBlock

	return p
}

func (s *DefaultBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultBlockContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *DefaultBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterDefaultBlock(s)
	}
}

func (s *DefaultBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitDefaultBlock(s)
	}
}

func (s *DefaultBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitDefaultBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) DefaultBlock() (localctx IDefaultBlockContext) {
	localctx = NewDefaultBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GramaticaParserRULE_defaultBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		p.Match(GramaticaParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(315)
		p.Match(GramaticaParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(316)
		p.Block()
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

// IForContext is an interface to support dynamic dispatch.
type IForContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsForContext differentiates from other interfaces.
	IsForContext()
}

type ForContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForContext() *ForContext {
	var p = new(ForContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_for
	return p
}

func InitEmptyForContext(p *ForContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_for
}

func (*ForContext) IsForContext() {}

func NewForContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForContext {
	var p = new(ForContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_for

	return p
}

func (s *ForContext) GetParser() antlr.Parser { return s.parser }

func (s *ForContext) CopyAll(ctx *ForContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForDeclarationContext struct {
	ForContext
}

func NewForDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForDeclarationContext {
	var p = new(ForDeclarationContext)

	InitEmptyForContext(&p.ForContext)
	p.parser = parser
	p.CopyAll(ctx.(*ForContext))

	return p
}

func (s *ForDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForDeclarationContext) FOR() antlr.TerminalNode {
	return s.GetToken(GramaticaParserFOR, 0)
}

func (s *ForDeclarationContext) Declaration() IDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *ForDeclarationContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForDeclarationContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForDeclarationContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *ForDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterForDeclaration(s)
	}
}

func (s *ForDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitForDeclaration(s)
	}
}

func (s *ForDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitForDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForInSliceContext struct {
	ForContext
}

func NewForInSliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForInSliceContext {
	var p = new(ForInSliceContext)

	InitEmptyForContext(&p.ForContext)
	p.parser = parser
	p.CopyAll(ctx.(*ForContext))

	return p
}

func (s *ForInSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInSliceContext) FOR() antlr.TerminalNode {
	return s.GetToken(GramaticaParserFOR, 0)
}

func (s *ForInSliceContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(GramaticaParserID)
}

func (s *ForInSliceContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, i)
}

func (s *ForInSliceContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForInSliceContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForInSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterForInSlice(s)
	}
}

func (s *ForInSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitForInSlice(s)
	}
}

func (s *ForInSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitForInSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForSimpleContext struct {
	ForContext
}

func NewForSimpleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForSimpleContext {
	var p = new(ForSimpleContext)

	InitEmptyForContext(&p.ForContext)
	p.parser = parser
	p.CopyAll(ctx.(*ForContext))

	return p
}

func (s *ForSimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForSimpleContext) FOR() antlr.TerminalNode {
	return s.GetToken(GramaticaParserFOR, 0)
}

func (s *ForSimpleContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForSimpleContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForSimpleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterForSimple(s)
	}
}

func (s *ForSimpleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitForSimple(s)
	}
}

func (s *ForSimpleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitForSimple(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) For_() (localctx IForContext) {
	localctx = NewForContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GramaticaParserRULE_for)
	var _la int

	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForSimpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(318)
			p.Match(GramaticaParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(319)
			p.expression(0)
		}
		{
			p.SetState(320)
			p.Match(GramaticaParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(321)
			p.Block()
		}
		{
			p.SetState(322)
			p.Match(GramaticaParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewForDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(324)
			p.Match(GramaticaParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(325)
			p.Declaration()
		}
		{
			p.SetState(326)
			p.Match(GramaticaParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(327)
			p.expression(0)
		}
		{
			p.SetState(328)
			p.Match(GramaticaParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(330)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9221120237048954882) != 0 {
			{
				p.SetState(329)
				p.Assignment()
			}

		}
		{
			p.SetState(332)
			p.Match(GramaticaParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(333)
			p.Block()
		}
		{
			p.SetState(334)
			p.Match(GramaticaParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewForInSliceContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(336)
			p.Match(GramaticaParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(337)
			p.Match(GramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(338)
			p.Match(GramaticaParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(339)
			p.Match(GramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(340)
			p.Match(GramaticaParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(341)
			p.expression(0)
		}
		{
			p.SetState(342)
			p.Match(GramaticaParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)
			p.Block()
		}
		{
			p.SetState(344)
			p.Match(GramaticaParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
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

// IIndividualBlockContext is an interface to support dynamic dispatch.
type IIndividualBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext

	// IsIndividualBlockContext differentiates from other interfaces.
	IsIndividualBlockContext()
}

type IndividualBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndividualBlockContext() *IndividualBlockContext {
	var p = new(IndividualBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_individualBlock
	return p
}

func InitEmptyIndividualBlockContext(p *IndividualBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_individualBlock
}

func (*IndividualBlockContext) IsIndividualBlockContext() {}

func NewIndividualBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndividualBlockContext {
	var p = new(IndividualBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_individualBlock

	return p
}

func (s *IndividualBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *IndividualBlockContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IndividualBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndividualBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndividualBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterIndividualBlock(s)
	}
}

func (s *IndividualBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitIndividualBlock(s)
	}
}

func (s *IndividualBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitIndividualBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) IndividualBlock() (localctx IIndividualBlockContext) {
	localctx = NewIndividualBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GramaticaParserRULE_individualBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(348)
		p.Match(GramaticaParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(349)
		p.Block()
	}
	{
		p.SetState(350)
		p.Match(GramaticaParserT__3)
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

// ITransferStatemenContext is an interface to support dynamic dispatch.
type ITransferStatemenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTransferStatemenContext differentiates from other interfaces.
	IsTransferStatemenContext()
}

type TransferStatemenContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransferStatemenContext() *TransferStatemenContext {
	var p = new(TransferStatemenContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_transferStatemen
	return p
}

func InitEmptyTransferStatemenContext(p *TransferStatemenContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_transferStatemen
}

func (*TransferStatemenContext) IsTransferStatemenContext() {}

func NewTransferStatemenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransferStatemenContext {
	var p = new(TransferStatemenContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_transferStatemen

	return p
}

func (s *TransferStatemenContext) GetParser() antlr.Parser { return s.parser }

func (s *TransferStatemenContext) CopyAll(ctx *TransferStatemenContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TransferStatemenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransferStatemenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BreakStatementContext struct {
	TransferStatemenContext
}

func NewBreakStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStatementContext {
	var p = new(BreakStatementContext)

	InitEmptyTransferStatemenContext(&p.TransferStatemenContext)
	p.parser = parser
	p.CopyAll(ctx.(*TransferStatemenContext))

	return p
}

func (s *BreakStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatementContext) TRANSFER_STATEMENT_BREAK() antlr.TerminalNode {
	return s.GetToken(GramaticaParserTRANSFER_STATEMENT_BREAK, 0)
}

func (s *BreakStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterBreakStatement(s)
	}
}

func (s *BreakStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitBreakStatement(s)
	}
}

func (s *BreakStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitBreakStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStatementContext struct {
	TransferStatemenContext
}

func NewReturnStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	InitEmptyTransferStatemenContext(&p.TransferStatemenContext)
	p.parser = parser
	p.CopyAll(ctx.(*TransferStatemenContext))

	return p
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) TRANSFER_STATEMENT_RETURN() antlr.TerminalNode {
	return s.GetToken(GramaticaParserTRANSFER_STATEMENT_RETURN, 0)
}

func (s *ReturnStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ContinueStatementContext struct {
	TransferStatemenContext
}

func NewContinueStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueStatementContext {
	var p = new(ContinueStatementContext)

	InitEmptyTransferStatemenContext(&p.TransferStatemenContext)
	p.parser = parser
	p.CopyAll(ctx.(*TransferStatemenContext))

	return p
}

func (s *ContinueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStatementContext) TRANSFER_STATEMENT_CONTINUE() antlr.TerminalNode {
	return s.GetToken(GramaticaParserTRANSFER_STATEMENT_CONTINUE, 0)
}

func (s *ContinueStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterContinueStatement(s)
	}
}

func (s *ContinueStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitContinueStatement(s)
	}
}

func (s *ContinueStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitContinueStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) TransferStatemen() (localctx ITransferStatemenContext) {
	localctx = NewTransferStatemenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GramaticaParserRULE_transferStatemen)
	var _la int

	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserTRANSFER_STATEMENT_BREAK:
		localctx = NewBreakStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(352)
			p.Match(GramaticaParserTRANSFER_STATEMENT_BREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(354)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramaticaParserT__5 {
			{
				p.SetState(353)
				p.Match(GramaticaParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case GramaticaParserTRANSFER_STATEMENT_CONTINUE:
		localctx = NewContinueStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(356)
			p.Match(GramaticaParserTRANSFER_STATEMENT_CONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(358)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramaticaParserT__5 {
			{
				p.SetState(357)
				p.Match(GramaticaParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case GramaticaParserTRANSFER_STATEMENT_RETURN:
		localctx = NewReturnStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(360)
			p.Match(GramaticaParserTRANSFER_STATEMENT_RETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(362)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(361)
				p.expression(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(365)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramaticaParserT__5 {
			{
				p.SetState(364)
				p.Match(GramaticaParserT__5)
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

// ICallContext is an interface to support dynamic dispatch.
type ICallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCallContext differentiates from other interfaces.
	IsCallContext()
}

type CallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallContext() *CallContext {
	var p = new(CallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_call
	return p
}

func InitEmptyCallContext(p *CallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_call
}

func (*CallContext) IsCallContext() {}

func NewCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallContext {
	var p = new(CallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_call

	return p
}

func (s *CallContext) GetParser() antlr.Parser { return s.parser }

func (s *CallContext) CopyAll(ctx *CallContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *CallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctionCallContext struct {
	CallContext
}

func NewFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallContext {
	var p = new(FunctionCallContext)

	InitEmptyCallContext(&p.CallContext)
	p.parser = parser
	p.CopyAll(ctx.(*CallContext))

	return p
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *FunctionCallContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) Call() (localctx ICallContext) {
	localctx = NewCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GramaticaParserRULE_call)
	var _la int

	localctx = NewFunctionCallContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(369)
		p.Match(GramaticaParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(370)
		p.Match(GramaticaParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9221120237048954882) != 0 {
		{
			p.SetState(371)
			p.Arguments()
		}

	}
	{
		p.SetState(374)
		p.Match(GramaticaParserT__1)
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

	if _la == GramaticaParserT__5 {
		{
			p.SetState(375)
			p.Match(GramaticaParserT__5)
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

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_arguments
	return p
}

func InitEmptyArgumentsContext(p *ArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_arguments
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentsContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (s *ArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, GramaticaParserRULE_arguments)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(378)
		p.expression(0)
	}
	p.SetState(383)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GramaticaParserT__8 {
		{
			p.SetState(379)
			p.Match(GramaticaParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(380)
			p.expression(0)
		}

		p.SetState(385)
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

// IPrimitiveTypeContext is an interface to support dynamic dispatch.
type IPrimitiveTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRIMITIVE_INT() antlr.TerminalNode
	PRIMITIVE_FLOAT() antlr.TerminalNode
	PRIMITIVE_BOOL() antlr.TerminalNode
	PRIMITIVE_STRING() antlr.TerminalNode

	// IsPrimitiveTypeContext differentiates from other interfaces.
	IsPrimitiveTypeContext()
}

type PrimitiveTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveTypeContext() *PrimitiveTypeContext {
	var p = new(PrimitiveTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_primitiveType
	return p
}

func InitEmptyPrimitiveTypeContext(p *PrimitiveTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_primitiveType
}

func (*PrimitiveTypeContext) IsPrimitiveTypeContext() {}

func NewPrimitiveTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveTypeContext {
	var p = new(PrimitiveTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_primitiveType

	return p
}

func (s *PrimitiveTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveTypeContext) PRIMITIVE_INT() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPRIMITIVE_INT, 0)
}

func (s *PrimitiveTypeContext) PRIMITIVE_FLOAT() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPRIMITIVE_FLOAT, 0)
}

func (s *PrimitiveTypeContext) PRIMITIVE_BOOL() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPRIMITIVE_BOOL, 0)
}

func (s *PrimitiveTypeContext) PRIMITIVE_STRING() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPRIMITIVE_STRING, 0)
}

func (s *PrimitiveTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterPrimitiveType(s)
	}
}

func (s *PrimitiveTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitPrimitiveType(s)
	}
}

func (s *PrimitiveTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitPrimitiveType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) PrimitiveType() (localctx IPrimitiveTypeContext) {
	localctx = NewPrimitiveTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, GramaticaParserRULE_primitiveType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(386)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2111062325329920) != 0) {
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

// ISlicePrimitiveValueFunctionContext is an interface to support dynamic dispatch.
type ISlicePrimitiveValueFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSlicePrimitiveValueFunctionContext differentiates from other interfaces.
	IsSlicePrimitiveValueFunctionContext()
}

type SlicePrimitiveValueFunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySlicePrimitiveValueFunctionContext() *SlicePrimitiveValueFunctionContext {
	var p = new(SlicePrimitiveValueFunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_slicePrimitiveValueFunction
	return p
}

func InitEmptySlicePrimitiveValueFunctionContext(p *SlicePrimitiveValueFunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_slicePrimitiveValueFunction
}

func (*SlicePrimitiveValueFunctionContext) IsSlicePrimitiveValueFunctionContext() {}

func NewSlicePrimitiveValueFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SlicePrimitiveValueFunctionContext {
	var p = new(SlicePrimitiveValueFunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_slicePrimitiveValueFunction

	return p
}

func (s *SlicePrimitiveValueFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *SlicePrimitiveValueFunctionContext) CopyAll(ctx *SlicePrimitiveValueFunctionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SlicePrimitiveValueFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SlicePrimitiveValueFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type JoinFunctionContext struct {
	SlicePrimitiveValueFunctionContext
}

func NewJoinFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JoinFunctionContext {
	var p = new(JoinFunctionContext)

	InitEmptySlicePrimitiveValueFunctionContext(&p.SlicePrimitiveValueFunctionContext)
	p.parser = parser
	p.CopyAll(ctx.(*SlicePrimitiveValueFunctionContext))

	return p
}

func (s *JoinFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JoinFunctionContext) JOIN() antlr.TerminalNode {
	return s.GetToken(GramaticaParserJOIN, 0)
}

func (s *JoinFunctionContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *JoinFunctionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *JoinFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterJoinFunction(s)
	}
}

func (s *JoinFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitJoinFunction(s)
	}
}

func (s *JoinFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitJoinFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

type LengthFunctionContext struct {
	SlicePrimitiveValueFunctionContext
}

func NewLengthFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LengthFunctionContext {
	var p = new(LengthFunctionContext)

	InitEmptySlicePrimitiveValueFunctionContext(&p.SlicePrimitiveValueFunctionContext)
	p.parser = parser
	p.CopyAll(ctx.(*SlicePrimitiveValueFunctionContext))

	return p
}

func (s *LengthFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthFunctionContext) LENGTH() antlr.TerminalNode {
	return s.GetToken(GramaticaParserLENGTH, 0)
}

func (s *LengthFunctionContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *LengthFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterLengthFunction(s)
	}
}

func (s *LengthFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitLengthFunction(s)
	}
}

func (s *LengthFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitLengthFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

type IndexOfFunctionContext struct {
	SlicePrimitiveValueFunctionContext
}

func NewIndexOfFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexOfFunctionContext {
	var p = new(IndexOfFunctionContext)

	InitEmptySlicePrimitiveValueFunctionContext(&p.SlicePrimitiveValueFunctionContext)
	p.parser = parser
	p.CopyAll(ctx.(*SlicePrimitiveValueFunctionContext))

	return p
}

func (s *IndexOfFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexOfFunctionContext) INDEXOF() antlr.TerminalNode {
	return s.GetToken(GramaticaParserINDEXOF, 0)
}

func (s *IndexOfFunctionContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *IndexOfFunctionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IndexOfFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterIndexOfFunction(s)
	}
}

func (s *IndexOfFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitIndexOfFunction(s)
	}
}

func (s *IndexOfFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitIndexOfFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) SlicePrimitiveValueFunction() (localctx ISlicePrimitiveValueFunctionContext) {
	localctx = NewSlicePrimitiveValueFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, GramaticaParserRULE_slicePrimitiveValueFunction)
	p.SetState(406)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserINDEXOF:
		localctx = NewIndexOfFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(388)
			p.Match(GramaticaParserINDEXOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(389)
			p.Match(GramaticaParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(390)
			p.Match(GramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(391)
			p.Match(GramaticaParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(392)
			p.expression(0)
		}
		{
			p.SetState(393)
			p.Match(GramaticaParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GramaticaParserJOIN:
		localctx = NewJoinFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(395)
			p.Match(GramaticaParserJOIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(396)
			p.Match(GramaticaParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(397)
			p.Match(GramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(398)
			p.Match(GramaticaParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(399)
			p.expression(0)
		}
		{
			p.SetState(400)
			p.Match(GramaticaParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GramaticaParserLENGTH:
		localctx = NewLengthFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(402)
			p.Match(GramaticaParserLENGTH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(403)
			p.Match(GramaticaParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(404)
			p.Match(GramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(405)
			p.Match(GramaticaParserT__1)
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramaticaParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BoolExprContext struct {
	ExpressionContext
}

func NewBoolExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolExprContext {
	var p = new(BoolExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BoolExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExprContext) BOOL() antlr.TerminalNode {
	return s.GetToken(GramaticaParserBOOL, 0)
}

func (s *BoolExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterBoolExpr(s)
	}
}

func (s *BoolExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitBoolExpr(s)
	}
}

func (s *BoolExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitBoolExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatExprContext struct {
	ExpressionContext
}

func NewFloatExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatExprContext {
	var p = new(FloatExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *FloatExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatExprContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(GramaticaParserFLOAT, 0)
}

func (s *FloatExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterFloatExpr(s)
	}
}

func (s *FloatExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitFloatExpr(s)
	}
}

func (s *FloatExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitFloatExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdExprContext struct {
	ExpressionContext
}

func NewIdExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdExprContext {
	var p = new(IdExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IdExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdExprContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *IdExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterIdExpr(s)
	}
}

func (s *IdExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitIdExpr(s)
	}
}

func (s *IdExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitIdExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type DerefExprContext struct {
	ExpressionContext
	right antlr.Token
}

func NewDerefExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DerefExprContext {
	var p = new(DerefExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *DerefExprContext) GetRight() antlr.Token { return s.right }

func (s *DerefExprContext) SetRight(v antlr.Token) { s.right = v }

func (s *DerefExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DerefExprContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *DerefExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterDerefExpr(s)
	}
}

func (s *DerefExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitDerefExpr(s)
	}
}

func (s *DerefExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitDerefExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParseFloatExprContext struct {
	ExpressionContext
}

func NewParseFloatExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParseFloatExprContext {
	var p = new(ParseFloatExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ParseFloatExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseFloatExprContext) PARSEFLOAT() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPARSEFLOAT, 0)
}

func (s *ParseFloatExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParseFloatExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterParseFloatExpr(s)
	}
}

func (s *ParseFloatExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitParseFloatExpr(s)
	}
}

func (s *ParseFloatExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitParseFloatExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NegExprContext struct {
	ExpressionContext
	right IExpressionContext
}

func NewNegExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegExprContext {
	var p = new(NegExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *NegExprContext) GetRight() IExpressionContext { return s.right }

func (s *NegExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *NegExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NegExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterNegExpr(s)
	}
}

func (s *NegExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitNegExpr(s)
	}
}

func (s *NegExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitNegExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AppendExprContext struct {
	ExpressionContext
}

func NewAppendExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AppendExprContext {
	var p = new(AppendExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AppendExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppendExprContext) APPEND() antlr.TerminalNode {
	return s.GetToken(GramaticaParserAPPEND, 0)
}

func (s *AppendExprContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *AppendExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AppendExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterAppendExpr(s)
	}
}

func (s *AppendExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitAppendExpr(s)
	}
}

func (s *AppendExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitAppendExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type OpExprContext struct {
	ExpressionContext
	left  IExpressionContext
	op    antlr.Token
	right IExpressionContext
}

func NewOpExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OpExprContext {
	var p = new(OpExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *OpExprContext) GetOp() antlr.Token { return s.op }

func (s *OpExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *OpExprContext) GetLeft() IExpressionContext { return s.left }

func (s *OpExprContext) GetRight() IExpressionContext { return s.right }

func (s *OpExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *OpExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *OpExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *OpExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *OpExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterOpExpr(s)
	}
}

func (s *OpExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitOpExpr(s)
	}
}

func (s *OpExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitOpExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SlicePrimitiveValueContext struct {
	ExpressionContext
}

func NewSlicePrimitiveValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SlicePrimitiveValueContext {
	var p = new(SlicePrimitiveValueContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *SlicePrimitiveValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SlicePrimitiveValueContext) SlicePrimitiveValueFunction() ISlicePrimitiveValueFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISlicePrimitiveValueFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISlicePrimitiveValueFunctionContext)
}

func (s *SlicePrimitiveValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterSlicePrimitiveValue(s)
	}
}

func (s *SlicePrimitiveValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitSlicePrimitiveValue(s)
	}
}

func (s *SlicePrimitiveValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitSlicePrimitiveValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeOfExprContext struct {
	ExpressionContext
}

func NewTypeOfExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeOfExprContext {
	var p = new(TypeOfExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *TypeOfExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeOfExprContext) TYPEOF() antlr.TerminalNode {
	return s.GetToken(GramaticaParserTYPEOF, 0)
}

func (s *TypeOfExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TypeOfExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterTypeOfExpr(s)
	}
}

func (s *TypeOfExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitTypeOfExpr(s)
	}
}

func (s *TypeOfExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitTypeOfExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructAccessContext struct {
	ExpressionContext
}

func NewStructAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructAccessContext {
	var p = new(StructAccessContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *StructAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructAccessContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StructAccessContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *StructAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterStructAccess(s)
	}
}

func (s *StructAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitStructAccess(s)
	}
}

func (s *StructAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitStructAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallExprContext struct {
	ExpressionContext
}

func NewFunctionCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallExprContext {
	var p = new(FunctionCallExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *FunctionCallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallExprContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *FunctionCallExprContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *FunctionCallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterFunctionCallExpr(s)
	}
}

func (s *FunctionCallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitFunctionCallExpr(s)
	}
}

func (s *FunctionCallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitFunctionCallExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AtoiExprContext struct {
	ExpressionContext
}

func NewAtoiExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtoiExprContext {
	var p = new(AtoiExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AtoiExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtoiExprContext) ATOI() antlr.TerminalNode {
	return s.GetToken(GramaticaParserATOI, 0)
}

func (s *AtoiExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AtoiExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterAtoiExpr(s)
	}
}

func (s *AtoiExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitAtoiExpr(s)
	}
}

func (s *AtoiExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitAtoiExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParExprContext struct {
	ExpressionContext
}

func NewParExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParExprContext {
	var p = new(ParExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ParExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterParExpr(s)
	}
}

func (s *ParExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitParExpr(s)
	}
}

func (s *ParExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitParExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type StrExprContext struct {
	ExpressionContext
}

func NewStrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StrExprContext {
	var p = new(StrExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *StrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(GramaticaParserSTRING, 0)
}

func (s *StrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterStrExpr(s)
	}
}

func (s *StrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitStrExpr(s)
	}
}

func (s *StrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitStrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExprContext struct {
	ExpressionContext
	right IExpressionContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *NotExprContext) GetRight() IExpressionContext { return s.right }

func (s *NotExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntExprContext struct {
	ExpressionContext
}

func NewIntExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntExprContext {
	var p = new(IntExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IntExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntExprContext) INT() antlr.TerminalNode {
	return s.GetToken(GramaticaParserINT, 0)
}

func (s *IntExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterIntExpr(s)
	}
}

func (s *IntExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitIntExpr(s)
	}
}

func (s *IntExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitIntExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type RefExprContext struct {
	ExpressionContext
	right antlr.Token
}

func NewRefExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RefExprContext {
	var p = new(RefExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *RefExprContext) GetRight() antlr.Token { return s.right }

func (s *RefExprContext) SetRight(v antlr.Token) { s.right = v }

func (s *RefExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RefExprContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *RefExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterRefExpr(s)
	}
}

func (s *RefExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitRefExpr(s)
	}
}

func (s *RefExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitRefExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SliceAccessContext struct {
	ExpressionContext
}

func NewSliceAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceAccessContext {
	var p = new(SliceAccessContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *SliceAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceAccessContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *SliceAccessContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SliceAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterSliceAccess(s)
	}
}

func (s *SliceAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitSliceAccess(s)
	}
}

func (s *SliceAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramaticaVisitor:
		return t.VisitSliceAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramaticaParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *GramaticaParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 54
	p.EnterRecursionRule(localctx, 54, GramaticaParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(460)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(409)
			p.Match(GramaticaParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(410)

			var _x = p.expression(25)

			localctx.(*NotExprContext).right = _x
		}

	case 2:
		localctx = NewNegExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(411)
			p.Match(GramaticaParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(412)

			var _x = p.expression(24)

			localctx.(*NegExprContext).right = _x
		}

	case 3:
		localctx = NewRefExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(413)
			p.Match(GramaticaParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(414)

			var _m = p.Match(GramaticaParserID)

			localctx.(*RefExprContext).right = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewDerefExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(415)
			p.Match(GramaticaParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(416)

			var _m = p.Match(GramaticaParserID)

			localctx.(*DerefExprContext).right = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewParExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(417)
			p.Match(GramaticaParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(418)
			p.expression(0)
		}
		{
			p.SetState(419)
			p.Match(GramaticaParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewSlicePrimitiveValueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(421)
			p.SlicePrimitiveValueFunction()
		}

	case 7:
		localctx = NewAppendExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(422)
			p.Match(GramaticaParserAPPEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(423)
			p.Match(GramaticaParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(424)
			p.Match(GramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(425)
			p.Match(GramaticaParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(426)
			p.expression(0)
		}
		{
			p.SetState(427)
			p.Match(GramaticaParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewSliceAccessContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(429)
			p.Match(GramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(430)
			p.Match(GramaticaParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(431)
			p.expression(0)
		}
		{
			p.SetState(432)
			p.Match(GramaticaParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewFunctionCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(434)
			p.Match(GramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(435)
			p.Match(GramaticaParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(437)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9221120237048954882) != 0 {
			{
				p.SetState(436)
				p.Arguments()
			}

		}
		{
			p.SetState(439)
			p.Match(GramaticaParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewAtoiExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(440)
			p.Match(GramaticaParserATOI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(441)
			p.Match(GramaticaParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(442)
			p.expression(0)
		}
		{
			p.SetState(443)
			p.Match(GramaticaParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		localctx = NewParseFloatExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(445)
			p.Match(GramaticaParserPARSEFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(446)
			p.Match(GramaticaParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(447)
			p.expression(0)
		}
		{
			p.SetState(448)
			p.Match(GramaticaParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		localctx = NewTypeOfExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(450)
			p.Match(GramaticaParserTYPEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(451)
			p.Match(GramaticaParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(452)
			p.expression(0)
		}
		{
			p.SetState(453)
			p.Match(GramaticaParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		localctx = NewIdExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(455)
			p.Match(GramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		localctx = NewIntExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(456)
			p.Match(GramaticaParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 15:
		localctx = NewFloatExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(457)
			p.Match(GramaticaParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 16:
		localctx = NewBoolExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(458)
			p.Match(GramaticaParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 17:
		localctx = NewStrExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(459)
			p.Match(GramaticaParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(488)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 51, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(486)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 50, p.GetParserRuleContext()) {
			case 1:
				localctx = NewOpExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expression)
				p.SetState(462)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(463)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&54525952) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(464)

					var _x = p.expression(21)

					localctx.(*OpExprContext).right = _x
				}

			case 2:
				localctx = NewOpExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expression)
				p.SetState(465)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(466)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramaticaParserT__19 || _la == GramaticaParserT__25) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(467)

					var _x = p.expression(20)

					localctx.(*OpExprContext).right = _x
				}

			case 3:
				localctx = NewOpExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expression)
				p.SetState(468)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(469)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramaticaParserT__26 || _la == GramaticaParserT__27) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(470)

					var _x = p.expression(19)

					localctx.(*OpExprContext).right = _x
				}

			case 4:
				localctx = NewOpExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expression)
				p.SetState(471)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(472)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramaticaParserT__28 || _la == GramaticaParserT__29) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(473)

					var _x = p.expression(18)

					localctx.(*OpExprContext).right = _x
				}

			case 5:
				localctx = NewOpExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expression)
				p.SetState(474)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(475)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramaticaParserT__30 || _la == GramaticaParserT__31) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(476)

					var _x = p.expression(17)

					localctx.(*OpExprContext).right = _x
				}

			case 6:
				localctx = NewOpExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expression)
				p.SetState(477)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(478)

					var _m = p.Match(GramaticaParserT__32)

					localctx.(*OpExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(479)

					var _x = p.expression(16)

					localctx.(*OpExprContext).right = _x
				}

			case 7:
				localctx = NewOpExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expression)
				p.SetState(480)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(481)

					var _m = p.Match(GramaticaParserT__33)

					localctx.(*OpExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(482)

					var _x = p.expression(15)

					localctx.(*OpExprContext).right = _x
				}

			case 8:
				localctx = NewStructAccessContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expression)
				p.SetState(483)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
					goto errorExit
				}
				{
					p.SetState(484)
					p.Match(GramaticaParserT__22)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(485)
					p.Match(GramaticaParserID)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(490)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 51, p.GetParserRuleContext())
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

func (p *GramaticaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 27:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GramaticaParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 21)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
