package cosplaystart

import (
	"bytes"
	"strings"
	"testing"
)

func intializeTest() {
	InitializeTemplates("../templates/html/")
}

// Basic Implementation Tests
func TestCreatePost(t *testing.T) {
	basicImpl := basicImpl{make(map[string]Post)}
	basicImpl.CreatePost("TEST", "TEST", "TEST")

	if len(basicImpl.posts) != 1 {
		t.Fail()
	}
}

func TestDeletePost(t *testing.T) {
	basicImpl := basicImpl{make(map[string]Post)}
	basicImpl.posts["TEST"] = Post{} // Add a fake post
	basicImpl.DeletePost("TEST")
	if len(basicImpl.posts) != 0 {
		t.Fail()
	}

	//Should not delete post if not here
	basicImpl.posts["TEST"] = Post{} // Add a fake post
	basicImpl.DeletePost("DOESNT EXIST")
	if len(basicImpl.posts) != 1 {
		t.Fail()
	}
}

func TestGetSeries(t *testing.T) {
	basicImpl := basicImpl{make(map[string]Post)}
	basicImpl.posts["TEST1"] = Post{series: "NARUTO"}
	basicImpl.posts["TEST2"] = Post{series: "ONE PIECE"}
	basicImpl.posts["TEST3"] = Post{series: "DEMON SLAYER"}
	basicImpl.posts["TEST4"] = Post{series: "Dragon Ball Z"}
	basicImpl.posts["TEST5"] = Post{series: "NARUTO"}

	m := basicImpl.GetSeries()
	if len(m) != 4 {
		t.Fail()
	}

}

func TestGetCharacters(t *testing.T) {
	basicImpl := basicImpl{make(map[string]Post)}
	basicImpl.posts["TEST1"] = Post{series: "NARUTO", character: "Sauske"}
	basicImpl.posts["TEST2"] = Post{series: "ONE PIECE", character: "Luffy"}
	basicImpl.posts["TEST3"] = Post{series: "DEMON SLAYER", character: "Nezuko"}
	basicImpl.posts["TEST4"] = Post{series: "Dragon Ball Z", character: "Gohan"}
	basicImpl.posts["TEST5"] = Post{series: "NARUTO", character: "Sakura"}

	if len(basicImpl.GetCharacters("NARUTO")) != 2 {
		t.Fail()
	}
	if len(basicImpl.GetCharacters("DEMON SLAYER")) != 1 {
		t.Fail()
	}
	if len(basicImpl.GetCharacters("NOT HERE")) != 0 {
		t.Fail()
	}
}

func TestGenerateShowsListTemplate(t *testing.T) {
	intializeTest()
	basicImpl := basicImpl{make(map[string]Post)}

	basicImpl.posts["TEST1"] = Post{series: "NARUTO", character: "Sauske"}
	basicImpl.posts["TEST3"] = Post{series: "NARUTO", character: "Sauske"}
	basicImpl.posts["TEST2"] = Post{series: "ONE PIECE", character: "Luffy"}

	b := new(bytes.Buffer)
	GenerateShowsView(b, basicImpl)
	if strings.Contains(b.String(), "NARUTO") != true {
		t.Fail()
	}
	if strings.Contains(b.String(), "ONE PIECE") != true {
		t.Fail()
	}
}

func TestGenerateCharactersListTemplate(t *testing.T) {
	intializeTest()
	basicImpl := basicImpl{make(map[string]Post)}

	basicImpl.posts["TEST1"] = Post{series: "NARUTO", character: "Sauske"}
	basicImpl.posts["TEST2"] = Post{series: "NARUTO", character: "Sauske"}
	basicImpl.posts["TEST3"] = Post{series: "NARUTO", character: "Sakura"}
	basicImpl.posts["TEST4"] = Post{series: "ONE PIECE", character: "Luffy"}
	basicImpl.posts["TEST5"] = Post{series: "DEMON SLAYER", character: "Nezuko"}
	basicImpl.posts["TEST6"] = Post{series: "Dragon Ball Z", character: "Gohan"}

	b := new(bytes.Buffer)
	GenerateCharactersView("NARUTO", b, basicImpl)
	if strings.Contains(b.String(), "Sauske") != true {
		t.Fail()
	}
	if strings.Contains(b.String(), "Sakura") != true {
		t.Fail()
	}
	if strings.Contains(b.String(), "Gohan") == true {
		t.Fail()
	}
}

func TestGeneratePostsForCharacterAndSeriesListTemplate(t *testing.T) {
	intializeTest()
	basicImpl := basicImpl{make(map[string]Post)}

	basicImpl.posts["TEST1"] = Post{series: "NARUTO", character: "Sauske", Author: "Chris", Title: "Budget Sauske"}
	basicImpl.posts["TEST2"] = Post{series: "NARUTO", character: "Sauske", Author: "Alex", Title: "Edge Personified"}
	basicImpl.posts["TEST3"] = Post{series: "NARUTO", character: "Sauske", Author: "N'dia", Title: "GenderBender Sauske Love!"}
	basicImpl.posts["TEST4"] = Post{series: "NARUTO", character: "Sakura", Author: "N'dia"}

	b := new(bytes.Buffer)

	GeneratePostsView("Sauske", "NARUTO", b, basicImpl)
	if strings.Contains(b.String(), "Chris") != true {
		t.Fail()
	}
	if strings.Contains(b.String(), "Alex") != true {
		t.Fail()
	}
	if strings.Contains(b.String(), "N&#39;dia") != true {
		t.Fail()
	}
	if strings.Contains(b.String(), "Budget Sauske") != true {
		t.Fail()
	}
	if strings.Contains(b.String(), "Sakura") {
		t.Fail()
	}
}
