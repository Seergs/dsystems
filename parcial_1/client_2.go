package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

const (
	join = iota
	username
	message
	messages
)


type client struct {
	id string
	Username string
	Port int
}

type request struct {
	ClientId string
	Action int
	Username string
	Message string
}

type response struct {
	Message string
	Messages [] Message
}

type Message struct {
	From client
	Text string
	Date time.Time
}

func newClient() *client {
	return &client {
		id: generateId(10),
	}
}


func (c *client) setup() {
	conn, err := net.Dial("tcp", ":5000")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	c.encode(conn, request {c.id, join, c.id, ""})
	port, err := strconv.Atoi(c.decode(conn).Message)

	c.Port = port
}

func (c *client) listen() {
	listener, err := net.Listen("tcp", ":" + strconv.Itoa(c.Port))
	if err != nil {
		log.Fatalf("Unable to client listener: %s", err.Error())
		return
	}
	defer listener.Close()

	for {
		msg, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %s", err.Error())
			continue
		}
		go c.messageHandler(msg)
	}
}

func (c *client) messageHandler(msg net.Conn) {
	message := c.decode(msg)
	fmt.Println("\n" + message.Message)
}

func (c *client) isValidUsername(u string) bool {
	conn, err := net.Dial("tcp", ":5000")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	c.encode(conn, request {c.id, username, u, ""})
	isValidUsername := c.decode(conn)
	return isValidUsername.Message == "1"
}

func (c *client) sendMessage(msg string) {
	conn, err := net.Dial("tcp", ":5000")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	c.encode(conn, request {c.id, message, c.Username, msg})
}

func (c *client) getAllMessages() {
	conn, err := net.Dial("tcp", ":5000")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	c.encode(conn, request {c.id, messages, c.Username, ""})
	msgs := c.decode(conn).Messages
	displayAllMessages(msgs)
}

func (c *client) decode(conn net.Conn) response {
	var res response
	err := gob.NewDecoder(conn).Decode(&res)
	if err != nil {
		fmt.Printf("Algo salio mal: %s", err.Error())
	}
	return res
}

func (c *client) encode(conn net.Conn, data interface {}) {
	err := gob.NewEncoder(conn).Encode(data)
	if err != nil {
		log.Fatalf("Algo salio mal")
	}
}

func displayAllMessages(msgs []Message) {
	fmt.Printf("\n\n\n\n\n\n\n\n")
	fmt.Printf("------------------------\n")
	fmt.Println("-> Todos los mensajes")
	fmt.Print("------------------------")
	for _, m := range msgs{
		printMessage(m)
	}
	fmt.Println("------------------------")
	fmt.Printf("\n\n\n\n\n\n")
}

func printMessage(m Message) {
	fmt.Printf("\n\nDe ")
	fmt.Println(m.From.Username)
	fmt.Println(m.Text)
	fmt.Printf("el ")
	fmt.Print(m.Date.Format("06-Jan-02"))
	fmt.Printf("\n\n")
}

func main() {
	rand.Seed(time.Now().Unix())

	client := newClient()
	client.setup()
	go client.listen()
	fmt.Print("Antes de empezar a chatear, elige tu username: ")
	username := getStringFromUser()
	isValidUsername := client.isValidUsername(username)
	for !isValidUsername {
		fmt.Print("Ya esta ocupado, elige otro: ")
		username = getStringFromUser()
		isValidUsername = client.isValidUsername(username)
	}
	client.Username = username
	fmt.Println("Bienvenido", client.Username)
	exit := false

	for !exit {
		displayMenu()
		option := getIntFromUser()
		if option == 1 {
			fmt.Print("Chat: ")
			msg := getStringFromUser()
			client.sendMessage(msg)
		} else if option == 2 {

		} else if option == 3 {
			client.getAllMessages()
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

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateId(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}