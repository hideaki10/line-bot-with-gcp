package linebot

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func Webook(w http.ResponseWriter, r *http.Request) {

	LineChannelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	LineChannelAccessToken := os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")

	client, err := linebot.New(LineChannelSecret, LineChannelAccessToken)
	if err != nil {
		http.Error(w, "Error client", http.StatusBadRequest)
		log.Fatal(err)
		return
	}

	events, err := client.ParseRequest(r)
	if err != nil {
		http.Error(w, "Error parse request", http.StatusBadRequest)
		log.Fatal(err)
		return
	}

	for _, e := range events {
		switch e.Type {
		case linebot.EventTypeMessage:
			message := linebot.NewTextMessage("test")
			_, err := client.ReplyMessage(e.ReplyToken, message).Do()
			if err != nil {
				log.Println(err)
				return
			}
		}
	}

	fmt.Fprintln(w, "ok")

}
