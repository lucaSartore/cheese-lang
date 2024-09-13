package expressions

import (
    "fmt"
    "strconv"
)


type ActionFunction[T VariableContainer] func(T) (VariableContainer, error)

func Wrapper[T VariableContainer](action ActionFunction[T], context *Context, inputName string)(ExpressionResult, error){
    x, ok := context.GetVariable(inputName)
    // if everything is done correctly in the setup stage there should be no possibility to fail this 2 checks
    if !ok {
        panic("std  function assertion fail 1") 
    }
    cast, ok := x.Value.(T)
    if ! ok {
        panic("std function assertion fail 2")
    }
    res, err := action(cast) 
    if err != nil {
        return NullExpressionResult, err
    }
    return ExpressionResult{Value: res, Brake: false, Return: false}, nil
}

type parmesanToGorgonzola struct{}
func (fc *parmesanToGorgonzola) Evaluate(globalContext *Context, localContext *Context) (ExpressionResult, error) {
    return Wrapper(
        func(v *ParmesanVariable) (VariableContainer, error) {
            return &GorgonzolaVariable{Value: float64(v.Value)}, nil
        },
        localContext,
        "x",
    )
}
var parmesanToGorgonzolaFunc = Function{
    Name: "p_to_g",
    ArgumentsType: []VariableType{Parmesan},
    ArgumentsNames: []string{"x"},
    Code: &parmesanToGorgonzola{}, 
}



type parmesanToMozzarella struct{}
func (fc *parmesanToMozzarella ) Evaluate(globalContext *Context, localContext *Context) (ExpressionResult, error) {
    return Wrapper(
        func(v *ParmesanVariable) (VariableContainer, error) {
            return &MozzarellaVariable{Value: fmt.Sprintf("%v",v.Value)}, nil
        },
        localContext,
        "x",
    )
}
var parmesanToMozzarellaFunc = Function{
    Name: "p_to_m",
    ArgumentsType: []VariableType{Parmesan},
    ArgumentsNames: []string{"x"},
    Code: &parmesanToMozzarella{}, 
}


type gorgonzolaToParmesan struct{}
func (fc *gorgonzolaToParmesan ) Evaluate(globalContext *Context, localContext *Context) (ExpressionResult, error) {
    return Wrapper(
        func(v *GorgonzolaVariable) (VariableContainer, error) {
            return &ParmesanVariable{Value: int(v.Value)}, nil
        },
        localContext,
        "x",
    )
}
var gorgonzolaToParmesanFunc = Function{
    Name: "g_to_p",
    ArgumentsType: []VariableType{Gorgonzola},
    ArgumentsNames: []string{"x"},
    Code: &gorgonzolaToParmesan{}, 
}


type gorgonzolaToMozzarella struct{}
func (fc *gorgonzolaToMozzarella ) Evaluate(globalContext *Context, localContext *Context) (ExpressionResult, error) {
    return Wrapper(
        func(v *GorgonzolaVariable) (VariableContainer, error) {
            return &MozzarellaVariable{Value: fmt.Sprintf("%v",v.Value)}, nil
        },
        localContext,
        "x",
    )
}
var gorgonzolaToMozzarellaFunc = Function{
    Name: "g_to_m",
    ArgumentsType: []VariableType{Gorgonzola},
    ArgumentsNames: []string{"x"},
    Code: &gorgonzolaToMozzarella{}, 
}


type mozzarellaToParmesan struct{}
func (fc *mozzarellaToParmesan ) Evaluate(globalContext *Context, localContext *Context) (ExpressionResult, error) {
    return Wrapper(
        func(v *MozzarellaVariable) (VariableContainer, error) {
            val, err := strconv.Atoi(v.Value)
            success := err == nil
            return &TupleVariableType{
                Variables: []VariableContainer{
                    &ParmesanVariable{Value: val},
                    &MilkVariable{Value: success},
                },
            }, nil
        },
        localContext,
        "x",
    )
}
var mozzarellaToParmesanFunc = Function{
    Name: "m_to_p",
    ArgumentsType: []VariableType{Mozzarella},
    ArgumentsNames: []string{"x"},
    Code: &mozzarellaToParmesan{}, 
}

type mozzarellaToGorgonzola struct{}
func (fc *mozzarellaToGorgonzola ) Evaluate(globalContext *Context, localContext *Context) (ExpressionResult, error) {
    return Wrapper(
        func(v *MozzarellaVariable) (VariableContainer, error) {
            val, err := strconv.ParseFloat(v.Value, 64)
            success := err == nil
            return &TupleVariableType{
                Variables: []VariableContainer{
                    &GorgonzolaVariable{Value: val},
                    &MilkVariable{Value: success},
                },
            }, nil
        },
        localContext,
        "x",
    )
}
var mozzarellaToGorgonzolaFunc = Function{
    Name: "m_to_p",
    ArgumentsType: []VariableType{Mozzarella},
    ArgumentsNames: []string{"x"},
    Code: &mozzarellaToGorgonzola{}, 
}

var StandardLibrary = []Function{parmesanToGorgonzolaFunc,parmesanToMozzarellaFunc,gorgonzolaToParmesanFunc,gorgonzolaToMozzarellaFunc, mozzarellaToGorgonzolaFunc, mozzarellaToParmesanFunc}
