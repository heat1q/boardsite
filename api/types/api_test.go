package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalMsg(t *testing.T) {
	tests := []struct {
		msg  string
		want *Message
		err  error
	}{
		{"", nil, assert.AnError},
		{`{{"content": {}`, nil, assert.AnError},
		{`{"content": {"type": "0}}`, nil, assert.AnError},
		{`{}`, &Message{}, nil},
		{`{"type":"sometype"}`, &Message{Type: "sometype"}, nil},
		{
			`{"type":"sometype","sender":"heat","content":{}}`,
			&Message{Type: "sometype", Sender: "heat"},
			nil,
		},
	}

	for _, test := range tests {
		m, err := UnmarshalMessage([]byte(test.msg))
		if test.err != nil {
			assert.Error(t, err)
			continue
		} else {
			assert.NoError(t, err)
		}
		assert.Equal(t, test.want.Type, m.Type, "incorrect marshalling of message type")
		assert.Equal(t, test.want.Sender, m.Sender, "incorrect marshalling of message sender")
	}
}

func TestUnmarshalMsgContent(t *testing.T) {
	tests := []struct {
		msg  string
		want Stroke
		err  error
	}{
		{"", Stroke{}, assert.AnError},
		{"{}", Stroke{}, assert.AnError},
		{`{{"content": {}`, Stroke{}, assert.AnError},
		{`{"content": null}`, Stroke{}, assert.AnError},
		{`{"content": {"type": "0}}`, Stroke{}, assert.AnError},
		{
			`{"type":"sometype","sender":"heat","content":{"type":0,"id":"id1","userId":"user1"}}`,
			Stroke{ID: "id1", UserID: "user1", Type: 0},
			nil,
		},
	}

	for _, test := range tests {
		var c Stroke
		if test.err != nil {
			assert.Error(t, UnmarshalMsgContent([]byte(test.msg), &c))
		} else {
			assert.NoError(t, UnmarshalMsgContent([]byte(test.msg), &c))
		}
		assert.Equal(t, test.want, c, "incorrect unmarshalling of message content")
	}
}

func TestMarshalMessage(t *testing.T) {
	tests := []struct {
		msgType string
		sender  string
		content interface{}
		want    string
	}{
		{"", "", []string{}, `{"type":"","content":[]}`},
		{"", "", []string{"pid1", "pid2"}, `{"type":"","content":["pid1","pid2"]}`},
	}

	for _, test := range tests {
		m := NewMessage(test.content, test.msgType, test.sender)
		menc, err := json.Marshal(m)
		assert.NoError(t, err)
		assert.Equal(t, test.want, string(menc), "incorrect marshalling of message")
	}
}
