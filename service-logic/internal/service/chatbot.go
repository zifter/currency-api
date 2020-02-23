package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zifter/currency/lib"
)

type ChatGateImpl struct {
	url string
}

func NewChatGateImpl(url string) *ChatGateImpl {
	return &ChatGateImpl{
		url: url,
	}
}

func (g *ChatGateImpl) SendMsg(msg *lib.ChatMessage) error {
	if msg == nil {
		return fmt.Errorf("message is nil")
	}

	r, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("cant marshal message: %w", err)
	}
	reader := bytes.NewReader(r)
	url := g.url + "/chat/message/"
	res, err := http.Post(url, "application/json", reader)
	if err != nil {
		return fmt.Errorf("cant post message to %v: %w", url, err)
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("wrong status code for post message %v: %v", url, res.Status)
	}
	return nil
}
