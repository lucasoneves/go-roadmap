package main

import "fmt"
import "os"
import "net/http"

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
	site := "https://httpbin.org/status/200"

	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("O site", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("O site", site, "está fora do ar! =(")
		fmt.Println("Código do erro:", resp.StatusCode)
	}
}