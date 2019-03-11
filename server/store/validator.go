package store

import (
	"net/url"
	"regexp"

	"gopkg.in/mgo.v2/bson"
)

//StoryRequestBody - schema for expected params in story request body
type StoryRequestBody struct {
	Title       string `json:"title"`
	PlayerName  string `json:"playerName"`
	PlayerEmail string `json:"playerEmail"`
}

//JoinStoryRequestBody - schema for expected params in story request body
type JoinStoryRequestBody struct {
	Code        string `json:"code"`
	PlayerName  string `json:"playerName"`
	PlayerEmail string `json:"playerEmail"`
}

// AddParagraphRequestBody ...
type AddParagraphRequestBody struct {
	StoryID  string `json:"storyID"`
	PlayerID string `json:"playerID"`
	Content  string `json:"content"`
}

//Basic validation for now.
//eventually change this to use go-validator or a better alternative.
func (s *StoryRequestBody) validate() url.Values {
	errs := url.Values{}

	if s.Title == "" {
		errs.Add("title", "The title is required.")
	}

	if s.PlayerName == "" {
		errs.Add("playerName", "Please enter a valid player name.")
	}

	match, _ := regexp.MatchString(EmailRegex, s.PlayerEmail)

	if s.PlayerEmail == "" || !match {
		errs.Add("PlayerEmail", "Please enter a valid player email.")
	}

	return errs
}

func (s *JoinStoryRequestBody) validate() url.Values {
	errs := url.Values{}

	if s.Code == "" {
		errs.Add("code", "Please enter a valid code.")
	}

	if s.PlayerName == "" {
		errs.Add("playerName", "Please enter a valid player name.")
	}

	match, _ := regexp.MatchString(EmailRegex, s.PlayerEmail)

	if s.PlayerEmail == "" || !match {
		errs.Add("PlayerEmail", "Please enter a valid player email.")
	}

	return errs
}

func (s *AddParagraphRequestBody) validate() url.Values {
	errs := url.Values{}

	if s.StoryID == "" || !bson.IsObjectIdHex(s.StoryID) {
		errs.Add("story", "Please enter a valid story.")
	}

	if s.PlayerID == "" {
		errs.Add("playerID", "Please enter a valid player ID.")
	}

	if s.Content == "" {
		errs.Add("content", "Please enter valid content.")
	}

	return errs
}
