package main

import (
	"log"

	"github.com/gorilla/websocket"
)

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

func (c *Client) readMessages() {
	defer func() { // esta es una funcion anonima que se ejecuta cuando ya se termino de ejecutar toda la funcion readMessages()
		c.manager.removeClient(c)
	}()

	for {
		messageType, payload, err := c.connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure ) {
				log.Printf("error reading message: %v", err)
			}

			break
		}

		log.Println(messageType)
		log.Println(payload)
	}
}
// TODO: HACER LO DE LEER Y ESCRIBIR MENSAJES TOMANDO EN CUENTA QUE GORILLA PACKAGE ACTUALMENTE SOLO PERMITE UN ESCRITOR DE 
// HILO A LA VEZ, POR QUE TOCA APRENDER BIEN QUE PEDO CON LA CONCURRENCIA