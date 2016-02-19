package web

import (
	"fmt"
	"log"
	"time"

	"github.com/cenan/mergen/engine"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

type Client struct {
	send   chan Message
	socket *websocket.Conn
	board  *engine.Board
}

type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

func (client *Client) Read() {
	var message Message
	for {
		if err := client.socket.ReadJSON(&message); err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("Read message: ", message.Name)
		if message.Name == "move" {
			move := engine.Move{}
			err := mapstructure.Decode(message.Data, &move)
			if err != nil {
				panic(err)
			}
			fmt.Println("User moved: ", move)
			client.send <- Message{Name: "move", Data: move.String()}
			client.board.MakeMove(move)
			start := time.Now()
			scr, moves := engine.ParallelNegaMax(client.board, 4, engine.BLACK)
			elapsed := time.Since(start)
			fmt.Println(moves[0])
			if scr < -900 {
				client.send <- Message{Name: "status", Data: "Resign"}
			} else {
				client.board.MakeMove(moves[0])
				client.send <- Message{Name: "move", Data: moves[0].String()}
				client.send <- Message{Name: "status", Data: fmt.Sprintf("time: %v score: %f", elapsed, scr)}
			}
		}
	}
	client.socket.Close()
}

func (client *Client) Write() {
	for msg := range client.send {
		fmt.Println("Writing message: ", msg.Name)
		if err := client.socket.WriteJSON(msg); err != nil {
			fmt.Println(err)
			break
		}
	}
	client.socket.Close()
}

func (c *Client) Close() {
	log.Println("Client Close")
	close(c.send)
}

func NewClient(socket *websocket.Conn, board *engine.Board) *Client {
	return &Client{
		send:   make(chan Message),
		socket: socket,
		board:  board,
	}
}
