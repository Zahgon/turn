// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package turn

import (
	"net"

	"github.com/pion/transport/v4"
)

// RelayAddressGeneratorNone returns the listener with no modifications.
type RelayAddressGeneratorNone struct {
	// Address is passed to Listen/ListenPacket when creating the Relay
	Address string

	Net transport.Net
}

// Validate is called on server startup and confirms the RelayAddressGenerator is properly configured.
func (r *RelayAddressGeneratorNone) Validate() error { _ = "STUB: not implemented"; return nil }

// AllocatePacketConn generates a new PacketConn to receive traffic on and the IP/Port
// to populate the allocation response with.
func (r *RelayAddressGeneratorNone) AllocatePacketConn(conf AllocateListenerConfig) (
	net.PacketConn,
	net.Addr,
	error,
) {
	_ = "STUB: not implemented"
	return *new(net.PacketConn), *new(net.Addr), nil
}

// nolint: noctx

// AllocateListener generates a new Listener to receive traffic on and the IP/Port
// to populate the allocation response with.
func (r *RelayAddressGeneratorNone) AllocateListener(conf AllocateListenerConfig) (net.Listener, net.Addr, error) {
	_ = "STUB: not implemented"
	// AllocateListener can be called independently of Validate (e.g. in tests),
	// so ensure we're initialized to avoid nil dereferences.
	return *new(net.Listener), *new(net.Addr), nil
}

// Enable SO_REUSEADDR and SO_REUSEPORT where needed to let multiple connnections
// bind to the same relay address.

// AllocateConn creates a new outgoing TCP connection bound to the relay address to send traffic to a peer.
func (r *RelayAddressGeneratorNone) AllocateConn(conf AllocateConnConfig) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// Enable SO_REUSEADDR and SO_REUSEPORT where needed to let multiple connnections
// bind to the same relay address.
