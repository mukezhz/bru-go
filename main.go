package main

import (
	_ "embed"
	"log"

	"github.com/mukezhz/bru-go/parser"
	"github.com/mukezhz/bru-go/utils"
)

//go:embed create.bru
var bruFile string

func main() {
	tokens := GetTokens(bruFile)
	ast := GetAST(tokens)
	n := GetTagNode[parser.HTTPNode](ast)
	log.Println(n.Method, n.URL, n.Body, n.Auth)
	headers := GetTagNode[[]parser.HeaderNode](ast)
	log.Println(headers)
	mapHeaders := utils.FromHeaderNodeToMap(headers)
	log.Println(mapHeaders)
	for k, v := range mapHeaders {
		log.Println(k, v)
	}

}

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
