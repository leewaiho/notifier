package dingtalk

import (
	"gitee.com/LeeWaiHo/tool-dingtalk/message"
	"testing"
)

const (
	MarkdownText = "## HelloWorld\n * foo\n * bar\n"
	TestURL      = "http://www.bilibili.com"
)

var (
	robot = NewRobot(&Option{
		AccessToken: "",
		SecretKey:   "",
	})
)

func TestSendMarkdown_Text(t *testing.T) {
	t.Log(robot.Send(message.NewMarkdownMessage("消息预览显示", MarkdownText, true)))
}

func TestSendActionCard_Normal(t *testing.T) {
	msg := message.NewActionCardMessage("消息预览显示", MarkdownText)
	msg.AddButton("好看", TestURL)
	msg.AddButton("还行", TestURL)
	t.Log(robot.Send(msg))
}

func TestSendActionCard_SingleAction(t *testing.T) {
	msg := message.NewActionCardMessage("消息预览显示", MarkdownText)
	msg.SetSingleAction("阅读全文", TestURL)
	t.Log(robot.Send(msg))
}

func TestSendActionCard_HideAvatar(t *testing.T) {
	msg := message.NewActionCardMessage("消息预览显示", MarkdownText)
	msg.SetHideAvatar(message.True)
	msg.SetSingleAction("阅读全文", TestURL)
	t.Log(robot.Send(msg))
}

func TestSendActionCard_HorizontalButtons(t *testing.T) {
	msg := message.NewActionCardMessage("消息预览显示", MarkdownText)
	msg.AddButton("好看", TestURL)
	msg.AddButton("还行", TestURL)
	msg.SetButtonHorizontal(message.True)
	t.Log(robot.Send(msg))
}

func TestSendTextMessage_Normal(t *testing.T) {
	t.Log(robot.Send(message.NewTextMessage("HelloWorld", true)))
}

func TestSendLinkMessage_Normal(t *testing.T) {
	t.Log(robot.Send(message.NewLinkMessage("Hi", "test", "http://www.bilibili.com", "https://s1.hdslb.com/bfs/static/jinkela/home/asserts/ic_launcher.png")))
}

func TestSendFeedCardMessage_Single(t *testing.T) {
	m := message.NewFeedCardMessage("你好鸭", "http://www.bilibili.com", "https://s1.hdslb.com/bfs/static/jinkela/home/asserts/ic_launcher.png")
	t.Log(robot.Send(m))
}

func TestSendFeedCardMessage_Normal(t *testing.T) {
	m := message.NewFeedCardMessage("你好鸭", "http://www.bilibili.com", "https://s1.hdslb.com/bfs/static/jinkela/home/asserts/ic_launcher.png")
	m.AddItem("鸭鸭", "http://www.bilibili.com", "https://s1.hdslb.com/bfs/static/jinkela/home/asserts/ic_launcher.png")
	t.Log(robot.Send(m))
}
