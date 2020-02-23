package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zifter/currency/lib"
	"github.com/zifter/currency/service-logic/internal/types"
)

const (
	getCurrencyPath = "/currency/"
	postPath        = "/chat/message/"
)

type ChatGate interface {
	SendMsg(msg *lib.ChatMessage) error
}

type Repo interface {
	IsEntryExists(ctx context.Context, e *types.RatePost) (bool, error)
	CreateEntry(ctx context.Context, e *types.RatePost) error
}

type CurrencyWatcher struct {
	CurrencyAPI string
	Chat        ChatGate
	Repo        Repo
}

func (w *CurrencyWatcher) CreateMessage(info *lib.FullCurrencyInfo) *lib.ChatMessage {
	msg := ""
	for k, data := range info.CurrencyAggregation {
		msg += k + "\n"
		msg += fmt.Sprintf("Курс Нацбанка РБ: %v\n", data.NationalBank.OfficialRate)
		msg += fmt.Sprintf("Лучший курс продажи в %v: %v\n", data.BankBest.Sell.BankName, data.BankBest.Sell.Value)
		msg += fmt.Sprintf("Лучший курс покупки в %v: %v\n", data.BankBest.Buy.BankName, data.BankBest.Buy.Value)
	}

	return &lib.ChatMessage{
		Message: msg,
	}
}

func (w *CurrencyWatcher) CheckUpdate() error {
	ctx := context.Background()
	url := w.CurrencyAPI + getCurrencyPath
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to get current currencu info: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("get currency wrong status code from %v: %v", url, resp.Status)
	}
	info := lib.NewFullCurrencyInfo()
	err = json.NewDecoder(resp.Body).Decode(info)
	if err != nil {
		return fmt.Errorf("cant decode response: %w", err)
	}

	msg := w.CreateMessage(info)
	entry := types.NewRatePost(msg.Message)
	yes, err := w.Repo.IsEntryExists(ctx, entry)
	if err != nil {
		return fmt.Errorf("cant check entry: %w", err)
	}
	if yes {
		return nil
	}

	err = w.Chat.SendMsg(msg)
	if err != nil {
		return fmt.Errorf("cant send message to chat: %w", err)
	}

	err = w.Repo.CreateEntry(ctx, entry)
	if err != nil {
		return fmt.Errorf("cant create post entry in db: %w", err)
	}

	return nil
}
