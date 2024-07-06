package parser

import (
	"log"
	"strings"
)

func parsePost(tokens []Token, i *int) PostNode {
	post := PostNode{}
	*i++
	for tokens[*i].Type != TOKEN_RBRACE {
		switch tokens[*i].Type {
		case TOKEN_IDENTIFIER:
			switch strings.TrimSpace(tokens[*i].Value) {
			case "url":
				*i++
				post.URL = tokens[*i+1].Value
			case "body":
				*i++
				log.Println("Body: ", tokens[*i+1].Value)
				post.Body = tokens[*i+1].Value
			case "auth":
				*i++
				post.Auth = tokens[*i+1].Value
			}
		}
		*i++
	}
	log.Println("Parsed post node: ", post)
	return post
}
