package expressions

import (
	"cheese-lang/internal/parser"
)

// CuddleExpression is the basic loop control flow of the language

type CuddleExpression struct {
	codeInside CodeExpression
}

func (ce *CuddleExpression) Evaluate(globalContext *parser.Context, localContext *parser.Context) (parser.ExpressionResult, error) {

	for {
		result, err := ce.codeInside.Evaluate(globalContext, localContext)

		if err != nil {
			return parser.NullExpressionResult, err
		}

		if result.Return {
			return result, nil
		}

		if result.Brake {
			// to avoid thai i endlessly brake around loops
			result.Brake = false
			return result, nil
		}
	}

}
