package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"strconv"
)

var initialTcpPort = 5000
var initialRpcPort = 8000

const (
	JOIN = iota
	MESSAGE
	LEAVE
)

type Room struct {
	host string
	port string
	topic string
	users []Client
}

type Client struct {
	Username string
	Port string
}

type request struct {
	Action int
	User Client
	Message string
}

func newRoom() *Room {
	room := &Room{}
	room.parseInitialFlags()
	room.host = "http://localhost"

	return room
}

func (r *Room) rpcServer() {
	rpc.Register(r)
	var ln net.Listener
	var err error
	for {
		ln, err = net.Listen("tcp", ":"+strconv.Itoa(initialRpcPort))
		if err == nil {
			break
		}
		initialRpcPort++
	}
	log.Println("Rpc connection setup completed on port", initialRpcPort)
	for {
		c, err := ln.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		go rpc.ServeConn(c)
	}
}

func (r *Room) tcpServer() {
	var ln net.Listener
	var err error
	for {
		ln, err = net.Listen("tcp", ":"+strconv.Itoa(initialTcpPort))
		if err == nil {
			break
		}
		initialTcpPort++
	}
	log.Println("Tcp connection setup completed on port",initialTcpPort)
	log.Println("Room /" + r.topic + " is active, waiting for messages")
	r.port = strconv.Itoa(initialTcpPort)

	for {
		c, err := ln.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		go r.handleClient(c)
	}
}

func (r *Room) handleClient(conn net.Conn) {
	defer conn.Close()
	req := r.decode(conn)

	if req.Action == JOIN {
		r.join(req.User)
	} else if req.Action == MESSAGE {
		r.message(req.User, req.Message)
	} else if req.Action == LEAVE {
		r.leaveRoom(req.User)
	}
}

func (r *Room) join(user Client) {
	r.users = append(r.users, user)
	log.Println("User " + user.Username + " has joined")
	r.broadcastMessage(user.Username, "Hello, I just joined!")
}

func (r *Room) message(user Client, msg string) {
	log.Println(user.Username + ":" + msg)
	r.broadcastMessage(user.Username, msg)
}

func (r *Room) leaveRoom(user Client) {
	i, _ := r.findClientByUsername(user.Username)
	r.users = removeClient(r.users, i)
	r.broadcastMessage(user.Username, "Adios")
	log.Println("User " + user.Username + " has disconnected")
}

func (r *Room) broadcastMessage(from string, msg string) {
	for _, user := range(r.users) {
		if from != user.Username {
			c, err := net.Dial("tcp", ":" + user.Port)
			if err != nil {
				log.Println(err.Error())
				continue
			}
			r.encode(c, from + ": " + msg)
			c.Close()
		}
	}
}

func (r *Room) findClientByUsername(username string) (int, Client) {
	for i, user := range(r.users) {
		if user.Username == username {
			return i, user
		}
	}
	return -1, Client{}
}

func removeClient(s []Client, index int) []Client {
    s[len(s)-1], s[index] = s[index], s[len(s)-1]
    return s[:len(s)-1]
}

func (r *Room) getOnlineUsersCount() int {
	return len(r.users)
}

func (s *Room) decode(conn net.Conn) request {
	var req request
	err := gob.NewDecoder(conn).Decode(&req)
	if err != nil {
		log.Printf("Unable to decode request: %s", err.Error())
	}
	return req
}

func (r *Room) encode(conn net.Conn, data interface {}) {
	err := gob.NewEncoder(conn).Encode(data)
	if err != nil {
		log.Printf("Unable to encode request: %s", err.Error())
	}
}

func (r *Room) GetRoomInfo(empty bool, reply *map[string]string) error {
	log.Println("GetRoomInfo")
	room := make(map[string]string)

	room["Topic"] = r.topic
	room["OnlineUsers"] = strconv.Itoa(r.getOnlineUsersCount())
	room["Host"] = r.host
	room["Port"] = r.port

	*reply = room
	return nil
}

func (r *Room) parseInitialFlags() {
	topic := flag.String("t","general", "The topic of this chatroom")
	flag.Parse()
	r.topic = *topic
}



func main() {
	server := newRoom()
	go server.rpcServer()
	go server.tcpServer()

	var input string
	fmt.Scanln(&input)
}
