package main

import (
	"bufio"
	"encoding/gob"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
)

var initialPort = 3000

const (
	MIDDLEWARE_URL="http://localhost:8080"
)

const (
	JOIN = iota
	MESSAGE
	LEAVE
)

type request struct {
	Action int
	User Client
	Message string
}


type Message struct {
	From string
	Text string
}

type Client struct {
	Username string
	Port string
	host string
	roomHost string
	roomPort string
}

type Room struct{
	Topic string
	OnlineUsers string
	Host string
	Port string
}

func newClient() *Client {
	c := &Client{}
	c.parseInitialFlags()
	fmt.Printf("\n\nHola " + c.Username + ", bienvenido!\n\n")

	return c
}

func (c *Client) parseInitialFlags() {
	username := flag.String("u","unknown", "The username")
	flag.Parse()
	c.Username = *username
}

func (c *Client) listen() {
	var ln net.Listener
	var err error
	for {
		ln, err = net.Listen("tcp", ":"+strconv.Itoa(initialPort))
		if err == nil {
			break
		}
		initialPort++
	}
	c.Port = strconv.Itoa(initialPort)
	defer ln.Close()

	for {
		msg, err := ln.Accept()
		if err != nil {
			fmt.Println("Algo salio mal")
			continue
		}
		go c.handleIncomingMessage(msg)
	}
}

func (c *Client) handleIncomingMessage(msg net.Conn) {
	response := c.decode(msg)
	fmt.Println(response)
}

func (c *Client) getAllRooms() []Room {
	res, err := http.Get(MIDDLEWARE_URL+"/rooms")
	if err != nil {
		fmt.Println("Algo salió mal obteniendo los chatrooms")
		return nil
	}
	defer res.Body.Close()

	var rooms []Room
	err = json.NewDecoder(res.Body).Decode(&rooms)
	if err != nil {
		fmt.Println(err.Error())
	}
	return rooms
}

func (c *Client) printChatrooms(rooms []Room) {
	fmt.Println("Estos son los chatrooms disponibles")
	for _, roomInfo := range(rooms) {
		fmt.Println("/" + roomInfo.Topic + " (" + roomInfo.OnlineUsers + " usuarios conectados)")
	}
}

func (c *Client) getRoomToJoin(rooms []Room) Room {
	fmt.Printf("\nA cual deseas ingresar? /")
	roomName := getStringFromUser()

	for _, room := range(rooms) {
		if room.Topic == roomName {
			return room
		}
	}
	return Room{}
}

func (c *Client) joinRoom() {
	rooms := c.getAllRooms()
	c.printChatrooms(rooms)
	roomToJoin := c.getRoomToJoin(rooms)
	conn, err := net.Dial("tcp", ":"+roomToJoin.Port)
	if err != nil {
		fmt.Println("Algo salió mal conectando con la sala de chat",err.Error())
		return
	}
	defer conn.Close()

	c.encode(conn, request {JOIN, *c, ""})
	c.roomHost = roomToJoin.Host
	c.roomPort = roomToJoin.Port
	fmt.Println("Listo! Puedes empezar a chatear")
	fmt.Printf("Para salir usa el comando /exit\n\n")
}

func (c *Client) sendMessage(msg string) {
	conn, err := net.Dial("tcp", ":"+c.roomPort)
	if err != nil {
		fmt.Println("Algo salio mal al enviar el mensaje")
	}
	defer conn.Close()

	c.encode(conn, request {MESSAGE, *c, msg})
}

func (c *Client) leaveRoom() {
	conn, err := net.Dial("tcp", ":"+c.roomPort)
	if err != nil {
		fmt.Println("Algo salio mal al intentar salir del chat")
	}
	defer conn.Close()
	c.encode(conn, request {LEAVE, *c, ""})
	fmt.Println("Hasta luego " + c.Username + ", vuelve pronto")
}

func (c *Client) encode(conn net.Conn, data interface {}) {
	err := gob.NewEncoder(conn).Encode(data)
	if err != nil {
		fmt.Println("Algo salio mal")
	}
}

func (c *Client) decode(conn net.Conn) string {
	var res string
	err := gob.NewDecoder(conn).Decode(&res)
	if err != nil {
		fmt.Printf("Algo salio mal: %s", err.Error())
	}
	return res
}

func (c *Client) getMessageFromClient() string {
	return getStringFromUser()
}


func main() {
	client := newClient()
	go client.listen()

	client.getAllRooms()
	client.joinRoom()

	exitCmd := "/exit"
	message := client.getMessageFromClient()
	for  {
		if message == exitCmd {
			client.leaveRoom()
			break
		}
		client.sendMessage(message)
		message = client.getMessageFromClient()
	}
	
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