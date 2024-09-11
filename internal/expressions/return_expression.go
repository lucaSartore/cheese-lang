package expressions

type ReturnExpression struct {
	Expressions []Expression
}

func (exp *ReturnExpression) Evaluate(globalContext *Context, localContext *Context) (ExpressionResult, error) {
    values := make([]VariableContainer,0)
    
    for _, expression := range exp.Expressions{
	    result, err := expression.Evaluate(globalContext, localContext)
        if err != nil {
            return result, err
        }
        if result.Value != nil {
            values = append(values, result.Value)
        }
    }

    if len(values) == 0 {
        return  ExpressionResult{Value: &RicottaVariable{}, Return: true, Brake: false}, nil
    }

    if len(values) == 1 {
        return  ExpressionResult{Value: values[0], Return: true, Brake: false}, nil
    }

    return ExpressionResult{Value: &TupleVariableType{Variables: values}, Return: true, Brake: false}, nil
}
