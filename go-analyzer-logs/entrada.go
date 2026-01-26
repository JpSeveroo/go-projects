package main

import (
	"fmt"
	"time"
	"os"
	"bufio"
)

func entrada() string{
	hello := "\n===== WELCOME TO LOGS ANALYZER ====="

	for _, caractereH := range hello{
		fmt.Printf("%c", caractereH)
		time.Sleep(50 * time.Millisecond)
	}
	
	como := "\nComo funciona >>> Esta ferramenta analisa logs brutos para \nextrair estatísticas de performance e erros."
	for _, caractereC := range como {
		fmt.Printf("%c", caractereC)
		time.Sleep(50 * time.Millisecond)
	}
	
	descri := "\n\nCole o conteúdo do seu log abaixo e pressione ENTER seguido \nde Ctrl+D para iniciar o processamento: "
	for _, caractereD := range descri {
		fmt.Printf("%c", caractereD)
		time.Sleep(50 * time.Millisecond)
	}

	fmt.Println("\n")

	scanner := bufio.NewScanner(os.Stdin)
	var logCapturado string

	return logCapturado()
	
}