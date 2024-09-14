package tokenizer

type TokenType int

const (
	NewLineToken TokenType = -2
	NullToken TokenType = -1

	Identifier TokenType = 0

	ParmesanType   TokenType = 1
	GorgonzolaType TokenType = 2
	MozzarellaType TokenType = 3
	MilkType       TokenType = 4
	RicottaType    TokenType = 5

	TasteKeyword   TokenType = 6
	RecipeKeyword  TokenType = 7
	PrepareKeyword TokenType = 8
	CurdleKeyword  TokenType = 9
	DrainKeyword   TokenType = 10

	AssignOperator       TokenType = 11
	EqualOperator        TokenType = 12
	UnEqualOperator      TokenType = 13
	AddOperator          TokenType = 14
	SubOperator          TokenType = 15
	DivOperator          TokenType = 16
	MulOperator          TokenType = 17
	ModOperator          TokenType = 18
	GreaterOperator      TokenType = 19
	LessOperator         TokenType = 20
	GreaterEqualOperator TokenType = 21
	LessEqualOperator    TokenType = 22
	NotOperator          TokenType = 23
	AndOperator          TokenType = 24
	OrOperator           TokenType = 25
	ExorOperator         TokenType = 26

	Comment           TokenType = 27
	SemiColon         TokenType = 28
	Comma             TokenType = 29
	OpenBracket       TokenType = 30
	CloseBracket      TokenType = 31
	OpenCurlyBracket  TokenType = 32
	CloseCurlyBracket TokenType = 33
	Arrow             TokenType = 34

	ParmesanLiteral   TokenType = 35
	GorgonzolaLiteral TokenType = 36
	MozzarellaLiteral TokenType = 37
	SpoiledMilk       TokenType = 38
	FreshMilk         TokenType = 39
)

func (t TokenType) String() string {
	switch t {
	case ParmesanType:
		return "ParmesanType"
	case GorgonzolaType:
		return "GorgonzolaType"
	case MozzarellaType:
		return "MozzarellaType"
	case MilkType:
		return "MilkType"
	case RicottaType:
		return "RicottaType"
	case TasteKeyword:
		return "TasteKeyword"
	case RecipeKeyword:
		return "RecipeKeyword"
	case PrepareKeyword:
		return "PrepareKeyword"
	case CurdleKeyword:
		return "CurdleKeyword"
	case DrainKeyword:
		return "DrainKeyword"
	case AssignOperator:
		return "AssignOperator"
	case EqualOperator:
		return "EqualOperator"
	case UnEqualOperator:
		return "UnEqualOperator"
	case AddOperator:
		return "AddOperator"
	case SubOperator:
		return "SubOperator"
	case DivOperator:
		return "DivOperator"
	case MulOperator:
		return "MulOperator"
	case ModOperator:
		return "ModOperator"
	case GreaterOperator:
		return "GreaterOperator"
	case LessOperator:
		return "LessOperator"
	case GreaterEqualOperator:
		return "GreaterEqualOperator"
	case LessEqualOperator:
		return "LessEqualOperator"
	case NotOperator:
		return "NotOperator"
	case AndOperator:
		return "AndOperator"
	case OrOperator:
		return "OrOperator"
	case ExorOperator:
		return "ExorOperator"
	case Comment:
		return "Comment"
	case SemiColon:
		return "SemiColon"
	case Comma:
		return "Comma"
	case OpenBracket:
		return "OpenBracket"
	case CloseBracket:
		return "CloseBracket"
	case OpenCurlyBracket:
		return "OpenCurlyBracket"
	case CloseCurlyBracket:
		return "CloseCurlyBracket"
	case Arrow:
		return "Arrow"
	case ParmesanLiteral:
		return "ParmesanLiteral"
	case GorgonzolaLiteral:
		return "GorgonzolaLiteral"
	case MozzarellaLiteral:
		return "MozzarellaLiteral"
	default:
		return "Unknown"
	}
}

type Token struct {
	TokenType TokenType
	TokenVal  string
    Line uint
    Colum uint
}

func MakeToken(t TokenType) Token {
	return Token{
		t,
		"",
        0,
        0,
	}
}

func MakeTokenWithMessage(t TokenType, v string) Token {
	return Token{
		t,
		v,
        0,
        0,
	}
}
