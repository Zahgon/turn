// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package turn contains the public API for pion/turn, a toolkit for building TURN clients and servers
package turn

import (
	"crypto/tls"
	"net"
	"time"

	"github.com/pion/logging"
	"github.com/pion/turn/v5/internal/allocation"
	"github.com/pion/turn/v5/internal/server"
)

const (
	defaultInboundMTU = 1600
)

// Server is an instance of the Pion TURN Server.
type Server struct {
	log                logging.LeveledLogger
	authHandler        AuthHandler
	quotaHandler       QuotaHandler
	realm              string
	channelBindTimeout time.Duration
	permissionTimeout  time.Duration
	allocationLifetime time.Duration
	nonceHash          server.NonceManager
	eventHandler       EventHandler

	packetConnConfigs  []PacketConnConfig
	listenerConfigs    []ListenerConfig
	allocationManagers []*allocation.Manager
	inboundMTU         int
}

// NewServer creates the Pion TURN server.
func NewServer(config ServerConfig) (*Server, error) {
	_ = "STUB: not implemented" //nolint:gocognit,cyclop
	return nil, nil
}

// AllocationCount returns the number of active allocations.
// It can be used to drain the server before closing.
func (s *Server) AllocationCount() int { _ = "STUB: not implemented"; return 0 }

// Close stops the TURN Server.
// It cleans up any associated state and closes all connections it is managing.
func (s *Server) Close() error { _ = "STUB: not implemented"; return nil }

//nolint:errorlint

func (s *Server) readListener(l net.Listener, am *allocation.Manager) {
	_ = "STUB: not implemented"
	return
}

// Extract tls connection state if possible

// Force TLS handshake to complete before extracting connection state.
// Per crypto/tls docs: handshakes are lazy (complete on first I/O).
// We must call Handshake() explicitly to populate PeerCertificates.

// Delete allocation

// fixed UDP

type nilAddressGenerator struct{}

func (n *nilAddressGenerator) Validate() error { _ = "STUB: not implemented"; return nil }

func (n *nilAddressGenerator) AllocatePacketConn(AllocateListenerConfig) (net.PacketConn, net.Addr, error) {
	_ = "STUB: not implemented"
	return *new(net.PacketConn), *new(net.Addr), nil
}

func (n *nilAddressGenerator) AllocateListener(AllocateListenerConfig) (net.Listener, net.Addr, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), *new(net.Addr), nil
}

func (n *nilAddressGenerator) AllocateConn(AllocateConnConfig) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (s *Server) createAllocationManager(
	addrGenerator RelayAddressGenerator,
	handler PermissionHandler,
) (*allocation.Manager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) readLoop(conn net.PacketConn, allocationManager *allocation.Manager, tlsState *tls.ConnectionState) {
	_ = "STUB: not implemented"
	return
}
