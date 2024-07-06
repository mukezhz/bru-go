package parser

import (
	"fmt"
)

func parseBody(tokens []Token, i *int) BodyNode {
	body := BodyNode{}
	content := []string{}
	*i++
	for tokens[*i].Type != TOKEN_RBRACE {
		switch tokens[*i].Type {
		case TOKEN_STRING:
			content = append(content, tokens[*i].Value)
		case TOKEN_INTEGER:
			content = append(content, tokens[*i].Value)
		case TOKEN_COLON:
			if len(content) != 0 {
				content = append(content, tokens[*i].Value)
			}
		}
		*i++
	}
	if len(content) == 0 {
		body.Content = ""
		return body
	}

	jsonContent := ""
	colonCount := 0
	for i, c := range content {
		jsonContent += c
		if colonCount == 1 && i != len(content)-1 {
			jsonContent += ","
			colonCount = 0
		}
		if c == ":" {
			jsonContent += " "
			colonCount++
		}
	}
	body.Content += fmt.Sprintf("{%s}", jsonContent)
	return body
}
