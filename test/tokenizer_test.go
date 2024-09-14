package test 

import (
	"cheese-lang/internal/tokenizer"
	"fmt"
	"testing"
)

func TestTokenizer1(t *testing.T) {
	s := "//this is a comment\n == = != && ^{    }\t	\n\n \t  -="
	tokens, err := tokenizer.Tokenize(s,false)

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
	s := "   Mozzarella Gorgonzola 	\"literal\" \"literal\\n \\\\ \\t \\\"\""

	tokens, err := tokenizer.Tokenize(s,false)

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

func TestTokenizer3(t *testing.T) {
	s := `
	
	//this is a comment

	recipe doSomething(Mozzarella m, Parmesan p, Gorgonzola g) -> Mozzarella, Parmesan, Gorgonzola{
		prepare m, p, g;
	}

	recipe main() -> Ricotta{
		//this is a comment
		Milk MyMilk = spoiled;
		
		Mozzarella MyMozzarella = "my mozzarella";
		Parmesan MyParmesan = 10;
		Gorgonzola MyGorgonzola = 10.0;

		taste(!MyMilk){
			MyMozerella = MyMozzarella + " is good";
			MyParmesan = MyParmesan * 10;
			MyGorgonzola = MyGorgonzola / 10.0;
		}
		
		prepare;		
	}
	`

	expected_tokens := []tokenizer.Token{
		tokenizer.MakeTokenWithMessage(tokenizer.Comment, "this is a comment"),
		tokenizer.MakeToken(tokenizer.RecipeKeyword),
		tokenizer.MakeTokenWithMessage(tokenizer.Identifier, "doSomething"),
		tokenizer.MakeToken(tokenizer.OpenBracket),
		tokenizer.MakeToken(tokenizer.MozzarellaType),
		tokenizer.MakeTokenWithMessage(tokenizer.Identifier, "m"),
		tokenizer.MakeToken(tokenizer.Comma),
		tokenizer.MakeToken(tokenizer.ParmesanType),
		tokenizer.MakeTokenWithMessage(tokenizer.Identifier, "p"),
		tokenizer.MakeToken(tokenizer.Comma),
		tokenizer.MakeToken(tokenizer.GorgonzolaType),
		tokenizer.MakeTokenWithMessage(tokenizer.Identifier, "g"),
		tokenizer.MakeToken(tokenizer.CloseBracket),
		tokenizer.MakeToken(tokenizer.Arrow),
		tokenizer.MakeToken(tokenizer.MozzarellaType),
		tokenizer.MakeToken(tokenizer.Comma),
		tokenizer.MakeToken(tokenizer.ParmesanType),
		tokenizer.MakeToken(tokenizer.Comma),
		tokenizer.MakeToken(tokenizer.GorgonzolaType),
		tokenizer.MakeToken(tokenizer.OpenCurlyBracket),
		tokenizer.MakeToken(tokenizer.PrepareKeyword),
		tokenizer.MakeTokenWithMessage(tokenizer.Identifier, "m"),
		tokenizer.MakeToken(tokenizer.Comma),
		tokenizer.MakeTokenWithMessage(tokenizer.Identifier, "p"),
		tokenizer.MakeToken(tokenizer.Comma),
		tokenizer.MakeTokenWithMessage(tokenizer.Identifier, "g"),
		tokenizer.MakeToken(tokenizer.SemiColon),
		tokenizer.MakeToken(tokenizer.CloseCurlyBracket),
		tokenizer.MakeToken(tokenizer.RecipeKeyword),
		tokenizer.MakeTokenWithMessage(tokenizer.Identifier, "main"),
		tokenizer.MakeToken(tokenizer.OpenBracket),
		tokenizer.MakeToken(tokenizer.CloseBracket),
		tokenizer.MakeToken(tokenizer.Arrow),
		tokenizer.MakeToken(tokenizer.RicottaType),
		tokenizer.MakeToken(tokenizer.OpenCurlyBracket),
		tokenizer.MakeTokenWithMessage(tokenizer.Comment, "this is a comment"),
		tokenizer.MakeToken(tokenizer.MilkType),
		tokenizer.MakeTokenWithMessage(tokenizer.Identifier, "MyMilk"),
		tokenizer.MakeToken(tokenizer.AssignOperator),
		tokenizer.MakeToken(tokenizer.SpoiledMilk),
		tokenizer.MakeToken(tokenizer.SemiColon),
		tokenizer.MakeToken(tokenizer.MozzarellaType),
		tokenizer.MakeTokenWithMessage(tokenizer.Identifier, "MyMozzarella"),
		tokenizer.MakeToken(tokenizer.AssignOperator),
		tokenizer.MakeTokenWithMessage(tokenizer.MozzarellaLiteral, "my mozzarella"),
		tokenizer.MakeToken(tokenizer.SemiColon),
		tokenizer.MakeToken(tokenizer.ParmesanType),
		tokenizer.MakeTokenWithMessage(tokenizer.Identifier, "MyParmesan"),
		tokenizer.MakeToken(tokenizer.AssignOperator),
		tokenizer.MakeTokenWithMessage(tokenizer.ParmesanLiteral, "10"),
		tokenizer.MakeToken(tokenizer.SemiColon),
		tokenizer.MakeToken(tokenizer.GorgonzolaType),
		tokenizer.MakeTokenWithMessage(tokenizer.Identifier, "MyGorgonzola"),
		tokenizer.MakeToken(tokenizer.AssignOperator),
		tokenizer.MakeTokenWithMessage(tokenizer.GorgonzolaLiteral, "10.0"),
		tokenizer.MakeToken(tokenizer.SemiColon),
		tokenizer.MakeToken(tokenizer.TasteKeyword),
		tokenizer.MakeToken(tokenizer.OpenBracket),
		tokenizer.MakeToken(tokenizer.NotOperator),
		tokenizer.MakeTokenWithMessage(tokenizer.Identifier, "MyMilk"),
		tokenizer.MakeToken(tokenizer.CloseBracket),
		tokenizer.MakeToken(tokenizer.OpenCurlyBracket),
		tokenizer.MakeTokenWithMessage(tokenizer.Identifier, "MyMozerella"),
		tokenizer.MakeToken(tokenizer.AssignOperator),
		tokenizer.MakeTokenWithMessage(tokenizer.Identifier, "MyMozzarella"),
		tokenizer.MakeToken(tokenizer.AddOperator),
		tokenizer.MakeTokenWithMessage(tokenizer.MozzarellaLiteral, " is good"),
		tokenizer.MakeToken(tokenizer.SemiColon),
		tokenizer.MakeTokenWithMessage(tokenizer.Identifier, "MyParmesan"),
		tokenizer.MakeToken(tokenizer.AssignOperator),
		tokenizer.MakeTokenWithMessage(tokenizer.Identifier, "MyParmesan"),
		tokenizer.MakeToken(tokenizer.MulOperator),
		tokenizer.MakeTokenWithMessage(tokenizer.ParmesanLiteral, "10"),
		tokenizer.MakeToken(tokenizer.SemiColon),
		tokenizer.MakeTokenWithMessage(tokenizer.Identifier, "MyGorgonzola"),
		tokenizer.MakeToken(tokenizer.AssignOperator),
		tokenizer.MakeTokenWithMessage(tokenizer.Identifier, "MyGorgonzola"),
		tokenizer.MakeToken(tokenizer.DivOperator),
		tokenizer.MakeTokenWithMessage(tokenizer.GorgonzolaLiteral, "10.0"),
		tokenizer.MakeToken(tokenizer.SemiColon),
		tokenizer.MakeToken(tokenizer.CloseCurlyBracket),
		tokenizer.MakeToken(tokenizer.PrepareKeyword),
		tokenizer.MakeToken(tokenizer.SemiColon),
		tokenizer.MakeToken(tokenizer.CloseCurlyBracket),
	}

	tokens, err := tokenizer.Tokenize(s,false)

	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(tokens)

	VerityTokens(tokens, expected_tokens, t)
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
