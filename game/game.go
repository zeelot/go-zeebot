package game

import (
	"github.com/zeelot/go-ircevent"
	"log"
	"strconv"
	"strings"
	"time"
)

type Player struct {
	Bot      Bot
	Strategy Strategy
}

func CreateIRCPlayer(bot Bot, strategy Strategy) {
	irccon1 := irc.IRC(bot.Name, bot.Name)
	irccon1.VerboseCallbackHandler = true
	irccon1.Debug = false

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
		time.Sleep(3 * time.Second)

		event := OftbotEvent(*e)
		log.Printf("Got event %v\n\n", event.GetMessage())
		if event.IsGameSuggestion() {
			log.Println("Detected game suggestion")
			strategy.Reset()
			irccon1.Privmsg("#cosmic-rift", "@oftbot join")
			irccon1.Privmsg("#cosmic-rift", "Strategy: Safe")
		}

		if event.IsTimeToRoll(bot) {
			log.Println("Detected time to roll")
			irccon1.Privmsg("#cosmic-rift", "@oftbot roll")
		}

		if event.IsTimeToKeep(bot) {
			log.Println("Detected time to keep")
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
