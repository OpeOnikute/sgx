package store

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User holds the main user info
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// JwtToken as implied.
type JwtToken struct {
	Token string `json:"token"`
}

// Exception is an attempt to define errors.
type Exception struct {
	Message string `json:"message"`
}

// Player defines the info we need for players.
type Player struct {
	UID   string `json:"uid"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// StoryFormat defines how paragraph requests are structured.
type StoryFormat struct {
	UserID string    `json:"userID"`
	Text   string    `json:"text"`
	Time   time.Time `json:"time"`
}

// Story defines the fields for story objects.
type Story struct {
	ID         bson.ObjectId `bson:"_id" json:"_id,omitempty"`
	Title      string        `json:"title"`
	InviteCode string        `json:"inviteCode,omitempty"`
	PlayerOne  Player        `json:"playerOne"`
	PlayerTwo  Player        `json:"playerTwo"`
	Content    []StoryFormat `json:"content"`
	SocketURL  string        `json:"socketurl,omitempty"`
	Status     string        `json:"status"`
	Created    time.Time     `json:"created"`
}

// Stories is just a list of stories, really.
type Stories []Story
