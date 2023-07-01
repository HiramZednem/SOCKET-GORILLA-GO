package main

import "github.com/gorilla/websocket"

// aqui estaremos tratando todo lo relacionado con un usuario en especifico


// aqui, map[*Client]bool, esto esta haciendo un mapa donde la clave es el cliente, 
// y el valor es un booleano {hiram: true, brodely: false}
type ClientList map[*Client]bool

type Client struct {
	// Todo cliente tiene a huevo, un conexion socket asignada, por lo que le vamos a guardar la conexion
	connection *websocket.Conn


	//Aqui agregamos el manager, porque un cliente puede hacerla de broadcasting con los usuarios
	manager *Manager
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: conn,
		manager: manager,
	}
} 