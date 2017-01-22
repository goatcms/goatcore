package messages

// MessageMap contains object errors
type MessageMap interface {
	Get(key string) MessageList
	GetAll() map[string]MessageList
	Add(key, msg string)
}

// MessageList represent list of field errors
type MessageList interface {
	GetAll() []string
	Add(msgkey string)
}
