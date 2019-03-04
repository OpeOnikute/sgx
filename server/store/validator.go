package store

import (
	"net/url"
)

//StoryRequestBody - schema for expected params in story request body
type StoryRequestBody struct {
	Title       string `json:"title"`
	PlayerName  string `json:"playerName"`
	PlayerEmail string `json:"playerEmail"`
}

//JoinStoryRequestBody - schema for expected params in story request body
type JoinStoryRequestBody struct {
	StoryID     string `json:"storyID"`
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

	if s.PlayerEmail == "" {
		errs.Add("PlayerEmail", "Please enter a valid player email.")
	}

	return errs
}

func (s *JoinStoryRequestBody) validate() url.Values {
	errs := url.Values{}

	if s.StoryID == "" {
		errs.Add("story", "Please enter a valid story.")
	}

	if s.PlayerName == "" {
		errs.Add("playerName", "Please enter a valid player name.")
	}

	if s.PlayerEmail == "" {
		errs.Add("PlayerEmail", "Please enter a valid player email.")
	}

	return errs
}

func (s *AddParagraphRequestBody) validate() url.Values {
	errs := url.Values{}

	if s.StoryID == "" {
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
