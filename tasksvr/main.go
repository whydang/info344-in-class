package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/info344-s17/in-class/tasksvr/handlers"
	"github.com/info344-s17/in-class/tasksvr/models/tasks"
	"gopkg.in/mgo.v2"
)

const defaultPort = "80"

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}
	addr := host + ":" + port

	//create Mongo Session
	mongoAddr := os.Getenv("MONGOADDR")
	fmt.Printf("dialing mongo server at %s...\n", mongoAddr)
	mongoSession, err := mgo.Dial(mongoAddr)
	if err != nil {
		log.Fatalf("error dialing mongo: %v", err)
	}

	//create TasksStore
	tstore := &tasks.MongoStore{
		Session:        mongoSession,
		DatabaseName:   "tasksdemo",
		CollectionName: "tasks",
	}

	//create handler context
	hctx := &handlers.Context{
		TasksStore: tstore,
	}

	//add handlers
	http.HandleFunc("/v1/tasks", hctx.HandleTasks)
	http.HandleFunc("/v1/tasks/", hctx.HandleSpecificTask)

	fmt.Printf("listening at %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
