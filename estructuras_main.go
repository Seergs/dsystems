package main

/*
func main() {
	shouldExit := false
	cw := ContenidoWeb{}
	for !shouldExit {
		displayMenu()
		op := getIntFromUser()
		switch op {
		case 1:
			displayAddMenu()
			contenido := getIntFromUser()
			if contenido == 1 {
				cw.M = append(cw.M, addNewImage())
			} else if contenido == 2{
				cw.M = append(cw.M, addNewAudio())
			} else {
				cw.M = append(cw.M, addNewVideo())
			}
			fmt.Println("Contenido agregado!\n\n")
			break
		case 2:
			mostrar(cw.M...)
			break
		case 3:
			shouldExit = true
			break
		default:
			fmt.Print("Opcion invalida")
			break
		}
	}
}

func displayMenu() {
	fmt.Printf("\nContenido Web\n")
	fmt.Print("1. Agregar nuevo\n")
	fmt.Print("2. Mostrar\n")
	fmt.Print("3. Salir\n")
	fmt.Print("Opcion: ")
}

func displayAddMenu() {
	fmt.Print("\t1. Imagen\n")
	fmt.Print("\t2. Audio\n")
	fmt.Print("\t3. Video\n")
	fmt.Print("\tOpcion: ")
}

func getIntFromUser() int64 {
	var op int64
	fmt.Scan(&op)

	return op
}

func addNewImage() *Imagen {
	var titulo, formato, canales string
	fmt.Printf("\nCrea una nueva imagen\n")
	fmt.Print("Titulo: ")
	titulo = getStringFromUser()
	fmt.Print("Formato: ")
	formato = getStringFromUser()
	fmt.Print("Canales: ")
	canales = getStringFromUser()

	return &Imagen{titulo, formato, canales}
}

func addNewAudio() *Audio {
	var titulo, formato string
	var duracion int64

	fmt.Printf("\nCrea un audio\n")
	fmt.Print("Titulo: ")
	titulo = getStringFromUser()
	fmt.Print("Formato: ")
	formato = getStringFromUser()
	fmt.Print("Duracion: ")
	duracion = getIntFromUser()

	return &Audio{titulo, formato, duracion}
}

func addNewVideo() *Video {
	var titulo, formato string
	var frames int64

	fmt.Printf("\nCrea un video\n")
	fmt.Print("Titulo: ")
	titulo = getStringFromUser()
	fmt.Print("Formato: ")
	formato = getStringFromUser()
	fmt.Print("Frames: ")
	frames = getIntFromUser()

	return &Video{titulo, formato, frames}
}

func mostrar(contenido ...Multimedia) {
	for _, c := range contenido {
		c.Mostrar()
	}
}

func getStringFromUser() string {
	var line string

	fmt.Scan(&line)

	return line
}
*/