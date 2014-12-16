package game

import (
	"github.com/zeelot/go-ircevent"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Sheriff struct {
	Name string
}

func NewSheriff(name string) {
	sheriff := &Sheriff{
		Name: name,
	}

	irccon := irc.IRC(sheriff.Name, sheriff.Name)
	irccon.VerboseCallbackHandler = true
	irccon.Debug = false

	err := irccon.Connect("irc.freenode.net:6667")
	if err != nil {
		log.Fatal(err.Error())
	}

	irccon.AddCallback("001", func(e *irc.Event) { irccon.Join("#cosmic-rift") })

	irccon.AddCallback("PRIVMSG", func(e *irc.Event) {
		// Just prevent bots from hammering IRC.
		time.Sleep(1 * time.Second)

		event := OftbotEvent(*e)
		if event.IsSummonPosseCommand() {
			irccon.Privmsg("#cosmic-rift", "Hold up! Making a phone call…")
			number := rand.Intn(100000)
			botNameTemplate := "zeebotposse:number"
			botName := strings.Replace(botNameTemplate, ":number", strconv.Itoa(number), 1)
			bot := Bot{Name: botName}
			log.Printf("STARTING BOT: %v\n\n\n", bot)
			strategy := &SafeStrategy{}
			go CreateIRCPlayer(bot, strategy)
		}
	})

	irccon.Loop()
}

func (self *Sheriff) SummonPosse() {

}
