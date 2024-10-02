package parser

import (
	"task/reader"
	"task/writer"
)

type ParserFactory interface {
	CreateParser() Parser
}

type MyParserFactory struct{}

func (f *MyParserFactory) CreateParser() Parser {
	return &MyParser{
		r: reader.NewMyReader(),
		w: writer.NewMyWriter(),
	}
}
