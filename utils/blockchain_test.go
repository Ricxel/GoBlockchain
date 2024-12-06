package utils

import (
	"testing"
)

func TestBlockChain(t *testing.T) {
	//creiamo la blockchain
	bc := NewBlockchain()
	//aggiungiamo un paio di blocchi
	bc.InsertBlock([]byte("Primo blocco!"))
	bc.InsertBlock([]byte("Secondo blocco!"))
	t.Log("Blocchi inseriti")
	bc.Print()
	t.Log()
}
