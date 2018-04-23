package firestore

import (
	"cloud.google.com/go/firestore"
)

type TransactionDelta struct {
	u map[*firestore.DocumentRef][]firestore.Update
	s map[*firestore.DocumentRef][]interface{}
	tx *firestore.Transaction
}

func (writes *TransactionDelta) Update(doc *firestore.DocumentRef, u firestore.Update) {
	if _, ok := writes.u[doc]; ok {
		writes.u[doc] = append(writes.u[doc], u)
	} else {
		writes.u[doc] = []firestore.Update{u}
	}
}

func (writes *TransactionDelta) Set(doc *firestore.DocumentRef, d interface{}) {
	if _, ok := writes.s[doc]; ok {
		writes.s[doc] = append(writes.s[doc], d)
	} else {
		writes.s[doc] = []interface{}{d}
	}
}

func (writes *TransactionDelta) Get(doc *firestore.DocumentRef) ([]firestore.Update, []interface{}) {
	return writes.u[doc], writes.s[doc]
}

func (writes *TransactionDelta) Apply() error {
	for ref, value := range writes.u {
		err := writes.tx.Update(ref, value)
		if err != nil {
			return err
		}
	}
	for ref, value := range writes.s {
		for _, newDoc := range value {
			err := writes.tx.Set(ref, newDoc)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func NewTransactionDelta(tx *firestore.Transaction) *TransactionDelta {
	return &TransactionDelta{
		u: make(map[*firestore.DocumentRef][]firestore.Update),
		s: make(map[*firestore.DocumentRef][]interface{}),
		tx: tx,
	}
}
