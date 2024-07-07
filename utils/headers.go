package utils

import "github.com/mukezhz/bru-go/parser"

func FromHeaderNodeToMap(headers []parser.HeaderNode) map[string]string {
	m := map[string]string{}
	for _, v := range headers {
		m[v.Key] = v.Value
	}
	return m
}
