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
	for _, event := range cb.Events {
		switch e := event.(type) {
		case webhook.MessageEvent:
			var uId string = getMsgUid(e.Source)
			fmt.Println("uId: ", uId)
			SendTextMessage(uId, "test")
		}
	}
}

func getMsgUid(s webhook.SourceInterface) (uid string) {
	switch source := s.(type) {
	case webhook.UserSource:
		uid = source.UserId
	case webhook.GroupSource:
		uid = source.GroupId
	case webhook.RoomSource:
		uid = source.RoomId
	}
	return uid
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
