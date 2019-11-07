package message

type MarkdownMessage struct {
	MsgType MsgType    `json:"msgtype"`
	Content msgContent `json:"markdown"`
	At      *At        `json:"at"`
}

func NewMarkdownMessage(title, content string, atAll bool, atMobiles ...string) *MarkdownMessage {
	return &MarkdownMessage{
		MsgType: Markdown,
		Content: msgContent{
			Title: title,
			Text:  content,
		},
		At: NewAt(atAll, atMobiles...),
	}
}
