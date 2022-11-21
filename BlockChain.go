package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	timestamp   time.Time
	Klassenbuch []string
	prevHash    []byte
	Hash        []byte
}

func main() {
	genesisKlassenbuch := []string{"Michel hat am 21.11.2022 angefangen mit dem Code", "Michel ist erst in die Falsche Richtung gelaufen"}
	genesisBlock := NewBlock(genesisKlassenbuch, []byte{})
	fmt.Println("--- First Block ---")
	printBlockInformation(genesisBlock)

	block2Klassenbuch := []string{"Michel hat das erste erfolgs ergebniss mit dem Code", "Die BlockChain hat erst nicht richtig funktioniert wegen dem Virusprogramm"}
	block2 := NewBlock(block2Klassenbuch, genesisBlock.Hash)
	fmt.Println("--- Second Block ---")
	printBlockInformation(block2)

	block3Klassenbuch := []string{"Michel zerbricht sich den Kopf über die Readme datei", "Michel hat einen Tipp von Christian bekommen"}
	block3 := NewBlock(block3Klassenbuch, block2.Hash)
	fmt.Println("--- Third Block ---")
	printBlockInformation(block3)
}

func NewBlock(Klassenbuch []string, prevHash []byte) *Block {
	currentTime := time.Now()
	return &Block{
		timestamp:   currentTime,
		Klassenbuch: Klassenbuch,
		prevHash:    prevHash,
		Hash:        NewHash(currentTime, Klassenbuch, prevHash),
	}
}

func NewHash(time time.Time, klassenbuch []string, prevHash []byte) []byte {
	input := append(prevHash, time.String()...)
	for Klassenbuch := range klassenbuch {
		input = append(input, string(rune(Klassenbuch))...)
	}
	hash := sha256.Sum256(input)
	return hash[:]
}

func printBlockInformation(block *Block) {
	fmt.Printf("\ttime: %s\n", block.timestamp.String())
	fmt.Printf("\tprevHash: %x\n", block.prevHash)
	fmt.Printf("\tHash: %x\n", block.Hash)
	printTransactions(block)
}

func printTransactions(block *Block) {
	fmt.Println("\tTransactions:")
	for i, klassenbuch := range block.Klassenbuch {
		fmt.Printf("\t\t%v: %q\n", i, klassenbuch)
	}
}
