package main

import (
	"log"
	"net/http"
	"sync"

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
	clients ClientList
	
	// sync RWMutex sirve para manejar la concurrencia, para que no se alteren datos al mismo tiempo
	sync.RWMutex
}

func NewManager() *Manager{
	return &Manager{
		clients: make(ClientList),
	}
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


	//Una vez que abrimos una conexion, tenemos que saber que lo tenemos, por lo que vamos a crear un cliente
	client := NewClient(conn, m)
	m.addClient(client)

	//Una vez que el cliente esta aniadido, le agregamos los procesos
	go client.readMessages()
}

func (m *Manager) addClient(client *Client) {
	m.Lock() // esto sirve para que no editemos dos archivos al mismo tiempo
	defer m.Unlock()

	m.clients[client] = true
}

func (m *Manager) removeClient(client *Client) {
	m.Lock() // esto sirve para que no editemos dos archivos al mismo tiempo
	defer m.Unlock()

	if _, ok := m.clients[client]; ok {
		client.connection.Close()
		delete(m.clients, client)
	}
}