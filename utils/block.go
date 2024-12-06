package utils

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"time"
)

type Block struct {
	PrevHash  []byte
	Hash      []byte
	Data      []byte //potrei fare una struttura apposita per i dati (ad esempio per salvare delle transazioni nel caso di una criptovaluta)
	TimeStamp int64
}

// metodo per calcolare l'hash
func (b *Block) CalculateHash() {
	//bisogna calcolare l'hash del prevHash+Data+Timestamp e inserirlo nell'hash del current block
	hash := sha256.New()
	hash.Write([]byte(b.PrevHash))
	hash.Write(b.Data)
	//ora per il timestamp bisogna fare un macello perchè questo linguaggio fa caà
	buf := make([]byte, 8)                               //prima creo un buffer
	binary.BigEndian.PutUint64(buf, uint64(b.TimeStamp)) // gli metto dentro il timestamp in bigendian

	hash.Write(buf) //ora lo posso scrivere :)

	digest := hash.Sum(nil) //calcolo l'hash
	b.Hash = digest[:]
}
func newBlock(prevHash []byte, data []byte) *Block {
	b := Block{ //faccio il blocco
		PrevHash:  prevHash,
		Data:      data,
		TimeStamp: time.Now().Unix(),
	}
	b.CalculateHash()
	return &b //ritorno l'indirizzo del blocco
}
func (b Block) Print() {
	fmt.Print("{" + string(b.PrevHash) + ", " + string(b.Data) + ", " + string(b.Hash) + "}")
}
