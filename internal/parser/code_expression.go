package parser

// a code expression is a list of expressions that are evaluated in order

type CodeExpression struct {
	Expressions []Expression
}
