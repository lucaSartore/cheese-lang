package expressions

import (
	"cheese-lang/internal/parser"
)

// CuddleExpression is the basic loop control flow of the language

type BrakeExpression struct {
	CodeInside parser.Expression
}

func (ce *BrakeExpression) Evaluate(globalContext *parser.Context, localContext *parser.Context) (parser.ExpressionResult, error) {
	v := parser.VoidExpressionResult
	v.Brake = true
	return v, nil
}
