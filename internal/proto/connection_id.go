// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package proto

import (
	"github.com/pion/stun/v3"
)

// ConnectionID represents CONNECTION-ID attribute.
//
// The CONNECTION-ID attribute uniquely identifies a peer data
// connection.  It is a 32-bit unsigned integral value.
//
// RFC 6062 Section 6.2.1.
type ConnectionID uint32

const connectionIDSize = 4 // uint32: 4 bytes, 32 bits

// AddTo adds CONNECTION-ID to message.
func (c ConnectionID) AddTo(m *stun.Message) error { _ = "STUB: not implemented"; return nil }

// GetFrom decodes CONNECTION-ID from message.
func (c *ConnectionID) GetFrom(m *stun.Message) error { _ = "STUB: not implemented"; return nil }

// Asserting length
