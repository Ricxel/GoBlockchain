package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type TransactionInput struct {
	PrevId    string //ID della transazione precedente
	Output    int
	Signature string //firma del mittente
}
type TransactionOutput struct {
	Amount             int64
	DestinationAddress string
}

type Transaction struct {
	ID      string              // ID della transazione
	Inputs  []TransactionInput  // Input
	Outputs []TransactionOutput // Output
}

func calculateTransactionID(tx Transaction) string { //hash di tutta la transazione
	hash := sha256.New()
	data := fmt.Sprintf("%v%v", tx.Inputs, tx.Outputs)
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
func newTransaction(inputs []TransactionInput, outputs []TransactionOutput) Transaction {
	tx := Transaction{
		Inputs:  inputs,
		Outputs: outputs,
	}
	tx.ID = calculateTransactionID(tx)
	return tx
}
