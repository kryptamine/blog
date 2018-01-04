package base

import (
	"gopkg.in/mgo.v2"
)

var DB *mgo.Database

func InitDB(dbServer string, dbName string) *mgo.Session {
	session, err := mgo.Dial(dbServer)

	if err != nil {
		panic(err)
	}

	DB = session.DB(dbName)

	return session
}
