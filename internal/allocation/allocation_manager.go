// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package allocation

import (
	"net"
	"sync"
	"time"

	"github.com/pion/logging"
	"github.com/pion/turn/v5/internal/proto"
)

// If no ConnectionBind request associated with this peer data
// connection is received after 30 seconds, the peer data connection
// MUST be closed.
const defaultTCPConnectionBindTimeout = time.Second * 30

// AllocateListenerConfig contains the parameters passed to the relay address allocator
// when creating a new UDP or TCP allocation.
type AllocateListenerConfig struct {
	// Network specifies the network type for the allocation: "udp4", "udp6", "tcp4", or "tcp6".
	Network string
	// UserID is the authenticated user's identifier as returned by the AuthHandler.
	//
	// Note: The UserID is typcally the same as the TURN username, except for authentication
	// schemes that overload the username field with additional info (e.g., the lifetime of the
	// credential, as in the time-windowed credential mechanism in
	// https://datatracker.ietf.org/doc/html/draft-uberti-behave-turn-rest-00.
	UserID string
	// Realm is the TURN realm for this allocation.
	Realm string
	// RequestedPort is the port requested by the client in the TURN Allocate request.
	// A value of 0 indicates that the client did not request a specific port and any
	// available port may be used.
	RequestedPort int
}

// AllocateConnConfig contains the parameters passed to the relay address allocator
// when creating a new outbound TCP connection for RFC 6062 (TURN TCP) Connect requests.
type AllocateConnConfig struct {
	// Network specifies the network type for the connection: "tcp4" or "tcp6".
	Network string
	// UserID is the authenticated user's identifier as returned by the AuthHandler.
	//
	// Note: The UserID is typcally the same as the TURN username, except for authentication
	// schemes that overload the username field with additional info (e.g., the lifetime of the
	// credential, as in the time-windowed credential mechanism in
	// https://datatracker.ietf.org/doc/html/draft-uberti-behave-turn-rest-00.
	UserID string
	// Realm is the TURN realm for this allocation.
	Realm string
	// LocalAddr is the relay address to bind the local side of the connection
	// to. Implementations must allocate the local address as requested.
	LocalAddr net.Addr
	// RemoteAddr is the peer address to connect to.
	RemoteAddr net.Addr
}

// ManagerConfig a bag of config params for Manager.
type ManagerConfig struct {
	LeveledLogger      logging.LeveledLogger
	AllocatePacketConn func(info AllocateListenerConfig) (net.PacketConn, net.Addr, error)
	AllocateListener   func(info AllocateListenerConfig) (net.Listener, net.Addr, error)
	AllocateConn       func(info AllocateConnConfig) (net.Conn, error)
	PermissionHandler  func(sourceAddr net.Addr, peerIP net.IP) bool
	EventHandler       EventHandler

	tcpConnectionBindTimeout time.Duration
}

type reservation struct {
	token string
	port  int
}

// Manager is used to hold active allocations.
type Manager struct {
	lock                     sync.RWMutex
	log                      logging.LeveledLogger
	tcpConnectionBindTimeout time.Duration

	allocations  map[FiveTupleFingerprint]*Allocation
	reservations []*reservation

	allocatePacketConn func(conf AllocateListenerConfig) (net.PacketConn, net.Addr, error)
	allocateListener   func(conf AllocateListenerConfig) (net.Listener, net.Addr, error)
	allocateConn       func(conf AllocateConnConfig) (net.Conn, error)
	permissionHandler  func(sourceAddr net.Addr, peerIP net.IP) bool
	EventHandler       EventHandler
}

// NewManager creates a new instance of Manager.
func NewManager(config ManagerConfig) (*Manager, error) { _ = "STUB: not implemented"; return nil, nil }

// GetAllocation fetches the allocation matching the passed FiveTuple.
func (m *Manager) GetAllocation(fiveTuple *FiveTuple) *Allocation {
	_ = "STUB: not implemented"
	return nil
}

// GetAllocationForUserID fetches the allocation matching the passed FiveTuple and Username.
func (m *Manager) GetAllocationForUserID(fiveTuple *FiveTuple, userID string) *Allocation {
	_ = "STUB: not implemented"
	return nil
}

// AllocationCount returns the number of existing allocations.
func (m *Manager) AllocationCount() int { _ = "STUB: not implemented"; return 0 }

// Close closes the manager and closes all allocations it manages.
func (m *Manager) Close() error { _ = "STUB: not implemented"; return nil }

// CreateAllocation creates a new allocation and starts relaying.
func (m *Manager) CreateAllocation( // nolint: cyclop
	fiveTuple *FiveTuple,
	turnSocket net.PacketConn,
	protocol proto.Protocol,
	requestedPort int,
	lifetime time.Duration,
	userID, realm string,
	addressFamily proto.RequestedAddressFamily,
) (*Allocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Only start the UDP relay loop for UDP allocations.

// For TCP allocations, accept inbound connections on the relayed listener and notify the client.

// DeleteAllocation removes an allocation.
func (m *Manager) DeleteAllocation(fiveTuple *FiveTuple) { _ = "STUB: not implemented"; return }

// CreateReservation stores the reservation for the token+port.
func (m *Manager) CreateReservation(reservationToken string, port int) {
	_ = "STUB: not implemented"
	return
}

// GetReservation returns the port for a given reservation if it exists.
func (m *Manager) GetReservation(reservationToken string) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// GetRandomEvenPort returns a random un-allocated udp4 port.
func (m *Manager) GetRandomEvenPort() (int, error) { _ = "STUB: not implemented"; return 0, nil }

// GrantPermission handles permission requests by calling the permission handler callback
// associated with the TURN server listener socket.
func (m *Manager) GrantPermission(sourceAddr net.Addr, peerIP net.IP) error {
	_ = "STUB: not implemented"
	// No permission handler: open
	return nil
}

// CreateTCPConnection creates a new outbound TCP Connection and returns the Connection-ID
// if it succeeds.
func (m *Manager) CreateTCPConnection( // nolint: cyclop
	allocation *Allocation,
	peerAddress proto.PeerAddress,
) (proto.ConnectionID, error) {
	_ = "STUB: not implemented"
	return *new(proto.ConnectionID), nil
}

// RFC 6156:
// "After the request has been successfully authenticated, the TURN
// server allocates a transport address of the type indicated in the
// REQUESTED-ADDRESS-FAMILY attribute."

// nolint: noctx

func (m *Manager) addTCPConnection(allocation *Allocation, conn net.Conn) (proto.ConnectionID, error) {
	_ = "STUB: not implemented"
	return *new(proto.ConnectionID), nil
}

// nolint: gosec

func (m *Manager) isDupeTCPConnection(allocation *Allocation, remoteAddr *net.TCPAddr) bool {
	_ = "STUB: not implemented"
	return false
}

// GetTCPConnection returns the TCP Connection for the given ConnectionID.
func (m *Manager) GetTCPConnection(userID string, connectionID proto.ConnectionID) net.Conn {
	_ = "STUB: not implemented"
	return *new(net.Conn)
}

func (m *Manager) RemoveTCPConnection(connectionID proto.ConnectionID) {
	_ = "STUB: not implemented"
	return
}
