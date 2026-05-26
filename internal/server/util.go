// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package server

import (
	"net"
	"time"

	"github.com/pion/stun/v3"
)

const (
	// See: https://tools.ietf.org/html/rfc5766#section-6.2 defines 3600 seconds recommendation.
	maximumAllocationLifetime = time.Hour
)

func buildAndSend(conn net.PacketConn, dst net.Addr, attrs ...stun.Setter) error {
	_ = "STUB: not implemented"
	return nil
}

// Send a STUN packet and return the original error to the caller.
func buildAndSendErr(conn net.PacketConn, dst net.Addr, err error, attrs ...stun.Setter) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errorlint

func buildMsg(
	transactionID [stun.TransactionIDSize]byte,
	msgType stun.MessageType,
	additional ...stun.Setter,
) []stun.Setter {
	_ = "STUB: not implemented"
	return nil
}

func authenticateRequest(req Request, stunMsg *stun.Message, callingMethod stun.Method) (
	messageIntegrity stun.MessageIntegrity,
	hasAuth bool,
	username string,
	err error,
) {
	_ = "STUB: not implemented"
	return *new(stun.MessageIntegrity), false, "", nil
}

// No Auth handler is set, server is running in STUN only mode
// Respond with 400 so clients don't retry.

// Assert Nonce is signed and is not expired.

func genAuthEvent(req Request, stunMsg *stun.Message, callingMethod stun.Method, verdict bool) {
	_ = "STUB: not implemented"
	return
}

// Auth event is generated per the username, not the user-id.

func allocationLifeTime(req Request, m *stun.Message) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
