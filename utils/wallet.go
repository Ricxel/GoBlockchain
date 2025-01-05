package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
)

type Wallet struct {
	Address string
	pubK    []byte
	privK   []byte
}

// Funzione per generare una chiave privata
func generatePrivateKey() (*ecdsa.PrivateKey, error) {
	curve := elliptic.P256() // Curve utilizzata in Bitcoin è secp256k1, ma P256 è simile.
	privKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}
	return privKey, nil
}

// Funzione per ottenere la chiave pubblica dalla chiave privata
func getPublicKey(privKey *ecdsa.PrivateKey) []byte {
	// La chiave pubblica è il punto sulla curva generato dalla chiave privata
	pubKey := privKey.PublicKey
	x, y := pubKey.X, pubKey.Y

	// La chiave pubblica è la concatenazione dei componenti X e Y
	pubKeyBytes := append(x.Bytes(), y.Bytes()...)
	return pubKeyBytes
}

// Funzione per generare l'indirizzo Bitcoin, è lo sha256 della firma pubblica
func createAddress(pubKey []byte) string {
	// SHA256 della chiave pubblica
	hash := sha256.New()
	hash.Write(pubKey)
	pubKeyHash := hash.Sum(nil)
	addressHash := pubKeyHash

	return hex.EncodeToString(addressHash)
}

// Funzione per creare un Wallet Bitcoin
func createWallet() {
	privKey, err := generatePrivateKey()
	if err != nil {
		log.Fatalf("Errore nella generazione della chiave privata: %v", err)
	}

	pubKey := getPublicKey(privKey)
	address := createAddress(pubKey)

	// Convertiamo la chiave privata in formato esadecimale
	privKeyHex := fmt.Sprintf("%x", privKey.D.Bytes())

	// Stampa la chiave privata, la chiave pubblica e l'indirizzo
	fmt.Println("Chiave Privata (Hex):", privKeyHex)
	fmt.Println("Indirizzo Bitcoin:", address)
}
