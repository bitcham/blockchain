package main

import "fmt"

type Chain []Block

func NewChain() Chain {
	return Chain{NewGenesis()}
}

func (c Chain) Add(tx Tx) Chain {
	prev := c[len(c)-1]
	return append(c, NewBlock(prev, tx))
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
	if genesis.Tx != (Tx{}) {
		return fmt.Errorf("block 0 must not contain a transaction")
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
		if err := curr.Tx.IsValid(); err != nil {
			return fmt.Errorf("block %d: %w", i, err)
		}
	}
	return nil
}
