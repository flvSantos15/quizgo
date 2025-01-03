package main

import (
	"bufio"
	"fmt"
	"os"
)

type Question struct {
	Text string
	Options []string
	Answer int
}

type GameState struct {
	Name string
	Points string
	Questions []Question
}

func (g *GameState) Int() {
	fmt.Println("Seja bem vindo(a) ao quiz")
	fmt.Println("Escreva o seu nome: ")
	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n')

	if err != nil {
		panic("Error ao ler a string")
	}

	g.Name = name

	fmt.Printf("Vamos ao jogo %s", g.Name)
}

func main() {
	game := &GameState{}
	game.Int()
}