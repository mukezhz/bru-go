package main

import (
	_ "embed"
	"log"

	"github.com/mukezhz/bru-go/bru"
	"github.com/mukezhz/bru-go/parser"
	"github.com/mukezhz/bru-go/utils"
)

//go:embed create.bru
var bruFile string

func main() {
	tokens := bru.GetTokens(bruFile)
	ast := bru.GetAST(tokens)
	n := bru.GetTagNode[parser.HTTPNode](ast)
	log.Println(n.Method, n.URL, n.Body, n.Auth)
	headers := bru.GetTagNode[[]parser.HeaderNode](ast)
	log.Println(headers)
	mapHeaders := utils.FromHeaderNodeToMap(headers)
	log.Println(mapHeaders)
	for k, v := range mapHeaders {
		log.Println(k, v)
	}

}
