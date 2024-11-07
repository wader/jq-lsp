// Code generated by goyacc -o parser.go parser.go.y. DO NOT EDIT.

//line parser.go.y:2
package gojqparser

import __yyfmt__ "fmt"

//line parser.go.y:2

func reverseFuncDef(xs []*FuncDef) []*FuncDef {
	for i, j := 0, len(xs)-1; i < j; i, j = i+1, j-1 {
		xs[i], xs[j] = xs[j], xs[i]
	}
	return xs
}

func prependFuncDef(xs []*FuncDef, x *FuncDef) []*FuncDef {
	xs = append(xs, nil)
	copy(xs[1:], xs)
	xs[0] = x
	return xs
}

//line parser.go.y:19
type yySymType struct {
	yys      int
	value    any
	token    *Token
	operator Operator
}

const tokAltOp = 57346
const tokUpdateOp = 57347
const tokDestAltOp = 57348
const tokCompareOp = 57349
const tokOrOp = 57350
const tokAndOp = 57351
const tokModule = 57352
const tokImport = 57353
const tokInclude = 57354
const tokDef = 57355
const tokAs = 57356
const tokLabel = 57357
const tokBreak = 57358
const tokNull = 57359
const tokTrue = 57360
const tokFalse = 57361
const tokIf = 57362
const tokThen = 57363
const tokElif = 57364
const tokElse = 57365
const tokEnd = 57366
const tokTry = 57367
const tokCatch = 57368
const tokReduce = 57369
const tokForeach = 57370
const tokIdent = 57371
const tokVariable = 57372
const tokModuleIdent = 57373
const tokModuleVariable = 57374
const tokRecurse = 57375
const tokIndex = 57376
const tokNumber = 57377
const tokFormat = 57378
const tokString = 57379
const tokStringStart = 57380
const tokStringQuery = 57381
const tokStringEnd = 57382
const tokInvalid = 57383
const tokInvalidEscapeSequence = 57384
const tokUnterminatedString = 57385
const tokFuncDefQuery = 57386
const tokExpr = 57387
const tokTerm = 57388
const tokEmptyCatch = 57389

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"tokAltOp",
	"tokUpdateOp",
	"tokDestAltOp",
	"tokCompareOp",
	"tokOrOp",
	"tokAndOp",
	"tokModule",
	"tokImport",
	"tokInclude",
	"tokDef",
	"tokAs",
	"tokLabel",
	"tokBreak",
	"tokNull",
	"tokTrue",
	"tokFalse",
	"tokIf",
	"tokThen",
	"tokElif",
	"tokElse",
	"tokEnd",
	"tokTry",
	"tokCatch",
	"tokReduce",
	"tokForeach",
	"tokIdent",
	"tokVariable",
	"tokModuleIdent",
	"tokModuleVariable",
	"tokRecurse",
	"tokIndex",
	"tokNumber",
	"tokFormat",
	"tokString",
	"tokStringStart",
	"tokStringQuery",
	"tokStringEnd",
	"tokInvalid",
	"tokInvalidEscapeSequence",
	"tokUnterminatedString",
	"tokFuncDefQuery",
	"tokExpr",
	"tokTerm",
	"'|'",
	"','",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'.'",
	"'?'",
	"tokEmptyCatch",
	"'['",
	"';'",
	"':'",
	"'('",
	"')'",
	"']'",
	"'{'",
	"'}'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:677

//line yacctab:1
var yyExca = [...]int16{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 148,
	5, 0,
	-2, 30,
	-1, 151,
	7, 0,
	-2, 33,
	-1, 202,
	59, 117,
	-2, 52,
}

const yyPrivate = 57344

const yyLast = 785

var yyAct = [...]int16{
	78, 137, 189, 178, 106, 10, 214, 198, 134, 179,
	111, 105, 48, 5, 81, 232, 32, 50, 138, 6,
	157, 158, 247, 14, 183, 184, 185, 250, 113, 241,
	162, 231, 240, 100, 118, 101, 16, 73, 74, 124,
	127, 249, 181, 107, 182, 139, 161, 211, 117, 204,
	210, 140, 203, 115, 116, 246, 126, 102, 120, 120,
	120, 183, 184, 185, 186, 268, 73, 74, 258, 223,
	6, 119, 121, 122, 131, 132, 244, 283, 251, 181,
	282, 182, 73, 74, 234, 144, 135, 73, 74, 73,
	74, 230, 73, 74, 73, 74, 233, 257, 141, 192,
	142, 186, 238, 160, 166, 229, 285, 6, 281, 165,
	120, 120, 120, 120, 120, 120, 120, 120, 120, 120,
	164, 73, 74, 147, 148, 149, 150, 151, 152, 153,
	154, 155, 156, 187, 188, 256, 73, 74, 50, 163,
	177, 196, 73, 74, 130, 199, 205, 206, 73, 74,
	248, 129, 73, 74, 73, 74, 217, 200, 207, 73,
	74, 209, 172, 276, 45, 274, 73, 74, 218, 128,
	273, 216, 220, 221, 213, 79, 222, 239, 243, 107,
	146, 86, 87, 76, 90, 88, 89, 120, 120, 75,
	225, 171, 194, 120, 227, 80, 228, 123, 226, 135,
	215, 215, 235, 43, 44, 237, 219, 91, 92, 93,
	94, 95, 242, 43, 44, 83, 82, 85, 84, 93,
	94, 95, 80, 84, 190, 191, 91, 92, 93, 94,
	95, 277, 252, 193, 167, 254, 255, 199, 73, 74,
	253, 169, 83, 82, 259, 84, 270, 265, 266, 200,
	260, 261, 133, 174, 267, 175, 173, 73, 74, 269,
	73, 74, 97, 98, 271, 272, 90, 120, 120, 99,
	275, 263, 264, 3, 278, 279, 96, 90, 280, 89,
	215, 215, 25, 24, 284, 52, 53, 54, 55, 56,
	57, 58, 59, 60, 61, 62, 63, 64, 65, 66,
	67, 68, 69, 70, 71, 72, 109, 110, 91, 92,
	93, 94, 95, 114, 43, 44, 224, 180, 13, 91,
	92, 93, 94, 95, 9, 47, 104, 13, 17, 168,
	15, 37, 21, 22, 23, 33, 262, 108, 77, 245,
	34, 212, 35, 36, 39, 41, 40, 42, 19, 20,
	28, 31, 43, 44, 159, 125, 197, 86, 87, 195,
	90, 88, 89, 136, 29, 30, 208, 170, 7, 18,
	8, 4, 27, 2, 145, 38, 1, 143, 26, 52,
	53, 54, 55, 56, 57, 58, 59, 60, 61, 62,
	63, 64, 65, 66, 67, 68, 69, 70, 71, 72,
	109, 110, 91, 92, 93, 94, 95, 0, 43, 44,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	11, 12, 17, 0, 15, 37, 21, 22, 23, 33,
	0, 108, 0, 0, 34, 103, 35, 36, 39, 41,
	40, 42, 19, 20, 28, 31, 43, 44, 0, 0,
	0, 0, 86, 87, 0, 90, 88, 89, 29, 30,
	90, 88, 89, 18, 0, 0, 27, 0, 0, 38,
	0, 17, 26, 15, 37, 21, 22, 23, 33, 0,
	0, 0, 0, 34, 0, 35, 36, 39, 41, 40,
	42, 19, 20, 28, 31, 43, 44, 91, 92, 93,
	94, 95, 91, 92, 93, 94, 95, 29, 30, 0,
	0, 0, 18, 0, 0, 27, 0, 0, 38, 0,
	236, 26, 17, 0, 15, 37, 21, 22, 23, 33,
	0, 0, 0, 0, 34, 0, 35, 36, 39, 41,
	40, 42, 19, 20, 28, 31, 43, 44, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 29, 30,
	0, 0, 0, 18, 0, 0, 27, 0, 0, 38,
	0, 112, 26, 17, 0, 15, 37, 21, 22, 23,
	33, 0, 0, 0, 0, 34, 0, 35, 36, 39,
	41, 40, 42, 19, 20, 28, 31, 43, 44, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 29,
	30, 0, 0, 0, 18, 0, 0, 27, 0, 0,
	38, 0, 0, 26, 52, 53, 54, 55, 56, 57,
	58, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	68, 69, 70, 71, 72, 49, 0, 0, 0, 0,
	0, 0, 0, 51, 52, 53, 54, 55, 56, 57,
	58, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	68, 69, 70, 71, 72, 49, 0, 0, 0, 0,
	176, 0, 0, 51, 37, 21, 22, 23, 33, 0,
	0, 0, 0, 34, 0, 35, 36, 39, 41, 40,
	42, 19, 20, 28, 31, 43, 44, 0, 0, 0,
	46, 0, 0, 0, 0, 0, 0, 29, 30, 0,
	0, 0, 18, 0, 0, 27, 0, 0, 38, 0,
	0, 26, 52, 53, 54, 55, 56, 57, 58, 59,
	60, 61, 62, 63, 64, 65, 66, 67, 68, 69,
	70, 71, 72, 109, 202, 0, 0, 0, 0, 0,
	0, 43, 44, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 201,
}

var yyPact = [...]int16{
	263, -1000, -1000, -44, 409, 106, 646, -1000, -1000, -1000,
	191, 152, 146, 560, 161, 187, 448, 233, 166, -1000,
	-1000, -1000, -1000, -1000, -3, -1000, 371, 509, -1000, 668,
	668, 176, -1000, 560, 668, 668, 668, 167, 560, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -8, -1000, 110,
	92, 85, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 560, 560, 238, -44, -1000, 191, -12,
	-1000, -1000, -1000, 166, 315, 133, 668, 668, 668, 668,
	668, 668, 668, 668, 668, 668, -39, -1000, -1000, -1000,
	-1000, -1000, 560, -1000, -18, -1000, 80, 61, 560, -1000,
	-1000, -1000, -1000, 42, 560, 188, 188, -1000, 213, 215,
	188, 353, 177, -1000, 101, 216, -1000, 616, 44, 44,
	44, 191, -1000, 195, 41, -1000, 186, -1000, -1000, -12,
	724, -1000, -1000, -1000, -10, 560, 560, 448, 453, 270,
	259, 158, 168, 168, -1000, -1000, -1000, 560, 195, -11,
	191, -1000, 277, 668, 668, 95, -1000, 560, -1000, 668,
	-12, -12, -1000, -1000, -1000, 560, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 7, -1000, -1000, -44,
	-1000, -1000, -1000, 560, -12, 43, -1000, -33, -1000, 37,
	25, 560, -1000, -1000, 458, 40, 191, 119, -29, -1000,
	-1000, 560, -1000, -1000, 131, 448, 131, 17, 191, -1000,
	-5, -38, 89, -1000, -21, -1000, 20, 191, -1000, -1000,
	-12, -1000, 724, -12, -12, 74, -1000, 35, -1000, -1000,
	9, 195, 191, 668, 668, 249, 560, 560, -1000, -1000,
	44, -1000, -1000, -1000, -1000, -1000, 6, -1000, 560, -1000,
	131, 131, 222, 560, 560, 112, 107, -1000, -12, 105,
	-1000, 210, 191, 560, 560, -1000, -1000, 560, 47, 19,
	191, -1000, -1000, 560, 45, -1000,
}

var yyPgo = [...]int16{
	0, 376, 373, 371, 370, 8, 368, 324, 313, 366,
	0, 363, 1, 359, 356, 7, 36, 23, 16, 355,
	14, 354, 339, 336, 329, 326, 11, 6, 3, 9,
	325, 12, 317, 316, 2, 283, 282, 10, 4, 276,
}

var yyR1 = [...]int8{
	0, 1, 2, 2, 3, 3, 4, 4, 5, 5,
	6, 6, 7, 7, 8, 8, 9, 9, 34, 34,
	39, 39, 39, 10, 10, 10, 10, 10, 10, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	11, 11, 12, 12, 12, 13, 13, 14, 14, 15,
	15, 15, 15, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 18, 18, 19,
	19, 19, 35, 35, 36, 36, 20, 20, 20, 20,
	20, 21, 21, 22, 22, 23, 23, 24, 24, 25,
	25, 26, 26, 26, 26, 26, 38, 38, 38, 27,
	27, 28, 28, 28, 28, 28, 28, 28, 29, 29,
	29, 30, 30, 31, 31, 31, 32, 32, 33, 33,
	37, 37, 37, 37, 37, 37, 37, 37, 37, 37,
	37, 37, 37, 37, 37, 37, 37, 37, 37, 37,
	37,
}

var yyR2 = [...]int8{
	0, 3, 0, 3, 0, 2, 6, 4, 0, 1,
	1, 1, 0, 2, 5, 8, 1, 3, 1, 1,
	1, 1, 1, 2, 3, 5, 4, 3, 1, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 1,
	1, 3, 1, 3, 3, 1, 3, 1, 3, 3,
	3, 5, 1, 1, 1, 1, 2, 2, 1, 1,
	1, 1, 4, 1, 2, 3, 4, 2, 3, 1,
	2, 2, 1, 2, 1, 7, 3, 9, 9, 11,
	2, 3, 2, 2, 2, 3, 3, 1, 3, 0,
	2, 4, 1, 1, 1, 1, 2, 3, 4, 4,
	5, 1, 3, 0, 5, 0, 2, 0, 2, 1,
	3, 3, 3, 5, 1, 1, 1, 1, 1, 3,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 3,
	4, 1, 3, 3, 3, 3, 2, 3, 1, 3,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1,
}

var yyChk = [...]int16{
	-1000, -1, -2, 10, -3, -29, 63, -6, -4, -7,
	-10, 11, 12, -8, -17, 15, -16, 13, 54, 33,
	34, 17, 18, 19, -35, -36, 63, 57, 35, 49,
	50, 36, -18, 20, 25, 27, 28, 16, 60, 29,
	31, 30, 32, 37, 38, 58, 64, -30, -31, 29,
	-37, 37, 8, 9, 10, 11, 12, 13, 14, 15,
	16, 17, 18, 19, 20, 21, 22, 23, 24, 25,
	26, 27, 28, 47, 48, 37, 37, -7, -10, 14,
	34, -20, 55, 54, 57, 30, 4, 5, 8, 9,
	7, 49, 50, 51, 52, 53, -39, 29, 30, 36,
	-20, -18, 60, 64, -25, -26, -38, -18, 60, 29,
	30, -37, 62, -10, -8, -17, -17, -18, -10, -16,
	-17, -16, -16, 30, -10, -19, 64, 48, 59, 59,
	59, -10, -10, 14, -5, -29, -11, -12, 30, 57,
	63, -20, -18, 62, -10, 59, 47, -16, -16, -16,
	-16, -16, -16, -16, -16, -16, -16, 59, 60, -21,
	-10, 64, 48, 59, 59, -10, 62, 21, -24, 26,
	14, 14, 61, 40, 37, 39, 64, -31, -28, -29,
	-32, 35, 37, 17, 18, 19, 57, -28, -28, -34,
	29, 30, 58, 47, 6, -13, -12, -14, -15, -38,
	-18, 60, 30, 62, 59, -10, -10, -10, -9, -34,
	61, 58, 64, -26, -27, -16, -27, 61, -10, -16,
	-12, -12, -10, 62, -33, -28, -5, -10, -12, 62,
	48, 64, 48, 59, 59, -10, 62, -10, 62, 58,
	61, 58, -10, 47, 59, -22, 60, 60, 61, 62,
	48, 58, -12, -15, -12, -12, 61, 62, 59, -34,
	-27, -27, -23, 22, 23, -10, -10, -28, 59, -10,
	24, -10, -10, 58, 58, -12, 58, 21, -10, -10,
	-10, 61, 61, 58, -10, 61,
}

var yyDef = [...]int16{
	2, -2, 4, 0, 12, 0, 0, 1, 5, 10,
	11, 0, 0, 12, 39, 0, 28, 0, 53, 54,
	55, 58, 59, 60, 61, 63, 0, 0, 69, 0,
	0, 72, 74, 0, 0, 0, 0, 0, 0, 92,
	93, 94, 95, 87, 89, 3, 128, 0, 131, 0,
	0, 0, 140, 141, 142, 143, 144, 145, 146, 147,
	148, 149, 150, 151, 152, 153, 154, 155, 156, 157,
	158, 159, 160, 0, 0, 0, 8, 13, 23, 0,
	82, 83, 84, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 20, 21, 22,
	56, 57, 0, 64, 0, 109, 114, 115, 0, 116,
	117, 118, 67, 0, 0, 70, 71, 73, 0, 107,
	39, 0, 0, 80, 0, 0, 129, 0, 0, 0,
	0, 24, 27, 0, 0, 9, 0, 40, 42, 0,
	0, 85, 86, 96, 0, 0, 0, 29, -2, 31,
	32, -2, 34, 35, 36, 37, 38, 0, 0, 0,
	101, 65, 0, 0, 0, 0, 68, 0, 76, 0,
	0, 0, 81, 88, 90, 0, 130, 132, 133, 121,
	122, 123, 124, 125, 126, 127, 0, 134, 135, 8,
	18, 19, 7, 0, 0, 0, 45, 0, 47, 0,
	0, 0, -2, 97, 0, 0, 26, 0, 0, 16,
	62, 0, 66, 110, 111, 120, 112, 0, 103, 108,
	0, 0, 0, 136, 0, 138, 0, 25, 41, 43,
	0, 44, 0, 0, 0, 0, 98, 0, 99, 14,
	0, 0, 102, 0, 0, 105, 0, 0, 91, 137,
	0, 6, 46, 48, 49, 50, 0, 100, 0, 17,
	119, 113, 0, 0, 0, 0, 0, 139, 0, 0,
	75, 0, 106, 0, 0, 51, 15, 0, 0, 0,
	104, 77, 78, 0, 0, 79,
}

var yyTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 53, 3, 3,
	60, 61, 51, 49, 48, 50, 54, 52, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 59, 58,
	3, 3, 3, 55, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 57, 3, 62, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 63, 47, 64,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 56,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:59
		{
			query := yyDollar[3].value.(*Query)
			query.Meta = yyDollar[1].value.(*ConstObject)
			query.Imports = yyDollar[2].value.([]*Import)
			yylex.(*lexer).result = query
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:68
		{
			yyVAL.value = (*ConstObject)(nil)
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:72
		{
			yyVAL.value = yyDollar[2].value
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:78
		{
			yyVAL.value = []*Import(nil)
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:82
		{
			yyVAL.value = append(yyDollar[1].value.([]*Import), yyDollar[2].value.(*Import))
		}
	case 6:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.go.y:88
		{
			yyVAL.value = &Import{ImportPath: yyDollar[2].token, ImportAlias: yyDollar[4].token, Meta: yyDollar[5].value.(*ConstObject)}
		}
	case 7:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:92
		{
			yyVAL.value = &Import{IncludePath: yyDollar[2].token, Meta: yyDollar[3].value.(*ConstObject)}
		}
	case 8:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:98
		{
			yyVAL.value = (*ConstObject)(nil)
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:105
		{
			yyVAL.value = &Query{FuncDefs: reverseFuncDef(yyDollar[1].value.([]*FuncDef))}
		}
	case 12:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:112
		{
			yyVAL.value = []*FuncDef(nil)
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:116
		{
			yyVAL.value = append(yyDollar[2].value.([]*FuncDef), yyDollar[1].value.(*FuncDef))
		}
	case 14:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:122
		{
			yyVAL.value = &FuncDef{Name: yyDollar[2].token, Body: yyDollar[4].value.(*Query)}
		}
	case 15:
		yyDollar = yyS[yypt-8 : yypt+1]
//line parser.go.y:126
		{
			yyVAL.value = &FuncDef{yyDollar[2].token, yyDollar[4].value.([]*Token), yyDollar[7].value.(*Query)}
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:132
		{
			yyVAL.value = []*Token{yyDollar[1].token}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:136
		{
			yyVAL.value = append(yyDollar[1].value.([]*Token), yyDollar[3].token)
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:151
		{
			query := yyDollar[2].value.(*Query)
			query.FuncDefs = prependFuncDef(query.FuncDefs, yyDollar[1].value.(*FuncDef))
			yyVAL.value = query
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:157
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpPipe, Right: yyDollar[3].value.(*Query)}
		}
	case 25:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:161
		{
			term := yyDollar[1].value.(*Term)
			term.SuffixList = append(term.SuffixList, &Suffix{Bind: &Bind{yyDollar[3].value.([]*Pattern), yyDollar[5].value.(*Query)}})
			yyVAL.value = &Query{Term: term}
		}
	case 26:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:167
		{
			yyVAL.value = &Query{Term: &Term{Type: TermTypeLabel, Label: &Label{yyDollar[2].token, yyDollar[4].value.(*Query)}}}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:171
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpComma, Right: yyDollar[3].value.(*Query)}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:178
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: yyDollar[2].operator, Right: yyDollar[3].value.(*Query)}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:182
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: yyDollar[2].operator, Right: yyDollar[3].value.(*Query)}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:186
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpOr, Right: yyDollar[3].value.(*Query)}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:190
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpAnd, Right: yyDollar[3].value.(*Query)}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:194
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: yyDollar[2].operator, Right: yyDollar[3].value.(*Query)}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:198
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpAdd, Right: yyDollar[3].value.(*Query)}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:202
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpSub, Right: yyDollar[3].value.(*Query)}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:206
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpMul, Right: yyDollar[3].value.(*Query)}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:210
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpDiv, Right: yyDollar[3].value.(*Query)}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:214
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpMod, Right: yyDollar[3].value.(*Query)}
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:218
		{
			yyVAL.value = &Query{Term: yyDollar[1].value.(*Term)}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:224
		{
			yyVAL.value = []*Pattern{yyDollar[1].value.(*Pattern)}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:228
		{
			yyVAL.value = append(yyDollar[1].value.([]*Pattern), yyDollar[3].value.(*Pattern))
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:234
		{
			yyVAL.value = &Pattern{Name: yyDollar[1].token}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:238
		{
			yyVAL.value = &Pattern{Array: yyDollar[2].value.([]*Pattern)}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:242
		{
			yyVAL.value = &Pattern{Object: yyDollar[2].value.([]*PatternObject)}
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:248
		{
			yyVAL.value = []*Pattern{yyDollar[1].value.(*Pattern)}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:252
		{
			yyVAL.value = append(yyDollar[1].value.([]*Pattern), yyDollar[3].value.(*Pattern))
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:258
		{
			yyVAL.value = []*PatternObject{yyDollar[1].value.(*PatternObject)}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:262
		{
			yyVAL.value = append(yyDollar[1].value.([]*PatternObject), yyDollar[3].value.(*PatternObject))
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:268
		{
			yyVAL.value = &PatternObject{Key: yyDollar[1].token, Val: yyDollar[3].value.(*Pattern)}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:272
		{
			yyVAL.value = &PatternObject{KeyString: yyDollar[1].value.(*String), Val: yyDollar[3].value.(*Pattern)}
		}
	case 51:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:276
		{
			yyVAL.value = &PatternObject{KeyQuery: yyDollar[2].value.(*Query), Val: yyDollar[5].value.(*Pattern)}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:280
		{
			yyVAL.value = &PatternObject{Key: yyDollar[1].token}
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:286
		{
			yyVAL.value = &Term{Type: TermTypeIdentity}
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:290
		{
			yyVAL.value = &Term{Type: TermTypeRecurse}
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:294
		{
			yyVAL.value = &Term{Type: TermTypeIndex, Index: &Index{Name: yyDollar[1].token}}
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:298
		{
			suffix := yyDollar[2].value.(*Suffix)
			if suffix.Iter {
				yyVAL.value = &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{suffix}}
			} else {
				yyVAL.value = &Term{Type: TermTypeIndex, Index: suffix.Index}
			}
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:307
		{
			yyVAL.value = &Term{Type: TermTypeIndex, Index: &Index{Str: yyDollar[2].value.(*String)}}
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:311
		{
			yyVAL.value = &Term{Type: TermTypeNull}
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:315
		{
			yyVAL.value = &Term{Type: TermTypeTrue}
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:319
		{
			yyVAL.value = &Term{Type: TermTypeFalse}
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:323
		{
			yyVAL.value = &Term{Type: TermTypeFunc, Func: &Func{Name: yyDollar[1].token}}
		}
	case 62:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:327
		{
			yyVAL.value = &Term{Type: TermTypeFunc, Func: &Func{Name: yyDollar[1].token, Args: yyDollar[3].value.([]*Query)}}
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:331
		{
			yyVAL.value = &Term{Type: TermTypeFunc, Func: &Func{Name: yyDollar[1].token}}
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:335
		{
			yyVAL.value = &Term{Type: TermTypeObject, Object: &Object{}}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:339
		{
			yyVAL.value = &Term{Type: TermTypeObject, Object: &Object{yyDollar[2].value.([]*ObjectKeyVal)}}
		}
	case 66:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:343
		{
			yyVAL.value = &Term{Type: TermTypeObject, Object: &Object{yyDollar[2].value.([]*ObjectKeyVal)}}
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:347
		{
			yyVAL.value = &Term{Type: TermTypeArray, Array: &Array{}}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:351
		{
			yyVAL.value = &Term{Type: TermTypeArray, Array: &Array{yyDollar[2].value.(*Query)}}
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:355
		{
			yyVAL.value = &Term{Type: TermTypeNumber, Number: yyDollar[1].token}
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:359
		{
			yyVAL.value = &Term{Type: TermTypeUnary, Unary: &Unary{OpAdd, yyDollar[2].value.(*Term)}}
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:363
		{
			yyVAL.value = &Term{Type: TermTypeUnary, Unary: &Unary{OpSub, yyDollar[2].value.(*Term)}}
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:367
		{
			yyVAL.value = &Term{Type: TermTypeFormat, Format: yyDollar[1].token}
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:371
		{
			yyVAL.value = &Term{Type: TermTypeFormat, Format: yyDollar[1].token, Str: yyDollar[2].value.(*String)}
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:375
		{
			yyVAL.value = &Term{Type: TermTypeString, Str: yyDollar[1].value.(*String)}
		}
	case 75:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.go.y:379
		{
			yyVAL.value = &Term{Type: TermTypeIf, If: &If{yyDollar[2].value.(*Query), yyDollar[4].value.(*Query), yyDollar[5].value.([]*IfElif), yyDollar[6].value.(*Query)}}
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:383
		{
			yyVAL.value = &Term{Type: TermTypeTry, Try: &Try{yyDollar[2].value.(*Query), yyDollar[3].value.(*Query)}}
		}
	case 77:
		yyDollar = yyS[yypt-9 : yypt+1]
//line parser.go.y:387
		{
			yyVAL.value = &Term{Type: TermTypeReduce, Reduce: &Reduce{yyDollar[2].value.(*Query), yyDollar[4].value.(*Pattern), yyDollar[6].value.(*Query), yyDollar[8].value.(*Query)}}
		}
	case 78:
		yyDollar = yyS[yypt-9 : yypt+1]
//line parser.go.y:391
		{
			yyVAL.value = &Term{Type: TermTypeForeach, Foreach: &Foreach{yyDollar[2].value.(*Query), yyDollar[4].value.(*Pattern), yyDollar[6].value.(*Query), yyDollar[8].value.(*Query), nil}}
		}
	case 79:
		yyDollar = yyS[yypt-11 : yypt+1]
//line parser.go.y:395
		{
			yyVAL.value = &Term{Type: TermTypeForeach, Foreach: &Foreach{yyDollar[2].value.(*Query), yyDollar[4].value.(*Pattern), yyDollar[6].value.(*Query), yyDollar[8].value.(*Query), yyDollar[10].value.(*Query)}}
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:399
		{
			yyVAL.value = &Term{Type: TermTypeBreak, Break: yyDollar[2].token}
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:403
		{
			yyVAL.value = &Term{Type: TermTypeQuery, Query: yyDollar[2].value.(*Query)}
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:407
		{
			yyDollar[1].value.(*Term).SuffixList = append(yyDollar[1].value.(*Term).SuffixList, &Suffix{Index: &Index{Name: yyDollar[2].token}})
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:411
		{
			yyDollar[1].value.(*Term).SuffixList = append(yyDollar[1].value.(*Term).SuffixList, yyDollar[2].value.(*Suffix))
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:415
		{
			yyDollar[1].value.(*Term).SuffixList = append(yyDollar[1].value.(*Term).SuffixList, &Suffix{Optional: true})
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:419
		{
			yyDollar[1].value.(*Term).SuffixList = append(yyDollar[1].value.(*Term).SuffixList, yyDollar[3].value.(*Suffix))
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:423
		{
			yyDollar[1].value.(*Term).SuffixList = append(yyDollar[1].value.(*Term).SuffixList, &Suffix{Index: &Index{Str: yyDollar[3].value.(*String)}})
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:429
		{
			yyVAL.value = &String{Str: yyDollar[1].token}
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:433
		{
			yyVAL.value = &String{Queries: yyDollar[2].value.([]*Query)}
		}
	case 89:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:439
		{
			yyVAL.value = []*Query{}
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:443
		{
			yyVAL.value = append(yyDollar[1].value.([]*Query), &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: yyDollar[2].token}}})
		}
	case 91:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:447
		{
			yylex.(*lexer).inString = true
			yyVAL.value = append(yyDollar[1].value.([]*Query), &Query{Term: &Term{Type: TermTypeQuery, Query: yyDollar[3].value.(*Query)}})
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:462
		{
			yyVAL.value = &Suffix{Iter: true}
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:466
		{
			yyVAL.value = &Suffix{Index: &Index{Start: yyDollar[2].value.(*Query)}}
		}
	case 98:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:470
		{
			yyVAL.value = &Suffix{Index: &Index{Start: yyDollar[2].value.(*Query), IsSlice: true}}
		}
	case 99:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:474
		{
			yyVAL.value = &Suffix{Index: &Index{End: yyDollar[3].value.(*Query), IsSlice: true}}
		}
	case 100:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:478
		{
			yyVAL.value = &Suffix{Index: &Index{Start: yyDollar[2].value.(*Query), End: yyDollar[4].value.(*Query), IsSlice: true}}
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:484
		{
			yyVAL.value = []*Query{yyDollar[1].value.(*Query)}
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:488
		{
			yyVAL.value = append(yyDollar[1].value.([]*Query), yyDollar[3].value.(*Query))
		}
	case 103:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:494
		{
			yyVAL.value = []*IfElif(nil)
		}
	case 104:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:498
		{
			yyVAL.value = append(yyDollar[1].value.([]*IfElif), &IfElif{yyDollar[3].value.(*Query), yyDollar[5].value.(*Query)})
		}
	case 105:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:504
		{
			yyVAL.value = (*Query)(nil)
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:508
		{
			yyVAL.value = yyDollar[2].value
		}
	case 107:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:514
		{
			yyVAL.value = (*Query)(nil)
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:518
		{
			yyVAL.value = yyDollar[2].value
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:524
		{
			yyVAL.value = []*ObjectKeyVal{yyDollar[1].value.(*ObjectKeyVal)}
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:528
		{
			yyVAL.value = append(yyDollar[1].value.([]*ObjectKeyVal), yyDollar[3].value.(*ObjectKeyVal))
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:534
		{
			yyVAL.value = &ObjectKeyVal{Key: yyDollar[1].token, Val: yyDollar[3].value.(*Query)}
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:538
		{
			yyVAL.value = &ObjectKeyVal{KeyString: yyDollar[1].value.(*String), Val: yyDollar[3].value.(*Query)}
		}
	case 113:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:542
		{
			yyVAL.value = &ObjectKeyVal{KeyQuery: yyDollar[2].value.(*Query), Val: yyDollar[5].value.(*Query)}
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:546
		{
			yyVAL.value = &ObjectKeyVal{Key: yyDollar[1].token}
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:550
		{
			yyVAL.value = &ObjectKeyVal{KeyString: yyDollar[1].value.(*String)}
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:561
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpPipe, Right: yyDollar[3].value.(*Query)}
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:568
		{
			yyVAL.value = &ConstTerm{Object: yyDollar[1].value.(*ConstObject)}
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:572
		{
			yyVAL.value = &ConstTerm{Array: yyDollar[1].value.(*ConstArray)}
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:576
		{
			yyVAL.value = &ConstTerm{Number: yyDollar[1].token}
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:580
		{
			yyVAL.value = &ConstTerm{Str: yyDollar[1].token}
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:584
		{
			yyVAL.value = &ConstTerm{Null: true}
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:588
		{
			yyVAL.value = &ConstTerm{True: true}
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:592
		{
			yyVAL.value = &ConstTerm{False: true}
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:598
		{
			yyVAL.value = &ConstObject{}
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:602
		{
			yyVAL.value = &ConstObject{yyDollar[2].value.([]*ConstObjectKeyVal)}
		}
	case 130:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:606
		{
			yyVAL.value = &ConstObject{yyDollar[2].value.([]*ConstObjectKeyVal)}
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:612
		{
			yyVAL.value = []*ConstObjectKeyVal{yyDollar[1].value.(*ConstObjectKeyVal)}
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:616
		{
			yyVAL.value = append(yyDollar[1].value.([]*ConstObjectKeyVal), yyDollar[3].value.(*ConstObjectKeyVal))
		}
	case 133:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:622
		{
			yyVAL.value = &ConstObjectKeyVal{Key: yyDollar[1].token, Val: yyDollar[3].value.(*ConstTerm)}
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:626
		{
			yyVAL.value = &ConstObjectKeyVal{Key: yyDollar[1].token, Val: yyDollar[3].value.(*ConstTerm)}
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:630
		{
			yyVAL.value = &ConstObjectKeyVal{KeyString: yyDollar[1].token, Val: yyDollar[3].value.(*ConstTerm)}
		}
	case 136:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:636
		{
			yyVAL.value = &ConstArray{}
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:640
		{
			yyVAL.value = &ConstArray{yyDollar[2].value.([]*ConstTerm)}
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:646
		{
			yyVAL.value = []*ConstTerm{yyDollar[1].value.(*ConstTerm)}
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:650
		{
			yyVAL.value = append(yyDollar[1].value.([]*ConstTerm), yyDollar[3].value.(*ConstTerm))
		}
	}
	goto yystack /* stack new state and value */
}
