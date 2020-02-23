//go:generate mockgen -source=currency_watcher.go -destination=mock/currency_watcher.go
package service

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_service "github.com/zifter/currency/service-logic/internal/service/mock"

	"github.com/golang/mock/gomock"
	"github.com/zifter/currency/lib"
)

type getCurrencyHandler struct {
	testCase string
	data     string
}

func (h *getCurrencyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch h.testCase {
	case "error":
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "")
	default:
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, h.data)
	}
}

func Test_CheckUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	currencyHandler := &getCurrencyHandler{}
	currencySvr := httptest.NewServer(currencyHandler)
	defer currencySvr.Close()

	mockChat := mock_service.NewMockChatGate(ctrl)
	mockRepo := mock_service.NewMockRepo(ctrl)
	w := CurrencyWatcher{
		CurrencyAPI: currencySvr.URL,
		Chat:        mockChat,
		Repo:        mockRepo,
	}

	cases := []struct {
		testCase string
		wantErr  bool
		data     string
		prepare  func()
	}{
		{
			"error",
			true,
			"",
			func() {},
		},
		{
			"correct",
			false,
			`{"CurrencyAggregation":{"usd":{"NationalBank":{"Cur_ID":145,"Cur_Abbreviation":"USD","Cur_Name":"Доллар США","Cur_OfficialRate":2.1986},"BankBest":{"Sell":{"BankName":"","Value":0},"Buy":{"BankName":"","Value":0}}}}}`,
			func() {
				msg := &lib.ChatMessage{
					Message: "usd\nКурс Нацбанка РБ: 2.1986\nЛучший курс продажи в : 0\nЛучший курс покупки в : 0\n",
				}
				mockChat.EXPECT().SendMsg(msg).Return(nil)
				mockRepo.EXPECT().IsEntryExists(gomock.Any(), gomock.Any()).Return(false, nil)
				mockRepo.EXPECT().CreateEntry(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			"entry exists failed",
			true,
			`{"CurrencyAggregation":{"usd":{"NationalBank":{"Cur_ID":145,"Cur_Abbreviation":"USD","Cur_Name":"Доллар США","Cur_OfficialRate":2.1986},"BankBest":{"Sell":{"BankName":"","Value":0},"Buy":{"BankName":"","Value":0}}}}}`,
			func() {
				mockRepo.EXPECT().IsEntryExists(gomock.Any(), gomock.Any()).Return(false, fmt.Errorf("test"))
			},
		},
		{
			"already exists",
			false,
			`{"CurrencyAggregation":{"usd":{"NationalBank":{"Cur_ID":145,"Cur_Abbreviation":"USD","Cur_Name":"Доллар США","Cur_OfficialRate":2.1986},"BankBest":{"Sell":{"BankName":"","Value":0},"Buy":{"BankName":"","Value":0}}}}}`,
			func() {
				mockRepo.EXPECT().IsEntryExists(gomock.Any(), gomock.Any()).Return(true, nil)
			},
		},
	}

	for _, ts := range cases {
		t.Run(ts.testCase, func(t *testing.T) {
			currencyHandler.testCase = ts.testCase
			currencyHandler.data = ts.data
			ts.prepare()

			err := w.CheckUpdate()
			if (err != nil) != ts.wantErr {
				t.Errorf(`Error mismatch %v`, err)
			}
		})
	}
}
