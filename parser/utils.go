package parser

func FromHeaderNodeToMap(headers []HeaderNode) map[string]string {
	m := map[string]string{}
	for _, v := range headers {
		m[v.Key] = v.Value
	}
	return m
}
