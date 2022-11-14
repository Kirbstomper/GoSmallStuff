package main

/***
THe idea is to have someting keeping track of all messages
sent to a channel and backing them up

THe backup will be in a local database which is periodically
backed up to an external instance using lightstream
https://mtlynch.io/litestream/

seems real cool, so lets try this shit out!
***/

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {

	db = startDatabase()

	//First lets get connecting to the discord client out of the way
	dg, err := discordgo.New("Bot " + "BOT_API_KEY")
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	dg.AddHandler(listenForWords)
	dg.Identify.Intents = discordgo.IntentsGuildMessages
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}

// Open the db and create tables if they do not exist already
func startDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "messages.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS messages(id VARCHAR(32), message TEXT, timestamp TEXT)`)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// Handle listening and writing
func insertMessage(id string, msg string) error {
	_, err := db.Exec(`INSERT INTO messages(id, message, timestamp) values (?, ?, ?)`, id, msg, time.Now().Format(time.RFC3339))

	return err
}

//Function to record all messages with the keyword in it

func listenForWords(s *discordgo.Session, m *discordgo.MessageCreate) {

	//Ignore the bot, fuck the bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.Contains(m.Content, "fuck") {
		err := insertMessage(m.Author.ID, m.Content)
		if err != nil {
			log.Fatal("Error inserting message into database. What the fuck")
		}
	}
}
