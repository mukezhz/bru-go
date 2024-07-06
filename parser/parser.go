package parser

type ASTNode interface{}

type MetaNode struct {
	Name string
	Type string
	Seq  int
}

type HTTPNode struct {
	Method  string
	URL     string
	Body    string
	Auth    string
	Headers map[string]string
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

func Parse(tokens []Token) []ASTNode {
	var nodes []ASTNode
	for i := range tokens {
		i := i
		switch tokens[i].Type {
		case TOKEN_META:
			nodes = append(nodes, parseMeta(tokens, i))
		case TOKEN_POST, TOKEN_GET, TOKEN_PUT, TOKEN_DELETE, TOKEN_OPTIONS, TOKEN_PATCH, TOKEN_HEAD:
			nodes = append(nodes, parseHttpMethod(tokens, i))
		case TOKEN_BODY:
			nodes = append(nodes, parseBody(tokens, i))
		case TOKEN_AUTH:
			nodes = append(nodes, parseAuth(tokens, i))
		}
	}
	return nodes
}
