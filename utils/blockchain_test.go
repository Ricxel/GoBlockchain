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

func TestMine(t *testing.T) {
	//creiamo la blockchain
	bc := NewBlockchain()
	//aggiungo un blocco
	prevHash := bc.GetLastHash()
	block := newBlock(prevHash, []byte("Primo blocco da minare!"))

	block.Mine()
	t.Log("Blocco minato")
	//lo inseriamo nella blockchain
	bc.insertBlock(block)
	t.Log("Blocchi inseriti")
	bc.Print()
	t.Log()
}
