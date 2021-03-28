package main

import "fmt"

type Imagen struct {
	titulo  string
	formato string
	canales string
}

type Audio struct {
	titulo   string
	formato  string
	duracion int64
}

type Video struct {
	titulo  string
	formato string
	frames  int64
}

type Multimedia interface {
	Mostrar()
}

type ContenidoWeb struct {
	M []Multimedia
}

func (i *Imagen) Mostrar() {
	fmt.Println("\nIMAGEN")
	fmt.Printf("Titulo: %s\n", i.titulo)
	fmt.Printf("Formato: %s\n", i.formato)
	fmt.Printf("Canales: %s\n", i.canales)
}
func (a *Audio) Mostrar() {
	fmt.Println("\nAUDIO")
	fmt.Printf("Titulo: %s\n", a.titulo)
	fmt.Printf("Formato: %s\n", a.formato)
	fmt.Printf("Duracion: %d\n", a.duracion)
}
func (v *Video) Mostrar() {
	fmt.Println("\nVIDEO")
	fmt.Printf("Titulo: %s\n", v.titulo)
	fmt.Printf("Formato: %s\n", v.formato)
	fmt.Printf("Frames: %d\n", v.frames)
}