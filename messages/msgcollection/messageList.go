package msgcollection

import (
	"strconv"
	"strings"

	"github.com/goatcms/goatcore/messages"
)

// MessageList represent a list of messages
type MessageList struct {
	list []string
}

// NewMessageList create new message list for a field
func NewMessageList() messages.MessageList {
	return &MessageList{
		list: []string{},
	}
}

// Add insert new element to message list
func (ml *MessageList) Add(msg string) {
	ml.list = append(ml.list, msg)
}

// GetAll return all messages
func (ml *MessageList) GetAll() []string {
	return ml.list
}

func (ml *MessageList) String() string {
	return "MessageList<" + strings.Join(ml.list, ", ") + ">"
}

// ToJSON return MessageList as json object
func (ml *MessageList) ToJSON() (json string) {
	tmp := make([]string, len(ml.list))
	for i, msg := range ml.list {
		tmp[i] = strconv.Quote(msg)
	}
	return "[" + strings.Join(tmp, ",") + "]"
}
