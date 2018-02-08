package msgcollection

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/goatcms/goatcore/messages"
)

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

func (mm *MessageMap) String() string {
	var list = []string{}
	for k, v := range mm.m {
		list = append(list, fmt.Sprintf("%v:%v", k, v))
	}
	return "MessageMap<" + strings.Join(list, ", ") + ">"
}

// ToJSON return MessageMap as json object
func (mm *MessageMap) ToJSON() (json string) {
	tmp := make([]string, len(mm.m))
	i := 0
	for key, msg := range mm.m {
		tmp[i] = strconv.Quote(key) + ":" + msg.ToJSON()
		i++
	}
	return "{" + strings.Join(tmp, ",") + "}"
}
