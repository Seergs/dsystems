package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"time"
)

type Server struct {
	Clients []string
	Messages []Message
}

type Message struct {
	Text string
	Date time.Time
	From string
}

type Request struct {
	Action string
	Message Message
	Username string
}

func NewServer() *Server {
	s := Server {}
	go s.start()
	return &s
}

func (s *Server) start() {
	server, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := server.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go s.clientHandler(c)
	}
}

func (s *Server) clientHandler(c net.Conn) {
	defer c.Close()

	var request Request
	err := gob.NewDecoder(c).Decode(&request) 
	if err != nil {
		fmt.Println(err)
	}
	log.Println("Request received with action", request.Action)
	if request.Action == "validateUsername" {
		s.validateUsername(request.Username, c)
	} else if request.Action == "sendMessage" {
		s.handleMessage(request, c)
	} else if request.Action == "sendFile" {
		s.handleFile(c)
	} else if request.Action == "getMessages" {
		s.getAllMessages(request, c)
	} else if request.Action == "disconnect" {
		s.removeClient(c)
	} else {
		fmt.Println("Unhandled action")
	}
}

func (s *Server) addClient(id string) {
	s.Clients = append(s.Clients, id)
}

func (s *Server) validateUsername(u string, c net.Conn) {
	log.Println("Validating username", u)
	if existsInSlice(u, s.Clients) {
		log.Println("Username", u, "already used")
		err := gob.NewEncoder(c).Encode(false)
		if err != nil {
			return
		}
	} else {
		err := gob.NewEncoder(c).Encode(true)
		if err != nil {
			return
		}
		s.addClient(u)
		log.Println("Username", u, "is not taken")
	}
}

func  (s *Server) handleMessage(r Request, c net.Conn,) {
	log.Println("Received a message")
	message := Message {r.Message.Text, r.Message.Date, r.Username}
	s.Messages = append(s.Messages, message)
}

func (s *Server) handleFile(c net.Conn) {

}

func (s *Server) getAllMessages(r Request, c net.Conn) {
	log.Println("Sending all messages to ", r.Username)
	err := gob.NewEncoder(c).Encode(s.Messages)
	if err != nil {
		return
	}
	log.Println("Messages sent")
}

func (s *Server) removeClient(c net.Conn) {

}

func (s *Server) showAllMessages() {
	fmt.Printf("\n\n\n\n\n\n\n\n")
	fmt.Printf("------------------------\n")
	fmt.Println("-> Todos los mensajes")
	fmt.Print("------------------------")
	for _, m := range s.Messages {
		printMessage(m)
	}
	fmt.Println("------------------------")
	fmt.Printf("\n\n\n\n\n\n")
}

func printMessage(m Message) {
	fmt.Printf("\n\nDe ")
	fmt.Println(m.From)
	fmt.Println(m.Text)
	fmt.Printf("el ")
	fmt.Print(m.Date.Format("06-Jan-02"))
	fmt.Printf("\n\n")
}

func main() {
	server := NewServer()
	fmt.Println("Servidor escuchando mensajes...")

	exit := false

	for !exit {
		displayMenu()
		option := getIntFromUser()
		
		if option == 1 {
			server.showAllMessages()
		} else if option == 2 {

		} else if option == 3 {  
			exit = true
		} else {
			fmt.Println("Opcion invalida")
		}
	}
}

func displayMenu() {
	fmt.Println("1. Mostrar mensajes/archivos")
	fmt.Println("2. Hacer backup de mensajes")
	fmt.Println("3. Terminar servidor")
	fmt.Println("Opcion: ")
}

func getStringFromUser() string {
	var line string

	fmt.Scan(&line)

	return line
}

func getIntFromUser() int64 {
	var op int64
	fmt.Scan(&op)

	return op
}

func existsInSlice(s string, slice []string) bool {
    for _, a := range slice {
        if a == s {
            return true
        }
    }
    return false
}