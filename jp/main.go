package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type TokenType string

var (
	LeftBrace    TokenType = "{"
	RightBrace   TokenType = "}"
	LeftBracket  TokenType = "["
	RightBracket TokenType = "]"
	Colon        TokenType = ":"
	Comma        TokenType = ","
	String       TokenType = "string"
	Number       TokenType = "number"
	Boolean      TokenType = "boolean"
	Null         TokenType = "null"
)

type Token struct {
	Type  TokenType
	Value string
}

func lex(input string) ([]Token, error) {
	// Lexical analysis breaks down an input string into tokens.
	tokens := make([]Token, 0)

	for _, v := range input {
		switch v {
		case '{':
			tokens = append(tokens, Token{Type: LeftBrace, Value: "{"})
		case '}':
			tokens = append(tokens, Token{Type: RightBrace, Value: "}"})
		}
	}

	return tokens, nil
}

func parse(tokens []Token) (interface{}, error) {
	for i, token := range tokens {
		switch token.Type {
		case LeftBrace:
			{
				var obj map[string]interface{}
				if i+1 <= len(tokens) && tokens[i+1].Type == RightBrace {
					return obj, nil
				}
			}
		case RightBrace:
			// End object
		}
	}

	return nil, errors.New("Invalid JSON")
}

func main() {
	var fileFlag string

	flag.StringVar(&fileFlag, "f", "", "File to parse")

	flag.Parse()

	if fileFlag == "" {
		fmt.Println("USAGE ERROR: -f flag is not provided")
		os.Exit(1)
	}

	file, err := os.Open(fileFlag)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	var buf bytes.Buffer

	if _, err := io.Copy(&buf, file); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tokens, err := lex(buf.String())

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jsonValue, err := parse(tokens)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("TOKENS: %+v", tokens)
	fmt.Println()
	fmt.Printf("PARSED JSON: %+v", jsonValue)
	fmt.Println()
}
