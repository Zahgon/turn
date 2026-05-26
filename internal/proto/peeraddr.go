// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package proto

import (
	"net"

	"github.com/pion/stun/v3"
)

// PeerAddress implements XOR-PEER-ADDRESS attribute.
//
// The XOR-PEER-ADDRESS specifies the address and port of the peer as
// seen from the TURN server. (For example, the peer's server-reflexive
// transport address if the peer is behind a NAT.)
//
// RFC 5766 Section 14.3.
type PeerAddress struct {
	IP   net.IP
	Port int
}

func (a PeerAddress) String() string { _ = "STUB: not implemented"; return "" }

// AddTo adds XOR-PEER-ADDRESS to message.
func (a PeerAddress) AddTo(m *stun.Message) error { _ = "STUB: not implemented"; return nil }

// GetFrom decodes XOR-PEER-ADDRESS from message.
func (a *PeerAddress) GetFrom(m *stun.Message) error { _ = "STUB: not implemented"; return nil }

// XORPeerAddress implements XOR-PEER-ADDRESS attribute.
//
// The XOR-PEER-ADDRESS specifies the address and port of the peer as
// seen from the TURN server. (For example, the peer's server-reflexive
// transport address if the peer is behind a NAT.)
//
// RFC 5766 Section 14.3.
type XORPeerAddress = PeerAddress
