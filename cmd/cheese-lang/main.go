package main
	
import (
    "os"
    "fmt"
	"cheese-lang/internal/tokenizer"
	"cheese-lang/internal/parser"
	"cheese-lang/internal/expressions"
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

    tokens, err := tokenizer.Tokenize(source_str)

    if err != nil {
        fmt.Printf("error in tokenization: %v",err)
        return
    }
    p := parser.MakeParser(tokens)

    parsed_code := p.ParseAnything(true)  
    
    if parsed_code.Error != nil {
        fmt.Printf("unable to parse the code because of error: %v",parsed_code.Error)
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
    
    _, err = main.Code.Evaluate(&global_context,&local_context)
    
    if err != nil {
        fmt.Printf("runtime error: %v",err)
    }

}
