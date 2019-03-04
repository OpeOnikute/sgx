package store

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sgx/server/db"
	"time"

	"github.com/kjk/betterguid"
	"gopkg.in/mgo.v2/bson"
)

//Handler ...
type Handler struct{}

// DBNAME the name of the DB instance
const DBNAME = "sgx"

// AddStory adds a Story in the DB
func (r Handler) AddStory(params StoryRequestBody) bool {

	var playerOne Player
	collection := "story"

	playerOne.UID = betterguid.New()
	playerOne.Name = params.PlayerName
	playerOne.Email = params.PlayerEmail

	var story = Story{
		ID:        bson.NewObjectId(),
		Title:     params.Title,
		PlayerOne: playerOne,
		Status:    "open",
		Created:   time.Now(),
	}

	if err := db.Session.C(collection).Insert(story); err != nil {
		log.Println(err)
		return false
	}

	return true
}

func (r Handler) GetStoryByID(ID string) (Story, error) {

	var story Story

	c := db.Session.C("story")

	if err := c.FindId(bson.ObjectIdHex(ID)).One(&story); err != nil {
		log.Println(err)
		return story, err
	}

	return story, nil
}

func (r Handler) ParseStory(ID string) (string, error) {

	var story Story

	var parsed string

	c := db.Session.C("story")

	err := c.FindId(bson.ObjectIdHex(ID)).One(&story)

	if err != nil {
		log.Println(err)
		return "", err
	}

	var content = story.Content

	//TODO: Add capitalization and pretty formating. Maybe using a package.
	for i := 0; i < len(content); i++ {
		parsed += (content[i].Text + "\n\n")
	}

	return parsed, nil
}

// JoinStory adds player two
func (r Handler) JoinStory(params JoinStoryRequestBody) bool {

	var playerTwo = Player{}

	playerTwo.UID = betterguid.New()
	playerTwo.Name = params.PlayerName
	playerTwo.Email = params.PlayerEmail

	var story Story

	c := db.Session.C("story")

	if err := c.FindId(bson.ObjectIdHex(params.StoryID)).One(&story); err != nil {
		log.Println(err)
		return false
	}

	story.PlayerTwo = playerTwo
	story.Status = "active"

	err := c.UpdateId(story.ID, story)

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// AddParagraph ...
func (r Handler) AddParagraph(params AddParagraphRequestBody) bool {

	var story Story
	c := db.Session.C("story")

	if err := c.FindId(bson.ObjectIdHex(params.StoryID)).One(&story); err != nil {
		log.Println(err)
		return false
	}

	var content = StoryFormat{
		UserID: params.PlayerID,
		Text:   params.Content,
		Time:   time.Now(),
	}

	story.Content = append(story.Content, content)

	err := c.UpdateId(story.ID, story)

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// GetStories returns the list of Stories
func (r Handler) GetStories() Stories {
	collection := "story"
	c := db.Session.C(collection)
	results := Stories{}

	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}

// SendSuccess ...
func (r Handler) SendSuccess(w http.ResponseWriter, data Response) bool {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
	return true
}

// SendError ...
func (r Handler) SendError(w http.ResponseWriter, code int, response Response, err error) {
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch code {
	case 400:
		w.WriteHeader(http.StatusBadRequest)
	case 500:
		w.WriteHeader(http.StatusInternalServerError)
	default:
	}

	json.NewEncoder(w).Encode(response)
}
