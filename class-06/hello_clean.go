package main

import (
	"fmt"
	"os"
	"net/http"
	"time"
	"bufio"
	"strings"
	"io"
)

const monitoramentos = 2
const delay = 5

func main() {

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
}

func exibeMenu() {
	fmt.Println("1-Iniciar Monitoramente")
	fmt.Println("2-Exibir logs")
	fmt.Println("0-Sair do programa")
}

func leComando() int  {
	var comando int
	fmt.Scan(&comando)
	return comando;
}

func inicarMonitoramente() {
	fmt.Println("Iniciando Monitoramente...")
	sites := leSiteDoArquivo()

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

func leSiteDoArquivo() []string {
	var sites []string
	arquivo2, err2 := os.Open("sites.txt")

	if err2 != nil {
		fmt.Println("Ocorreu um erro 2: ", err2)

		return sites
	}

	leitor := bufio.NewReader(arquivo2)

	for {
		linha, err3 := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		if err3 == io.EOF {
			break
		}

		if err3 != nil {
			fmt.Println("Ocorreu um erro 3: ", err3)
			return sites
		}
	
		fmt.Println(linha)
		sites = append(sites, linha)
	}

	arquivo2.Close()

	return sites
}