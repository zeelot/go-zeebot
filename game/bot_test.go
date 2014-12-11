package game

import (
	//"fmt"
	"github.com/zeelot/go-ircevent"
	"testing"
)

func GetSampleGameSuggestionEvent() OftbotEvent {
	return OftbotEvent(irc.Event{
		Code:   "NOTICE",
		Raw:    ":oftbot!~oftbot@WoPRCentral.jonathan-hanson.org NOTICE #cosmic-rift :@Zeelot suggests a new game of 1, 4, 24!  Who's in?  Type '@oftbot join' to play!",
		Nick:   "oftbot",
		Host:   "WoPRCentral.jonathan-hanson.org",
		Source: "oftbot!~oftbot@WoPRCentral.jonathan-hanson.org",
		User:   "~oftbot",
		Arguments: []string{
			"#cosmic-rift",
			"@Zeelot suggests a new game of 1, 4, 24!  Who's in?  Type '@oftbot join' to play!",
		},
	})
}

func GetSampleTimeToRollEvent() OftbotEvent {
	return OftbotEvent(irc.Event{
		Code:   "NOTICE",
		Raw:    ":oftbot!~oftbot@WoPRCentral.jonathan-hanson.org NOTICE #cosmic-rift :@triplepoint, you're up first.  Type '@oftbot roll' to take your first roll.",
		Nick:   "oftbot",
		Host:   "WoPRCentral.jonathan-hanson.org",
		Source: "oftbot!~oftbot@WoPRCentral.jonathan-hanson.org",
		User:   "~oftbot",
		Arguments: []string{
			"#cosmic-rift",
			"@triplepoint, you're up first.  Type '@oftbot roll' to take your first roll.",
		},
	})
}

func GetSampleTimeToKeepEvent() OftbotEvent {
	return OftbotEvent(irc.Event{
		Code:   "NOTICE",
		Raw:    ":oftbot!~oftbot@WoPRCentral.jonathan-hanson.org NOTICE #cosmic-rift :@triplepoint rolled: 2, 1, 2, 2, 6, 6.",
		Nick:   "oftbot",
		Host:   "WoPRCentral.jonathan-hanson.org",
		Source: "oftbot!~oftbot@WoPRCentral.jonathan-hanson.org",
		User:   "~oftbot",
		Arguments: []string{
			"#cosmic-rift",
			"@triplepoint rolled: 2, 1, 2, 2, 6, 6.",
		},
	})
}

func GetSampleSecondRollEvent() OftbotEvent {
	return OftbotEvent(irc.Event{
		Code:   "NOTICE",
		Raw:    ":oftbot!~oftbot@WoPRCentral.jonathan-hanson.org NOTICE #cosmic-rift :@triplepoint rolled: 2, 1, 2, 2, 6.",
		Nick:   "oftbot",
		Host:   "WoPRCentral.jonathan-hanson.org",
		Source: "oftbot!~oftbot@WoPRCentral.jonathan-hanson.org",
		User:   "~oftbot",
		Arguments: []string{
			"#cosmic-rift",
			"@triplepoint rolled: 2, 1, 2, 2, 6.",
		},
	})
}

func TestOftbotEventCanBeBuilt(t *testing.T) {
	GetSampleGameSuggestionEvent()
}

func TestMatchGameSuggestion(t *testing.T) {
	event := GetSampleGameSuggestionEvent()

	if !event.IsGameSuggestion() {
		t.Fatal("Was not able to match game suggestion text")
	}
}

func TestIsTimeToRollDetection(t *testing.T) {
	event := GetSampleTimeToRollEvent()
	jonBot := Bot{Name: "triplepoint"}
	zoBot := Bot{Name: "zeelot"}

	if !event.IsTimeToRoll(jonBot) {
		t.Fatal("Time to roll detection failed for jon bot")
	}
	if event.IsTimeToRoll(zoBot) {
		t.Fatal("Time to roll detection failed for zo bot")
	}
}

func TestIsTimeToKeepDetection(t *testing.T) {
	event := GetSampleTimeToKeepEvent()
	jonBot := Bot{Name: "triplepoint"}
	zoBot := Bot{Name: "zeelot"}

	if !event.IsTimeToKeep(jonBot) {
		t.Fatal("Time to keep detection failed for jon bot")
	}
	if event.IsTimeToKeep(zoBot) {
		t.Fatal("Time to keep detection failed for zo bot")
	}
}

func TestSecondRollDetection(t *testing.T) {
	event := GetSampleSecondRollEvent()
	jonBot := Bot{Name: "triplepoint"}
	zoBot := Bot{Name: "zeelot"}

	event.GetRollValues()

	if !event.IsTimeToKeep(jonBot) {
		t.Fatal("Time to keep detection failed for jon bot")
	}
	if event.IsTimeToKeep(zoBot) {
		t.Fatal("Time to keep detection failed for zo bot")
	}
}
