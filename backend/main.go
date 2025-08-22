package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/pernydev/the-resistance/backend/room"
	"github.com/pernydev/the-resistance/backend/room/game"
	cors "github.com/rs/cors/wrapper/gin"
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
	r.Use(cors.Default())

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
		token, err := r.AddPlayer(params.Name)
		if err != nil {
			fmt.Println(err)
			c.Status(500)
			return
		}
		c.JSON(200, gin.H{
			"room_id": roomID,
			"token":   token,
		})
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

	if rm.Players[tokenData.ID].Sender != nil {
		rm.Players[tokenData.ID].Sender.Close()
		rm.Players[tokenData.ID].Sender = nil
	}

	wsSender := room.NewWSSender(conn)
	rm.Players[tokenData.ID].Sender = wsSender
	rm.Update()

	isHost := rm.HostID == tokenData.ID

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}

		type MessageData struct {
			Command string `json:"command"`
			Data    string `json:"data"`
		}

		var messageData MessageData
		err = json.Unmarshal(message, &messageData)
		if err != nil {
			continue
		}

		switch messageData.Command {
		case "start":
			if !isHost {
				continue
			}
			rm.CreateGame()
		case "settings":
			{
				if !isHost {
					continue
				}
				var data room.GameSettings
				err := json.Unmarshal([]byte(messageData.Data), &data)
				if err != nil {
					fmt.Println("not ok")
					continue
				}
				rm.Settings = data
			}
		case "continue":
			{
				if !isHost {
					continue
				}
				switch rm.Game.State {
				case game.GameStateCompositionCreation:
					rm.Game.SetState(game.GameStateVoting)
				case game.GameStateVotingReveal:
					if rm.Game.CurrentCompositionRejected {
						rm.Game.SetState(game.GameStateVoting)
						continue
					}
					rm.Game.SetState(game.GameStateMission)
				}

			}
		case "next":
			{
				if !isHost {
					continue
				}
				rm.Game.NextMission()
			}
		case "add":
			{
				if rm.Game.PlayerOrder[rm.Game.CurrentPlayer] != tokenData.ID {
					continue
				}

				rm.Game.Players[messageData.Data].IsInComposition = !rm.Game.Players[messageData.Data].IsInComposition
				if rm.Game.AmountInComposition() > rm.Game.Missions[rm.Game.Mission].ParticipantCount {
					rm.Game.Players[messageData.Data].IsInComposition = !rm.Game.Players[messageData.Data].IsInComposition
				}
			}
		case "vote":
			{
				vote := game.Vote(messageData.Data == "approve")
				rm.Game.Vote(tokenData.ID, vote)
			}

		case "mission_submit":
			{
				if !rm.Game.Players[tokenData.ID].IsInComposition {
					continue
				}
				result := game.MissionResult(messageData.Data)
				rm.Game.Missions[rm.Game.Mission].Submit(result)
			}
		}

		rm.Update()

	}
}
