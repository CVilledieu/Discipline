package messages

type Message struct {
	content string
	UserID  uint16
}

func (m *Message) GetContent() string {
	return m.content
}

func NewMessage(content string, userID uint16) *Message {
	return &Message{content: content, UserID: userID}
}

func (m *Message) UpdateContent(content string) {
	m.content = content
}
