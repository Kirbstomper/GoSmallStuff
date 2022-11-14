package saltybet

import (
	"errors"
	"strconv"
	"strings"

	"github.com/go-rod/rod"
)

type SaltyBet struct {
	page               *rod.Page // The page used by the betting engine
	username, password string    // The username and password used to login
}

// Logs into saltybet and returns the page to the homepage,
// returns an error if encountered
func (s SaltyBet) Login() error {

	wait := s.page.MustWaitNavigation()
	//s.page.MustEmulate(devices.IPadMini)
	s.page.MustNavigate("https://www.saltybet.com/authenticate?signin=1")
	wait()

	s.page.MustElement(`[id="email"]`).MustInput(s.username)
	s.page.MustElement(`[id="pword"]`).MustInput(s.password)
	s.page.MustElement(`[type="submit"]`).MustClick()
	s.page.MustWaitLoad()
	return nil //Should everything be fine. Return nil
}

// Returns the balance for the account currently logged in
func (s SaltyBet) GetBalance() (int, error) {

	if !s.isLoggedIn() {
		return -1, errors.New("User not Logged in!")
	}
	balEl, err := s.page.Element(`[id="balance"]`)

	if err != nil {
		return -1, err
	}
	balString := balEl.MustEval(`() => this.innerText`).String()

	balString = strings.ReplaceAll(balString, ",", "")

	bal, err := strconv.Atoi(balString)

	return bal, err
}
func (s SaltyBet) PlaceBet(int) (error){

	bal = s.GetBalance()

	return nil
}
//Checks if the user is logged in
func (s SaltyBet) isLoggedIn() bool {

	html, _ := s.page.HTML()
	return strings.Contains(html, `<a href="/logout">Logout</a>`)
}
