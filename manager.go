package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	// EstE objeto toma el http request y lo transforma en una conexion websocket
	websocketUpgrader = websocket.Upgrader{
		// Estas propiedades se aseguran que nuestros clientes no manden un paquete super gigante de cosas, 
		// aunque en neo hay que ver como manejar esta parte
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {

}

func NewManager() *Manager{
	return &Manager{}
}

func (m *Manager) serveWS(w http.ResponseWriter, r *http.Request) {
	log.Println("new connection")

	websocketUpgrader.CheckOrigin = func(r *http.Request) bool { return true }

	//aqui le pasamos a la funcion que quiere volver la conecion un websocket y aqui pasa la magia
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		// En un futuro habria que trabajar el futuro de errores, como esto es una demo solo se imprime el error
		log.Print(err)
		return
	}

	conn.Close()
}