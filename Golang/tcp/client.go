package tcp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func Client() {
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		fmt.Println("Failed to connect to server:", err)
		return
	}
	defer conn.Close()
	for {
		reader := bufio.NewReader(os.Stdin)
		msg, _ := reader.ReadString('\n')
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("Failed to send messdage to server:", err)
			return
		}
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Server closed connection:", err)
			return
		}
		var data []string
		err = json.Unmarshal(buffer[:n], &data)
		if err != nil {
			fmt.Println("Failed to unmarshal JSON:", err)
			return
		}
		for _, value := range data {
			fmt.Println(value)
		}
	}
}
