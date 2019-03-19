package store

import (
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/websocket"
)

//SocketHandler ...
type SocketHandler struct{}

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
	Signal  string  `json:"signal"`
	StoryID string  `json:"storyID"`
	Content Content `json:"content"`
}

// Content defines the structure of the json message content
type Content struct {
	Email      string    `json:"email"`
	Username   string    `json:"username"`
	Message    string    `json:"message"`
	StoryOwner bool      `json:"storyOwner"`
	Time       time.Time `json:"time"`
}

func (m *Message) validate() ValidationErrs {
	errs := ValidationErrs{}

	if m.Signal == "" {
		errs["signal"] = "Please enter a valid signal."
	}

	if !bson.IsObjectIdHex(m.StoryID) {
		errs["storyID"] = "Please enter a valid story."
	}

	return errs
}

// HandleSocketMessages is a Go routine that listens infinitely on the broadcast channel, maps the story and
// performs any actions/transformations needed before broadcasting out to all connected clients for that
// story.
func HandleSocketMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-Broadcast
		// Send it out to every client for a specific chat.
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
}
