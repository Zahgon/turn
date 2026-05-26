// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package server

import (
	"time"
)

// NonceManager interface that both implementations satisfy.
type NonceManager interface {
	Generate() (string, error)
	Validate(nonce string) error
}

const (
	shortNonceLifetime     = time.Hour // Same as original
	shortNonceKeyLength    = 64        // Same as original
	shortNonceTimestampLen = 4         // 6 bytes for timestamp (minutes) - optimal size
	shortNonceMinHMACLen   = 2         // Minimum HMAC length for security
	shortNonceMaxHMACLen   = 32        // Maximum HMAC length (full SHA256)
	defaultNonceHMACLen    = 12        // Default HMAC length
)

// NewShortNonceHash creates a ShortNonceHash. The hmacLen argument specifies the number of HMAC
// bytes to include (2-32 bytes).  The total nonce size will be 4 + hmacLen bytes, default hmaclen
// is 12 bytes. The 4 bytes timestamp gives about ~8000 years before nonces would start to repeat
// (safe until year 10,135).
func NewShortNonceHash(hmacLen int) (NonceManager, error) {
	_ = "STUB: not implemented"
	return *new(NonceManager), nil
}

// ShortNonceHash is used to create and verify short nonces.
type ShortNonceHash struct {
	key     []byte
	hmacLen int
}

// Generate a short nonce (4 + hmacLen bytes encoded as base36).
func (s *ShortNonceHash) Generate() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Convert to bytes and trim to 4 bytes.  This safely handles the conversion since we know
// current values fit in 4 bytes until year 10,135.

// nolint:gosec // G115

// Validate checks that nonce is signed and is not expired.
func (s *ShortNonceHash) Validate(nonce string) error { _ = "STUB: not implemented"; return nil }

// Pad with leadnign zeros if leading zeros were stripped during encoding/decoding.

// Check if nonce is expired (older than 1 hour).

// Recompute HMAC and compare.
