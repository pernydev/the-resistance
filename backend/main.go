package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/pernydev/the-resistance/backend/game"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Allow all origins (for development only)
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	r := gin.Default()

	// WebSocket endpoint
	r.GET("/ws", func(c *gin.Context) {
		wsHandler(c.Writer, c.Request)
	})

	r.POST("/rooms/new", func(c *gin.Context) {
		type createRoomParams struct {
			Name string `json:"string"`
		}

		var params createRoomParams
		c.BindJSON(params)

		room := game.NewRoom()
		token, err := room.AddPlayer(params.Name)

	})

	r.Run(":8080")
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		// Handle error
		return
	}
	defer conn.Close()

	// WebSocket handling loop
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			// Handle error
			break
		}

		// Echo the message back to client
		if err := conn.WriteMessage(messageType, message); err != nil {
			// Handle error
			break
		}
	}
}
