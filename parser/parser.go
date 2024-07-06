package parser

type ASTNode interface{}

type MetaNode struct {
	Name string
	Type string
	Seq  int
}

type HTTPNode struct {
	URL  string
	Body string
	Auth string
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
	var nodes []ASTNode
	for i := 0; i < len(tokens); i++ {
		switch tokens[i].Type {
		case TOKEN_META:
			nodes = append(nodes, parseMeta(tokens, &i))
		case TOKEN_POST:
			nodes = append(nodes, parsePost(tokens, &i))
		case TOKEN_AUTH:
			nodes = append(nodes, parseAuth(tokens, &i))
		case TOKEN_BODY:
			nodes = append(nodes, parseBody(tokens, &i))
		}
	}
	return nodes
}
