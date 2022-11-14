package main

import (
	"testing"
)

//Ensures that unmarshalling works as expected
func TestUnmarshal(t *testing.T) {
	testData := `
username: testUsername
password: testPassword
hostname: testHostname
`
	config := unmarshallConfig([]byte(testData))
	if config.Hostname != "testHostname" {
		t.Errorf("Hostname not unmarshalled correctly")
	}
	if config.Password != "testPassword" {
		t.Error("Password not unmarshalled correctly")
	}
	if config.Username != "testUsername" {
		t.Error("Username not unmarshalled correctly")
	}
}
