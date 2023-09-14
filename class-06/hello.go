package main

import (
	"fmt"
	"reflect"
	"os"
	"net/http"
	"time"
	"bufio"
)

const monitoramentos = 2
const delay = 5

func main() {
	
	nome, idade := devolveNomeEIdade()
	fmt.Println("Seu nome e ", nome, " e idade e ", idade)

	exibeIntroducao()
	//leSiteDoArquivo()

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

	// site := "https://www.alura.com.br/"
	// resp, _ := http.Get(site)
	// fmt.Println(resp)

	// if resp.StatusCode == 200 {
	// 	fmt.Println("Site: ", site, " Foi carregado com sucesso")
	// } else {
	// 	fmt.Println("Site: ", site, " esta com problemas, com error", resp.StatusCode)
	// }

		// var sites [4]string
	// sites[0] = "https://www.alura.com.br/"
	// sites[1] = "https://www.caelum.com.br/"
	// sites[2] = "https://www.google.at/?client=safari"

	//sites := []string{"https://www.alura.com.br/", "https://www.caelum.com.br/", "https://www.google.at/?client=safari"}

	sites := leSiteDoArquivo()

	fmt.Println(sites)
	fmt.Println(reflect.TypeOf(sites))
	fmt.Println("O meu array tem ", len(sites), " posicoes")
	fmt.Println("O meu array tem ", cap(sites), " capacidade")

	// for i:=0; i < len(sites); i++ {
	// 	site := sites[i]
	// 	resp, _ := http.Get(site)
	// 	fmt.Println(resp)

	// 	if resp.StatusCode == 200 {
	// 		fmt.Println("Site: ", site, " Foi carregado com sucesso")
	// 	} else {
	// 		fmt.Println("Site: ", site, " esta com problemas, com error => ", resp.StatusCode)
	// 	}
	// }

	// for i, link := range sites {
	// 	site := link
	// 	resp, _ := http.Get(site)
	// 	fmt.Println(resp)

	// 	if resp.StatusCode == 200 {
	// 		fmt.Println("Site: ", site, " Foi carregado com sucesso")
	// 	} else {
	// 		fmt.Println("Site: ", site, " esta com problemas, com error => ", resp.StatusCode)
	// 	}
	// }

	for j:=0; j < monitoramentos; j++ {
		for i, link := range sites {
			fmt.Println("Testando site : ", i, " => ", link)
			testSite(link)
			fmt.Println("")
		}
		time.Sleep(delay * time.Second)
	}

	fmt.Println("")
	fmt.Println("")
}

func testSite(site string) {
	resp, err := http.Get(site)
	fmt.Println(resp)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		return
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, " Foi carregado com sucesso")
	} else {
		fmt.Println("Site: ", site, " esta com problemas, com error => ", resp.StatusCode)
	}
}

func devolveNomeEIdade() (string, int){
	nome := "Dougles"
	idade := 24

	return nome,  idade
}

func exibeNomes() {
	nomes := []string{"test1", "test2", "test3"}
	nomes = append(nomes, "Aparecida")
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nome))
	fmt.Println("O meu slice tem ", len(nomes), " posicoes")
	fmt.Println("O meu slice tem ", cap(nomes), " capacidade")
	
	nomes = append(nomes, "Aparecidad")
	fmt.Println("O meu slice tem ", len(nomes), " posicoes")
	fmt.Println("O meu slice tem ", cap(nomes), " capacidade")
}

func leSiteDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	arquivo2, err2 := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)

		return
	}

	if err2 != nil {
		fmt.Println("Ocorreu um erro 2: ", err2)

		return
	}

	fmt.Println(arquivo)
	fmt.Println(string(arquivo2))

	leitor := bufio.NewReader(arquivo2)

	for {
		linha, err3 := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		if err3 == io.EOF {
			break
		}

		if err3 != nil {
			fmt.Println("Ocorreu um erro 3: ", err3)
			return
		}
	
		fmt.Println(linha)
		sites = append(sites, linha)
	}

	arquivo.Close()
	arquivo2.Close()

	fmt.Println(sites)

	return sites
}