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

	rootMux := goji.NewMux()
	watchMux := s.createWatchMux(params.HTTPUser, params.HTTPPass)
	rootMux.HandleFunc(pat.Get(params.BasePath+"/health"), s.health)
	rootMux.Handle(pat.New(params.BasePath+"/*"), watchMux)

	log.Fatal(http.ListenAndServe(params.Addr, rootMux))
}

func (s *server) createWatchMux(user, pass string) *goji.Mux {
	mux := goji.SubMux()
	if user != "" && pass != "" {
		fmt.Println("basic auth enabled")
		mux.Use(httpauth.SimpleBasicAuth(user, pass))
	}
	mux.HandleFunc(pat.Get("/watch"), s.watch)
	return mux
}

func (s *server) watch(w http.ResponseWriter, r *http.Request) {
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

func (s *server) health(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"status\":\"up\"}")
}
