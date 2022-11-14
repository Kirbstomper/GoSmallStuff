package cosplaystart

import (
	"log"

	"github.com/rs/xid"
)

type CosplayOperations interface {
	CreatePost(string, string, string)                      //Creates a post
	DeletePost(string)                                      //Deletes a post
	GetSeries() map[string]int                              //Gets all the series that posts exist for
	GetCharacters(string) map[string]int                    //Gets all characters for a series
	GetAllPostsForSeries(string) []Post                     //Gets all posts for a given series
	GetAllPostsForCharacterAndSeries(string, string) []Post //Gets all posts for a given character and series
	GetPost(string) Post                                    //Looksup a post given its id
}

type basicImpl struct {
	posts map[string]Post
}
type Post struct {
	series    string
	character string
	Author    string
	Title     string
	Content   string
	Id        string
}

func NewBasicOperationsImpl() CosplayOperations {
	basicImpl := basicImpl{make(map[string]Post)}
	basicImpl.posts["TEST1"] = Post{Id: "TEST1", series: "NARUTO", character: "Sauske", Author: "Chris", Title: "Budget Sauske"}
	basicImpl.posts["TEST2"] = Post{Id: "TEST2", series: "NARUTO", character: "Sauske", Author: "Alex", Title: "Edge Personified"}
	basicImpl.posts["TEST3"] = Post{Id: "TEST3", series: "NARUTO", character: "Sauske", Author: "N'dia", Title: "GenderBender Sauske Love!"}
	basicImpl.posts["TEST4"] = Post{Id: "TEST4", series: "NARUTO", character: "Sakura", Author: "N'dia", Title: "GenderBender Sauske Love! FIXED", Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."}
	return basicImpl
}

//Creates a post and adds it to the slice of Posts
func (b basicImpl) CreatePost(series string, character string, content string) {
	guid := xid.New()
	id := guid.String()
	post := Post{series: series, character: character, Content: content, Id: id}
	log.Print("Creating Post: ", post)
	b.posts[id] = post
}

//Deletes a post given its unique ID, I should prob use a fucking map here...
func (b basicImpl) DeletePost(postid string) {
	log.Print("Removing Post: ", b.posts[postid])
	delete(b.posts, postid)
}

//Returns all series currently present in the slice of series
func (b basicImpl) GetSeries() map[string]int {
	m := make(map[string]int)
	for _, post := range b.posts {
		m[post.series] = m[post.series] + 1
	}
	return m
}

//Returns all characters for a given series
func (b basicImpl) GetCharacters(series string) map[string]int {
	m := make(map[string]int)
	for _, post := range b.posts {
		if series == post.series {
			m[post.character] = m[post.character] + 1
		}
	}
	return m
}

func (b basicImpl) GetAllPostsForSeries(s string) []Post {
	posts := make([]Post, 0)
	for _, p := range b.posts {
		if s == p.series {
			posts = append(posts, p)
		}
	}

	return posts
}
func (b basicImpl) GetAllPostsForCharacterAndSeries(c string, s string) []Post {
	posts := make([]Post, 0)
	for _, p := range b.posts {
		if s == p.series && c == p.character {
			posts = append(posts, p)
		}
	}

	return posts
}

func (b basicImpl) GetPost(p string) Post {
	return b.posts[p]
}
