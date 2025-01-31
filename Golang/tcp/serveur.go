package tcp

import (
	"Golang/utils"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"
)

const PORT = "3000"

func Main() {
	// Écoute des connexions sur le port spécifié
	listener, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()
	log.Printf("Server is listening on port %s...\n", PORT)

	const MaxConnections = 100
	sem := make(chan struct{}, MaxConnections)

	for {
		sem <- struct{}{}
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			<-sem
			continue
		}

		log.Printf("New connection established from %s\n", conn.RemoteAddr())
		go func() {
			HandleConnection(conn)
			<-sem
		}()
	}
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(5 * time.Minute)) // Timeout après 5 minutes

	buff := make([]byte, 2048)
	for {
		n, err := conn.Read(buff)
		if err != nil {
			if err.Error() == "EOF" {
				log.Printf("Client disconnected: %s\n", conn.RemoteAddr())
			} else {
				log.Printf("Error reading from client %s: %v\n", conn.RemoteAddr(), err)
			}
			return
		}

		msg := string(buff[:n])
		if len(msg) == 0 {
			log.Printf("Received empty message from %s\n", conn.RemoteAddr())
			continue
		}

		dict := utils.ExtraireListe("Base_de_donnees/dict_anglais.json")
		message := utils.Txt_to_liste(msg)
		reponse := utils.Corrections(message, dict)
		fmt.Println(reponse)
		jsonData, err := json.Marshal(reponse)
		if err != nil {
			log.Printf("Failed to marshal JSON: %v\n", err)
			return
		}

		_, err = conn.Write(jsonData)
		if err != nil {
			log.Printf("Failed to send response to %s: %v\n", conn.RemoteAddr(), err)
			return
		}
	}
}
