package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

/*
TODO: Implement the code in this file, according to the comments.
If you haven't yet read the assigned reading, now would be a
good time to do so:
- Read the Overview section of the Gorilla WebSockets package
https://godoc.org/github.com/gorilla/websocket
- Read the Writing WebSocket Client Application
https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API/Writing_WebSocket_client_applications
*/

//WebSocketsHandler is a handler for WebSocket upgrade requests
type WebSocketsHandler struct {
	notifier *Notifier
	upgrader *websocket.Upgrader
}

//NewWebSocketsHandler constructs a new WebSocketsHandler
func NewWebSocketsHandler(notifer *Notifier) *WebSocketsHandler {
	return &WebSocketsHandler{
		notifier: notifer,
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
	}
}

//ServeHTTP implements the http.Handler interface for the WebSocketsHandler
func (wsh *WebSocketsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("received websocket upgrade request")
	//TODO: Upgrade the connection to a WebSocket, and add the
	//new websock.Conn to the Notifier. See
	//https://godoc.org/github.com/gorilla/websocket#hdr-Overview
	conn, err := wsh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("error upgrading websocker: %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Println(wsh.notifier.clients)
	wsh.notifier.AddClient(conn)
	fmt.Println(wsh.notifier.clients)
}

//Notifier is an object that handles WebSocket notifications
type Notifier struct {
	clients []*websocket.Conn
	eventQ  chan []byte
	//TODO: add a mutex or other channels to
	//protect the `clients` slice from concurrent use.
	mx *sync.RWMutex
}

//NewNotifier constructs a new Notifier
func NewNotifier() *Notifier {
	//TODO: construct a new Notifier
	//and call the .start() method on
	//a new goroutine to start the
	//event notification loop
	notifier := &Notifier{
		clients: make([]*websocket.Conn, 0),
		eventQ:  make(chan []byte),
		mx:      &sync.RWMutex{},
	}
	go notifier.start()
	return notifier
}

//AddClient adds a new client to the Notifier
func (n *Notifier) AddClient(client *websocket.Conn) {
	log.Println("adding new WebSockets client")
	//TODO: add the client to the `clients` slice
	//but since this can be called from multiple
	//goroutines, make sure you protect the `clients`
	//slice while you add a new connection to it!
	n.mx.Lock()
	n.clients = append(n.clients, client)
	fmt.Println(n.clients)
	n.mx.Unlock()

	//also process incoming control messages from
	//the client, as described in this section of the docs:
	//https://godoc.org/github.com/gorilla/websocket#hdr-Control_Messages
	for {
		if _, _, err := client.NextReader(); err != nil {
			client.Close()
			break
		}
	}
}

//Notify broadcasts the event to all WebSocket clients
func (n *Notifier) Notify(event []byte) {
	log.Printf("adding event to the queue")
	//TODO: add `event` to the `n.eventQ`
	n.eventQ <- event
}

//start starts the notification loop
func (n *Notifier) start() {
	log.Println("starting notifier loop")
	//TODO: start a never-ending loop that reads
	//new events out of the `n.eventQ` and broadcasts
	//them to all WebSocket clients.
	//If you use additional channels instead of a mutex
	//to protext the `clients` slice, also process those
	//channels here using a non-blocking `select` statement
	for {
		event := <-n.eventQ
		for _, client := range n.clients {
			err := client.WriteMessage(websocket.TextMessage, event)
			if err != nil {
				fmt.Errorf("error writing message: %v", err)
				return
			}
		}
	}
}
