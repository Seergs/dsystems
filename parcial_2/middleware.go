package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"net/rpc"
	"strings"
)


type Middleware struct {
	rooms []string
}

func newMiddleware() *Middleware {
	middleware := Middleware{}
	middleware.parseInitialFlags()
	return &middleware
}

func (m *Middleware) getAllRooms(res http.ResponseWriter, req *http.Request) {
	var rooms []map[string]string
	for _, url := range(m.rooms) {
		c, err := rpc.Dial("tcp", ":"+url)
		if err != nil {
			log.Println(err.Error())
		}
		var roomInfo map[string]string
		err = c.Call("Room.GetRoomInfo", true, &roomInfo)
		if err != nil {
			log.Println(err.Error())
		} else {
			log.Println("Room info for " + url + " is ")
			log.Println(roomInfo)
			rooms = append(rooms, roomInfo)
		}
	}
	roomsJson, _ := toJSON(rooms)
	res.Write(roomsJson)
}

func (m *Middleware) parseInitialFlags() {
	rooms := flag.String("rooms","8000,8001,8002", "The ports where the chatrooms are running")
	flag.Parse()
	m.rooms = strings.Split(*rooms, ",")
}

func toJSON(data interface {}) ([]byte, error) {
	dataJson, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		return dataJson, nil
	}
	return dataJson, err
}

func main() {
	middleware := newMiddleware()
	http.HandleFunc("/rooms", middleware.getAllRooms)

	log.Println("Middleware started")
	http.ListenAndServe(":8080", logRequest(http.DefaultServeMux))
}


func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}