package tcp

import (
	"elp_projet/utils"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type Server struct {
	listenAddr string
	ln         net.Listener
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}
func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln
	go s.AcceptLoop()
	return nil
}
func (s *Server) AcceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}
		go HandleConnection(conn)
	}

}
func HandleConnection(conn net.Conn) {
	defer conn.Close()
	buff := make([]byte, 2048)
	for {
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Printf("Connection closed by client:%s\n", conn.RemoteAddr().String())
			return
		}

		msg := string(buff[:n])
		dict := utils.ExtraireListe("Base_de_donnees/dict_anglais.json")
		message := utils.Txt_to_liste(msg)
		reponse := utils.Corrections(message, dict)
		jsonData, err := json.Marshal(reponse)
		if err != nil {
			fmt.Println("Failed to marshal JSON:", err)
			return
		}
		_, err = conn.Write(jsonData)
		if err != nil {
			fmt.Println("Failed to send response to client:", err)
			return
		}
	}
}
func Main_server() {
	server := NewServer(":3000")
	log.Fatal(server.Start())

}
