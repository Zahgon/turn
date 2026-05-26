// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package allocation

import (
	"net"
)

// Protocol is an enum for relay protocol.
type Protocol uint8

// Network protocols for relay.
const (
	UDP Protocol = iota
	TCP
)

func (p Protocol) String() string { _ = "STUB: not implemented"; return "" }

// FiveTuple is the combination (client IP address and port, server IP
// address and port, and transport protocol (currently one of UDP,
// TCP, or TLS)) used to communicate between the client and the
// server.  The 5-tuple uniquely identifies this communication
// stream.  The 5-tuple also uniquely identifies the Allocation on
// the server.
type FiveTuple struct {
	Protocol
	SrcAddr, DstAddr net.Addr
}

// Equal asserts if two FiveTuples are equal.
func (f *FiveTuple) Equal(b *FiveTuple) bool { _ = "STUB: not implemented"; return false }

// FiveTupleFingerprint is a comparable representation of a FiveTuple.
type FiveTupleFingerprint struct {
	srcIP, dstIP     [16]byte
	srcPort, dstPort uint16
	protocol         Protocol
}

// Fingerprint is the identity of a FiveTuple.
func (f *FiveTuple) Fingerprint() (fp FiveTupleFingerprint) {
	_ = "STUB: not implemented"
	return *new(FiveTupleFingerprint)
}

func netAddrIPAndPort(addr net.Addr) (net.IP, uint16) {
	_ = "STUB: not implemented"
	return *new(net.IP), 0
}

// nolint:gosec // G115

// nolint:gosec // G115
