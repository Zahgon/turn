// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package proto

import (
	"net"

	"github.com/pion/stun/v3"
)

// RelayedAddress implements XOR-RELAYED-ADDRESS attribute.
//
// It specifies the address and port that the server allocated to the
// client. It is encoded in the same way as XOR-MAPPED-ADDRESS.
//
// RFC 5766 Section 14.5.
type RelayedAddress struct {
	IP   net.IP
	Port int
}

func (a RelayedAddress) String() string { _ = "STUB: not implemented"; return "" }

// AddTo adds XOR-PEER-ADDRESS to message.
func (a RelayedAddress) AddTo(m *stun.Message) error { _ = "STUB: not implemented"; return nil }

// GetFrom decodes XOR-PEER-ADDRESS from message.
func (a *RelayedAddress) GetFrom(m *stun.Message) error { _ = "STUB: not implemented"; return nil }

// XORRelayedAddress implements XOR-RELAYED-ADDRESS attribute.
//
// It specifies the address and port that the server allocated to the
// client. It is encoded in the same way as XOR-MAPPED-ADDRESS.
//
// RFC 5766 Section 14.5.
type XORRelayedAddress = RelayedAddress
