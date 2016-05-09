package main

import (
	"log"

	"github.com/gotschmarcel/goserv"
	"gopkg.in/mgo.v2"
)

func main() {

	dbSession := initAnStartDB()
	defer dbSession.Close()
	setRoutings(goserv.NewServer(), dbSession)
}

func initAnStartDB() *mgo.Session {
	mongo, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatalln(err)
	}
	return mongo
}

func setRoutings(server *goserv.Server, dbSession *mgo.Session) {
	controller := &Controller{dbSession.DB("test").C("newsletters")}
	server.Get("/", controller.ViewAllIssues)
	log.Fatalln(server.Listen(":12345"))
}
