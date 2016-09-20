package typemsg

import "github.com/goatcms/goat-core/types"

// MessageMap represent messages for object
type MessageMap struct {
	m map[string]types.MessageList
}

// NewMessageMap create new message map
func NewMessageMap() types.MessageMap {
	return &MessageMap{
		m: make(map[string]types.MessageList),
	}
}

// Get return single message list
func (mm *MessageMap) Get(key string) types.MessageList {
	v, ok := mm.m[key]
	if !ok {
		v = NewMessageList()
		mm.m[key] = v
	}
	return v
}

// GetAll return map of MessageList
func (mm *MessageMap) GetAll() map[string]types.MessageList {
	return mm.m
}

// Add add new message
func (mm *MessageMap) Add(key, msg string) {
	mm.Get(key).Add(msg)
}
