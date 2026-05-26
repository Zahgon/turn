// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package server

import (
	"time"
)

const (
	nonceLifetime  = time.Hour // See: https://tools.ietf.org/html/rfc5766#section-4
	nonceLength    = 40
	nonceKeyLength = 64
)

// NewNonceHash creates a NonceHash.
func NewNonceHash() (NonceManager, error) {
	_ = "STUB: not implemented"
	return *new(NonceManager), nil
}

// NonceHash is used to create and verify nonces.
type NonceHash struct {
	key []byte
}

// Generate a nonce.
func (n *NonceHash) Generate() (string, error) { _ = "STUB: not implemented"; return "", nil }

// nolint:gosec // G115

//nolint:errorlint

// Validate checks that nonce is signed and is not expired.
func (n *NonceHash) Validate(nonce string) error { _ = "STUB: not implemented"; return nil }

//nolint:errorlint

// nolint:gosec // G115

//nolint:errorlint
