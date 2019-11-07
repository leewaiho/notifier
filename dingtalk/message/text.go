package message

type TextMessage struct {
	MsgType MsgType     `json:"msgtype"`
	Content TextContent `json:"text"`
	At      *At         `json:"at"`
}

type TextContent struct {
	Content string `json:"content"`
}

func NewTextMessage(content string, atAll bool, atMobiles ...string) *TextMessage {
	message := &TextMessage{
		MsgType: Text,
		Content: TextContent{
			Content: content,
		},
		At: NewAt(atAll, atMobiles...),
	}
	if message.At.AtAll {
		message.Content.Content += "\n\n"
	}
	return message
}
