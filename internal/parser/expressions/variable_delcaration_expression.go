package expressions

import (
	"cheese-lang/internal/parser"
	"fmt"
)

type VariableDeclarationExpression struct {
	Type     parser.VariableType
	Name     string
	ToAssign parser.Expression
	Global   bool
}

func (exp *VariableDeclarationExpression) Evaluate(globalContext *parser.Context, localContext *parser.Context) (parser.ExpressionResult, error) {

	var context *parser.Context
	if exp.Global {
		context = globalContext
	} else {
		context = localContext
	}

	_, ok := context.GetVariable(exp.Name)
	if ok {
		return parser.NullExpressionResult, fmt.Errorf("unable to declare variable %v as the name is already existing in this scope", exp.Name)
	}

	res, err := exp.ToAssign.Evaluate(globalContext, localContext)

	if err != nil {
		return parser.NullExpressionResult, err
	}

	if res.Value.GetVariableType() != exp.Type {
		return parser.NullExpressionResult, fmt.Errorf("expected %v got type %v in creation ov variable %v", exp.Type, res.Value.GetVariableType(), exp.Name)
	}

	context.AddVariable(parser.MakeVariable(exp.Name, res.Value))
	return parser.VoidExpressionResult, nil
}
