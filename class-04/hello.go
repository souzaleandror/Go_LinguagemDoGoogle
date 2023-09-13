package main

import (
	"fmt"
	"reflect"
	"os"
	"net/http"
)

func main() {
	
	nome, idade := devolveNomeEIdade()
	fmt.Println("Seu nome e ", nome, " e idade e ", idade)

	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			fmt.Println("Iniciar Monitoramente")
			inicarMonitoramente()
		case 2:fmt.Println("Exibir logs")
		case 0:
			fmt.Println("Sair do programa")
			os.Exit(0)
		default:
			fmt.Println("Nao conheco este comando")
			os.Exit(-1)
	}
	}


	// if comando == 1 {
	// 	fmt.Println("Iniciar Monitoramente")
	// } else if comando == 2 {
	// 	fmt.Println("Exibir logs")
	// } else if comando == 0 {
	// 	fmt.Println("Sair do programa")
	// 	os.Exit(0)
	// } else {
	// 	fmt.Println("Nao conheco este comando")
	// 	os.Exit(-1)
	// }


}

func exibeIntroducao() {
	nome := "Douglas"
	idade := 24
	versao := 1.1

	fmt.Println("Ola, Sr.", nome, "sua idade e ", idade)
	fmt.Println("Ese programa esta na versao", versao)

	fmt.Println("O tipo da variavel nome e ", reflect.TypeOf(nome))
	fmt.Println("O tipo da variavel idade e ", reflect.TypeOf(idade))
	fmt.Println("O tipo da variavel versao e ", reflect.TypeOf(versao))
}

func exibeMenu() {
	fmt.Println("1-Iniciar Monitoramente")
	fmt.Println("2-Exibir logs")
	fmt.Println("0-Sair do programa")
}

func leComando() int  {
	var comando int

	//fmt.Scanf("%d", &comando)
	fmt.Scan(&comando)
	fmt.Println("Endereco da variavel/ponteiro", &comando)
	fmt.Println("O comando escolhido foi ", comando)
	return comando;
}

func inicarMonitoramente() {
	fmt.Println("Iniciando Monitoramente...")
	site := "https://www.alura.com.br/"
	resp, _ := http.Get(site)
	fmt.Println(resp)

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, " Foi carregado com sucesso")
	} else {
		fmt.Println("Site: ", site, " esta com problemas, com error", resp.StatusCode)
	}
}

func devolveNomeEIdade() (string, int){
	nome := "Dougles"
	idade := 24

	return nome,  idade
}