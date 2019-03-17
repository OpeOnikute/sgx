package store

import (
	"testing"
)

func TestStoryRequestBodyValidator(t *testing.T) {

	request := StoryRequestBody{
		Title:       "ope",
		PlayerName:  "onikute",
		PlayerEmail: "kjfbd@yahoo.com",
	}

	errs := request.validate()

	if len(errs) > 0 {
		t.Errorf("Story request: %v", errs)
	}
}

func TestJoinStoryRequestBodyValidator(t *testing.T) {

	request := JoinStoryRequestBody{
		Code:        "ope",
		PlayerName:  "onikute",
		PlayerEmail: "kj-fbd@yahoo.com",
	}

	errs := request.validate()

	if len(errs) > 0 {
		t.Errorf("Join story request: %v", errs)
	}
}

func TestAddParagraphRequestBodyValidator(t *testing.T) {

	request := AddParagraphRequestBody{
		StoryID:  "5bd6fff6cf23a172915e9e48",
		PlayerID: "sdsd",
		Content:  "Sdsdsd",
	}

	errs := request.validate()

	if len(errs) > 0 {
		t.Errorf("Add paragraph error: %v", errs)
	}
}
