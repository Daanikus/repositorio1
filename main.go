package main

import "fmt"

const helloWorld string ="Hola %s %s, bienvenido a GO "
const testConst = "Test"

func main() {

	name := getName()
	lastname:="<Modificar con el apellido>"
	var number = 100

	a,b,c := getVariables()

	fmt.Printf(helloWorld,name,lastname)
	fmt.Println("hola mundo ")
	fmt.Println(number,a,b,c)
	//Comentariooooossss
	//Comentario- Se reinicia capacitacion GO
}

func getName() string {
	var name string
	name ="Sin nombres"
	fmt.Print("Ingresa tu nombre")
	fmt.Scanf("%s",&name)
	return name

}

func getVariables()(int ,int,int)  {
	return 4,5,6

}
