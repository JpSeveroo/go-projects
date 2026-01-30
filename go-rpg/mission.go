package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Mission struct {
	ID          string
	Nome        string
	Dificuldade int
	Recompensa  int
}

var reader = bufio.NewReader(os.Stdin)

func readLine() string {
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func readInt() int {
	for {
		text := readLine()
		n, err := strconv.Atoi(text)
		if err == nil {
			return n
		}
		fmt.Print("Digite um número válido: ")
	}
}

func createMission() Mission {
	var m Mission

	fmt.Print("ID: ")
	m.ID = readLine()

	fmt.Print("Nome: ")
	m.Nome = readLine()

	fmt.Print("Dificuldade: ")
	m.Dificuldade = readInt()

	fmt.Print("Recompensa: ")
	m.Recompensa = readInt()

	return m
}

func listMissions() {
	fmt.Println("\n--- MISSÕES ---")

	if len(missions) == 0 {
		fmt.Println("Nenhuma missão cadastrada.")
		return
	}

	for _, m := range missions {
		fmt.Printf(
			"ID: %s | Nome: %s | Dif: %d | Ouro: %d\n",
			m.ID, m.Nome, m.Dificuldade, m.Recompensa,
		)
	}
}