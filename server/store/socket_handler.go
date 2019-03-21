package store

import (
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/websocket"
)

const (
	ErrorSignal          = "error"
	JoinStorySignal      = "story:join"
	CreateStorySignal    = "story:create"
	CreatedStorySignal   = "story:created"
	AddParagraphSignal   = "story:paragraph"
	AddedParagraphSignal = "story:new-paragraph"
	EndStorySignal       = "story:end"
)

var handler = &Handler{}

//SocketHandler implements the strategy pattern
type SocketHandler struct {
}

// Clients stores connected clients. Nested Maps story ID to the connections
// structure:
// {
//	storyID: {
//		*websocket.Conn: bool
//	}
//}
var Clients = make(map[bson.ObjectId]map[*websocket.Conn]bool)

// Broadcast channel
var Broadcast = make(chan Message)

// Message defines the structure of messages expected from the client
type Message struct {
	Client     *websocket.Conn `json:"client,omitempty"`
	Signal     string          `json:"signal"`
	InviteCode string          `json:"inviteCode,omitempty"`
	StoryID    bson.ObjectId   `json:"storyID,omitempty"`
	Content    Content         `json:"content"`
}

// ErrorMessage defines the structure of ws error messages sent to clients
type ErrorMessage struct {
	Signal  string      `json:"signal"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

// Content defines the structure of the json message content
type Content struct {
	PlayerID    string    `json:"playerID,omitempty"`
	PlayerEmail string    `json:"playerEmail"`
	PlayerName  string    `json:"playerName"`
	Message     string    `json:"message,omitempty"`
	StoryOwner  bool      `json:"storyOwner,omitempty"`
	StoryTitle  string    `json:"storyTitle,omitempty"`
	Time        time.Time `json:"time"`
}

func (m *Message) validate() ValidationErrs {
	errs := ValidationErrs{}

	if m.Signal == "" {
		errs["signal"] = "Please enter a valid signal."
	}

	if m.Signal == JoinStorySignal && m.InviteCode == "" {
		errs["inviteCode"] = "Please enter a valid invite code."
	}

	if m.StoryID == "" {
		errs["storyID"] = "Please enter a valid story."
	}

	content := Content{}
	if m.Content == content {
		errs["content"] = "Please enter valid content."
	}

	return errs
}

func createStory(msg Message) {

	request := StoryRequestBody{
		msg.Content.StoryTitle,
		msg.Content.PlayerName,
		msg.Content.PlayerEmail,
	}

	story, err := handler.AddStory(request)

	if err != nil {
		log.Println(err)
		WriteToClient(msg.Client, err.Error(), nil)
		return
	}

	//register the client for the story here.
	//this is especially important here because the story id can only be gotten for the first time here.
	registerClient(story.ID, msg.Client)

	responseMessage := Message{
		Signal:     CreatedStorySignal,
		InviteCode: story.InviteCode,
		StoryID:    story.ID,
		Content: Content{
			StoryTitle:  story.Title,
			PlayerID:    story.PlayerOne.UID,
			PlayerEmail: story.PlayerOne.Email,
			PlayerName:  story.PlayerOne.Name,
			Time:        msg.Content.Time,
		},
	}

	broadcastMessage(responseMessage)
}

// joinStory method for player two.
func joinStory(msg Message) {

	if msg.InviteCode == "" {
		log.Println("No invite code passed.")
		return
	}

	request := JoinStoryRequestBody{
		msg.InviteCode,
		msg.Content.PlayerName,
		msg.Content.PlayerEmail,
	}

	story, err := handler.JoinStory(request)

	if err != nil {
		log.Println(err)
		WriteToClient(msg.Client, err.Error(), nil)
		return
	}

	//register the client for the story here.
	registerClient(story.ID, msg.Client)

	// mutate message signal to indicate what just happened.
	msg.Signal = "player:new"
	msg.StoryID = story.ID
	msg.Content.PlayerID = story.PlayerTwo.UID
	broadcastMessage(msg)
}

func addParagraph(msg Message) {

	request := AddParagraphRequestBody{
		msg.StoryID,
		msg.Content.PlayerID,
		msg.Content.Message,
	}

	story, err := handler.AddParagraph(request)

	if err != nil {
		WriteToClient(msg.Client, err.Error(), nil)
		log.Println(err)
		return
	}

	msg.Signal = AddedParagraphSignal
	msg.StoryID = story.ID
	broadcastMessage(msg)
}

func endStory(msg Message) {

	request := EndStoryRequestBody{
		msg.StoryID,
		msg.Content.PlayerID,
	}

	story, err := handler.EndStory(request)

	if err != nil {
		log.Println(err)
		WriteToClient(msg.Client, err.Error(), nil)
		return
	}

	msg.Signal = EndStorySignal
	msg.StoryID = story.ID
	broadcastMessage(msg)
}

func handleMessage(msg Message) {
	switch msg.Signal {
	case JoinStorySignal:
		joinStory(msg)
		break
	case CreateStorySignal:
		createStory(msg)
		break
	case AddParagraphSignal:
		addParagraph(msg)
		break
	case EndStorySignal:
		endStory(msg)
		break
	}
}

func broadcastMessage(msg Message) {
	if msg.StoryID == "" {
		log.Println("No story ID passed to broadcast to.")
		return
	}
	chatClients := Clients[msg.StoryID]
	for client := range chatClients {
		err := client.WriteJSON(msg)
		if err != nil {
			log.Printf("error: %v", err)
			client.Close()
			delete(chatClients, client)
		}
	}
}

//Remains to be seen if this need to be done for every message handled like such.
func registerClient(storyID bson.ObjectId, client *websocket.Conn) {
	//initialize the new chat map if it doesn't exist.
	if _, ok := Clients[storyID]; !ok {
		Clients[storyID] = make(map[*websocket.Conn]bool)
	}
	Clients[storyID][client] = true
}

// HandleSocketMessages is a Go routine that listens infinitely on the broadcast channel, maps the story and
// performs any actions/transformations needed before broadcasting out to all connected clients for that
// story.
func HandleSocketMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-Broadcast
		handleMessage(msg)
	}
}
