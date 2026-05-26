// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package ipnet contains helper functions around net and IP
package ipnet

import (
	"errors"
	"net"
)

var errFailedToCastAddr = errors.New("failed to cast net.Addr to *net.UDPAddr or *net.TCPAddr")

// AddrIPPort extracts the IP and Port from a net.Addr.
func AddrIPPort(a net.Addr) (net.IP, int, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), 0, nil
}

// AddrEqual asserts that two net.Addrs are equal.
// Compares IP and Port and intentionally ignores IPv6 Zone.
func AddrEqual(addrA, addrB net.Addr) bool { _ = "STUB: not implemented"; return false }

// FingerprintAddr generates a fingerprint from net.UDPAddr or net.TCPAddr's
// which can be used for indexing maps.
func FingerprintAddr(addr net.Addr) string { _ = "STUB: not implemented"; return "" }

// Do we really need this case?

// Should never happen
