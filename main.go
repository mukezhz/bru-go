package main

import (
	_ "embed"
	"log"

	"github.com/mukezhz/bru-go/brugo"
	"github.com/mukezhz/bru-go/parser"
	"github.com/mukezhz/bru-go/utils"
)

//go:embed create.bru
var bruFile string

func main() {
	tokens := brugo.GetTokens(bruFile)
	ast := brugo.GetAST(tokens)
	n := brugo.GetTagNode[parser.HTTPNode](ast)
	log.Println(n.Method, n.URL, n.Body, n.Auth)
	headers := brugo.GetTagNode[[]parser.HeaderNode](ast)
	log.Println(headers)
	mapHeaders := utils.FromHeaderNodeToMap(headers)
	log.Println(mapHeaders)
	for k, v := range mapHeaders {
		log.Println(k, v)
	}

}
