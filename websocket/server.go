package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/goji/httpauth"
	"github.com/gorilla/websocket"
	"goji.io"
	"goji.io/pat"
)

const (
	writeWait  = 10 * time.Second
	pongWait   = 5 * time.Second
	pingPeriod = pongWait * 9 / 10
)

var upgrader = websocket.Upgrader{} // use default options

// Params for a websocket server
type Params struct {
	Addr     string
	BasePath string
	HTTPUser string
	HTTPPass string
}

type server struct {
	hub      *hub
	basePath string
}

// Run a websocket server
func Run(params *Params, c chan []byte) {
	fmt.Println("initializing websocket...")

	s := &server{
		createHub(c),
		params.BasePath,
	}

	mux := goji.NewMux()
	if params.HTTPUser != "" && params.HTTPPass != "" {
		fmt.Println("basic auth enabled")
		mux.Use(httpauth.SimpleBasicAuth(params.HTTPUser, params.HTTPPass))
	}
	mux.HandleFunc(pat.Get(params.BasePath+"/watch"), s.echo)

	log.Fatal(http.ListenAndServe(params.Addr, mux))
}

func (s *server) echo(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	conn := createConn(ws)
	s.hub.addDestination(conn.id, conn.channel)
	conn.setCloseHandler(func() {
		s.hub.removeDestination(conn.id)
	})
}
