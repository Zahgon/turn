// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package server

import (
	"net"

	"github.com/pion/stun/v3"
	"github.com/pion/turn/v5/internal/proto"
)

const runesAlpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// See: https://tools.ietf.org/html/rfc5766#section-6.2
// .
func handleAllocateRequest(req Request, stunMsg *stun.Message) error {
	_ = "STUB: not implemented" //nolint:cyclop,gocyclo,maintidx,gocognit
	return nil
}

// 1. The server MUST require that the request be authenticated.  This
//    authentication MUST be done using the long-term credential
//    mechanism of [https://tools.ietf.org/html/rfc5389#section-10.2.2]
//    unless the client and server agree to use another mechanism through
//    some procedure outside the scope of this document.

// 2. The server checks if the 5-tuple is currently in use by an
//    existing allocation.  If yes, the server rejects the request with
//    a 437 (Allocation Mismatch) error.

// A retry allocation

// 3. The server checks if the request contains a REQUESTED-TRANSPORT
//    attribute.  If the REQUESTED-TRANSPORT attribute is not included
//    or is malformed, the server rejects the request with a 400 (Bad
//    Request) error.  Otherwise, if the attribute is included but
//    specifies a protocol other that UDP/TCP, the server rejects the
//    request with a 442 (Unsupported Transport Protocol) error.

// 4. The request may contain a DONT-FRAGMENT attribute.  If it does,
//    but the server does not support sending UDP datagrams with the DF
//    bit set to 1 (see Section 12), then the server treats the DONT-
//    FRAGMENT attribute in the Allocate request as an unknown
//    comprehension-required attribute.

// 5.  The server checks if the request contains a RESERVATION-TOKEN
//     attribute.  If yes, and the request also contains an EVEN-PORT
//     attribute, then the server rejects the request with a 400 (Bad
//     Request) error.  Otherwise, it checks to see if the token is
//     valid (i.e., the token is in range and has not expired and the
//     corresponding relayed transport address is still available).  If
//     the token is not valid for some reason, the server rejects the
//     request with a 508 (Insufficient Capacity) error.

// 6. The server checks if the request contains an EVEN-PORT attribute.
//    If yes, then the server checks that it can satisfy the request
//    (i.e., can allocate a relayed transport address as described
//    below).  If the server cannot satisfy the request, then the
//    server rejects the request with a 508 (Insufficient Capacity)
//    error.

// Parse realm (already checked in authenticateRequest)

// RFC 6156: Parse REQUESTED-ADDRESS-FAMILY attribute if present.
// If absent, default to IPv4 per RFC 6156 Section 4.1.1.
// "If the REQUESTED-ADDRESS-FAMILY attribute is absent, the server MUST
// allocate an IPv4-relayed transport address for the TURN client."
// "Length:  this 16-bit field contains the length of the attribute in
// bytes.  The length of this attribute is 4 bytes."
// "If the server does not support the address family requested by the
// client, it MUST generate an Allocate error response, and it MUST
// include an ERROR-CODE attribute with the 440 (Address Family not
// Supported) response code, which is defined in Section 4.2.1."

// RFC 6156: Check if the requested address family is supported.
// If not, reject with 440 (Address Family not Supported) error.

// RFC 6156: REQUESTED-ADDRESS-FAMILY and RESERVATION-TOKEN are mutually exclusive.

// 7. At any point, the server MAY choose to reject the request with a
//    486 (Allocation Quota Reached) error if it feels the client is
//    trying to exceed some locally defined allocation quota.  The
//    server is free to define this allocation quota any way it wishes,
//    but SHOULD define it based on the username used to authenticate
//    the request, and not on the client's transport address.

// 8. Also at any point, the server MAY choose to reject the request
//    with a 300 (Try Alternate) error if it wishes to redirect the
//    client to a different server.  The use of this error code and
//    attribute follow the specification in [RFC5389].

// Once the allocation is created, the server replies with a success
// response.
// The success response contains:
//   * An XOR-RELAYED-ADDRESS attribute containing the relayed transport
//     address.
//   * A LIFETIME attribute containing the current value of the time-to-
//     expiry timer.
//   * A RESERVATION-TOKEN attribute (if a second relayed transport
//     address was reserved).
//   * An XOR-MAPPED-ADDRESS attribute containing the client's IP address
//     and port (from the 5-tuple).

func handleRefreshRequest(req Request, stunMsg *stun.Message) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:varnamelen

// RFC 6156: If REQUESTED-ADDRESS-FAMILY is present in Refresh request,
// it must match the allocation's address family.
// "The client MUST NOT include any REQUESTED-ADDRESS-FAMILY attribute in
// its Refresh Request."
// "If a server receives a Refresh Request with a REQUESTED-ADDRESS-
// FAMILY attribute, and the attribute's value doesn't match the address
// family of the allocation, the server MUST reply with a 443 (Peer
// Address Family Mismatch) Refresh error response."

func handleCreatePermissionRequest(req Request, stunMsg *stun.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// RFC 6156: Peer address must match allocation's address family.

func handleSendIndication(req Request, stunMsg *stun.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func handleChannelBindRequest(req Request, stunMsg *stun.Message) error {
	_ = "STUB: not implemented" // nolint:cyclop
	return nil
}

// RFC 6156: Peer address must match allocation's address family.

// nolint:nestif

func handleChannelData(req Request, channelData *proto.ChannelData) error {
	_ = "STUB: not implemented"
	return nil
}

func handleConnectRequest(req Request, stunMsg *stun.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// If the request does not contain an XOR-PEER-ADDRESS attribute, or if
// such attribute is invalid, the server MUST return a 400 (Bad Request)
// error.

// If the server is currently processing a Connect request for this
// allocation with the same XOR-PEER-ADDRESS, it MUST return a 446
// (Connection Already Exists) error.

// If the server has already successfully processed a Connect request
// for this allocation with the same XOR-PEER-ADDRESS, and the resulting
// client and peer data connections are either pending or active, it
// MUST return a 446 (Connection Already Exists) error.

// Otherwise, the server MUST initiate an outgoing TCP connection.  The
// local endpoint is the relayed transport address associated with the
// allocation.  The remote endpoint is the one indicated by the XOR-
// PEER-ADDRESS attribute.  If the connection attempt fails or times
// out, the server MUST return a 447 (Connection Timeout or Failure)
// error.  The timeout value MUST be at least 30 seconds.

// The server MUST include the CONNECTION-ID attribute in the Connect
// success response.  The attribute's value MUST uniquely identify the
// peer data connection.

func handleConnectionBindRequest(req Request, stunMsg *stun.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// Authentication of the client by the server MUST use the same method
// and credentials as for the control connection.
//
// GetTCPConnection asserts that userName used for auth is same as allocation

// When either side has failed close both

// ipMatchesFamily checks if an IP address matches the given address family.
func ipMatchesFamily(ip net.IP, family proto.RequestedAddressFamily) bool {
	_ = "STUB: not implemented"
	return false
}
