package main

import "fmt"

var (
	player   Player
	missions = make(map[string]Mission)
)

func main() {
	state := loadGame()
	player = state.Player

	if state.Mission != nil {
		missions = state.Mission
	}

	for {
		fmt.Println("\n1 - Criar jogador")
		fmt.Println("2 - Criar missão")
		fmt.Println("3 - Listar missões")
		fmt.Println("4 - Salvar")
		fmt.Println("5 - Carregar")
		fmt.Println("0 - Sair")

		fmt.Print("Escolha: ")
		op := readInt()

		switch op {
		case 1:
			player = createPlayer()

		case 2:
			m := createMission()
			missions[m.ID] = m

		case 3:
			listMissions()

		case 4:
			saveGame(GameState{Player: player, Mission: missions})

		case 5:
			state := loadGame()
			player = state.Player
			if state.Mission != nil {
				missions = state.Mission
			}

		case 0:
			fmt.Println("Saindo...")
			return

		default:
			fmt.Println("Opção inválida")
		}
	}
}