package main

import "fmt"
import "os"

func main() {

	showIntro()
	showMenu()	

	comando := readCommand()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
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
