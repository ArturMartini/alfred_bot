package control

import (
	"github.com/nlopes/slack"
	"strings"
)

var (
	commands = map[string]func(Message) Message{
		"hello":     hello,
		"help":      help,
		"configure": configure,
		"send":      send,
	}

	users = map[string]*User{
	}

)

type Message struct {
	Name string
	User string
	Text string
}

type User struct {
	User         string
	Email        string
	Password     string
	Template     string
	TemplateArgs string
	EmailTo      string
	EmailCco     string
}

func hello(_ Message) Message {
	return Message{Text: "My name is Alfred and i am is your new butler. What do you want?"}
}

func help(_ Message) Message {
	return Message{Text: "I tell you what i can do, but you not is Batman HaHaHa, then:\n" +
		"- help\n" +
		"- configure\n" +
		"- send\n"}
}

func configure(message Message) Message {
	text := "Hello Sr. " + message.Name + "\n" +
		"I need some information for configure automation of send email\n" +
		"- email (your@email.com)\n" +
		"- password \n" +
		"- template \n" +
		"- args template \n" +
		"- to \n" +
		"- copy to \n"

	return Message{
		Name: message.Name,
		Text: text,
	}
}

func send(_ Message) Message {
	return Message{Text: "My name is Alfred and i am is your new butler. What do you want?"}
}

func ProcessMessage(ev *slack.MessageEvent) Message {
	message := translateEventToMessage(ev)
	cmd := commands[message.Text];
	if cmd == nil {
		return Message{
			message.Name,
			extractUser(message.Text),
			"Sorry Sr. " + message.Name + " Im not understand what you say",
		}
	}
	return cmd(message)
}

func translateEventToMessage(event *slack.MessageEvent) Message {
	text := event.Text
	name := event.Name
	user := extractUser(text)
	text = removeUser(text, user)
	return Message{
		Name: name,
		User: user,
		Text: text,
	}
}

func removeUser(text, user string) string {
	newText := strings.Replace(text, "<"+user+">", "", -1)
	return strings.TrimSpace(newText)
}

func extractUser(text string) string {
	end := strings.LastIndexAny(text, ">")
	if end < 0 {
		return ""
	} else {
		return text[1:end]
	}
}
