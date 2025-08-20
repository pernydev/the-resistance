package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/pernydev/the-resistance/backend/room"
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
	godotenv.Overload()
	r := gin.Default()

	// WebSocket endpoint
	r.GET("/ws", func(c *gin.Context) {
		wsHandler(c.Writer, c.Request)
	})

	r.POST("/rooms/:id", func(c *gin.Context) {
		type joinRoomParams struct {
			Name string `json:"name"`
		}

		var params joinRoomParams
		err := c.BindJSON(&params)
		if err != nil {
			fmt.Println(err)
			return
		}

		roomID := c.Param("id")

		r, err := room.GetRoom(roomID)
		if err != nil {
			fmt.Println(err)
			return
		}
		r.AddPlayer(params.Name)
	})

	r.POST("/rooms/new", func(c *gin.Context) {
		type createRoomParams struct {
			Name string `json:"name"`
		}

		var params createRoomParams
		err := c.BindJSON(&params)
		if err != nil {
			fmt.Println(err)
			return
		}

		room := room.NewRoom()
		token, err := room.AddPlayer(params.Name)
		if err != nil {
			fmt.Println(err)
			c.Status(500)
			return
		}
		c.JSON(200, gin.H{
			"room_id": room.ID,
			"token":   token,
		})
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

	_, token, err := conn.ReadMessage()
	if err != nil {
		return
	}

	tokenData, err := room.VerifyToken(string(token))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(tokenData.ID)

	rm, err := room.GetRoom(tokenData.RoomID)
	if err != nil {
		fmt.Println(err)
		return
	}

	wsSender := room.NewWSSender(conn)
	rm.Players[tokenData.ID].Sender = wsSender
	rm.Update()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}

		type MessageData struct {
			Command string         `json:"command"`
			Data    map[string]any `json:"data"`
		}

		var messageData MessageData
		err = json.Unmarshal(message, &messageData)
		if err != nil {
			continue
		}

		switch messageData.Command {
		case "start":
			rm.CreateGame()
		}

	}
}
