package saltybet

import (
	"log"
	"testing"

	"github.com/go-rod/rod"
)

func TestLogin(t *testing.T) {
	testPage := rod.New().MustConnect().MustPage()
	sb := SaltyBet{page: testPage, username: "wing.edkraby@gmail.com", password: "password"}
	err := sb.Login()
	if err != nil {
		t.Fail()
	}
	//The page should now be on the salty bet homepage, logged in
	if !sb.isLoggedIn() { //There is prob a quicker way to test this...
		t.Fail()
	}
}

func TestCheckBalance(t *testing.T) {

	testPage := rod.New().MustConnect().MustPage()
	sb := SaltyBet{page: testPage, username: "wing.edkraby@gmail.com", password: "password"}
	sb.Login()

	//Given we are logged in. We should be able to see balance

	bal, err := sb.GetBalance()

	if err != nil {
		log.Panic(err)
		t.Fail()
	}

	if bal < 0 {
		t.Fail()
	}

}

func TestCheckBalanceFailsWhenLoggedOut(t *testing.T) {

	testPage := rod.New().MustConnect().MustPage()
	sb := SaltyBet{page: testPage, username: "wing.edkraby@gmail.com", password: "password"}

	_, err := sb.GetBalance()

	if err == nil {
		t.Fail()
	}

}

func TestPlaceBet(t *testing.T) {

}
