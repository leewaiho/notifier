package message

type FeedCardMessage struct {
	MsgType MsgType         `json:"msgtype"`
	Content feedCardContent `json:"feedCard"`
}

type feedCardContent struct {
	ItemList []*feedCardItem `json:"links"`
}

type feedCardItem struct {
	Title      string `json:"title"`
	MessageURL string `json:"messageURL"`
	PictureURL string `json:"picURL"`
}

func NewFeedCardMessage(title, url, picUrl string) *FeedCardMessage {
	message := &FeedCardMessage{
		MsgType: FeedCard,
		Content: feedCardContent{
			ItemList: []*feedCardItem{},
		},
	}
	message.AddItem(title, url, picUrl)
	return message
}

func (f *FeedCardMessage) AddItem(title, url, picUrl string) *FeedCardMessage {
	f.Content.ItemList = append(f.Content.ItemList, &feedCardItem{
		Title:      title,
		MessageURL: url,
		PictureURL: picUrl,
	})
	return f
}
