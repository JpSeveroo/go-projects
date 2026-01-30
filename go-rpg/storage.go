package main

import(
	"fmt"
	"os"
	"encoding/json"
)

type GameState struct { // Criacao da "classe" que vai salvar novos objetos e retorna-los
	Player Player
	Mission map[string]Mission 
}

func saveGame(state GameState) { //Funcao que ta recebendo um pacote
	data, err := json.MarshalIndent(state, "", "  ") /* Essa funcionalidade de json é interessante pois passa o state 
	como um json que futuramente vai residir num data.json dentro das aspas eu passo o erro e as ultimas aspas é pra identacao*/

	if err != nil {
		fmt.Println("Erro ao salvar arquivo: ", err)
		return
	}

	err = os.WriteFile("data.json", data, 0644)
	if err != nil { //se desse certo nao retornaria algo
		fmt.Println("Erro ao salvar arquivo:", err)
		return
	}

	fmt.Println("Jogo salvo com sucesso!")
}

func loadGame() GameState{ // Nao recebe mas devolve
	var state GameState

	data, err := os.ReadFile("data.json") // Leitura do arquivo

	if err != nil { // Senao tiver nada ele imprime
		fmt.Println("Nenhum save encontrado.")
		return state
	}

	err = json.Unmarshal(data, &state) //Transforma json em struct pra imprimir no terminal

	if err != nil {
		fmt.Println("Erro ao carregar JSON:", err)
	}

	fmt.Println("Jogo carregado!")
	return state
}