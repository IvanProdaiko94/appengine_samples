package firestore

import (
	"cloud.google.com/go/firestore"
)

type TransactionDelta struct {
	tx *firestore.Transaction
	delta []func() error
}

func (writes *TransactionDelta) Create(doc *firestore.DocumentRef, d interface{}) *TransactionDelta {
	writes.delta = append(writes.delta, func() error {
		return writes.tx.Create(doc, d)
	})
	return writes
}

func (writes *TransactionDelta) Set(doc *firestore.DocumentRef, d interface{}, opts... firestore.SetOption) *TransactionDelta {
	writes.delta = append(writes.delta, func() error {
		return writes.tx.Set(doc, d, opts...)
	})
	return writes
}

func (writes *TransactionDelta) Update(doc *firestore.DocumentRef, u []firestore.Update, precondition... firestore.Precondition) *TransactionDelta {
	writes.delta = append(writes.delta, func() error {
		return writes.tx.Update(doc, u, precondition...)
	})
	return writes
}

func (writes *TransactionDelta) Delete(doc *firestore.DocumentRef, precondition... firestore.Precondition) *TransactionDelta {
	writes.delta = append(writes.delta, func() error {
		return writes.tx.Delete(doc, precondition...)
	})
	return writes
}

func (writes *TransactionDelta) Clear() *TransactionDelta {
	writes.delta = []func() error{}
	return writes
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
