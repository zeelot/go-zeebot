package game

import (
	"github.com/zeelot/go-ircevent"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Bot struct {
	Name string
}

type OftbotEvent irc.Event

func (event OftbotEvent) IsGameSuggestion() bool {
	r, _ := regexp.Compile(`^@([a-zA-Z\-_]+) suggests a new game of 1, 4, 24!`)
	return r.MatchString(event.GetMessage())
}

func (event OftbotEvent) IsTimeToRoll(bot Bot) bool {
	if !event.IsMessageIntendedFor(bot) {
		return false
	}

	r, _ := regexp.Compile(`^@([a-zA-Z\-_]+), it's your turn next`)
	if r.MatchString(event.GetMessage()) {
		return true
	}

	r, _ = regexp.Compile(`^@([a-zA-Z\-_]+), you're up first`)
	if r.MatchString(event.GetMessage()) {
		return true
	}

	return false
}

func (event OftbotEvent) IsTimeToKeep(bot Bot) bool {
	return event.IsRollBy(bot) && len(event.GetRollValues()) > 1
}

func (event OftbotEvent) GetRollValues() []int {
	var values []int
	parts := strings.Split(event.GetMessage(), ":")
	r, _ := regexp.Compile(`([0-9])`)
	match := r.FindAllString(parts[1], 6)

	for _, value := range match {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err.Error())
		}
		values = append(values, intValue)
	}

	return values
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
