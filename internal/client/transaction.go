// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package client

import (
	"net"
	"sync"
	"time"

	"github.com/pion/stun/v3"
)

const (
	maxRtxInterval time.Duration = 1600 * time.Millisecond
)

// TransactionResult is a bag of result values of a transaction.
type TransactionResult struct {
	Msg     *stun.Message
	From    net.Addr
	Retries int
	Err     error
}

// TransactionConfig is a set of config params used by NewTransaction.
type TransactionConfig struct {
	Key          string
	Raw          []byte
	To           net.Addr
	Interval     time.Duration
	IgnoreResult bool // True to throw away the result of this transaction (it will not be readable using WaitForResult)
}

// Transaction represents a transaction.
type Transaction struct {
	Key      string                 // Read-only
	Raw      []byte                 // Read-only
	To       net.Addr               // Read-only
	nRtx     int                    // Modified only by the timer thread
	interval time.Duration          // Modified only by the timer thread
	timer    *time.Timer            // Thread-safe, set only by the creator, and stopper
	resultCh chan TransactionResult // Thread-safe
	mutex    sync.RWMutex
}

// NewTransaction creates a new instance of Transaction.
func NewTransaction(config *TransactionConfig) *Transaction { _ = "STUB: not implemented"; return nil }

// Read-only
// Read-only
// Read-only
// Modified only by the timer thread
// Thread-safe

// StartRtxTimer starts the transaction timer.
func (t *Transaction) StartRtxTimer(onTimeout func(trKey string, nRtx int)) {
	_ = "STUB: not implemented"
	return
}

// StopRtxTimer stop the transaction timer.
func (t *Transaction) StopRtxTimer() { _ = "STUB: not implemented"; return }

// WriteResult writes the result to the result channel.
func (t *Transaction) WriteResult(res TransactionResult) bool {
	_ = "STUB: not implemented"
	return false
}

// WaitForResult waits for the transaction result.
func (t *Transaction) WaitForResult() TransactionResult {
	_ = "STUB: not implemented"
	return *new(TransactionResult)
}

// Close closes the transaction.
func (t *Transaction) Close() { _ = "STUB: not implemented"; return }

// Retries returns the number of retransmission it has made.
func (t *Transaction) Retries() int { _ = "STUB: not implemented"; return 0 }

// TransactionMap is a thread-safe transaction map.
type TransactionMap struct {
	trMap map[string]*Transaction
	mutex sync.RWMutex
}

// NewTransactionMap create a new instance of the transaction map.
func NewTransactionMap() *TransactionMap { _ = "STUB: not implemented"; return nil }

// Insert inserts a transaction to the map.
func (m *TransactionMap) Insert(key string, tr *Transaction) bool {
	_ = "STUB: not implemented"
	return false
}

// Find looks up a transaction by its key.
func (m *TransactionMap) Find(key string) (*Transaction, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Delete deletes a transaction by its key.
func (m *TransactionMap) Delete(key string) { _ = "STUB: not implemented"; return }

// CloseAndDeleteAll closes and deletes all transactions.
func (m *TransactionMap) CloseAndDeleteAll() { _ = "STUB: not implemented"; return }

// Size returns the length of the transaction map.
func (m *TransactionMap) Size() int { _ = "STUB: not implemented"; return 0 }
