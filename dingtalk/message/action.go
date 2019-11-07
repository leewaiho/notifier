package message

type ActionCardMessage struct {
	MsgType MsgType           `json:"msgtype"`
	Content actionCardContent `json:"actionCard"`
}

type actionCardContent struct {
	msgContent
	Buttons          []*actionCardButton `json:"btns"`
	HideAvatar       DingTalkBool        `json:"hideAvatar"`
	ButtonHorizontal DingTalkBool        `json:"btnOrientation"`
	SingleTitle      string              `json:"singleTitle,omitempty"`
	SingleURL        string              `json:"singleURL,omitempty"`
}

type actionCardButton struct {
	Title string `json:"title"`
	Url   string `json:"actionURL"`
}

func NewActionCardMessage(title, content string) *ActionCardMessage {
	return &ActionCardMessage{
		MsgType: ActionCard,
		Content: actionCardContent{
			msgContent: msgContent{
				Title: title,
				Text:  content,
			},
			HideAvatar:       False,
			ButtonHorizontal: False,
		},
	}
}

func (acm *ActionCardMessage) SetButtonHorizontal(talkBool DingTalkBool) *ActionCardMessage {
	acm.Content.ButtonHorizontal = talkBool
	return acm
}

func (acm *ActionCardMessage) SetHideAvatar(talkBool DingTalkBool) *ActionCardMessage {
	acm.Content.HideAvatar = talkBool
	return acm
}

func (acm *ActionCardMessage) AddButton(title, url string) *ActionCardMessage {
	acm.Content.Buttons = append(acm.Content.Buttons, &actionCardButton{
		Title: title,
		Url:   url,
	})
	acm.Content.SingleTitle = ""
	acm.Content.SingleURL = ""
	return acm
}

func (acm *ActionCardMessage) SetSingleAction(singleText, singleURL string) *ActionCardMessage {
	acm.Content.SingleTitle = singleText
	acm.Content.SingleURL = singleURL
	acm.Content.Buttons = nil
	return acm
}
