package tokenizer

type TokenType int

const (
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

	Comment           TokenType = 23
	SemiColon         TokenType = 24
	Comma             TokenType = 25
	OpenBracket       TokenType = 26
	CloseBracket      TokenType = 27
	OpenCurlyBracket  TokenType = 28
	CloseCurlyBracket TokenType = 29
	Arrow             TokenType = 30
)

type Token struct {
	TokenType TokenType
	TokenVal  string
}
