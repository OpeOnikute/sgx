package store

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"
)

//Controller ...
type Controller struct {
	Handler Handler
}

// Response ...
type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

var upgrader = websocket.Upgrader{}

// AddStory POST /
func (c *Controller) AddStory(w http.ResponseWriter, r *http.Request) {
	var story StoryRequestBody
	var response Response
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&story); err != nil {
		response.Message = "We could not parse your data. Please try again."
		c.Handler.SendError(w, 500, response, err)
		return
	}

	if validErrs := story.validate(); len(validErrs) > 0 {
		response.Data = map[string]interface{}{"validationError": validErrs}
		response.Message = "The data you entered is invalid."
		c.Handler.SendError(w, 400, response, nil)
		return
	}

	//from here, add better data validation
	createdStory, err := c.Handler.AddStory(story) // adds the product to the DB
	if err != nil {
		response.Message = "There was a problem creating the story. Please try again"
		c.Handler.SendError(w, 500, response, nil)
		return
	}

	response.Data = createdStory

	c.Handler.SendSuccess(w, response)
}

// AddParagraph ..
func (c *Controller) AddParagraph(w http.ResponseWriter, r *http.Request) {
	var body AddParagraphRequestBody
	var response Response
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response.Message = "We could not parse your data. Please try again."
		c.Handler.SendError(w, 500, response, err)
		return
	}

	if validErrs := body.validate(); len(validErrs) > 0 {
		response.Data = map[string]interface{}{"validationError": validErrs}
		response.Message = "The data you entered is invalid."
		c.Handler.SendError(w, 400, response, nil)
		return
	}

	// check for story
	story, err := c.Handler.GetStoryByID(body.StoryID)

	if err != nil {
		response.Message = err.Error()
		c.Handler.SendError(w, 400, response, err)
		return
	}

	if story.Status != "active" {
		response.Message = "Sorry, this story is not active."
		c.Handler.SendError(w, 400, response, nil)
		return
	}

	if story.PlayerOne.UID != body.PlayerID && story.PlayerTwo.UID != body.PlayerID {
		response.Message = "You don't seem to be a valid player."
		c.Handler.SendError(w, 400, response, nil)
		return
	}

	story, err = c.Handler.AddParagraph(body)
	if err != nil {
		response.Message = "There was a problem adding the paragraph. Please try again."
		c.Handler.SendError(w, 500, response, err)
		return
	}

	response.Data = story
	c.Handler.SendSuccess(w, response)

}

// GetStoryByField POST /
func (c *Controller) GetStoryByField(w http.ResponseWriter, r *http.Request) {
	var response Response
	var story Story

	queryParams := r.URL.Query()

	if len(queryParams["f"]) < 1 || len(queryParams["v"]) < 1 {
		response.Message = "The parameters you entered are invalid."
		c.Handler.SendError(w, 442, response, nil)
		return
	}

	field := queryParams["f"][0]
	value := queryParams["v"][0]

	//confirm the story exists first
	story, err := c.Handler.GetStoryByField(field, value)

	if err != nil {
		response.Message = err.Error()
		c.Handler.SendError(w, 400, response, err)
		return
	}

	response.Data = story
	c.Handler.SendSuccess(w, response)
}

// JoinStory POST /
func (c *Controller) JoinStory(w http.ResponseWriter, r *http.Request) {
	var validRequest JoinStoryRequestBody
	var response Response
	var story Story
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&validRequest); err != nil {
		response.Message = "We could not parse your data. Please try again."
		c.Handler.SendError(w, 500, response, err)
		return
	}

	if validErrs := validRequest.validate(); len(validErrs) > 0 {
		response.Data = map[string]interface{}{"validationError": validErrs}
		response.Message = "The data you entered is invalid."
		c.Handler.SendError(w, 400, response, nil)
		return
	}

	//confirm the story exists first
	story, err := c.Handler.GetStoryByField("invitecode", validRequest.Code)

	if err != nil {
		response.Message = err.Error()
		c.Handler.SendError(w, 400, response, err)
		return
	}

	if story.Status != "open" {
		response.Message = "Sorry, this story is not available."
		c.Handler.SendError(w, 400, response, nil)
		return
	}

	//from here, add better data validation
	story, err = c.Handler.JoinStory(validRequest)
	if err != nil {
		response.Message = "There was a problem joining the story. Please try again"
		c.Handler.SendError(w, 500, response, err)
		return
	}

	response.Data = story
	c.Handler.SendSuccess(w, response)
}

// GetAllStories ...
func (c *Controller) GetAllStories(w http.ResponseWriter, r *http.Request) {
	products := c.Handler.GetStories() // list of all products

	var response = Response{"", products}

	c.Handler.SendSuccess(w, response)
}

// GetParsedStory GET /
func (c *Controller) GetParsedStory(w http.ResponseWriter, r *http.Request) {
	var response Response
	var storyID string
	defer r.Body.Close()

	queryParams := r.URL.Query()

	if len(queryParams["s"]) < 1 || queryParams["s"][0] == "" {
		response.Message = "Please enter a valid story."
		c.Handler.SendError(w, 400, response, nil)
		return
	}

	storyID = queryParams["s"][0]

	if !bson.IsObjectIdHex(storyID) {
		response.Message = "Please enter a valid story ID."
		c.Handler.SendError(w, 400, response, nil)
		return
	}

	parsed, err := c.Handler.ParseStory(storyID)

	if err != nil {
		response.Message = err.Error()
		c.Handler.SendError(w, 400, response, err)
		return
	}

	response.Data = parsed

	c.Handler.SendSuccess(w, response)
}

func (c *Controller) handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			errMsg := fmt.Sprintf("error: %v", err)
			log.Println(errMsg)
			WriteToClient(ws, errMsg, nil)
			break
		}

		if validErrs := msg.validate(); len(validErrs) > 0 {
			//TODO: Write the error to the client here.
			errMsg := fmt.Sprintf("Invalid data entered: %v", validErrs)
			log.Println(errMsg)
			WriteToClient(ws, errMsg, nil)
			break
		}
		// Send the newly received message to the broadcast channel
		msg.Client = ws
		Broadcast <- msg
	}
}

// WriteToClient sends a json ws message to a connected client
func WriteToClient(client *websocket.Conn, msg string, data interface{}) {
	errMsg := ErrorMessage{
		ErrorSignal,
		msg,
		data,
	}
	err := client.WriteJSON(errMsg)
	if err != nil {
		log.Printf("error: %v", err)
		client.Close()
	}
}
