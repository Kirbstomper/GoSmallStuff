package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type config struct {
	Username string
	Password string
	Hostname string
}

var publicip string

func main() {
	//Read configuration file
	configfile, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	config := unmarshallConfig(configfile)
	fmt.Println(config)

	for true == true {
		//Get current IP address
		ipReq, err := http.Get("https://domains.google.com/checkip")
		if err != nil {
			log.Fatalf("Error getting IP: %v", err)
		}
		ipResp, _ := io.ReadAll(ipReq.Body)
		newIp := string(ipResp)
		fmt.Printf(newIp)

		//If the IP has changed, we should update it!
		if publicip != newIp {
			publicip = newIp
			query := "https://" + config.Username + ":" + config.Password + "@domains.google.com/nic/update?hostname=" + config.Hostname + "&myip=" + publicip
			resp, _ := http.Post(query, "", nil)
			resps, _ := io.ReadAll(resp.Body)
			fmt.Print(string(resps))
		}
		//wait 10 minutes before doing again
		time.Sleep(10 * time.Minute)
	}
}

// Umarshalls the configuration byte slice passed to a config struct
// and returns it
func unmarshallConfig(yml []byte) config {
	c := config{}
	err := yaml.Unmarshal(yml, &c)
	if err != nil {
		log.Fatalf("error: %v , \n there was a problem parsing your configuration file", err)
	}
	return c
}
