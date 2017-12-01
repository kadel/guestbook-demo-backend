package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	database      *mgo.Database
	mongoUser     string
	mongoPassword string
	mongoDb       string
	mongoServer   string
	port          string
	host          string
)

const key string = "guestbook"

type comment struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty" `
	Author    string        `json:"author"`
	Text      string        `json:"text"`
	Timestamp time.Time     `json:"timestamp"`
}

func handleComments(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "OPTIONS":
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Write([]byte{})
	case "POST":
		var data comment
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), 400)
			log.Fatal(err)
		}

		log.Printf("data %#v\n", data)

		c := database.C("comments")

		data.ID = bson.NewObjectId()
		data.Timestamp = time.Now()

		err = c.Insert(&data)
		if err != nil {
			log.Fatal(err)
		}

		output, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Write(output)

	case "GET":
		c := database.C("comments")

		allComments := []comment{}

		err := c.Find(nil).Sort("-timestamp").All(&allComments)
		if err != nil {
			log.Fatal(err)
		}

		output, err := json.MarshalIndent(allComments, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Write(output)

	default:
		errText := fmt.Sprintf("Unsupported method: %s", r.Method)
		fmt.Println(errText)
		http.Error(w, errText, http.StatusMethodNotAllowed)
	}
}

func main() {
	flag.StringVar(&port, "port", "3000", "port")
	flag.StringVar(&host, "host", "0.0.0.0", "host")
	flag.Parse()

	mongoPassword = os.Getenv("MONGODB_PASSWORD")
	if mongoPassword == "" {
		log.Fatal("You have to set MONGODB_PASSWORD")
	}

	mongoUser = os.Getenv("MONGODB_USERNAME")
	if mongoUser == "" {
		log.Fatal("You have to set MONGODB_USERNAME")
	}

	mongoDb = os.Getenv("MONGODB_DATABASE")
	if mongoDb == "" {
		log.Fatal("You have to set MONGODB_DATABASE")
	}

	mongoServer = os.Getenv("MONGODB_SERVER")
	if mongoServer == "" {
		log.Fatal("You have to set MONGODB_SERVER")
	}

	// wait for DB
	var mongo *mgo.Session
	err := errors.New("fake error")
	retry := 0
	maxRetry := 10
	for err != nil {
		retry = retry + 1
		log.Printf("trying connect to db (attempt: %d/%d)", retry, maxRetry)
		mongo, err = mgo.Dial(mongoServer)
		log.Printf("Error: %#v\n", err)
		if err == nil {
			break
		}

		time.Sleep(5 * time.Second)
		if retry >= maxRetry {
			log.Fatal(err)
		}
	}

	err = mongo.DB("db").Login("user", "pass")
	if err != nil {
		log.Fatal(err)
	}
	database = mongo.DB("db")

	http.HandleFunc("/api/comments", handleComments)
	log.Println("Server started: http://" + host + ":" + port)
	log.Fatal(http.ListenAndServe(host+":"+port, nil))
}
