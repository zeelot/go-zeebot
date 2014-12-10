package main

import (
	"github.com/zeelot/go-ircevent"
	"log"
	"regexp"
	"strings"
)

type OftbotEvent irc.Event

func (event OftbotEvent) IsGameSuggestion() bool {
	r, _ := regexp.Compile("^@([a-zA-Z]+) suggests a new game of 1, 4, 24!")
	return r.MatchString(event.GetMessage())
}

func (event OftbotEvent) IsTimeToRoll(bot Bot) bool {
	if !event.IsMessageIntendedFor(bot) {
		return false
	}

	r, _ := regexp.Compile("^@([a-zA-Z]+), it's your turn next")
	if r.MatchString(event.GetMessage()) {
		return true
	}

	r, _ = regexp.Compile("^@([a-zA-Z]+), you're up first")
	if r.MatchString(event.GetMessage()) {
		return true
	}

	return false
}

func (event OftbotEvent) IsTimeToKeep(bot Bot) bool {
	return event.IsRollBy(bot)
}

func (event OftbotEvent) GetValuesToKeep() []string {
	//@Zeelot rolled: 5, 3, 5, 4, 4, 6.
	r, _ := regexp.Compile("^@[a-zA-Z]+ rolled: ([0-9]), ([0-9]), ([0-9]), ([0-9]), ([0-9]), ([0-9]).$")
	match := r.FindStringSubmatch(event.GetMessage())
	return match[1:]
}

func (event OftbotEvent) IsMessageIntendedFor(bot Bot) bool {
	pattern := "^@:name, "
	r, _ := regexp.Compile(strings.Replace(pattern, ":name", bot.Name, 1))
	return r.MatchString(event.GetMessage())
}

func (event OftbotEvent) IsRollBy(bot Bot) bool {
	pattern := "^@:name rolled"
	r, _ := regexp.Compile(strings.Replace(pattern, ":name", bot.Name, 1))
	return r.MatchString(event.GetMessage())
}

func (event OftbotEvent) GetMessage() string {
	return event.Arguments[1]
}

type Bot struct {
	Name string
}

func main() {
	bot := Bot{Name: "zeebot"}

	irccon1 := irc.IRC(bot.Name, bot.Name)
	irccon1.VerboseCallbackHandler = true
	irccon1.Debug = true

	err := irccon1.Connect("irc.freenode.net:6667")
	if err != nil {
		log.Fatal(err.Error())
	}

	irccon1.AddCallback("001", func(e *irc.Event) { irccon1.Join("#cosmic-rift") })

	log.Println("hello")

	irccon1.AddCallback("NOTICE", func(e *irc.Event) {
		if e.Nick != "oftbot" {
			return
		}

		event := OftbotEvent(*e)
		if event.IsGameSuggestion() {
			irccon1.Privmsg("#cosmic-rift", "@oftbot join")
		}

		if event.IsTimeToRoll(bot) {
			irccon1.Privmsg("#cosmic-rift", "@oftbot roll")
		}

		if event.IsTimeToKeep(bot) {
			template := "@oftbot keep :numbers"
			numbers := strings.Join(event.GetValuesToKeep(), "")
			response := strings.Replace(template, ":numbers", numbers, 1)
			irccon1.Privmsg("#cosmic-rift", response)
		}
	})

	irccon1.Loop()
}
