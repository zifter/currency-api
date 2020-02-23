package service

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zifter/currency/lib"
)

type postBotHandler struct {
	testCase string
	data     string
}

func (h *postBotHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch h.testCase {
	case "failed post":
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "")
	}

	body, _ := ioutil.ReadAll(r.Body)
	h.data = string(body)
}

func Test_Send(t *testing.T) {
	postBotHandler := &postBotHandler{}
	botAPI := httptest.NewServer(postBotHandler)
	defer botAPI.Close()

	gate := NewChatGateImpl(botAPI.URL)

	cases := []struct {
		testCase string
		wantErr  bool
		data     string
		prepare  func()
		msg      *lib.ChatMessage
	}{
		{
			"error",
			true,
			"",
			func() {},
			nil,
		},
		{
			"failed post",
			true,
			`{"Message":"failed post body"}`,
			func() {},
			&lib.ChatMessage{
				Message: "failed post body",
			},
		},
		{
			"correct",
			false,
			`{"Message":"correct body"}`,
			func() {},
			&lib.ChatMessage{
				Message: "correct body",
			},
		},
	}

	for _, ts := range cases {
		t.Run(ts.testCase, func(t *testing.T) {
			postBotHandler.testCase = ts.testCase
			ts.prepare()

			err := gate.SendMsg(ts.msg)
			if (err != nil) != ts.wantErr {
				t.Errorf(`Error mismatch %v`, err)
			}
			if postBotHandler.data != ts.data {
				t.Errorf(`Data mismatch %v != %v`, postBotHandler.data, ts.data)
			}
		})
	}
}
