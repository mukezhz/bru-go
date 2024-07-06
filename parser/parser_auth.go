package parser

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
