package test

import (
	"fmt"
	"testing"

	"github.com/lucaSartore/cheese-lang/internal/tokenizer"
)

func TestTokenizer1(t *testing.T) {
	s := "\\\\this is a comment\n == = != && ^{    }\t	\n\n \t  -="
	tokens, err := tokenizer.Tokenize(s)

	if err != nil {
		t.Error(err)
		return
	}

	expected_tokens := []tokenizer.Token{
		tokenizer.MakeTokenWithMessage(tokenizer.Comment, "this is a comment"),
		tokenizer.MakeToken(tokenizer.EqualOperator),
		tokenizer.MakeToken(tokenizer.AssignOperator),
		tokenizer.MakeToken(tokenizer.UnEqualOperator),
		tokenizer.MakeToken(tokenizer.AndOperator),
		tokenizer.MakeToken(tokenizer.ExorOperator),
		tokenizer.MakeToken(tokenizer.OpenCurlyBracket),
		tokenizer.MakeToken(tokenizer.CloseCurlyBracket),
		tokenizer.MakeToken(tokenizer.SubOperator),
		tokenizer.MakeToken(tokenizer.AssignOperator),
	}

	VerityTokens(tokens, expected_tokens, t)
}

func TestTokenizer2(t *testing.T) {
	s := "   mozzarella gorgonzola 	\"literal\" \"literal\\n \\\\ \\t \\\"\""

	tokens, err := tokenizer.Tokenize(s)

	if err != nil {
		t.Error(err)
	}

	expected_tokens := []tokenizer.Token{
		tokenizer.MakeToken(tokenizer.MozzarellaType),
		tokenizer.MakeToken(tokenizer.GorgonzolaType),
		tokenizer.MakeTokenWithMessage(tokenizer.MozzarellaLiteral, "literal"),
		tokenizer.MakeTokenWithMessage(tokenizer.MozzarellaLiteral, "literal\n \\ \t \""),
	}

	VerityTokens(tokens, expected_tokens, t)
}

func TestTokenizer4(t *testing.T) {
	s := `
	
	//this is a comment

	recipie doSomething(Mozzarella m, Parmesan p, Gorgonzola g) -> Mozzarella, Parmesan, Gorgonzola{
		prepare m, p, g;
	}

	recipie main() -> Ricotta{
		//this is a comment
		milk MyMilk = spoiled;
		
		Mozzarella MyMozzarella = "my mozzarella";
		Parmesan MyParmesan = 10;
		Gorgonzola MyGorgonzola = 10.0;

		if(!MyMilk){
			MyMozerella = MyMozzarella + " is good";
			MyParmesan = MyParmesan * 10;
			MyGorgonzola = MyGorgonzola / 10.0;
		}
		
		prepare;		
	}
	
	`
	tokens, err := tokenizer.Tokenize(s)

	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(tokens)

	// VerityTokens(tokens, expected_tokens, t)
}

func VerityTokens(tokens []tokenizer.Token, expected_tokens []tokenizer.Token, t *testing.T) {

	for i, token := range tokens {
		if token.TokenType != expected_tokens[i].TokenType {
			t.Errorf("Expected token type %v, got %v", expected_tokens[i].TokenType, token.TokenType)
		}
		if token.TokenVal != expected_tokens[i].TokenVal {
			t.Errorf("Expected token value %v, got %v", expected_tokens[i].TokenVal, token.TokenVal)
		}
	}
}
