package main

import "fmt"

type Servidor struct {
	Nombre string
	IP     string
	Activo bool
}

func configurar(nombre string, ip string, canal chan *Servidor) {
	s := Servidor{
		Nombre: nombre,
		IP:     ip,
		Activo: true,
	}

	canal <- &s
}

func main() {
	canal := make(chan *Servidor)

	go configurar("Servidor 1", "192.168.1.1", canal)
	go configurar("Servidor 2", "192.168.1.2", canal)

	for i := 0; i < 2; i++ {
		s := <-canal

		fmt.Println("Nombre:", s.Nombre)
		fmt.Println("IP:", s.IP)
		fmt.Println("Activo:", s.Activo)
		fmt.Println("----------------")
	}
}
