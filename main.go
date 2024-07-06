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
	// for i, token := range tokens {
	// 	fmt.Println(i, token.Type, token.Value)
	// }
	ast := parser.Parse(tokens)
	// log.Println(ast)
	for _, node := range ast {
		if g, ok := node.(parser.HTTPNode); ok {
			log.Println(g.Method, g.URL, g.Body, g.Auth, g.Headers)
		}
		if g, ok := node.(parser.MetaNode); ok {
			log.Println(g.Name, g.Type, g.Seq)
		}
		if g, ok := node.(parser.BodyNode); ok {
			log.Println(g.Content)
		}
	}

}
