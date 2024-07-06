package parser

import (
	"fmt"
	"regexp"
)

const (
	BRUNO_TAGS  = `(meta|post|auth|body)`
	OTHER_REGEX = `(\w+\s?\w+|".*?"|{|}|\d+|:)`
	URL_REGEX   = `(http[s]?:\/\/)?(\{\{\w+\}\})(\/[^\s]*)?`
)

func Tokenize(input string) []Token {
	stack := []rune{}
	var tokens []Token
	re := regexp.MustCompile(fmt.Sprintf(`%s|%s|%s`, BRUNO_TAGS, URL_REGEX, OTHER_REGEX))
	matches := re.FindAllString(input, -1)
	for _, match := range matches {
		if match == "{" {
			stack = append(stack, '{')
		} else if match == "}" {
			stack = stack[:len(stack)-1]
		}

		bruRe := regexp.MustCompile(BRUNO_TAGS)
		if len(stack) > 0 && bruRe.MatchString(match) {
			match = match + " "
		}
		switch match {
		case "meta":
			tokens = append(tokens, Token{Type: TOKEN_META, Value: match})
		case "put":
			tokens = append(tokens, Token{Type: TOKEN_PUT, Value: match})
		case "post":
			tokens = append(tokens, Token{Type: TOKEN_POST, Value: match})
		case "get":
			tokens = append(tokens, Token{Type: TOKEN_GET, Value: match})
		case "delete":
			tokens = append(tokens, Token{Type: TOKEN_DELETE, Value: match})
		case "options":
			tokens = append(tokens, Token{Type: TOKEN_OPTIONS, Value: match})
		case "patch":
			tokens = append(tokens, Token{Type: TOKEN_PATCH, Value: match})
		case "head":
			tokens = append(tokens, Token{Type: TOKEN_HEAD, Value: match})
		case "headers":
			tokens = append(tokens, Token{Type: TOKEN_HEADERS, Value: match})
		case "auth":
			tokens = append(tokens, Token{Type: TOKEN_AUTH, Value: match})
		case "body":
			tokens = append(tokens, Token{Type: TOKEN_BODY, Value: match})
		case ":":
			tokens = append(tokens, Token{Type: TOKEN_COLON, Value: match})
		case "{":
			tokens = append(tokens, Token{Type: TOKEN_LBRACE, Value: match})
		case "}":
			tokens = append(tokens, Token{Type: TOKEN_RBRACE, Value: match})
		default:
			if regexp.MustCompile(`^\d+$`).MatchString(match) {
				tokens = append(tokens, Token{Type: TOKEN_INTEGER, Value: match})
			} else if regexp.MustCompile(`^".*"$`).MatchString(match) {
				tokens = append(tokens, Token{Type: TOKEN_STRING, Value: match})
			} else {
				tokens = append(tokens, Token{Type: TOKEN_IDENTIFIER, Value: match})
			}
		}
	}
	return tokens
}
