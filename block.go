package main

import "fmt"

type Block struct {
	Index    int
	Tx       Tx
	PrevHash string
	Hash     string
}

func (b Block) computeHash() string {
	payload := fmt.Sprintf("%d:%s:%s:%s", b.Index, b.Tx.payload(), b.Tx.Signature, b.PrevHash)
	return hash(payload)
}

func NewGenesis() Block {
	b := Block{
		Index:    0,
		PrevHash: "",
	}
	b.Hash = b.computeHash()
	return b
}

func NewBlock(prev Block, tx Tx) Block {
	b := Block{
		Index:    prev.Index + 1,
		Tx:       tx,
		PrevHash: prev.Hash,
	}
	b.Hash = b.computeHash()
	return b
}
