// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package allocation contains all CRUD operations for allocations
package allocation

import (
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/pion/logging"
	"github.com/pion/stun/v3"
	"github.com/pion/turn/v5/internal/proto"
)

type allocationResponse struct {
	transactionID [stun.TransactionIDSize]byte
	responseAttrs []stun.Setter
}

type tcpConnection struct {
	net.Conn

	isBound   atomic.Bool // ConnectionBind has been done for this TCP Connection
	bindTimer *time.Timer
}

// Allocation is tied to a FiveTuple and relays traffic
// use CreateAllocation and GetAllocation to operate.
type Allocation struct {
	RelayAddr           net.Addr
	Protocol            Protocol
	TurnSocket          net.PacketConn
	fiveTuple           *FiveTuple
	permissionsLock     sync.RWMutex
	permissions         map[string]*Permission
	channelBindingsLock sync.RWMutex
	channelBindings     []*ChannelBind
	lifetimeTimer       *time.Timer
	closed              chan any
	userID, realm       string
	eventHandler        EventHandler
	log                 logging.LeveledLogger
	addressFamily       proto.RequestedAddressFamily // RFC 6156

	// Relay Transport for UDP
	relayPacketConn net.PacketConn

	// Relay Transport for TCP
	relayListener net.Listener

	tcpConnections map[proto.ConnectionID]*tcpConnection // Guarded by AllocationManager lock

	// Some clients (Firefox or others using resiprocate's nICE lib) may retry allocation
	// with same 5 tuple when received 413, for compatible with these clients,
	// cache for response lost and client retry to implement 'stateless stack approach'
	// See: https://datatracker.ietf.org/doc/html/rfc5766#section-6.2
	responseCache atomic.Value // *allocationResponse
}

// NewAllocation creates a new instance of NewAllocation.
func NewAllocation(
	turnSocket net.PacketConn,
	fiveTuple *FiveTuple,
	eventHandler EventHandler,
	log logging.LeveledLogger,
) *Allocation {
	_ = "STUB: not implemented"
	return nil
}

// GetPermission gets the Permission from the allocation.
func (a *Allocation) GetPermission(addr net.Addr) *Permission {
	_ = "STUB: not implemented"
	return nil
}

// AddPermission adds a new permission to the allocation.
func (a *Allocation) AddPermission(perms *Permission) { _ = "STUB: not implemented"; return }

// RemovePermission removes the net.Addr's fingerprint from the allocation's permissions.
func (a *Allocation) RemovePermission(addr net.Addr) { _ = "STUB: not implemented"; return }

// ListPermissions returns the permissions associated with an allocation.
func (a *Allocation) ListPermissions() []*Permission { _ = "STUB: not implemented"; return nil }

// AddChannelBind adds a new ChannelBind to the allocation, it also updates the
// permissions needed for this ChannelBind.
func (a *Allocation) AddChannelBind(chanBind *ChannelBind, channelLifetime, permissionLifetime time.Duration) error {
	_ = "STUB: not implemented"
	// Check that this channel id isn't bound to another transport address, and
	// that this transport address isn't bound to another channel number.
	return nil
}

// Peer already bound to a different channel number.

// Channel number already bound to a different peer.

// Add or refresh this channel.

// Channel binds also refresh permissions.

// Channel binds also refresh permissions.

// RemoveChannelBind removes the ChannelBind from this allocation by id.
func (a *Allocation) RemoveChannelBind(number proto.ChannelNumber) bool {
	_ = "STUB: not implemented"
	return false
}

// GetChannelByNumber gets the ChannelBind from this allocation by id.
func (a *Allocation) GetChannelByNumber(number proto.ChannelNumber) *ChannelBind {
	_ = "STUB: not implemented"
	return nil
}

// GetChannelByAddr gets the ChannelBind from this allocation by net.Addr.
func (a *Allocation) GetChannelByAddr(addr net.Addr) *ChannelBind {
	_ = "STUB: not implemented"
	return nil
}

// ListChannelBindings returns the channel bindings associated with an allocation.
func (a *Allocation) ListChannelBindings() []*ChannelBind { _ = "STUB: not implemented"; return nil }

// Refresh updates the allocations lifetime.
func (a *Allocation) Refresh(lifetime time.Duration) { _ = "STUB: not implemented"; return }

// AddressFamily returns the address family of the allocation (RFC 6156).
func (a *Allocation) AddressFamily() proto.RequestedAddressFamily {
	_ = "STUB: not implemented"
	return *

	// SetResponseCache cache allocation response for retransmit allocation request.
	new(proto.RequestedAddressFamily)
}

func (a *Allocation) SetResponseCache(transactionID [stun.TransactionIDSize]byte, attrs []stun.Setter) {
	_ = "STUB: not implemented"
	return
}

// GetResponseCache return response cache for retransmit allocation request.
func (a *Allocation) GetResponseCache() (id [stun.TransactionIDSize]byte, attrs []stun.Setter) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *Allocation) removeTCPConnection(connectionID proto.ConnectionID) {
	_ = "STUB: not implemented"
	return
}

// RemoveTCPConnection closes and removes the TCP Connection.
func (a *Allocation) RemoveTCPConnection(m *Manager, connectionID proto.ConnectionID) {
	_ = "STUB: not implemented"
	return
}

// Close closes the allocation.
func (a *Allocation) Close() error { _ = "STUB: not implemented"; return nil }

// WriteTo writes a packet with payload p to addr via the Relay socket.
func (a *Allocation) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//  https://tools.ietf.org/html/rfc5766#section-10.3
//  When the server receives a UDP datagram at a currently allocated
//  relayed transport address, the server looks up the allocation
//  associated with the relayed transport address.  The server then
//  checks to see whether the set of permissions for the allocation allow
//  the relaying of the UDP datagram as described in Section 8.
//
//  If relaying is permitted, then the server checks if there is a
//  channel bound to the peer that sent the UDP datagram (see
//  Section 11).  If a channel is bound, then processing proceeds as
//  described in Section 11.7.
//
//  If relaying is permitted but no channel is bound to the peer, then
//  the server forms and sends a Data indication.  The Data indication
//  MUST contain both an XOR-PEER-ADDRESS and a DATA attribute.  The DATA
//  attribute is set to the value of the 'data octets' field from the
//  datagram, and the XOR-PEER-ADDRESS attribute is set to the source
//  transport address of the received UDP datagram.  The Data indication
//  is then sent on the 5-tuple associated with the allocation.

const rtpMTU = 1600

func (a *Allocation) packetConnHandler(manager *Manager) { _ = "STUB: not implemented"; return }

// nolint:nestif

func (a *Allocation) connHandler(manager *Manager) { _ = "STUB: not implemented"; return }
