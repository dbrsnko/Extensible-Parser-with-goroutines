package writer

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

type Writer interface {
	WriteContent(ctx context.Context, contentChan <-chan string, fileName string) <-chan error
}

type MyWriter struct{}

func NewMyWriter() *MyWriter {
	return &MyWriter{}
}

func (p *MyWriter) WriteContent(ctx context.Context, contentChan <-chan string, fileName string) <-chan error {
	errChan := make(chan error)
	go func() {
		defer close(errChan)
		select {
		case <-ctx.Done():
			return
		default:
			file, err := os.Create(fileName)
			if err != nil {
				errChan <- err
				return
			}
			defer file.Close()
			defer fmt.Println(fileName, "Closed in writer")
			w := bufio.NewWriter(file)
			defer w.Flush()
			chunk, ok := <-contentChan
			if !ok {
				return
			}
			fmt.Println("Chunk", chunk, "collected in writer")
			_, err = w.WriteString(chunk)
			if err != nil {
				errChan <- err
				return
			}
			fmt.Println("File ", fileName, " saved")
		}
	}()
	return errChan
}
