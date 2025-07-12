package service

import (
	"context"
	"fmt"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	arkmodel "github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"io"
)

type OpenAIMessage struct {
	Role    string
	Content string
}

// GetAIReplyStreamed 接收用户输入，流式返回AI回复内容（每次收到就返回，不等全部完成）
func GetAIReplyStreamed(userInput string, out chan<- string, errChan chan<- error) {
	go func() {
		client := arkruntime.NewClientWithApiKey("7f83334b-d834-486d-b2b1-002ac5d6819d")
		ctx := context.Background()

		messages := []*arkmodel.ChatCompletionMessage{
			{
				Role: "user",
				Content: &arkmodel.ChatCompletionMessageContent{
					StringValue: volcengine.String(userInput),
				},
			},
		}

		req := arkmodel.CreateChatCompletionRequest{
			Model:    "doubao-1-5-pro-32k-250115",
			Messages: messages,
			Stream:   volcengine.Bool(true),
		}

		stream, err := client.CreateChatCompletionStream(ctx, req)
		if err != nil {
			errChan <- fmt.Errorf("豆包流式接口调用失败: %v", err)
			close(out)
			close(errChan)
			return
		}
		defer stream.Close()

		for {
			recv, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				errChan <- fmt.Errorf("豆包流式读取失败: %v", err)
				break
			}
			if len(recv.Choices) > 0 {
				out <- recv.Choices[0].Delta.Content
			}
		}
		close(out)
		close(errChan)
	}()
}
