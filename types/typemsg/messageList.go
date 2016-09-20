package typemsg

import "github.com/goatcms/goat-core/types"

// MessageList represent a list of messages
type MessageList struct {
	list []string
}

// NewMessageList create new message list for a field
func NewMessageList() types.MessageList {
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
