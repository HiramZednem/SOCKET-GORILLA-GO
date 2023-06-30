package main

import (
	"log"
	"net/http"
)

func main() {
	setupAPI()

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupAPI() {
	manager := NewManager()

	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/ws", manager.serveWS)
	
	http.HandleFunc("/prueba", func(w http.ResponseWriter, r *http.Request) {
		log.Println("accedi al endpoint")
	})
}

func SetupResponse(w *http.ResponseWriter, req *http.Request) {
	/*
	   (*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	   (*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	   (*w).Header().Set("Access-Control-Allow-Origin", "https://neodatalicencias.web.app")
	   (*w).Header().Set("Access-Control-Allow-Origin", "https://neodatacontrollicencias.web.app")
	   (*w).Header().Set("Access-Control-Allow-Origin", "https://neodatalicenciasdesarrollo.web.app")
	   (*w).Header().Set("Access-Control-Allow-Origin", "https://neocontrollicenciasdesarrollo.web.app")
	   c := cors.New(cors.Options{
		   AllowedOrigins: []string{"http://foo.com", "http://foo.com:8080"},
		   AllowCredentials: true,
		   // Enable Debugging for testing, consider disabling in production
		   Debug: true,
	   })
 
	   // Insert the middleware
	   handler = c.Handler(handler)
	*/
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
 
 }