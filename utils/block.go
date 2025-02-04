package utils

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"time"
)

type Block struct {
	PrevHash     []byte
	Transactions []Transaction //insieme di transazioni
	Hash         []byte
	TimeStamp    int64
	Nonce        uint64
	Target       uint16
}

var ZeroNumber = 18

// metodo per calcolare l'hash
func (b *Block) CalculateHash() []byte {
	//bisogna calcolare l'hash del prevHash+Data+Timestamp e inserirlo nell'hash del current block
	hash := sha256.New()
	hash.Write([]byte(b.PrevHash))
	data := fmt.Sprintf("%s%v", b.TimeStamp, b.Transactions)
	hash.Write([]byte(data))

	buf2 := make([]byte, 8)
	binary.BigEndian.PutUint64(buf2, uint64(b.Nonce))

	hash.Write(buf2) //aggiungo anche il nonce

	digest := hash.Sum(nil) //calcolo l'hash
	return digest
}
func CheckFirstNBitsZero(data []byte, target uint16) bool {
	//iteriamo sui bits da 0 a n-1
	for i := uint16(0); i < target; i++ {
		byteIndex := i / 8 //indice di byte
		bitIndex := i % 8  //bit specifico
		if (data[byteIndex] & (1 << (7 - bitIndex))) != 0 {
			return false
		}
	}
	return true
}
func (b *Block) Mine() {
	//faccio un []byte con "n" byte messi a 0
	compArray := []byte{}
	for i := 0; i < ZeroNumber; i++ {
		compArray = append(compArray, 0)
	}
	print(compArray)
	var digest []byte
	for {
		digest = b.CalculateHash()
		fmt.Println(b.Nonce)
		fmt.Println(string(digest))
		//li confronto
		if CheckFirstNBitsZero(digest, b.Target) {
			break //hash trovato
		}
		b.Nonce++
	}
	//ora lo metto nel blocco
	b.Hash = digest
}
func newBlock(prevHash []byte, transactions []Transaction) *Block {
	b := Block{ //faccio il blocco
		PrevHash:     prevHash,
		Transactions: transactions,
		TimeStamp:    time.Now().Unix(),
		Nonce:        0, //numero per avere i primi n valori dell'hash a 0
		Target:       uint16(ZeroNumber),
	}
	return &b //ritorno l'indirizzo del blocco
}
func (b Block) verify() bool { //funziona che controlla se il nonce trovato è giusto calcolando l'hash e verficandolo
	digest := b.CalculateHash()
	//controlliamo se ha i primi 'target' bit a 0
	return CheckFirstNBitsZero(digest, b.Target)
}
func (b Block) Print() {
	fmt.Print("{" + string(b.PrevHash) + ", " + string(b.Hash) + "}")
}
