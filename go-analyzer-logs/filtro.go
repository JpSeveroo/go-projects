package main

import (
	"strings"
)

type logResultado struct{
	logInfo int
	logErro int
	linhasCriticas []string
}

func filtrar(textoBruto string) logResultado{
	var res logResultado

	linhas := strings.Split(textoBruto, "\n") // aqui eu organizo meu texto em relação a quebra de linhas

	for _, linha := range linhas{ // Faço um for pra percorrer todas as linhas 
		linhaMaiuscula := strings.ToUpper(linha) // Deixo maiúsculo para identificar palavras chaves

		// Abaixo eu procuro erros
		if strings.Contains(linhaMaiuscula, "ERROR"){
			res.logErro ++
			res.linhasCriticas = append(res.linhasCriticas, linha)
		}

		if strings.Contains(linhaMaiuscula, "INFO"){
			res.logInfo ++
		}
	}

	return res
}