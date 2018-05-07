package firestore

import (
	"cloud.google.com/go/firestore"
)

type TransactionDelta struct {
	tx *firestore.Transaction
	delta []func() error
}

func (writes *TransactionDelta) Create(doc *firestore.DocumentRef, d interface{}) {
	writes.delta = append(writes.delta, func() error {
		return writes.tx.Create(doc, d)
	})
}

func (writes *TransactionDelta) Set(doc *firestore.DocumentRef, d interface{}) {
	writes.delta = append(writes.delta, func() error {
		return writes.tx.Set(doc, d)
	})
}

func (writes *TransactionDelta) Update(doc *firestore.DocumentRef, u firestore.Update) {
	writes.delta = append(writes.delta, func() error {
		return writes.tx.Update(doc, []firestore.Update{u})
	})
}

func (writes *TransactionDelta) Delete(doc *firestore.DocumentRef) {
	writes.delta = append(writes.delta, func() error {
		return writes.tx.Delete(doc)
	})
}

func (writes *TransactionDelta) Apply() error {
	for _, fn := range writes.delta {
		err := fn()
		if err != nil {
			return err
		}
	}
	return nil
}

func NewTransactionDelta(tx *firestore.Transaction) *TransactionDelta {
	return &TransactionDelta{
		tx: tx,
		delta: []func() error{},
	}
}
