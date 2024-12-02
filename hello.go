package main

import "fmt"

func main() {
	userName := "Lucas"
	version := 1.1
	fmt.Println("Olá, sr.", userName, version)
	fmt.Println("Este programa está na versão", version)

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("Comando escolhido:", comando)

	if comando == 0 {
		fmt.Println("Saindo do programa...")
	} else if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo logs...")
	} else if comando == 0 {
		fmt.Println("Saindo do programa.")
	} else {
		fmt.Println("Não conheço esse código")
	}
}
