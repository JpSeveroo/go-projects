package main

import (
	"fmt"
)

func main(){
	texto := entrada()
	dadosProcessados := filtrar(texto)

	fmt.Println("\n" + "===============================")
    fmt.Println("       RELATÓRIO DE LOGS       ")
    fmt.Println("===============================")
	fmt.Printf("ERROS DETECTADOS: %d\n", dadosProcessados.logErro)
    fmt.Printf("INFOS DETECTADAS: %d\n", dadosProcessados.logInfo)
    fmt.Println("-------------------------------")

	if dadosProcessados.logErro > 0 {
		fmt.Println("DETALHES DOS ERROS:")
		for i, linha := range dadosProcessados.linhasCriticas{
			fmt.Println(i, "-->", linha, "\n")
		}
	}

	gerarGrafico(dadosProcessados)
}