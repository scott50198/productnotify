package model

type Request struct {
	Url       string
	Method    string
	PostBody  map[string]string
	ParseFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}
