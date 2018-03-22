package main

import "fmt"

var resultado int

func main() {
	resultado := funcionCalculadora()
		fmt.Printf("El resultado es : %d ",resultado)
}

func funcionCalculadora() int  {
	var numero1 int
	var numero2 int
	var operacion int
	var resultadoCalculadora int

	fmt.Print("Ingresa primer numero: ")
	fmt.Scanf("%d",&numero1)
	fmt.Print("Ingresa segundo numero: ")
	fmt.Scanf("%d",&numero2)
	fmt.Print("Ingresa operacion a realizar: \n1:Suma - 2:Resta - 3:Multiplicacion - 4:Division")
	fmt.Scanf("%d",&operacion)

  resultadoCalculadora = gestorOperacion(operacion, numero1,numero2)
	return resultadoCalculadora
}

func gestorOperacion(operacion int, numero1 int, numero2 int) int {
	var operacionTerminada int
	if (operacion == 1) {
		operacionTerminada = numero1 + numero2
	}else if (operacion == 2)  {
		operacionTerminada = numero1 - numero2
	}else if (operacion == 3) {
		operacionTerminada =  numero1 * numero2
	}else if (operacion == 4) {
		operacionTerminada = numero1 / numero2
	}else  {
		fmt.Print("Operacion no permitida \n")
	}
	return operacionTerminada
}
