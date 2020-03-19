package config

import (
	"fmt"
	"os"
	"time"

	"github.com/n9e/wechat-sender/corp"
	"github.com/toolkits/pkg/logger"
)

// InitLogger init logger toolkits
func InitLogger() {
	c := Get().Logger

	lb, err := logger.NewFileBackend(c.Dir)
	if err != nil {
		fmt.Println("cannot init logger:", err)
		os.Exit(1)
	}

	lb.SetRotateByHour(true)
	lb.SetKeepHours(c.KeepHours)

	logger.SetLogging(c.Level, lb)
}

func Test(args []string) {
	c := Get()

	chatClient := corp.New(c.WeChat.CorpID, c.WeChat.AgentID, c.WeChat.Secret)

	if len(args) == 0 {
		fmt.Println("mail address not given")
		os.Exit(1)
	}

	for i := 0; i < len(args); i++ {
		err := chatClient.Send(corp.Message{
			ToUser:  args[i],
			MsgType: "text",
			Text:    corp.Content{Content: fmt.Sprintf("test message from n9e at %v", time.Now())},
		})

		if err != nil {
			fmt.Printf("send to %s fail: %v\n", args[i], err)
		} else {
			fmt.Printf("send to %s succ\n", args[i])
		}
	}
}
