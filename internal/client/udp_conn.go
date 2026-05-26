// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package client implements the API for a TURN client
package client

import (
	"net"
	"time"

	"github.com/pion/turn/v5/internal/proto"
)

const (
	maxReadQueueSize              = 1024
	defaultPermRefreshInterval    = 120 * time.Second
	defaultBindingRefreshInterval = 5 * time.Minute
	defaultBindingCheckInterval   = 30 * time.Second
	maxRetryAttempts              = 3
)

const (
	timerIDRefreshAlloc int = iota
	timerIDRefreshPerms
	timerIDCheckBindings
)

type inboundData struct {
	data []byte
	from net.Addr
}

// UDPConn is the implementation of the Conn and PacketConn interfaces for UDP network connections.
// compatible with net.PacketConn and net.Conn.
type UDPConn struct {
	bindingMgr             *bindingManager   // Thread-safe
	checkBindingsTimer     *PeriodicTimer    // Thread-safe
	readCh                 chan *inboundData // Thread-safe
	closeCh                chan struct{}     // Thread-safe
	bindingRefreshInterval time.Duration     // Read-only
	allocation
}

// NewUDPConn creates a new instance of UDPConn.
func NewUDPConn(config *AllocationConfig) *UDPConn { _ = "STUB: not implemented"; return nil }

// ReadFrom reads a packet from the connection,
// copying the payload into p. It returns the number of
// bytes copied into p and the return address that
// was on the packet.
// It returns the number of bytes read (0 <= n <= len(p))
// and any error encountered. Callers should always process
// the n > 0 bytes returned before considering the error err.
// ReadFrom can be made to time out and return
// an Error with Timeout() == true after a fixed time limit;
// see SetDeadline and SetReadDeadline.
func (c *UDPConn) ReadFrom(p []byte) (n int, addr net.Addr, err error) {
	_ = "STUB: not implemented"
	return 0, *new(net.Addr), nil
}

func (a *allocation) createPermission(perm *permission, addr net.Addr) error {
	_ = "STUB: not implemented"
	return nil
}

// Punch a hole! (this would block a bit..)

// WriteTo writes a packet with payload to addr.
// WriteTo can be made to time out and return
// an Error with Timeout() == true after a fixed time limit;
// see SetDeadline and SetWriteDeadline.
// On packet-oriented connections, write timeouts are rare.
func (c *UDPConn) WriteTo(payload []byte, addr net.Addr) (int, error) {
	_ = "STUB: not implemented" //nolint:gocognit,cyclop
	return 0, nil
}

// Check if we have a permission for the destination IP addr

// c.createPermission() would block, per destination IP (, or perm),
// until the perm state becomes "requested". Purpose of this is to
// guarantee the order of packets (within the same perm).
// Note that CreatePermission transaction may not be complete before
// all the data transmission. This is done assuming that the request
// will be most likely successful and we can tolerate some loss of
// UDP packet (or reorder), inorder to minimize the latency in most cases.

// Bind channel

//nolint:nestif

// Try to establish an initial binding with the server.
// Writes still occur via indications meanwhile.

// Send data using SendIndication

// Binding is ready beyond this point, so send over it.

// Close closes the connection.
// Any blocked ReadFrom or WriteTo operations will be unblocked and return errors.
func (c *UDPConn) Close() error { _ = "STUB: not implemented"; return nil }

/* dontWait=true */

// LocalAddr returns the local network address.
func (c *UDPConn) LocalAddr() net.Addr {
	_ = "STUB: not implemented"
	return *

	// SetDeadline sets the read and write deadlines associated
	// with the connection. It is equivalent to calling both
	// SetReadDeadline and SetWriteDeadline.
	//
	// A deadline is an absolute time after which I/O operations
	// fail with a timeout (see type Error) instead of
	// blocking. The deadline applies to all future and pending
	// I/O, not just the immediately following call to ReadFrom or
	// WriteTo. After a deadline has been exceeded, the connection
	// can be refreshed by setting a deadline in the future.
	//
	// An idle timeout can be implemented by repeatedly extending
	// the deadline after successful ReadFrom or WriteTo calls.
	//
	// A zero value for t means I/O operations will not time out.
	new(net.Addr)
}

func (c *UDPConn) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// SetReadDeadline sets the deadline for future ReadFrom calls
// and any currently-blocked ReadFrom call.
// A zero value for t means ReadFrom will not time out.
func (c *UDPConn) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// SetWriteDeadline sets the deadline for future WriteTo calls
// and any currently-blocked WriteTo call.
// Even if write times out, it may return n > 0, indicating that
// some of the data was successfully written.
// A zero value for t means WriteTo will not time out.
func (c *UDPConn) SetWriteDeadline(time.Time) error {
	_ = "STUB: not implemented"
	// Write never blocks.
	return nil
}

func addr2PeerAddress(addr net.Addr) proto.PeerAddress {
	_ = "STUB: not implemented"
	return *new(proto.PeerAddress)
}

// CreatePermissions Issues a CreatePermission request for the supplied addresses
// as described in https://datatracker.ietf.org/doc/html/rfc5766#section-9
func (a *allocation) CreatePermissions(addrs ...net.Addr) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint // dynamic errors

// HandleInbound passes inbound data in UDPConn.
func (c *UDPConn) HandleInbound(data []byte, from net.Addr) {
	_ = "STUB: not implemented"
	// Copy data
	return
}

// FindAddrByChannelNumber returns a peer address associated with the
// channel number on this UDPConn.
func (c *UDPConn) FindAddrByChannelNumber(chNum uint16) (net.Addr, bool) {
	_ = "STUB: not implemented"
	return *new(net.Addr), false
}

func (c *UDPConn) maybeBind(bound *binding) { _ = "STUB: not implemented"; return }

// Block only callers with the same binding until
// the binding transaction has been complete

// Establish binding with the server if eligible
// with regard to cases right above.

func (c *UDPConn) bind(bound *binding) error { _ = "STUB: not implemented"; return nil }

// nolint:err113

// nolint:err113

// Success.

func (c *UDPConn) sendChannelData(data []byte, chNum uint16) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
