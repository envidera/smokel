package messenger

type Message struct {
	text    string
	example string
}

func NewMessage(text, example string) *Message {
	return &Message{
		text:    text,
		example: example,
	}
}
