package parser

func parseHeaders(tokens []Token, i int) []HeaderNode {
	n := []HeaderNode{}
	i++
	c := 0
	h := HeaderNode{}
	for tokens[i].Type != TOKEN_RBRACE {
		switch tokens[i].Type {
		case TOKEN_IDENTIFIER:
			if c == 1 && h.Key != "" {
				h.Value = tokens[i].Value
				c = 0
			} else {
				h.Key = tokens[i].Value
				c = 1
			}
			if h.Key != "" && h.Value != "" {
				n = append(n, h)
				h = HeaderNode{}
			}
		}
		i++
	}
	return n
}
