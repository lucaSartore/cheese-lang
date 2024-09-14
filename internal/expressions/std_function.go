package expressions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


type ActionFunction[T VariableContainer] func(T) (VariableContainer, error)

func GetVariableOrPanic[T VariableContainer] (context *Context, name string) T {
    x, ok := context.GetVariable(name)
    // if everything is done correctly in the setup stage there should be no possibility to fail this 2 checks
    if !ok {
        panic("std  function assertion fail 1") 
    }
    cast, ok := x.Value.(T)
    if ! ok {
        panic("std function assertion fail 2")
    }
    return cast
}

func Wrapper[T VariableContainer](action ActionFunction[T], context *Context, inputName string)(ExpressionResult, error){
    variable := GetVariableOrPanic[T](context, inputName)
    res, err := action(variable) 
    if err != nil {
        return NullExpressionResult, err
    }
    return ExpressionResult{Value: res, Brake: false, Return: true}, nil
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
    Name: "m_to_g",
    ArgumentsType: []VariableType{Mozzarella},
    ArgumentsNames: []string{"x"},
    Code: &mozzarellaToGorgonzola{}, 
}


type serve struct{}
func (fc *serve ) Evaluate(globalContext *Context, localContext *Context) (ExpressionResult, error) {
    return Wrapper(
        func(v *MozzarellaVariable) (VariableContainer, error) {
            fmt.Print(v.Value)
            return &RicottaVariable{}, nil
        },
        localContext,
        "x",
    )
}
var serveFunction = Function{
    Name: "serve",
    ArgumentsType: []VariableType{Mozzarella},
    ArgumentsNames: []string{"x"},
    Code: &serve{}, 
}

var reader = bufio.NewReader(os.Stdin)

type eat struct{}
func (fc *eat ) Evaluate(globalContext *Context, localContext *Context) (ExpressionResult, error) {
    text, err := reader.ReadString('\n')   
    if err != nil {
        return NullExpressionResult, err
    }
    text = text[0:len(text)-1]
    if len(text) != 0 && text[len(text)-1] == '\r'{
        text = text[0:len(text)-1]
    }
    return ExpressionResult{Value: &MozzarellaVariable{Value: text},Return: true, Brake: false}, nil
}
var eatFunction = Function{
    Name: "eat",
    ArgumentsType: []VariableType{},
    ArgumentsNames: []string{},
    Code: &eat{}, 
}

type weight struct{}
func (fc *weight ) Evaluate(globalContext *Context, localContext *Context) (ExpressionResult, error) {
    return Wrapper(
        func(v *MozzarellaVariable) (VariableContainer, error) {
            return &ParmesanVariable{len(v.Value)}, nil
        },
        localContext,
        "x",
    )
}
var weightFunction = Function{
    Name: "weight",
    ArgumentsType: []VariableType{Mozzarella},
    ArgumentsNames: []string{"x"},
    Code: &weight{}, 
}

type slice struct{}
func (fc *slice) Evaluate(globalContext *Context, localContext *Context) (ExpressionResult, error) {
    input := GetVariableOrPanic[*MozzarellaVariable](localContext,"input").Value
    start := GetVariableOrPanic[*ParmesanVariable](localContext,"start").Value
    end := GetVariableOrPanic[*ParmesanVariable](localContext,"end").Value
    
    if start < 0 {
        return NullExpressionResult, fmt.Errorf("start must be greater than 0")
    }

    if end < start {
        return NullExpressionResult, fmt.Errorf("end shall not be smaller than start")
    }

    if end > len(input) {
        return NullExpressionResult, fmt.Errorf("end shall not be greater than the mozzarella string itself")
    }

    return ExpressionResult{Value: &MozzarellaVariable{Value: input[start:end]}, Return: true, Brake: false}, nil
}
var sliceFunction = Function{
    Name: "slice",
    ArgumentsType: []VariableType{Mozzarella, Parmesan, Parmesan},
    ArgumentsNames: []string{"input", "start", "end"},
    Code: &slice{}, 
}

var StandardLibraryFunctions = []Function{parmesanToGorgonzolaFunc,parmesanToMozzarellaFunc,gorgonzolaToParmesanFunc,gorgonzolaToMozzarellaFunc, mozzarellaToGorgonzolaFunc, mozzarellaToParmesanFunc,serveFunction, eatFunction, weightFunction, sliceFunction}


