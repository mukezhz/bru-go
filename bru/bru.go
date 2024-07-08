package bru

import "github.com/mukezhz/bru-go/parser"

func GetTokens(file string) []parser.Token {
	tokens := parser.Tokenize(file)
	return tokens
}

func GetAST(tokens []parser.Token) []parser.ASTNode {
	ast := parser.Parse(tokens)
	return ast
}

func GetTagNode[T any](ast []parser.ASTNode) T {
	for _, n := range ast {
		switch n := n.(type) {
		case T:
			return n
		}
	}
	return *new(T)
}

func Unmarshal(b []byte) (bru Bru, err error) {
	if err := bru.Unmarshal(b); err != nil {
		return Bru{}, err
	}
	return bru, nil
}
