// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package proto

import (
	"time"

	"github.com/pion/stun/v3"
)

// DefaultLifetime in RFC 5766 is 10 minutes.
//
// RFC 5766 Section 2.2.
const DefaultLifetime = time.Minute * 10

// Lifetime represents LIFETIME attribute.
//
// The LIFETIME attribute represents the duration for which the server
// will maintain an allocation in the absence of a refresh. The value
// portion of this attribute is 4-bytes long and consists of a 32-bit
// unsigned integral value representing the number of seconds remaining
// until expiration.
//
// RFC 5766 Section 14.2.
type Lifetime struct {
	time.Duration
}

// Seconds in uint32.
const lifetimeSize = 4 // 4 bytes, 32 bits

// AddTo adds LIFETIME to message.
func (l Lifetime) AddTo(m *stun.Message) error { _ = "STUB: not implemented"; return nil }

// GetFrom decodes LIFETIME from message.
func (l *Lifetime) GetFrom(m *stun.Message) error { _ = "STUB: not implemented"; return nil }

// Asserting length
