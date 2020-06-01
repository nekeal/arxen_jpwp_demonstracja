package main

import (
	"testing"
	"time"
)

func TestHandler_Handle(t *testing.T) {
	msgs := []Message{
		{
			Text:   "message Uno",
			ChatID: "1",
		},
		{
			Text:   "message duo",
			ChatID: "2",
		},
		{
			Text:   "message tre",
			ChatID: "2",
		},
		{
			Text:   "message erro",
			ChatID: "4",
		},
	}

	chats := make(map[string]*Chat)

	chats["1"] = &Chat{
		ChatID:      "1",
		MessageList: []Message{},
	}

	chats["2"] = &Chat{
		ChatID:      "2",
		MessageList: []Message{},
	}

	type fields struct {
		handler *Handler
	}
	tests := []struct {
		name     string
		messages []Message
		chats    map[string]*Chat
		fields   fields
		wantErr  bool
		want     string
	}{
		{
			name:     "basic_test",
			messages: msgs,
			chats:    chats,
			fields: struct {
				handler *Handler
			}{handler: NewHandler()},
			wantErr: true,
			want:    "message Uno",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timeout := time.After(3 * time.Second)
			done := make(chan bool)

			h := tt.fields.handler
			h.ChatList = chats
			go func() {
				for _, msg := range msgs {
					h.receivedMessageChan <- msg
				}
				done <- true
			}()

			go h.Handle()

			select {
			case <-timeout:
				t.Errorf("Handle() function timedout")
			case <-done:
				got := h.ChatList["1"].MessageList[0].Text
				if got != tt.want {
					t.Errorf("Handle() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
