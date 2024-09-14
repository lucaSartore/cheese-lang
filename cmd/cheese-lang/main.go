package main

import (
	"cheese-lang/internal/expressions"
	"cheese-lang/internal/parser"
	"cheese-lang/internal/tokenizer"
	"fmt"
	"os"
)

func main() {

    if len(os.Args) < 2 {
        fmt.Print("you need to specify a source file in order to run cheeselang")
        return
    }

    file := os.Args[1]

    source, err := os.ReadFile(file)

    if err != nil {
        fmt.Printf("impossible to open the file %v because of error: %v",file, err)
        return
    }
        
    source_str := string(source)
    source_str = "{" + source_str + "}"

    tokens, err := tokenizer.Tokenize(source_str,true)

    if err != nil {
        fmt.Printf("error in tokenization: %v",err)
        return
    }
    p := parser.MakeParser(tokens)

    parsed_code := p.ParseAnything(true)  
    
    if parsed_code.Error != nil {
        fmt.Printf("unable to parse the code at line %v:%v because of error: %v", parsed_code.Line, parsed_code.Colum, parsed_code.Error)
        return
    }

    local_context := expressions.MakeContextWithStd()
    global_context := expressions.MakeContext()
    
    _, err = parsed_code.Expression.Evaluate(&global_context, &local_context)

    if err != nil{
        fmt.Printf("runtime error: %v",err)
    }
    
    main, success := global_context.GetFunction("main")
    
    if !success {
        fmt.Print("unable to find \"main\" entry point")
        return
    }
    
    return_code, err := main.Code.Evaluate(&global_context,&local_context)
    
    if err != nil {
        fmt.Printf("runtime error: %v",err)
        return
    }
    
    switch return_code.Value.GetVariableType(){
        case expressions.Mozzarella:
        fmt.Printf("Program exited with message: %s\n",return_code.Value.(*expressions.MozzarellaVariable).Value)
        case expressions.Parmesan:
        fmt.Printf("Program exited with return code: %v\n",return_code.Value.(*expressions.ParmesanVariable).Value)
        default:
    }

}
