package parser

import (
	"context"
	"fmt"
	"sync"
	"task/reader"
	"task/writer"
)

type Parser interface {
	Parse(fileNames map[string]string)
}

type MyParser struct {
	r reader.Reader
	w writer.Writer
}

func (p *MyParser) Parse(fileNames map[string]string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	placeholder := context.WithoutCancel(context.Background())
	var wg sync.WaitGroup
	for inFileName, outFileName := range fileNames {
		wg.Add(1)
		go func(inFileName, outFileName string) {
			defer wg.Done()
			resChan, errChan := p.r.GetContent(ctx, inFileName)
			wErrChan := p.w.WriteContent(placeholder, resChan, outFileName)
			var pErrDone, errDone bool
			for {
				select {
				case err, ok := <-errChan:
					if err != nil {
						cancel()
						fmt.Println(err)
					}
					if !ok {
						errDone = true
					}
				case err, ok := <-wErrChan:
					if err != nil {
						cancel()
						fmt.Println(err)
					}
					if !ok {
						pErrDone = true
					}
				}
				if errDone && pErrDone {
					break
				}
			}
		}(inFileName, outFileName)
	}
	wg.Wait()
}
