package bru

import (
	"fmt"

	"github.com/mukezhz/bru-go/parser"
)

type Bru struct {
	Meta    parser.MetaNode
	Body    parser.BodyNode
	Auth    parser.AuthNode
	HTTP    parser.HTTPNode
	Headers []parser.HeaderNode
}

func (bru *Bru) Unmarshal(b []byte) (err error) {
	if bru == nil {
		return fmt.Errorf("file is empty")
	}

	if len(b) == 0 {
		return fmt.Errorf("file is empty")
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Unmarshal", r)
			err = fmt.Errorf("unable to parse file")
		}
	}()

	d := string(b)
	tokens := GetTokens(d)
	ast := GetAST(tokens)
	for _, n := range ast {
		switch n := n.(type) {
		case parser.MetaNode:
			bru.Meta = n
		case parser.BodyNode:
			bru.Body = n
		case parser.AuthNode:
			bru.Auth = n
		case parser.HTTPNode:
			bru.HTTP = n
		case parser.HeaderNode:
			bru.Headers = append(bru.Headers, n)
		}
	}

	return nil
}
