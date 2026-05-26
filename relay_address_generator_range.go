// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package turn

import (
	"net"

	"github.com/pion/randutil"
	"github.com/pion/transport/v4"
)

// RelayAddressGeneratorPortRange can be used to only allocate connections inside a defined port range.
// Similar to the RelayAddressGeneratorStatic a static ip address can be set.
type RelayAddressGeneratorPortRange struct {
	// RelayAddress is the IP returned to the user when the relay is created
	RelayAddress net.IP

	// MinPort the minimum port to allocate
	MinPort uint16
	// MaxPort the maximum (inclusive) port to allocate
	MaxPort uint16

	// MaxRetries the amount of tries to allocate a random port in the defined range
	MaxRetries int

	// Rand the random source of numbers
	Rand randutil.MathRandomGenerator

	// Address is passed to Listen/ListenPacket when creating the Relay
	Address string

	Net transport.Net
}

// Validate is called on server startup and confirms the RelayAddressGenerator is properly configured.
func (r *RelayAddressGeneratorPortRange) Validate() error { _ = "STUB: not implemented"; return nil }

// AllocatePacketConn generates a new PacketConn to receive traffic on and the IP/Port
// to populate the allocation response with.
func (r *RelayAddressGeneratorPortRange) AllocatePacketConn(
	conf AllocateListenerConfig,
) (net.PacketConn, net.Addr, error) {
	_ = "STUB: not implemented"
	return *new(net.PacketConn), *new(net.Addr), nil
}

// nolint: noctx

// nolint:gosec // G115 false positive

// nolint: noctx

// AllocateListener generates a new Listener to receive traffic on and the IP/Port
// to populate the allocation response with.
func (r *RelayAddressGeneratorPortRange) AllocateListener( // nolint: cyclop
	conf AllocateListenerConfig,
) (net.Listener, net.Addr, error) {
	_ = "STUB: not implemented"
	// AllocateListener can be called independently of Validate (e.g. in tests),
	// so ensure we're initialized to avoid nil dereferences.
	return *new(net.Listener), *new(net.Addr), nil
}

// Enable SO_REUSEADDR and SO_REUSEPORT where needed to let multiple connnections
// bind to the same relay address.

// nolint:gosec // G115 false positive

// AllocateConn creates a new outgoing TCP connection bound to the relay address to send traffic to a peer.
func (r *RelayAddressGeneratorPortRange) AllocateConn(conf AllocateConnConfig) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// Enable SO_REUSEADDR and SO_REUSEPORT where needed to let multiple connnections
// bind to the same relay address.
