package loam

import (
	"fmt"
	"lineOAM/config"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

var chanSecret string = config.ChannelSecret
var bot *messaging_api.MessagingApiAPI

type source struct {
	Type   string `json:"type"`
	UserId string `json:"userId"`
}

func init() {
	var err error
	bot, err = messaging_api.NewMessagingApiAPI(
		config.ChannelToken,
	)
	if err != nil {
		panic(err)
	}
}

func ReceiveCallBack(ctx *gin.Context) {
	cb, err := webhook.ParseRequest(chanSecret, ctx.Request)
	if err != nil {
		return
	}
	fmt.Println(cb)
	for _, event := range cb.Events {
		switch e := event.(type) {
		case webhook.MessageEvent:
			source := e.Source
			fmt.Println(source)
			fmt.Printf("%T, %T, %+v", e, e.Source, e.Source)
		}
	}
}

func SendTextMessage(to string, text string) {
	bot.PushMessage(
		&messaging_api.PushMessageRequest{
			To: to,
			Messages: []messaging_api.MessageInterface{
				messaging_api.TextMessage{
					Text: text,
				},
			},
		},
		"",
	)

}
