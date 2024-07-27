package expressions

import (
	"fmt"
)

type VariableDeclarationExpression struct {
	Type     VariableType
	Name     string
	ToAssign Expression
	Global   bool
}

func (exp *VariableDeclarationExpression) Evaluate(globalContext *Context, localContext *Context) (ExpressionResult, error) {

	var context *Context
	if exp.Global {
		context = globalContext
	} else {
		context = localContext
	}

	_, ok := context.GetVariable(exp.Name)
	if ok {
		return NullExpressionResult, fmt.Errorf("unable to declare variable %v as the name is already existing in this scope", exp.Name)
	}

	res, err := exp.ToAssign.Evaluate(globalContext, localContext)

	if err != nil {
		return NullExpressionResult, err
	}

	if res.Value.GetVariableType() != exp.Type {
		return NullExpressionResult, fmt.Errorf("expected %v got type %v in creation ov variable %v", exp.Type, res.Value.GetVariableType(), exp.Name)
	}

	context.AddVariable(MakeVariable(exp.Name, res.Value))
	return VoidExpressionResult, nil
}
