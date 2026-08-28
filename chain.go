package main

import "fmt"

type Chain []Block

func NewChain(genesisData string) Chain {
	return Chain{NewGenesis(genesisData)}
}

func (c Chain) Add(data string) Chain {
	prev := c[len(c)-1]
	return append(c, NewBlock(prev, data))
}

func (c Chain) IsValid() error {
	if len(c) == 0 {
		return fmt.Errorf("empty chain")
	}

	genesis := c[0]
	if genesis.Index != 0 || genesis.PrevHash != "" {
		return fmt.Errorf("block 0 is not a valid genesis")
	}
	if genesis.Hash != genesis.computeHash() {
		return fmt.Errorf("block 0: stored hash does not match contents")
	}

	for i := 1; i < len(c); i++ {
		prev := c[i-1]
		curr := c[i]

		if curr.Index != prev.Index+1 {
			return fmt.Errorf("block %d: index should be %d", i, prev.Index+1)
		}
		if curr.Hash != curr.computeHash() {
			return fmt.Errorf("block %d: stored hash does not match contents", i)
		}
		if curr.PrevHash != prev.Hash {
			return fmt.Errorf("block %d: prev_hash does not match previous block", i)
		}
	}
	return nil
}
