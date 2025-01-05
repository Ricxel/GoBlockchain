package utils

import (
	"fmt"
)

type Digest []byte

type Blockchain struct {
	//l'idea è di avere un array che contiene gli hash, e una hash map che associa un hash al blocco
	//in questo modo possiamo trovare subito un blocco in base al suo hash
	UTXOs   map[string]TransactionOutput //Mappa di tutte le transazioni non spese
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
func (bc *Blockchain) InsertBlock(transactions []Transaction) { //inserimento nodo
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
	b = newBlock(prevHash, transactions)
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
func (bc *Blockchain) addTransaction(tx Transaction) error {
	// Controlla input
	for _, input := range tx.Inputs {
		utxoKey := fmt.Sprintf("%s:%d", input.PrevId, input.Output)
		utxo, exists := bc.UTXOs[utxoKey]
		if !exists { //esistenza
			return fmt.Errorf("UTXO non trovato: %s", utxoKey)
		}
		if utxo.DestinationAddress != input.Signature { //controllo firma, controlla che la tx sia effettivamente del mittente
			return fmt.Errorf("Firma non valida per l'UTXO: %s", utxoKey)
		}
		// Elimina l'UTXO speso
		delete(bc.UTXOs, utxoKey)
	}

	// Aggiungi gli output come nuovi UTXO
	// for i, output := range tx.Outputs {
	// utxoKey := fmt.Sprint	bc.Chain = append(bc.Chain, tx)
	//TODO: Aggiungere ogni transazione a un blocco
	// f("%s:%d", tx.ID, i)
	// bc.UTXOs[utxoKey] = output
	// }

	// TODO: Aggiungere la transazione a un blocco

	return nil
}

// funzione per avere il bilancio di un wallet (satoshi)
func (bc Blockchain) getBilance(address string) int64 {
	//guardo tutte le UTXO e vedo quelle che hanno l'address giusto
	sum := int64(0)
	for _, utxo := range bc.UTXOs {
		if utxo.DestinationAddress == address {
			sum += utxo.Amount
		}
	}
	return sum
}
