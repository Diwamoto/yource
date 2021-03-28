package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocket サーバーにつなぎにいくクライアント
var clients = make(map[*websocket.Conn]bool)

// クライアントから受け取るメッセージを格納
var broadcast = make(chan Message)

// WebSocket 更新用
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true //TODO:ドメイン判定を作る
	},
}

// クライアントからは JSON 形式で受け取る
type Message struct {
	Message string
}

// クライアントのハンドラ
func HandleClients(w http.ResponseWriter, r *http.Request) {
	// ゴルーチンで起動
	go broadcastMessagesToClients()
	// websocket の状態を更新
	websocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("error upgrading GET request to a websocket::", err)
	}
	// websocket を閉じる
	defer websocket.Close()

	clients[websocket] = true

	for {
		var message Message
		// メッセージ読み込み
		err := websocket.ReadJSON(&message)
		if err != nil {
			log.Printf("error occurred while reading message: %v", err)
			delete(clients, websocket)
			break
		}
		// メッセージを受け取る
		broadcast <- message
	}
}

func WSserver() {
	// localhost:8080 でアクセスした時に index.html を読み込む
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/chat", HandleClients)
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal("error starting http server::", err)
		return
	}
}

func broadcastMessagesToClients() {
	for {
		// メッセージ受け取り
		message := <-broadcast
		// クライアントの数だけループ
		for client := range clients {
			//　書き込む
			err := client.WriteJSON(message)
			if err != nil {
				log.Printf("error occurred while writing message to client: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
