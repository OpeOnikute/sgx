package store

import (
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

//EndStoryRequestBody - schema for expected params
type EndStoryRequestBody struct {
	StoryID  bson.ObjectId `json:"storyID"`
	PlayerID string        `json:"playerID"`
}

// AddParagraphRequestBody ...
type AddParagraphRequestBody struct {
	StoryID  bson.ObjectId `json:"storyID"`
	PlayerID string        `json:"playerID"`
	Content  string        `json:"content"`
}

// ValidationErrs ...
type ValidationErrs map[string]string

//Basic validation for now.
//eventually change this to use go-validator or a better alternative.
func (s *StoryRequestBody) validate() ValidationErrs {
	errs := ValidationErrs{}

	if s.Title == "" {
		errs["title"] = "The title is required."
	}

	if s.PlayerName == "" {
		errs["playerName"] = "Please enter a valid player name."
	}

	match, _ := regexp.MatchString(EmailRegex, s.PlayerEmail)

	if s.PlayerEmail == "" || !match {
		errs["PlayerEmail"] = "Please enter a valid player email."
	}

	return errs
}

func (s *JoinStoryRequestBody) validate() ValidationErrs {
	errs := ValidationErrs{}

	if s.Code == "" {
		errs["code"] = "Please enter a valid code."
	}

	if s.PlayerName == "" {
		errs["playerName"] = "Please enter a valid player name."
	}

	match, _ := regexp.MatchString(EmailRegex, s.PlayerEmail)

	if s.PlayerEmail != "" && !match {
		errs["playerEmail"] = "Please enter a valid player email."
	}

	return errs
}

func (s *AddParagraphRequestBody) validate() ValidationErrs {
	errs := ValidationErrs{}

	if s.StoryID == "" {
		errs["story"] = "Please enter a valid story."
	}

	if s.PlayerID == "" {
		errs["playerID"] = "Please enter a valid player ID."
	}

	if s.Content == "" {
		errs["content"] = "Please enter valid content."
	}

	return errs
}
