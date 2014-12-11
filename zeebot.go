package main

import (
	"github.com/zeelot/go-ircevent"
	"github.com/zeelot/zeebot/bot"
	"github.com/zeelot/zeebot/strategy"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	theBot := bot.Bot{Name: "zeebot-clever"}
	strategy := strategy.SafeStrategy{}

	irccon1 := irc.IRC(theBot.Name, theBot.Name)
	irccon1.VerboseCallbackHandler = true
	irccon1.Debug = true

	err := irccon1.Connect("irc.freenode.net:6667")
	if err != nil {
		log.Fatal(err.Error())
	}

	irccon1.AddCallback("001", func(e *irc.Event) { irccon1.Join("#cosmic-rift") })

	irccon1.AddCallback("NOTICE", func(e *irc.Event) {
		if e.Nick != "oftbot" {
			return
		}

		// Just prevent bots from hammering IRC.
		time.Sleep(1 * time.Second)

		event := bot.OftbotEvent(*e)
		if event.IsGameSuggestion() {
			strategy.Reset()
			irccon1.Privmsg("#cosmic-rift", "@oftbot join")
			irccon1.Privmsg("#cosmic-rift", "Strategy: Safe")
		}

		if event.IsTimeToRoll(theBot) {
			irccon1.Privmsg("#cosmic-rift", "@oftbot roll")
		}

		if event.IsTimeToKeep(theBot) {
			template := "@oftbot keep :numbers"
			toKeep := strategy.ChooseDice(event.GetRollValues())
			stringNumbers := []string{}
			for _, intNumber := range toKeep {
				stringNumbers = append(stringNumbers, strconv.Itoa(intNumber))
			}
			numbers := strings.Join(stringNumbers, "")
			response := strings.Replace(template, ":numbers", numbers, 1)
			irccon1.Privmsg("#cosmic-rift", response)
		}
	})

	irccon1.Loop()
}
