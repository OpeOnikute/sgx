package db

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

type DBObject struct {
	Server   string
	Database string
}

// Session ...
var Session *mgo.Database

const (
	COLLECTION = "story"
	SERVER     = "localhost"
	DATABASE   = "sgx"
)

// Connect ...
func Connect() {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		log.Fatal(err)
	}
	Session = session.DB(DATABASE)
}
