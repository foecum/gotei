package sockets

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/foecum/gotei2.0/logger"
)

var (
	hub *Hub
	// The port on which we are hosting the reload server has to be hardcoded on the client-side too.
	reloadAddress = ":12450"
)

var log = logger.New()

// StartReloadServer ...
func StartReloadServer() {
	hub = newHub()
	go hub.run()
	http.HandleFunc("/reload", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	go startServer()
	log.Notice(fmt.Sprintf("Reload server listening at %v", reloadAddress))
}

func startServer() {
	err := http.ListenAndServe(reloadAddress, nil)

	if err != nil {
		log.Error(fmt.Sprintf("Failed to start up the Reload server: %v", err.Error()))
		return
	}
}

// SendReload ...
func SendReload() {
	message := bytes.TrimSpace([]byte("reload"))
	hub.broadcast <- message
}
