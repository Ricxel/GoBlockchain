package utils

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	PrevHash     []byte
	Transactions []Transaction //insieme di transazioni
	Hash         []byte
	TimeStamp    int64
}

// metodo per calcolare l'hash
func (b *Block) CalculateHash() {
	//bisogna calcolare l'hash del prevHash+Data+Timestamp e inserirlo nell'hash del current block
	hash := sha256.New()
	hash.Write([]byte(b.PrevHash))
	data := fmt.Sprintf("%s%v", b.TimeStamp, b.Transactions)
	hash.Write([]byte(data))

	digest := hash.Sum(nil) //calcolo l'hash
	b.Hash = digest[:]
}
func newBlock(prevHash []byte, transactions []Transaction) *Block {
	b := Block{ //faccio il blocco
		PrevHash:     prevHash,
		Transactions: transactions,
		TimeStamp:    time.Now().Unix(),
	}
	b.CalculateHash()
	return &b //ritorno l'indirizzo del blocco
}
func (b Block) Print() {
	fmt.Print("{" + string(b.PrevHash) + ", " + string(b.Hash) + "}")
}
