package mockentity

// Entity represent simple entity
type Entity struct {
	Title   string `json:"title" db:"title" schema:"title"`
	Content string `json:"content" db:"content" schema:"content"`
}

// NewEntityI create new interface wrapped a entity
func NewEntityI() interface{} {
	return &Entity{}
}
