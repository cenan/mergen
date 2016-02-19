package web

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "net/http/pprof"

	"github.com/cenan/mergen/engine"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var homeTempl = template.Must(template.ParseFiles("web/templates/index.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	homeTempl.Execute(w, nil)
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("NEW CONNECTION")
	b := engine.NewBoard()
	b.Reset()
	client := NewClient(ws, b)
	defer client.Close()
	go client.Write()
	client.Read()
}

func StartServer(port int) {
	fmt.Println("starting server")
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", serveWs)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
