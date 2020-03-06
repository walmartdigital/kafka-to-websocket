package server

import (
	"fmt"
	"sync"
)

type hub struct {
	source       chan []byte
	destinations *sync.Map
}

func createHub(c chan []byte) *hub {
	hub := &hub{c, &sync.Map{}}
	go hub.readLoop()
	return hub
}

func (h *hub) readLoop() {
	for {
		select {
		case m := <-h.source:
			h.destinations.Range(func(key, value interface{}) bool {
				value.(chan []byte) <- m
				return true
			})
		}
	}
}

func (h *hub) addDestination(id string, c chan []byte) {
	h.destinations.Store(id, c)
	fmt.Printf("conn %s added\n", id)
}

func (h *hub) removeDestination(id string) {
	h.destinations.Delete(id)
	fmt.Printf("conn %s removed\n", id)
}
