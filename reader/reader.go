package reader

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

type Reader interface {
	GetContent(ctx context.Context, fileName string) (<-chan string, <-chan error)
}

type MyReader struct{}

func NewMyReader() *MyReader {
	return &MyReader{}
}

func (r *MyReader) GetContent(ctx context.Context, fileName string) (<-chan string, <-chan error) {
	errChan := make(chan error)
	resChan := make(chan string)
	go func() {
		defer close(errChan)
		defer close(resChan)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				file, err := os.Open(fileName)
				if err != nil {
					errChan <- err
					return
				}
				defer file.Close()
				defer fmt.Println(fileName, "Closed in reader")
				reader := bufio.NewReader(file)
				line, err := reader.ReadString('\n')
				if err == io.EOF {
					return
				} else if err != nil {
					errChan <- err
					return
				}
				type Message struct {
					Message   string
					Timestamp time.Time
				}
				var message Message
				err = json.Unmarshal([]byte(line), &message)
				if err != nil {
					errChan <- err
					return
				}
				timeToSleep := rand.Intn(5) * int(time.Second)
				time.Sleep(time.Duration(timeToSleep))
				fmt.Println("message is", "["+message.Timestamp.Format(time.DateTime)+"]: "+message.Message)
				resChan <- "[" + message.Timestamp.Format(time.DateTime) + "]: " + message.Message
				return
			}
		}
	}()
	return resChan, errChan
}
