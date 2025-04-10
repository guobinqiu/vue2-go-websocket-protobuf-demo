package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/guobinqiu/vue2-go-websocket-protobuf-demo/chat"
	"google.golang.org/protobuf/proto"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	for {
		_, msgBytes, err := ws.ReadMessage()
		if err != nil {
			log.Printf("error: %v", err)
			break
		}

		chatMsg := &chat.ChatMessage{}
		if err := proto.Unmarshal(msgBytes, chatMsg); err != nil {
			log.Printf("Failed to unmarshal: %v", err)
			continue
		}

		fmt.Printf("Received message from %s: %s\n", chatMsg.User, chatMsg.Text)

		// echo back
		if buf, err := proto.Marshal(chatMsg); err == nil {
			ws.WriteMessage(websocket.BinaryMessage, buf)
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)
	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
