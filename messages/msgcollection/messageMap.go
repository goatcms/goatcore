package msgcollection

import "github.com/goatcms/goatcore/messages"

// MessageMap represent messages for object
type MessageMap struct {
	m map[string]messages.MessageList
}

// NewMessageMap create new message map
func NewMessageMap() messages.MessageMap {
	return &MessageMap{
		m: make(map[string]messages.MessageList),
	}
}

// Get return single message list
func (mm *MessageMap) Get(key string) messages.MessageList {
	v, ok := mm.m[key]
	if !ok {
		v = NewMessageList()
		mm.m[key] = v
	}
	return v
}

// GetAll return map of MessageList
func (mm *MessageMap) GetAll() map[string]messages.MessageList {
	return mm.m
}

// Add add new message
func (mm *MessageMap) Add(key, msg string) {
	mm.Get(key).Add(msg)
}
