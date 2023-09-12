package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Douglas"
	idade := 24
	versao := 1.1

	fmt.Println("Ola, Sr.", nome, "sua idade e ", idade)
	fmt.Println("Ese programa esta na versao", versao)
	fmt.Println("O tipo da variavel nome e ", reflect.TypeOf(nome))
	fmt.Println("O tipo da variavel idade e ", reflect.TypeOf(idade))
	fmt.Println("O tipo da variavel versao e ", reflect.TypeOf(versao))

	fmt.Println("1-Iniciar Monitoramente")
	fmt.Println("2-Exibir logs")
	fmt.Println("0-Sair do programa")

	var comando int

	//fmt.Scanf("%d", &comando)
	fmt.Scan(&comando)
	fmt.Println("Endereco da variavel/ponteiro", &comando)
	fmt.Println("O comando escolhido foi ", comando)

}
