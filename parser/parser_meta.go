package parser

import "strconv"

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
				*i++
				seq, _ := strconv.Atoi(tokens[*i+1].Value)
				meta.Seq = seq
			}
		}
		*i++
	}
	return meta
}
