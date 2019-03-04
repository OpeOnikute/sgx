package store

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

type Player struct {
	UID   string `json:"uid"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type StoryFormat struct {
	UserID string    `json:"userID"`
	Text   string    `json:"text"`
	Time   time.Time `json:"time"`
}

type Story struct {
	ID        bson.ObjectId `bson:"_id" json:"_id,omitempty"`
	Title     string        `json:"title"`
	PlayerOne Player        `json:"playerOne"`
	PlayerTwo Player        `json:"playerTwo"`
	Content   []StoryFormat `json:"content"`
	SocketURL string        `json:"socketurl,omitempty"`
	Status    string        `json:"status"`
	Created   time.Time     `json:"created"`
}

type Stories []Story
