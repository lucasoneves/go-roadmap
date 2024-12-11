package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
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
			imprimeLogs()
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
	// var sites = []string{"https://httpbin.org/status/200", "https://www.frontendmasters.com", "https://www.caelum.com.br"}

	sites := readFileSites()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
	}

}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Error ocurred", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site", site, "foi carregado com sucesso!")
		fmt.Println("")
		registraLog(site, true)
	} else {
		fmt.Println("O site", site, "está fora do ar! =(")
		fmt.Println("Código do erro:", resp.StatusCode)
		registraLog(site, false)
	}
}

func readFileSites() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error ocurred", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 - 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Error ocurred", err)
	}

	fmt.Println(string(arquivo))

}
