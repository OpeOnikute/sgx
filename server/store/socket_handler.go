package store

import (
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/websocket"
)

const (
	ErrorSignal       = "error"
	JoinStorySignal   = "story:join"
	CreateStorySignal = "story:create"
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
var Clients = make(map[string]map[*websocket.Conn]bool)

// Broadcast channel
var Broadcast = make(chan Message)

// Message defines the structure of messages expected from the client
type Message struct {
	Signal     string  `json:"signal"`
	InviteCode string  `json:"inviteCode,omitempty"`
	StoryID    string  `json:"storyID,omitempty"`
	Content    Content `json:"content"`
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

	if m.StoryID != "" && !bson.IsObjectIdHex(m.StoryID) {
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
		return
	}

	responseMessage := Message{
		Signal:     "story:created",
		InviteCode: story.InviteCode,
		Content: Content{
			StoryTitle:  story.Title,
			PlayerID:    story.PlayerOne.UID,
			PlayerEmail: story.PlayerOne.Email,
			PlayerName:  story.PlayerOne.Name,
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
		return
	}

	// mutate message signal to indicate what just happened.
	msg.Signal = "player:new"
	msg.Content.PlayerID = story.PlayerTwo.UID
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
	}
}

func broadcastMessage(msg Message) {
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
