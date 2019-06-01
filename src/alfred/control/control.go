package control

import (
	"alfred/canonical"
	"alfred/email"
	"alfred/repository"
	"github.com/nlopes/slack"
	"strings"
)

var (
	commands = map[string]func(Message) Message{
		"hello":         hello,
		"help":          help,
		"configure":     configure,
		"configure add": configure_add,
		"send":          send,
	}

	users = map[string]*User{
	}

	repo repository.Repository
)

func init() {
	repo = repository.New()
}

type Message struct {
	Nick string
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

func hello(message Message) Message {
	return Message{Text: "Welcome " + message.Nick + "! My name is Alfred and i am is your new butler. What do you want? Little Robin.."}
}

func help(_ Message) Message {
	return Message{Text: "I tell you what i can do, but you not is Batman HaHaHa, then:\n" +
		"- help\n" +
		"- configure\n" +
		"- send\n"}
}

func configure(message Message) Message {
	text := "Hello Sr. " + message.Nick + "\n" +
		"I need some information for configure automation of send email\n" +
		"Execute command above:\n" +
		"configure add -email your@email.com -pass your_password -template your_template"

	return Message{
		Nick: message.Nick,
		Text: text,
	}
}

func configure_add(message Message) Message {
	text := ""
	text = strings.Replace(text, "configure add", "", -1)
	text = strings.Replace(text, "-", "", -1)
	text = strings.Replace(text, "  ", " ", -1)
	cmds := strings.Split(text, " ")

	indexEmail := 0
	indexPass := 0
	indexTemplate := 0

	for index, key := range cmds {
		switch key {
		case "email":
			indexEmail += index
		case "pass":
			indexPass += index
		case "template":
			indexTemplate += index
		}
	}

	config := canonical.Configuration{
		Email:    cmds[indexEmail],
		Pass:     cmds[indexPass],
		Template: cmds[indexTemplate],
	}

	err := repo.Update(config); if err != nil {
		return Message{
			Nick: message.Nick,
			Text: "Sorry Sr " + message.Nick + ". I'm busy with billions of other things you told me to do. You really think it's batman.",
		}
	}

	return Message{
		Nick: message.Nick,
		Text: "Your configuration haas been saved, Little Robin.",
	}

}

func send(message Message) Message {
	success := email.Send()
	if success {
		return Message{
			Nick: message.Nick,
			Text: "Your email has been sent.",
		}
	} else {
		return Message{
			Nick: message.Nick,
			Text: "Error ocurred when try send email, try again later.",
		}
	}
}

func ProcessMessage(ev *slack.MessageEvent) Message {
	message := translateEventToMessage(ev)
	cmd := commands[message.Text];
	if cmd == nil {
		return Message{
			message.Nick,
			message.User,
			"Sorry Sr " + message.Nick + ". Im not understand what you say",
		}
	}
	return cmd(message)
}

func translateEventToMessage(event *slack.MessageEvent) Message {
	text := event.Text
	nick := "<@" + event.User + ">"
	user := event.User
	return Message{
		Nick: nick,
		User: user,
		Text: text,
	}
}

func removeUser(text, user string) string {
	newText := strings.Replace(text, "<"+user+">", "", -1)
	return strings.TrimSpace(newText)
}
