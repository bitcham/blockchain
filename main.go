package main

import "fmt"

func printChain(c Chain) {
	for _, b := range c {
		fmt.Printf("  #%d  prev=%s\n", b.Index, short(b.PrevHash))
		if b.Index == 0 {
			fmt.Printf("       (genesis)\n")
		} else {
			fmt.Printf("       tx=%s -> %s amount=%d\n", short(b.Tx.From), short(b.Tx.To), b.Tx.Amount)
			fmt.Printf("       sig=%s\n", short(b.Tx.Signature))
		}
		fmt.Printf("       hash=%s\n", short(b.Hash))
	}
}

func short(h string) string {
	if h == "" {
		return "(empty)"
	}
	if len(h) < 12 {
		return h
	}
	return h[:12] + "..."
}

func report(label string, c Chain) {
	fmt.Println(label)
	printChain(c)
	if err := c.IsValid(); err != nil {
		fmt.Println("  valid? NO —", err)
	} else {
		fmt.Println("  valid? YES")
	}
	fmt.Println()
}

func main() {
	alice, err := NewKeyPair()
	if err != nil {
		panic(err)
	}
	bob, err := NewKeyPair()
	if err != nil {
		panic(err)
	}
	mallory, err := NewKeyPair()
	if err != nil {
		panic(err)
	}

	chain := NewChain()
	chain = chain.Add(NewTx(alice, bob.Address(), 10))
	chain = chain.Add(NewTx(bob, alice.Address(), 4))

	report("honest signed chain", chain)

	edited := make(Chain, len(chain))
	copy(edited, chain)
	edited[1].Tx.To = mallory.Address()
	edited[1].Hash = edited[1].computeHash()
	report("change destination, recompute block hash, keep Alice's signature", edited)

	forged := make(Chain, len(chain))
	copy(forged, chain)
	forged[1].Tx.To = mallory.Address()
	forged[1].Tx.Signature = mallory.Sign(forged[1].Tx.payload())
	forged[1].Hash = forged[1].computeHash()
	report("same edit, but Mallory signs while From is still Alice", forged)
}
