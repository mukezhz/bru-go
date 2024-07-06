package main

import (
	_ "embed"
	"log"

	"github.com/mukezhz/bru-go/parser"
)

//go:embed create.bru
var bruFile string

func main() {
	// fmt.Println(bruFile)
	tokens := parser.Tokenize(bruFile)
	// for _, token := range tokens {
	// 	fmt.Println(token)
	// }
	ast := parser.Parse(tokens)
	log.Println(ast)

}
