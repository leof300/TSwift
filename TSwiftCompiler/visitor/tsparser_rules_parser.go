// Code generated from TSParser_rules.g4 by ANTLR 4.13.1. DO NOT EDIT.

package TSVisitor // TSParser_rules
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

type TSParser_rulesParser struct {
	*antlr.BaseParser
}

var TSParser_rulesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tsparser_rulesParserInit() {
	staticData := &TSParser_rulesParserStaticData
	staticData.LiteralNames = []string{
		"", "'-'", "'!'", "'%'", "'*'", "'/'", "'+'", "'=='", "'!='", "'>'",
		"'<'", "'>='", "'<='", "'('", "')'", "'&&'", "'||'", "'='", "'+='",
		"'-='", "':'", "'['", "']'", "'?'", "','", "'{'", "'}'", "'..'", "'->'",
		"'_'", "'&'", "", "", "'var'", "'let'", "'nil'", "'String'", "'Int'",
		"'Float'", "'Bool'", "'Character'", "'if'", "'else'", "'switch'", "'default'",
		"'case'", "'while'", "'in'", "'for'", "'guard'", "'continue'", "'break'",
		"'return'", "'print'", "'struct'", "'self'", "'mutating'", "'func'",
		"'inout'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "NL", "WS",
		"VAR", "LET", "NIL", "STRING", "INT", "FLOAT", "BOOL", "CHARACTER",
		"IF", "ELSE", "SWITCH", "DEFAULT", "CASE", "WHILE", "IN", "FOR", "GUARD",
		"CONTINUE", "BREAK", "RETURN", "PRINT", "STRUCT", "SELF", "MUTATING",
		"FUNC", "INOUT", "VBOOL", "VSTRING", "VFLOAT", "VINTEGER", "ID", "SL_COMMENT",
		"ML_COMMENT",
	}
	staticData.RuleNames = []string{
		"start", "lins", "ins", "lsents", "sents", "expr", "declar", "arrayValue",
		"decons", "tstypes", "block", "if", "switch", "default", "while", "guard",
		"for", "rango", "strans", "tsprint", "functions", "parameter", "callFunction",
		"callParameter", "tsfunctions",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 65, 358, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0, 1, 0,
		1, 1, 5, 1, 55, 8, 1, 10, 1, 12, 1, 58, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 3, 2, 65, 8, 2, 1, 3, 5, 3, 68, 8, 3, 10, 3, 12, 3, 71, 9, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 83, 8, 4, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 103, 8, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5,
		5, 132, 8, 5, 10, 5, 12, 5, 135, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 3, 6, 145, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 161, 8, 6, 1, 6, 1,
		6, 1, 6, 3, 6, 166, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 172, 8, 7, 10,
		7, 12, 7, 175, 9, 7, 3, 7, 177, 8, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 192, 8, 8, 1, 9, 1,
		9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3,
		11, 216, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		4, 12, 226, 8, 12, 11, 12, 12, 12, 227, 1, 12, 3, 12, 231, 8, 12, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 253,
		8, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1,
		18, 3, 18, 265, 8, 18, 3, 18, 267, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 5, 19, 274, 8, 19, 10, 19, 12, 19, 277, 9, 19, 1, 19, 1, 19, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 287, 8, 20, 10, 20, 12, 20, 290,
		9, 20, 3, 20, 292, 8, 20, 1, 20, 1, 20, 1, 20, 3, 20, 297, 8, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 21, 3, 21, 304, 8, 21, 1, 21, 1, 21, 1, 21, 3,
		21, 309, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 316, 8, 21, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 323, 8, 22, 10, 22, 12, 22, 326,
		9, 22, 3, 22, 328, 8, 22, 1, 22, 1, 22, 1, 23, 1, 23, 3, 23, 334, 8, 23,
		1, 23, 3, 23, 337, 8, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		3, 24, 356, 8, 24, 1, 24, 0, 1, 10, 25, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 0, 5, 1,
		0, 4, 5, 2, 0, 1, 1, 6, 6, 1, 0, 7, 12, 1, 0, 36, 40, 2, 0, 29, 29, 63,
		63, 395, 0, 50, 1, 0, 0, 0, 2, 56, 1, 0, 0, 0, 4, 64, 1, 0, 0, 0, 6, 69,
		1, 0, 0, 0, 8, 82, 1, 0, 0, 0, 10, 102, 1, 0, 0, 0, 12, 165, 1, 0, 0, 0,
		14, 167, 1, 0, 0, 0, 16, 191, 1, 0, 0, 0, 18, 193, 1, 0, 0, 0, 20, 195,
		1, 0, 0, 0, 22, 215, 1, 0, 0, 0, 24, 217, 1, 0, 0, 0, 26, 234, 1, 0, 0,
		0, 28, 238, 1, 0, 0, 0, 30, 242, 1, 0, 0, 0, 32, 247, 1, 0, 0, 0, 34, 256,
		1, 0, 0, 0, 36, 266, 1, 0, 0, 0, 38, 268, 1, 0, 0, 0, 40, 280, 1, 0, 0,
		0, 42, 303, 1, 0, 0, 0, 44, 317, 1, 0, 0, 0, 46, 333, 1, 0, 0, 0, 48, 355,
		1, 0, 0, 0, 50, 51, 3, 2, 1, 0, 51, 52, 5, 0, 0, 1, 52, 1, 1, 0, 0, 0,
		53, 55, 3, 4, 2, 0, 54, 53, 1, 0, 0, 0, 55, 58, 1, 0, 0, 0, 56, 54, 1,
		0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 3, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 59,
		65, 3, 40, 20, 0, 60, 65, 3, 12, 6, 0, 61, 65, 3, 16, 8, 0, 62, 65, 3,
		44, 22, 0, 63, 65, 3, 38, 19, 0, 64, 59, 1, 0, 0, 0, 64, 60, 1, 0, 0, 0,
		64, 61, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 64, 63, 1, 0, 0, 0, 65, 5, 1, 0,
		0, 0, 66, 68, 3, 8, 4, 0, 67, 66, 1, 0, 0, 0, 68, 71, 1, 0, 0, 0, 69, 67,
		1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 7, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0,
		72, 83, 3, 22, 11, 0, 73, 83, 3, 24, 12, 0, 74, 83, 3, 28, 14, 0, 75, 83,
		3, 30, 15, 0, 76, 83, 3, 32, 16, 0, 77, 83, 3, 12, 6, 0, 78, 83, 3, 16,
		8, 0, 79, 83, 3, 38, 19, 0, 80, 83, 3, 10, 5, 0, 81, 83, 3, 36, 18, 0,
		82, 72, 1, 0, 0, 0, 82, 73, 1, 0, 0, 0, 82, 74, 1, 0, 0, 0, 82, 75, 1,
		0, 0, 0, 82, 76, 1, 0, 0, 0, 82, 77, 1, 0, 0, 0, 82, 78, 1, 0, 0, 0, 82,
		79, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 82, 81, 1, 0, 0, 0, 83, 9, 1, 0, 0,
		0, 84, 85, 6, 5, -1, 0, 85, 86, 5, 1, 0, 0, 86, 103, 3, 10, 5, 21, 87,
		88, 5, 2, 0, 0, 88, 103, 3, 10, 5, 20, 89, 103, 3, 48, 24, 0, 90, 103,
		3, 44, 22, 0, 91, 103, 3, 14, 7, 0, 92, 93, 5, 13, 0, 0, 93, 94, 3, 10,
		5, 0, 94, 95, 5, 14, 0, 0, 95, 103, 1, 0, 0, 0, 96, 103, 5, 60, 0, 0, 97,
		103, 5, 62, 0, 0, 98, 103, 5, 61, 0, 0, 99, 103, 5, 59, 0, 0, 100, 103,
		5, 63, 0, 0, 101, 103, 5, 35, 0, 0, 102, 84, 1, 0, 0, 0, 102, 87, 1, 0,
		0, 0, 102, 89, 1, 0, 0, 0, 102, 90, 1, 0, 0, 0, 102, 91, 1, 0, 0, 0, 102,
		92, 1, 0, 0, 0, 102, 96, 1, 0, 0, 0, 102, 97, 1, 0, 0, 0, 102, 98, 1, 0,
		0, 0, 102, 99, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 102, 101, 1, 0, 0, 0,
		103, 133, 1, 0, 0, 0, 104, 105, 10, 19, 0, 0, 105, 106, 5, 3, 0, 0, 106,
		132, 3, 10, 5, 20, 107, 108, 10, 18, 0, 0, 108, 109, 7, 0, 0, 0, 109, 132,
		3, 10, 5, 19, 110, 111, 10, 17, 0, 0, 111, 112, 7, 1, 0, 0, 112, 132, 3,
		10, 5, 18, 113, 114, 10, 16, 0, 0, 114, 115, 7, 2, 0, 0, 115, 132, 3, 10,
		5, 17, 116, 117, 10, 11, 0, 0, 117, 118, 5, 15, 0, 0, 118, 132, 3, 10,
		5, 12, 119, 120, 10, 10, 0, 0, 120, 121, 5, 16, 0, 0, 121, 132, 3, 10,
		5, 11, 122, 123, 10, 9, 0, 0, 123, 124, 5, 17, 0, 0, 124, 132, 3, 10, 5,
		9, 125, 126, 10, 8, 0, 0, 126, 127, 5, 18, 0, 0, 127, 132, 3, 10, 5, 8,
		128, 129, 10, 7, 0, 0, 129, 130, 5, 19, 0, 0, 130, 132, 3, 10, 5, 7, 131,
		104, 1, 0, 0, 0, 131, 107, 1, 0, 0, 0, 131, 110, 1, 0, 0, 0, 131, 113,
		1, 0, 0, 0, 131, 116, 1, 0, 0, 0, 131, 119, 1, 0, 0, 0, 131, 122, 1, 0,
		0, 0, 131, 125, 1, 0, 0, 0, 131, 128, 1, 0, 0, 0, 132, 135, 1, 0, 0, 0,
		133, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 11, 1, 0, 0, 0, 135, 133,
		1, 0, 0, 0, 136, 137, 5, 33, 0, 0, 137, 138, 5, 63, 0, 0, 138, 144, 5,
		20, 0, 0, 139, 145, 3, 18, 9, 0, 140, 141, 5, 21, 0, 0, 141, 142, 3, 18,
		9, 0, 142, 143, 5, 22, 0, 0, 143, 145, 1, 0, 0, 0, 144, 139, 1, 0, 0, 0,
		144, 140, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 147, 5, 23, 0, 0, 147,
		166, 1, 0, 0, 0, 148, 149, 5, 33, 0, 0, 149, 150, 5, 63, 0, 0, 150, 151,
		5, 17, 0, 0, 151, 166, 3, 10, 5, 0, 152, 153, 5, 33, 0, 0, 153, 154, 5,
		63, 0, 0, 154, 160, 5, 20, 0, 0, 155, 161, 3, 18, 9, 0, 156, 157, 5, 21,
		0, 0, 157, 158, 3, 18, 9, 0, 158, 159, 5, 22, 0, 0, 159, 161, 1, 0, 0,
		0, 160, 155, 1, 0, 0, 0, 160, 156, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162,
		163, 5, 17, 0, 0, 163, 164, 3, 10, 5, 0, 164, 166, 1, 0, 0, 0, 165, 136,
		1, 0, 0, 0, 165, 148, 1, 0, 0, 0, 165, 152, 1, 0, 0, 0, 166, 13, 1, 0,
		0, 0, 167, 176, 5, 21, 0, 0, 168, 173, 3, 10, 5, 0, 169, 170, 5, 24, 0,
		0, 170, 172, 3, 10, 5, 0, 171, 169, 1, 0, 0, 0, 172, 175, 1, 0, 0, 0, 173,
		171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 177, 1, 0, 0, 0, 175, 173,
		1, 0, 0, 0, 176, 168, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 178, 1, 0,
		0, 0, 178, 179, 5, 22, 0, 0, 179, 15, 1, 0, 0, 0, 180, 181, 5, 34, 0, 0,
		181, 182, 5, 63, 0, 0, 182, 183, 5, 17, 0, 0, 183, 192, 3, 10, 5, 0, 184,
		185, 5, 34, 0, 0, 185, 186, 5, 63, 0, 0, 186, 187, 5, 20, 0, 0, 187, 188,
		3, 18, 9, 0, 188, 189, 5, 17, 0, 0, 189, 190, 3, 10, 5, 0, 190, 192, 1,
		0, 0, 0, 191, 180, 1, 0, 0, 0, 191, 184, 1, 0, 0, 0, 192, 17, 1, 0, 0,
		0, 193, 194, 7, 3, 0, 0, 194, 19, 1, 0, 0, 0, 195, 196, 5, 25, 0, 0, 196,
		197, 3, 6, 3, 0, 197, 198, 5, 26, 0, 0, 198, 21, 1, 0, 0, 0, 199, 200,
		5, 41, 0, 0, 200, 201, 3, 10, 5, 0, 201, 202, 3, 20, 10, 0, 202, 216, 1,
		0, 0, 0, 203, 204, 5, 41, 0, 0, 204, 205, 3, 10, 5, 0, 205, 206, 3, 20,
		10, 0, 206, 207, 5, 42, 0, 0, 207, 208, 3, 20, 10, 0, 208, 216, 1, 0, 0,
		0, 209, 210, 5, 41, 0, 0, 210, 211, 3, 10, 5, 0, 211, 212, 3, 20, 10, 0,
		212, 213, 5, 42, 0, 0, 213, 214, 3, 22, 11, 0, 214, 216, 1, 0, 0, 0, 215,
		199, 1, 0, 0, 0, 215, 203, 1, 0, 0, 0, 215, 209, 1, 0, 0, 0, 216, 23, 1,
		0, 0, 0, 217, 218, 5, 43, 0, 0, 218, 219, 3, 10, 5, 0, 219, 225, 5, 25,
		0, 0, 220, 221, 5, 45, 0, 0, 221, 222, 3, 10, 5, 0, 222, 223, 5, 20, 0,
		0, 223, 224, 3, 6, 3, 0, 224, 226, 1, 0, 0, 0, 225, 220, 1, 0, 0, 0, 226,
		227, 1, 0, 0, 0, 227, 225, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 230,
		1, 0, 0, 0, 229, 231, 3, 26, 13, 0, 230, 229, 1, 0, 0, 0, 230, 231, 1,
		0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 233, 5, 26, 0, 0, 233, 25, 1, 0, 0,
		0, 234, 235, 5, 44, 0, 0, 235, 236, 5, 20, 0, 0, 236, 237, 3, 6, 3, 0,
		237, 27, 1, 0, 0, 0, 238, 239, 5, 46, 0, 0, 239, 240, 3, 10, 5, 0, 240,
		241, 3, 20, 10, 0, 241, 29, 1, 0, 0, 0, 242, 243, 5, 49, 0, 0, 243, 244,
		3, 10, 5, 0, 244, 245, 5, 42, 0, 0, 245, 246, 3, 20, 10, 0, 246, 31, 1,
		0, 0, 0, 247, 248, 5, 48, 0, 0, 248, 249, 5, 63, 0, 0, 249, 252, 5, 47,
		0, 0, 250, 253, 3, 10, 5, 0, 251, 253, 3, 34, 17, 0, 252, 250, 1, 0, 0,
		0, 252, 251, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 255, 3, 20, 10, 0,
		255, 33, 1, 0, 0, 0, 256, 257, 5, 62, 0, 0, 257, 258, 5, 27, 0, 0, 258,
		259, 5, 62, 0, 0, 259, 35, 1, 0, 0, 0, 260, 267, 5, 50, 0, 0, 261, 267,
		5, 51, 0, 0, 262, 264, 5, 52, 0, 0, 263, 265, 3, 10, 5, 0, 264, 263, 1,
		0, 0, 0, 264, 265, 1, 0, 0, 0, 265, 267, 1, 0, 0, 0, 266, 260, 1, 0, 0,
		0, 266, 261, 1, 0, 0, 0, 266, 262, 1, 0, 0, 0, 267, 37, 1, 0, 0, 0, 268,
		269, 5, 53, 0, 0, 269, 270, 5, 13, 0, 0, 270, 275, 3, 10, 5, 0, 271, 272,
		5, 24, 0, 0, 272, 274, 3, 10, 5, 0, 273, 271, 1, 0, 0, 0, 274, 277, 1,
		0, 0, 0, 275, 273, 1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 278, 1, 0, 0,
		0, 277, 275, 1, 0, 0, 0, 278, 279, 5, 14, 0, 0, 279, 39, 1, 0, 0, 0, 280,
		281, 5, 57, 0, 0, 281, 282, 5, 63, 0, 0, 282, 291, 5, 13, 0, 0, 283, 288,
		3, 42, 21, 0, 284, 285, 5, 24, 0, 0, 285, 287, 3, 42, 21, 0, 286, 284,
		1, 0, 0, 0, 287, 290, 1, 0, 0, 0, 288, 286, 1, 0, 0, 0, 288, 289, 1, 0,
		0, 0, 289, 292, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 291, 283, 1, 0, 0, 0,
		291, 292, 1, 0, 0, 0, 292, 293, 1, 0, 0, 0, 293, 296, 5, 14, 0, 0, 294,
		295, 5, 28, 0, 0, 295, 297, 3, 18, 9, 0, 296, 294, 1, 0, 0, 0, 296, 297,
		1, 0, 0, 0, 297, 298, 1, 0, 0, 0, 298, 299, 5, 25, 0, 0, 299, 300, 3, 6,
		3, 0, 300, 301, 5, 26, 0, 0, 301, 41, 1, 0, 0, 0, 302, 304, 7, 4, 0, 0,
		303, 302, 1, 0, 0, 0, 303, 304, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305,
		306, 5, 63, 0, 0, 306, 308, 5, 20, 0, 0, 307, 309, 5, 58, 0, 0, 308, 307,
		1, 0, 0, 0, 308, 309, 1, 0, 0, 0, 309, 315, 1, 0, 0, 0, 310, 316, 3, 18,
		9, 0, 311, 312, 5, 21, 0, 0, 312, 313, 3, 18, 9, 0, 313, 314, 5, 22, 0,
		0, 314, 316, 1, 0, 0, 0, 315, 310, 1, 0, 0, 0, 315, 311, 1, 0, 0, 0, 316,
		43, 1, 0, 0, 0, 317, 318, 5, 63, 0, 0, 318, 327, 5, 13, 0, 0, 319, 324,
		3, 46, 23, 0, 320, 321, 5, 24, 0, 0, 321, 323, 3, 46, 23, 0, 322, 320,
		1, 0, 0, 0, 323, 326, 1, 0, 0, 0, 324, 322, 1, 0, 0, 0, 324, 325, 1, 0,
		0, 0, 325, 328, 1, 0, 0, 0, 326, 324, 1, 0, 0, 0, 327, 319, 1, 0, 0, 0,
		327, 328, 1, 0, 0, 0, 328, 329, 1, 0, 0, 0, 329, 330, 5, 14, 0, 0, 330,
		45, 1, 0, 0, 0, 331, 332, 5, 63, 0, 0, 332, 334, 5, 20, 0, 0, 333, 331,
		1, 0, 0, 0, 333, 334, 1, 0, 0, 0, 334, 336, 1, 0, 0, 0, 335, 337, 5, 30,
		0, 0, 336, 335, 1, 0, 0, 0, 336, 337, 1, 0, 0, 0, 337, 338, 1, 0, 0, 0,
		338, 339, 3, 10, 5, 0, 339, 47, 1, 0, 0, 0, 340, 341, 5, 36, 0, 0, 341,
		342, 5, 13, 0, 0, 342, 343, 3, 10, 5, 0, 343, 344, 5, 14, 0, 0, 344, 356,
		1, 0, 0, 0, 345, 346, 5, 37, 0, 0, 346, 347, 5, 13, 0, 0, 347, 348, 3,
		10, 5, 0, 348, 349, 5, 14, 0, 0, 349, 356, 1, 0, 0, 0, 350, 351, 5, 38,
		0, 0, 351, 352, 5, 13, 0, 0, 352, 353, 3, 10, 5, 0, 353, 354, 5, 14, 0,
		0, 354, 356, 1, 0, 0, 0, 355, 340, 1, 0, 0, 0, 355, 345, 1, 0, 0, 0, 355,
		350, 1, 0, 0, 0, 356, 49, 1, 0, 0, 0, 31, 56, 64, 69, 82, 102, 131, 133,
		144, 160, 165, 173, 176, 191, 215, 227, 230, 252, 264, 266, 275, 288, 291,
		296, 303, 308, 315, 324, 327, 333, 336, 355,
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

// TSParser_rulesParserInit initializes any static state used to implement TSParser_rulesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTSParser_rulesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TSParser_rulesParserInit() {
	staticData := &TSParser_rulesParserStaticData
	staticData.once.Do(tsparser_rulesParserInit)
}

// NewTSParser_rulesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTSParser_rulesParser(input antlr.TokenStream) *TSParser_rulesParser {
	TSParser_rulesParserInit()
	this := new(TSParser_rulesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &TSParser_rulesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "TSParser_rules.g4"

	return this
}

// TSParser_rulesParser tokens.
const (
	TSParser_rulesParserEOF        = antlr.TokenEOF
	TSParser_rulesParserT__0       = 1
	TSParser_rulesParserT__1       = 2
	TSParser_rulesParserT__2       = 3
	TSParser_rulesParserT__3       = 4
	TSParser_rulesParserT__4       = 5
	TSParser_rulesParserT__5       = 6
	TSParser_rulesParserT__6       = 7
	TSParser_rulesParserT__7       = 8
	TSParser_rulesParserT__8       = 9
	TSParser_rulesParserT__9       = 10
	TSParser_rulesParserT__10      = 11
	TSParser_rulesParserT__11      = 12
	TSParser_rulesParserT__12      = 13
	TSParser_rulesParserT__13      = 14
	TSParser_rulesParserT__14      = 15
	TSParser_rulesParserT__15      = 16
	TSParser_rulesParserT__16      = 17
	TSParser_rulesParserT__17      = 18
	TSParser_rulesParserT__18      = 19
	TSParser_rulesParserT__19      = 20
	TSParser_rulesParserT__20      = 21
	TSParser_rulesParserT__21      = 22
	TSParser_rulesParserT__22      = 23
	TSParser_rulesParserT__23      = 24
	TSParser_rulesParserT__24      = 25
	TSParser_rulesParserT__25      = 26
	TSParser_rulesParserT__26      = 27
	TSParser_rulesParserT__27      = 28
	TSParser_rulesParserT__28      = 29
	TSParser_rulesParserT__29      = 30
	TSParser_rulesParserNL         = 31
	TSParser_rulesParserWS         = 32
	TSParser_rulesParserVAR        = 33
	TSParser_rulesParserLET        = 34
	TSParser_rulesParserNIL        = 35
	TSParser_rulesParserSTRING     = 36
	TSParser_rulesParserINT        = 37
	TSParser_rulesParserFLOAT      = 38
	TSParser_rulesParserBOOL       = 39
	TSParser_rulesParserCHARACTER  = 40
	TSParser_rulesParserIF         = 41
	TSParser_rulesParserELSE       = 42
	TSParser_rulesParserSWITCH     = 43
	TSParser_rulesParserDEFAULT    = 44
	TSParser_rulesParserCASE       = 45
	TSParser_rulesParserWHILE      = 46
	TSParser_rulesParserIN         = 47
	TSParser_rulesParserFOR        = 48
	TSParser_rulesParserGUARD      = 49
	TSParser_rulesParserCONTINUE   = 50
	TSParser_rulesParserBREAK      = 51
	TSParser_rulesParserRETURN     = 52
	TSParser_rulesParserPRINT      = 53
	TSParser_rulesParserSTRUCT     = 54
	TSParser_rulesParserSELF       = 55
	TSParser_rulesParserMUTATING   = 56
	TSParser_rulesParserFUNC       = 57
	TSParser_rulesParserINOUT      = 58
	TSParser_rulesParserVBOOL      = 59
	TSParser_rulesParserVSTRING    = 60
	TSParser_rulesParserVFLOAT     = 61
	TSParser_rulesParserVINTEGER   = 62
	TSParser_rulesParserID         = 63
	TSParser_rulesParserSL_COMMENT = 64
	TSParser_rulesParserML_COMMENT = 65
)

// TSParser_rulesParser rules.
const (
	TSParser_rulesParserRULE_start         = 0
	TSParser_rulesParserRULE_lins          = 1
	TSParser_rulesParserRULE_ins           = 2
	TSParser_rulesParserRULE_lsents        = 3
	TSParser_rulesParserRULE_sents         = 4
	TSParser_rulesParserRULE_expr          = 5
	TSParser_rulesParserRULE_declar        = 6
	TSParser_rulesParserRULE_arrayValue    = 7
	TSParser_rulesParserRULE_decons        = 8
	TSParser_rulesParserRULE_tstypes       = 9
	TSParser_rulesParserRULE_block         = 10
	TSParser_rulesParserRULE_if            = 11
	TSParser_rulesParserRULE_switch        = 12
	TSParser_rulesParserRULE_default       = 13
	TSParser_rulesParserRULE_while         = 14
	TSParser_rulesParserRULE_guard         = 15
	TSParser_rulesParserRULE_for           = 16
	TSParser_rulesParserRULE_rango         = 17
	TSParser_rulesParserRULE_strans        = 18
	TSParser_rulesParserRULE_tsprint       = 19
	TSParser_rulesParserRULE_functions     = 20
	TSParser_rulesParserRULE_parameter     = 21
	TSParser_rulesParserRULE_callFunction  = 22
	TSParser_rulesParserRULE_callParameter = 23
	TSParser_rulesParserRULE_tsfunctions   = 24
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Lins() ILinsContext
	EOF() antlr.TerminalNode

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_start
	return p
}

func InitEmptyStartContext(p *StartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_start
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Lins() ILinsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILinsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILinsContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Start_() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TSParser_rulesParserRULE_start)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Lins()
	}
	{
		p.SetState(51)
		p.Match(TSParser_rulesParserEOF)
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

// ILinsContext is an interface to support dynamic dispatch.
type ILinsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIns() []IInsContext
	Ins(i int) IInsContext

	// IsLinsContext differentiates from other interfaces.
	IsLinsContext()
}

type LinsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLinsContext() *LinsContext {
	var p = new(LinsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_lins
	return p
}

func InitEmptyLinsContext(p *LinsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_lins
}

func (*LinsContext) IsLinsContext() {}

func NewLinsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LinsContext {
	var p = new(LinsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_lins

	return p
}

func (s *LinsContext) GetParser() antlr.Parser { return s.parser }

func (s *LinsContext) AllIns() []IInsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInsContext); ok {
			len++
		}
	}

	tst := make([]IInsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInsContext); ok {
			tst[i] = t.(IInsContext)
			i++
		}
	}

	return tst
}

func (s *LinsContext) Ins(i int) IInsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInsContext); ok {
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

	return t.(IInsContext)
}

func (s *LinsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LinsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LinsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitLins(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Lins() (localctx ILinsContext) {
	localctx = NewLinsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TSParser_rulesParserRULE_lins)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-9070249623754375168) != 0 {
		{
			p.SetState(53)
			p.Ins()
		}

		p.SetState(58)
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

// IInsContext is an interface to support dynamic dispatch.
type IInsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsInsContext differentiates from other interfaces.
	IsInsContext()
}

type InsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInsContext() *InsContext {
	var p = new(InsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_ins
	return p
}

func InitEmptyInsContext(p *InsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_ins
}

func (*InsContext) IsInsContext() {}

func NewInsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InsContext {
	var p = new(InsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_ins

	return p
}

func (s *InsContext) GetParser() antlr.Parser { return s.parser }

func (s *InsContext) CopyAll(ctx *InsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *InsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IDeclContext struct {
	InsContext
}

func NewIDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IDeclContext {
	var p = new(IDeclContext)

	InitEmptyInsContext(&p.InsContext)
	p.parser = parser
	p.CopyAll(ctx.(*InsContext))

	return p
}

func (s *IDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IDeclContext) Declar() IDeclarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarContext)
}

func (s *IDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitIDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type IPrintContext struct {
	InsContext
}

func NewIPrintContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IPrintContext {
	var p = new(IPrintContext)

	InitEmptyInsContext(&p.InsContext)
	p.parser = parser
	p.CopyAll(ctx.(*InsContext))

	return p
}

func (s *IPrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IPrintContext) Tsprint() ITsprintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITsprintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITsprintContext)
}

func (s *IPrintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitIPrint(s)

	default:
		return t.VisitChildren(s)
	}
}

type IFuncContext struct {
	InsContext
}

func NewIFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IFuncContext {
	var p = new(IFuncContext)

	InitEmptyInsContext(&p.InsContext)
	p.parser = parser
	p.CopyAll(ctx.(*InsContext))

	return p
}

func (s *IFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IFuncContext) Functions() IFunctionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionsContext)
}

func (s *IFuncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitIFunc(s)

	default:
		return t.VisitChildren(s)
	}
}

type SCallFunctionContext struct {
	InsContext
}

func NewSCallFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SCallFunctionContext {
	var p = new(SCallFunctionContext)

	InitEmptyInsContext(&p.InsContext)
	p.parser = parser
	p.CopyAll(ctx.(*InsContext))

	return p
}

func (s *SCallFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SCallFunctionContext) CallFunction() ICallFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallFunctionContext)
}

func (s *SCallFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSCallFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

type IConsContext struct {
	InsContext
}

func NewIConsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IConsContext {
	var p = new(IConsContext)

	InitEmptyInsContext(&p.InsContext)
	p.parser = parser
	p.CopyAll(ctx.(*InsContext))

	return p
}

func (s *IConsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IConsContext) Decons() IDeconsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeconsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeconsContext)
}

func (s *IConsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitICons(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Ins() (localctx IInsContext) {
	localctx = NewInsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TSParser_rulesParserRULE_ins)
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TSParser_rulesParserFUNC:
		localctx = NewIFuncContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.Functions()
		}

	case TSParser_rulesParserVAR:
		localctx = NewIDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(60)
			p.Declar()
		}

	case TSParser_rulesParserLET:
		localctx = NewIConsContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(61)
			p.Decons()
		}

	case TSParser_rulesParserID:
		localctx = NewSCallFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(62)
			p.CallFunction()
		}

	case TSParser_rulesParserPRINT:
		localctx = NewIPrintContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(63)
			p.Tsprint()
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

// ILsentsContext is an interface to support dynamic dispatch.
type ILsentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSents() []ISentsContext
	Sents(i int) ISentsContext

	// IsLsentsContext differentiates from other interfaces.
	IsLsentsContext()
}

type LsentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLsentsContext() *LsentsContext {
	var p = new(LsentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_lsents
	return p
}

func InitEmptyLsentsContext(p *LsentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_lsents
}

func (*LsentsContext) IsLsentsContext() {}

func NewLsentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LsentsContext {
	var p = new(LsentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_lsents

	return p
}

func (s *LsentsContext) GetParser() antlr.Parser { return s.parser }

func (s *LsentsContext) AllSents() []ISentsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISentsContext); ok {
			len++
		}
	}

	tst := make([]ISentsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISentsContext); ok {
			tst[i] = t.(ISentsContext)
			i++
		}
	}

	return tst
}

func (s *LsentsContext) Sents(i int) ISentsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentsContext); ok {
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

	return t.(ISentsContext)
}

func (s *LsentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LsentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LsentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitLsents(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Lsents() (localctx ILsentsContext) {
	localctx = NewLsentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TSParser_rulesParserRULE_lsents)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-558645923742212090) != 0 {
		{
			p.SetState(66)
			p.Sents()
		}

		p.SetState(71)
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

// ISentsContext is an interface to support dynamic dispatch.
type ISentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSentsContext differentiates from other interfaces.
	IsSentsContext()
}

type SentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentsContext() *SentsContext {
	var p = new(SentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_sents
	return p
}

func InitEmptySentsContext(p *SentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_sents
}

func (*SentsContext) IsSentsContext() {}

func NewSentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentsContext {
	var p = new(SentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_sents

	return p
}

func (s *SentsContext) GetParser() antlr.Parser { return s.parser }

func (s *SentsContext) CopyAll(ctx *SentsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SForContext struct {
	SentsContext
}

func NewSForContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SForContext {
	var p = new(SForContext)

	InitEmptySentsContext(&p.SentsContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentsContext))

	return p
}

func (s *SForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SForContext) For_() IForContext {
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

func (s *SForContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSFor(s)

	default:
		return t.VisitChildren(s)
	}
}

type SIfContext struct {
	SentsContext
}

func NewSIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SIfContext {
	var p = new(SIfContext)

	InitEmptySentsContext(&p.SentsContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentsContext))

	return p
}

func (s *SIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SIfContext) If_() IIfContext {
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

func (s *SIfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSIf(s)

	default:
		return t.VisitChildren(s)
	}
}

type SSwitchContext struct {
	SentsContext
}

func NewSSwitchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SSwitchContext {
	var p = new(SSwitchContext)

	InitEmptySentsContext(&p.SentsContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentsContext))

	return p
}

func (s *SSwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SSwitchContext) Switch_() ISwitchContext {
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

func (s *SSwitchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSSwitch(s)

	default:
		return t.VisitChildren(s)
	}
}

type SentExprContext struct {
	SentsContext
}

func NewSentExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SentExprContext {
	var p = new(SentExprContext)

	InitEmptySentsContext(&p.SentsContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentsContext))

	return p
}

func (s *SentExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentExprContext) Expr() IExprContext {
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

func (s *SentExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSentExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SWhileContext struct {
	SentsContext
}

func NewSWhileContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SWhileContext {
	var p = new(SWhileContext)

	InitEmptySentsContext(&p.SentsContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentsContext))

	return p
}

func (s *SWhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SWhileContext) While() IWhileContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileContext)
}

func (s *SWhileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSWhile(s)

	default:
		return t.VisitChildren(s)
	}
}

type SGuardContext struct {
	SentsContext
}

func NewSGuardContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SGuardContext {
	var p = new(SGuardContext)

	InitEmptySentsContext(&p.SentsContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentsContext))

	return p
}

func (s *SGuardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SGuardContext) Guard() IGuardContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGuardContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGuardContext)
}

func (s *SGuardContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSGuard(s)

	default:
		return t.VisitChildren(s)
	}
}

type SPrintContext struct {
	SentsContext
}

func NewSPrintContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SPrintContext {
	var p = new(SPrintContext)

	InitEmptySentsContext(&p.SentsContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentsContext))

	return p
}

func (s *SPrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SPrintContext) Tsprint() ITsprintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITsprintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITsprintContext)
}

func (s *SPrintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSPrint(s)

	default:
		return t.VisitChildren(s)
	}
}

type SConsContext struct {
	SentsContext
}

func NewSConsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SConsContext {
	var p = new(SConsContext)

	InitEmptySentsContext(&p.SentsContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentsContext))

	return p
}

func (s *SConsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SConsContext) Decons() IDeconsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeconsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeconsContext)
}

func (s *SConsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSCons(s)

	default:
		return t.VisitChildren(s)
	}
}

type SentTransContext struct {
	SentsContext
}

func NewSentTransContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SentTransContext {
	var p = new(SentTransContext)

	InitEmptySentsContext(&p.SentsContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentsContext))

	return p
}

func (s *SentTransContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentTransContext) Strans() IStransContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStransContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStransContext)
}

func (s *SentTransContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSentTrans(s)

	default:
		return t.VisitChildren(s)
	}
}

type SDeclContext struct {
	SentsContext
}

func NewSDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SDeclContext {
	var p = new(SDeclContext)

	InitEmptySentsContext(&p.SentsContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentsContext))

	return p
}

func (s *SDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SDeclContext) Declar() IDeclarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarContext)
}

func (s *SDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Sents() (localctx ISentsContext) {
	localctx = NewSentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TSParser_rulesParserRULE_sents)
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TSParser_rulesParserIF:
		localctx = NewSIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)
			p.If_()
		}

	case TSParser_rulesParserSWITCH:
		localctx = NewSSwitchContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)
			p.Switch_()
		}

	case TSParser_rulesParserWHILE:
		localctx = NewSWhileContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(74)
			p.While()
		}

	case TSParser_rulesParserGUARD:
		localctx = NewSGuardContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(75)
			p.Guard()
		}

	case TSParser_rulesParserFOR:
		localctx = NewSForContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(76)
			p.For_()
		}

	case TSParser_rulesParserVAR:
		localctx = NewSDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(77)
			p.Declar()
		}

	case TSParser_rulesParserLET:
		localctx = NewSConsContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(78)
			p.Decons()
		}

	case TSParser_rulesParserPRINT:
		localctx = NewSPrintContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(79)
			p.Tsprint()
		}

	case TSParser_rulesParserT__0, TSParser_rulesParserT__1, TSParser_rulesParserT__12, TSParser_rulesParserT__20, TSParser_rulesParserNIL, TSParser_rulesParserSTRING, TSParser_rulesParserINT, TSParser_rulesParserFLOAT, TSParser_rulesParserVBOOL, TSParser_rulesParserVSTRING, TSParser_rulesParserVFLOAT, TSParser_rulesParserVINTEGER, TSParser_rulesParserID:
		localctx = NewSentExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(80)
			p.expr(0)
		}

	case TSParser_rulesParserCONTINUE, TSParser_rulesParserBREAK, TSParser_rulesParserRETURN:
		localctx = NewSentTransContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(81)
			p.Strans()
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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
	p.RuleIndex = TSParser_rulesParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EVStringContext struct {
	ExprContext
}

func NewEVStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EVStringContext {
	var p = new(EVStringContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EVStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EVStringContext) VSTRING() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserVSTRING, 0)
}

func (s *EVStringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitEVString(s)

	default:
		return t.VisitChildren(s)
	}
}

type EIDContext struct {
	ExprContext
}

func NewEIDContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EIDContext {
	var p = new(EIDContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EIDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EIDContext) ID() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserID, 0)
}

func (s *EIDContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitEID(s)

	default:
		return t.VisitChildren(s)
	}
}

type EVBOOLContext struct {
	ExprContext
}

func NewEVBOOLContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EVBOOLContext {
	var p = new(EVBOOLContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EVBOOLContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EVBOOLContext) VBOOL() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserVBOOL, 0)
}

func (s *EVBOOLContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitEVBOOL(s)

	default:
		return t.VisitChildren(s)
	}
}

type EMulDivContext struct {
	ExprContext
	op antlr.Token
}

func NewEMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EMulDivContext {
	var p = new(EMulDivContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EMulDivContext) GetOp() antlr.Token { return s.op }

func (s *EMulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *EMulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EMulDivContext) AllExpr() []IExprContext {
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

func (s *EMulDivContext) Expr(i int) IExprContext {
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

func (s *EMulDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitEMulDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type EFunctionContext struct {
	ExprContext
}

func NewEFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EFunctionContext {
	var p = new(EFunctionContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EFunctionContext) CallFunction() ICallFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallFunctionContext)
}

func (s *EFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitEFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

type ERelOrContext struct {
	ExprContext
	op antlr.Token
}

func NewERelOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ERelOrContext {
	var p = new(ERelOrContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ERelOrContext) GetOp() antlr.Token { return s.op }

func (s *ERelOrContext) SetOp(v antlr.Token) { s.op = v }

func (s *ERelOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ERelOrContext) AllExpr() []IExprContext {
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

func (s *ERelOrContext) Expr(i int) IExprContext {
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

func (s *ERelOrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitERelOr(s)

	default:
		return t.VisitChildren(s)
	}
}

type EVIntegerContext struct {
	ExprContext
}

func NewEVIntegerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EVIntegerContext {
	var p = new(EVIntegerContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EVIntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EVIntegerContext) VINTEGER() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserVINTEGER, 0)
}

func (s *EVIntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitEVInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

type ESubAddContext struct {
	ExprContext
	op antlr.Token
}

func NewESubAddContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ESubAddContext {
	var p = new(ESubAddContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ESubAddContext) GetOp() antlr.Token { return s.op }

func (s *ESubAddContext) SetOp(v antlr.Token) { s.op = v }

func (s *ESubAddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ESubAddContext) AllExpr() []IExprContext {
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

func (s *ESubAddContext) Expr(i int) IExprContext {
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

func (s *ESubAddContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitESubAdd(s)

	default:
		return t.VisitChildren(s)
	}
}

type ENegContext struct {
	ExprContext
	op antlr.Token
}

func NewENegContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ENegContext {
	var p = new(ENegContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ENegContext) GetOp() antlr.Token { return s.op }

func (s *ENegContext) SetOp(v antlr.Token) { s.op = v }

func (s *ENegContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ENegContext) Expr() IExprContext {
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

func (s *ENegContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitENeg(s)

	default:
		return t.VisitChildren(s)
	}
}

type ENILContext struct {
	ExprContext
}

func NewENILContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ENILContext {
	var p = new(ENILContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ENILContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ENILContext) NIL() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserNIL, 0)
}

func (s *ENILContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitENIL(s)

	default:
		return t.VisitChildren(s)
	}
}

type EAddSubContext struct {
	ExprContext
	op antlr.Token
}

func NewEAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EAddSubContext {
	var p = new(EAddSubContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EAddSubContext) GetOp() antlr.Token { return s.op }

func (s *EAddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *EAddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EAddSubContext) AllExpr() []IExprContext {
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

func (s *EAddSubContext) Expr(i int) IExprContext {
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

func (s *EAddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitEAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type EAsAddContext struct {
	ExprContext
	op antlr.Token
}

func NewEAsAddContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EAsAddContext {
	var p = new(EAsAddContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EAsAddContext) GetOp() antlr.Token { return s.op }

func (s *EAsAddContext) SetOp(v antlr.Token) { s.op = v }

func (s *EAsAddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EAsAddContext) AllExpr() []IExprContext {
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

func (s *EAsAddContext) Expr(i int) IExprContext {
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

func (s *EAsAddContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitEAsAdd(s)

	default:
		return t.VisitChildren(s)
	}
}

type ENotContext struct {
	ExprContext
	op antlr.Token
}

func NewENotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ENotContext {
	var p = new(ENotContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ENotContext) GetOp() antlr.Token { return s.op }

func (s *ENotContext) SetOp(v antlr.Token) { s.op = v }

func (s *ENotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ENotContext) Expr() IExprContext {
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

func (s *ENotContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitENot(s)

	default:
		return t.VisitChildren(s)
	}
}

type EModuleContext struct {
	ExprContext
	op antlr.Token
}

func NewEModuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EModuleContext {
	var p = new(EModuleContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EModuleContext) GetOp() antlr.Token { return s.op }

func (s *EModuleContext) SetOp(v antlr.Token) { s.op = v }

func (s *EModuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EModuleContext) AllExpr() []IExprContext {
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

func (s *EModuleContext) Expr(i int) IExprContext {
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

func (s *EModuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitEModule(s)

	default:
		return t.VisitChildren(s)
	}
}

type EVArrayContext struct {
	ExprContext
}

func NewEVArrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EVArrayContext {
	var p = new(EVArrayContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EVArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EVArrayContext) ArrayValue() IArrayValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayValueContext)
}

func (s *EVArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitEVArray(s)

	default:
		return t.VisitChildren(s)
	}
}

type EVFloatContext struct {
	ExprContext
}

func NewEVFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EVFloatContext {
	var p = new(EVFloatContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EVFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EVFloatContext) VFLOAT() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserVFLOAT, 0)
}

func (s *EVFloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitEVFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

type EParentContext struct {
	ExprContext
}

func NewEParentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EParentContext {
	var p = new(EParentContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EParentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EParentContext) Expr() IExprContext {
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

func (s *EParentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitEParent(s)

	default:
		return t.VisitChildren(s)
	}
}

type ERelContext struct {
	ExprContext
	op antlr.Token
}

func NewERelContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ERelContext {
	var p = new(ERelContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ERelContext) GetOp() antlr.Token { return s.op }

func (s *ERelContext) SetOp(v antlr.Token) { s.op = v }

func (s *ERelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ERelContext) AllExpr() []IExprContext {
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

func (s *ERelContext) Expr(i int) IExprContext {
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

func (s *ERelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitERel(s)

	default:
		return t.VisitChildren(s)
	}
}

type EAssignContext struct {
	ExprContext
	op antlr.Token
}

func NewEAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EAssignContext {
	var p = new(EAssignContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EAssignContext) GetOp() antlr.Token { return s.op }

func (s *EAssignContext) SetOp(v antlr.Token) { s.op = v }

func (s *EAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EAssignContext) AllExpr() []IExprContext {
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

func (s *EAssignContext) Expr(i int) IExprContext {
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

func (s *EAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitEAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

type ETSFunctionsContext struct {
	ExprContext
}

func NewETSFunctionsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ETSFunctionsContext {
	var p = new(ETSFunctionsContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ETSFunctionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ETSFunctionsContext) Tsfunctions() ITsfunctionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITsfunctionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITsfunctionsContext)
}

func (s *ETSFunctionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitETSFunctions(s)

	default:
		return t.VisitChildren(s)
	}
}

type ERelAndContext struct {
	ExprContext
	op antlr.Token
}

func NewERelAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ERelAndContext {
	var p = new(ERelAndContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ERelAndContext) GetOp() antlr.Token { return s.op }

func (s *ERelAndContext) SetOp(v antlr.Token) { s.op = v }

func (s *ERelAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ERelAndContext) AllExpr() []IExprContext {
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

func (s *ERelAndContext) Expr(i int) IExprContext {
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

func (s *ERelAndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitERelAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *TSParser_rulesParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, TSParser_rulesParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewENegContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(85)

			var _m = p.Match(TSParser_rulesParserT__0)

			localctx.(*ENegContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(86)
			p.expr(21)
		}

	case 2:
		localctx = NewENotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(87)

			var _m = p.Match(TSParser_rulesParserT__1)

			localctx.(*ENotContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)
			p.expr(20)
		}

	case 3:
		localctx = NewETSFunctionsContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(89)
			p.Tsfunctions()
		}

	case 4:
		localctx = NewEFunctionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(90)
			p.CallFunction()
		}

	case 5:
		localctx = NewEVArrayContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(91)
			p.ArrayValue()
		}

	case 6:
		localctx = NewEParentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(92)
			p.Match(TSParser_rulesParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(93)
			p.expr(0)
		}
		{
			p.SetState(94)
			p.Match(TSParser_rulesParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewEVStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(96)
			p.Match(TSParser_rulesParserVSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewEVIntegerContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(97)
			p.Match(TSParser_rulesParserVINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewEVFloatContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(98)
			p.Match(TSParser_rulesParserVFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewEVBOOLContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(99)
			p.Match(TSParser_rulesParserVBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		localctx = NewEIDContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(100)
			p.Match(TSParser_rulesParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		localctx = NewENILContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(101)
			p.Match(TSParser_rulesParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(131)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewEModuleContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(104)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(105)

					var _m = p.Match(TSParser_rulesParserT__2)

					localctx.(*EModuleContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(106)
					p.expr(20)
				}

			case 2:
				localctx = NewEMulDivContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(107)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(108)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EMulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TSParser_rulesParserT__3 || _la == TSParser_rulesParserT__4) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EMulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(109)
					p.expr(19)
				}

			case 3:
				localctx = NewEAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(110)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(111)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EAddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TSParser_rulesParserT__0 || _la == TSParser_rulesParserT__5) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EAddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(112)
					p.expr(18)
				}

			case 4:
				localctx = NewERelContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(113)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(114)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ERelContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8064) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ERelContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(115)
					p.expr(17)
				}

			case 5:
				localctx = NewERelAndContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(116)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(117)

					var _m = p.Match(TSParser_rulesParserT__14)

					localctx.(*ERelAndContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(118)
					p.expr(12)
				}

			case 6:
				localctx = NewERelOrContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(119)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(120)

					var _m = p.Match(TSParser_rulesParserT__15)

					localctx.(*ERelOrContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(121)
					p.expr(11)
				}

			case 7:
				localctx = NewEAssignContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(122)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(123)

					var _m = p.Match(TSParser_rulesParserT__16)

					localctx.(*EAssignContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(124)
					p.expr(9)
				}

			case 8:
				localctx = NewEAsAddContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(125)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(126)

					var _m = p.Match(TSParser_rulesParserT__17)

					localctx.(*EAsAddContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(127)
					p.expr(8)
				}

			case 9:
				localctx = NewESubAddContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSParser_rulesParserRULE_expr)
				p.SetState(128)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(129)

					var _m = p.Match(TSParser_rulesParserT__18)

					localctx.(*ESubAddContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(130)
					p.expr(7)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
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

// IDeclarContext is an interface to support dynamic dispatch.
type IDeclarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDeclarContext differentiates from other interfaces.
	IsDeclarContext()
}

type DeclarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarContext() *DeclarContext {
	var p = new(DeclarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_declar
	return p
}

func InitEmptyDeclarContext(p *DeclarContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_declar
}

func (*DeclarContext) IsDeclarContext() {}

func NewDeclarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarContext {
	var p = new(DeclarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_declar

	return p
}

func (s *DeclarContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarContext) CopyAll(ctx *DeclarContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DeclarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SDTypeContext struct {
	DeclarContext
	arr antlr.Token
}

func NewSDTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SDTypeContext {
	var p = new(SDTypeContext)

	InitEmptyDeclarContext(&p.DeclarContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclarContext))

	return p
}

func (s *SDTypeContext) GetArr() antlr.Token { return s.arr }

func (s *SDTypeContext) SetArr(v antlr.Token) { s.arr = v }

func (s *SDTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SDTypeContext) VAR() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserVAR, 0)
}

func (s *SDTypeContext) ID() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserID, 0)
}

func (s *SDTypeContext) Tstypes() ITstypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITstypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITstypesContext)
}

func (s *SDTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSDType(s)

	default:
		return t.VisitChildren(s)
	}
}

type SDecTAssignContext struct {
	DeclarContext
	arr antlr.Token
	op  antlr.Token
}

func NewSDecTAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SDecTAssignContext {
	var p = new(SDecTAssignContext)

	InitEmptyDeclarContext(&p.DeclarContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclarContext))

	return p
}

func (s *SDecTAssignContext) GetArr() antlr.Token { return s.arr }

func (s *SDecTAssignContext) GetOp() antlr.Token { return s.op }

func (s *SDecTAssignContext) SetArr(v antlr.Token) { s.arr = v }

func (s *SDecTAssignContext) SetOp(v antlr.Token) { s.op = v }

func (s *SDecTAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SDecTAssignContext) VAR() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserVAR, 0)
}

func (s *SDecTAssignContext) ID() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserID, 0)
}

func (s *SDecTAssignContext) Expr() IExprContext {
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

func (s *SDecTAssignContext) Tstypes() ITstypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITstypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITstypesContext)
}

func (s *SDecTAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSDecTAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

type SDecAssignContext struct {
	DeclarContext
}

func NewSDecAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SDecAssignContext {
	var p = new(SDecAssignContext)

	InitEmptyDeclarContext(&p.DeclarContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclarContext))

	return p
}

func (s *SDecAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SDecAssignContext) VAR() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserVAR, 0)
}

func (s *SDecAssignContext) ID() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserID, 0)
}

func (s *SDecAssignContext) Expr() IExprContext {
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

func (s *SDecAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSDecAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Declar() (localctx IDeclarContext) {
	localctx = NewDeclarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TSParser_rulesParserRULE_declar)
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSDTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(136)
			p.Match(TSParser_rulesParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.Match(TSParser_rulesParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Match(TSParser_rulesParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case TSParser_rulesParserSTRING, TSParser_rulesParserINT, TSParser_rulesParserFLOAT, TSParser_rulesParserBOOL, TSParser_rulesParserCHARACTER:
			{
				p.SetState(139)
				p.Tstypes()
			}

		case TSParser_rulesParserT__20:
			{
				p.SetState(140)

				var _m = p.Match(TSParser_rulesParserT__20)

				localctx.(*SDTypeContext).arr = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(141)
				p.Tstypes()
			}
			{
				p.SetState(142)
				p.Match(TSParser_rulesParserT__21)
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
			p.SetState(146)
			p.Match(TSParser_rulesParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewSDecAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(148)
			p.Match(TSParser_rulesParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)
			p.Match(TSParser_rulesParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)
			p.Match(TSParser_rulesParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)
			p.expr(0)
		}

	case 3:
		localctx = NewSDecTAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(152)
			p.Match(TSParser_rulesParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)
			p.Match(TSParser_rulesParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.Match(TSParser_rulesParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case TSParser_rulesParserSTRING, TSParser_rulesParserINT, TSParser_rulesParserFLOAT, TSParser_rulesParserBOOL, TSParser_rulesParserCHARACTER:
			{
				p.SetState(155)
				p.Tstypes()
			}

		case TSParser_rulesParserT__20:
			{
				p.SetState(156)

				var _m = p.Match(TSParser_rulesParserT__20)

				localctx.(*SDecTAssignContext).arr = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(157)
				p.Tstypes()
			}
			{
				p.SetState(158)
				p.Match(TSParser_rulesParserT__21)
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
			p.SetState(162)

			var _m = p.Match(TSParser_rulesParserT__16)

			localctx.(*SDecTAssignContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)
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

// IArrayValueContext is an interface to support dynamic dispatch.
type IArrayValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsArrayValueContext differentiates from other interfaces.
	IsArrayValueContext()
}

type ArrayValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyArrayValueContext() *ArrayValueContext {
	var p = new(ArrayValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_arrayValue
	return p
}

func InitEmptyArrayValueContext(p *ArrayValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_arrayValue
}

func (*ArrayValueContext) IsArrayValueContext() {}

func NewArrayValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayValueContext {
	var p = new(ArrayValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_arrayValue

	return p
}

func (s *ArrayValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayValueContext) GetOp() antlr.Token { return s.op }

func (s *ArrayValueContext) SetOp(v antlr.Token) { s.op = v }

func (s *ArrayValueContext) AllExpr() []IExprContext {
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

func (s *ArrayValueContext) Expr(i int) IExprContext {
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

func (s *ArrayValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitArrayValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) ArrayValue() (localctx IArrayValueContext) {
	localctx = NewArrayValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TSParser_rulesParserRULE_arrayValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)

		var _m = p.Match(TSParser_rulesParserT__20)

		localctx.(*ArrayValueContext).op = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-576460236905242618) != 0 {
		{
			p.SetState(168)
			p.expr(0)
		}
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TSParser_rulesParserT__23 {
			{
				p.SetState(169)
				p.Match(TSParser_rulesParserT__23)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(170)
				p.expr(0)
			}

			p.SetState(175)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(178)
		p.Match(TSParser_rulesParserT__21)
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

// IDeconsContext is an interface to support dynamic dispatch.
type IDeconsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDeconsContext differentiates from other interfaces.
	IsDeconsContext()
}

type DeconsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeconsContext() *DeconsContext {
	var p = new(DeconsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_decons
	return p
}

func InitEmptyDeconsContext(p *DeconsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_decons
}

func (*DeconsContext) IsDeconsContext() {}

func NewDeconsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeconsContext {
	var p = new(DeconsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_decons

	return p
}

func (s *DeconsContext) GetParser() antlr.Parser { return s.parser }

func (s *DeconsContext) CopyAll(ctx *DeconsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DeconsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeconsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SConsTypeAssContext struct {
	DeconsContext
	op antlr.Token
}

func NewSConsTypeAssContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SConsTypeAssContext {
	var p = new(SConsTypeAssContext)

	InitEmptyDeconsContext(&p.DeconsContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeconsContext))

	return p
}

func (s *SConsTypeAssContext) GetOp() antlr.Token { return s.op }

func (s *SConsTypeAssContext) SetOp(v antlr.Token) { s.op = v }

func (s *SConsTypeAssContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SConsTypeAssContext) LET() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserLET, 0)
}

func (s *SConsTypeAssContext) ID() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserID, 0)
}

func (s *SConsTypeAssContext) Tstypes() ITstypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITstypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITstypesContext)
}

func (s *SConsTypeAssContext) Expr() IExprContext {
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

func (s *SConsTypeAssContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSConsTypeAss(s)

	default:
		return t.VisitChildren(s)
	}
}

type SConsAssContext struct {
	DeconsContext
}

func NewSConsAssContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SConsAssContext {
	var p = new(SConsAssContext)

	InitEmptyDeconsContext(&p.DeconsContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeconsContext))

	return p
}

func (s *SConsAssContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SConsAssContext) LET() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserLET, 0)
}

func (s *SConsAssContext) ID() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserID, 0)
}

func (s *SConsAssContext) Expr() IExprContext {
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

func (s *SConsAssContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSConsAss(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Decons() (localctx IDeconsContext) {
	localctx = NewDeconsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TSParser_rulesParserRULE_decons)
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSConsAssContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.Match(TSParser_rulesParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(181)
			p.Match(TSParser_rulesParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.Match(TSParser_rulesParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)
			p.expr(0)
		}

	case 2:
		localctx = NewSConsTypeAssContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(184)
			p.Match(TSParser_rulesParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)
			p.Match(TSParser_rulesParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.Match(TSParser_rulesParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Tstypes()
		}
		{
			p.SetState(188)

			var _m = p.Match(TSParser_rulesParserT__16)

			localctx.(*SConsTypeAssContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
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

// ITstypesContext is an interface to support dynamic dispatch.
type ITstypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	CHARACTER() antlr.TerminalNode

	// IsTstypesContext differentiates from other interfaces.
	IsTstypesContext()
}

type TstypesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTstypesContext() *TstypesContext {
	var p = new(TstypesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_tstypes
	return p
}

func InitEmptyTstypesContext(p *TstypesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_tstypes
}

func (*TstypesContext) IsTstypesContext() {}

func NewTstypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TstypesContext {
	var p = new(TstypesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_tstypes

	return p
}

func (s *TstypesContext) GetParser() antlr.Parser { return s.parser }

func (s *TstypesContext) STRING() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserSTRING, 0)
}

func (s *TstypesContext) INT() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserINT, 0)
}

func (s *TstypesContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserFLOAT, 0)
}

func (s *TstypesContext) BOOL() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserBOOL, 0)
}

func (s *TstypesContext) CHARACTER() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserCHARACTER, 0)
}

func (s *TstypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TstypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TstypesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitTstypes(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Tstypes() (localctx ITstypesContext) {
	localctx = NewTstypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TSParser_rulesParserRULE_tstypes)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2130303778816) != 0) {
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTerm returns the term token.
	GetTerm() antlr.Token

	// SetTerm sets the term token.
	SetTerm(antlr.Token)

	// Getter signatures
	Lsents() ILsentsContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	term   antlr.Token
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) GetTerm() antlr.Token { return s.term }

func (s *BlockContext) SetTerm(v antlr.Token) { s.term = v }

func (s *BlockContext) Lsents() ILsentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILsentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILsentsContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TSParser_rulesParserRULE_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)

		var _m = p.Match(TSParser_rulesParserT__24)

		localctx.(*BlockContext).term = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(196)
		p.Lsents()
	}
	{
		p.SetState(197)
		p.Match(TSParser_rulesParserT__25)
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

// IIfContext is an interface to support dynamic dispatch.
type IIfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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
	p.RuleIndex = TSParser_rulesParserRULE_if
	return p
}

func InitEmptyIfContext(p *IfContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_if
}

func (*IfContext) IsIfContext() {}

func NewIfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfContext {
	var p = new(IfContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_if

	return p
}

func (s *IfContext) GetParser() antlr.Parser { return s.parser }

func (s *IfContext) CopyAll(ctx *IfContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RIfEIfContext struct {
	IfContext
}

func NewRIfEIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RIfEIfContext {
	var p = new(RIfEIfContext)

	InitEmptyIfContext(&p.IfContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfContext))

	return p
}

func (s *RIfEIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RIfEIfContext) IF() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserIF, 0)
}

func (s *RIfEIfContext) Expr() IExprContext {
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

func (s *RIfEIfContext) Block() IBlockContext {
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

func (s *RIfEIfContext) ELSE() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserELSE, 0)
}

func (s *RIfEIfContext) If_() IIfContext {
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

func (s *RIfEIfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitRIfEIf(s)

	default:
		return t.VisitChildren(s)
	}
}

type RIfElseContext struct {
	IfContext
}

func NewRIfElseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RIfElseContext {
	var p = new(RIfElseContext)

	InitEmptyIfContext(&p.IfContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfContext))

	return p
}

func (s *RIfElseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RIfElseContext) IF() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserIF, 0)
}

func (s *RIfElseContext) Expr() IExprContext {
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

func (s *RIfElseContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *RIfElseContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
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

	return t.(IBlockContext)
}

func (s *RIfElseContext) ELSE() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserELSE, 0)
}

func (s *RIfElseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitRIfElse(s)

	default:
		return t.VisitChildren(s)
	}
}

type RIfContext struct {
	IfContext
}

func NewRIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RIfContext {
	var p = new(RIfContext)

	InitEmptyIfContext(&p.IfContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfContext))

	return p
}

func (s *RIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RIfContext) IF() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserIF, 0)
}

func (s *RIfContext) Expr() IExprContext {
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

func (s *RIfContext) Block() IBlockContext {
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

func (s *RIfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitRIf(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) If_() (localctx IIfContext) {
	localctx = NewIfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, TSParser_rulesParserRULE_if)
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		localctx = NewRIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(199)
			p.Match(TSParser_rulesParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(200)
			p.expr(0)
		}
		{
			p.SetState(201)
			p.Block()
		}

	case 2:
		localctx = NewRIfElseContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(203)
			p.Match(TSParser_rulesParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(204)
			p.expr(0)
		}
		{
			p.SetState(205)
			p.Block()
		}
		{
			p.SetState(206)
			p.Match(TSParser_rulesParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)
			p.Block()
		}

	case 3:
		localctx = NewRIfEIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(209)
			p.Match(TSParser_rulesParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(210)
			p.expr(0)
		}
		{
			p.SetState(211)
			p.Block()
		}
		{
			p.SetState(212)
			p.Match(TSParser_rulesParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
			p.If_()
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

// ISwitchContext is an interface to support dynamic dispatch.
type ISwitchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInput returns the input rule contexts.
	GetInput() IExprContext

	// SetInput sets the input rule contexts.
	SetInput(IExprContext)

	// Getter signatures
	SWITCH() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCASE() []antlr.TerminalNode
	CASE(i int) antlr.TerminalNode
	AllLsents() []ILsentsContext
	Lsents(i int) ILsentsContext
	Default_() IDefaultContext

	// IsSwitchContext differentiates from other interfaces.
	IsSwitchContext()
}

type SwitchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	input  IExprContext
}

func NewEmptySwitchContext() *SwitchContext {
	var p = new(SwitchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_switch
	return p
}

func InitEmptySwitchContext(p *SwitchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_switch
}

func (*SwitchContext) IsSwitchContext() {}

func NewSwitchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchContext {
	var p = new(SwitchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_switch

	return p
}

func (s *SwitchContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchContext) GetInput() IExprContext { return s.input }

func (s *SwitchContext) SetInput(v IExprContext) { s.input = v }

func (s *SwitchContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserSWITCH, 0)
}

func (s *SwitchContext) AllExpr() []IExprContext {
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

func (s *SwitchContext) Expr(i int) IExprContext {
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

func (s *SwitchContext) AllCASE() []antlr.TerminalNode {
	return s.GetTokens(TSParser_rulesParserCASE)
}

func (s *SwitchContext) CASE(i int) antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserCASE, i)
}

func (s *SwitchContext) AllLsents() []ILsentsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILsentsContext); ok {
			len++
		}
	}

	tst := make([]ILsentsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILsentsContext); ok {
			tst[i] = t.(ILsentsContext)
			i++
		}
	}

	return tst
}

func (s *SwitchContext) Lsents(i int) ILsentsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILsentsContext); ok {
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

	return t.(ILsentsContext)
}

func (s *SwitchContext) Default_() IDefaultContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultContext)
}

func (s *SwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitSwitch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Switch_() (localctx ISwitchContext) {
	localctx = NewSwitchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TSParser_rulesParserRULE_switch)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(217)
		p.Match(TSParser_rulesParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(218)

		var _x = p.expr(0)

		localctx.(*SwitchContext).input = _x
	}
	{
		p.SetState(219)
		p.Match(TSParser_rulesParserT__24)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == TSParser_rulesParserCASE {
		{
			p.SetState(220)
			p.Match(TSParser_rulesParserCASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(221)
			p.expr(0)
		}
		{
			p.SetState(222)
			p.Match(TSParser_rulesParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)
			p.Lsents()
		}

		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TSParser_rulesParserDEFAULT {
		{
			p.SetState(229)
			p.Default_()
		}

	}
	{
		p.SetState(232)
		p.Match(TSParser_rulesParserT__25)
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

// IDefaultContext is an interface to support dynamic dispatch.
type IDefaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFAULT() antlr.TerminalNode
	Lsents() ILsentsContext

	// IsDefaultContext differentiates from other interfaces.
	IsDefaultContext()
}

type DefaultContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultContext() *DefaultContext {
	var p = new(DefaultContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_default
	return p
}

func InitEmptyDefaultContext(p *DefaultContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_default
}

func (*DefaultContext) IsDefaultContext() {}

func NewDefaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultContext {
	var p = new(DefaultContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_default

	return p
}

func (s *DefaultContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserDEFAULT, 0)
}

func (s *DefaultContext) Lsents() ILsentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILsentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILsentsContext)
}

func (s *DefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitDefault(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Default_() (localctx IDefaultContext) {
	localctx = NewDefaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, TSParser_rulesParserRULE_default)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		p.Match(TSParser_rulesParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(235)
		p.Match(TSParser_rulesParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(236)
		p.Lsents()
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

// IWhileContext is an interface to support dynamic dispatch.
type IWhileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expr() IExprContext
	Block() IBlockContext

	// IsWhileContext differentiates from other interfaces.
	IsWhileContext()
}

type WhileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileContext() *WhileContext {
	var p = new(WhileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_while
	return p
}

func InitEmptyWhileContext(p *WhileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_while
}

func (*WhileContext) IsWhileContext() {}

func NewWhileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileContext {
	var p = new(WhileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_while

	return p
}

func (s *WhileContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileContext) WHILE() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserWHILE, 0)
}

func (s *WhileContext) Expr() IExprContext {
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

func (s *WhileContext) Block() IBlockContext {
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

func (s *WhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitWhile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) While() (localctx IWhileContext) {
	localctx = NewWhileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, TSParser_rulesParserRULE_while)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.Match(TSParser_rulesParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(239)
		p.expr(0)
	}
	{
		p.SetState(240)
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

// IGuardContext is an interface to support dynamic dispatch.
type IGuardContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GUARD() antlr.TerminalNode
	Expr() IExprContext
	ELSE() antlr.TerminalNode
	Block() IBlockContext

	// IsGuardContext differentiates from other interfaces.
	IsGuardContext()
}

type GuardContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGuardContext() *GuardContext {
	var p = new(GuardContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_guard
	return p
}

func InitEmptyGuardContext(p *GuardContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_guard
}

func (*GuardContext) IsGuardContext() {}

func NewGuardContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GuardContext {
	var p = new(GuardContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_guard

	return p
}

func (s *GuardContext) GetParser() antlr.Parser { return s.parser }

func (s *GuardContext) GUARD() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserGUARD, 0)
}

func (s *GuardContext) Expr() IExprContext {
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

func (s *GuardContext) ELSE() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserELSE, 0)
}

func (s *GuardContext) Block() IBlockContext {
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

func (s *GuardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GuardContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GuardContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitGuard(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Guard() (localctx IGuardContext) {
	localctx = NewGuardContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, TSParser_rulesParserRULE_guard)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(TSParser_rulesParserGUARD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
		p.expr(0)
	}
	{
		p.SetState(244)
		p.Match(TSParser_rulesParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(245)
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

	// Getter signatures
	FOR() antlr.TerminalNode
	ID() antlr.TerminalNode
	IN() antlr.TerminalNode
	Block() IBlockContext
	Expr() IExprContext
	Rango() IRangoContext

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
	p.RuleIndex = TSParser_rulesParserRULE_for
	return p
}

func InitEmptyForContext(p *ForContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_for
}

func (*ForContext) IsForContext() {}

func NewForContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForContext {
	var p = new(ForContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_for

	return p
}

func (s *ForContext) GetParser() antlr.Parser { return s.parser }

func (s *ForContext) FOR() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserFOR, 0)
}

func (s *ForContext) ID() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserID, 0)
}

func (s *ForContext) IN() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserIN, 0)
}

func (s *ForContext) Block() IBlockContext {
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

func (s *ForContext) Expr() IExprContext {
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

func (s *ForContext) Rango() IRangoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangoContext)
}

func (s *ForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitFor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) For_() (localctx IForContext) {
	localctx = NewForContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, TSParser_rulesParserRULE_for)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.Match(TSParser_rulesParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(248)
		p.Match(TSParser_rulesParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.Match(TSParser_rulesParserIN)
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

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(250)
			p.expr(0)
		}

	case 2:
		{
			p.SetState(251)
			p.Rango()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(254)
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

// IRangoContext is an interface to support dynamic dispatch.
type IRangoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVINTEGER() []antlr.TerminalNode
	VINTEGER(i int) antlr.TerminalNode

	// IsRangoContext differentiates from other interfaces.
	IsRangoContext()
}

type RangoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangoContext() *RangoContext {
	var p = new(RangoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_rango
	return p
}

func InitEmptyRangoContext(p *RangoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_rango
}

func (*RangoContext) IsRangoContext() {}

func NewRangoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangoContext {
	var p = new(RangoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_rango

	return p
}

func (s *RangoContext) GetParser() antlr.Parser { return s.parser }

func (s *RangoContext) AllVINTEGER() []antlr.TerminalNode {
	return s.GetTokens(TSParser_rulesParserVINTEGER)
}

func (s *RangoContext) VINTEGER(i int) antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserVINTEGER, i)
}

func (s *RangoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitRango(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Rango() (localctx IRangoContext) {
	localctx = NewRangoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, TSParser_rulesParserRULE_rango)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.Match(TSParser_rulesParserVINTEGER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(257)
		p.Match(TSParser_rulesParserT__26)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)
		p.Match(TSParser_rulesParserVINTEGER)
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

// IStransContext is an interface to support dynamic dispatch.
type IStransContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTINUE() antlr.TerminalNode
	BREAK() antlr.TerminalNode
	RETURN() antlr.TerminalNode
	Expr() IExprContext

	// IsStransContext differentiates from other interfaces.
	IsStransContext()
}

type StransContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStransContext() *StransContext {
	var p = new(StransContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_strans
	return p
}

func InitEmptyStransContext(p *StransContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_strans
}

func (*StransContext) IsStransContext() {}

func NewStransContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StransContext {
	var p = new(StransContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_strans

	return p
}

func (s *StransContext) GetParser() antlr.Parser { return s.parser }

func (s *StransContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserCONTINUE, 0)
}

func (s *StransContext) BREAK() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserBREAK, 0)
}

func (s *StransContext) RETURN() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserRETURN, 0)
}

func (s *StransContext) Expr() IExprContext {
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

func (s *StransContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StransContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StransContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitStrans(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Strans() (localctx IStransContext) {
	localctx = NewStransContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, TSParser_rulesParserRULE_strans)
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TSParser_rulesParserCONTINUE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(260)
			p.Match(TSParser_rulesParserCONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TSParser_rulesParserBREAK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(261)
			p.Match(TSParser_rulesParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TSParser_rulesParserRETURN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(262)
			p.Match(TSParser_rulesParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(264)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(263)
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

// ITsprintContext is an interface to support dynamic dispatch.
type ITsprintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINT() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsTsprintContext differentiates from other interfaces.
	IsTsprintContext()
}

type TsprintContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTsprintContext() *TsprintContext {
	var p = new(TsprintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_tsprint
	return p
}

func InitEmptyTsprintContext(p *TsprintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_tsprint
}

func (*TsprintContext) IsTsprintContext() {}

func NewTsprintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TsprintContext {
	var p = new(TsprintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_tsprint

	return p
}

func (s *TsprintContext) GetParser() antlr.Parser { return s.parser }

func (s *TsprintContext) PRINT() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserPRINT, 0)
}

func (s *TsprintContext) AllExpr() []IExprContext {
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

func (s *TsprintContext) Expr(i int) IExprContext {
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

func (s *TsprintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TsprintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TsprintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitTsprint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Tsprint() (localctx ITsprintContext) {
	localctx = NewTsprintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, TSParser_rulesParserRULE_tsprint)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.Match(TSParser_rulesParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(269)
		p.Match(TSParser_rulesParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(270)
		p.expr(0)
	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TSParser_rulesParserT__23 {
		{
			p.SetState(271)
			p.Match(TSParser_rulesParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(272)
			p.expr(0)
		}

		p.SetState(277)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(278)
		p.Match(TSParser_rulesParserT__13)
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

// IFunctionsContext is an interface to support dynamic dispatch.
type IFunctionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	ID() antlr.TerminalNode
	Lsents() ILsentsContext
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext
	Tstypes() ITstypesContext

	// IsFunctionsContext differentiates from other interfaces.
	IsFunctionsContext()
}

type FunctionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionsContext() *FunctionsContext {
	var p = new(FunctionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_functions
	return p
}

func InitEmptyFunctionsContext(p *FunctionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_functions
}

func (*FunctionsContext) IsFunctionsContext() {}

func NewFunctionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionsContext {
	var p = new(FunctionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_functions

	return p
}

func (s *FunctionsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionsContext) FUNC() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserFUNC, 0)
}

func (s *FunctionsContext) ID() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserID, 0)
}

func (s *FunctionsContext) Lsents() ILsentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILsentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILsentsContext)
}

func (s *FunctionsContext) AllParameter() []IParameterContext {
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

func (s *FunctionsContext) Parameter(i int) IParameterContext {
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

func (s *FunctionsContext) Tstypes() ITstypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITstypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITstypesContext)
}

func (s *FunctionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitFunctions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Functions() (localctx IFunctionsContext) {
	localctx = NewFunctionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, TSParser_rulesParserRULE_functions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(280)
		p.Match(TSParser_rulesParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(281)
		p.Match(TSParser_rulesParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(282)
		p.Match(TSParser_rulesParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TSParser_rulesParserT__28 || _la == TSParser_rulesParserID {
		{
			p.SetState(283)
			p.Parameter()
		}
		p.SetState(288)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TSParser_rulesParserT__23 {
			{
				p.SetState(284)
				p.Match(TSParser_rulesParserT__23)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(285)
				p.Parameter()
			}

			p.SetState(290)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(293)
		p.Match(TSParser_rulesParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TSParser_rulesParserT__27 {
		{
			p.SetState(294)
			p.Match(TSParser_rulesParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(295)
			p.Tstypes()
		}

	}
	{
		p.SetState(298)
		p.Match(TSParser_rulesParserT__24)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(299)
		p.Lsents()
	}
	{
		p.SetState(300)
		p.Match(TSParser_rulesParserT__25)
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

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAlias returns the alias token.
	GetAlias() antlr.Token

	// GetArr returns the arr token.
	GetArr() antlr.Token

	// SetAlias sets the alias token.
	SetAlias(antlr.Token)

	// SetArr sets the arr token.
	SetArr(antlr.Token)

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	Tstypes() ITstypesContext
	INOUT() antlr.TerminalNode

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	alias  antlr.Token
	arr    antlr.Token
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_parameter
	return p
}

func InitEmptyParameterContext(p *ParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_parameter
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) GetAlias() antlr.Token { return s.alias }

func (s *ParameterContext) GetArr() antlr.Token { return s.arr }

func (s *ParameterContext) SetAlias(v antlr.Token) { s.alias = v }

func (s *ParameterContext) SetArr(v antlr.Token) { s.arr = v }

func (s *ParameterContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TSParser_rulesParserID)
}

func (s *ParameterContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserID, i)
}

func (s *ParameterContext) Tstypes() ITstypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITstypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITstypesContext)
}

func (s *ParameterContext) INOUT() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserINOUT, 0)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, TSParser_rulesParserRULE_parameter)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(303)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(302)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ParameterContext).alias = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == TSParser_rulesParserT__28 || _la == TSParser_rulesParserID) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ParameterContext).alias = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(305)
		p.Match(TSParser_rulesParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(306)
		p.Match(TSParser_rulesParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(308)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TSParser_rulesParserINOUT {
		{
			p.SetState(307)
			p.Match(TSParser_rulesParserINOUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TSParser_rulesParserSTRING, TSParser_rulesParserINT, TSParser_rulesParserFLOAT, TSParser_rulesParserBOOL, TSParser_rulesParserCHARACTER:
		{
			p.SetState(310)
			p.Tstypes()
		}

	case TSParser_rulesParserT__20:
		{
			p.SetState(311)

			var _m = p.Match(TSParser_rulesParserT__20)

			localctx.(*ParameterContext).arr = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(312)
			p.Tstypes()
		}
		{
			p.SetState(313)
			p.Match(TSParser_rulesParserT__21)
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

// ICallFunctionContext is an interface to support dynamic dispatch.
type ICallFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	AllCallParameter() []ICallParameterContext
	CallParameter(i int) ICallParameterContext

	// IsCallFunctionContext differentiates from other interfaces.
	IsCallFunctionContext()
}

type CallFunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallFunctionContext() *CallFunctionContext {
	var p = new(CallFunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_callFunction
	return p
}

func InitEmptyCallFunctionContext(p *CallFunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_callFunction
}

func (*CallFunctionContext) IsCallFunctionContext() {}

func NewCallFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallFunctionContext {
	var p = new(CallFunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_callFunction

	return p
}

func (s *CallFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *CallFunctionContext) ID() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserID, 0)
}

func (s *CallFunctionContext) AllCallParameter() []ICallParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICallParameterContext); ok {
			len++
		}
	}

	tst := make([]ICallParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICallParameterContext); ok {
			tst[i] = t.(ICallParameterContext)
			i++
		}
	}

	return tst
}

func (s *CallFunctionContext) CallParameter(i int) ICallParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallParameterContext); ok {
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

	return t.(ICallParameterContext)
}

func (s *CallFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitCallFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) CallFunction() (localctx ICallFunctionContext) {
	localctx = NewCallFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, TSParser_rulesParserRULE_callFunction)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(317)
		p.Match(TSParser_rulesParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(318)
		p.Match(TSParser_rulesParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-576460235831500794) != 0 {
		{
			p.SetState(319)
			p.CallParameter()
		}
		p.SetState(324)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TSParser_rulesParserT__23 {
			{
				p.SetState(320)
				p.Match(TSParser_rulesParserT__23)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(321)
				p.CallParameter()
			}

			p.SetState(326)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(329)
		p.Match(TSParser_rulesParserT__13)
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

// ICallParameterContext is an interface to support dynamic dispatch.
type ICallParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	Expr() IExprContext
	ID() antlr.TerminalNode

	// IsCallParameterContext differentiates from other interfaces.
	IsCallParameterContext()
}

type CallParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyCallParameterContext() *CallParameterContext {
	var p = new(CallParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_callParameter
	return p
}

func InitEmptyCallParameterContext(p *CallParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_callParameter
}

func (*CallParameterContext) IsCallParameterContext() {}

func NewCallParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallParameterContext {
	var p = new(CallParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_callParameter

	return p
}

func (s *CallParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *CallParameterContext) GetOp() antlr.Token { return s.op }

func (s *CallParameterContext) SetOp(v antlr.Token) { s.op = v }

func (s *CallParameterContext) Expr() IExprContext {
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

func (s *CallParameterContext) ID() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserID, 0)
}

func (s *CallParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitCallParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) CallParameter() (localctx ICallParameterContext) {
	localctx = NewCallParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, TSParser_rulesParserRULE_callParameter)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(333)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(331)
			p.Match(TSParser_rulesParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)
			p.Match(TSParser_rulesParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TSParser_rulesParserT__29 {
		{
			p.SetState(335)

			var _m = p.Match(TSParser_rulesParserT__29)

			localctx.(*CallParameterContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(338)
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

// ITsfunctionsContext is an interface to support dynamic dispatch.
type ITsfunctionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTsfunctionsContext differentiates from other interfaces.
	IsTsfunctionsContext()
}

type TsfunctionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTsfunctionsContext() *TsfunctionsContext {
	var p = new(TsfunctionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_tsfunctions
	return p
}

func InitEmptyTsfunctionsContext(p *TsfunctionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSParser_rulesParserRULE_tsfunctions
}

func (*TsfunctionsContext) IsTsfunctionsContext() {}

func NewTsfunctionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TsfunctionsContext {
	var p = new(TsfunctionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSParser_rulesParserRULE_tsfunctions

	return p
}

func (s *TsfunctionsContext) GetParser() antlr.Parser { return s.parser }

func (s *TsfunctionsContext) CopyAll(ctx *TsfunctionsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TsfunctionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TsfunctionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ConvertStringContext struct {
	TsfunctionsContext
}

func NewConvertStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConvertStringContext {
	var p = new(ConvertStringContext)

	InitEmptyTsfunctionsContext(&p.TsfunctionsContext)
	p.parser = parser
	p.CopyAll(ctx.(*TsfunctionsContext))

	return p
}

func (s *ConvertStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConvertStringContext) STRING() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserSTRING, 0)
}

func (s *ConvertStringContext) Expr() IExprContext {
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

func (s *ConvertStringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitConvertString(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConvertIntContext struct {
	TsfunctionsContext
}

func NewConvertIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConvertIntContext {
	var p = new(ConvertIntContext)

	InitEmptyTsfunctionsContext(&p.TsfunctionsContext)
	p.parser = parser
	p.CopyAll(ctx.(*TsfunctionsContext))

	return p
}

func (s *ConvertIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConvertIntContext) INT() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserINT, 0)
}

func (s *ConvertIntContext) Expr() IExprContext {
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

func (s *ConvertIntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitConvertInt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConvertFloatContext struct {
	TsfunctionsContext
}

func NewConvertFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConvertFloatContext {
	var p = new(ConvertFloatContext)

	InitEmptyTsfunctionsContext(&p.TsfunctionsContext)
	p.parser = parser
	p.CopyAll(ctx.(*TsfunctionsContext))

	return p
}

func (s *ConvertFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConvertFloatContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(TSParser_rulesParserFLOAT, 0)
}

func (s *ConvertFloatContext) Expr() IExprContext {
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

func (s *ConvertFloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSParser_rulesVisitor:
		return t.VisitConvertFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSParser_rulesParser) Tsfunctions() (localctx ITsfunctionsContext) {
	localctx = NewTsfunctionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, TSParser_rulesParserRULE_tsfunctions)
	p.SetState(355)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TSParser_rulesParserSTRING:
		localctx = NewConvertStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(340)
			p.Match(TSParser_rulesParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(341)
			p.Match(TSParser_rulesParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(342)
			p.expr(0)
		}
		{
			p.SetState(343)
			p.Match(TSParser_rulesParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TSParser_rulesParserINT:
		localctx = NewConvertIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(345)
			p.Match(TSParser_rulesParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(346)
			p.Match(TSParser_rulesParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(347)
			p.expr(0)
		}
		{
			p.SetState(348)
			p.Match(TSParser_rulesParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TSParser_rulesParserFLOAT:
		localctx = NewConvertFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(350)
			p.Match(TSParser_rulesParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(351)
			p.Match(TSParser_rulesParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(352)
			p.expr(0)
		}
		{
			p.SetState(353)
			p.Match(TSParser_rulesParserT__13)
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

func (p *TSParser_rulesParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TSParser_rulesParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
