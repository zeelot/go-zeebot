package main

import (
	"github.com/zeelot/go-zeebot/game"
	"time"
)

func main() {
	bot := game.Bot{Name: "zeebot-clever"}
	strategy := &game.SafeStrategy{}
	go game.CreateIRCPlayer(bot, strategy)

	bot2 := game.Bot{Name: "zeebot-dummy"}
	strategy2 := &game.DummyStrategy{}
	go game.CreateIRCPlayer(bot2, strategy2)

	for {
		// Just let the other functions do their thing forever.
		time.Sleep(1 * time.Second)
	}
}
