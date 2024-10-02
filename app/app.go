package app

import (
	"task/parser"
)

type App interface {
	parser.Parser
	run()
}

type MyApp struct {
	p parser.Parser
}

func (a *MyApp) Run() {
	fileNames := map[string]string{"file1.txt": "out1.txt", "file2.txt": "out2.txt", "file3.txt": "out3.txt"}
	a.p.Parse(fileNames)
}

func NewMyApp() *MyApp {
	factory := &parser.MyParserFactory{}
	return &MyApp{p: factory.CreateParser()}
}
