package main

import (
	"github.com/Ricxel/GoBlockchain.git/utils"
)

func main() {
	bc := utils.NewBlockchain()
	//aggiungiamo un paio di blocchi

	bc.Print()
}
