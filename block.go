package main

import "fmt"

type Block struct {
	Index    int
	Data     string
	PrevHash string
	Hash     string
}

func (b Block) computeHash() string {
	payload := fmt.Sprintf("%d:%s:%s", b.Index, b.Data, b.PrevHash)
	return hash(payload)
}

func NewGenesis(data string) Block {
	b := Block{
		Index:    0,
		Data:     data,
		PrevHash: "",
	}
	b.Hash = b.computeHash()
	return b
}

func NewBlock(prev Block, data string) Block {
	b := Block{
		Index:    prev.Index + 1,
		Data:     data,
		PrevHash: prev.Hash,
	}
	b.Hash = b.computeHash()
	return b
}
