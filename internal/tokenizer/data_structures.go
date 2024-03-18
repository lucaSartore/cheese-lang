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
)

type Token struct {
	TokenType TokenType
	TokenVal  string
}
