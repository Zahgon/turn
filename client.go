// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package turn

import (
	"math"
	"net"
	"sync"
	"time"

	"github.com/pion/logging"
	"github.com/pion/stun/v3"
	"github.com/pion/transport/v4"
	"github.com/pion/turn/v5/internal/client"
	"github.com/pion/turn/v5/internal/proto"
)

const (
	defaultRTO        = 200 * time.Millisecond
	maxRtxCount       = 7              // Total 7 requests (Rc)
	maxDataBufferSize = math.MaxUint16 // Message size limit for Chromium
)

//              interval [msec]
// 0: 0 ms      +500
// 1: 500 ms	+1000
// 2: 1500 ms   +2000
// 3: 3500 ms   +4000
// 4: 7500 ms   +8000
// 5: 15500 ms  +16000
// 6: 31500 ms  +32000
// -: 63500 ms  failed

// ClientConfig is a bag of config parameters for Client.
type ClientConfig struct {
	STUNServerAddr string // STUN server address (e.g. "stun.abc.com:3478")
	TURNServerAddr string // TURN server address (e.g. "turn.abc.com:3478")
	Username       string
	Password       string //nolint:gosec // runtime credential, not hardcoded.
	Realm          string
	Software       string
	RTO            time.Duration
	Conn           net.PacketConn // Listening socket (net.PacketConn)
	Net            transport.Net
	LoggerFactory  logging.LoggerFactory

	// PermissionTimeout sets the refresh interval of permissions. Defaults to 2 minutes.
	PermissionRefreshInterval time.Duration

	// RequestedAddressFamily is the address family to request in allocations (IPv4 or IPv6).
	// If not specified (zero value), the client will attempt to infer from the PacketConn's
	// local address, falling back to IPv4 if inference fails. See RFC 6156.
	RequestedAddressFamily RequestedAddressFamily

	evenPort               bool   // If EVEN-PORT Attribute should be sent in Allocation
	reservationToken       []byte // If Server responds with RESERVATION-TOKEN or if Client wishes to send one
	bindingRefreshInterval time.Duration
	bindingCheckInterval   time.Duration
}

// Client is a STUN server client.
type Client struct {
	conn           net.PacketConn // Read-only
	net            transport.Net  // Read-only
	stunServerAddr net.Addr       // Read-only
	turnServerAddr net.Addr       // Read-only

	username      stun.Username          // Read-only
	password      string                 // Read-only
	realm         stun.Realm             // Read-only
	integrity     stun.MessageIntegrity  // Read-only
	software      stun.Software          // Read-only
	trMap         *client.TransactionMap // Thread-safe
	rto           time.Duration          // Read-only
	relayedConn   *client.UDPConn        // Protected by mutex ***
	tcpAllocation *client.TCPAllocation  // Protected by mutex ***
	allocTryLock  client.TryLock         // Thread-safe
	listenTryLock client.TryLock         // Thread-safe
	mutex         sync.RWMutex           // Thread-safe
	mutexTrMap    sync.Mutex             // Thread-safe
	log           logging.LeveledLogger  // Read-only

	// If EVEN-PORT Attribute should be sent in Allocation
	evenPort bool

	// If Server responds with RESERVATION-TOKEN or if Client wishes to send one
	reservationToken []byte

	// REQUESTED-ADDRESS-FAMILY attribute for allocations (RFC 6156)
	requestedAddressFamily proto.RequestedAddressFamily

	permissionRefreshInterval time.Duration
	bindingRefreshInterval    time.Duration
	bindingCheckInterval      time.Duration
}

// inferAddressFamilyFromConn attempts to determine the address
// family (IPv4 or IPv6) from a PacketConn's local address.
// Returns an error if the address type is not IP-based.
func inferAddressFamilyFromConn(
	conn net.PacketConn,
) (proto.RequestedAddressFamily, error) {
	_ = "STUB: not implemented"
	return *new(proto.RequestedAddressFamily), nil
}

//nolint:err113

// getRequestedAddressFamily determines the address family to use
// for TURN allocations. It follows this priority:
//  1. Use explicitly configured RequestedAddressFamily if set
//  2. Try to infer from the PacketConn's local address
//  3. Fall back to IPv4 default per RFC 6156
func getRequestedAddressFamily(
	log logging.LeveledLogger,
	config *ClientConfig,
) proto.RequestedAddressFamily {
	_ = "STUB: not implemented"
	// If explicitly set, use it
	return *new(proto.RequestedAddressFamily)
}

// Try to infer from the PacketConn

// Default to IPv4 per RFC 6156

// appendRequestedAddressFamilyOrReservation adds either RESERVATION-TOKEN or
// REQUESTED-ADDRESS-FAMILY to the provided setters slice, respecting mutual
// exclusivity rules from RFC 6156.
// The REQUESTED-ADDRESS-FAMILY attribute is
// only included when IPv6 is desired.
func appendRequestedAddressFamilyOrReservation(
	setters []stun.Setter,
	requestedFamily proto.RequestedAddressFamily,
	reservationToken []byte,
) []stun.Setter {
	_ = "STUB: not implemented"
	// Clients MUST NOT include a REQUESTED-ADDRESS-FAMILY attribute in an
	// Allocate request that contains a RESERVATION-TOKEN attribute.
	// https://www.rfc-editor.org/rfc/rfc6156#section-4.1
	return nil
}

// Only include the attribute when IPv6 is explicitly requested.
// This indirectly implied by the specification:
// If the REQUESTED-ADDRESS-FAMILY attribute is absent, the server MUST
// allocate an IPv4-relayed transport address for the TURN client.
// https://www.rfc-editor.org/rfc/rfc6156#section-4.2

// NewClient returns a new Client instance. listeningAddress is the address and port to listen on,
// default "0.0.0.0:0".
func NewClient(config *ClientConfig) (*Client, error) {
	_ = "STUB: not implemented" //nolint:gocyclo,cyclop
	return nil, nil
}

// Determine the requested address family (RFC 6156)

// TURNServerAddr return the TURN server address.
func (c *Client) TURNServerAddr() net.Addr {
	_ = "STUB: not implemented"
	return *

	// STUNServerAddr return the STUN server address.
	new(net.Addr)
}

func (c *Client) STUNServerAddr() net.Addr {
	_ = "STUB: not implemented"
	return *

	// Username returns username.
	new(net.Addr)
}

func (c *Client) Username() stun.Username {
	_ = "STUB: not implemented"

	// Realm return realm.
	return *new(stun.Username)
}

func (c *Client) Realm() stun.Realm {
	_ = "STUB: not implemented"

	// WriteTo sends data to the specified destination using the base socket.
	return *new(stun.Realm)
}

func (c *Client) WriteTo(data []byte, to net.Addr) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Listen will have this client start listening on the conn provided via the config.
// This is optional. If not used, you will need to call HandleInbound method
// to supply incoming data, instead.
func (c *Client) Listen() error { _ = "STUB: not implemented"; return nil }

// Close closes this client.
func (c *Client) Close() { _ = "STUB: not implemented"; return }

// TransactionID & Base64: https://play.golang.org/p/EEgmJDI971P

// SendBindingRequestTo sends a new STUN request to the given transport address.
func (c *Client) SendBindingRequestTo(to net.Addr) (net.Addr, error) {
	_ = "STUB: not implemented"
	return *new(net.Addr), nil
}

// SendBindingRequest sends a new STUN request to the STUN server.
func (c *Client) SendBindingRequest() (net.Addr, error) {
	_ = "STUB: not implemented"
	return *new(net.Addr), nil
}

func (c *Client) sendAllocateRequest(protocol proto.Protocol) ( //nolint:cyclop
	relayed proto.RelayedAddress,
	lifetime proto.Lifetime,
	nonce stun.Nonce,
	reservationToken proto.ReservationToken,
	err error,
) {
	_ = "STUB: not implemented"
	return *new(proto.RelayedAddress), *new(proto.Lifetime), *new(stun.Nonce), *new(proto.ReservationToken), nil
}

// FINGERPRINT must be the last attribute per RFC 5389

// Anonymous allocate failed, trying to authenticate.

// Trying to authorize.

// FINGERPRINT must be the last attribute per RFC 5389

//nolint:err113

// Getting relayed addresses from response.

// Getting lifetime from response

// Getting reservation-token from response

// Allocate sends a TURN allocation request to the given transport address.
func (c *Client) Allocate() (net.PacketConn, error) {
	_ = "STUB: not implemented"
	return *new(net.PacketConn), nil
}

// AllocateTCP creates a new TCP allocation at the TURN server.
func (c *Client) AllocateTCP() (*client.TCPAllocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreatePermission Issues a CreatePermission request for the supplied addresses
// as described in https://datatracker.ietf.org/doc/html/rfc5766#section-9
func (c *Client) CreatePermission(addrs ...net.Addr) error { _ = "STUB: not implemented"; return nil }

// PerformTransaction performs STUN transaction.
func (c *Client) PerformTransaction(msg *stun.Message, to net.Addr, ignoreResult bool) (client.TransactionResult,
	error,
) {
	_ = "STUB: not implemented"
	return *new(client.TransactionResult), nil
}

// If ignoreResult is true, get the transaction going and return immediately

// OnDeallocated is called when de-allocation of relay address has been complete.
// (Called by UDPConn).
func (c *Client) OnDeallocated(net.Addr) { _ = "STUB: not implemented"; return }

// HandleInbound handles data received.
// This method handles incoming packet de-multiplex it by the source address
// and the types of the message.
// This return a boolean (handled or not) and if there was an error.
// Caller should check if the packet was handled by this client or not.
// If not handled, it is assumed that the packet is application data.
// If an error is returned, the caller should discard the packet regardless.
func (c *Client) HandleInbound(data []byte, from net.Addr) (bool, error) {
	_ = "STUB: not implemented"
	// +-------------------+-------------------------------+
	// |   Return Values   |                               |
	// +-------------------+       Meaning / Action        |
	// | handled |  error  |                               |
	// |=========+=========+===============================+
	// |  false  |   nil   | Handle the packet as app data |
	// |---------+---------+-------------------------------+
	// |  true   |   nil   |        Nothing to do          |
	// |---------+---------+-------------------------------+
	// |  false  |  error  |     (shouldn't happen)        |
	// |---------+---------+-------------------------------+
	// |  true   |  error  | Error occurred while handling |
	// +---------+---------+-------------------------------+
	// Possible causes of the error:
	//  - Malformed packet (parse error)
	//  - STUN message was a request
	//  - Non-STUN message from the STUN server
	return false, nil
}

// Received from STUN server but it is not a STUN message

// Assume, this is an application data

func (c *Client) handleSTUNMessage(data []byte, from net.Addr) error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}

// nolint:nestif

// Silently discard

// Silently discard

// This is a STUN response message (transactional)
// The type is either:
// - stun.ClassSuccessResponse
// - stun.ClassErrorResponse

// Silently discard

// End the transaction

func (c *Client) handleChannelData(data []byte) error { _ = "STUB: not implemented"; return nil }

// Silently discard

func (c *Client) onRtxTimeout(trKey string, nRtx int) { _ = "STUB: not implemented"; return }

// Already gone

// All retransmissions failed

func (c *Client) setRelayedUDPConn(conn *client.UDPConn) { _ = "STUB: not implemented"; return }

func (c *Client) relayedUDPConn() *client.UDPConn { _ = "STUB: not implemented"; return nil }

func (c *Client) setTCPAllocation(alloc *client.TCPAllocation) { _ = "STUB: not implemented"; return }

func (c *Client) getTCPAllocation() *client.TCPAllocation { _ = "STUB: not implemented"; return nil }

func (c *Client) setReservationToken(reservationToken []byte) { _ = "STUB: not implemented"; return }

func (c *Client) getReservationToken() []byte { _ = "STUB: not implemented"; return nil }
