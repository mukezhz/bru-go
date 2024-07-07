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
	for _, n := range ast {
		switch n := n.(type) {
		case parser.HTTPNode:
			log.Println("HTTPNode")
			log.Println(n.Method, n.URL, n.Body, n.Auth)
		case parser.MetaNode:
			log.Println("MetaNode")
			log.Println(n.Name, n.Type, n.Seq)
		case parser.BodyNode:
			log.Println("BodyNode")
			log.Println(n.Content)
		case []parser.HeaderNode:
			log.Println("HeaderNode")
			for _, v := range n {
				log.Println(v.Key, v.Value)
			}

			v := parser.FromHeaderNodeToMap(n)
			log.Println(v)
		}
	}

}
