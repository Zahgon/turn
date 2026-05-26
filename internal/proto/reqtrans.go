// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package proto

import (
	"github.com/pion/stun/v3"
)

// Protocol is IANA assigned protocol number.
type Protocol byte

const (
	// ProtoTCP is IANA assigned protocol number for TCP.
	ProtoTCP Protocol = 6
	// ProtoUDP is IANA assigned protocol number for UDP.
	ProtoUDP Protocol = 17
)

func (p Protocol) String() string { _ = "STUB: not implemented"; return "" }

// RequestedTransport represents REQUESTED-TRANSPORT attribute.
//
// This attribute is used by the client to request a specific transport
// protocol for the allocated transport address. RFC 5766 only allows the use of
// code point 17 (User Datagram Protocol).
//
// RFC 5766 Section 14.7.
type RequestedTransport struct {
	Protocol Protocol
}

func (t RequestedTransport) String() string { _ = "STUB: not implemented"; return "" }

const requestedTransportSize = 4

// AddTo adds REQUESTED-TRANSPORT to message.
func (t RequestedTransport) AddTo(m *stun.Message) error { _ = "STUB: not implemented"; return nil }

// b[1:4] is RFFU = 0.
// The RFFU field MUST be set to zero on transmission and MUST be
// ignored on reception. It is reserved for future uses.

// GetFrom decodes REQUESTED-TRANSPORT from message.
func (t *RequestedTransport) GetFrom(m *stun.Message) error { _ = "STUB: not implemented"; return nil }
