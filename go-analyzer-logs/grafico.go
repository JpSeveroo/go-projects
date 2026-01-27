package main

import (
	"fmt"
	"strings"
)

func gerarGrafico(dadosProcessados logResultado){
	fmt.Println(">>>>> Deseja gerar um gráfico visual referente aos eventos decorridos? ")
	fmt.Println("[ 0 ] Não")
	fmt.Println("[ 1 ] Sim")
	fmt.Print("Digite a operação desejada a partir de um algarismo: ")

	var escolha int

	for {
		fmt.Scan(&escolha)

		if escolha == 0 || escolha == 1 {
			if escolha == 1 {
				fmt.Println("\n===============================")
				fmt.Println("      REPRESENTAÇÃO VISUAL     ")
				fmt.Println("===============================")

				total := dadosProcessados.logErro + dadosProcessados.logInfo

				if total == 0 {
					fmt.Println("Sem dados suficientes para gerar o gráfico.")
					return
				}

				const larguraMax = 20

				tamanhoErro := (dadosProcessados.logErro * larguraMax) / total
				tamanhoInfo := (dadosProcessados.logInfo * larguraMax) / total


				barraErro := strings.Repeat("█", tamanhoErro)
				barraInfo := strings.Repeat("█", tamanhoInfo)

				fmt.Printf("ERROS: [%-20s] %d\n", barraErro, dadosProcessados.logErro)
				fmt.Printf("INFOS: [%-20s] %d\n", barraInfo, dadosProcessados.logInfo)
				fmt.Println("===============================")
			}
			break
		}

		fmt.Println("[!] Valor inválido, tente novamente.")
	}
}