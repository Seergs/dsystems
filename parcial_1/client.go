package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"time"
)

type Client struct {
	identifier string
	messages []Message
	port string
}

type Request struct {
	Action string
	Message Message
	Username string
}

type Message struct {
	Text string
	Date time.Time
	From string
}

func NewClient() *Client {
	c := Client{identifier: ""}
	c.getPort()
	c.createUsername()
	go c.start()
	return &c
}

func (c *Client) start() {
	client, err := net.Listen("tcp", ":50001")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		_, err := client.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Mensaje recibido desde el servidor")
	}
}

func (c *Client) createUsername() {
	fmt.Print("Antes de empezar a chatear, elige tu username: ")
	identifier := getStringFromUser()
	c.validateUsername(identifier)
	for c.identifier == "" {
		fmt.Print("Username ocupado, elige otro: ")
		c.validateUsername(getStringFromUser())
	}
}

func (c *Client) getPort() {
	client, err := net.Dial("tcp", ":5000")
	defer client.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	request := Request {"getPort", Message{}, ""}
	err = gob.NewEncoder(client).Encode(request)
	if err != nil {
		fmt.Println(err)
	}

	err = gob.NewDecoder(client).Decode(&c.port)
	if err != nil {
		fmt.Println("Algo salio mal")
	}
}

func (c *Client) validateUsername(u string) {
	client, err := net.Dial("tcp", ":5000")
	defer client.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	request := Request {"validateUsername", Message{}, u}
	err = gob.NewEncoder(client).Encode(request)
	if err != nil {
		fmt.Println(err)
	}

	var isValidUsername bool
	err = gob.NewDecoder(client).Decode(&isValidUsername)
	if err != nil {
		fmt.Println("Username ya usado, elige otro")
	}
	if isValidUsername {
		c.identifier = u
	}
}

func (c *Client) sendTextMessage(s string) {
	client, err := net.Dial("tcp", ":5000")
	defer client.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	message := Message {s, time.Now(), c.identifier}
	request := Request {"sendMessage", message, c.identifier}
	err = gob.NewEncoder(client).Encode(request)
	if err != nil {
		fmt.Println(err)
	}
}

func (c *Client) getAllMessages() {
	client, err := net.Dial("tcp", ":5000")
	defer client.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	request := Request {"getMessages", Message{}, c.identifier}
	err = gob.NewEncoder(client).Encode(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = gob.NewDecoder(client).Decode(&c.messages)
	if err != nil {
		fmt.Println(err)
	}
	c.showAllMessages()
}

func (c *Client) showAllMessages() {
	fmt.Printf("\n\n\n\n\n\n\n\n")
	fmt.Printf("----------------------------\n")
	fmt.Println("-> Todos los mensajes")
	fmt.Print("------------------------")
	for _, m := range c.messages {
		printMessage(m, c.identifier)
	}
	fmt.Println("----------------------------")
	fmt.Printf("\n\n\n\n\n\n")
}

func printMessage(m Message, clientUsername string) {
	fmt.Println()
	if m.From != clientUsername {
		fmt.Print("de ")
		fmt.Println(m.From)
		fmt.Println(m.Text)
		fmt.Printf("el ")
		fmt.Print(m.Date.Format("06-Jan-02"))
		fmt.Printf("\n")
	}
}

func main() {
	client := NewClient()
	//client.getPort()
	exit := false

	for !exit {
		displayMenu()
		option := getIntFromUser()
		if option == 1 {
			fmt.Print("Que quieres decir?: ")
			client.sendTextMessage(getStringFromUser())
		} else if option == 2 {

		} else if option == 3 {
			go client.getAllMessages()
		} else if option == 4 {
			exit = true
		} else {
			fmt.Println("Opcion no valida")
		}
	}
	
	var input string
	fmt.Scanln(&input)
}

func displayMenu() {
	fmt.Println("1. Enviar mensaje")
	fmt.Println("2. Enviar archivo")
	fmt.Println("3. Mostrar chat")
	fmt.Println("4. Salir")
	fmt.Print("Opcion: ")
}

const inputDelimiter = '\n'
func getStringFromUser() string {
	s := ""
	for s == "" {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		s = scanner.Text()
	}
	return s
}

func getIntFromUser() int64 {
	var op int64
	fmt.Scan(&op)

	return op
}