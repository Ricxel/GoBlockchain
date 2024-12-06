package utils

import (
	"testing"
)

func TestCreateBlock(t *testing.T) {
	block := Block{
		PrevHash:  []byte("È il primo blocco quindi no hash :("),
		Data:      []byte("Prova di dato"),
		TimeStamp: 123982303,
	}
	block.CalculateHash()
	t.Log("Hash: " + string(block.Hash))
}
