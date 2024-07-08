package parser_test

import (
	_ "embed"
	"log"
	"testing"

	"github.com/mukezhz/bru-go/parser"
)

//go:embed bru/body.bru
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
		t.Errorf("Expected type __body__, got %s", tokens[0].Type)
	}
}

//go:embed bru/auth.bru
var authBru string

func TestParseAuth(t *testing.T) {
	auth := parser.Token{
		Type:  parser.TOKEN_AUTH,
		Value: "auth",
	}
	tokens := parser.Tokenize(authBru)
	if len(tokens) != 8 {
		t.Errorf("Expected 8 tokens, got %d", len(tokens))
	}

	if tokens[0].Value != auth.Value {
		t.Errorf("Expected bru tag auth, got %s", tokens[0].Type)
	}

	if tokens[0].Type != auth.Type {
		t.Errorf("Expected type __auth__, got %s", tokens[0].Type)
	}
}

//go:embed bru/meta.bru
var metaBru string

func TestParseMeta(t *testing.T) {
	meta := parser.Token{
		Type:  parser.TOKEN_META,
		Value: "meta",
	}
	tokens := parser.Tokenize(metaBru)
	if len(tokens) != 12 {
		t.Errorf("Expected 12 tokens, got %d", len(tokens))
	}

	if tokens[0].Value != meta.Value {
		t.Errorf("Expected bru tag meta, got %s", tokens[0].Type)
	}

	if tokens[0].Type != meta.Type {
		t.Errorf("Expected type __meta__, got %s", tokens[0].Type)
	}
}

//go:embed bru/post.bru
var postBru string

func TestParsePost(t *testing.T) {
	post := parser.Token{
		Type:  parser.TOKEN_POST,
		Value: "post",
	}
	tokens := parser.Tokenize(postBru)
	if len(tokens) != 12 {
		t.Errorf("Expected 12 tokens, got %d", len(tokens))
	}

	if tokens[0].Value != post.Value {
		t.Errorf("Expected bru tag post, got %s", tokens[0].Type)
	}

	if tokens[0].Type != post.Type {
		t.Errorf("Expected type __post__, got %s", tokens[0].Type)
	}
}

//go:embed bru/headers.bru
var headersBru string

func TestHeaders(t *testing.T) {
	header := parser.Token{
		Type:  parser.TOKEN_HEADERS,
		Value: "headers",
	}
	tokens := parser.Tokenize(headersBru)
	if len(tokens) != 9 {
		t.Errorf("Expected 9 tokens, got %d", len(tokens))
	}

	if tokens[0].Value != header.Value {
		t.Errorf("Expected bru tag headers, got %s", tokens[0].Type)
	}

	if tokens[0].Type != header.Type {
		t.Errorf("Expected type __headers__, got %s", tokens[0].Type)
	}
}
