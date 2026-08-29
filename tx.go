package main

import "fmt"

type Tx struct {
	From      string
	To        string
	Amount    int
	Signature string
}

func (t Tx) payload() string {
	return fmt.Sprintf("%s:%s:%d", t.From, t.To, t.Amount)
}

func NewTx(from KeyPair, to string, amount int) Tx {
	t := Tx{
		From:   from.Address(),
		To:     to,
		Amount: amount,
	}
	t.Signature = from.Sign(t.payload())
	return t
}

func (t Tx) IsValid() error {
	if t.From == "" {
		return fmt.Errorf("missing from")
	}
	if !Verify(t.From, t.payload(), t.Signature) {
		return fmt.Errorf("bad signature")
	}
	return nil
}
