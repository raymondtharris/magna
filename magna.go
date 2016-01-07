package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/magna/magnagraph"
	"github.com/magna/magnalanguage"
)

//current query being processed
var queryTokens []magnagraph.Node
var stemString string

//graph used to solve queries
var magnasMind magnagraph.Graph

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func loadMagnasMind() bool {
	magnasMind = magnagraph.Graph{0, 0, nil, true}
	return true
}

// Handlers for Magana
func testHandler(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		messageType, p, err1 := connection.ReadMessage()
		if err1 != nil {
			return
		}
		fmt.Println(string(p))
		err1 = connection.WriteMessage(messageType, p)
		if err1 != nil {
			return
		}
	}
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Query Received.")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var query magnalanguage.MagnaQueryObject
	err = json.Unmarshal([]byte(string(body[:])), &query)
	if err != nil {
		panic(err)
	}
	fmt.Println(query)
	queryTokens = magnalanguage.TokenizeQuery(query)
	stemString = ""
	magnalanguage.PartOfSpeachTag(queryTokens)
	for _, aNode := range queryTokens {
		magnalanguage.ProcessNode(aNode)
		magnalanguage.AppendToStemmedSring(stemString, aNode)
	}
}

func main() {
	loadMagnasMind()
	http.HandleFunc("/", testHandler)
	http.HandleFunc("/query", queryHandler)
	http.ListenAndServe("localhost:8080", nil)
}
