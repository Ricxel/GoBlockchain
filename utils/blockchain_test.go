package utils

import (
	"testing"
)

func TestBlockChain(t *testing.T) {
	//creiamo la blockchain
	bc := newBlockchain()
	//aggiungiamo un paio di blocchi
	bc.insertBlock([]byte("Primo blocco!"))

}
