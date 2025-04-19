package lexer

import (
	"fmt"
	"testing"

	"github.com/Martin-Martinez4/compiler_in_go/token"
)

func TestNextTokenSimpleInput(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		name            string
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{"assign/equal sign", token.ASSIGN, "="},
		{"plus", token.PLUS, "+"},
		{"Left Paren", token.LPAREN, "("},
		{"Right Paren", token.RPAREN, ")"},
		{"Left Brace", token.LBRACE, "{"},
		{"Right Brace", token.RBRACE, "}"},
		{"Comma", token.COMMA, ","},
		{"Semicolon", token.SEMICOLON, ";"},
		{"End of File", token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		fmt.Println(i, tok)

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d, %s] - literal wrong, expected: %q got : %q", i, tt.name, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {

			t.Fatalf("tests[%d, %s] - literal wrong, expected: %q got : %q", i, tt.name, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	
	let add = fn(x,y){
		x+y;
	};
	
	let result = add(five, ten);
	!-/*5;
	5 < 10 > 5;

	if(5 < 10){
		return true;
	}else {
		return false;
	}

	10 == 10;
	10 != 9;
	"foobar"
	"foo bar"
	[1, 2];
	{"foo": "bar"}
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},

		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},

		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RBRACKET, "]"},
		{token.SEMICOLON, ";"},

		{token.LBRACE, "{"},
		{token.STRING, "foo"},
		{token.COLON, ":"},
		{token.STRING, "bar"},
		{token.RBRACE, "}"},

		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		fmt.Println(i, tok)

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - literal wrong, expected: %q got : %q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {

			t.Fatalf("tests[%d] - literal wrong, expected: %q got : %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
