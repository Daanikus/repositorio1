package main

import "fmt"

const helloWorld string ="Hola %s %s, bienvenido a GO "
const testConst = "Test"

func main() {
	var name string
	name = "sin nombre"
	lastname:="<Modificar con el apellido>"
	var number = 100
	var (
		a=1
		b=2
		c=3
	)

	fmt.Println("ingresa tu nombre: ")
	fmt.Scanf("%s",&name)
	fmt.Printf(helloWorld,name,lastname)
	fmt.Println("hola mundo ")
	fmt.Println(number,a,b,c)
	//Comentariooooo
	//Comentario- Se reinicia capacitacion GO
}
