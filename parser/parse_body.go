package parser

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
