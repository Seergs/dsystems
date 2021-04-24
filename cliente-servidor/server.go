package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

type Server struct {
	processes []Process
	addresses []string
	counter uint64
}

type Process struct {
	Id      uint64
	Counter uint64
}

func NewServer() *Server {
	s := Server{}
	s.buildInitialProcesses(5)
	
	go s.start()
	time.Sleep(time.Second)
	go s.updateProcesses()
	return &s
}

func (s *Server) addProcess() {
	s.counter++
	process := Process{s.counter, 0}
	s.processes = append(s.processes, process)
}

func (s *Server) updateProcesses() {
	for {
		fmt.Println("-------------")
		for i, p := range (s.processes) {
			fmt.Printf("Proceso %d: %d\n", p.Id, p.Counter)
			s.processes[i].Counter++
		}
		time.Sleep(time.Millisecond * 500)
	}
}


func (s *Server) buildInitialProcesses(count uint64) {
	var i uint64 = 0
	for i < count {
		s.addProcess()
		i++
	}
}

func (s *Server) start() {
	server, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Servidor iniciado en puerto 5000")
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
	var kind string
	err := gob.NewDecoder(c).Decode(&kind) 
	if err != nil {
		fmt.Println(err)
	}
	if kind == "get" {
		s.handleGet(c)
	} else {
		s.handlePost(c)
	}
}

func (s *Server) handleGet(c net.Conn) {
	err := gob.NewEncoder(c).Encode(s.processes[0])
	fmt.Printf("\n\n\n\nCliente conectado, asignando proceso con ID %d...\n\n\n\n\n\n\n\n", s.processes[0].Id)
	if err != nil {
		fmt.Println(err)
	}
	s.processes = s.processes[1:]
}

func (s *Server) handlePost(c net.Conn) {
	var process Process
	err := gob.NewDecoder(c).Decode(&process)
	fmt.Print(process)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n\n\n\nCliente conectado, retornando proceso con ID %d...\n\n\n\n\n\n\n\n", process.Id)
	s.processes = append(s.processes, process)
}

func main() {
	NewServer()

	var quit string
	fmt.Scanln(&quit)
}

