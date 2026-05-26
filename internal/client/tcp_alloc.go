// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package client

import (
	"net"
	"time"

	"github.com/pion/transport/v4"
	"github.com/pion/turn/v5/internal/proto"
)

var (
	_ transport.TCPListener = (*TCPAllocation)(nil) // Includes type check for net.Listener
	_ transport.Dialer      = (*TCPAllocation)(nil)
)

func noDeadline() time.Time {
	_ = "STUB: not implemented"

	// TCPAllocation is an active TCP allocation on the TURN server
	// as specified by RFC 6062.
	// The allocation can be used to Dial/Accept relayed outgoing/incoming TCP connections.
	return *new(time.Time)
}

type TCPAllocation struct {
	connAttemptCh chan *connectionAttempt
	acceptTimer   *time.Timer
	allocation
}

// NewTCPAllocation creates a new instance of TCPConn.
func NewTCPAllocation(config *AllocationConfig) *TCPAllocation {
	_ = "STUB: not implemented"
	return nil
}

// Connect sends a Connect request to the turn server and returns a chosen connection ID.
func (a *TCPAllocation) Connect(peer net.Addr) (proto.ConnectionID, error) {
	_ = "STUB: not implemented"
	return *new(proto.ConnectionID), nil
}

//nolint // dynamic errors

//nolint // dynamic errors

// Dial connects to the address on the named network.
func (a *TCPAllocation) Dial(network, rAddrStr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// DialWithConn connects to the address on the named network with an already existing connection.
// The provided connection must be an already connected TCP connection to the TURN server.
func (a *TCPAllocation) DialWithConn(conn net.Conn, network, rAddrStr string) (*TCPConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RFC 6156:
// "TURN can run over UDP and TCP, and it allows for a client to request
// address/port pairs for receiving both UDP and TCP."
// "This document adds IPv6 support to TURN, which includes IPv4-to-IPv6,
// IPv6-to-IPv6, and IPv6-to-IPv4 relaying.".
func (a *TCPAllocation) serverTCPAddr() (*net.TCPAddr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DialTCP acts like Dial for TCP networks.
func (a *TCPAllocation) DialTCP(network string, lAddr, rAddr *net.TCPAddr) (*TCPConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck,gosec

// DialTCPWithConn acts like DialWithConn for TCP networks.
func (a *TCPAllocation) DialTCPWithConn(conn net.Conn, _ string, rAddr *net.TCPAddr) (*TCPConn, error) {
	_ = "STUB: not implemented"

	// Check if we have a permission for the destination IP addr
	return nil, nil
}

// Send connect request if haven't done so.

// BindConnection associates the provided connection.
func (a *TCPAllocation) BindConnection(dataConn *TCPConn, cid proto.ConnectionID) error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}

// Read exactly one STUN message, any data after belongs to the user

//nolint // dynamic errors

//nolint // dynamic errors

// Accept waits for and returns the next connection to the listener.
func (a *TCPAllocation) Accept() (net.Conn, error) {
	_ = "STUB: not implemented"
	return *

	// AcceptTCP accepts the next incoming call and returns the new connection.
	new(net.Conn), nil
}

func (a *TCPAllocation) AcceptTCP() (transport.TCPConn, error) {
	_ = "STUB: not implemented"
	return *new(transport.TCPConn), nil
}

//nolint:errcheck,gosec

// AcceptTCPWithConn accepts the next incoming call and returns the new connection.
func (a *TCPAllocation) AcceptTCPWithConn(conn net.Conn) (*TCPConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetDeadline sets the deadline associated with the listener. A zero time value disables the deadline.
func (a *TCPAllocation) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// Close releases the allocation
// Any blocked Accept operations will be unblocked and return errors.
// Any opened connection via Dial/Accept will be closed.
func (a *TCPAllocation) Close() error { _ = "STUB: not implemented"; return nil }

/* dontWait=true */

// Addr returns the relayed address of the allocation.
func (a *TCPAllocation) Addr() net.Addr {
	_ = "STUB: not implemented"
	return *

	// HandleConnectionAttempt is called by the TURN client
	// when it receives a ConnectionAttempt indication.
	new(net.Addr)
}

func (a *TCPAllocation) HandleConnectionAttempt(from *net.TCPAddr, cid proto.ConnectionID) {
	_ = "STUB: not implemented"
	return
}
