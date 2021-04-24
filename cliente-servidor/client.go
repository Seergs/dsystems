package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

type Process struct {
	Id      uint64
	Counter uint64
}

type Client struct {
	process Process
}

func NewClient() *Client {
	c := Client{}

	return &c
}

func (c *Client) getNewProcessFromServer() {
	client, err := net.Dial("tcp", ":5000")
	defer client.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	kind := "get"
	err = gob.NewEncoder(client).Encode(&kind)
	if err != nil {
		fmt.Println(err)
		return
	}

	var process Process
	err = gob.NewDecoder(client).Decode(&process)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\n\n\n\nProceso con ID %d recibido del servidor\n", process.Id)
	fmt.Println("Para detener la ejecución presione Enter")
	fmt.Printf("\n\n\n\n\n\n\n\n\n")
	c.process = process

}

func (c *Client) updateProcesses() {
	for {
		if c.process.Id != 0 {
			fmt.Printf("Proceso %d: %d\n", c.process.Id, c.process.Counter)
			c.process.Counter++
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func (c *Client) returnProcessToServer() {
	client, err := net.Dial("tcp", ":5000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	kind := "post"
	err = gob.NewEncoder(client).Encode(&kind)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = gob.NewEncoder(client).Encode(c.process)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Proceso con ID", c.process.Id, "retornado al servidor")
	fmt.Println("Presione Enter para salir")
	c.process = Process{}
}


func main() {
	client := NewClient()
	go client.getNewProcessFromServer()
	go client.updateProcesses()

	var input string
	fmt.Scanln(&input)
	go client.returnProcessToServer()
	fmt.Scanln(&input)
}

