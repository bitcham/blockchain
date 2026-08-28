package main

import "fmt"

func printChain(c Chain) {
	for _, b := range c {
		fmt.Printf("  #%d  prev=%s\n", b.Index, short(b.PrevHash))
		fmt.Printf("       data=%q\n", b.Data)
		fmt.Printf("       hash=%s\n", short(b.Hash))
	}
}

func short(h string) string {
	if h == "" {
		return "(genesis)"
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
	chain := NewChain("genesis")
	chain = chain.Add("alice sends 10 to bob")
	chain = chain.Add("bob sends 4 to carol")
	// 3 chain blocks

	report("honest chain", chain)

	tampered := make(Chain, len(chain))
	copy(tampered, chain)
	tampered[1].Data = "alice sends 11 to bob"
	report("tamper block 1 data, leave hashes", tampered)

	tampered[1].Hash = tampered[1].computeHash()
	report("same tamper, but recompute only block 1 hash", tampered)
}
