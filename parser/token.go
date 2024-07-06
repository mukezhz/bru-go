package parser

type TokenType string

const (
	TOKEN_META       TokenType = "__meta__"
	TOKEN_GET        TokenType = "__get__"
	TOKEN_POST       TokenType = "__post__"
	TOKEN_PUT        TokenType = "__put__"
	TOKEN_DELETE     TokenType = "__delete__"
	TOKEN_PATCH      TokenType = "__patch__"
	TOKEN_OPTIONS    TokenType = "__options__"
	TOKEN_HEAD       TokenType = "__head__"
	TOKEN_HEADERS    TokenType = "__headers__"
	TOKEN_AUTH       TokenType = "__auth__"
	TOKEN_BODY       TokenType = "__body__"
	TOKEN_IDENTIFIER TokenType = "__identifier__"
	TOKEN_STRING     TokenType = "__string__"
	TOKEN_INTEGER    TokenType = "__integer__"
	TOKEN_COLON      TokenType = "__colon__"
	TOKEN_LBRACE     TokenType = "__lbrace__"
	TOKEN_RBRACE     TokenType = "__rbrace__"
)

type Token struct {
	Type  TokenType
	Value string
}
