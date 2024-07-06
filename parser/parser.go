package parser

import (
	"strconv"
)

type ASTNode interface{}

type MetaNode struct {
	Name string
	Type string
	Seq  int
}

type HTTPNode struct {
	URL  string
	Body *string
	Auth *string
}

type PostNode struct {
	HTTPNode
}

type AuthNode struct {
	Token string
}

type BodyNode struct {
	Content string
}

func Parse(tokens []Token) ASTNode {
	var node ASTNode
	for i := 0; i < len(tokens); i++ {
		switch tokens[i].Type {
		case TOKEN_META:
			node = parseMeta(tokens, &i)
		case TOKEN_POST:
			node = parsePost(tokens, &i)
		case TOKEN_AUTH:
			node = parseAuth(tokens, &i)
		case TOKEN_BODY:
			node = parseBody(tokens, &i)
		}
	}
	return node
}

func parseMeta(tokens []Token, i *int) MetaNode {
	meta := MetaNode{}
	*i++
	for tokens[*i].Type != TOKEN_RBRACE {
		switch tokens[*i].Type {
		case TOKEN_IDENTIFIER:
			switch tokens[*i].Value {
			case "name":
				*i++
				meta.Name = tokens[*i+1].Value
			case "type":
				*i++
				meta.Type = tokens[*i+1].Value
			case "seq":
				seq, _ := strconv.Atoi(tokens[*i+1].Value)
				*i++
				meta.Seq = seq
			}
		}
		*i++
	}
	return meta
}

func parsePost(tokens []Token, i *int) PostNode {
	post := PostNode{}
	*i++
	for tokens[*i].Type != TOKEN_RBRACE {
		switch tokens[*i].Type {
		case TOKEN_IDENTIFIER:
			switch tokens[*i].Value {
			case "url":
				*i++
				post.URL = tokens[*i+1].Value
			case "body":
				*i++
				post.Body = &tokens[*i+1].Value
			case "auth":
				*i++
				post.Auth = &tokens[*i+1].Value
			}
		}
		*i++
	}
	return post
}

func parseAuth(tokens []Token, i *int) AuthNode {
	auth := AuthNode{}
	*i++
	for tokens[*i].Type != TOKEN_RBRACE {
		switch tokens[*i].Type {
		case TOKEN_IDENTIFIER:
			if tokens[*i].Value == "token" {
				*i++
				auth.Token = tokens[*i+1].Value
			}
		}
		*i++
	}
	return auth
}

func parseBody(tokens []Token, i *int) BodyNode {
	body := BodyNode{}
	*i++
	for tokens[*i].Type != TOKEN_RBRACE {
		switch tokens[*i].Type {
		case TOKEN_STRING:
			body.Content = tokens[*i].Value
		}
		*i++
	}
	return body
}
