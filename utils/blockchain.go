package utils

import "fmt"

type Digest []byte

type Blockchain struct {
	//l'idea è di avere un array che contiene gli hash, e una hash map che associa un hash al blocco
	//in questo modo possiamo trovare subito un blocco in base al suo hash
	Digests []Digest
	Blocks  map[string]*Block //uso string e non Digest perchè []byte non supporta direttamente l'operatore di eguaglianza ==
}

func NewBlockchain() *Blockchain {
	bc := Blockchain{
		Digests: []Digest{},
		Blocks:  make(map[string]*Block), //inizializzazione della map
	}
	bc.GenerateGenesisBlock()
	return &bc
}
func (bc *Blockchain) insertBlock(b *Block) {
	//se è valido lo mettiamo dentro
	if !b.verify() {
		return
	}
	fmt.Println("Blocco valido, inserito!")
	bc.Digests = append(bc.Digests, b.Hash)
	bc.Blocks[string(b.Hash)] = b
}
func (bc *Blockchain) InsertBlock(data []byte) { //inserimento nodo
	//steps:
	//1) Mettere nel nodo current l'hash del nodo precedente e i dati passati
	//2) Aggiungerlo l'hash al'array
	//3) Aggiungere il blocco alla map
	var b *Block
	prevHash := []byte{} //inizializzo a array vuoto, se è il primo blocco andrà bene

	if len(bc.Blocks) != 0 {
		//0) Troviamo il prev hash
		prevHash = bc.Digests[len(bc.Digests)-1]
	}
	//1) Creiamo il blocco
	b = newBlock(prevHash, data)
	//Ora aggiungiamolo alla block chain
	bc.Digests = append(bc.Digests, b.Hash)
	bc.Blocks[string(b.Hash)] = b
}
func (bc Blockchain) GetBlock(digest []byte) *Block { //ottieni nodo
	return bc.Blocks[string(digest)]
}
func (bc *Blockchain) GetLastHash() []byte {
	if len(bc.Digests) == 0 {
		return []byte{} //non ci sono blocchi
	}
	return bc.Digests[len(bc.Digests)-1]
}
func (bc *Blockchain) GenerateGenesisBlock() {
	block := Block{
		PrevHash:  []byte{},
		Data:      []byte("Genesis Block"),
		TimeStamp: 0,
		Nonce:     0,
		Target:    0,
	}
	block.Hash = block.CalculateHash()
	bc.Digests = append(bc.Digests, block.Hash)
	bc.Blocks[string(block.Hash)] = &block
}
func (bc Blockchain) Print() {
	for i := 0; i < len(bc.Digests); i++ {
		digest := bc.Digests[i]
		bc.GetBlock(digest).Print()
		if i != len(bc.Digests)-1 {
			fmt.Print(" <- ")
		}
	}
	fmt.Println()
}
