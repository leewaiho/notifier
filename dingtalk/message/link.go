package message

type LinkMessage struct {
	MsgType MsgType            `json:"msgtype"`
	Content linkMessageContent `json:"link"`
}

type linkMessageContent struct {
	msgContent
	PicUrl     string `json:"picUrl"`
	MessageUrl string `json:"messageUrl"`
}

func NewLinkMessage(title, content, messageUrl, picUrl string) *LinkMessage {
	return &LinkMessage{
		MsgType: Link,
		Content: linkMessageContent{
			msgContent: msgContent{
				Title: title,
				Text:  content,
			},
			MessageUrl: messageUrl,
			PicUrl:     picUrl,
		},
	}
}
