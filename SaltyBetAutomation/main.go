package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/devices"
)

func main() {
	page := rod.New().MustConnect().MustPage()
	wait := page.MustWaitNavigation()
	page.MustEmulate(devices.IPadMini)
	page.MustNavigate("https://www.saltybet.com/authenticate?signin=1")
	wait()

	//Login flow
	page.MustElement(`[id="email"]`).MustInput("usernameconst")
	page.MustElement(`[id="pword"]`).MustInput("password")
	page.MustScreenshot("input.png")
	page.MustElement(`[type="submit"]`).MustClick()
	page.MustWaitLoad()

	//Bet 10% on player one

	el, err := page.Element(`[id="player1"]`)

	if err != nil {
		println(err)
	}

	var canBet = true
	for 1 == 1 {
		s, err := el.Attribute("disabled")
		if err != nil {
			println(err)
		}
		if s == nil && canBet {
			page.MustScreenshot("betplaced.png")
			//println(page.HTML())
			page.MustElement(`[id="interval1"]`).MustClick()
			page.MustElement(`[id="player1"]`).MustClick()
			canBet = false
		}
		if s != nil {
			canBet = true
		}
	}

	//
}
