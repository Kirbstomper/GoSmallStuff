package main

import (
	"log"
	"net/http"

	cosplaystart "github.com/Kirbstomper/CosplayStart/cosplaystartBackend"
	"github.com/gorilla/mux"
)

var operations = cosplaystart.NewBasicOperationsImpl()

func listSeries(rw http.ResponseWriter, r *http.Request) {
	cosplaystart.GenerateShowsView(rw, operations)
}

func listCharactersForSeries(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cosplaystart.GenerateCharactersView(vars["sname"], rw, operations)
}

func listPostsForCharacterInSeries(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cosplaystart.GeneratePostsView(vars["cname"], vars["sname"], rw, operations)
}

func listAllPosts(rw http.ResponseWriter, r *http.Request) {}
func displayPost(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cosplaystart.GeneratePostView(vars["postid"], rw, operations)
}

func main() {
	log.Println("Starting Application!")

	log.Println("Initializing Templates")
	cosplaystart.InitializeTemplates("templates/html/")

	g := mux.NewRouter()

	g.HandleFunc("/series/{sname}/", listCharactersForSeries)
	g.HandleFunc("/series/{sname}/{cname}", listPostsForCharacterInSeries)
	g.HandleFunc("/series/", listSeries)
	g.HandleFunc("/posts/", listAllPosts)
	g.HandleFunc("/posts/{postid}", displayPost)

	http.Handle("/", g)

	log.Println("Starting Server!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
