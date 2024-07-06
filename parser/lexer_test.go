package parser_test

import (
	_ "embed"
	"log"
	"testing"

	"github.com/mukezhz/bru-go/parser"
)

//go:embed body.bru
var bodyBru string

func TestParseBody(t *testing.T) {
	body := parser.Token{
		Type:  parser.TOKEN_BODY,
		Value: "body",
	}
	tokens := parser.Tokenize(bodyBru)
	log.Println(len(tokens))
	if len(tokens) != 10 {
		t.Errorf("Expected 10 tokens, got %d", len(tokens))
	}

	if tokens[0].Value != body.Value {
		t.Errorf("Expected bru tag body, got %s", tokens[0].Type)
	}

	if tokens[0].Type != body.Type {
		t.Errorf("Expected bru tag body, got %s", tokens[0].Type)
	}
}
