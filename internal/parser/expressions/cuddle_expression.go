package expressions

import (
	"cheese-lang/internal/parser"
)

// CuddleExpression is the basic loop control flow of the language

type CuddleExpression struct {
	codeInside CodeExpression
}

func (ce *CuddleExpression) Evaluate(globalContext *parser.Context, localContext *parser.Context) (ExpressionResult, error) {

	for {
		result, err := ce.codeInside.Evaluate(globalContext, localContext)

		if err != nil {
			return NullExpressionResult, err
		}

		if result.Return != nil {
			return result, nil
		}

		if result.Brake != nil {
			return result, nil
		}
	}

}
