package parser

import (
	"strings"
)

func parseHttpMethod(tokens []Token, i int) HTTPNode {
	req := HTTPNode{}
	req.Method = tokens[i].Value
	i++
	for tokens[i].Type != TOKEN_RBRACE {
		switch tokens[i].Type {
		case TOKEN_IDENTIFIER:
			switch strings.TrimSpace(tokens[i].Value) {
			case "url":
				i++
				req.URL = tokens[i+1].Value
			case "body":
				i++
				req.Body = tokens[i+1].Value
			case "auth":
				i++
				req.Auth = tokens[i+1].Value
			}
		}
		i++
	}
	return req
}
