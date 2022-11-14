package cosplaystart

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
)

//This exists to contain actually producing pages based on certain endpoints

//Generates the view for when you want to list all shows with cosplay availible

var BasefilePath = ""

type CharactersView struct {
	Series string
	Data   map[string]int
}

type PostsView struct {
	Series    string
	Character string
	Data      []Post
}

var tempEngine template.Template

func readTemplates(dir string) *template.Template {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal("Error occured reading directory", err)
	}
	filenames := make([]string, 0)
	for _, u := range files {
		filenames = append(filenames, dir+u.Name())
	}
	tempeng, err := template.ParseFiles(filenames...)
	if err != nil {
		log.Fatal("Error parsing files", err)
	}
	return tempeng
}

func InitializeTemplates(dir string) {
	tempEngine = *readTemplates(dir)
}
func GenerateShowsView(out io.Writer, op CosplayOperations) {
	showsList := op.GetSeries()
	err := tempEngine.ExecuteTemplate(out, "showlist.html", showsList)
	if err != nil {
		log.Fatal("Error occured executing template! GenerateShowsView", err)
	}
}

//Generates the view for list of characters with cosplay from a given
func GenerateCharactersView(series string, out io.Writer, op CosplayOperations) {
	err := tempEngine.ExecuteTemplate(out, "characterlist.html", CharactersView{Series: series, Data: op.GetCharacters(series)})
	if err != nil {
		log.Fatal("Error occured executing template! GenerateCharactersView", err)
	}
}

//Generates the view for listing posts for a character, for a given show
//Rename this shit, that's complaining
func GeneratePostsView(character string, series string, out io.Writer, op CosplayOperations) {
	posts := op.GetAllPostsForCharacterAndSeries(character, series)
	err := tempEngine.ExecuteTemplate(out, "postslist.html", PostsView{Series: series, Character: character, Data: posts})
	if err != nil {
		log.Fatal("Error occured executing template! GenerateCharactersView", err)
	}
}

func GeneratePostView(p string, out io.Writer, op CosplayOperations) {
	err := tempEngine.ExecuteTemplate(out, "viewpost.html", op.GetPost(p))
	if err != nil {
		log.Fatal("Error occured executing template! GeneratePostView", err)
	}
}
