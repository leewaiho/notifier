package message

type MsgType string

type DingTalkBool string

const (
	False DingTalkBool = "0"
	True  DingTalkBool = "1"
)

const (
	Text       MsgType = "text"
	Link       MsgType = "link"
	Markdown   MsgType = "markdown"
	ActionCard MsgType = "actionCard"
	FeedCard   MsgType = "feedCard"
)

type At struct {
	AtMobiles []string `json:"atMobiles"`
	AtAll     bool     `json:"isAtAll"`
}

func NewAt(atAll bool, atMobiles ...string) *At {
	return &At{
		AtMobiles: atMobiles,
		AtAll:     atAll,
	}
}

type msgContent struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
