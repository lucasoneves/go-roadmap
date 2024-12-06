package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	showIntro()

	for {
		showMenu()

		comando := readCommand()

		switch comando {
		case 1:
			initMonotiring()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)

		default:
			fmt.Println("Não conheço esse comando...")
			os.Exit(-1)
		}
	}

}

func showIntro() {
	userName := "Lucas"
	version := 1.1
	fmt.Println("Olá, sr.", userName)
	fmt.Println("Este programa está na versão", version)
}

func showMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func readCommand() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("Comando escolhido:", comando)

	return comando
}

func initMonotiring() {
	fmt.Println("Monitorando...")
	var sites = []string{"https://httpbin.org/status/200", "https://www.frontendmasters.com", "https://www.caelum.com.br"}

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
	}

}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("O site", site, "foi carregado com sucesso!")
		fmt.Println("")
	} else {
		fmt.Println("O site", site, "está fora do ar! =(")
		fmt.Println("Código do erro:", resp.StatusCode)
	}
}
