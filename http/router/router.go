package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/toolkits/pkg/logger"

	"github.com/n9e/wechat-sender/config"
	"github.com/n9e/wechat-sender/corp"
	"github.com/n9e/wechat-sender/http/render"
)

var chatClient *corp.Client

func ConfigRoutes(r *mux.Router) {
	r.HandleFunc("/send/wechat", apiSendWechat)
	r.HandleFunc("/ping", ping)
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

type Message struct {
	Tos     []string `json:"tos"`
	Content string   `json:"content"`
}

func apiSendWechat(w http.ResponseWriter, r *http.Request) {
	var message Message
	bindJSON(r, &message)

	cnt := len(message.Tos)
	if cnt == 0 {
		logger.Warningf("api send wechat fail, empty tos, message: %+v", message)
		render.Message(w, "api empty tos")
		return
	}

	c := config.Get()
	client := corp.New(c.WeChat.CorpID, c.WeChat.AgentID, c.WeChat.Secret)

	var err error
	for i := 0; i < cnt; i++ {
		err = client.Send(corp.Message{
			ToUser:  message.Tos[i],
			MsgType: "text",
			Text:    corp.Content{Content: message.Content},
		})

		if err != nil {
			logger.Warningf("api send to %s fail: %v", message.Tos[i], err)
		} else {
			logger.Infof("api send to %s succ", message.Tos[i])
		}
	}

	render.Message(w, err)
}
