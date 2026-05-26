// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package proto

import (
	"net"
)

// Addr is ip:port.
type Addr struct {
	IP   net.IP
	Port int
}

// Network implements net.Addr.
func (Addr) Network() string {
	_ = "STUB: not implemented"

	// FromUDPAddr sets addr to UDPAddr.
	return ""
}

func (a *Addr) FromUDPAddr(n *net.UDPAddr) { _ = "STUB: not implemented"; return }

// Equal returns true if b == a.
func (a Addr) Equal(b Addr) bool { _ = "STUB: not implemented"; return false }

// EqualIP returns true if a and b have equal IP addresses.
func (a Addr) EqualIP(b Addr) bool { _ = "STUB: not implemented"; return false }

func (a Addr) String() string { _ = "STUB: not implemented"; return "" }

// FiveTuple represents 5-TUPLE value.
type FiveTuple struct {
	Client Addr
	Server Addr
	Proto  Protocol
}

func (t FiveTuple) String() string { _ = "STUB: not implemented"; return "" }

// Equal returns true if b == t.
func (t FiveTuple) Equal(b FiveTuple) bool { _ = "STUB: not implemented"; return false }
