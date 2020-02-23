package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"github.com/zifter/currency/lib"
)

var log = logrus.New().WithFields(logrus.Fields{
	"service-name": "chatbot",
})

func SendMessage(bot *tgbotapi.BotAPI, groupID int64, msg *lib.ChatMessage) error {
	log.Printf("Send message %v\n", msg)
	m := tgbotapi.NewMessage(groupID, msg.Message)
	_, err := bot.Send(m)

	if err != nil {
		log.Println(err)
	}
	return err
}

func ReadChatMessage(w http.ResponseWriter, req *http.Request) (*lib.ChatMessage, error) {
	msg := &lib.ChatMessage{}
	err := json.NewDecoder(req.Body).Decode(msg)

	if err != nil {
		return nil, fmt.Errorf("cant decode message: %w", err)
	}

	return msg, nil
}

func main() {
	log.Println("Chat bot start")
	config := lib.LoadConfig()
	botToken := config.Chatbot.TelegramBotToken
	groupID := config.Chatbot.TelegramGroupID

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	quit := make(chan bool)
	c := make(chan *lib.ChatMessage, 1)
	go func() {
		for {
			select {
			case msg := <-c:
				go SendMessage(bot, groupID, msg)
			case <-quit:
				quit <- true
				return
			}
		}
	}()

	http.HandleFunc("/chat/message/", func(w http.ResponseWriter, req *http.Request) {
		msg, err := ReadChatMessage(w, req)
		if err != nil {
			log.Errorf("Cant read message: %v", err)
			return
		}
		log.Info("Receive message")
		go func() {
			c <- msg
		}()
	})

	log.Println("Start http on: ", config.Chatbot.Port)

	// start server
	{
		err := http.ListenAndServe(config.Chatbot.Host+":"+config.Chatbot.Port, nil)
		if err != nil {
			log.Panic(err)
		}
		quit <- true
	}

	<-quit
}
