package main

import (
	"fmt"
)

type Player struct {
	Nick   string
	Classe string
	Nivel  int
}

func createPlayer() Player {
	var p Player
	
	fmt.Print("Nickname: ")
	p.Nick = readLine()

	fmt.Print("Classe: ")
	p.Classe = readLine()

	fmt.Print("Nivel: ")
	p.Nivel = readInt()

	return p
}