// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package proto

import (
	"errors"
	"net"
	"time"
)

var (
	errInvalidTURNFrame    = errors.New("data is not a valid TURN frame, no STUN or ChannelData found")
	errIncompleteTURNFrame = errors.New("data contains incomplete STUN or TURN frame")
)

// STUNConn wraps a net.Conn and implements
// net.PacketConn by being STUN aware and
// packetizing the stream.
type STUNConn struct {
	nextConn net.Conn
	buff     []byte
}

const (
	stunHeaderSize     = 20
	channelDataPadding = 4
)

// Given a buffer give the last offset of the TURN frame
// If the buffer isn't a valid STUN or ChannelData packet,
// or the length doesn't match return false.
func consumeSingleTURNFrame(b []byte) (int, error) {
	_ = "STUB: not implemented"
	// Too short to determine if ChannelData or STUN
	return 0, nil
}

// ReadFrom implements ReadFrom from net.PacketConn.
func (s *STUNConn) ReadFrom(payload []byte) (n int, addr net.Addr, err error) {
	_ = "STUB: not implemented"
	// First pass any buffered data from previous reads
	return 0, *new(net.Addr), nil
}

// Then read from the nextConn, appending to our buff

// WriteTo implements WriteTo from net.PacketConn.
func (s *STUNConn) WriteTo(payload []byte, _ net.Addr) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Close implements Close from net.PacketConn.
func (s *STUNConn) Close() error { _ = "STUB: not implemented"; return nil }

// LocalAddr implements LocalAddr from net.PacketConn.
func (s *STUNConn) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

// SetDeadline implements SetDeadline from net.PacketConn.
func (s *STUNConn) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// SetReadDeadline implements SetReadDeadline from net.PacketConn.
func (s *STUNConn) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// SetWriteDeadline implements SetWriteDeadline from net.PacketConn.
func (s *STUNConn) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// Conn returns the net.Conn used for this STUNConn.
func (s *STUNConn) Conn() net.Conn {
	_ = "STUB: not implemented"

	// NewSTUNConn creates a STUNConn.
	return *new(net.Conn)
}

func NewSTUNConn(nextConn net.Conn) *STUNConn { _ = "STUB: not implemented"; return nil }
