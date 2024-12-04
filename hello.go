package main

import "fmt"

func main() {
	userName := "Lucas"
	version := 1.1
	var comando int

	fmt.Println("Olá, sr.", userName)
	fmt.Println("Este programa está na versão", version)

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")

	fmt.Scan(&comando)

	fmt.Println("Comando escolhido:", comando)

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa...")
	default:
		fmt.Println("Não conheço esse comando...")
	}
}
